// Package models defines the different object models used in EdgeCluster
package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// Validate validates the EdgeCluster and return error if the validation failes
// Returns error if validation failes
func (val EdgeCluster) Validate() error {
	return validation.ValidateStruct(&val,
		// ProjectID cannot be empty
		validation.Field(&val.ProjectID, validation.Required),
		// Name cannot be empty
		validation.Field(&val.Name, validation.Required),
		// clusterSecret cannot be empty
		validation.Field(&val.ClusterSecret, validation.Required),
	)
}
