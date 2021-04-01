// Package business implements different business services required by the edge-cluster service
package business

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// Validate validates the CreateEdgeClusterRequest model and return error if the validation failes
// Returns error if validation failes
func (val CreateEdgeClusterRequest) Validate() error {
	return validation.ValidateStruct(&val,
		// Email must be provided
		validation.Field(&val.UserEmail, validation.Required, is.Email),
		// Validate EdgeCluster using its own validation rules
		validation.Field(&val.EdgeCluster),
	)
}

// Validate validates the ReadEdgeClusterRequest model and return error if the validation failes
// Returns error if validation failes
func (val ReadEdgeClusterRequest) Validate() error {
	return validation.ValidateStruct(&val,
		// Email must be provided
		validation.Field(&val.UserEmail, validation.Required, is.Email),
		// EdgeClusterID cannot be empty
		validation.Field(&val.EdgeClusterID, validation.Required),
	)
}

// Validate validates the UpdateEdgeClusterRequest model and return error if the validation failes
// Returns error if validation failes
func (val UpdateEdgeClusterRequest) Validate() error {
	return validation.ValidateStruct(&val,
		// Email must be provided
		validation.Field(&val.UserEmail, validation.Required, is.Email),
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
		// Email must be provided
		validation.Field(&val.UserEmail, validation.Required, is.Email),
		// EdgeClusterID cannot be empty
		validation.Field(&val.EdgeClusterID, validation.Required),
	)
}

// Validate validates the ListEdgeClustersRequest model and return error if the validation failes
// Returns error if validation failes
func (val ListEdgeClustersRequest) Validate() error {
	return validation.ValidateStruct(&val,
		// Email must be provided
		validation.Field(&val.UserEmail, validation.Required, is.Email),
	)
}

// Validate validates the ListEdgeClusterNodesRequest model and return error if the validation failes
// Returns error if validation failes
func (val ListEdgeClusterNodesRequest) Validate() error {
	return validation.ValidateStruct(&val,
		// Email must be provided
		validation.Field(&val.UserEmail, validation.Required, is.Email),
		// EdgeClusterID cannot be empty
		validation.Field(&val.EdgeClusterID, validation.Required),
	)
}

// Validate validates the ListEdgeClusterPodsRequest model and return error if the validation failes
// Returns error if validation failes
func (val ListEdgeClusterPodsRequest) Validate() error {
	return validation.ValidateStruct(&val,
		// Email must be provided
		validation.Field(&val.UserEmail, validation.Required, is.Email),
		// EdgeClusterID cannot be empty
		validation.Field(&val.EdgeClusterID, validation.Required),
	)
}

// Validate validates the ListEdgeClusterServicesRequest model and return error if the validation failes
// Returns error if validation failes
func (val ListEdgeClusterServicesRequest) Validate() error {
	return validation.ValidateStruct(&val,
		// Email must be provided
		validation.Field(&val.UserEmail, validation.Required, is.Email),
		// EdgeClusterID cannot be empty
		validation.Field(&val.EdgeClusterID, validation.Required),
	)
}
