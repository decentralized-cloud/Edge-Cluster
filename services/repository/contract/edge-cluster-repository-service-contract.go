// Package contract defines the different EdgeCluster repository contracts
package contract

import "context"

// EdgeClusterRepositoryServiceContract declares the repository service that can create new edge cluster, read, update
// and delete existing edge clusters.
type EdgeClusterRepositoryServiceContract interface {
	// CreateEdgeCluster creates a new edge cluster.
	// context: Mandatory The reference to the context
	// request: Mandatory. The request to create a new edge cluster
	// Returns either the result of creating new edge cluster or error if something goes wrong.
	CreateEdgeCluster(
		ctx context.Context,
		request *CreateEdgeClusterRequest) (*CreateEdgeClusterResponse, error)

	// ReadEdgeCluster read an existing edge cluster
	// context: Mandatory The reference to the context
	// request: Mandatory. The request to read an esiting edge cluster
	// Returns either the result of reading an exiting edge cluster or error if something goes wrong.
	ReadEdgeCluster(
		ctx context.Context,
		request *ReadEdgeClusterRequest) (*ReadEdgeClusterResponse, error)

	// UpdateEdgeCluster update an existing edge cluster
	// context: Mandatory The reference to the context
	// request: Mandatory. The request to update an esiting edge cluster
	// Returns either the result of updateing an exiting edge cluster or error if something goes wrong.
	UpdateEdgeCluster(
		ctx context.Context,
		request *UpdateEdgeClusterRequest) (*UpdateEdgeClusterResponse, error)

	// DeleteEdgeCluster delete an existing edge cluster
	// context: Mandatory The reference to the context
	// request: Mandatory. The request to delete an esiting edge cluster
	// Returns either the result of deleting an exiting edge cluster or error if something goes wrong.
	DeleteEdgeCluster(
		ctx context.Context,
		request *DeleteEdgeClusterRequest) (*DeleteEdgeClusterResponse, error)
}
