// Package business implements different business services required by the edge-cluster service
package business

import (
	"context"

	edgeClusterTypes "github.com/decentralized-cloud/edge-cluster/services/edgecluster/types"
	"github.com/decentralized-cloud/edge-cluster/services/repository"
	commonErrors "github.com/micro-business/go-core/system/errors"
)

type businessService struct {
	repositoryService         repository.RepositoryContract
	edgeClusterFactoryService edgeClusterTypes.EdgeClusterFactoryContract
}

// NewBusinessService creates new instance of the BusinessService, setting up all dependencies and returns the instance
// repositoryService: Mandatory. Reference to the repository service that can persist the edge cluster related data
// edgeClusterFactoryService: Mandatory. Reference to the factory service that can that can create different type of supported
// edge cluster provisioner
// Returns the new service or error if something goes wrong
func NewBusinessService(
	repositoryService repository.RepositoryContract,
	edgeClusterFactoryService edgeClusterTypes.EdgeClusterFactoryContract) (BusinessContract, error) {
	if repositoryService == nil {
		return nil, commonErrors.NewArgumentNilError("repositoryService", "repositoryService is required")
	}

	if edgeClusterFactoryService == nil {
		return nil, commonErrors.NewArgumentNilError("edgeClusterFactoryService", "edgeClusterFactoryService is required")
	}

	return &businessService{
		repositoryService:         repositoryService,
		edgeClusterFactoryService: edgeClusterFactoryService,
	}, nil
}

// CreateEdgeCluster creates a new edge cluster.
// context: Mandatory The reference to the context
// request: Mandatory. The request to create a new edge cluster
// Returns either the result of creating new edge cluster or error if something goes wrong.
func (service *businessService) CreateEdgeCluster(
	ctx context.Context,
	request *CreateEdgeClusterRequest) (*CreateEdgeClusterResponse, error) {

	repositoryResponse, err := service.repositoryService.CreateEdgeCluster(ctx, &repository.CreateEdgeClusterRequest{
		EdgeCluster: request.EdgeCluster,
	})

	if err != nil {
		return &CreateEdgeClusterResponse{
			Err: mapRepositoryError(err, ""),
		}, nil
	}

	edgeClusterProvisioner, err := service.edgeClusterFactoryService.Create(ctx, request.EdgeCluster.ClusterType)
	if err != nil {
		return nil, NewUnknownErrorWithError("Failed to create egde cluster provisioner", err)
	}

	if _, err = edgeClusterProvisioner.CreateProvision(
		ctx,
		&edgeClusterTypes.CreateProvisionRequest{
			EdgeClusterID: repositoryResponse.EdgeClusterID,
			ClusterSecret: request.EdgeCluster.ClusterSecret,
		}); err != nil {

		return nil, NewUnknownErrorWithError("Failed to provision egde cluster", err)
	}

	response := &CreateEdgeClusterResponse{
		EdgeClusterID: repositoryResponse.EdgeClusterID,
		EdgeCluster:   repositoryResponse.EdgeCluster,
		Cursor:        repositoryResponse.Cursor,
	}

	if provisionDetailsReponse, err := edgeClusterProvisioner.GetProvisionDetails(
		ctx,
		&edgeClusterTypes.GetProvisionDetailsRequest{EdgeClusterID: repositoryResponse.EdgeClusterID}); err == nil {
		response.ProvisionDetails = provisionDetailsReponse.ProvisionDetails
	}

	return response, nil
}

// ReadEdgeCluster read an existing edge cluster
// context: Mandatory The reference to the context
// request: Mandatory. The request to read an existing edge cluster
// Returns either the result of reading an existing edge cluster or error if something goes wrong.
func (service *businessService) ReadEdgeCluster(
	ctx context.Context,
	request *ReadEdgeClusterRequest) (*ReadEdgeClusterResponse, error) {
	repositoryResponse, err := service.repositoryService.ReadEdgeCluster(ctx, &repository.ReadEdgeClusterRequest{
		EdgeClusterID: request.EdgeClusterID,
	})

	if err != nil {
		return &ReadEdgeClusterResponse{
			Err: mapRepositoryError(err, request.EdgeClusterID),
		}, nil
	}

	edgeClusterProvisioner, err := service.edgeClusterFactoryService.Create(ctx, repositoryResponse.EdgeCluster.ClusterType)
	if err != nil {
		return nil, NewUnknownErrorWithError("Failed to create egde cluster provisioner", err)
	}

	response := &ReadEdgeClusterResponse{
		EdgeCluster: repositoryResponse.EdgeCluster,
	}

	if provisionDetailsReponse, err := edgeClusterProvisioner.GetProvisionDetails(
		ctx,
		&edgeClusterTypes.GetProvisionDetailsRequest{EdgeClusterID: request.EdgeClusterID}); err == nil {
		response.ProvisionDetails = provisionDetailsReponse.ProvisionDetails
	}

	return response, nil
}

