// Package models defines the different object models used in EdgeCluster
package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// Validate validates the EdgeCluster and return error if the validation failes
// Returns error if validation failes
func (val EdgeCluster) Validate() error {
	return validation.ValidateStruct(&val,
		// TenantID cannot be empty
		validation.Field(&val.TenantID, validation.Required),
		// Name cannot be empty
		validation.Field(&val.Name, validation.Required),
		// ClusterIP cannot be empty and must follow IP v4 format
		validation.Field(&val.ClusterPublicIPAddress, validation.Required, is.IPv4),
	)
}
