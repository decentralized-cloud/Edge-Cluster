// Package business implements different business services required by the edge-cluster service
package business

import (
	"github.com/decentralized-cloud/edge-cluster/models"
	"github.com/micro-business/go-core/common"
)

// CreateEdgeClusterRequest contains the request to create a new edge cluster
type CreateEdgeClusterRequest struct {
	EdgeCluster models.EdgeCluster
}

// CreateEdgeClusterResponse contains the result of creating a new edge cluster
type CreateEdgeClusterResponse struct {
	Err              error
	EdgeClusterID    string
	EdgeCluster      models.EdgeCluster
	Cursor           string
	ProvisionDetails models.ProvisionDetails
}

// ReadEdgeClusterRequest contains the request to read an existing edge cluster
type ReadEdgeClusterRequest struct {
	EdgeClusterID string
}

// ReadEdgeClusterResponse contains the result of reading an existing edge cluster
type ReadEdgeClusterResponse struct {
	Err              error
	EdgeCluster      models.EdgeCluster
	ProvisionDetails models.ProvisionDetails
}

// UpdateEdgeClusterRequest contains the request to update an existing edge cluster
type UpdateEdgeClusterRequest struct {
	EdgeClusterID string
	EdgeCluster   models.EdgeCluster
}

// UpdateEdgeClusterResponse contains the result of updating an existing edge cluster
type UpdateEdgeClusterResponse struct {
	Err              error
	EdgeCluster      models.EdgeCluster
	Cursor           string
	ProvisionDetails models.ProvisionDetails
}

// DeleteEdgeClusterRequest contains the request to delete an existing edge cluster
type DeleteEdgeClusterRequest struct {
	EdgeClusterID string
	EdgeCluster   models.EdgeCluster
}

// DeleteEdgeClusterResponse contains the result of deleting an existing edge cluster
type DeleteEdgeClusterResponse struct {
	Err error
}

// SearchRequest contains the filter criteria to look for existing edge clusters
type SearchRequest struct {
	Pagination     common.Pagination
	SortingOptions []common.SortingOptionPair
	EdgeClusterIDs []string
	TenantIDs      []string
}

// SearchResponse contains the list of the edge clusters that matched the result
type SearchResponse struct {
	Err             error
	HasPreviousPage bool
	HasNextPage     bool
	TotalCount      int64
	EdgeClusters    []models.EdgeClusterWithCursor
}
