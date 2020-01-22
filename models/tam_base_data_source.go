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

// TamBaseDataSource Tam:Base Data Source
//
// Represents source for the data needed to analyze the alerts. this could one of several supported source types (textFsmTemplates/Intersight API/external API). TextFsmTemplates and Intersight API are the only ones supported currently.
//
// swagger:model tamBaseDataSource
type TamBaseDataSource struct {
	MoBaseComplexType

	TamBaseDataSourceAO1P1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TamBaseDataSource) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseComplexType
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseComplexType = aO0

	// AO1
	var aO1 TamBaseDataSourceAO1P1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.TamBaseDataSourceAO1P1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TamBaseDataSource) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseComplexType)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.TamBaseDataSourceAO1P1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this tam base data source
func (m *TamBaseDataSource) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseComplexType
	if err := m.MoBaseComplexType.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TamBaseDataSourceAO1P1
	if err := m.TamBaseDataSourceAO1P1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *TamBaseDataSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TamBaseDataSource) UnmarshalBinary(b []byte) error {
	var res TamBaseDataSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TamBaseDataSourceAO1P1 tam base data source a o1 p1
// swagger:model TamBaseDataSourceAO1P1
type TamBaseDataSourceAO1P1 struct {

	// Name is used to unique identify and refer a given data source in an alert definition.
	//
	Name string `json:"Name,omitempty"`

	// Type of data source (for e.g. TextFsmTempalate based, Intersight API based etc.).
	//
	// Enum: [nxos intersightApi]
	Type *string `json:"Type,omitempty"`

	// tam base data source a o1 p1
	TamBaseDataSourceAO1P1 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *TamBaseDataSourceAO1P1) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// Name is used to unique identify and refer a given data source in an alert definition.
		//
		Name string `json:"Name,omitempty"`

		// Type of data source (for e.g. TextFsmTempalate based, Intersight API based etc.).
		//
		// Enum: [nxos intersightApi]
		Type *string `json:"Type,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv TamBaseDataSourceAO1P1

	rcv.Name = stage1.Name

	rcv.Type = stage1.Type

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "Name")

	delete(stage2, "Type")

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
		m.TamBaseDataSourceAO1P1 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m TamBaseDataSourceAO1P1) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// Name is used to unique identify and refer a given data source in an alert definition.
		//
		Name string `json:"Name,omitempty"`

		// Type of data source (for e.g. TextFsmTempalate based, Intersight API based etc.).
		//
		// Enum: [nxos intersightApi]
		Type *string `json:"Type,omitempty"`
	}

	stage1.Name = m.Name

	stage1.Type = m.Type

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.TamBaseDataSourceAO1P1) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.TamBaseDataSourceAO1P1)
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

// Validate validates this tam base data source a o1 p1
func (m *TamBaseDataSourceAO1P1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var tamBaseDataSourceAO1P1TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["nxos","intersightApi"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tamBaseDataSourceAO1P1TypeTypePropEnum = append(tamBaseDataSourceAO1P1TypeTypePropEnum, v)
	}
}

const (

	// TamBaseDataSourceAO1P1TypeNxos captures enum value "nxos"
	TamBaseDataSourceAO1P1TypeNxos string = "nxos"

	// TamBaseDataSourceAO1P1TypeIntersightAPI captures enum value "intersightApi"
	TamBaseDataSourceAO1P1TypeIntersightAPI string = "intersightApi"
)

// prop value enum
func (m *TamBaseDataSourceAO1P1) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, tamBaseDataSourceAO1P1TypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TamBaseDataSourceAO1P1) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("Type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TamBaseDataSourceAO1P1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TamBaseDataSourceAO1P1) UnmarshalBinary(b []byte) error {
	var res TamBaseDataSourceAO1P1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
