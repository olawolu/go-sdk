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

// V1CreateInvitationsResult v1 create invitations result
//
// swagger:model v1CreateInvitationsResult
type V1CreateInvitationsResult struct {

	// A list of Invitation IDs
	// Required: true
	InvitationIds []string `json:"invitationIds"`
}

// Validate validates this v1 create invitations result
func (m *V1CreateInvitationsResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInvitationIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1CreateInvitationsResult) validateInvitationIds(formats strfmt.Registry) error {

	if err := validate.Required("invitationIds", "body", m.InvitationIds); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 create invitations result based on context it is used
func (m *V1CreateInvitationsResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1CreateInvitationsResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CreateInvitationsResult) UnmarshalBinary(b []byte) error {
	var res V1CreateInvitationsResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
