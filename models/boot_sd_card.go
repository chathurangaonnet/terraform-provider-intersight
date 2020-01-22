// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BootSdCard SD Card
//
// Device type used when booting from SD Card device.
//
// swagger:model bootSdCard
type BootSdCard struct {
	BootDeviceBase

	// The Logical Unit Number (LUN) of the device.
	//
	Lun int64 `json:"Lun,omitempty"`

	// The subtype for the selected device type.
	//
	// Enum: [None flex-util flex-flash SDCARD]
	Subtype *string `json:"Subtype,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *BootSdCard) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BootDeviceBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BootDeviceBase = aO0

	// AO1
	var dataAO1 struct {
		Lun int64 `json:"Lun,omitempty"`

		Subtype *string `json:"Subtype,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Lun = dataAO1.Lun

	m.Subtype = dataAO1.Subtype

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m BootSdCard) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.BootDeviceBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Lun int64 `json:"Lun,omitempty"`

		Subtype *string `json:"Subtype,omitempty"`
	}

	dataAO1.Lun = m.Lun

	dataAO1.Subtype = m.Subtype

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this boot sd card
func (m *BootSdCard) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BootDeviceBase
	if err := m.BootDeviceBase.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubtype(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var bootSdCardTypeSubtypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["None","flex-util","flex-flash","SDCARD"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		bootSdCardTypeSubtypePropEnum = append(bootSdCardTypeSubtypePropEnum, v)
	}
}

// property enum
func (m *BootSdCard) validateSubtypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, bootSdCardTypeSubtypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *BootSdCard) validateSubtype(formats strfmt.Registry) error {

	if swag.IsZero(m.Subtype) { // not required
		return nil
	}

	// value enum
	if err := m.validateSubtypeEnum("Subtype", "body", *m.Subtype); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BootSdCard) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BootSdCard) UnmarshalBinary(b []byte) error {
	var res BootSdCard
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}