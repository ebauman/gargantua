package rpc

import (
	"context"
	"crypto/sha256"
	"encoding/base32"
	"fmt"
	"github.com/golang/glog"
	hfv1 "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"
	"github.com/hobbyfarm/gargantua/pkg/auth"
	hfClientset "github.com/hobbyfarm/gargantua/pkg/client/clientset/versioned"
	"github.com/hobbyfarm/gargantua/pkg/protobuf"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/util/retry"
	"strings"
)

type AuthServer struct {
	authClient *auth.AuthClient
	hfClientset *hfClientset.Clientset
}

func setupAuthServer(g *grpc.Server, hfClientset *hfClientset.Clientset, authClient *auth.AuthClient) {
	a := AuthServer{}
	a.authClient = authClient
	a.hfClientset = hfClientset

	protobuf.RegisterAuthServiceServer(g, a)
}

func (a AuthServer) RegisterWithAccessCode(ctx context.Context, registration *protobuf.Registration) (*protobuf.RegistrationResponse, error) {
	var accessCodeSuccessful bool
	userId, err := a.newUser(registration.Email, registration.Password)

	if err != nil {
		return nil, err
	}

	accessCode := strings.ToLower(registration.AccessCode)
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		user, err := a.hfClientset.HobbyfarmV1().Users().Get(userId, metav1.GetOptions{})
		if err != nil {
			return fmt.Errorf("error retrieving user")
		}

		if len(user.Spec.AccessCodes) == 0 {
			user.Spec.AccessCodes = []string{}
		} else {
			for _, ac := range user.Spec.AccessCodes {
				if ac == accessCode {
					return fmt.Errorf("access code already added to user")
				}
			}
		}

		user.Spec.AccessCodes = append(user.Spec.AccessCodes, accessCode)
		_, updateErr := a.hfClientset.HobbyfarmV1().Users().Update(user)
		return updateErr
	})

	if retryErr != nil {
		accessCodeSuccessful = false
	}

	glog.V(2).Infof("created user %s", registration.Email)
	return &protobuf.RegistrationResponse{AccessCodeSuccessful: accessCodeSuccessful}, nil
}

func (a AuthServer) ListAccessCodes(ctx context.Context, _ *protobuf.Empty) (*protobuf.UserAccessCodes, error) {
	user, err := getUser(ctx)
	if err != nil {
		return nil, err
	}

	latestUser, err := a.hfClientset.HobbyfarmV1().Users().Get(user, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return &protobuf.UserAccessCodes{AccessCodes: latestUser.Spec.AccessCodes}, nil
}

func (a AuthServer) UpdateAccessCodes(ctx context.Context, codes *protobuf.UserAccessCodes) (*protobuf.UserAccessCodes, error) {
	user, err := getUser(ctx)
	if err != nil {
		return nil, err
	}

	latestUser, err := a.hfClientset.HobbyfarmV1().Users().Get(user, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	latestUser.Spec.AccessCodes = codes.AccessCodes

	_, err = a.hfClientset.HobbyfarmV1().Users().Update(latestUser)
	if err != nil {
		return nil, err
	}

	return &protobuf.UserAccessCodes{AccessCodes: latestUser.Spec.AccessCodes}, nil
}

func (a AuthServer) ChangePassword(ctx context.Context, request *protobuf.ChangePasswordRequest) (*protobuf.Status, error) {
	user, err := getUser(ctx)
	if err != nil {
		return nil, err
	}

	latestUser, err := a.hfClientset.HobbyfarmV1().Users().Get(user, metav1.GetOptions{})
	if err != nil {
		glog.Errorf("error retrieving user: %v", err)
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(latestUser.Spec.Password), []byte(request.OldPassword))
	if err != nil {
		glog.Errorf("old password incorrect for user %s: %v", user, err)
		return nil, status.Error(codes.InvalidArgument, "bad password")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		glog.Errorf("error hashing password for user %s: %v", user, err)
		return nil, status.Error(codes.Internal, "error creating new password")
	}

	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		latestUser.Spec.Password = string(passwordHash)
		_, updateErr := a.hfClientset.HobbyfarmV1().Users().Update(latestUser)
		return updateErr
	})

	if retryErr != nil {
		return nil, retryErr
	}

	return &protobuf.Status{Status: true}, nil
}

func (a AuthServer) Authenticate(ctx context.Context, authentication *protobuf.Authentication) (*protobuf.Token, error) {
	user, err := a.authClient.GetUserByEmail(authentication.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Spec.Password), []byte(authentication.Password))
	if err != nil {
		glog.Errorf("password incorrect for user %s: %v", user.Name, err)
		return nil, status.Error(codes.Unauthenticated, "authentication failed")
	}

	token, err := a.authClient.GenerateJWT(user)
	if err != nil {
		return nil, err
	}

	return &protobuf.Token{Token: token}, nil
}

func (a AuthServer) newUser(email string, password string) (string, error) {
	// check if user exists
	users, err := a.hfClientset.HobbyfarmV1().Users().List(metav1.ListOptions{})
	if err != nil {
		return "", status.Errorf(codes.Internal, "error while retrieving user list: %s", err)
	}

	for _, user := range users.Items {
		if user.Spec.Email == email {
			return "", status.Errorf(codes.AlreadyExists, "user already exists with email %s", email)
		}
	}

	// at this point, user does not exist. so, create one!
	newUser := hfv1.User{}

	hasher := sha256.New()
	hasher.Write([]byte(email))
	sha := base32.StdEncoding.WithPadding(-1).EncodeToString(hasher.Sum(nil))[:10]
	id := "u-" + strings.ToLower(sha)
	newUser.Name = id
	newUser.Spec.Id = id
	newUser.Spec.Email = email

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", status.Errorf(codes.Internal,"error while hashing password for email %s", email)
	}

	newUser.Spec.Password = string(passwordHash)

	_, err = a.hfClientset.HobbyfarmV1().Users().Create(&newUser)
	if err != nil {
		glog.Errorf("error creating user: %s", err)
		return "", status.Errorf(codes.Internal, "error creating user")
	}

	return id, nil
}

