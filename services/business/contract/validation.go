// Package contract defines the different EdgeCluster business contracts
package contract

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// Validate validates the CreateEdgeClusterRequest model and return error if the validation failes
// Returns error if validation failes
func (val *CreateEdgeClusterRequest) Validate() error {
	if err := validation.ValidateStruct(val,
		// TenantID cannot be empty
		validation.Field(&val.TenantID, validation.Required),
	); err != nil {
		return err
	}

	// TODO: mortezaalizadeh: 16/09/2019: Should replace following code with nested validation
	return val.EdgeCluster.Validate()
}

// Validate validates the CreateEdgeClusterResponse model and return error if the validation failes
// Returns error if validation failes
func (val *CreateEdgeClusterResponse) Validate() error {
	return validation.ValidateStruct(val,
		// EdgeClusterID cannot be empty
		validation.Field(&val.EdgeClusterID, validation.Required),
	)
}

// Validate validates the ReadEdgeClusterRequest model and return error if the validation failes
// Returns error if validation failes
func (val *ReadEdgeClusterRequest) Validate() error {
	return validation.ValidateStruct(val,
		// TenantID cannot be empty
		validation.Field(&val.TenantID, validation.Required),
		// EdgeClusterID cannot be empty
		validation.Field(&val.EdgeClusterID, validation.Required),
	)
}

// Validate validates the ReadEdgeClusterResponse model and return error if the validation failes
// Returns error if validation failes
func (val *ReadEdgeClusterResponse) Validate() error {
	return validation.ValidateStruct(val,
		// Validate EdgeCluster using its own validation rules
		validation.Field(&val.EdgeCluster),
	)
}

// Validate validates the UpdateEdgeClusterRequest model and return error if the validation failes
// Returns error if validation failes
func (val *UpdateEdgeClusterRequest) Validate() error {
	return validation.ValidateStruct(val,
		// TenantID cannot be empty
		validation.Field(&val.TenantID, validation.Required),
		// EdgeClusterID cannot be empty
		validation.Field(&val.EdgeClusterID, validation.Required),
		// Validate EdgeCluster using its own validation rules
		validation.Field(&val.EdgeCluster),
	)
}

// Validate validates the UpdateEdgeClusterResponse model and return error if the validation failes
// Returns error if validation failes
func (val *UpdateEdgeClusterResponse) Validate() error {
	return nil
}

// Validate validates the DeleteEdgeClusterRequest model and return error if the validation failes
// Returns error if validation failes
func (val *DeleteEdgeClusterRequest) Validate() error {
	return validation.ValidateStruct(val,
		// TenantID cannot be empty
		validation.Field(&val.TenantID, validation.Required),
		// EdgeClusterID cannot be empty
		validation.Field(&val.EdgeClusterID, validation.Required),
	)
}

// Validate validates the DeleteEdgeClusterResponse model and return error if the validation failes
// Returns error if validation failes
func (val *DeleteEdgeClusterResponse) Validate() error {
	return nil
}
