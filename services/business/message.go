// Package business implements different business services required by the edge-cluster service
package business

import (
	"github.com/decentralized-cloud/edge-cluster/models"
	"github.com/micro-business/go-core/common"
)

// CreateEdgeClusterRequest contains the request to create a new edge cluster
type CreateEdgeClusterRequest struct {
	UserEmail   string
	EdgeCluster models.EdgeCluster
}

// CreateEdgeClusterResponse contains the result of creating a new edge cluster
type CreateEdgeClusterResponse struct {
	Err           error
	EdgeClusterID string
	EdgeCluster   models.EdgeCluster
	Cursor        string
}

// ReadEdgeClusterRequest contains the request to read an existing edge cluster
type ReadEdgeClusterRequest struct {
	UserEmail     string
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
	UserEmail     string
	EdgeClusterID string
	EdgeCluster   models.EdgeCluster
}

// UpdateEdgeClusterResponse contains the result of updating an existing edge cluster
type UpdateEdgeClusterResponse struct {
	Err         error
	EdgeCluster models.EdgeCluster
	Cursor      string
}

// DeleteEdgeClusterRequest contains the request to delete an existing edge cluster
type DeleteEdgeClusterRequest struct {
	UserEmail     string
	EdgeClusterID string
	EdgeCluster   models.EdgeCluster
}

// DeleteEdgeClusterResponse contains the result of deleting an existing edge cluster
type DeleteEdgeClusterResponse struct {
	Err error
}

// ListEdgeClustersRequest contains the filter criteria to look for existing edge clusters
type ListEdgeClustersRequest struct {
	UserEmail      string
	Pagination     common.Pagination
	SortingOptions []common.SortingOptionPair
	EdgeClusterIDs []string
	ProjectIDs     []string
}

// ListEdgeClustersResponse contains the list of the edge clusters that matched the result
type ListEdgeClustersResponse struct {
	Err             error
	HasPreviousPage bool
	HasNextPage     bool
	TotalCount      int64
	EdgeClusters    []models.EdgeClusterWithCursor
}

// ListEdgeClusterNodesRequest contains the request to list an existing edge cluster nodes details
type ListEdgeClusterNodesRequest struct {
	UserEmail     string
	EdgeClusterID string
}

// ListEdgeClusterNodesResponse contains the result of listing an existing edge cluster nodes details
type ListEdgeClusterNodesResponse struct {
	Err   error
	Nodes []models.EdgeClusterNode
}

// ListEdgeClusterPodsRequest contains the request to list an existing edge cluster pods details
type ListEdgeClusterPodsRequest struct {
	UserEmail     string
	EdgeClusterID string
	Namespace     string
	NodeName      string
}

// ListEdgeClusterPodsResponse contains the result of listing an existing edge cluster pods details
type ListEdgeClusterPodsResponse struct {
	Err  error
	Pods []models.EdgeClusterPod
}

// ListEdgeClusterServicesRequest contains the request to list an existing edge cluster services details
type ListEdgeClusterServicesRequest struct {
	UserEmail     string
	EdgeClusterID string
	Namespace     string
}

// ListEdgeClusterServicesResponse contains the result of listing an existing edge cluster services details
type ListEdgeClusterServicesResponse struct {
	Err      error
	Services []models.EdgeClusterService
}
