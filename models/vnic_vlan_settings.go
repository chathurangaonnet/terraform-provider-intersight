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

// VnicVlanSettings VLAN Settings
//
// VLAN configuration for the virtual interface.
//
// swagger:model vnicVlanSettings
type VnicVlanSettings struct {
	MoBaseComplexType

	VnicVlanSettingsAO1P1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *VnicVlanSettings) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseComplexType
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseComplexType = aO0

	// AO1
	var aO1 VnicVlanSettingsAO1P1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.VnicVlanSettingsAO1P1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m VnicVlanSettings) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseComplexType)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.VnicVlanSettingsAO1P1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this vnic vlan settings
func (m *VnicVlanSettings) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseComplexType
	if err := m.MoBaseComplexType.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with VnicVlanSettingsAO1P1
	if err := m.VnicVlanSettingsAO1P1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *VnicVlanSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VnicVlanSettings) UnmarshalBinary(b []byte) error {
	var res VnicVlanSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VnicVlanSettingsAO1P1 vnic vlan settings a o1 p1
// swagger:model VnicVlanSettingsAO1P1
type VnicVlanSettingsAO1P1 struct {

	// Default VLAN ID of the virtual interface. Setting the ID to 0 will not associate any default VLAN to the traffic on the virtual interface.
	//
	DefaultVlan int64 `json:"DefaultVlan,omitempty"`

	// Option to determine if the port can carry single VLAN (Access) or multiple VLANs (Trunk) traffic.
	//
	// Enum: [ACCESS TRUNK]
	Mode *string `json:"Mode,omitempty"`

	// vnic vlan settings a o1 p1
	VnicVlanSettingsAO1P1 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *VnicVlanSettingsAO1P1) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// Default VLAN ID of the virtual interface. Setting the ID to 0 will not associate any default VLAN to the traffic on the virtual interface.
		//
		DefaultVlan int64 `json:"DefaultVlan,omitempty"`

		// Option to determine if the port can carry single VLAN (Access) or multiple VLANs (Trunk) traffic.
		//
		// Enum: [ACCESS TRUNK]
		Mode *string `json:"Mode,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv VnicVlanSettingsAO1P1

	rcv.DefaultVlan = stage1.DefaultVlan

	rcv.Mode = stage1.Mode

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "DefaultVlan")

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
		m.VnicVlanSettingsAO1P1 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m VnicVlanSettingsAO1P1) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// Default VLAN ID of the virtual interface. Setting the ID to 0 will not associate any default VLAN to the traffic on the virtual interface.
		//
		DefaultVlan int64 `json:"DefaultVlan,omitempty"`

		// Option to determine if the port can carry single VLAN (Access) or multiple VLANs (Trunk) traffic.
		//
		// Enum: [ACCESS TRUNK]
		Mode *string `json:"Mode,omitempty"`
	}

	stage1.DefaultVlan = m.DefaultVlan

	stage1.Mode = m.Mode

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.VnicVlanSettingsAO1P1) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.VnicVlanSettingsAO1P1)
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

// Validate validates this vnic vlan settings a o1 p1
func (m *VnicVlanSettingsAO1P1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var vnicVlanSettingsAO1P1TypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ACCESS","TRUNK"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vnicVlanSettingsAO1P1TypeModePropEnum = append(vnicVlanSettingsAO1P1TypeModePropEnum, v)
	}
}

const (

	// VnicVlanSettingsAO1P1ModeACCESS captures enum value "ACCESS"
	VnicVlanSettingsAO1P1ModeACCESS string = "ACCESS"

	// VnicVlanSettingsAO1P1ModeTRUNK captures enum value "TRUNK"
	VnicVlanSettingsAO1P1ModeTRUNK string = "TRUNK"
)

// prop value enum
func (m *VnicVlanSettingsAO1P1) validateModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, vnicVlanSettingsAO1P1TypeModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *VnicVlanSettingsAO1P1) validateMode(formats strfmt.Registry) error {

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
func (m *VnicVlanSettingsAO1P1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VnicVlanSettingsAO1P1) UnmarshalBinary(b []byte) error {
	var res VnicVlanSettingsAO1P1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
