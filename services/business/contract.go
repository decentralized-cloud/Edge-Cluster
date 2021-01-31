// Package business implements different business services required by the edge-cluster service
package business

import "context"

// BusinessContract declares the service that can create new edge cluster, read, update
// and delete existing edge clusters.
type BusinessContract interface {
	// CreateEdgeCluster creates a new edge cluster.
	// context: Mandatory The reference to the context
	// request: Mandatory. The request to create a new edge cluster
	// Returns either the result of creating new edge cluster or error if something goes wrong.
	CreateEdgeCluster(
		ctx context.Context,
		request *CreateEdgeClusterRequest) (*CreateEdgeClusterResponse, error)

	// ReadEdgeCluster read an existing edge cluster
	// context: Mandatory The reference to the context
	// request: Mandatory. The request to read an existing edge cluster
	// Returns either the result of reading an exiting edge cluster or error if something goes wrong.
	ReadEdgeCluster(
		ctx context.Context,
		request *ReadEdgeClusterRequest) (*ReadEdgeClusterResponse, error)

	// UpdateEdgeCluster update an existing edge cluster
	// context: Mandatory The reference to the context
	// request: Mandatory. The request to update an existing edge cluster
	// Returns either the result of updateing an exiting edge cluster or error if something goes wrong.
	UpdateEdgeCluster(
		ctx context.Context,
		request *UpdateEdgeClusterRequest) (*UpdateEdgeClusterResponse, error)

	// DeleteEdgeCluster delete an existing edge cluster
	// context: Mandatory The reference to the context
	// request: Mandatory. The request to delete an existing edge cluster
	// Returns either the result of deleting an exiting edge cluster or error if something goes wrong.
	DeleteEdgeCluster(
		ctx context.Context,
		request *DeleteEdgeClusterRequest) (*DeleteEdgeClusterResponse, error)

	// Search returns the list of edge clusters that matched the criteria
	// ctx: Mandatory The reference to the context
	// request: Mandatory. The request contains the search criteria
	// Returns the list of edge clusters that matched the criteria
	Search(
		ctx context.Context,
		request *SearchRequest) (*SearchResponse, error)

	// ListEdgeClusterNodes lists an existing edge cluster nodes details
	// ctx: Mandatory The reference to the context
	// request: Mandatory. The request to list an existing edge cluster nodes details
	// Returns an existing edge cluster nodes details or error if something goes wrong.
	ListEdgeClusterNodes(
		ctx context.Context,
		request *ListEdgeClusterNodesRequest) (*ListEdgeClusterNodesResponse, error)
}
