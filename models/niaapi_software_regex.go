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

// NiaapiSoftwareRegex Niaapi:Software Regex
//
// The regular expression to parse the software version strings.
//
// swagger:model niaapiSoftwareRegex
type NiaapiSoftwareRegex struct {
	NiaapiSoftwareRegexAO0P0
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *NiaapiSoftwareRegex) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 NiaapiSoftwareRegexAO0P0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.NiaapiSoftwareRegexAO0P0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m NiaapiSoftwareRegex) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.NiaapiSoftwareRegexAO0P0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this niaapi software regex
func (m *NiaapiSoftwareRegex) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with NiaapiSoftwareRegexAO0P0
	if err := m.NiaapiSoftwareRegexAO0P0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *NiaapiSoftwareRegex) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NiaapiSoftwareRegex) UnmarshalBinary(b []byte) error {
	var res NiaapiSoftwareRegex
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NiaapiSoftwareRegexAO0P0 niaapi software regex a o0 p0
// swagger:model NiaapiSoftwareRegexAO0P0
type NiaapiSoftwareRegexAO0P0 struct {

	// The concrete type of this complex type.
	//
	// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
	// ObjectType is optional.
	// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
	// are heterogeneous, i.e. the array can contain nested documents of different types.
	//
	//
	ObjectType string `json:"ObjectType,omitempty"`

	// Regular Expression pattern used to reconginze the version string.
	//
	Regex string `json:"Regex,omitempty"`

	// Software release. A set of Software releases seperated by comma which can be recongized by according Regex pattern.
	//
	SoftwareVersion string `json:"SoftwareVersion,omitempty"`

	// niaapi software regex a o0 p0
	NiaapiSoftwareRegexAO0P0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *NiaapiSoftwareRegexAO0P0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// Regular Expression pattern used to reconginze the version string.
		//
		Regex string `json:"Regex,omitempty"`

		// Software release. A set of Software releases seperated by comma which can be recongized by according Regex pattern.
		//
		SoftwareVersion string `json:"SoftwareVersion,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv NiaapiSoftwareRegexAO0P0

	rcv.ObjectType = stage1.ObjectType

	rcv.Regex = stage1.Regex

	rcv.SoftwareVersion = stage1.SoftwareVersion

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "ObjectType")

	delete(stage2, "Regex")

	delete(stage2, "SoftwareVersion")

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
		m.NiaapiSoftwareRegexAO0P0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m NiaapiSoftwareRegexAO0P0) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// Regular Expression pattern used to reconginze the version string.
		//
		Regex string `json:"Regex,omitempty"`

		// Software release. A set of Software releases seperated by comma which can be recongized by according Regex pattern.
		//
		SoftwareVersion string `json:"SoftwareVersion,omitempty"`
	}

	stage1.ObjectType = m.ObjectType

	stage1.Regex = m.Regex

	stage1.SoftwareVersion = m.SoftwareVersion

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.NiaapiSoftwareRegexAO0P0) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.NiaapiSoftwareRegexAO0P0)
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

// Validate validates this niaapi software regex a o0 p0
func (m *NiaapiSoftwareRegexAO0P0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NiaapiSoftwareRegexAO0P0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NiaapiSoftwareRegexAO0P0) UnmarshalBinary(b []byte) error {
	var res NiaapiSoftwareRegexAO0P0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}