// Package types defines the contracts that are used to provision a supported edge cluster and managing them
package types

import (
	"context"

	"github.com/decentralized-cloud/edge-cluster/models"
)

// EdgeClusterFactoryContract defines the factory method that are used to create provisioner
// for different supported type of edge cluster (e.g. K3S)
type EdgeClusterFactoryContract interface {
	// Create instantiates a new edge cluster provisioner of a requested edge cluster type and returns
	// it to the caller.
	// ctx: Mandatory The reference to the context
	// clusterType: Mandatory. The type of edge cluster provisioner to be instantiated
	// Returns either the result of instantiating a edge cluster provisioner or error if something goes wrong.
	Create(
		ctx context.Context,
		clusterType models.ClusterType) (EdgeClusterProvisionerContract, error)
}

// EdgeClusterProvisionerContract defines the methods that are required to provision a supported
// type of edge cluster
type EdgeClusterProvisionerContract interface {
	// CreateProvision provisions a new edge cluster.
	// ctx: Mandatory The reference to the context
	// request: Mandatory. The request to provision a new edge cluster
	// Returns either the result of provisioning new edge cluster or error if something goes wrong.
	CreateProvision(
		ctx context.Context,
		request *CreateProvisionRequest) (*CreateProvisionResponse, error)

	// UpdateProvisionWithRetry updates an existing provision.
	// ctx: Mandatory The reference to the context
	// request: Mandatory. The request to update an existing provision
	// Returns either the result of updating an existing provision or error if something goes wrong.
	UpdateProvisionWithRetry(
		ctx context.Context,
		request *UpdateProvisionRequest) (*UpdateProvisionResponse, error)

	// DeleteProvision deletes an existing provision.
	// ctx: Mandatory The reference to the context
	// request: Mandatory. The request to delete an existing provision
	// Returns either the result of deleting an existing provision or error if something goes wrong.
	DeleteProvision(
		ctx context.Context,
		request *DeleteProvisionRequest) (*DeleteProvisionResponse, error)

	// GetProvisionDetails retrieves information on an existing provision.
	// ctx: Mandatory The reference to the context
	// request: Mandatory. The request to retrieve information on an existing provision
	// Returns either the result of retrieving information on an provision or error if something goes wrong.
	GetProvisionDetails(
		ctx context.Context,
		request *GetProvisionDetailsRequest) (*GetProvisionDetailsResponse, error)

	// ListEdgeClusterNodes lists an existing edge cluster nodes details
	// ctx: Mandatory The reference to the context
	// request: Mandatory. The request to list an existing edge cluster nodes details
	// Returns an existing edge cluster nodes details or error if something goes wrong.
	ListEdgeClusterNodes(
		ctx context.Context,
		request *ListEdgeClusterNodesRequest) (*ListEdgeClusterNodesResponse, error)
}
