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
		EdgeCluster: request.EdgeCluster,
	})

	if err != nil {
		return &CreateEdgeClusterResponse{
			Err: mapRepositoryError(err, ""),
		}, nil
	}

	return &CreateEdgeClusterResponse{
		EdgeClusterID: response.EdgeClusterID,
		EdgeCluster:   response.EdgeCluster,
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
		EdgeClusterID: request.EdgeClusterID,
	})

	if err != nil {
		return &ReadEdgeClusterResponse{
			Err: mapRepositoryError(err, request.EdgeClusterID),
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
	response, err := service.repositoryService.UpdateEdgeCluster(ctx, &repository.UpdateEdgeClusterRequest{
		EdgeClusterID: request.EdgeClusterID,
		EdgeCluster:   request.EdgeCluster,
	})

	if err != nil {
		return &UpdateEdgeClusterResponse{
			Err: mapRepositoryError(err, request.EdgeClusterID),
		}, nil
	}

	return &UpdateEdgeClusterResponse{
		EdgeCluster: response.EdgeCluster,
	}, nil
}

// DeleteEdgeCluster delete an existing edge cluster
// context: Mandatory The reference to the context
// request: Mandatory. The request to delete an existing edge cluster
// Returns either the result of deleting an existing edge cluster or error if something goes wrong.
func (service *businessService) DeleteEdgeCluster(
	ctx context.Context,
	request *DeleteEdgeClusterRequest) (*DeleteEdgeClusterResponse, error) {
	_, err := service.repositoryService.DeleteEdgeCluster(ctx, &repository.DeleteEdgeClusterRequest{
		EdgeClusterID: request.EdgeClusterID,
	})

	if err != nil {
		return &DeleteEdgeClusterResponse{
			Err: mapRepositoryError(err, request.EdgeClusterID),
		}, nil
	}

	return &DeleteEdgeClusterResponse{}, nil
}

// Search returns the list of edge clusters that matched the criteria
// ctx: Mandatory The reference to the context
// request: Mandatory. The request contains the search criteria
// Returns the list of edge clusters that matched the criteria
func (service *businessService) Search(
	ctx context.Context,
	request *SearchRequest) (*SearchResponse, error) {
	result, err := service.repositoryService.Search(ctx, &repository.SearchRequest{
		Pagination:     request.Pagination,
		SortingOptions: request.SortingOptions,
		EdgeClusterIDs: request.EdgeClusterIDs,
		TenantIDs:      request.TenantIDs,
	})

	if err != nil {
		return &SearchResponse{
			Err: mapRepositoryError(err, ""),
		}, nil
	}

	return &SearchResponse{
		HasPreviousPage: result.HasPreviousPage,
		HasNextPage:     result.HasNextPage,
		EdgeClusters:    result.EdgeClusters,
	}, nil
}

func mapRepositoryError(err error, edgeClusterID string) error {
	if repository.IsEdgeClusterAlreadyExistsError(err) {
		return NewEdgeClusterAlreadyExistsErrorWithError(err)
	}

	if repository.IsEdgeClusterNotFoundError(err) {
		return NewEdgeClusterNotFoundErrorWithError(edgeClusterID, err)
	}

	return NewUnknownErrorWithError("", err)
}
