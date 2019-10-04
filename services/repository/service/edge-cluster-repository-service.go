// Package service implements the different EdgeCluster repository services
package service

import (
	"context"

	"github.com/decentralized-cloud/edge-cluster/models"
	"github.com/decentralized-cloud/edge-cluster/services/repository/contract"

	"github.com/lucsky/cuid"
)

var tenants map[string]map[string]models.EdgeCluster

type edgeClusterRepositoryService struct {
}

func init() {
	tenants = make(map[string]map[string]models.EdgeCluster)
}

// NewEdgeClusterRepositoryService creates new instance of the EdgeClusterRepositoryService, setting up all dependencies and returns the instance
// Returns the new service or error if something goes wrong
func NewEdgeClusterRepositoryService() (contract.EdgeClusterRepositoryServiceContract, error) {
	return &edgeClusterRepositoryService{}, nil
}

// CreateEdgeCluster creates a new edge cluster.
// context: Optional The reference to the context
// request: Mandatory. The request to create a new edge cluster
// Returns either the result of creating new edge cluster or error if something goes wrong.
func (service *edgeClusterRepositoryService) CreateEdgeCluster(
	ctx context.Context,
	request *contract.CreateEdgeClusterRequest) (*contract.CreateEdgeClusterResponse, error) {

	edgeClusters, ok := tenants[request.TenantID]
	if !ok {
		edgeClusters = make(map[string]models.EdgeCluster)
		tenants[request.TenantID] = edgeClusters
	}

	edgeClusterID := cuid.New()
	edgeClusters[edgeClusterID] = request.EdgeCluster

	return &contract.CreateEdgeClusterResponse{
		EdgeClusterID: edgeClusterID,
	}, nil
}

// ReadEdgeCluster read an existing edge cluster
// context: Optional The reference to the context
// request: Mandatory. The request to read an existing edge cluster
// Returns either the result of reading an exiting edge cluster or error if something goes wrong.
func (service *edgeClusterRepositoryService) ReadEdgeCluster(
	ctx context.Context,
	request *contract.ReadEdgeClusterRequest) (*contract.ReadEdgeClusterResponse, error) {

	edgeClusters, ok := tenants[request.TenantID]
	if !ok {
		return nil, contract.NewTenantNotFoundError(request.TenantID)
	}

	edgeCluster, ok := edgeClusters[request.EdgeClusterID]
	if !ok {
		return nil, contract.NewEdgeClusterNotFoundError(request.EdgeClusterID)
	}

	return &contract.ReadEdgeClusterResponse{EdgeCluster: edgeCluster}, nil
}

// UpdateEdgeCluster update an existing edge cluster
// context: Optional The reference to the context
// request: Mandatory. The request to update an existing edge cluster
// Returns either the result of updateing an exiting edge cluster or error if something goes wrong.
func (service *edgeClusterRepositoryService) UpdateEdgeCluster(
	ctx context.Context,
	request *contract.UpdateEdgeClusterRequest) (*contract.UpdateEdgeClusterResponse, error) {

	edgeClusters, ok := tenants[request.TenantID]
	if !ok {
		return nil, contract.NewTenantNotFoundError(request.TenantID)
	}

	_, ok = edgeClusters[request.EdgeClusterID]
	if !ok {
		return nil, contract.NewEdgeClusterNotFoundError(request.EdgeClusterID)
	}

	edgeClusters[request.EdgeClusterID] = request.EdgeCluster

	return &contract.UpdateEdgeClusterResponse{}, nil
}

// DeleteEdgeCluster delete an existing edge cluster
// context: Optional The reference to the context
// request: Mandatory. The request to delete an existing edge cluster
// Returns either the result of deleting an exiting edge cluster or error if something goes wrong.
func (service *edgeClusterRepositoryService) DeleteEdgeCluster(
	ctx context.Context,
	request *contract.DeleteEdgeClusterRequest) (*contract.DeleteEdgeClusterResponse, error) {

	edgeClusters, ok := tenants[request.TenantID]
	if !ok {
		return nil, contract.NewTenantNotFoundError(request.TenantID)
	}

	_, ok = edgeClusters[request.EdgeClusterID]
	if !ok {
		return nil, contract.NewEdgeClusterNotFoundError(request.EdgeClusterID)
	}

	delete(edgeClusters, request.EdgeClusterID)

	return &contract.DeleteEdgeClusterResponse{}, nil
}
