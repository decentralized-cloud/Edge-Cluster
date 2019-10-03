// Package contract defines configuration service contracts
package contract

import "fmt"

// UnknownError indicates that the edge cluster with the given information already exists
type UnknownError struct {
	errorMessage string
	message      string
}

// Error returns message for the UnknownError error type
// Returns the error nessage
func (e UnknownError) Error() string {
	return e.message
}

// NewUnknownError creates a new UnknownError error
func NewUnknownError(errorMessage string) error {
	return UnknownError{
		errorMessage: errorMessage,
		message:      fmt.Sprintf("Unknown error occurs. Error message is: %s", errorMessage),
	}
}
