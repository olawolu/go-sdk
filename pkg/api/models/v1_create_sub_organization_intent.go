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

// V1CreateSubOrganizationIntent v1 create sub organization intent
//
// swagger:model v1CreateSubOrganizationIntent
type V1CreateSubOrganizationIntent struct {

	// @inject_tag: validate:"omitempty,tk_label,tk_label_length"
	//
	// Name for this sub-organization
	// Required: true
	Name *string `json:"name"`

	// root authenticator
	// Required: true
	RootAuthenticator *V1AuthenticatorParamsV2 `json:"rootAuthenticator"`
}

// Validate validates this v1 create sub organization intent
func (m *V1CreateSubOrganizationIntent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRootAuthenticator(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1CreateSubOrganizationIntent) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *V1CreateSubOrganizationIntent) validateRootAuthenticator(formats strfmt.Registry) error {

	if err := validate.Required("rootAuthenticator", "body", m.RootAuthenticator); err != nil {
		return err
	}

	if m.RootAuthenticator != nil {
		if err := m.RootAuthenticator.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rootAuthenticator")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rootAuthenticator")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 create sub organization intent based on the context it is used
func (m *V1CreateSubOrganizationIntent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRootAuthenticator(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1CreateSubOrganizationIntent) contextValidateRootAuthenticator(ctx context.Context, formats strfmt.Registry) error {

	if m.RootAuthenticator != nil {
		if err := m.RootAuthenticator.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rootAuthenticator")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rootAuthenticator")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1CreateSubOrganizationIntent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CreateSubOrganizationIntent) UnmarshalBinary(b []byte) error {
	var res V1CreateSubOrganizationIntent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
