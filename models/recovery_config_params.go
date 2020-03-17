// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RecoveryConfigParams Recovery:Config Params
//
// Encapsulates the restore workflows configuration parameters.
//
// swagger:model recoveryConfigParams
type RecoveryConfigParams struct {
	MoBaseComplexType

	// a o1
	AO1 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *RecoveryConfigParams) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseComplexType
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseComplexType = aO0

	// AO1
	var aO1 map[string]interface{}
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.AO1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m RecoveryConfigParams) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseComplexType)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	aO1, err := swag.WriteJSON(m.AO1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this recovery config params
func (m *RecoveryConfigParams) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseComplexType
	if err := m.MoBaseComplexType.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with map[string]interface{}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *RecoveryConfigParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoveryConfigParams) UnmarshalBinary(b []byte) error {
	var res RecoveryConfigParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
