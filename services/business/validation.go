// Package business implements different business services required by the edge-cluster service
package business

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// Validate validates the CreateEdgeClusterRequest model and return error if the validation failes
// Returns error if validation failes
func (val CreateEdgeClusterRequest) Validate() error {
	return validation.ValidateStruct(&val,
		// Validate EdgeCluster using its own validation rules
		validation.Field(&val.EdgeCluster),
	)
}

// Validate validates the ReadEdgeClusterRequest model and return error if the validation failes
// Returns error if validation failes
func (val ReadEdgeClusterRequest) Validate() error {
	return validation.ValidateStruct(&val,
		// EdgeClusterID cannot be empty
		validation.Field(&val.EdgeClusterID, validation.Required),
	)
}

// Validate validates the UpdateEdgeClusterRequest model and return error if the validation failes
// Returns error if validation failes
func (val UpdateEdgeClusterRequest) Validate() error {
	return validation.ValidateStruct(&val,
		// EdgeClusterID cannot be empty
		validation.Field(&val.EdgeClusterID, validation.Required),
		// Validate EdgeCluster using its own validation rules
		validation.Field(&val.EdgeCluster),
	)
}

// Validate validates the DeleteEdgeClusterRequest model and return error if the validation failes
// Returns error if validation failes
func (val DeleteEdgeClusterRequest) Validate() error {
	return validation.ValidateStruct(&val,
		// EdgeClusterID cannot be empty
		validation.Field(&val.EdgeClusterID, validation.Required),
	)
}

// Validate validates the SearchRequest model and return error if the validation failes
// Returns error if validation failes
func (val SearchRequest) Validate() error {
	return nil
}
