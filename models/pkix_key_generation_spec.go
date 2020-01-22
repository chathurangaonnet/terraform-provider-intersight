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

// PkixKeyGenerationSpec Pkix:Key Generation Spec
//
// The key generation spec provides the algorithm and the parameters required for this algorithm to generate a key.
//
// swagger:model pkixKeyGenerationSpec
type PkixKeyGenerationSpec struct {
	MoBaseComplexType

	PkixKeyGenerationSpecAO1P1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PkixKeyGenerationSpec) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseComplexType
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseComplexType = aO0

	// AO1
	var aO1 PkixKeyGenerationSpecAO1P1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.PkixKeyGenerationSpecAO1P1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PkixKeyGenerationSpec) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseComplexType)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.PkixKeyGenerationSpecAO1P1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this pkix key generation spec
func (m *PkixKeyGenerationSpec) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseComplexType
	if err := m.MoBaseComplexType.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with PkixKeyGenerationSpecAO1P1
	if err := m.PkixKeyGenerationSpecAO1P1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PkixKeyGenerationSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PkixKeyGenerationSpec) UnmarshalBinary(b []byte) error {
	var res PkixKeyGenerationSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PkixKeyGenerationSpecAO1P1 pkix key generation spec a o1 p1
// swagger:model PkixKeyGenerationSpecAO1P1
type PkixKeyGenerationSpecAO1P1 struct {

	// Name of the key generation algorithm.
	//
	// Read Only: true
	// Enum: [RSA]
	Name string `json:"Name,omitempty"`

	// pkix key generation spec a o1 p1
	PkixKeyGenerationSpecAO1P1 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *PkixKeyGenerationSpecAO1P1) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// Name of the key generation algorithm.
		//
		// Read Only: true
		// Enum: [RSA]
		Name string `json:"Name,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv PkixKeyGenerationSpecAO1P1

	rcv.Name = stage1.Name

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "Name")

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
		m.PkixKeyGenerationSpecAO1P1 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m PkixKeyGenerationSpecAO1P1) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// Name of the key generation algorithm.
		//
		// Read Only: true
		// Enum: [RSA]
		Name string `json:"Name,omitempty"`
	}

	stage1.Name = m.Name

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.PkixKeyGenerationSpecAO1P1) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.PkixKeyGenerationSpecAO1P1)
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

// Validate validates this pkix key generation spec a o1 p1
func (m *PkixKeyGenerationSpecAO1P1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var pkixKeyGenerationSpecAO1P1TypeNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["RSA"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pkixKeyGenerationSpecAO1P1TypeNamePropEnum = append(pkixKeyGenerationSpecAO1P1TypeNamePropEnum, v)
	}
}

const (

	// PkixKeyGenerationSpecAO1P1NameRSA captures enum value "RSA"
	PkixKeyGenerationSpecAO1P1NameRSA string = "RSA"
)

// prop value enum
func (m *PkixKeyGenerationSpecAO1P1) validateNameEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, pkixKeyGenerationSpecAO1P1TypeNamePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PkixKeyGenerationSpecAO1P1) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	// value enum
	if err := m.validateNameEnum("Name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PkixKeyGenerationSpecAO1P1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PkixKeyGenerationSpecAO1P1) UnmarshalBinary(b []byte) error {
	var res PkixKeyGenerationSpecAO1P1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}