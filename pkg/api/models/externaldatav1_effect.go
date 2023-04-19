// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// Externaldatav1Effect externaldatav1 effect
//
// swagger:model externaldatav1Effect
type Externaldatav1Effect string

func NewExternaldatav1Effect(value Externaldatav1Effect) *Externaldatav1Effect {
	return &value
}

// Pointer returns a pointer to a freshly-allocated Externaldatav1Effect.
func (m Externaldatav1Effect) Pointer() *Externaldatav1Effect {
	return &m
}

const (

	// Externaldatav1EffectEFFECTALLOW captures enum value "EFFECT_ALLOW"
	Externaldatav1EffectEFFECTALLOW Externaldatav1Effect = "EFFECT_ALLOW"

	// Externaldatav1EffectEFFECTDENY captures enum value "EFFECT_DENY"
	Externaldatav1EffectEFFECTDENY Externaldatav1Effect = "EFFECT_DENY"
)

// for schema
var externaldatav1EffectEnum []interface{}

func init() {
	var res []Externaldatav1Effect
	if err := json.Unmarshal([]byte(`["EFFECT_ALLOW","EFFECT_DENY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		externaldatav1EffectEnum = append(externaldatav1EffectEnum, v)
	}
}

func (m Externaldatav1Effect) validateExternaldatav1EffectEnum(path, location string, value Externaldatav1Effect) error {
	if err := validate.EnumCase(path, location, value, externaldatav1EffectEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this externaldatav1 effect
func (m Externaldatav1Effect) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateExternaldatav1EffectEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this externaldatav1 effect based on context it is used
func (m Externaldatav1Effect) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
