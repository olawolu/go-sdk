// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1CreatePrivateKeysResultV2 v1 create private keys result v2
//
// swagger:model v1CreatePrivateKeysResultV2
type V1CreatePrivateKeysResultV2 struct {

	// A list of Private Key IDs and addresses.
	// Required: true
	PrivateKeys []*V1PrivateKeyResult `json:"privateKeys"`
}

// Validate validates this v1 create private keys result v2
func (m *V1CreatePrivateKeysResultV2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrivateKeys(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1CreatePrivateKeysResultV2) validatePrivateKeys(formats strfmt.Registry) error {

	if err := validate.Required("privateKeys", "body", m.PrivateKeys); err != nil {
		return err
	}

	for i := 0; i < len(m.PrivateKeys); i++ {
		if swag.IsZero(m.PrivateKeys[i]) { // not required
			continue
		}

		if m.PrivateKeys[i] != nil {
			if err := m.PrivateKeys[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("privateKeys" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("privateKeys" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1 create private keys result v2 based on the context it is used
func (m *V1CreatePrivateKeysResultV2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePrivateKeys(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1CreatePrivateKeysResultV2) contextValidatePrivateKeys(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PrivateKeys); i++ {

		if m.PrivateKeys[i] != nil {
			if err := m.PrivateKeys[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("privateKeys" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("privateKeys" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1CreatePrivateKeysResultV2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CreatePrivateKeysResultV2) UnmarshalBinary(b []byte) error {
	var res V1CreatePrivateKeysResultV2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
