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
		return nil, commonErrors.NewArgumentNilError("businessService", "businessService is required")
	}

	return &endpointCreatorService{
		businessService: businessService,
	}, nil
}

// CreateEdgeClusterEndpoint creates Create Edge Cluster endpoint
// Returns the Create Edge Cluster endpoint
func (service *endpointCreatorService) CreateEdgeClusterEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		if ctx == nil {
			return &businessContract.CreateEdgeClusterResponse{
				Err: commonErrors.NewArgumentNilError("ctx", "ctx is required"),
			}, nil
		}

		if request == nil {
			return &businessContract.CreateEdgeClusterResponse{
				Err: commonErrors.NewArgumentNilError("request", "request is required"),
			}, nil
		}

		castedRequest := request.(*businessContract.CreateEdgeClusterRequest)
		if err := castedRequest.Validate(); err != nil {
			return &businessContract.CreateEdgeClusterResponse{
				Err: commonErrors.NewArgumentErrorWithError("request", "", err),
			}, nil
		}

		return service.businessService.CreateEdgeCluster(ctx, castedRequest)
	}
}

// ReadEdgeClusterEndpoint creates Read Edge Cluster endpoint
// Returns the Read Edge Cluster endpoint
func (service *endpointCreatorService) ReadEdgeClusterEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		if ctx == nil {
			return &businessContract.ReadEdgeClusterResponse{
				Err: commonErrors.NewArgumentNilError("ctx", "ctx is required"),
			}, nil
		}

		if request == nil {
			return &businessContract.ReadEdgeClusterResponse{
				Err: commonErrors.NewArgumentNilError("request", "request is required"),
			}, nil
		}

		castedRequest := request.(*businessContract.ReadEdgeClusterRequest)
		if err := castedRequest.Validate(); err != nil {
			return &businessContract.ReadEdgeClusterResponse{
				Err: commonErrors.NewArgumentErrorWithError("request", "", err),
			}, nil
		}

		return service.businessService.ReadEdgeCluster(ctx, castedRequest)
	}
}

// UpdateEdgeClusterEndpoint creates Update Edge Cluster endpoint
// Returns the Update Edge Cluster endpoint
func (service *endpointCreatorService) UpdateEdgeClusterEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		if ctx == nil {
			return &businessContract.UpdateEdgeClusterResponse{
				Err: commonErrors.NewArgumentNilError("ctx", "ctx is required"),
			}, nil
		}

		if request == nil {
			return &businessContract.UpdateEdgeClusterResponse{
				Err: commonErrors.NewArgumentNilError("request", "request is required"),
			}, nil
		}

		castedRequest := request.(*businessContract.UpdateEdgeClusterRequest)
		if err := castedRequest.Validate(); err != nil {
			return &businessContract.UpdateEdgeClusterResponse{
				Err: commonErrors.NewArgumentErrorWithError("request", "", err),
			}, nil
		}

		return service.businessService.UpdateEdgeCluster(ctx, castedRequest)
	}
}

// DeleteEdgeClusterEndpoint creates Delete Edge Cluster endpoint
// Returns the Delete Edge Cluster endpoint
func (service *endpointCreatorService) DeleteEdgeClusterEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		if ctx == nil {
			return &businessContract.DeleteEdgeClusterResponse{
				Err: commonErrors.NewArgumentNilError("ctx", "ctx is required"),
			}, nil
		}

		if request == nil {
			return &businessContract.DeleteEdgeClusterResponse{
				Err: commonErrors.NewArgumentNilError("request", "request is required"),
			}, nil
		}

		castedRequest := request.(*businessContract.DeleteEdgeClusterRequest)
		if err := castedRequest.Validate(); err != nil {
			return &businessContract.DeleteEdgeClusterResponse{
				Err: commonErrors.NewArgumentErrorWithError("request", "", err),
			}, nil
		}

		return service.businessService.DeleteEdgeCluster(ctx, castedRequest)
	}
}
