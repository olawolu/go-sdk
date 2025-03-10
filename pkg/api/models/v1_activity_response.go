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

// V1ActivityResponse v1 activity response
//
// swagger:model v1ActivityResponse
type V1ActivityResponse struct {

	// An action that can that can be taken within the Turnkey infrastructure.
	// Required: true
	Activity *V1Activity `json:"activity"`
}

// Validate validates this v1 activity response
func (m *V1ActivityResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActivity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ActivityResponse) validateActivity(formats strfmt.Registry) error {

	if err := validate.Required("activity", "body", m.Activity); err != nil {
		return err
	}

	if m.Activity != nil {
		if err := m.Activity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("activity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("activity")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 activity response based on the context it is used
func (m *V1ActivityResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActivity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ActivityResponse) contextValidateActivity(ctx context.Context, formats strfmt.Registry) error {

	if m.Activity != nil {
		if err := m.Activity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("activity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("activity")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ActivityResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ActivityResponse) UnmarshalBinary(b []byte) error {
	var res V1ActivityResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
