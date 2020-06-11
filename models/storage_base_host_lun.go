// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StorageBaseHostLun Storage:Base Host Lun
//
// Generic storage host lun object. It exists only if the volume is associated to host initiator.
//
// swagger:model storageBaseHostLun
type StorageBaseHostLun struct {
	MoBaseMo

	// Logical unit number (LUN) by which hosts address specified volume. Hlu is a decimal representation of the LUN from the endpoint.
	// Read Only: true
	Hlu int64 `json:"Hlu,omitempty"`

	// Name of the host associated with LUN.
	// Read Only: true
	HostName string `json:"HostName,omitempty"`

	// Name of the storage volume associated with LUN.
	// Read Only: true
	VolumeName string `json:"VolumeName,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *StorageBaseHostLun) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		Hlu int64 `json:"Hlu,omitempty"`

		HostName string `json:"HostName,omitempty"`

		VolumeName string `json:"VolumeName,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Hlu = dataAO1.Hlu

	m.HostName = dataAO1.HostName

	m.VolumeName = dataAO1.VolumeName

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m StorageBaseHostLun) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Hlu int64 `json:"Hlu,omitempty"`

		HostName string `json:"HostName,omitempty"`

		VolumeName string `json:"VolumeName,omitempty"`
	}

	dataAO1.Hlu = m.Hlu

	dataAO1.HostName = m.HostName

	dataAO1.VolumeName = m.VolumeName

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this storage base host lun
func (m *StorageBaseHostLun) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *StorageBaseHostLun) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageBaseHostLun) UnmarshalBinary(b []byte) error {
	var res StorageBaseHostLun
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
