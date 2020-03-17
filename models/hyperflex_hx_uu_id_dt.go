// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HyperflexHxUuIDDt Hyperflex:Hx Uu Id Dt
//
// swagger:model hyperflexHxUuIdDt
type HyperflexHxUuIDDt struct {
	MoBaseComplexType

	HyperflexHxUuIDDtAO1P1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *HyperflexHxUuIDDt) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseComplexType
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseComplexType = aO0

	// AO1
	var aO1 HyperflexHxUuIDDtAO1P1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.HyperflexHxUuIDDtAO1P1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m HyperflexHxUuIDDt) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseComplexType)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.HyperflexHxUuIDDtAO1P1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this hyperflex hx uu Id dt
func (m *HyperflexHxUuIDDt) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseComplexType
	if err := m.MoBaseComplexType.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with HyperflexHxUuIDDtAO1P1
	if err := m.HyperflexHxUuIDDtAO1P1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *HyperflexHxUuIDDt) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HyperflexHxUuIDDt) UnmarshalBinary(b []byte) error {
	var res HyperflexHxUuIDDt
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HyperflexHxUuIDDtAO1P1 hyperflex hx uu ID dt a o1 p1
//
// swagger:model HyperflexHxUuIDDtAO1P1
type HyperflexHxUuIDDtAO1P1 struct {

	// links
	// Read Only: true
	Links []*HyperflexHxLinkDt `json:"Links"`

	// Uuid
	// Read Only: true
	UUID string `json:"Uuid,omitempty"`

	// hyperflex hx uu ID dt a o1 p1
	HyperflexHxUuIDDtAO1P1 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *HyperflexHxUuIDDtAO1P1) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// links
		// Read Only: true
		Links []*HyperflexHxLinkDt `json:"Links"`

		// Uuid
		// Read Only: true
		UUID string `json:"Uuid,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv HyperflexHxUuIDDtAO1P1

	rcv.Links = stage1.Links
	rcv.UUID = stage1.UUID
	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "Links")
	delete(stage2, "Uuid")
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
		m.HyperflexHxUuIDDtAO1P1 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m HyperflexHxUuIDDtAO1P1) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// links
		// Read Only: true
		Links []*HyperflexHxLinkDt `json:"Links"`

		// Uuid
		// Read Only: true
		UUID string `json:"Uuid,omitempty"`
	}

	stage1.Links = m.Links
	stage1.UUID = m.UUID

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.HyperflexHxUuIDDtAO1P1) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.HyperflexHxUuIDDtAO1P1)
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

// Validate validates this hyperflex hx uu ID dt a o1 p1
func (m *HyperflexHxUuIDDtAO1P1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HyperflexHxUuIDDtAO1P1) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	for i := 0; i < len(m.Links); i++ {
		if swag.IsZero(m.Links[i]) { // not required
			continue
		}

		if m.Links[i] != nil {
			if err := m.Links[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Links" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HyperflexHxUuIDDtAO1P1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HyperflexHxUuIDDtAO1P1) UnmarshalBinary(b []byte) error {
	var res HyperflexHxUuIDDtAO1P1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
