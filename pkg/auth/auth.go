package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/golang/glog"
	hfv1 "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"
	hfClientset "github.com/hobbyfarm/gargantua/pkg/client/clientset/versioned"
	hfInformers "github.com/hobbyfarm/gargantua/pkg/client/informers/externalversions"
	"k8s.io/client-go/tools/cache"
	"net/http"
	"time"
)

const (
	emailIndex = "authclient.hobbyfarm.io/user-email-index"
)

type AuthClient struct {
	hfClientset *hfClientset.Clientset
	userIndexer cache.Indexer
}

func New(hfClientset *hfClientset.Clientset, factory hfInformers.SharedInformerFactory) (*AuthClient, error) {
	a := AuthClient{}
	a.hfClientset = hfClientset
	inf := factory.Hobbyfarm().V1().Users().Informer()
	indexers := map[string]cache.IndexFunc{emailIndex: emailIndexer}
	err := inf.AddIndexers(indexers)
	if err != nil {
		return nil, err
	}
	a.userIndexer = inf.GetIndexer()
	return &a, nil
}

func emailIndexer(obj interface{}) ([]string, error) {
	user, ok := obj.(*hfv1.User)
	if !ok {
		return []string{}, nil
	}
	return []string{user.Spec.Email}, nil
}

func (a AuthClient) GetUserByEmail(email string) (hfv1.User, error) {
	if len(email) == 0 {
		return hfv1.User{}, fmt.Errorf("email passed in was empty")
	}

	obj, err := a.userIndexer.ByIndex(emailIndex, email)
	if err != nil {
		return hfv1.User{}, fmt.Errorf("error while retrieving user by e-mail: %s with error: %v", email, err)
	}

	if len(obj) < 1 {
		return hfv1.User{}, fmt.Errorf("user not found by email: %s", email)
	}

	user, ok := obj[0].(*hfv1.User)

	if !ok {
		return hfv1.User{}, fmt.Errorf("error while converting user found by email to object: %s", email)
	}

	return *user, nil
}

func (a AuthClient) AuthWS(w http.ResponseWriter, r *http.Request) (hfv1.User, error) {
	token := r.URL.Query().Get("auth")

	if len(token) == 0 {
		glog.Errorf("no auth token passed in websocket query string")
		//util.ReturnHTTPMessage(w, r, 403, "forbidden", "no token passed")
		return hfv1.User{}, fmt.Errorf("authentication failed")
	}

	return a.performAuth(token, false)
}

func (a AuthClient) AuthN(token string, admin bool) (hfv1.User, error) {
	return a.performAuth(token, admin)
}

func (a AuthClient) performAuth(token string, admin bool) (hfv1.User, error) {
	//glog.V(2).Infof("token passed in was: %s", token)

	user, err := a.ValidateJWT(token)

	if err != nil {
		glog.Errorf("error validating user %v", err)
		//util.ReturnHTTPMessage(w, r, 403, "forbidden", "forbidden")
		return hfv1.User{}, fmt.Errorf("authentication failed")
	}

	glog.V(2).Infof("validated user %s!", user.Spec.Email)
	if admin {
		if user.Spec.Admin {
			return user, nil
		} else {
			glog.Errorf("AUDIT: User %s attempted to access an admin protected resource.", user.Spec.Email)
			return hfv1.User{}, fmt.Errorf("authentication failed")
		}
	}
	//util.ReturnHTTPMessage(w, r, 200, "success", "test successful. valid token")
	return user, nil
}

func (a AuthClient) ValidateJWT(tokenString string) (hfv1.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		var user hfv1.User
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			var err error
			user, err = a.GetUserByEmail(fmt.Sprint(claims["email"]))
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
		return hfv1.User{}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user, err := a.GetUserByEmail(fmt.Sprint(claims["email"]))
		if err != nil {
			return hfv1.User{}, err
		} else {
			return user, nil
		}
	}
	glog.Errorf("error while validating user")
	return hfv1.User{}, fmt.Errorf("error while validating user")
}

func (a AuthClient) GenerateJWT(user hfv1.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Spec.Email,
		"nbf":   time.Now().Unix(),                     // not valid before now
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // expire in 24 hours
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(user.Spec.Password))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}