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

// V1Policy v1 policy
//
// swagger:model v1Policy
type V1Policy struct {

	// A condition expression that evalutes to true or false.
	// Required: true
	Condition *string `json:"condition"`

	// A consensus expression that evalutes to true or false.
	// Required: true
	Consensus *string `json:"consensus"`

	// created at
	// Required: true
	CreatedAt *V1Timestamp `json:"createdAt"`

	// The instruction to DENY or ALLOW a particular activity following policy selector(s).
	// Required: true
	Effect *Externaldatav1Effect `json:"effect"`

	// Human-readable notes added by a User to describe a particular policy.
	// Required: true
	Notes *string `json:"notes"`

	// Unique identifier for a given Policy.
	// Required: true
	PolicyID *string `json:"policyId"`

	// Human-readable name for a Policy.
	// Required: true
	PolicyName *string `json:"policyName"`

	// A list of simple functions each including a subject, target and boolean. See Policy Engine Language section for additional details.
	// Required: true
	Selectors []*Externaldatav1Selector `json:"selectors"`

	// updated at
	// Required: true
	UpdatedAt *V1Timestamp `json:"updatedAt"`
}

// Validate validates this v1 policy
func (m *V1Policy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCondition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConsensus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEffect(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicyID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicyName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelectors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Policy) validateCondition(formats strfmt.Registry) error {

	if err := validate.Required("condition", "body", m.Condition); err != nil {
		return err
	}

	return nil
}

func (m *V1Policy) validateConsensus(formats strfmt.Registry) error {

	if err := validate.Required("consensus", "body", m.Consensus); err != nil {
		return err
	}

	return nil
}

func (m *V1Policy) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("createdAt", "body", m.CreatedAt); err != nil {
		return err
	}

	if m.CreatedAt != nil {
		if err := m.CreatedAt.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createdAt")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createdAt")
			}
			return err
		}
	}

	return nil
}

func (m *V1Policy) validateEffect(formats strfmt.Registry) error {

	if err := validate.Required("effect", "body", m.Effect); err != nil {
		return err
	}

	if err := validate.Required("effect", "body", m.Effect); err != nil {
		return err
	}

	if m.Effect != nil {
		if err := m.Effect.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("effect")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("effect")
			}
			return err
		}
	}

	return nil
}

func (m *V1Policy) validateNotes(formats strfmt.Registry) error {

	if err := validate.Required("notes", "body", m.Notes); err != nil {
		return err
	}

	return nil
}

func (m *V1Policy) validatePolicyID(formats strfmt.Registry) error {

	if err := validate.Required("policyId", "body", m.PolicyID); err != nil {
		return err
	}

	return nil
}

func (m *V1Policy) validatePolicyName(formats strfmt.Registry) error {

	if err := validate.Required("policyName", "body", m.PolicyName); err != nil {
		return err
	}

	return nil
}

func (m *V1Policy) validateSelectors(formats strfmt.Registry) error {

	if err := validate.Required("selectors", "body", m.Selectors); err != nil {
		return err
	}

	for i := 0; i < len(m.Selectors); i++ {
		if swag.IsZero(m.Selectors[i]) { // not required
			continue
		}

		if m.Selectors[i] != nil {
			if err := m.Selectors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("selectors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("selectors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1Policy) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updatedAt", "body", m.UpdatedAt); err != nil {
		return err
	}

	if m.UpdatedAt != nil {
		if err := m.UpdatedAt.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updatedAt")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updatedAt")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 policy based on the context it is used
func (m *V1Policy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEffect(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelectors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Policy) contextValidateCreatedAt(ctx context.Context, formats strfmt.Registry) error {

	if m.CreatedAt != nil {
		if err := m.CreatedAt.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createdAt")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createdAt")
			}
			return err
		}
	}

	return nil
}

func (m *V1Policy) contextValidateEffect(ctx context.Context, formats strfmt.Registry) error {

	if m.Effect != nil {
		if err := m.Effect.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("effect")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("effect")
			}
			return err
		}
	}

	return nil
}

func (m *V1Policy) contextValidateSelectors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Selectors); i++ {

		if m.Selectors[i] != nil {
			if err := m.Selectors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("selectors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("selectors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1Policy) contextValidateUpdatedAt(ctx context.Context, formats strfmt.Registry) error {

	if m.UpdatedAt != nil {
		if err := m.UpdatedAt.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updatedAt")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updatedAt")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Policy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Policy) UnmarshalBinary(b []byte) error {
	var res V1Policy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
