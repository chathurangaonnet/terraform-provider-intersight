// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ApplianceKeyValuePair Appliance:Key Value Pair
//
// An arbitrary key and value pair that can be used to for having pair of key and values in application such as Appliance Capabilities.
//
// swagger:model applianceKeyValuePair
type ApplianceKeyValuePair struct {
	MoBaseComplexType

	ApplianceKeyValuePairAO1P1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ApplianceKeyValuePair) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseComplexType
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseComplexType = aO0

	// AO1
	var aO1 ApplianceKeyValuePairAO1P1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.ApplianceKeyValuePairAO1P1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ApplianceKeyValuePair) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseComplexType)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.ApplianceKeyValuePairAO1P1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this appliance key value pair
func (m *ApplianceKeyValuePair) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseComplexType
	if err := m.MoBaseComplexType.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with ApplianceKeyValuePairAO1P1
	if err := m.ApplianceKeyValuePairAO1P1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ApplianceKeyValuePair) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplianceKeyValuePair) UnmarshalBinary(b []byte) error {
	var res ApplianceKeyValuePair
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ApplianceKeyValuePairAO1P1 appliance key value pair a o1 p1
// swagger:model ApplianceKeyValuePairAO1P1
type ApplianceKeyValuePairAO1P1 struct {

	// The string representation of a tag key.
	//
	// Read Only: true
	Key string `json:"Key,omitempty"`

	// The string representation of a tag value.
	//
	// Read Only: true
	Value string `json:"Value,omitempty"`

	// appliance key value pair a o1 p1
	ApplianceKeyValuePairAO1P1 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *ApplianceKeyValuePairAO1P1) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// The string representation of a tag key.
		//
		// Read Only: true
		Key string `json:"Key,omitempty"`

		// The string representation of a tag value.
		//
		// Read Only: true
		Value string `json:"Value,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv ApplianceKeyValuePairAO1P1

	rcv.Key = stage1.Key

	rcv.Value = stage1.Value

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "Key")

	delete(stage2, "Value")

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
		m.ApplianceKeyValuePairAO1P1 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m ApplianceKeyValuePairAO1P1) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// The string representation of a tag key.
		//
		// Read Only: true
		Key string `json:"Key,omitempty"`

		// The string representation of a tag value.
		//
		// Read Only: true
		Value string `json:"Value,omitempty"`
	}

	stage1.Key = m.Key

	stage1.Value = m.Value

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.ApplianceKeyValuePairAO1P1) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.ApplianceKeyValuePairAO1P1)
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

// Validate validates this appliance key value pair a o1 p1
func (m *ApplianceKeyValuePairAO1P1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ApplianceKeyValuePairAO1P1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplianceKeyValuePairAO1P1) UnmarshalBinary(b []byte) error {
	var res ApplianceKeyValuePairAO1P1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
