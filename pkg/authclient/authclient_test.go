package authclient

import (
	hfv2 "github.com/hobbyfarm/gargantua/v3/pkg/apis/hobbyfarm.io/v2"
	hfFake "github.com/hobbyfarm/gargantua/v3/pkg/client/clientset/versioned/fake"
	hfInformers "github.com/hobbyfarm/gargantua/v3/pkg/client/informers/externalversions"
	"github.com/hobbyfarm/gargantua/v3/pkg/rbacclient"
	"k8s.io/client-go/informers"
	kfake "k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/tools/cache"
	"testing"
)

var (
	validEmail   = "user@example.com"
	invalidEmail = "nonexistent@example.com"
)

func createFakeUser(email string) *hfv2.User {
	return &hfv2.User{
		Spec: hfv2.UserSpec{
			Email: email,
		},
	}
}

func CreateAuthClient(t *testing.T) *AuthClient {
	hfFakeClientset := hfFake.NewSimpleClientset(createFakeUser(validEmail))
	kFakeClientset := kfake.NewSimpleClientset()

	hfFakeInformer := hfInformers.NewSharedInformerFactory(hfFakeClientset, 0)

	kFakeInformer := informers.NewSharedInformerFactory(kFakeClientset, 0)

	fakeRbacClient, err := rbacclient.NewRbacClient("default", kFakeInformer)
	if err != nil {
		t.Errorf("error creating fake rbac client: %v", err)
	}

	stopCh := make(chan struct{})
	defer close(stopCh)

	hfFakeInformer.Start(stopCh)
	kFakeInformer.Start(stopCh)

	cache.WaitForCacheSync(stopCh, hfFakeInformer.Hobbyfarm().V2().Users().Informer().HasSynced,
		kFakeInformer.Core().V1().ServiceAccounts().Informer().HasSynced)

	authClient, err := NewAuthClient(hfFakeClientset, hfFakeInformer, fakeRbacClient)
	if err != nil {
		t.Errorf("error creating auth client: %v", err)
	}

	return authClient
}

func TestAuthClient_GetUserByEmail(t *testing.T) {
	authClient := CreateAuthClient(t)

	t.Run("valid email", func(t *testing.T) {
		_, err := authClient.getUserByEmail(validEmail)
		if err != nil {
			t.Errorf("error getting user by email: %v", err)
		}
	})

	t.Run("invalid email", func(t *testing.T) {
		_, err := authClient.getUserByEmail(invalidEmail)
		if err == nil {
			t.Errorf("expected error getting user by email: %v", err)
		}
	})
}
