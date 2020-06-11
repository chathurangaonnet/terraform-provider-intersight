// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StorageBaseCapacity Storage:Base Capacity
//
// Storage capacity information which includes, total capacity, available capacity, used capacity and free capacity.
//
// swagger:model storageBaseCapacity
type StorageBaseCapacity struct {
	MoBaseComplexType

	StorageBaseCapacityAO1P1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *StorageBaseCapacity) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseComplexType
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseComplexType = aO0

	// AO1
	var aO1 StorageBaseCapacityAO1P1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.StorageBaseCapacityAO1P1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m StorageBaseCapacity) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseComplexType)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.StorageBaseCapacityAO1P1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this storage base capacity
func (m *StorageBaseCapacity) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseComplexType
	if err := m.MoBaseComplexType.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with StorageBaseCapacityAO1P1
	if err := m.StorageBaseCapacityAO1P1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *StorageBaseCapacity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageBaseCapacity) UnmarshalBinary(b []byte) error {
	var res StorageBaseCapacity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// StorageBaseCapacityAO1P1 storage base capacity a o1 p1
//
// swagger:model StorageBaseCapacityAO1P1
type StorageBaseCapacityAO1P1 struct {

	// Total consumable storage capacity represented in bytes. System may reserve some space for internal purposes which is excluded from total capacity.
	// Read Only: true
	Available int64 `json:"Available,omitempty"`

	// Unused space available for applications to consume, represented in bytes.
	// Read Only: true
	Free int64 `json:"Free,omitempty"`

	// Total storage capacity, represented in bytes. It is set by the component manufacturer.
	// Read Only: true
	Total int64 `json:"Total,omitempty"`

	// Used or consumed storage capacity, represented in bytes.
	// Read Only: true
	Used int64 `json:"Used,omitempty"`

	// storage base capacity a o1 p1
	StorageBaseCapacityAO1P1 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *StorageBaseCapacityAO1P1) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// Total consumable storage capacity represented in bytes. System may reserve some space for internal purposes which is excluded from total capacity.
		// Read Only: true
		Available int64 `json:"Available,omitempty"`

		// Unused space available for applications to consume, represented in bytes.
		// Read Only: true
		Free int64 `json:"Free,omitempty"`

		// Total storage capacity, represented in bytes. It is set by the component manufacturer.
		// Read Only: true
		Total int64 `json:"Total,omitempty"`

		// Used or consumed storage capacity, represented in bytes.
		// Read Only: true
		Used int64 `json:"Used,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv StorageBaseCapacityAO1P1

	rcv.Available = stage1.Available
	rcv.Free = stage1.Free
	rcv.Total = stage1.Total
	rcv.Used = stage1.Used
	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "Available")
	delete(stage2, "Free")
	delete(stage2, "Total")
	delete(stage2, "Used")
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
		m.StorageBaseCapacityAO1P1 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m StorageBaseCapacityAO1P1) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// Total consumable storage capacity represented in bytes. System may reserve some space for internal purposes which is excluded from total capacity.
		// Read Only: true
		Available int64 `json:"Available,omitempty"`

		// Unused space available for applications to consume, represented in bytes.
		// Read Only: true
		Free int64 `json:"Free,omitempty"`

		// Total storage capacity, represented in bytes. It is set by the component manufacturer.
		// Read Only: true
		Total int64 `json:"Total,omitempty"`

		// Used or consumed storage capacity, represented in bytes.
		// Read Only: true
		Used int64 `json:"Used,omitempty"`
	}

	stage1.Available = m.Available
	stage1.Free = m.Free
	stage1.Total = m.Total
	stage1.Used = m.Used

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.StorageBaseCapacityAO1P1) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.StorageBaseCapacityAO1P1)
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

// Validate validates this storage base capacity a o1 p1
func (m *StorageBaseCapacityAO1P1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StorageBaseCapacityAO1P1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageBaseCapacityAO1P1) UnmarshalBinary(b []byte) error {
	var res StorageBaseCapacityAO1P1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}