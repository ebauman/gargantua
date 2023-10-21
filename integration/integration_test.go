//go:build integration

package integration

import (
	"context"
	"github.com/ebauman/crder"
	hfClientset "github.com/hobbyfarm/gargantua/v3/pkg/client/clientset/versioned"
	"github.com/hobbyfarm/gargantua/v3/pkg/crd"
	"github.com/hobbyfarm/gargantua/v3/pkg/preinstall"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"strings"
	"testing"
)

func CreateTestClients(kubeconfig string, t *testing.T) (kubernetes.Interface, hfClientset.Interface) {
	cfg, err := clientcmd.RESTConfigFromKubeConfig([]byte(kubeconfig))
	if err != nil {
		t.Fatalf("error building RESTConfig from kubeconfig: %v", err)
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		t.Fatalf("error setting up kClient: %v", err)
	}

	hfClient, err := hfClientset.NewForConfig(cfg)
	if err != nil {
		t.Fatalf("error setting up hfClient: %v", err)
	}

	if err := InstallCRDs(kubeconfig); err != nil {
		t.Fatalf("error installing CRDs: %v", err)
	}

	// do preinstall as that's required for anything in gargantua to work
	preinstall.Preinstall(context.TODO(), hfClient)

	return kubeClient, hfClient
}

func CreateTestCluster(t *testing.T) (kubeconfig string) {
	var err error
	kubeconfig, err = CreateCluster()

	if err != nil && strings.Contains(err.Error(), "a cluster with that name already exists") {
		DestroyTestCluster(t)
		kubeconfig, err = CreateCluster()
	}

	if err != nil {
		t.Fatalf("error creating cluster :\n%s", err.Error())
	}

	return
}

func DestroyTestCluster(t *testing.T) {
	if err := DestroyCluster(); err != nil {
		t.Error(err)
	}
}

func InstallCRDs(kubeConfig string) error {
	cfg, err := clientcmd.RESTConfigFromKubeConfig([]byte(kubeConfig))
	if err != nil {
		return err
	}

	crds := crd.GenerateCRDsWithCABundleAndServiceReference("", crd.ServiceReference{
		Namespace: namespace,
		Name:      "conversion-service", // garbage, nonexistent filler
	})

	return crder.InstallUpdateCRDs(cfg, crds...)
}
