// Package edgecluster implements the services that are used to provision a supported edge cluster and managing them
package edgecluster

import (
	"context"
	"os"
	"path/filepath"

	"github.com/decentralized-cloud/edge-cluster/services/edgecluster/k3s"
	"github.com/decentralized-cloud/edge-cluster/services/edgecluster/types"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"github.com/savsgio/go-logger"
	"go.uber.org/zap"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type edgeClusterFactoryService struct {
	logger        *zap.Logger
	k8sRestConfig *rest.Config
}

// NewEdgeClusterFactoryService creates new instance of the edgeClusterFactoryService, setting up all dependencies and returns the instance
// logger: Mandatory. Reference to the logger service
// Returns the new service or error if something goes wrong
func NewEdgeClusterFactoryService(logger *zap.Logger) (types.EdgeClusterFactoryContract, error) {
	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	service := edgeClusterFactoryService{
		logger: logger,
	}

	k8sRestConfig, err := service.getRestConfig()
	if err != nil {
		return nil, types.NewUnknownErrorWithError("Failed to retrieve rest config", err)
	}

	service.k8sRestConfig = k8sRestConfig

	return &service, nil
}

// Create instantiates a new edge cluster provisioner of a requested edge cluster type and returns
// it to the caller.
// ctx: Optional The reference to the context
// edgeClusterType: Mandatory. The type of edge cluster provisioner to be instantiated
// Returns either the result of instantiating a edge cluster provisioner or error if something goes wrong.
func (service *edgeClusterFactoryService) Create(
	ctx context.Context,
	edgeClusterType types.EdgeClusterType) (types.EdgeClusterProvisionerContract, error) {
	if edgeClusterType == types.K3S {
		return k3s.NewK3SProvisioner(
			service.logger,
			service.k8sRestConfig)
	}

	return nil, types.NewEdgeClusterNotSupportedError(edgeClusterType)
}

func (service *edgeClusterFactoryService) getRestConfig() (*rest.Config, error) {
	if kubeConfig := os.Getenv("KUBECONFIG"); kubeConfig != "" {
		service.logger.Info("path ", zap.String("KUBECONFIG", kubeConfig))

		return clientcmd.BuildConfigFromFlags("", kubeConfig)
	}

	homePath, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	service.logger.Info("homePath ", zap.String("path", homePath))

	kubeConfigFilePath := filepath.Join(homePath, ".kube", "config")
	logger.Info("kubePath ", zap.String("kube path", kubeConfigFilePath))

	_, err = os.Stat(kubeConfigFilePath)
	if !os.IsNotExist(err) {
		return clientcmd.BuildConfigFromFlags("", kubeConfigFilePath)
	}

	return rest.InClusterConfig()
}
