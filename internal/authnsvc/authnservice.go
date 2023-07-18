package authnservice

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/golang/glog"
	"github.com/gorilla/mux"
	hfv2 "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v2"
	"github.com/hobbyfarm/gargantua/pkg/microservices"
	"github.com/hobbyfarm/gargantua/pkg/util"
	"github.com/hobbyfarm/gargantua/protos/authn"
	rbacProto "github.com/hobbyfarm/gargantua/protos/rbac"
	userProto "github.com/hobbyfarm/gargantua/protos/user"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PreparedUser struct {
	ID string `json:"id"`
	hfv2.UserSpec
}

func (a AuthServer) ChangePasswordFunc(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	user, err := a.internalAuthnServer.AuthN(r.Context(), &authn.AuthNRequest{
		Token: token,
	})
	if err != nil {
		util.ReturnHTTPMessage(w, r, 403, "forbidden", "no access to change password")
		return
	}

	r.ParseForm()

	oldPassword := r.PostFormValue("old_password")
	newPassword := r.PostFormValue("new_password")

	err = a.ChangePassword(user, oldPassword, newPassword, r.Context())

	if err != nil {
		util.ReturnHTTPMessage(w, r, 500, "error", fmt.Sprintf("error changing password for user %s", user.Id))
		return
	}

	util.ReturnHTTPMessage(w, r, 200, "success", "password changed")

	glog.V(2).Infof("changed password for user %s", user.Email)
}

func (a AuthServer) UpdateSettingsFunc(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	user, err := a.internalAuthnServer.AuthN(r.Context(), &authn.AuthNRequest{
		Token: token,
	})
	if err != nil {
		util.ReturnHTTPMessage(w, r, 403, "forbidden", "no access to update settings")
		return
	}

	r.ParseForm()

	newSettings := make(map[string]string)
	for key := range r.Form {
		newSettings[key] = r.FormValue(key) //Ignore when multiple values were set for one argument. Just take the first one
	}

	err = a.UpdateSettings(user, newSettings, r.Context())

	if err != nil {
		util.ReturnHTTPMessage(w, r, 500, "error", fmt.Sprintf("error updating settings for user %s", user.Id))
		return
	}

	util.ReturnHTTPMessage(w, r, 200, "success", "settings updated")

	glog.V(2).Infof("settings updated for user %s", user.Email)
}

func (a AuthServer) ListAccessCodeFunc(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	user, err := a.internalAuthnServer.AuthN(r.Context(), &authn.AuthNRequest{
		Token: token,
	})
	if err != nil {
		util.ReturnHTTPMessage(w, r, 403, "forbidden", "no access to get accesscode")
		return
	}

	encodedACList, err := json.Marshal(user.GetAccessCodes())
	if err != nil {
		glog.Error(err)
	}
	util.ReturnHTTPContent(w, r, 200, "success", encodedACList)

	glog.V(2).Infof("retrieved accesscode list for user %s", user.GetEmail())
}

func (a AuthServer) RetreiveSettingsFunc(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	user, err := a.internalAuthnServer.AuthN(r.Context(), &authn.AuthNRequest{
		Token: token,
	})
	if err != nil {
		util.ReturnHTTPMessage(w, r, 403, "forbidden", "no access to get settings")
		return
	}

	encodedSettings, err := json.Marshal(user.GetSettings())

	if err != nil {
		glog.Error(err)
	}
	util.ReturnHTTPContent(w, r, 200, "success", encodedSettings)

	glog.V(2).Infof("retrieved settings list for user %s", user.Email)
}

