// Package service implements the service that provides endpoints to be used by the transport layer
package service

import (
	"context"

	businessContract "github.com/decentralized-cloud/edge-cluster/services/business/contract"
	"github.com/decentralized-cloud/edge-cluster/services/endpoint/contract"
	"github.com/go-kit/kit/endpoint"
	commonErrors "github.com/micro-business/go-core/system/errors"
)

type endpointCreatorService struct {
	businessService businessContract.EdgeClusterServiceContract
}

// NewEndpointCreatorService creates new instance of the EndpointCreatorService, setting up all dependencies and returns the instance
// businessService: Mandatory. Reference to the instance of the Edge Cluster  service
// Returns the new service or error if something goes wrong
func NewEndpointCreatorService(
	businessService businessContract.EdgeClusterServiceContract) (contract.EndpointCreatorContract, error) {
	if businessService == nil {
		return nil, commonErrors.NewArgumentError("businessService", "businessService is required")
	}

	return &endpointCreatorService{
		businessService: businessService,
	}, nil
}

// CreateEdgeClusterEndpoint creates Create Edge Cluster endpoint
// Returns the Create Edge Cluster endpoint
func (service *endpointCreatorService) CreateEdgeClusterEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return service.businessService.CreateEdgeCluster(ctx, request.(*businessContract.CreateEdgeClusterRequest))
	}
}

// ReadEdgeClusterEndpoint creates Read Edge Cluster endpoint
// Returns the Read Edge Cluster endpoint
func (service *endpointCreatorService) ReadEdgeClusterEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return service.businessService.ReadEdgeCluster(ctx, request.(*businessContract.ReadEdgeClusterRequest))
	}
}

// UpdateEdgeClusterEndpoint creates Update Edge Cluster endpoint
// Returns the Update Edge Cluster endpoint
func (service *endpointCreatorService) UpdateEdgeClusterEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return service.businessService.UpdateEdgeCluster(ctx, request.(*businessContract.UpdateEdgeClusterRequest))
	}
}

// DeleteEdgeClusterEndpoint creates Delete Edge Cluster endpoint
// Returns the Delete Edge Cluster endpoint
func (service *endpointCreatorService) DeleteEdgeClusterEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return service.businessService.DeleteEdgeCluster(ctx, request.(*businessContract.DeleteEdgeClusterRequest))
	}
}
