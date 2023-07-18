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

// V1AcceptInvitationIntentV2 v1 accept invitation intent v2
//
// swagger:model v1AcceptInvitationIntentV2
type V1AcceptInvitationIntentV2 struct {

	// authenticator
	// Required: true
	Authenticator *V1AuthenticatorParamsV2 `json:"authenticator"`

	// @inject_tag: validate:"required,uuid"
	//
	// Unique identifier for a given Invitation object.
	// Required: true
	InvitationID *string `json:"invitationId"`

	// @inject_tag: validate:"required,uuid"
	//
	// Unique identifier for a given User.
	// Required: true
	UserID *string `json:"userId"`
}

// Validate validates this v1 accept invitation intent v2
func (m *V1AcceptInvitationIntentV2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthenticator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInvitationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AcceptInvitationIntentV2) validateAuthenticator(formats strfmt.Registry) error {

	if err := validate.Required("authenticator", "body", m.Authenticator); err != nil {
		return err
	}

	if m.Authenticator != nil {
		if err := m.Authenticator.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authenticator")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("authenticator")
			}
			return err
		}
	}

	return nil
}

func (m *V1AcceptInvitationIntentV2) validateInvitationID(formats strfmt.Registry) error {

	if err := validate.Required("invitationId", "body", m.InvitationID); err != nil {
		return err
	}

	return nil
}

func (m *V1AcceptInvitationIntentV2) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("userId", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1 accept invitation intent v2 based on the context it is used
func (m *V1AcceptInvitationIntentV2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuthenticator(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AcceptInvitationIntentV2) contextValidateAuthenticator(ctx context.Context, formats strfmt.Registry) error {

	if m.Authenticator != nil {
		if err := m.Authenticator.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authenticator")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("authenticator")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1AcceptInvitationIntentV2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AcceptInvitationIntentV2) UnmarshalBinary(b []byte) error {
	var res V1AcceptInvitationIntentV2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
