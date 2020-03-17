// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StorageVdMemberEp Storage:Vd Member Ep
//
// Reference to LocalDisk to build up a VirtualDrive.
//
// swagger:model storageVdMemberEp
type StorageVdMemberEp struct {
	InventoryBase

	// oper qualifier reason
	// Read Only: true
	OperQualifierReason string `json:"OperQualifierReason,omitempty"`

	// presence
	// Read Only: true
	Presence string `json:"Presence,omitempty"`

	// The Device to which this Managed Object is associated.
	// Read Only: true
	RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

	// Role of the disk normal or hot-spare, used by virtual-drive.
	// Read Only: true
	Role string `json:"Role,omitempty"`

	// span Id
	// Read Only: true
	SpanID string `json:"SpanId,omitempty"`

	// A collection of references to the [storage.VirtualDrive](mo://storage.VirtualDrive) Managed Object.
	// When this managed object is deleted, the referenced [storage.VirtualDrive](mo://storage.VirtualDrive) MO unsets its reference to this deleted MO.
	// Read Only: true
	StorageVirtualDrive *StorageVirtualDriveRef `json:"StorageVirtualDrive,omitempty"`

	// It shows local disk slot number as id.
	// Read Only: true
	VdMemberEpID int64 `json:"VdMemberEpId,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *StorageVdMemberEp) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 InventoryBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.InventoryBase = aO0

	// AO1
	var dataAO1 struct {
		OperQualifierReason string `json:"OperQualifierReason,omitempty"`

		Presence string `json:"Presence,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

		Role string `json:"Role,omitempty"`

		SpanID string `json:"SpanId,omitempty"`

		StorageVirtualDrive *StorageVirtualDriveRef `json:"StorageVirtualDrive,omitempty"`

		VdMemberEpID int64 `json:"VdMemberEpId,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.OperQualifierReason = dataAO1.OperQualifierReason

	m.Presence = dataAO1.Presence

	m.RegisteredDevice = dataAO1.RegisteredDevice

	m.Role = dataAO1.Role

	m.SpanID = dataAO1.SpanID

	m.StorageVirtualDrive = dataAO1.StorageVirtualDrive

	m.VdMemberEpID = dataAO1.VdMemberEpID

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m StorageVdMemberEp) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.InventoryBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		OperQualifierReason string `json:"OperQualifierReason,omitempty"`

		Presence string `json:"Presence,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

		Role string `json:"Role,omitempty"`

		SpanID string `json:"SpanId,omitempty"`

		StorageVirtualDrive *StorageVirtualDriveRef `json:"StorageVirtualDrive,omitempty"`

		VdMemberEpID int64 `json:"VdMemberEpId,omitempty"`
	}

	dataAO1.OperQualifierReason = m.OperQualifierReason

	dataAO1.Presence = m.Presence

	dataAO1.RegisteredDevice = m.RegisteredDevice

	dataAO1.Role = m.Role

	dataAO1.SpanID = m.SpanID

	dataAO1.StorageVirtualDrive = m.StorageVirtualDrive

	dataAO1.VdMemberEpID = m.VdMemberEpID

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this storage vd member ep
func (m *StorageVdMemberEp) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with InventoryBase
	if err := m.InventoryBase.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegisteredDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageVirtualDrive(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StorageVdMemberEp) validateRegisteredDevice(formats strfmt.Registry) error {

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

func (m *StorageVdMemberEp) validateStorageVirtualDrive(formats strfmt.Registry) error {

	if swag.IsZero(m.StorageVirtualDrive) { // not required
		return nil
	}

	if m.StorageVirtualDrive != nil {
		if err := m.StorageVirtualDrive.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("StorageVirtualDrive")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StorageVdMemberEp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageVdMemberEp) UnmarshalBinary(b []byte) error {
	var res StorageVdMemberEp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
