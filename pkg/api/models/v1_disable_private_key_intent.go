// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1DisablePrivateKeyIntent v1 disable private key intent
//
// swagger:model v1DisablePrivateKeyIntent
type V1DisablePrivateKeyIntent struct {

	// Unique identifier for a given Private Key.
	// Required: true
	PrivateKeyID *string `json:"privateKeyId"`
}

// Validate validates this v1 disable private key intent
func (m *V1DisablePrivateKeyIntent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrivateKeyID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1DisablePrivateKeyIntent) validatePrivateKeyID(formats strfmt.Registry) error {

	if err := validate.Required("privateKeyId", "body", m.PrivateKeyID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 disable private key intent based on context it is used
func (m *V1DisablePrivateKeyIntent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1DisablePrivateKeyIntent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1DisablePrivateKeyIntent) UnmarshalBinary(b []byte) error {
	var res V1DisablePrivateKeyIntent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
