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

// V1UpdatePrivateKeyTagResult v1 update private key tag result
//
// swagger:model v1UpdatePrivateKeyTagResult
type V1UpdatePrivateKeyTagResult struct {

	// Unique identifier for a given Private Key Tag.
	// Required: true
	PrivateKeyTagID *string `json:"privateKeyTagId"`
}

// Validate validates this v1 update private key tag result
func (m *V1UpdatePrivateKeyTagResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrivateKeyTagID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1UpdatePrivateKeyTagResult) validatePrivateKeyTagID(formats strfmt.Registry) error {

	if err := validate.Required("privateKeyTagId", "body", m.PrivateKeyTagID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 update private key tag result based on context it is used
func (m *V1UpdatePrivateKeyTagResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1UpdatePrivateKeyTagResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1UpdatePrivateKeyTagResult) UnmarshalBinary(b []byte) error {
	var res V1UpdatePrivateKeyTagResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
