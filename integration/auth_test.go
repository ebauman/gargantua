//go:build integration

package integration

import (
	"context"
	"fmt"
	"github.com/hobbyfarm/gargantua/v3/pkg/accesscode"
	hfv2 "github.com/hobbyfarm/gargantua/v3/pkg/apis/hobbyfarm.io/v2"
	"github.com/hobbyfarm/gargantua/v3/pkg/authclient"
	"github.com/hobbyfarm/gargantua/v3/pkg/authserver"
	hfClientset "github.com/hobbyfarm/gargantua/v3/pkg/client/clientset/versioned"
	hfInformers "github.com/hobbyfarm/gargantua/v3/pkg/client/informers/externalversions"
	"github.com/hobbyfarm/gargantua/v3/pkg/rbacclient"
	"golang.org/x/crypto/bcrypt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
)

var (
	userObjectName = "test-user"
	namespace      = "default"
	email          = "user@example.com"
	password       = "testing123"
	hashedPassword string
	accessCode     = "nonsense"

	userToken string
)

func init() {
	// must do this
	os.Setenv("HF_NAMESPACE", namespace)

	hpBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	hashedPassword = string(hpBytes)
}

func deleteUser(ctx context.Context, client hfClientset.Interface) error {
	return client.HobbyfarmV2().Users(namespace).Delete(ctx, userObjectName, v1.DeleteOptions{})
}

func createUser(ctx context.Context, client hfClientset.Interface) error {
	user := &hfv2.User{
		TypeMeta: v1.TypeMeta{},
		ObjectMeta: v1.ObjectMeta{
			Name:      userObjectName,
			Namespace: namespace,
		},
		Spec: hfv2.UserSpec{
			Email:       email,
			Password:    hashedPassword,
			AccessCodes: nil,
			Settings:    nil,
		},
	}

	_, err := client.HobbyfarmV2().Users(namespace).Create(ctx, user, v1.CreateOptions{})
	return err
}

func Test(t *testing.T) {
	ctx := context.Background()

	stopCh := make(chan struct{})
	defer close(stopCh)

	kubeconfig := CreateTestCluster(t)
	defer DestroyTestCluster(t)

	kubeClient, hfClient := CreateTestClients(kubeconfig, t)

	hfInformerFactory := hfInformers.NewSharedInformerFactoryWithOptions(hfClient, time.Second*30, hfInformers.WithNamespace(namespace))
	kubeInformerFactory := informers.NewSharedInformerFactoryWithOptions(kubeClient, time.Second*30, informers.WithNamespace(namespace))

	hfInformerFactory.Start(stopCh)
	kubeInformerFactory.Start(stopCh)

	rbacClient, err := rbacclient.NewRbacClient(namespace, kubeInformerFactory)
	if err != nil {
		t.Fatal(err)
	}

	authClient, err := authclient.NewAuthClient(hfClient, hfInformerFactory, rbacClient)
	if err != nil {
		t.Fatal(err)
	}

	accessCodeClient, err := accesscode.NewAccessCodeClient(hfClient, ctx)
	if err != nil {
		t.Fatal(err)
	}

	authServer, err := authserver.NewAuthServer(authClient, hfClient, ctx, accessCodeClient, rbacClient)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("it should create a user", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/auth/registerwithaccesscode",
			strings.NewReader(fmt.Sprintf("email=%s&access_code=%s&password=%s", email, accessCode, password)))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		w := httptest.NewRecorder()

		authServer.RegisterWithAccessCodeFunc(w, req)

		if w.Code != 201 {
			t.Errorf("return code not 201: %d", w.Code)
			t.Errorf("error content: %s", w.Body.String())
		}
	})

	t.Run("it should login a valid user", func(t *testing.T) {
		if err := createUser(ctx, hfClient); err != nil {
			t.Error(err)
		}

		defer func() {
			if err := deleteUser(ctx, hfClient); err != nil {
				t.Errorf("error deleting user: %s", err.Error())
			}
		}()

		req := httptest.NewRequest(http.MethodPost, "/auth",
			strings.NewReader(fmt.Sprintf("email=%s&password=%s", email, password)))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		w := httptest.NewRecorder()

		authServer.AuthNFunc(w, req)

		if w.Code != 200 {
			t.Errorf("return code not 200: %d", w.Code)
			t.Errorf("error message: %s", w.Body.String())
		}
	})
}
