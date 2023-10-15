package sessionserver

import (
	"github.com/hobbyfarm/gargantua/v3/pkg/authclient"
	"github.com/hobbyfarm/gargantua/v3/pkg/client/clientset/versioned/fake"
	"testing"
)

func Test_NewSessionFunc(t *testing.T) {
	fakeHfClientset := fake.NewSimpleClientset()

	authClient, err := authclient.NewAuthClient(fakeHfClientset)

	sessionServer := NewSessionServer()
}
