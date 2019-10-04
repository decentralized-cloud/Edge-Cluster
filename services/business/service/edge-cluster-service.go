// Package service implements the different EdgeCluster business services
package service

import (
	"context"

	"github.com/decentralized-cloud/edge-cluster/services/business/contract"
	repositoryContract "github.com/decentralized-cloud/edge-cluster/services/repository/contract"
	commonErrors "github.com/micro-business/go-core/system/errors"
)

type edgeClusterService struct {
	repositoryService repositoryContract.EdgeClusterRepositoryServiceContract
}

// NewEdgeClusterService creates new instance of the EdgeClusterService, setting up all dependencies and returns the instance
// repositoryService: Mandatory. Reference to the repository service that can persist the edge cluster related data
// Returns the new service or error if something goes wrong
func NewEdgeClusterService(
	repositoryService repositoryContract.EdgeClusterRepositoryServiceContract) (contract.EdgeClusterServiceContract, error) {
	if repositoryService == nil {
		return nil, commonErrors.NewArgumentNilError("repositoryService", "repositoryService is required")
	}

	return &edgeClusterService{
		repositoryService: repositoryService,
	}, nil
}

// CreateEdgeCluster creates a new edge cluster.
// context: Mandatory The reference to the context
// request: Mandatory. The request to create a new edge cluster
// Returns either the result of creating new edge cluster or error if something goes wrong.
func (service *edgeClusterService) CreateEdgeCluster(
	ctx context.Context,
	request *contract.CreateEdgeClusterRequest) (*contract.CreateEdgeClusterResponse, error) {
	if ctx == nil {
		return &contract.CreateEdgeClusterResponse{
			Err: commonErrors.NewArgumentNilError("ctx", "ctx is required"),
		}, nil
	}

	if request == nil {
		return &contract.CreateEdgeClusterResponse{
			Err: commonErrors.NewArgumentNilError("request", "request is required"),
		}, nil
	}

	if err := request.Validate(); err != nil {
		return &contract.CreateEdgeClusterResponse{
			Err: commonErrors.NewArgumentErrorWithError("request", "", err),
		}, nil
	}

	response, err := service.repositoryService.CreateEdgeCluster(ctx, &repositoryContract.CreateEdgeClusterRequest{
		TenantID:    request.TenantID,
		EdgeCluster: request.EdgeCluster,
	})

	if err != nil {
		return &contract.CreateEdgeClusterResponse{
			Err: mapRepositoryError(err, request.TenantID, ""),
		}, nil
	}

	return &contract.CreateEdgeClusterResponse{
		EdgeClusterID: response.EdgeClusterID,
	}, nil
}

// ReadEdgeCluster read an existing edge cluster
// context: Mandatory The reference to the context
// request: Mandatory. The request to read an existing edge cluster
// Returns either the result of reading an existing edge cluster or error if something goes wrong.
func (service *edgeClusterService) ReadEdgeCluster(
	ctx context.Context,
	request *contract.ReadEdgeClusterRequest) (*contract.ReadEdgeClusterResponse, error) {
	if ctx == nil {
		return &contract.ReadEdgeClusterResponse{
			Err: commonErrors.NewArgumentNilError("ctx", "ctx is required"),
		}, nil
	}

	if request == nil {
		return &contract.ReadEdgeClusterResponse{
			Err: commonErrors.NewArgumentNilError("request", "request is required"),
		}, nil
	}

	if err := request.Validate(); err != nil {
		return &contract.ReadEdgeClusterResponse{
			Err: commonErrors.NewArgumentErrorWithError("request", "", err),
		}, nil
	}

	response, err := service.repositoryService.ReadEdgeCluster(ctx, &repositoryContract.ReadEdgeClusterRequest{
		TenantID:      request.TenantID,
		EdgeClusterID: request.EdgeClusterID,
	})

	if err != nil {
		return &contract.ReadEdgeClusterResponse{
			Err: mapRepositoryError(err, request.TenantID, request.EdgeClusterID),
		}, nil
	}

	return &contract.ReadEdgeClusterResponse{
		EdgeCluster: response.EdgeCluster,
	}, nil
}

