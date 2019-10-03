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
		EdgeCluster: request.EdgeCluster,
	})

	if err != nil {
		if repositoryContract.IsEdgeClusterAlreadyExistsError(err) {
			return &contract.CreateEdgeClusterResponse{
				Err: contract.NewEdgeClusterAlreadyExistsErrorWithError(err),
			}, nil
		}

		return &contract.CreateEdgeClusterResponse{
			Err: contract.NewUnknownErrorWithError("", err),
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
		EdgeClusterID: request.EdgeClusterID,
	})

	if err != nil {
		if repositoryContract.IsEdgeClusterNotFoundError(err) {
			return &contract.ReadEdgeClusterResponse{
				Err: contract.NewEdgeClusterNotFoundErrorWithError(request.EdgeClusterID, err),
			}, nil
		}

		return &contract.ReadEdgeClusterResponse{
			Err: contract.NewUnknownErrorWithError("", err),
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
		EdgeClusterID: request.EdgeClusterID,
		EdgeCluster:   request.EdgeCluster,
	})

	if err != nil {
		if repositoryContract.IsEdgeClusterNotFoundError(err) {
			return &contract.UpdateEdgeClusterResponse{
				Err: contract.NewEdgeClusterNotFoundErrorWithError(request.EdgeClusterID, err),
			}, nil
		}

		return &contract.UpdateEdgeClusterResponse{
			Err: contract.NewUnknownErrorWithError("", err),
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
		EdgeClusterID: request.EdgeClusterID,
	})

	if err != nil {
		if repositoryContract.IsEdgeClusterNotFoundError(err) {
			return &contract.DeleteEdgeClusterResponse{
				Err: contract.NewEdgeClusterNotFoundErrorWithError(request.EdgeClusterID, err),
			}, nil
		}

		return &contract.DeleteEdgeClusterResponse{
			Err: contract.NewUnknownErrorWithError("", err),
		}, nil
	}

	return &contract.DeleteEdgeClusterResponse{}, nil
}