func (a AuthServer) AddAccessCodeFunc(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	user, err := a.internalAuthnServer.AuthN(r.Context(), &authn.AuthNRequest{
		Token: token,
	})
	if err != nil {
		util.ReturnHTTPMessage(w, r, 403, "forbidden", "no access to get accesscode")
		return
	}

	r.ParseForm()

	accessCode := strings.ToLower(r.PostFormValue("access_code"))

	err = a.AddAccessCode(user, accessCode, r.Context())

	if err != nil {
		glog.Error(err)
		util.ReturnHTTPMessage(w, r, 500, "error", "error adding access code")
		return
	}

	util.ReturnHTTPMessage(w, r, 200, "success", accessCode)

	glog.V(2).Infof("added accesscode %s to user %s", accessCode, user.Email)
}

func (a AuthServer) RemoveAccessCodeFunc(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	user, err := a.internalAuthnServer.AuthN(r.Context(), &authn.AuthNRequest{
		Token: token,
	})
	if err != nil {
		util.ReturnHTTPMessage(w, r, 403, "forbidden", "no access to get accesscode")
		return
	}

	vars := mux.Vars(r)

	accessCode := strings.ToLower(vars["access_code"])

	err = a.RemoveAccessCode(user, accessCode, r.Context())

	if err != nil {
		glog.Error(err)
		util.ReturnHTTPMessage(w, r, 500, "error", "error removing access code")
		return
	}

	util.ReturnHTTPMessage(w, r, 200, "success", accessCode)

	glog.V(2).Infof("removed accesscode %s to user %s", accessCode, user.Email)
}

func (a AuthServer) AddAccessCode(user *userProto.User, accessCode string, ctx context.Context) error {
	if len(user.GetId()) == 0 || len(accessCode) == 0 {
		return fmt.Errorf("bad parameters passed, %s:%s", user.GetId(), accessCode)
	}

	accessCode = strings.ToLower(accessCode)

	if len(user.GetAccessCodes()) == 0 {
		user.AccessCodes = []string{}
	} else {
		for _, ac := range user.GetAccessCodes() {
			if ac == accessCode {
				return fmt.Errorf("access code already added to user")
			}
		}
	}

	// Important: user.GetPassword() contains the hashed password. Hence, it can and should not be updated!
	// Otherwise the password would be updated to the current password hash value.
	// To not update the password, we therefore need to provide an empty string or a user object without password.
	user = &userProto.User{
		Id:          user.Id,
		Email:       user.Email,
		AccessCodes: append(user.AccessCodes, accessCode),
		Settings:    user.Settings,
	}

	userConn, err := microservices.EstablishConnection("user-service", a.tlsCaPath)
	if err != nil {
		glog.Error("failed connecting to service user-service")
		return err
	}
	userClient := userProto.NewUserSvcClient(userConn)
	defer userConn.Close()

	_, err = userClient.UpdateUser(ctx, user)

	if err != nil {
		return err
	}

	return nil
}

func (a AuthServer) RemoveAccessCode(user *userProto.User, accessCode string, ctx context.Context) error {
	if len(user.GetId()) == 0 || len(accessCode) == 0 {
		return fmt.Errorf("bad parameters passed, %s:%s", user.GetId(), accessCode)
	}

	accessCode = strings.ToLower(accessCode)

	var newAccessCodes []string

	newAccessCodes = make([]string, 0)

	if len(user.AccessCodes) == 0 {
		// there were no access codes at this point so what are we doing
		return fmt.Errorf("accesscode %s for user %s was not found", accessCode, user.GetId())
	} else {
		found := false
		for _, ac := range user.AccessCodes {
			if ac == accessCode {
				found = true
			} else {
				newAccessCodes = append(newAccessCodes, ac)
			}
		}
		if !found {
			// the access code wasn't found so no update required
			return nil
		}
	}

	// Important: user.GetPassword() contains the hashed password. Hence, it can and should not be updated!
	// Otherwise the password would be updated to the current password hash value.
	// To not update the password, we therefore need to provide an empty string or a user object without password.
	user = &userProto.User{
		Id:          user.Id,
		Email:       user.Email,
		AccessCodes: newAccessCodes,
		Settings:    user.Settings,
	}

	userConn, err := microservices.EstablishConnection("user-service", a.tlsCaPath)
	if err != nil {
		glog.Error("failed connecting to service user-service")
		return err
	}
	userClient := userProto.NewUserSvcClient(userConn)
	defer userConn.Close()

	_, err = userClient.UpdateUser(ctx, user)

	if err != nil {
		return err
	}

	return nil
}

