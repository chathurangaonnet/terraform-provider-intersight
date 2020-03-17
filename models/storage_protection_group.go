// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StorageProtectionGroup Storage:Protection Group
//
// A protection group contains list of members that are protected together through snapshots with point-in-time consistency across the members.
// Members within the protection group have common data protection requirements and also the same snapshot, replication, and retention schedules.
//
// swagger:model storageProtectionGroup
type StorageProtectionGroup struct {
	MoBaseMo

	// Name of the protection Group.
	Name string `json:"Name,omitempty"`

	// Prefix used for all generated snapshots from the protection group.
	Prefix string `json:"Prefix,omitempty"`

	// Flag to determine if replication is enabled. Snapshots are replicated to the target array if this flag is set.
	ReplicationEnabled *bool `json:"ReplicationEnabled,omitempty"`

	// Flag to determine if snapshot creation is enabled. Snapshots are created on local array if this flag is set.
	SnapshotEnabled *bool `json:"SnapshotEnabled,omitempty"`

	// Storage array managed object.
	// Read Only: true
	StorageArray *StorageGenericArrayRef `json:"StorageArray,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *StorageProtectionGroup) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		Name string `json:"Name,omitempty"`

		Prefix string `json:"Prefix,omitempty"`

		ReplicationEnabled *bool `json:"ReplicationEnabled,omitempty"`

		SnapshotEnabled *bool `json:"SnapshotEnabled,omitempty"`

		StorageArray *StorageGenericArrayRef `json:"StorageArray,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Name = dataAO1.Name

	m.Prefix = dataAO1.Prefix

	m.ReplicationEnabled = dataAO1.ReplicationEnabled

	m.SnapshotEnabled = dataAO1.SnapshotEnabled

	m.StorageArray = dataAO1.StorageArray

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m StorageProtectionGroup) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Name string `json:"Name,omitempty"`

		Prefix string `json:"Prefix,omitempty"`

		ReplicationEnabled *bool `json:"ReplicationEnabled,omitempty"`

		SnapshotEnabled *bool `json:"SnapshotEnabled,omitempty"`

		StorageArray *StorageGenericArrayRef `json:"StorageArray,omitempty"`
	}

	dataAO1.Name = m.Name

	dataAO1.Prefix = m.Prefix

	dataAO1.ReplicationEnabled = m.ReplicationEnabled

	dataAO1.SnapshotEnabled = m.SnapshotEnabled

	dataAO1.StorageArray = m.StorageArray

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this storage protection group
func (m *StorageProtectionGroup) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageArray(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StorageProtectionGroup) validateStorageArray(formats strfmt.Registry) error {

	if swag.IsZero(m.StorageArray) { // not required
		return nil
	}

	if m.StorageArray != nil {
		if err := m.StorageArray.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("StorageArray")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StorageProtectionGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageProtectionGroup) UnmarshalBinary(b []byte) error {
	var res StorageProtectionGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
