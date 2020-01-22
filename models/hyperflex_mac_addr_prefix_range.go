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

// HyperflexMacAddrPrefixRange Hyperflex:Mac Addr Prefix Range
//
// A MAC address prefix range.
//
// The range is inclusive and comprised of a start and end MAC addresses.
// A single address can be specified by setting it as the start and end of the range.
//
// swagger:model hyperflexMacAddrPrefixRange
type HyperflexMacAddrPrefixRange struct {
	MoBaseComplexType

	HyperflexMacAddrPrefixRangeAO1P1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *HyperflexMacAddrPrefixRange) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseComplexType
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseComplexType = aO0

	// AO1
	var aO1 HyperflexMacAddrPrefixRangeAO1P1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.HyperflexMacAddrPrefixRangeAO1P1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m HyperflexMacAddrPrefixRange) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseComplexType)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.HyperflexMacAddrPrefixRangeAO1P1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this hyperflex mac addr prefix range
func (m *HyperflexMacAddrPrefixRange) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseComplexType
	if err := m.MoBaseComplexType.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with HyperflexMacAddrPrefixRangeAO1P1
	if err := m.HyperflexMacAddrPrefixRangeAO1P1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *HyperflexMacAddrPrefixRange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HyperflexMacAddrPrefixRange) UnmarshalBinary(b []byte) error {
	var res HyperflexMacAddrPrefixRange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HyperflexMacAddrPrefixRangeAO1P1 hyperflex mac addr prefix range a o1 p1
// swagger:model HyperflexMacAddrPrefixRangeAO1P1
type HyperflexMacAddrPrefixRangeAO1P1 struct {

	// The end MAC address prefix of a MAC address prefix range in the form of 00:25:B5:XX.
	//
	EndAddr string `json:"EndAddr,omitempty"`

	// The start MAC address prefix of a MAC address prefix range in the form of 00:25:B5:XX.
	//
	StartAddr string `json:"StartAddr,omitempty"`

	// hyperflex mac addr prefix range a o1 p1
	HyperflexMacAddrPrefixRangeAO1P1 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *HyperflexMacAddrPrefixRangeAO1P1) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// The end MAC address prefix of a MAC address prefix range in the form of 00:25:B5:XX.
		//
		EndAddr string `json:"EndAddr,omitempty"`

		// The start MAC address prefix of a MAC address prefix range in the form of 00:25:B5:XX.
		//
		StartAddr string `json:"StartAddr,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv HyperflexMacAddrPrefixRangeAO1P1

	rcv.EndAddr = stage1.EndAddr

	rcv.StartAddr = stage1.StartAddr

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "EndAddr")

	delete(stage2, "StartAddr")

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
		m.HyperflexMacAddrPrefixRangeAO1P1 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m HyperflexMacAddrPrefixRangeAO1P1) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// The end MAC address prefix of a MAC address prefix range in the form of 00:25:B5:XX.
		//
		EndAddr string `json:"EndAddr,omitempty"`

		// The start MAC address prefix of a MAC address prefix range in the form of 00:25:B5:XX.
		//
		StartAddr string `json:"StartAddr,omitempty"`
	}

	stage1.EndAddr = m.EndAddr

	stage1.StartAddr = m.StartAddr

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.HyperflexMacAddrPrefixRangeAO1P1) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.HyperflexMacAddrPrefixRangeAO1P1)
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

// Validate validates this hyperflex mac addr prefix range a o1 p1
func (m *HyperflexMacAddrPrefixRangeAO1P1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HyperflexMacAddrPrefixRangeAO1P1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HyperflexMacAddrPrefixRangeAO1P1) UnmarshalBinary(b []byte) error {
	var res HyperflexMacAddrPrefixRangeAO1P1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
