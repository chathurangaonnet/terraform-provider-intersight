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

// VnicFcInterruptSettings Interrupt Settings
//
// Interrupt Settings for the virtual fibre channel interface.
//
// swagger:model vnicFcInterruptSettings
type VnicFcInterruptSettings struct {
	MoBaseComplexType

	VnicFcInterruptSettingsAO1P1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *VnicFcInterruptSettings) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseComplexType
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseComplexType = aO0

	// AO1
	var aO1 VnicFcInterruptSettingsAO1P1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.VnicFcInterruptSettingsAO1P1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m VnicFcInterruptSettings) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseComplexType)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.VnicFcInterruptSettingsAO1P1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this vnic fc interrupt settings
func (m *VnicFcInterruptSettings) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseComplexType
	if err := m.MoBaseComplexType.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with VnicFcInterruptSettingsAO1P1
	if err := m.VnicFcInterruptSettingsAO1P1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *VnicFcInterruptSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VnicFcInterruptSettings) UnmarshalBinary(b []byte) error {
	var res VnicFcInterruptSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VnicFcInterruptSettingsAO1P1 vnic fc interrupt settings a o1 p1
// swagger:model VnicFcInterruptSettingsAO1P1
type VnicFcInterruptSettingsAO1P1 struct {

	// The preferred driver interrupt mode. This can be one of the following:- MSIx - Message Signaled Interrupts (MSI) with the optional extension. MSI  - MSI only. INTx - PCI INTx interrupts. MSIx is the recommended option.
	//
	// Enum: [MSIx MSI INTx]
	Mode *string `json:"Mode,omitempty"`

	// vnic fc interrupt settings a o1 p1
	VnicFcInterruptSettingsAO1P1 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *VnicFcInterruptSettingsAO1P1) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// The preferred driver interrupt mode. This can be one of the following:- MSIx - Message Signaled Interrupts (MSI) with the optional extension. MSI  - MSI only. INTx - PCI INTx interrupts. MSIx is the recommended option.
		//
		// Enum: [MSIx MSI INTx]
		Mode *string `json:"Mode,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv VnicFcInterruptSettingsAO1P1

	rcv.Mode = stage1.Mode

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "Mode")

	// stage 3, add additional properties values
	if len(stage2) > 0 {
		result := make(map[string]interface{})
		for k, v := range stage2 {
			var toadd interface{}
			if err := json.Unmarshal(v, &toadd); err != nil {
				return err
			}
			result[k] = toadd
		}
		m.VnicFcInterruptSettingsAO1P1 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m VnicFcInterruptSettingsAO1P1) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// The preferred driver interrupt mode. This can be one of the following:- MSIx - Message Signaled Interrupts (MSI) with the optional extension. MSI  - MSI only. INTx - PCI INTx interrupts. MSIx is the recommended option.
		//
		// Enum: [MSIx MSI INTx]
		Mode *string `json:"Mode,omitempty"`
	}

	stage1.Mode = m.Mode

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.VnicFcInterruptSettingsAO1P1) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.VnicFcInterruptSettingsAO1P1)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 {
		return additional, nil
	}

	// concatenate the 2 objects
	props[len(props)-1] = ','
	return append(props, additional[1:]...), nil
}

// Validate validates this vnic fc interrupt settings a o1 p1
func (m *VnicFcInterruptSettingsAO1P1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var vnicFcInterruptSettingsAO1P1TypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["MSIx","MSI","INTx"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vnicFcInterruptSettingsAO1P1TypeModePropEnum = append(vnicFcInterruptSettingsAO1P1TypeModePropEnum, v)
	}
}

const (

	// VnicFcInterruptSettingsAO1P1ModeMSIx captures enum value "MSIx"
	VnicFcInterruptSettingsAO1P1ModeMSIx string = "MSIx"

	// VnicFcInterruptSettingsAO1P1ModeMSI captures enum value "MSI"
	VnicFcInterruptSettingsAO1P1ModeMSI string = "MSI"

	// VnicFcInterruptSettingsAO1P1ModeINTx captures enum value "INTx"
	VnicFcInterruptSettingsAO1P1ModeINTx string = "INTx"
)

// prop value enum
func (m *VnicFcInterruptSettingsAO1P1) validateModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, vnicFcInterruptSettingsAO1P1TypeModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *VnicFcInterruptSettingsAO1P1) validateMode(formats strfmt.Registry) error {

	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	// value enum
	if err := m.validateModeEnum("Mode", "body", *m.Mode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VnicFcInterruptSettingsAO1P1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VnicFcInterruptSettingsAO1P1) UnmarshalBinary(b []byte) error {
	var res VnicFcInterruptSettingsAO1P1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
