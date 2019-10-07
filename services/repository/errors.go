// Package repository implements different repository services required by the edge-cluster service
package repository

import "fmt"

// UnknownError indicates that an unknown error has happened<Paste>
type UnknownError struct {
	Message string
	Err     error
}

// Error returns message for the UnknownError error type
// Returns the error nessage
func (e UnknownError) Error() string {
	if e.Err == nil {
		return fmt.Sprintf("Unknown error occurred. Error message: %s.", e.Message)
	}

	return fmt.Sprintf("Unknown error occurred. Error message: %s. Error: %s", e.Message, e.Err.Error())
}

// Unwrap returns the err if provided through NewUnknownErrorWithError function, otherwise returns nil
func (e UnknownError) Unwrap() error {
	return e.Err
}

// IsUnknownError indicates whether the error is of type UnknownError
func IsUnknownError(err error) bool {
	_, ok := err.(UnknownError)

	return ok
}

// NewUnknownError creates a new UnknownError error
func NewUnknownError(message string) error {
	return UnknownError{
		Message: message,
	}
}

// NewUnknownErrorWithError creates a new UnknownError error
func NewUnknownErrorWithError(message string, err error) error {
	return UnknownError{
		Message: message,
		Err:     err,
	}
}

// EdgeClusterAlreadyExistsError indicates that the edge cluster with the given information already exists
type EdgeClusterAlreadyExistsError struct {
	Err error
}

// Error returns message for the EdgeClusterAlreadyExistsError error type
// Returns the error nessage
func (e EdgeClusterAlreadyExistsError) Error() string {
	if e.Err == nil {
		return "Edge Cluster already exists."
	}

	return fmt.Sprintf("Edge Cluster already exists. Error: %s", e.Err.Error())
}

// Unwrap returns the err if provided through NewEdgeClusterAlreadyExistsErrorWithError function, otherwise returns nil
func (e EdgeClusterAlreadyExistsError) Unwrap() error {
	return e.Err
}

// IsEdgeClusterAlreadyExistsError indicates whether the error is of type EdgeClusterAlreadyExistsError
func IsEdgeClusterAlreadyExistsError(err error) bool {
	_, ok := err.(EdgeClusterAlreadyExistsError)

	return ok
}

// NewEdgeClusterAlreadyExistsError creates a new EdgeClusterAlreadyExistsError error
func NewEdgeClusterAlreadyExistsError() error {
	return EdgeClusterAlreadyExistsError{}
}

// NewEdgeClusterAlreadyExistsErrorWithError creates a new EdgeClusterAlreadyExistsError error
func NewEdgeClusterAlreadyExistsErrorWithError(err error) error {
	return EdgeClusterAlreadyExistsError{
		Err: err,
	}
}

// TenantNotFoundError indicates that the tenant with the given edge tenant ID does not exist
type TenantNotFoundError struct {
	TenantID string
	Err      error
}

// Error returns message for the TenantNotFoundError error type
// Returns the error nessage
func (e TenantNotFoundError) Error() string {
	if e.Err == nil {
		return fmt.Sprintf("Tenant not found. TenantID: %s.", e.TenantID)
	}

	return fmt.Sprintf("Tenant not found. TenantID: %s. Error: %s", e.TenantID, e.Err.Error())
}

// Unwrap returns the err if provided through NewTenantNotFoundErrorWithError function, otherwise returns nil
func (e TenantNotFoundError) Unwrap() error {
	return e.Err
}

// IsTenantNotFoundError indicates whether the error is of type TenantNotFoundError
func IsTenantNotFoundError(err error) bool {
	_, ok := err.(TenantNotFoundError)

	return ok
}

// NewTenantNotFoundError creates a new TenantNotFoundError error
// tenantID: Mandatory. The tenant ID that did not match any existing tenant
func NewTenantNotFoundError(tenantID string) error {
	return TenantNotFoundError{
		TenantID: tenantID,
	}
}

// NewTenantNotFoundErrorWithError creates a new TenantNotFoundError error
// tenantID: Mandatory. The tenant ID that did not match any existing tenant
func NewTenantNotFoundErrorWithError(tenantID string, err error) error {
	return TenantNotFoundError{
		TenantID: tenantID,
		Err:      err,
	}
}

// EdgeClusterNotFoundError indicates that the edge cluster with the given edge cluster ID does not exist
type EdgeClusterNotFoundError struct {
	EdgeClusterID string
	Err           error
}

// Error returns message for the EdgeClusterNotFoundError error type
// Returns the error nessage
func (e EdgeClusterNotFoundError) Error() string {
	if e.Err == nil {
		return fmt.Sprintf("Edge Cluster not found. EdgeClusterID: %s.", e.EdgeClusterID)
	}

	return fmt.Sprintf("Edge Cluster not found. EdgeClusterID: %s. Error: %s", e.EdgeClusterID, e.Err.Error())
}

// Unwrap returns the err if provided through NewEdgeClusterNotFoundErrorWithError function, otherwise returns nil
func (e EdgeClusterNotFoundError) Unwrap() error {
	return e.Err
}

// IsEdgeClusterNotFoundError indicates whether the error is of type EdgeClusterNotFoundError
func IsEdgeClusterNotFoundError(err error) bool {
	_, ok := err.(EdgeClusterNotFoundError)

	return ok
}

// NewEdgeClusterNotFoundError creates a new EdgeClusterNotFoundError error
// edgeClusterID: Mandatory. The edge clusterID that did not match any existing edge cluster
func NewEdgeClusterNotFoundError(edgeClusterID string) error {
	return EdgeClusterNotFoundError{
		EdgeClusterID: edgeClusterID,
	}
}

// NewEdgeClusterNotFoundErrorWithError creates a new EdgeClusterNotFoundError error
// edgeClusterID: Mandatory. The edge clusterID that did not match any existing edge cluster
func NewEdgeClusterNotFoundErrorWithError(edgeClusterID string, err error) error {
	return EdgeClusterNotFoundError{
		EdgeClusterID: edgeClusterID,
		Err:           err,
	}
}
