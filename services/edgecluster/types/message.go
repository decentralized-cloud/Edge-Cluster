// Package types defines the contracts that are used to provision a supported edge cluster and managing them
package types

import "github.com/decentralized-cloud/edge-cluster/models"

// CreateProvisionRequest contains the request to provision a new supported edge cluser
type CreateProvisionRequest struct {
	EdgeClusterID string
	ClusterSecret string
}

// CreateProvisionResponse contains the result of provisioning a new supported edge cliuster
type CreateProvisionResponse struct {
}

// UpdateProvisionRequest contains the request to update an existing provision
type UpdateProvisionRequest struct {
	EdgeClusterID string
	ClusterSecret string
}

// UpdateProvisionResponse contains the result of updating an existing provision
type UpdateProvisionResponse struct {
}

// DeleteProvisionRequest contains the request to delete an existing provision
type DeleteProvisionRequest struct {
	EdgeClusterID string
}

// DeleteProvisionResponse contains the result of deleting an existing provision
type DeleteProvisionResponse struct {
}

// GetProvisionDetailsRequest contains the request to retrieve an existing provision details
type GetProvisionDetailsRequest struct {
	EdgeClusterID string
}

// GetProvisionDetailsResponse contains the result of retrieving an existing provision
type GetProvisionDetailsResponse struct {
	ProvisionDetails models.ProvisionDetails
}

// ListNodesRequest contains the request to list an existing edge cluster nodes details
type ListNodesRequest struct {
	EdgeClusterID string
}

// ListNodesResponse contains the result of listing an existing edge cluster nodes details
type ListNodesResponse struct {
	Nodes []models.EdgeClusterNode
}

// ListPodsRequest contains the request to list an existing edge cluster pods
type ListPodsRequest struct {
	EdgeClusterID string
	Namespace     string
	NodeName      string
}

// ListPodsResponse contains the result of listing an existing edge cluster pods
type ListPodsResponse struct {
	Pods []models.EdgeClusterPod
}

// ListServicesRequest contains the request to list an existing edge cluster services
type ListServicesRequest struct {
	EdgeClusterID string
	Namespace     string
}

// ListServicesResponse contains the result of listing an existing edge cluster services
type ListServicesResponse struct {
	Services []models.EdgeClusterService
}