func (a AuthServer) ChangePassword(user *userProto.User, oldPassword string, newPassword string, ctx context.Context) error {
	if len(user.GetId()) == 0 || len(oldPassword) == 0 || len(newPassword) == 0 {
		return fmt.Errorf("bad parameters passed, %s", user.GetId())
	}

	userConn, err := microservices.EstablishConnection("user-service", a.tlsCaPath)
	if err != nil {
		glog.Error("failed connecting to service user-service")
		return err
	}
	userClient := userProto.NewUserSvcClient(userConn)
	defer userConn.Close()

	err = bcrypt.CompareHashAndPassword([]byte(user.GetPassword()), []byte(oldPassword))

	if err != nil {
		glog.Errorf("old password incorrect for user ID %s: %v", user.GetId(), err)
		return fmt.Errorf("bad password change")
	}

	user.Password = newPassword

	_, err = userClient.UpdateUser(ctx, user)

	if err != nil {
		return err
	}

	return nil
}

func (a AuthServer) UpdateSettings(user *userProto.User, newSettings map[string]string, ctx context.Context) error {
	if len(user.GetId()) == 0 {
		return fmt.Errorf("bad parameters passed, %s", user.GetId())
	}

	userConn, err := microservices.EstablishConnection("user-service", a.tlsCaPath)
	if err != nil {
		glog.Error("failed connecting to service user-service")
		return err
	}
	userClient := userProto.NewUserSvcClient(userConn)
	defer userConn.Close()

	user = &userProto.User{
		Id:       user.GetId(),
		Settings: newSettings,
	}

	_, err = userClient.UpdateUser(ctx, user)

	if err != nil {
		return err
	}

	return nil
}

func (a AuthServer) RegisterWithAccessCodeFunc(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	email := r.PostFormValue("email")
	accessCode := strings.ToLower(r.PostFormValue("access_code"))
	password := r.PostFormValue("password")
	// should we reconcile based on the access code posted in? nah

	if len(email) == 0 || len(accessCode) == 0 || len(password) == 0 {
		util.ReturnHTTPMessage(w, r, 400, "error", "invalid input. required fields: email, access_code, password")
		return
	}

	userConn, err := microservices.EstablishConnection("user-service", a.tlsCaPath)
	if err != nil {
		glog.Error("failed connecting to service user-service")
		util.ReturnHTTPMessage(w, r, 500, "internal server error", "user service unreachable")
		return
	}
	userClient := userProto.NewUserSvcClient(userConn)
	defer userConn.Close()

	userId, err := userClient.CreateUser(r.Context(), &userProto.CreateUserRequest{
		Email:    email,
		Password: password,
	})

	if err != nil {
		if s, ok := status.FromError(err); ok {
			details := s.Details()[0].(*userProto.CreateUserRequest)
			if s.Code() == codes.InvalidArgument {
				glog.Errorf("error creating user, invalid argument for user with email: %s", details.Email)
				util.ReturnHTTPMessage(w, r, 400, "error", s.Message())
				return
			} else if s.Code() == codes.AlreadyExists {
				glog.Errorf("user with email %s already exists", details.Email)
				util.ReturnHTTPMessage(w, r, 409, "error", s.Message())
				return
			}
			glog.Errorf("error creating user: %s", s.Message())
			util.ReturnHTTPMessage(w, r, 500, "error", "error creating user")
			return
		}
		glog.Errorf("error creating user: %s", err.Error())
		util.ReturnHTTPMessage(w, r, 500, "error", "error creating user")
		return
	}

	// from this point, the user is created
	// we are now trying to add the access code he provided

	user, err := userClient.GetUserById(r.Context(), &userProto.UserId{
		Id: userId.GetId(),
	})

	if err != nil {
		if s, ok := status.FromError(err); ok {
			details := s.Details()[0].(*userProto.UserId)
			if s.Code() == codes.InvalidArgument {
				glog.Error("error retrieving created user, no id passed in")
			} else {
				glog.Errorf("error while retrieving created user %s: %s", details.Id, s.Message())
			}
		}
		glog.Errorf("error while retrieving created user: %s", err.Error())
		util.ReturnHTTPMessage(w, r, 500, "error", "error creating user with accesscode")
	}

	err = a.AddAccessCode(user, accessCode, r.Context())

	if err != nil {
		glog.Errorf("error creating user %s %v", email, err)
		util.ReturnHTTPMessage(w, r, 500, "error", "error creating user with accesscode")
		return
	}

	glog.V(2).Infof("created user %s", email)
	util.ReturnHTTPMessage(w, r, 201, "info", "created user")
}

