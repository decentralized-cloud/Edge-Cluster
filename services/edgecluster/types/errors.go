// Package types defines the contracts that are used to provision a supported edge cluster and managing them
package types

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

// EdgeClusterNotSupportedError indicates that the requested edge cluster is not suipported
type EdgeClusterNotSupportedError struct {
	Err             error
	EdgeClusterType EdgeClusterType
}

// Error returns message for the EdgeClusterNotSupportedError error type
// Returns the error nessage
func (e EdgeClusterNotSupportedError) Error() string {
	if e.Err == nil {
		return "Edge Cluster not supported."
	}

	return fmt.Sprintf("Edge Cluster not supported. Error: %v", e.Err)
}

// Unwrap returns the err if provided through NewEdgeClusterNotSupportedErrorWithError function, otherwise returns nil
func (e EdgeClusterNotSupportedError) Unwrap() error {
	return e.Err
}

// IsEdgeClusterNotSupportedError indicates whether the error is of type EdgeClusterNotSupportedError
func IsEdgeClusterNotSupportedError(err error) bool {
	_, ok := err.(EdgeClusterNotSupportedError)

	return ok
}

// NewEdgeClusterNotSupportedError creates a new EdgeClusterNotSupportedError error
// edgeClusterType: Mandatory. The type of the edge cluster that is not suppoorted
func NewEdgeClusterNotSupportedError(edgeClusterType EdgeClusterType) error {
	return EdgeClusterNotSupportedError{
		EdgeClusterType: edgeClusterType,
	}
}

// NewEdgeClusterNotSupportedErrorWithError creates a new EdgeClusterNotSupportedError error
// edgeClusterType: Mandatory. The type of the edge cluster that is not suppoorted
func NewEdgeClusterNotSupportedErrorWithError(
	edgeClusterType EdgeClusterType,
	err error) error {
	return EdgeClusterNotSupportedError{
		Err:             err,
		EdgeClusterType: edgeClusterType,
	}
}
