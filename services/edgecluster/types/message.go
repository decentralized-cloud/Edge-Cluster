// Package types defines the contracts that are used to provision a supported edge cluster and managing them
package types

// NewProvisionRequest contains the request to provision a new supported edge cluser
type NewProvisionRequest struct {
	EdgeClusterID string
	ClusterSecret string
}

// NewProvisionResponse contains the result of provisioning a new supported edge cliuster
type NewProvisionResponse struct {
}

// UpdateProvisionRequest contains the request to update an existing supported edge cluser
type UpdateProvisionRequest struct {
	EdgeClusterID string
	ClusterSecret string
}

// UpdateProvisionResponse contains the result of updating an existing supported edge cliuster
type UpdateProvisionResponse struct {
}

// DeleteProvisionRequest contains the request to delete an existing provision
type DeleteProvisionRequest struct {
	EdgeClusterID string
}

// DeleteProvisionResponse contains the result of deleting an existing edge cluster provision
type DeleteProvisionResponse struct {
}
