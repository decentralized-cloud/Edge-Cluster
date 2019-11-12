// Package types defines the contracts that are used to provision a supported edge cluster and managing them
package types

import "context"

type EdgeClusterType int

const (
	K3S EdgeClusterType = iota
)

// EdgeClusterFactoryContract defines the factory method that are used to create provisioner
// for different supported type of edge cluster (e.g. K3S)
type EdgeClusterFactoryContract interface {
	// Create instantiates a new edge cluster provisioner of a requested edge cluster type and returns
	// it to the caller.
	// ctx: Mandatory The reference to the context
	// edgeClusterType: Mandatory. The type of edge cluster provisioner to be instantiated
	// Returns either the result of instantiating a edge cluster provisioner or error if something goes wrong.
	Create(
		ctx context.Context,
		edgeClusterType EdgeClusterType) (EdgeClusterProvisionerContract, error)
}

// EdgeClusterProvisionerContract defines the methods that are required to provision a supported
// type of edge cluster
type EdgeClusterProvisionerContract interface {
	// NewProvision provisions a new edge cluster.
	// ctx: Mandatory The reference to the context
	// request: Mandatory. The request to provision a new edge cluster
	// Returns either the result of provisioning new edge cluster or error if something goes wrong.
	NewProvision(
		ctx context.Context,
		request *NewProvisionRequest) (*NewProvisionResponse, error)

	// UpdateProvisionWithRetry updates an existing edge cluster.
	// ctx: Mandatory The reference to the context
	// request: Mandatory. The request to update an existing edge cluster.
	// Returns either the result of updating an existing edge cluster or error if something goes wrong.
	UpdateProvisionWithRetry(
		ctx context.Context,
		request *UpdateProvisionRequest) (response *UpdateProvisionResponse, err error)

	// DeleteProvision deletes an edge cluster.
	// ctx: Mandatory The reference to the context
	// request: Mandatory. The request to delete an edge cluster
	// Returns either the result of deleting an edge cluster or error if something goes wrong.
	DeleteProvision(
		ctx context.Context,
		request *NewProvisionRequest) (response *NewProvisionResponse, err error)
}
