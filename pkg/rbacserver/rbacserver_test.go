package rbacserver

import (
	"k8s.io/client-go/kubernetes/fake"
	"testing"
)

func CreateRbacServer(t *testing.T) {
	fakeClient := fake.NewSimpleClientset()

	rbacServer := NewRbacServer(fakeClient)
}
