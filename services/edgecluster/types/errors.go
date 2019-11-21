// Package types defines the contracts that are used to provision a supported edge cluster and managing them
package types

import (
	"fmt"

	"github.com/decentralized-cloud/edge-cluster/models"
)

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

	return fmt.Sprintf("Unknown error occurred. Error message: %s. Error: %v", e.Message, e.Err)
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

// EdgeClusterTypeNotSupportedError indicates that the edge cluster type is not supported
type EdgeClusterTypeNotSupportedError struct {
	ClusterType models.ClusterType
	Err         error
}

// Error returns message for the EdgeClusterTypeNotSupportedError error type
// Returns the error nessage
func (e EdgeClusterTypeNotSupportedError) Error() string {
	if e.Err == nil {
		return fmt.Sprintf("Edge Cluster Type is not supported. Clustertype: %v.", e.ClusterType)
	}

	return fmt.Sprintf("Edge Cluster Type is not supported. ClusterType: %v. Error: %v", e.ClusterType, e.Err)
}

// Unwrap returns the err if provided through NewEdgeClusterTypeNotSupportedErrorWithError function, otherwise returns nil
func (e EdgeClusterTypeNotSupportedError) Unwrap() error {
	return e.Err
}

// IsEdgeClusterTypeNotSupportedError indicates whether the error is of type EdgeClusterTypeNotSupportedError
func IsEdgeClusterTypeNotSupportedError(err error) bool {
	_, ok := err.(EdgeClusterTypeNotSupportedError)

	return ok
}

// NewEdgeClusterTypeNotSupportedError creates a new EdgeClusterTypeNotSupportedError error
// clusterType: Mandatory. The edge cluster type that is not supported
func NewEdgeClusterTypeNotSupportedError(clusterType models.ClusterType) error {
	return EdgeClusterTypeNotSupportedError{
		ClusterType: clusterType,
	}
}

// NewEdgeClusterTypeNotSupportedErrorWithError creates a new EdgeClusterTypeNotSupportedError error
// clusterType: Mandatory. The edge cluster type that is not supported
func NewEdgeClusterTypeNotSupportedErrorWithError(clusterType models.ClusterType, err error) error {
	return EdgeClusterTypeNotSupportedError{
		ClusterType: clusterType,
		Err:         err,
	}
}
