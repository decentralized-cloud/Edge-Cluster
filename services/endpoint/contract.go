// Package endpoint implements different endpoint services required by the edge-cluster service
package endpoint

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

	// SearchEndpoint creates Search Edge Cluster endpoint
	// Returns the Search Edge Cluster endpoint
	SearchEndpoint() endpoint.Endpoint

	// ListEdgeClusterNodesEndpoint creates List Edge Cluster Nodes endpoint
	// Returns the List Edge Cluster Nodes endpoint
	ListEdgeClusterNodesEndpoint() endpoint.Endpoint

	// ListEdgeClusterPodsEndpoint creates List Edge Cluster Pods endpoint
	// Returns the List Edge Cluster Pods endpoint
	ListEdgeClusterPodsEndpoint() endpoint.Endpoint
}