// UpdateEdgeCluster update an existing edge cluster
// context: Mandatory The reference to the context
// request: Mandatory. The request to update an existing edge cluster
// Returns either the result of updateing an existing edge cluster or error if something goes wrong.
func (service *businessService) UpdateEdgeCluster(
	ctx context.Context,
	request *UpdateEdgeClusterRequest) (*UpdateEdgeClusterResponse, error) {

	repositoryResponse, err := service.repositoryService.UpdateEdgeCluster(ctx, &repository.UpdateEdgeClusterRequest{
		EdgeClusterID: request.EdgeClusterID,
		EdgeCluster:   request.EdgeCluster,
	})

	if err != nil {
		return &UpdateEdgeClusterResponse{
			Err: mapRepositoryError(err, request.EdgeClusterID),
		}, nil
	}

	edgeClusterProvisioner, err := service.edgeClusterFactoryService.Create(ctx, request.EdgeCluster.ClusterType)
	if err != nil {
		return nil, NewUnknownErrorWithError("Failed to create egde cluster provisioner", err)
	}

	_, err = edgeClusterProvisioner.UpdateProvisionWithRetry(
		ctx,
		&edgeClusterTypes.UpdateProvisionRequest{
			EdgeClusterID: request.EdgeClusterID,
			ClusterSecret: request.EdgeCluster.ClusterSecret,
		})

	if err != nil {
		return nil, NewUnknownErrorWithError("Failed to update the existing edge cluster provision", err)
	}

	response := &UpdateEdgeClusterResponse{
		EdgeCluster: repositoryResponse.EdgeCluster,
		Cursor:      repositoryResponse.Cursor,
	}

	if provisionDetailsReponse, err := edgeClusterProvisioner.GetProvisionDetails(
		ctx,
		&edgeClusterTypes.GetProvisionDetailsRequest{EdgeClusterID: request.EdgeClusterID}); err == nil {
		response.ProvisionDetails = provisionDetailsReponse.ProvisionDetails
	}

	return response, nil
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

	edgeClusterProvisioner, err := service.edgeClusterFactoryService.Create(ctx, request.EdgeCluster.ClusterType)
	if err != nil {
		return nil, NewUnknownErrorWithError("Failed to create egde cluster provisioner", err)
	}

	_, err = edgeClusterProvisioner.DeleteProvision(
		ctx,
		&edgeClusterTypes.DeleteProvisionRequest{
			EdgeClusterID: request.EdgeClusterID,
		})

	if err != nil {
		return nil, NewUnknownErrorWithError("Failed to delete the existing edge cluster provision", err)
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

	for idx, edgeCluster := range result.EdgeClusters {
		edgeClusterProvisioner, err := service.edgeClusterFactoryService.Create(ctx, edgeCluster.EdgeCluster.ClusterType)
		if err != nil {
			return nil, NewUnknownErrorWithError("Failed to create egde cluster provisioner", err)
		}

		if provisionDetailsReponse, err := edgeClusterProvisioner.GetProvisionDetails(
			ctx,
			&edgeClusterTypes.GetProvisionDetailsRequest{EdgeClusterID: edgeCluster.EdgeClusterID}); err == nil {
			result.EdgeClusters[idx].ProvisionDetails = provisionDetailsReponse.ProvisionDetails
		}
	}

	return &SearchResponse{
		HasPreviousPage: result.HasPreviousPage,
		HasNextPage:     result.HasNextPage,
		TotalCount:      result.TotalCount,
		EdgeClusters:    result.EdgeClusters,
	}, nil
}

// ListEdgeClusterNodes lists an existing edge cluster nodes details
// ctx: Mandatory The reference to the context
// request: Mandatory. The request to list an existing edge cluster nodes details
// Returns an existing edge cluster nodes details or error if something goes wrong.
func (service *businessService) ListEdgeClusterNodes(
	ctx context.Context,
	request *ListEdgeClusterNodesRequest) (*ListEdgeClusterNodesResponse, error) {
	repositoryResponse, err := service.repositoryService.ReadEdgeCluster(ctx, &repository.ReadEdgeClusterRequest{
		EdgeClusterID: request.EdgeClusterID,
	})

	if err != nil {
		return &ListEdgeClusterNodesResponse{
			Err: mapRepositoryError(err, request.EdgeClusterID),
		}, nil
	}

	edgeClusterProvisioner, err := service.edgeClusterFactoryService.Create(ctx, repositoryResponse.EdgeCluster.ClusterType)
	if err != nil {
		return nil, NewUnknownErrorWithError("Failed to create egde cluster provisioner", err)
	}

	response, err := edgeClusterProvisioner.ListEdgeClusterNodes(
		ctx,
		&edgeClusterTypes.ListEdgeClusterNodesRequest{EdgeClusterID: request.EdgeClusterID})

	if err != nil {
		return &ListEdgeClusterNodesResponse{
			Err: mapRepositoryError(err, request.EdgeClusterID),
		}, nil
	}

	return &ListEdgeClusterNodesResponse{
		Nodes: response.Nodes,
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
