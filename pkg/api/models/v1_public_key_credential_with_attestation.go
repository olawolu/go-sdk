// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1PublicKeyCredentialWithAttestation v1 public key credential with attestation
//
// swagger:model v1PublicKeyCredentialWithAttestation
type V1PublicKeyCredentialWithAttestation struct {

	// authenticator attachment
	// Enum: [cross-platform platform]
	AuthenticatorAttachment *string `json:"authenticatorAttachment,omitempty"`

	// client extension results
	// Required: true
	ClientExtensionResults *V1SimpleClientExtensionResults `json:"clientExtensionResults"`

	// id
	// Required: true
	ID *string `json:"id"`

	// raw Id
	// Required: true
	RawID *string `json:"rawId"`

	// response
	// Required: true
	Response *V1AuthenticatorAttestationResponse `json:"response"`

	// type
	// Required: true
	// Enum: [public-key]
	Type *string `json:"type"`
}

// Validate validates this v1 public key credential with attestation
func (m *V1PublicKeyCredentialWithAttestation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthenticatorAttachment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientExtensionResults(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRawID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponse(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var v1PublicKeyCredentialWithAttestationTypeAuthenticatorAttachmentPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cross-platform","platform"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1PublicKeyCredentialWithAttestationTypeAuthenticatorAttachmentPropEnum = append(v1PublicKeyCredentialWithAttestationTypeAuthenticatorAttachmentPropEnum, v)
	}
}

const (

	// V1PublicKeyCredentialWithAttestationAuthenticatorAttachmentCrossDashPlatform captures enum value "cross-platform"
	V1PublicKeyCredentialWithAttestationAuthenticatorAttachmentCrossDashPlatform string = "cross-platform"

	// V1PublicKeyCredentialWithAttestationAuthenticatorAttachmentPlatform captures enum value "platform"
	V1PublicKeyCredentialWithAttestationAuthenticatorAttachmentPlatform string = "platform"
)

// prop value enum
func (m *V1PublicKeyCredentialWithAttestation) validateAuthenticatorAttachmentEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1PublicKeyCredentialWithAttestationTypeAuthenticatorAttachmentPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1PublicKeyCredentialWithAttestation) validateAuthenticatorAttachment(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthenticatorAttachment) { // not required
		return nil
	}

	// value enum
	if err := m.validateAuthenticatorAttachmentEnum("authenticatorAttachment", "body", *m.AuthenticatorAttachment); err != nil {
		return err
	}

	return nil
}

func (m *V1PublicKeyCredentialWithAttestation) validateClientExtensionResults(formats strfmt.Registry) error {

	if err := validate.Required("clientExtensionResults", "body", m.ClientExtensionResults); err != nil {
		return err
	}

	if m.ClientExtensionResults != nil {
		if err := m.ClientExtensionResults.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clientExtensionResults")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clientExtensionResults")
			}
			return err
		}
	}

	return nil
}

func (m *V1PublicKeyCredentialWithAttestation) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *V1PublicKeyCredentialWithAttestation) validateRawID(formats strfmt.Registry) error {

	if err := validate.Required("rawId", "body", m.RawID); err != nil {
		return err
	}

	return nil
}

func (m *V1PublicKeyCredentialWithAttestation) validateResponse(formats strfmt.Registry) error {

	if err := validate.Required("response", "body", m.Response); err != nil {
		return err
	}

	if m.Response != nil {
		if err := m.Response.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("response")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("response")
			}
			return err
		}
	}

	return nil
}

var v1PublicKeyCredentialWithAttestationTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["public-key"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1PublicKeyCredentialWithAttestationTypeTypePropEnum = append(v1PublicKeyCredentialWithAttestationTypeTypePropEnum, v)
	}
}

const (

	// V1PublicKeyCredentialWithAttestationTypePublicDashKey captures enum value "public-key"
	V1PublicKeyCredentialWithAttestationTypePublicDashKey string = "public-key"
)

// prop value enum
func (m *V1PublicKeyCredentialWithAttestation) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1PublicKeyCredentialWithAttestationTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1PublicKeyCredentialWithAttestation) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1 public key credential with attestation based on the context it is used
func (m *V1PublicKeyCredentialWithAttestation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClientExtensionResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResponse(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PublicKeyCredentialWithAttestation) contextValidateClientExtensionResults(ctx context.Context, formats strfmt.Registry) error {

	if m.ClientExtensionResults != nil {
		if err := m.ClientExtensionResults.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clientExtensionResults")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clientExtensionResults")
			}
			return err
		}
	}

	return nil
}

func (m *V1PublicKeyCredentialWithAttestation) contextValidateResponse(ctx context.Context, formats strfmt.Registry) error {

	if m.Response != nil {
		if err := m.Response.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("response")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("response")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1PublicKeyCredentialWithAttestation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PublicKeyCredentialWithAttestation) UnmarshalBinary(b []byte) error {
	var res V1PublicKeyCredentialWithAttestation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
