// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MemoryPersistentMemoryUnit Memory:Persistent Memory Unit
//
// Persistent Memory Module on a server.
//
// swagger:model memoryPersistentMemoryUnit
type MemoryPersistentMemoryUnit struct {
	MemoryAbstractUnit

	// AppDirect capacity in GiB of the Persistent Memory Module on a server.
	// Read Only: true
	AppDirectCapacity string `json:"AppDirectCapacity,omitempty"`

	// Count status of the Persistent Memory Module on a server.
	// Read Only: true
	CountStatus string `json:"CountStatus,omitempty"`

	// Firmware version of the firware running on the Persistent Memory Module on a server.
	// Read Only: true
	FirmwareVersion string `json:"FirmwareVersion,omitempty"`

	// Frozen status of the Persistent Memory Module on a server.
	// Read Only: true
	FrozenStatus string `json:"FrozenStatus,omitempty"`

	// Health state of the Persistent Memory Module on a server.
	// Read Only: true
	HealthState string `json:"HealthState,omitempty"`

	// Lock status of the Persistent Memory Module on a server.
	// Read Only: true
	LockStatus string `json:"LockStatus,omitempty"`

	// A collection of references to the [memory.Array](mo://memory.Array) Managed Object.
	// When this managed object is deleted, the referenced [memory.Array](mo://memory.Array) MO unsets its reference to this deleted MO.
	// Read Only: true
	MemoryArray *MemoryArrayRef `json:"MemoryArray,omitempty"`

	// Memory capacity in GiB of the Persistent Memory Module on a server.
	// Read Only: true
	MemoryCapacity string `json:"MemoryCapacity,omitempty"`

	// ID of the Persistent Memory Module on a server.
	// Read Only: true
	MemoryID int64 `json:"MemoryId,omitempty"`

	// Persistent Memory capacity in GiB of the Persistent Memory Module on a server.
	// Read Only: true
	PersistentMemoryCapacity string `json:"PersistentMemoryCapacity,omitempty"`

	// The Device to which this Managed Object is associated.
	// Read Only: true
	RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

	// Reserved capacity in GiB of the Persistent Memory Module on a server.
	// Read Only: true
	ReservedCapacity string `json:"ReservedCapacity,omitempty"`

	// Security status of the Persistent Memory Module on a server.
	// Read Only: true
	SecurityStatus string `json:"SecurityStatus,omitempty"`

	// Socket ID of the Persistent Memory Module on a server.
	// Read Only: true
	SocketID string `json:"SocketId,omitempty"`

	// Socket Memory ID of the Persistent Memory Module on a server.
	// Read Only: true
	SocketMemoryID string `json:"SocketMemoryId,omitempty"`

	// Total capacity in GiB of the Persistent Memory Module on a server.
	// Read Only: true
	TotalCapacity string `json:"TotalCapacity,omitempty"`

	// UID of the Persistent Memory Module on a server.
	// Read Only: true
	UID string `json:"Uid,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *MemoryPersistentMemoryUnit) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MemoryAbstractUnit
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MemoryAbstractUnit = aO0

	// AO1
	var dataAO1 struct {
		AppDirectCapacity string `json:"AppDirectCapacity,omitempty"`

		CountStatus string `json:"CountStatus,omitempty"`

		FirmwareVersion string `json:"FirmwareVersion,omitempty"`

		FrozenStatus string `json:"FrozenStatus,omitempty"`

		HealthState string `json:"HealthState,omitempty"`

		LockStatus string `json:"LockStatus,omitempty"`

		MemoryArray *MemoryArrayRef `json:"MemoryArray,omitempty"`

		MemoryCapacity string `json:"MemoryCapacity,omitempty"`

		MemoryID int64 `json:"MemoryId,omitempty"`

		PersistentMemoryCapacity string `json:"PersistentMemoryCapacity,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

		ReservedCapacity string `json:"ReservedCapacity,omitempty"`

		SecurityStatus string `json:"SecurityStatus,omitempty"`

		SocketID string `json:"SocketId,omitempty"`

		SocketMemoryID string `json:"SocketMemoryId,omitempty"`

		TotalCapacity string `json:"TotalCapacity,omitempty"`

		UID string `json:"Uid,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.AppDirectCapacity = dataAO1.AppDirectCapacity

	m.CountStatus = dataAO1.CountStatus

	m.FirmwareVersion = dataAO1.FirmwareVersion

	m.FrozenStatus = dataAO1.FrozenStatus

	m.HealthState = dataAO1.HealthState

	m.LockStatus = dataAO1.LockStatus

	m.MemoryArray = dataAO1.MemoryArray

	m.MemoryCapacity = dataAO1.MemoryCapacity

	m.MemoryID = dataAO1.MemoryID

	m.PersistentMemoryCapacity = dataAO1.PersistentMemoryCapacity

	m.RegisteredDevice = dataAO1.RegisteredDevice

	m.ReservedCapacity = dataAO1.ReservedCapacity

	m.SecurityStatus = dataAO1.SecurityStatus

	m.SocketID = dataAO1.SocketID

	m.SocketMemoryID = dataAO1.SocketMemoryID

	m.TotalCapacity = dataAO1.TotalCapacity

	m.UID = dataAO1.UID

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m MemoryPersistentMemoryUnit) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MemoryAbstractUnit)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		AppDirectCapacity string `json:"AppDirectCapacity,omitempty"`

		CountStatus string `json:"CountStatus,omitempty"`

		FirmwareVersion string `json:"FirmwareVersion,omitempty"`

		FrozenStatus string `json:"FrozenStatus,omitempty"`

		HealthState string `json:"HealthState,omitempty"`

		LockStatus string `json:"LockStatus,omitempty"`

		MemoryArray *MemoryArrayRef `json:"MemoryArray,omitempty"`

		MemoryCapacity string `json:"MemoryCapacity,omitempty"`

		MemoryID int64 `json:"MemoryId,omitempty"`

		PersistentMemoryCapacity string `json:"PersistentMemoryCapacity,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

		ReservedCapacity string `json:"ReservedCapacity,omitempty"`

		SecurityStatus string `json:"SecurityStatus,omitempty"`

		SocketID string `json:"SocketId,omitempty"`

		SocketMemoryID string `json:"SocketMemoryId,omitempty"`

		TotalCapacity string `json:"TotalCapacity,omitempty"`

		UID string `json:"Uid,omitempty"`
	}

	dataAO1.AppDirectCapacity = m.AppDirectCapacity

	dataAO1.CountStatus = m.CountStatus

	dataAO1.FirmwareVersion = m.FirmwareVersion

	dataAO1.FrozenStatus = m.FrozenStatus

	dataAO1.HealthState = m.HealthState

	dataAO1.LockStatus = m.LockStatus

	dataAO1.MemoryArray = m.MemoryArray

	dataAO1.MemoryCapacity = m.MemoryCapacity

	dataAO1.MemoryID = m.MemoryID

	dataAO1.PersistentMemoryCapacity = m.PersistentMemoryCapacity

	dataAO1.RegisteredDevice = m.RegisteredDevice

	dataAO1.ReservedCapacity = m.ReservedCapacity

	dataAO1.SecurityStatus = m.SecurityStatus

	dataAO1.SocketID = m.SocketID

	dataAO1.SocketMemoryID = m.SocketMemoryID

	dataAO1.TotalCapacity = m.TotalCapacity

	dataAO1.UID = m.UID

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this memory persistent memory unit
func (m *MemoryPersistentMemoryUnit) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MemoryAbstractUnit
	if err := m.MemoryAbstractUnit.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemoryArray(formats); err != nil {
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

func (m *MemoryPersistentMemoryUnit) validateMemoryArray(formats strfmt.Registry) error {

	if swag.IsZero(m.MemoryArray) { // not required
		return nil
	}

	if m.MemoryArray != nil {
		if err := m.MemoryArray.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MemoryArray")
			}
			return err
		}
	}

	return nil
}

func (m *MemoryPersistentMemoryUnit) validateRegisteredDevice(formats strfmt.Registry) error {

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
func (m *MemoryPersistentMemoryUnit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MemoryPersistentMemoryUnit) UnmarshalBinary(b []byte) error {
	var res MemoryPersistentMemoryUnit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
