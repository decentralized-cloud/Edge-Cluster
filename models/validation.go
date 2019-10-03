// Package models defines the different object models used in EdgeCluster
package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// Validate validates the EdgeCluster and return error if the validation failes
// Returns error if validation failes
func (val *EdgeCluster) Validate() error {
	return validation.ValidateStruct(val,
		// Name cannot be empty
		validation.Field(&val.Name, validation.Required),
	)
}
