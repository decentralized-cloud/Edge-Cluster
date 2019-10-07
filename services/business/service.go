// Package business implements different business services required by the edge-cluster service
package business

import (
	"context"

	"github.com/decentralized-cloud/edge-cluster/services/repository"
	commonErrors "github.com/micro-business/go-core/system/errors"
)

type businessService struct {
	repositoryService repository.RepositoryContract
}

// NewBusinessService creates new instance of the BusinessService, setting up all dependencies and returns the instance
// repositoryService: Mandatory. Reference to the repository service that can persist the edge cluster related data
// Returns the new service or error if something goes wrong
func NewBusinessService(
	repositoryService repository.RepositoryContract) (BusinessContract, error) {
	if repositoryService == nil {
		return nil, commonErrors.NewArgumentNilError("repositoryService", "repositoryService is required")
	}

	return &businessService{
		repositoryService: repositoryService,
	}, nil
}

// CreateEdgeCluster creates a new edge cluster.
// context: Mandatory The reference to the context
// request: Mandatory. The request to create a new edge cluster
// Returns either the result of creating new edge cluster or error if something goes wrong.
func (service *businessService) CreateEdgeCluster(
	ctx context.Context,
	request *CreateEdgeClusterRequest) (*CreateEdgeClusterResponse, error) {
	response, err := service.repositoryService.CreateEdgeCluster(ctx, &repository.CreateEdgeClusterRequest{
		TenantID:    request.TenantID,
		EdgeCluster: request.EdgeCluster,
	})

	if err != nil {
		return &CreateEdgeClusterResponse{
			Err: mapRepositoryError(err, request.TenantID, ""),
		}, nil
	}

	return &CreateEdgeClusterResponse{
		EdgeClusterID: response.EdgeClusterID,
	}, nil
}

// ReadEdgeCluster read an existing edge cluster
// context: Mandatory The reference to the context
// request: Mandatory. The request to read an existing edge cluster
// Returns either the result of reading an existing edge cluster or error if something goes wrong.
func (service *businessService) ReadEdgeCluster(
	ctx context.Context,
	request *ReadEdgeClusterRequest) (*ReadEdgeClusterResponse, error) {
	response, err := service.repositoryService.ReadEdgeCluster(ctx, &repository.ReadEdgeClusterRequest{
		TenantID:      request.TenantID,
		EdgeClusterID: request.EdgeClusterID,
	})

	if err != nil {
		return &ReadEdgeClusterResponse{
			Err: mapRepositoryError(err, request.TenantID, request.EdgeClusterID),
		}, nil
	}

	return &ReadEdgeClusterResponse{
		EdgeCluster: response.EdgeCluster,
	}, nil
}

// UpdateEdgeCluster update an existing edge cluster
// context: Mandatory The reference to the context
// request: Mandatory. The request to update an existing edge cluster
// Returns either the result of updateing an existing edge cluster or error if something goes wrong.
func (service *businessService) UpdateEdgeCluster(
	ctx context.Context,
	request *UpdateEdgeClusterRequest) (*UpdateEdgeClusterResponse, error) {
	_, err := service.repositoryService.UpdateEdgeCluster(ctx, &repository.UpdateEdgeClusterRequest{
		TenantID:      request.TenantID,
		EdgeClusterID: request.EdgeClusterID,
		EdgeCluster:   request.EdgeCluster,
	})

	if err != nil {
		return &UpdateEdgeClusterResponse{
			Err: mapRepositoryError(err, request.TenantID, request.EdgeClusterID),
		}, nil
	}

	return &UpdateEdgeClusterResponse{}, nil
}

// DeleteEdgeCluster delete an existing edge cluster
// context: Mandatory The reference to the context
// request: Mandatory. The request to delete an existing edge cluster
// Returns either the result of deleting an existing edge cluster or error if something goes wrong.
func (service *businessService) DeleteEdgeCluster(
	ctx context.Context,
	request *DeleteEdgeClusterRequest) (*DeleteEdgeClusterResponse, error) {
	_, err := service.repositoryService.DeleteEdgeCluster(ctx, &repository.DeleteEdgeClusterRequest{
		TenantID:      request.TenantID,
		EdgeClusterID: request.EdgeClusterID,
	})

	if err != nil {
		return &DeleteEdgeClusterResponse{
			Err: mapRepositoryError(err, request.TenantID, request.EdgeClusterID),
		}, nil
	}

	return &DeleteEdgeClusterResponse{}, nil
}

func mapRepositoryError(err error, tenantID, edgeClusterID string) error {
	if repository.IsEdgeClusterAlreadyExistsError(err) {
		return NewEdgeClusterAlreadyExistsErrorWithError(err)
	}

	if repository.IsTenantNotFoundError(err) {
		return NewTenantNotFoundErrorWithError(tenantID, err)
	}

	if repository.IsEdgeClusterNotFoundError(err) {
		return NewEdgeClusterNotFoundErrorWithError(edgeClusterID, err)
	}

	return NewUnknownErrorWithError("", err)
}
