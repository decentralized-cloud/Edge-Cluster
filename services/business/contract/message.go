// Package contract defines the different EdgeCluster business contracts
package contract

import "github.com/decentralized-cloud/edge-cluster/models"

// CreateEdgeClusterRequest contains the request to create a new edge cluster
type CreateEdgeClusterRequest struct {
	EdgeCluster models.EdgeCluster
}

// CreateEdgeClusterResponse contains the result of creating a new edge cluster
type CreateEdgeClusterResponse struct {
	EdgeClusterID string
	Err           error
}

// ReadEdgeClusterRequest contains the request to read an existing edge cluster
type ReadEdgeClusterRequest struct {
	EdgeClusterID string
}

// ReadEdgeClusterResponse contains the result of reading an existing edge cluster
type ReadEdgeClusterResponse struct {
	EdgeCluster models.EdgeCluster
	Err         error
}

// UpdateEdgeClusterRequest contains the request to update an existing edge cluster
type UpdateEdgeClusterRequest struct {
	EdgeClusterID string
	EdgeCluster   models.EdgeCluster
}

// UpdateEdgeClusterResponse contains the result of updating an existing edge cluster
type UpdateEdgeClusterResponse struct {
	Err error
}

// DeleteEdgeClusterRequest contains the request to delete an existing edge cluster
type DeleteEdgeClusterRequest struct {
	EdgeClusterID string
}

// DeleteEdgeClusterResponse contains the result of deleting an existing edge cluster
type DeleteEdgeClusterResponse struct {
	Err error
}
