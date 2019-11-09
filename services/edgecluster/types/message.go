// Package types defines the contracts that are used to provision a supported edge cluster and managing them
package types

// NewProvisionRequest contains the request to provision a new supported edge cluser
type NewProvisionRequest struct {
	Name               string
	NameSpace          string
	ServicePort        int32
	TargetPort         int
	ContainerIpAddress string
}

// NewProvisionResponse contains the result of provisioning a new supported edge cliuster
type NewProvisionResponse struct {
}
