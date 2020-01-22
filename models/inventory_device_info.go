// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InventoryDeviceInfo Inventory:Device Info
//
// Information pertaining to a Registered Device in starship which is pertinent to Inventory Microservice.
//
// swagger:model inventoryDeviceInfo
type InventoryDeviceInfo struct {
	PolicyinventoryAbstractDeviceInfo

	InventoryDeviceInfoAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *InventoryDeviceInfo) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 PolicyinventoryAbstractDeviceInfo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.PolicyinventoryAbstractDeviceInfo = aO0

	// AO1
	var aO1 InventoryDeviceInfoAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.InventoryDeviceInfoAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m InventoryDeviceInfo) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.PolicyinventoryAbstractDeviceInfo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.InventoryDeviceInfoAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this inventory device info
func (m *InventoryDeviceInfo) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with PolicyinventoryAbstractDeviceInfo
	if err := m.PolicyinventoryAbstractDeviceInfo.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with InventoryDeviceInfoAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *InventoryDeviceInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InventoryDeviceInfo) UnmarshalBinary(b []byte) error {
	var res InventoryDeviceInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// InventoryDeviceInfoAllOf1 inventory device info all of1
// swagger:model InventoryDeviceInfoAllOf1
type InventoryDeviceInfoAllOf1 interface{}
