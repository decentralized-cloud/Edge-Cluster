// Package contract defines the contracts that provides endpoint to be used by the transport layer
package contract

import "github.com/go-kit/kit/endpoint"

// EndpointCreatorContract declares the contract that creates endpoints to create new edgeCluster,
// read, update and delete existing edgeClusters.
type EndpointCreatorContract interface {
	// CreateEdgeClusterEndpoint creates Create Edge Cluster endpoint
	// Returns the Create Edge Cluster endpoint
	CreateEdgeClusterEndpoint() endpoint.Endpoint

	// ReadEdgeClusterEndpoint creates Read Edge Cluster endpoint
	// Returns the Read Edge Cluster endpoint
	ReadEdgeClusterEndpoint() endpoint.Endpoint

	// UpdateEdgeClusterEndpoint creates Update Edge Cluster endpoint
	// Returns the Update Edge Cluster endpoint
	UpdateEdgeClusterEndpoint() endpoint.Endpoint

	// DeleteEdgeClusterEndpoint creates Delete Edge Cluster endpoint
	// Returns the Delete Edge Cluster endpoint
	DeleteEdgeClusterEndpoint() endpoint.Endpoint
}
