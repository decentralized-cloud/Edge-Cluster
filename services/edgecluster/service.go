// Package edgecluster implements the services that are used to provision a supported edge cluster and managing them
package edgecluster

import (
	"context"

	"github.com/decentralized-cloud/edge-cluster/services/edgecluster/k3s"
	"github.com/decentralized-cloud/edge-cluster/services/edgecluster/types"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
)

type edgeClusterFactoryService struct {
	logger *zap.Logger
}

// NewEdgeClusterFactoryService creates new instance of the edgeClusterFactoryService, setting up all dependencies and returns the instance
// logger: Mandatory. Reference to the logger service
// Returns the new service or error if something goes wrong
func NewEdgeClusterFactoryService(logger *zap.Logger) (types.EdgeClusterFactoryContract, error) {
	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	return &edgeClusterFactoryService{
		logger: logger,
	}, nil
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
		return k3s.NewK3SProvisioner(service.logger)
	}

	return nil, types.NewEdgeClusterNotSupportedError(edgeClusterType)
}
