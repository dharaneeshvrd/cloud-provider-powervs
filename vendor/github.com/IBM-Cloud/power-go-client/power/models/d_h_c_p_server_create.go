// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DHCPServerCreate d h c p server create
//
// swagger:model DHCPServerCreate
type DHCPServerCreate struct {

	// Optional cloud connection uuid to connect with DHCP private network
	CloudConnectionID *string `json:"cloudConnectionID,omitempty"`
}

// Validate validates this d h c p server create
func (m *DHCPServerCreate) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this d h c p server create based on context it is used
func (m *DHCPServerCreate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DHCPServerCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DHCPServerCreate) UnmarshalBinary(b []byte) error {
	var res DHCPServerCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
