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

// TamQueryEntry Tam:Query Entry
//
// Contains a set of queries, each with an integer priority. the queries are executed in the order of specified priority. The result of each query is used as an input to the query next in priority order.
//
// swagger:model tamQueryEntry
type TamQueryEntry struct {
	MoBaseComplexType

	TamQueryEntryAO1P1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TamQueryEntry) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseComplexType
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseComplexType = aO0

	// AO1
	var aO1 TamQueryEntryAO1P1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.TamQueryEntryAO1P1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TamQueryEntry) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseComplexType)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.TamQueryEntryAO1P1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this tam query entry
func (m *TamQueryEntry) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseComplexType
	if err := m.MoBaseComplexType.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TamQueryEntryAO1P1
	if err := m.TamQueryEntryAO1P1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *TamQueryEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TamQueryEntry) UnmarshalBinary(b []byte) error {
	var res TamQueryEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TamQueryEntryAO1P1 tam query entry a o1 p1
// swagger:model TamQueryEntryAO1P1
type TamQueryEntryAO1P1 struct {

	// Name is used to unique identify and result of the given query which can be used by subsequent queries as input data source.
	//
	Name string `json:"Name,omitempty"`

	// An integer value depicting the priority of the query among the queries that are part of the same QueryEntry collection.
	//
	Priority int64 `json:"Priority,omitempty"`

	// A SparkSQL query to be used on a given data source.
	//
	Query string `json:"Query,omitempty"`

	// tam query entry a o1 p1
	TamQueryEntryAO1P1 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *TamQueryEntryAO1P1) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// Name is used to unique identify and result of the given query which can be used by subsequent queries as input data source.
		//
		Name string `json:"Name,omitempty"`

		// An integer value depicting the priority of the query among the queries that are part of the same QueryEntry collection.
		//
		Priority int64 `json:"Priority,omitempty"`

		// A SparkSQL query to be used on a given data source.
		//
		Query string `json:"Query,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv TamQueryEntryAO1P1

	rcv.Name = stage1.Name

	rcv.Priority = stage1.Priority

	rcv.Query = stage1.Query

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "Name")

	delete(stage2, "Priority")

	delete(stage2, "Query")

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
		m.TamQueryEntryAO1P1 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m TamQueryEntryAO1P1) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// Name is used to unique identify and result of the given query which can be used by subsequent queries as input data source.
		//
		Name string `json:"Name,omitempty"`

		// An integer value depicting the priority of the query among the queries that are part of the same QueryEntry collection.
		//
		Priority int64 `json:"Priority,omitempty"`

		// A SparkSQL query to be used on a given data source.
		//
		Query string `json:"Query,omitempty"`
	}

	stage1.Name = m.Name

	stage1.Priority = m.Priority

	stage1.Query = m.Query

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.TamQueryEntryAO1P1) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.TamQueryEntryAO1P1)
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

// Validate validates this tam query entry a o1 p1
func (m *TamQueryEntryAO1P1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TamQueryEntryAO1P1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TamQueryEntryAO1P1) UnmarshalBinary(b []byte) error {
	var res TamQueryEntryAO1P1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
