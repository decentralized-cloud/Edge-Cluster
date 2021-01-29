// Package types defines the contracts that are used to provision a supported edge cluster and managing them
package types

import v1 "k8s.io/api/core/v1"

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
	Ingress []v1.LoadBalancerIngress
}
