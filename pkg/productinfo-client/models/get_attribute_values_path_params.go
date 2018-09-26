// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// GetAttributeValuesPathParams GetAttributeValuesPathParams is a placeholder for the get attribute values route's path parameters
// swagger:model GetAttributeValuesPathParams
type GetAttributeValuesPathParams struct {

	// in:path
	Attribute string `json:"attribute,omitempty"`

	// in:path
	Provider string `json:"provider,omitempty"`

	// in:path
	Region string `json:"region,omitempty"`

	// in:path
	Service string `json:"service,omitempty"`
}

// Validate validates this get attribute values path params
func (m *GetAttributeValuesPathParams) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetAttributeValuesPathParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetAttributeValuesPathParams) UnmarshalBinary(b []byte) error {
	var res GetAttributeValuesPathParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}