func (a AuthServer) LoginFunc(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	userConn, err := microservices.EstablishConnection("user-service", a.tlsCaPath)
	if err != nil {
		glog.Error("failed connecting to service user-service")
		util.ReturnHTTPMessage(w, r, 500, "internal server error", "user service unreachable")
		return
	}
	userClient := userProto.NewUserSvcClient(userConn)
	defer userConn.Close()

	user, err := userClient.GetUserByEmail(r.Context(), &userProto.GetUserByEmailRequest{Email: email})

	if err != nil {
		glog.Errorf("there was an error retrieving the user %s: %v", email, err)
		util.ReturnHTTPMessage(w, r, 401, "unauthorized", "login failed")
		return
	}

	glog.Infof("unhashed PW - authnsvc: %s", password)
	glog.Infof("hashed PW - authnsvc: %s", user.GetPassword())
	err = bcrypt.CompareHashAndPassword([]byte(user.GetPassword()), []byte(password))

	if err != nil {
		glog.Errorf("password incorrect for user %s: %v", email, err)
		util.ReturnHTTPMessage(w, r, 401, "unauthorized", "login failed")
		return
	}

	token, err := GenerateJWT(user)

	if err != nil {
		glog.Error(err)
	}

	util.ReturnHTTPMessage(w, r, 200, "authorized", token)
}

func GenerateJWT(user *userProto.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.GetEmail(),
		"nbf":   time.Now().Unix(),                     // not valid before now
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // expire in 24 hours
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(user.GetPassword()))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (a *AuthServer) GetAccessSet(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	user, err := a.internalAuthnServer.AuthN(r.Context(), &authn.AuthNRequest{
		Token: token,
	})
	if err != nil {
		util.ReturnHTTPMessage(w, r, http.StatusUnauthorized, "unauthorized", "unauthorized")
		return
	}

	rbacConn, err := microservices.EstablishConnection("rbac-service", a.tlsCaPath)
	if err != nil {
		glog.Error("failed connecting to service rbac-service")
		util.ReturnHTTPMessage(w, r, 500, "internal server error", "rbac service unreachable")
		return
	}
	rbacClient := rbacProto.NewRbacSvcClient(rbacConn)
	defer rbacConn.Close()

	// need to get the user's access set and publish to front end
	as, err := rbacClient.GetAccessSet(r.Context(), &userProto.UserId{Id: user.GetId()})
	if err != nil {
		util.ReturnHTTPMessage(w, r, http.StatusInternalServerError, "internalerror", "internal error fetching access set")
		glog.Error(err)
		return
	}

	encodedAS, err := util.GetProtoMarshaller().Marshal(as)
	if err != nil {
		util.ReturnHTTPMessage(w, r, http.StatusInternalServerError, "internalerror", "internal error encoding access set")
		glog.Error(err)
		return
	}

	util.ReturnHTTPContent(w, r, http.StatusOK, "access_set", encodedAS)
}
