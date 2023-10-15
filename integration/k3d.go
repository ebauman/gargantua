package integration

import (
	"fmt"
	"os/exec"
)

var (
	K3dPath     = "k3d"
	ClusterName = "hf-integration-testing"
)

func FindK3d(path string) error {
	if _, err := exec.LookPath(path); err != nil {
		return fmt.Errorf("could not find k3d with path %s", path)
	}

	return nil
}

func CreateCluster() (string, error) {
	if err := FindK3d(K3dPath); err != nil {
		return "", err
	}

	createCmd := exec.Command(K3dPath, "cluster", "create", ClusterName)
	if _, err := createCmd.Output(); err != nil {
		return "", fmt.Errorf("error creating cluster: %s", err.(*exec.ExitError).Stderr)
	}

	kubeconfigCmd := exec.Command(K3dPath, "kubeconfig", "get", ClusterName)
	out, err := kubeconfigCmd.Output()
	if err != nil {
		return "", fmt.Errorf("error getting kubeconfig for cluster: %s", err.(*exec.ExitError).Stderr)
	}

	return string(out), nil
}

func DestroyCluster() error {
	if err := FindK3d(K3dPath); err != nil {
		return err
	}

	destroyCmd := exec.Command(K3dPath, "cluster", "delete", ClusterName)
	if err := destroyCmd.Run(); err != nil {
		return err
	}

	return nil
}
