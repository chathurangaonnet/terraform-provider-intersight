// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BootLocalCdd Local CDD
//
// Device type used when booting from local CD drive.
//
// swagger:model bootLocalCdd
type BootLocalCdd struct {
	BootDeviceBase

	BootLocalCddAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *BootLocalCdd) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BootDeviceBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BootDeviceBase = aO0

	// AO1
	var aO1 BootLocalCddAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.BootLocalCddAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m BootLocalCdd) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.BootDeviceBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.BootLocalCddAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this boot local cdd
func (m *BootLocalCdd) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BootDeviceBase
	if err := m.BootDeviceBase.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with BootLocalCddAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *BootLocalCdd) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BootLocalCdd) UnmarshalBinary(b []byte) error {
	var res BootLocalCdd
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BootLocalCddAllOf1 boot local cdd all of1
//
// swagger:model BootLocalCddAllOf1
type BootLocalCddAllOf1 interface{}
