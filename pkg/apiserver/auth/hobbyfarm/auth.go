package hobbyfarm

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"github.com/golang/glog"
	hfv1 "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"
	"github.com/hobbyfarm/gargantua/pkg/apiserver/auth"
	"github.com/hobbyfarm/gargantua/pkg/apiserver/errors"
	hfClientset "github.com/hobbyfarm/gargantua/pkg/client/clientset/versioned"
	hfInformers "github.com/hobbyfarm/gargantua/pkg/client/informers/externalversions"
	"k8s.io/client-go/tools/cache"
	"strings"
)

const (
	emailIndex = "auth.hobbyfarm.io/user-email-index"
)

const (
	Default auth.AuthNProvider = "default"
	Allow   auth.AuthZProvider = "allow"
	Admin   auth.AuthZProvider = "admin"
)

type Server struct {
	hfClientSet *hfClientset.Clientset
	userIndexer cache.Indexer
}

func New(hfClientset *hfClientset.Clientset, hfInformerFactory hfInformers.SharedInformerFactory) (*Server, error) {
	s := Server{}
	s.hfClientSet = hfClientset
	inf := hfInformerFactory.Hobbyfarm().V1().Users().Informer()
	indexers := map[string]cache.IndexFunc{emailIndex: emailIndexer}
	_ = inf.AddIndexers(indexers)
	s.userIndexer = inf.GetIndexer()
	return &s, nil
}

func (s *Server) AuthNProviders() map[auth.AuthNProvider]func() func(*fiber.Ctx) {
	provides := make(map[auth.AuthNProvider]func() func(*fiber.Ctx))

	provides[Default] = s.AuthN

	return provides
}

func (s *Server) AuthZProviders() map[auth.AuthZProvider]func() func(*fiber.Ctx) {
	provides := make(map[auth.AuthZProvider]func() func(*fiber.Ctx))

	provides[Allow] = func() func(*fiber.Ctx) {
		return func(c *fiber.Ctx) {
			c.Next()
		}
	}

	provides[Admin] = s.AuthZ

	return provides
}

func emailIndexer(obj interface{}) ([]string, error) {
	user, ok := obj.(*hfv1.User)
	if !ok {
		return []string{}, nil
	}
	return []string{user.Spec.Email}, nil
}

func (s *Server) AuthZ() func(*fiber.Ctx) {
	return func(c *fiber.Ctx) {
		u := c.Locals("user")
		user, ok := u.(hfv1.User)
		if !ok {
			c.Next(errors.Internal("unable to assert user into hfv1.User"))
			return
		}

		if !user.Spec.Admin {
			c.Next(errors.Unauthorized("unauthorized"))
		}

		c.Next()
	}
}

func (s *Server) AuthN() func(*fiber.Ctx) {
	// return middleware handler
	return func(c *fiber.Ctx) {
		token := c.Get("Authorization")

		if len(token) == 0 {
			c.Next(errors.Unauthorized("authentication failed"))
			return
		}

		var finalToken string

		splitToken := strings.Split(token, "Bearer")
		finalToken = strings.TrimSpace(splitToken[1])

		user, err := s.performAuth(finalToken)

		if err != nil {
			c.Next(err)
			return
		}

		c.Locals("user", user)
		c.Next()
	}
}

func (s *Server) performAuth(token string) (hfv1.User, error) {
	user, err := s.validateJWT(token)

	if err != nil {
		glog.Errorf("error validating user %v", err)
		return hfv1.User{}, errors.Unauthorized("authentication failed")
	}

	glog.V(2).Infof("validated user %s!", user.Spec.Email)
	// this used to perform "if admin" check, but moved that to AuthZ component
	return user, nil
}

func (s *Server) validateJWT(tokenString string) (hfv1.User, error) {
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
		return hfv1.User{}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user, err := s.getUserByEmail(fmt.Sprint(claims["email"]))
		if err != nil {
			return hfv1.User{}, err
		} else {
			return user, nil
		}
	}
	glog.Errorf("error while validating user")
	return hfv1.User{}, fmt.Errorf("error while validating user")
}

func (s *Server) getUserByEmail(email string) (hfv1.User, error) {
	if len(email) == 0 {
		return hfv1.User{}, fmt.Errorf("email passed in was empty")
	}

	obj, err := s.userIndexer.ByIndex(emailIndex, email)
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
