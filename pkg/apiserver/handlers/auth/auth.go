package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/golang/glog"
	hfv1 "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"
	clientset "github.com/hobbyfarm/gargantua/pkg/client/clientset/versioned"
	"github.com/labstack/echo/v4"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

type Server struct {
	hfClientSet *clientset.Clientset
}

func NewServer(hfClientset *clientset.Clientset) *Server {
	server := &Server{
		hfClientSet: hfClientset,
	}

	return server
}

func (s *Server) MiddlwareDispatch(ctx echo.Context, provider string, scopes []string) error {
	switch provider {
	case "BearerAuth":
		return s.HandleBearerAuth(ctx, scopes)
	}

	return nil
}

func retrieveTokenString(ctx echo.Context) string {
	authString := ctx.Request().Header.Get("Authorization")
	if len(authString) == 0 {
		return ""
	}

	authSplit := strings.Split(authString, "Bearer ")
	return authSplit[1]
}

func (s *Server) getUserByEmail(email string) (hfv1.User, error) {
	if len(email) == 0 {
		return hfv1.User{}, fmt.Errorf("email passed in was empty")
	}

	users, err := s.hfClientSet.HobbyfarmV1().Users().List(metav1.ListOptions{})

	if err != nil {
		return hfv1.User{}, fmt.Errorf("error while retrieving user list")
	}

	for _, user := range users.Items {
		if user.Spec.Email == email {
			return user, nil
		}
	}

	return hfv1.User{}, fmt.Errorf("user not found")
}

func (s *Server) HandleBearerAuth(ctx echo.Context, scopes []string) error {
	// we're going to abuse the scopes functionality to provide
	// "roles" that a user should have to perform a certain action

	// this normally should be handled by the Echo JWT middleware (https://echo.labstack.com/middleware/jwt)
	tokenString := retrieveTokenString(ctx)

	var user hfv1.User

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		var user hfv1.User
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			var err error
			user, err = s.getUserByEmail(fmt.Sprint(claims["email"]))
			if err != nil {
				glog.Errorf("could not find user that matched token %s", fmt.Sprint(claims["email"]))
				return hfv1.User{}, fmt.Errorf("could not find user that matched token %s", fmt.Sprint(claims["email"]))
			}
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(user.Spec.Password), nil
	})

	if err != nil {
		glog.Errorf("error while validating user: %v", err)
		return err
	}

	if _, ok := token.Claims.(jwt.Claims); !ok || token.Valid {
		return fmt.Errorf("unable to validate user")
	}

	ctx.Set("user", user)

	for _, scope := range scopes {
		if scope == "admin" {
			if user.Spec.Admin {
				return nil
			} else {
				return fmt.Errorf("not authorized")
			}
		}
	}

	/*
	naive approach for potential rbac usage
	requires user.spec.roles to be a list with roles

	for _, scope := range scopes {
		for _, role := range user.Spec.Roles {
			if scope == role {
				return nil
			}
		}
	}
	 */

	return nil // user would have failed scope & role checks before getting here, so this is for auth-only endpoints
}