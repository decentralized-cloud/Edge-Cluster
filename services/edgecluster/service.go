// Package edgecluster implements the services that are used to provision a supported edge cluster and managing them
package edgecluster

import (
	"context"
	"os"
	"path/filepath"

	"github.com/decentralized-cloud/edge-cluster/models"
	"github.com/decentralized-cloud/edge-cluster/services/configuration"
	"github.com/decentralized-cloud/edge-cluster/services/edgecluster/k3s"
	"github.com/decentralized-cloud/edge-cluster/services/edgecluster/types"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"github.com/savsgio/go-logger"
	"go.uber.org/zap"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type edgeClusterFactoryService struct {
	logger               *zap.Logger
	k8sRestConfig        *rest.Config
	configurationService configuration.ConfigurationContract
}

// NewEdgeClusterFactoryService creates new instance of the edgeClusterFactoryService, setting up all dependencies and returns the instance
// logger: Mandatory. Reference to the logger service
// Returns the new service or error if something goes wrong
func NewEdgeClusterFactoryService(
	logger *zap.Logger,
	configurationService configuration.ConfigurationContract) (types.EdgeClusterFactoryContract, error) {
	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	if configurationService == nil {
		return nil, commonErrors.NewArgumentNilError("configurationService", "configurationService is required")
	}

	service := edgeClusterFactoryService{
		logger:               logger,
		configurationService: configurationService,
	}

	k8sRestConfig, err := service.getRestConfig()
	if err != nil {
		return nil, types.NewUnknownErrorWithError("failed to retrieve rest config", err)
	}

	service.k8sRestConfig = k8sRestConfig

	return &service, nil
}

// Create instantiates a new edge cluster provisioner of a requested edge cluster type and returns
// it to the caller.
// ctx: Mandatory The reference to the context
// clusterType: Mandatory. The type of edge cluster provisioner to be instantiated
// Returns either the result of instantiating a edge cluster provisioner or error if something goes wrong.
func (service *edgeClusterFactoryService) Create(
	ctx context.Context,
	clusterType models.ClusterType) (types.EdgeClusterProvisionerContract, error) {
	if clusterType == models.K3S {
		return k3s.NewK3SProvisioner(
			service.logger,
			service.k8sRestConfig,
			service.configurationService)
	}

	return nil, types.NewEdgeClusterTypeNotSupportedError(clusterType)
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
