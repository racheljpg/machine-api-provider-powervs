// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NetworkIDs network i ds
//
// swagger:model NetworkIDs
type NetworkIDs struct {

	// an array of network IDs
	// Example: ["7f950c76-8582-11qeb-8dcd-0242ac143","7f950c76-8582-11veb-8dcd-0242ac153","7f950c76-8582-11deb-8dcd-0242ac163"]
	NetworkIDs []string `json:"networkIDs"`
}

// Validate validates this network i ds
func (m *NetworkIDs) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this network i ds based on context it is used
func (m *NetworkIDs) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NetworkIDs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkIDs) UnmarshalBinary(b []byte) error {
	var res NetworkIDs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}