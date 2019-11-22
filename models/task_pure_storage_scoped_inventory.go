// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TaskPureStorageScopedInventory Task:Pure Storage Scoped Inventory
//
// API to trigger on-demand PureStorage FlashArray inventory to update modified objects in Intersight report.
//
// swagger:model taskPureStorageScopedInventory
type TaskPureStorageScopedInventory struct {
	MoBaseMo

	// PureStorage FlashArray managed object name. Example: storage.PureHost, storage.PureVolume etc.
	//
	//
	ManagedObject string `json:"ManagedObject,omitempty"`

	// Attribute name of PureStorage FlashArray managed object.
	// Example: For storage.PureVolume managed object, Name is an attribute.
	//
	//
	PropertyName string `json:"PropertyName,omitempty"`

	// Managed object property value. It is an object name for which inventory need to be collected or updated.
	//
	//
	PropertyValue []string `json:"PropertyValue"`

	// Device registration managed object that represents this storage array connection to Intersight.
	//
	RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TaskPureStorageScopedInventory) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		ManagedObject string `json:"ManagedObject,omitempty"`

		PropertyName string `json:"PropertyName,omitempty"`

		PropertyValue []string `json:"PropertyValue"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ManagedObject = dataAO1.ManagedObject

	m.PropertyName = dataAO1.PropertyName

	m.PropertyValue = dataAO1.PropertyValue

	m.RegisteredDevice = dataAO1.RegisteredDevice

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TaskPureStorageScopedInventory) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		ManagedObject string `json:"ManagedObject,omitempty"`

		PropertyName string `json:"PropertyName,omitempty"`

		PropertyValue []string `json:"PropertyValue"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`
	}

	dataAO1.ManagedObject = m.ManagedObject

	dataAO1.PropertyName = m.PropertyName

	dataAO1.PropertyValue = m.PropertyValue

	dataAO1.RegisteredDevice = m.RegisteredDevice

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this task pure storage scoped inventory
func (m *TaskPureStorageScopedInventory) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
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

func (m *TaskPureStorageScopedInventory) validateRegisteredDevice(formats strfmt.Registry) error {

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
func (m *TaskPureStorageScopedInventory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskPureStorageScopedInventory) UnmarshalBinary(b []byte) error {
	var res TaskPureStorageScopedInventory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}