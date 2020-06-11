// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NiatelemetryNiaLicenseState Niatelemetry:Nia License State
//
// Object available at device scope for license information. This determines the usage of this attribute.
//
// swagger:model niatelemetryNiaLicenseState
type NiatelemetryNiaLicenseState struct {
	MoBaseMo

	// A collection of references to the [niatelemetry.NiaInventory](mo://niatelemetry.NiaInventory) Managed Object.
	// When this managed object is deleted, the referenced [niatelemetry.NiaInventory](mo://niatelemetry.NiaInventory) MO unsets its reference to this deleted MO.
	Device *NiatelemetryNiaInventoryRef `json:"Device,omitempty"`

	// Features activated on device being inventoried. This determines which features are currently enabled on the device that the license API can check.
	FeatureActivated string `json:"FeatureActivated,omitempty"`

	// Licenses activated on device being inventoried. This determines which lienceses are currently enabled on the device.
	LicenseActivated string `json:"LicenseActivated,omitempty"`

	// PID of device being inventoried. This determines the hardware model type of the device.
	PidType string `json:"PidType,omitempty"`

	// Serial number of device being inventoried. The serial number is unique per device.
	Serial string `json:"Serial,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *NiatelemetryNiaLicenseState) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		Device *NiatelemetryNiaInventoryRef `json:"Device,omitempty"`

		FeatureActivated string `json:"FeatureActivated,omitempty"`

		LicenseActivated string `json:"LicenseActivated,omitempty"`

		PidType string `json:"PidType,omitempty"`

		Serial string `json:"Serial,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Device = dataAO1.Device

	m.FeatureActivated = dataAO1.FeatureActivated

	m.LicenseActivated = dataAO1.LicenseActivated

	m.PidType = dataAO1.PidType

	m.Serial = dataAO1.Serial

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m NiatelemetryNiaLicenseState) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Device *NiatelemetryNiaInventoryRef `json:"Device,omitempty"`

		FeatureActivated string `json:"FeatureActivated,omitempty"`

		LicenseActivated string `json:"LicenseActivated,omitempty"`

		PidType string `json:"PidType,omitempty"`

		Serial string `json:"Serial,omitempty"`
	}

	dataAO1.Device = m.Device

	dataAO1.FeatureActivated = m.FeatureActivated

	dataAO1.LicenseActivated = m.LicenseActivated

	dataAO1.PidType = m.PidType

	dataAO1.Serial = m.Serial

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this niatelemetry nia license state
func (m *NiatelemetryNiaLicenseState) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NiatelemetryNiaLicenseState) validateDevice(formats strfmt.Registry) error {

	if swag.IsZero(m.Device) { // not required
		return nil
	}

	if m.Device != nil {
		if err := m.Device.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Device")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NiatelemetryNiaLicenseState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NiatelemetryNiaLicenseState) UnmarshalBinary(b []byte) error {
	var res NiatelemetryNiaLicenseState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
