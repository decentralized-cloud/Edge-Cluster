// Package k3s provides functionality to provision a K3S edge cluster type and manage them
package k3s

import (
	"context"
	"os"
	"path/filepath"

	"github.com/decentralized-cloud/edge-cluster/services/edgecluster/types"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type k3sProvisioner struct {
	logger    *zap.Logger
	clientset *kubernetes.Clientset
}

// NewK3SProvisioner creates new instance of the k3sProvisioner, setting up all dependencies and returns the instance
// logger: Mandatory. Reference to the logger service
// Returns the new service or error if something goes wrong
func NewK3SProvisioner(logger *zap.Logger) (types.EdgeClusterProvisionerContract, error) {
	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	restConfig, err := getRestConfig()
	if err != nil {
		return nil, types.NewUnknownErrorWithError("Failed to retrieve rest config", err)
	}

	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, types.NewUnknownErrorWithError("Failed to create client set", err)
	}

	return &k3sProvisioner{
		logger:    logger,
		clientset: clientset,
	}, nil
}

// NewProvision provisions a new K3S edge cluster.
// ctx: Mandatory The reference to the context
// request: Mandatory. The request to provision a new edge cluster
// Returns either the result of provisioning new K3S edge cluster or error if something goes wrong.
func (service *k3sProvisioner) NewProvision(
	ctx context.Context,
	request *types.NewProvisionRequest) (*types.NewProvisionResponse, error) {
	return &types.NewProvisionResponse{}, nil
}

func getRestConfig() (*rest.Config, error) {
	kubeconfig := os.Getenv("KUBECONFIG")
	if kubeconfig != "" {
		return clientcmd.BuildConfigFromFlags("", kubeconfig)
	}

	homePath, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	kubeconfigFilePath := filepath.Join(homePath, ".kube", "config")
	_, err = os.Stat(kubeconfigFilePath)
	if !os.IsNotExist(err) {
		return clientcmd.BuildConfigFromFlags("", kubeconfigFilePath)
	}

	return rest.InClusterConfig()
}
