// Package types defines the contracts that are used to provision a supported edge cluster and managing them
package types

// NewProvisionRequest contains the request to provision a new supported edge cluser
type NewProvisionRequest struct {
	Name                   string
	NameSpace              string
	ClusterPublicIPAddress string
}

// NewProvisionResponse contains the result of provisioning a new supported edge cliuster
type NewProvisionResponse struct {
}

// UpdateProvisionResponse contains the result of updating an existing supported edge cliuster
type UpdateProvisionResponse struct {
}

// UpdateProvisionRequest contains the request to update an existing supported edge cluser
type UpdateProvisionRequest struct {
	Name                   string
	NameSpace              string
	ClusterPublicIPAddress string
	K3SClusterSecret       string
}

// DeleteProvisionRequest contains the request to delete an existing provision
type DeleteProvisionRequest struct {
	Name      string
	NameSpace string
}

// DeleteProvisionResponse contains the result of deleting an existing edge cluster provision
type DeleteProvisionResponse struct {
}
