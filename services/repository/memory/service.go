// Package memory implements im-memory repository services
package memory

import (
	"context"
	"sort"

	"github.com/decentralized-cloud/edge-cluster/models"
	"github.com/decentralized-cloud/edge-cluster/services/repository"
	"github.com/lucsky/cuid"
	"github.com/micro-business/go-core/common"
	"github.com/thoas/go-funk"
)

type repositoryService struct {
	edgeClusters map[string]models.EdgeCluster
}

// NewRepositoryService creates new instance of the RepositoryService, setting up all dependencies and returns the instance
// Returns the new service or error if something goes wrong
func NewRepositoryService() (repository.RepositoryContract, error) {
	return &repositoryService{
		edgeClusters: make(map[string]models.EdgeCluster),
	}, nil
}

// CreateEdgeCluster creates a new edge cluster.
// context: Optional The reference to the context
// request: Mandatory. The request to create a new edge cluster
// Returns either the result of creating new edge cluster or error if something goes wrong.
func (service *repositoryService) CreateEdgeCluster(
	ctx context.Context,
	request *repository.CreateEdgeClusterRequest) (*repository.CreateEdgeClusterResponse, error) {
	edgeClusterID := cuid.New()
	service.edgeClusters[edgeClusterID] = request.EdgeCluster

	return &repository.CreateEdgeClusterResponse{
		EdgeClusterID: edgeClusterID,
		EdgeCluster:   request.EdgeCluster,
	}, nil
}

// ReadEdgeCluster read an existing edge cluster
// context: Optional The reference to the context
// request: Mandatory. The request to read an existing edge cluster
// Returns either the result of reading an exiting edge cluster or error if something goes wrong.
func (service *repositoryService) ReadEdgeCluster(
	ctx context.Context,
	request *repository.ReadEdgeClusterRequest) (*repository.ReadEdgeClusterResponse, error) {
	edgeCluster, ok := service.edgeClusters[request.EdgeClusterID]
	if !ok {
		return nil, repository.NewEdgeClusterNotFoundError(request.EdgeClusterID)
	}

	return &repository.ReadEdgeClusterResponse{
		EdgeCluster: edgeCluster,
	}, nil
}

// UpdateEdgeCluster update an existing edge cluster
// context: Optional The reference to the context
// request: Mandatory. The request to update an existing edge cluster
// Returns either the result of updateing an exiting edge cluster or error if something goes wrong.
func (service *repositoryService) UpdateEdgeCluster(
	ctx context.Context,
	request *repository.UpdateEdgeClusterRequest) (*repository.UpdateEdgeClusterResponse, error) {
	_, ok := service.edgeClusters[request.EdgeClusterID]
	if !ok {
		return nil, repository.NewEdgeClusterNotFoundError(request.EdgeClusterID)
	}

	service.edgeClusters[request.EdgeClusterID] = request.EdgeCluster

	return &repository.UpdateEdgeClusterResponse{
		EdgeCluster: request.EdgeCluster,
	}, nil
}

// DeleteEdgeCluster delete an existing edge cluster
// context: Optional The reference to the context
// request: Mandatory. The request to delete an existing edge cluster
// Returns either the result of deleting an exiting edge cluster or error if something goes wrong.
func (service *repositoryService) DeleteEdgeCluster(
	ctx context.Context,
	request *repository.DeleteEdgeClusterRequest) (*repository.DeleteEdgeClusterResponse, error) {
	_, ok := service.edgeClusters[request.EdgeClusterID]
	if !ok {
		return nil, repository.NewEdgeClusterNotFoundError(request.EdgeClusterID)
	}

	delete(service.edgeClusters, request.EdgeClusterID)

	return &repository.DeleteEdgeClusterResponse{}, nil
}

// Search returns the list of edge clusters that matched the criteria
// ctx: Mandatory The reference to the context
// request: Mandatory. The request contains the search criteria
// Returns the list of edge clusters that matched the criteria
func (service *repositoryService) Search(
	ctx context.Context,
	request *repository.SearchRequest) (*repository.SearchResponse, error) {
	response := &repository.SearchResponse{
		HasPreviousPage: false,
		HasNextPage:     false,
	}

	edgeClustersWithCursor := funk.Map(service.edgeClusters, func(edgeClusterID string, edgeCluster models.EdgeCluster) models.EdgeClusterWithCursor {
		return models.EdgeClusterWithCursor{
			EdgeClusterID: edgeClusterID,
			EdgeCluster:   edgeCluster,
			Cursor:        "Not implemented",
		}
	})

	if len(request.EdgeClusterIDs) > 0 {
		edgeClustersWithCursor = funk.Filter(edgeClustersWithCursor, func(edgeClusterWithCursor models.EdgeClusterWithCursor) bool {
			return funk.Contains(request.EdgeClusterIDs, edgeClusterWithCursor.EdgeClusterID)
		})
	}

	response.EdgeClusters = edgeClustersWithCursor.([]models.EdgeClusterWithCursor)

	// Default sorting is acsending if not provided, aslo as we only have one field currenrly stored againsst a edge cluster, we are ignroing the provided field name to sort on
	sortingDirection := common.Ascending
	if len(request.SortingOptions) > 0 {
		sortingDirection = request.SortingOptions[0].Direction
	}

	sort.Slice(response.EdgeClusters, func(i, j int) bool {
		if sortingDirection == common.Ascending {
			return response.EdgeClusters[i].EdgeCluster.Name < response.EdgeClusters[j].EdgeCluster.Name
		}

		return response.EdgeClusters[i].EdgeCluster.Name > response.EdgeClusters[j].EdgeCluster.Name
	})

	return response, nil
}
