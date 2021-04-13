// Package business implements different business services required by the edge-cluster service
package business

import (
	"context"

	edgeClusterTypes "github.com/decentralized-cloud/edge-cluster/services/edgecluster/types"
	"github.com/decentralized-cloud/edge-cluster/services/repository"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
)

type businessService struct {
	logger                    *zap.Logger
	repositoryService         repository.RepositoryContract
	edgeClusterFactoryService edgeClusterTypes.EdgeClusterFactoryContract
}

// NewBusinessService creates new instance of the BusinessService, setting up all dependencies and returns the instance
// repositoryService: Mandatory. Reference to the repository service that can persist the edge cluster related data
// edgeClusterFactoryService: Mandatory. Reference to the factory service that can that can create different type of supported
// edge cluster provisioner
// logger: Mandatory. Reference to the logger service
// Returns the new service or error if something goes wrong
func NewBusinessService(
	logger *zap.Logger,
	repositoryService repository.RepositoryContract,
	edgeClusterFactoryService edgeClusterTypes.EdgeClusterFactoryContract) (BusinessContract, error) {
	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	if repositoryService == nil {
		return nil, commonErrors.NewArgumentNilError("repositoryService", "repositoryService is required")
	}

	if edgeClusterFactoryService == nil {
		return nil, commonErrors.NewArgumentNilError("edgeClusterFactoryService", "edgeClusterFactoryService is required")
	}

	return &businessService{
		logger:                    logger,
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
		UserEmail:   request.UserEmail,
		EdgeCluster: request.EdgeCluster,
	})

	if err != nil {
		return &CreateEdgeClusterResponse{
			Err: err,
		}, nil
	}

	edgeClusterProvisioner, err := service.edgeClusterFactoryService.Create(ctx, request.EdgeCluster.ClusterType)
	if err != nil {
		return nil, commonErrors.NewUnknownErrorWithError("failed to create egde cluster provisioner", err)
	}

	go func() {
		if _, err = edgeClusterProvisioner.CreateProvision(
			context.Background(),
			&edgeClusterTypes.CreateProvisionRequest{
				EdgeClusterID: repositoryResponse.EdgeClusterID,
				ClusterSecret: request.EdgeCluster.ClusterSecret,
			}); err != nil {

			service.logger.Error("failed to provision egde cluster", zap.Error(err))

			return
		}
	}()

	return &CreateEdgeClusterResponse{
		EdgeClusterID: repositoryResponse.EdgeClusterID,
		EdgeCluster:   repositoryResponse.EdgeCluster,
		Cursor:        repositoryResponse.Cursor,
	}, nil
}

// ReadEdgeCluster read an existing edge cluster
// context: Mandatory The reference to the context
// request: Mandatory. The request to read an existing edge cluster
// Returns either the result of reading an existing edge cluster or error if something goes wrong.
func (service *businessService) ReadEdgeCluster(
	ctx context.Context,
	request *ReadEdgeClusterRequest) (*ReadEdgeClusterResponse, error) {
	repositoryResponse, err := service.repositoryService.ReadEdgeCluster(ctx, &repository.ReadEdgeClusterRequest{
		UserEmail:     request.UserEmail,
		EdgeClusterID: request.EdgeClusterID,
	})

	if err != nil {
		return &ReadEdgeClusterResponse{
			Err: err,
		}, nil
	}

	edgeClusterProvisioner, err := service.edgeClusterFactoryService.Create(ctx, repositoryResponse.EdgeCluster.ClusterType)
	if err != nil {
		return nil, commonErrors.NewUnknownErrorWithError("failed to create egde cluster provisioner", err)
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
		UserEmail:     request.UserEmail,
		EdgeClusterID: request.EdgeClusterID,
		EdgeCluster:   request.EdgeCluster,
	})

	if err != nil {
		return &UpdateEdgeClusterResponse{
			Err: err,
		}, nil
	}

	edgeClusterProvisioner, err := service.edgeClusterFactoryService.Create(ctx, request.EdgeCluster.ClusterType)
	if err != nil {
		return nil, commonErrors.NewUnknownErrorWithError("failed to create egde cluster provisioner", err)
	}

	go func() {
		if _, err = edgeClusterProvisioner.UpdateProvisionWithRetry(
			context.Background(),
			&edgeClusterTypes.UpdateProvisionRequest{
				EdgeClusterID: request.EdgeClusterID,
				ClusterSecret: request.EdgeCluster.ClusterSecret,
			}); err != nil {
			service.logger.Error("failed to update the existing edge cluster provision", zap.Error(err))

			return
		}

	}()

	return &UpdateEdgeClusterResponse{
		EdgeCluster: repositoryResponse.EdgeCluster,
		Cursor:      repositoryResponse.Cursor,
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
		UserEmail:     request.UserEmail,
		EdgeClusterID: request.EdgeClusterID,
	})

	if err != nil {
		return &DeleteEdgeClusterResponse{
			Err: err,
		}, nil
	}

	edgeClusterProvisioner, err := service.edgeClusterFactoryService.Create(ctx, request.EdgeCluster.ClusterType)
	if err != nil {
		return nil, commonErrors.NewUnknownErrorWithError("failed to create egde cluster provisioner", err)
	}

	_, err = edgeClusterProvisioner.DeleteProvision(
		ctx,
		&edgeClusterTypes.DeleteProvisionRequest{
			EdgeClusterID: request.EdgeClusterID,
		})

	if err != nil {
		return nil, commonErrors.NewUnknownErrorWithError("failed to delete the existing edge cluster provision", err)
	}

	return &DeleteEdgeClusterResponse{}, nil
}