// UpdateEdgeCluster update an existing edge cluster
// context: Mandatory The reference to the context
// request: Mandatory. The request to update an existing edge cluster
// Returns either the result of updateing an existing edge cluster or error if something goes wrong.
func (service *edgeClusterService) UpdateEdgeCluster(
	ctx context.Context,
	request *contract.UpdateEdgeClusterRequest) (*contract.UpdateEdgeClusterResponse, error) {
	if ctx == nil {
		return &contract.UpdateEdgeClusterResponse{
			Err: commonErrors.NewArgumentNilError("ctx", "ctx is required"),
		}, nil
	}

	if request == nil {
		return &contract.UpdateEdgeClusterResponse{
			Err: commonErrors.NewArgumentNilError("request", "request is required"),
		}, nil
	}

	if err := request.Validate(); err != nil {
		return &contract.UpdateEdgeClusterResponse{
			Err: commonErrors.NewArgumentErrorWithError("request", "", err),
		}, nil
	}

	_, err := service.repositoryService.UpdateEdgeCluster(ctx, &repositoryContract.UpdateEdgeClusterRequest{
		TenantID:      request.TenantID,
		EdgeClusterID: request.EdgeClusterID,
		EdgeCluster:   request.EdgeCluster,
	})

	if err != nil {
		return &contract.UpdateEdgeClusterResponse{
			Err: mapRepositoryError(err, request.TenantID, request.EdgeClusterID),
		}, nil
	}

	return &contract.UpdateEdgeClusterResponse{}, nil
}

// DeleteEdgeCluster delete an existing edge cluster
// context: Mandatory The reference to the context
// request: Mandatory. The request to delete an existing edge cluster
// Returns either the result of deleting an existing edge cluster or error if something goes wrong.
func (service *edgeClusterService) DeleteEdgeCluster(
	ctx context.Context,
	request *contract.DeleteEdgeClusterRequest) (*contract.DeleteEdgeClusterResponse, error) {
	if ctx == nil {
		return &contract.DeleteEdgeClusterResponse{
			Err: commonErrors.NewArgumentNilError("ctx", "ctx is required"),
		}, nil
	}

	if request == nil {
		return &contract.DeleteEdgeClusterResponse{
			Err: commonErrors.NewArgumentNilError("request", "request is required"),
		}, nil
	}

	if err := request.Validate(); err != nil {
		return &contract.DeleteEdgeClusterResponse{
			Err: commonErrors.NewArgumentErrorWithError("request", "", err),
		}, nil
	}

	_, err := service.repositoryService.DeleteEdgeCluster(ctx, &repositoryContract.DeleteEdgeClusterRequest{
		TenantID:      request.TenantID,
		EdgeClusterID: request.EdgeClusterID,
	})

	if err != nil {
		return &contract.DeleteEdgeClusterResponse{
			Err: mapRepositoryError(err, request.TenantID, request.EdgeClusterID),
		}, nil
	}

	return &contract.DeleteEdgeClusterResponse{}, nil
}

func mapRepositoryError(err error, tenantID, edgeClusterID string) error {
	if repositoryContract.IsEdgeClusterAlreadyExistsError(err) {
		return contract.NewEdgeClusterAlreadyExistsErrorWithError(err)
	}

	if repositoryContract.IsTenantNotFoundError(err) {
		return contract.NewTenantNotFoundErrorWithError(tenantID, err)
	}

	if repositoryContract.IsEdgeClusterNotFoundError(err) {
		return contract.NewEdgeClusterNotFoundErrorWithError(edgeClusterID, err)
	}

	return contract.NewUnknownErrorWithError("", err)
}
