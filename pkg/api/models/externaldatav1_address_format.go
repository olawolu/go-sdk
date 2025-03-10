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

// Externaldatav1AddressFormat externaldatav1 address format
//
// swagger:model externaldatav1AddressFormat
type Externaldatav1AddressFormat string

func NewExternaldatav1AddressFormat(value Externaldatav1AddressFormat) *Externaldatav1AddressFormat {
	return &value
}

// Pointer returns a pointer to a freshly-allocated Externaldatav1AddressFormat.
func (m Externaldatav1AddressFormat) Pointer() *Externaldatav1AddressFormat {
	return &m
}

const (

	// Externaldatav1AddressFormatADDRESSFORMATUNCOMPRESSED captures enum value "ADDRESS_FORMAT_UNCOMPRESSED"
	Externaldatav1AddressFormatADDRESSFORMATUNCOMPRESSED Externaldatav1AddressFormat = "ADDRESS_FORMAT_UNCOMPRESSED"

	// Externaldatav1AddressFormatADDRESSFORMATCOMPRESSED captures enum value "ADDRESS_FORMAT_COMPRESSED"
	Externaldatav1AddressFormatADDRESSFORMATCOMPRESSED Externaldatav1AddressFormat = "ADDRESS_FORMAT_COMPRESSED"

	// Externaldatav1AddressFormatADDRESSFORMATETHEREUM captures enum value "ADDRESS_FORMAT_ETHEREUM"
	Externaldatav1AddressFormatADDRESSFORMATETHEREUM Externaldatav1AddressFormat = "ADDRESS_FORMAT_ETHEREUM"
)

// for schema
var externaldatav1AddressFormatEnum []interface{}

func init() {
	var res []Externaldatav1AddressFormat
	if err := json.Unmarshal([]byte(`["ADDRESS_FORMAT_UNCOMPRESSED","ADDRESS_FORMAT_COMPRESSED","ADDRESS_FORMAT_ETHEREUM"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		externaldatav1AddressFormatEnum = append(externaldatav1AddressFormatEnum, v)
	}
}

func (m Externaldatav1AddressFormat) validateExternaldatav1AddressFormatEnum(path, location string, value Externaldatav1AddressFormat) error {
	if err := validate.EnumCase(path, location, value, externaldatav1AddressFormatEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this externaldatav1 address format
func (m Externaldatav1AddressFormat) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateExternaldatav1AddressFormatEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this externaldatav1 address format based on context it is used
func (m Externaldatav1AddressFormat) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