// ListEdgeClusters returns the list of edge clusters that matched the criteria
// ctx: Mandatory The reference to the context
// request: Mandatory. The request contains the search criteria
// Returns the list of edge clusters that matched the criteria
func (service *businessService) ListEdgeClusters(
	ctx context.Context,
	request *ListEdgeClustersRequest) (*ListEdgeClustersResponse, error) {
	result, err := service.repositoryService.ListEdgeClusters(ctx, &repository.ListEdgeClustersRequest{
		UserEmail:      request.UserEmail,
		Pagination:     request.Pagination,
		SortingOptions: request.SortingOptions,
		EdgeClusterIDs: request.EdgeClusterIDs,
		ProjectIDs:     request.ProjectIDs,
	})

	if err != nil {
		return &ListEdgeClustersResponse{
			Err: err,
		}, nil
	}

	for idx, edgeCluster := range result.EdgeClusters {
		edgeClusterProvisioner, err := service.edgeClusterFactoryService.Create(ctx, edgeCluster.EdgeCluster.ClusterType)
		if err != nil {
			return nil, commonErrors.NewUnknownErrorWithError("failed to create egde cluster provisioner", err)
		}

		if provisionDetailsReponse, err := edgeClusterProvisioner.GetProvisionDetails(
			ctx,
			&edgeClusterTypes.GetProvisionDetailsRequest{EdgeClusterID: edgeCluster.EdgeClusterID}); err == nil {
			result.EdgeClusters[idx].ProvisionDetails = provisionDetailsReponse.ProvisionDetails
		}
	}

	return &ListEdgeClustersResponse{
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
		UserEmail:     request.UserEmail,
		EdgeClusterID: request.EdgeClusterID,
	})

	if err != nil {
		return &ListEdgeClusterNodesResponse{
			Err: err,
		}, nil
	}

	edgeClusterProvisioner, err := service.edgeClusterFactoryService.Create(ctx, repositoryResponse.EdgeCluster.ClusterType)
	if err != nil {
		return nil, commonErrors.NewUnknownErrorWithError("failed to create egde cluster provisioner", err)
	}

	response, err := edgeClusterProvisioner.ListNodes(
		ctx,
		&edgeClusterTypes.ListNodesRequest{EdgeClusterID: request.EdgeClusterID})

	if err != nil {
		return &ListEdgeClusterNodesResponse{
			Err: err,
		}, nil
	}

	return &ListEdgeClusterNodesResponse{
		Nodes: response.Nodes,
	}, nil
}

// ListEdgeClusterPods lists an existing edge cluster pods details
// ctx: Mandatory The reference to the context
// request: Mandatory. The request to list an existing edge cluster pods details
// Returns an existing edge cluster pods details or error if something goes wrong.
func (service *businessService) ListEdgeClusterPods(
	ctx context.Context,
	request *ListEdgeClusterPodsRequest) (*ListEdgeClusterPodsResponse, error) {
	repositoryResponse, err := service.repositoryService.ReadEdgeCluster(ctx, &repository.ReadEdgeClusterRequest{
		UserEmail:     request.UserEmail,
		EdgeClusterID: request.EdgeClusterID,
	})

	if err != nil {
		return &ListEdgeClusterPodsResponse{
			Err: err,
		}, nil
	}

	edgeClusterProvisioner, err := service.edgeClusterFactoryService.Create(ctx, repositoryResponse.EdgeCluster.ClusterType)
	if err != nil {
		return nil, commonErrors.NewUnknownErrorWithError("failed to create egde cluster provisioner", err)
	}

	response, err := edgeClusterProvisioner.ListPods(
		ctx,
		&edgeClusterTypes.ListPodsRequest{
			EdgeClusterID: request.EdgeClusterID,
			Namespace:     request.Namespace,
			NodeName:      request.NodeName,
		})

	if err != nil {
		return &ListEdgeClusterPodsResponse{
			Err: err,
		}, nil
	}

	return &ListEdgeClusterPodsResponse{
		Pods: response.Pods,
	}, nil
}

// ListEdgeClusterServices lists an existing edge cluster services details
// ctx: Mandatory The reference to the context
// request: Mandatory. The request to list an existing edge cluster services details
// Returns an existing edge cluster services details or error if something goes wrong.
func (service *businessService) ListEdgeClusterServices(
	ctx context.Context,
	request *ListEdgeClusterServicesRequest) (*ListEdgeClusterServicesResponse, error) {
	repositoryResponse, err := service.repositoryService.ReadEdgeCluster(ctx, &repository.ReadEdgeClusterRequest{
		UserEmail:     request.UserEmail,
		EdgeClusterID: request.EdgeClusterID,
	})

	if err != nil {
		return &ListEdgeClusterServicesResponse{
			Err: err,
		}, nil
	}

	edgeClusterProvisioner, err := service.edgeClusterFactoryService.Create(ctx, repositoryResponse.EdgeCluster.ClusterType)
	if err != nil {
		return nil, commonErrors.NewUnknownErrorWithError("failed to create egde cluster provisioner", err)
	}

	response, err := edgeClusterProvisioner.ListServices(
		ctx,
		&edgeClusterTypes.ListServicesRequest{
			EdgeClusterID: request.EdgeClusterID,
			Namespace:     request.Namespace,
		})

	if err != nil {
		return &ListEdgeClusterServicesResponse{
			Err: err,
		}, nil
	}

	return &ListEdgeClusterServicesResponse{
		Services: response.Services,
	}, nil
}
