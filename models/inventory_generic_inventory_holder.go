// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InventoryGenericInventoryHolder Inventory:Generic Inventory Holder
//
// A container class for generic inventory.
//
// swagger:model inventoryGenericInventoryHolder
type InventoryGenericInventoryHolder struct {
	InventoryBase

	// A collection of references to the [compute.Blade](mo://compute.Blade) Managed Object.
	// When this managed object is deleted, the referenced [compute.Blade](mo://compute.Blade) MO unsets its reference to this deleted MO.
	// Read Only: true
	ComputeBlade *ComputeBladeRef `json:"ComputeBlade,omitempty"`

	// A collection of references to the [compute.RackUnit](mo://compute.RackUnit) Managed Object.
	// When this managed object is deleted, the referenced [compute.RackUnit](mo://compute.RackUnit) MO unsets its reference to this deleted MO.
	// Read Only: true
	ComputeRackUnit *ComputeRackUnitRef `json:"ComputeRackUnit,omitempty"`

	// endpoint
	// Read Only: true
	Endpoint string `json:"Endpoint,omitempty"`

	// generic inventory
	// Read Only: true
	GenericInventory []*InventoryGenericInventoryRef `json:"GenericInventory"`

	// The Device to which this Managed Object is associated.
	// Read Only: true
	RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *InventoryGenericInventoryHolder) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 InventoryBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.InventoryBase = aO0

	// AO1
	var dataAO1 struct {
		ComputeBlade *ComputeBladeRef `json:"ComputeBlade,omitempty"`

		ComputeRackUnit *ComputeRackUnitRef `json:"ComputeRackUnit,omitempty"`

		Endpoint string `json:"Endpoint,omitempty"`

		GenericInventory []*InventoryGenericInventoryRef `json:"GenericInventory"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ComputeBlade = dataAO1.ComputeBlade

	m.ComputeRackUnit = dataAO1.ComputeRackUnit

	m.Endpoint = dataAO1.Endpoint

	m.GenericInventory = dataAO1.GenericInventory

	m.RegisteredDevice = dataAO1.RegisteredDevice

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m InventoryGenericInventoryHolder) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.InventoryBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		ComputeBlade *ComputeBladeRef `json:"ComputeBlade,omitempty"`

		ComputeRackUnit *ComputeRackUnitRef `json:"ComputeRackUnit,omitempty"`

		Endpoint string `json:"Endpoint,omitempty"`

		GenericInventory []*InventoryGenericInventoryRef `json:"GenericInventory"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`
	}

	dataAO1.ComputeBlade = m.ComputeBlade

	dataAO1.ComputeRackUnit = m.ComputeRackUnit

	dataAO1.Endpoint = m.Endpoint

	dataAO1.GenericInventory = m.GenericInventory

	dataAO1.RegisteredDevice = m.RegisteredDevice

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this inventory generic inventory holder
func (m *InventoryGenericInventoryHolder) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with InventoryBase
	if err := m.InventoryBase.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComputeBlade(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComputeRackUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGenericInventory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegisteredDevice(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InventoryGenericInventoryHolder) validateComputeBlade(formats strfmt.Registry) error {

	if swag.IsZero(m.ComputeBlade) { // not required
		return nil
	}

	if m.ComputeBlade != nil {
		if err := m.ComputeBlade.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ComputeBlade")
			}
			return err
		}
	}

	return nil
}

func (m *InventoryGenericInventoryHolder) validateComputeRackUnit(formats strfmt.Registry) error {

	if swag.IsZero(m.ComputeRackUnit) { // not required
		return nil
	}

	if m.ComputeRackUnit != nil {
		if err := m.ComputeRackUnit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ComputeRackUnit")
			}
			return err
		}
	}

	return nil
}

func (m *InventoryGenericInventoryHolder) validateGenericInventory(formats strfmt.Registry) error {

	if swag.IsZero(m.GenericInventory) { // not required
		return nil
	}

	for i := 0; i < len(m.GenericInventory); i++ {
		if swag.IsZero(m.GenericInventory[i]) { // not required
			continue
		}

		if m.GenericInventory[i] != nil {
			if err := m.GenericInventory[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("GenericInventory" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InventoryGenericInventoryHolder) validateRegisteredDevice(formats strfmt.Registry) error {

	if swag.IsZero(m.RegisteredDevice) { // not required
		return nil
	}

	if m.RegisteredDevice != nil {
		if err := m.RegisteredDevice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RegisteredDevice")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InventoryGenericInventoryHolder) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InventoryGenericInventoryHolder) UnmarshalBinary(b []byte) error {
	var res InventoryGenericInventoryHolder
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
