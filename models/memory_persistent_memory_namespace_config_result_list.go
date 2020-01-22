// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// MemoryPersistentMemoryNamespaceConfigResultList memory persistent memory namespace config result list
// swagger:model memoryPersistentMemoryNamespaceConfigResultList
type MemoryPersistentMemoryNamespaceConfigResultList struct {

	// The number of memoryPersistentMemoryNamespaceConfigResults matching your request in total for all pages.
	Count int32 `json:"Count,omitempty"`

	// The array of memoryPersistentMemoryNamespaceConfigResults matching your request.
	Results []*MemoryPersistentMemoryNamespaceConfigResult `json:"Results"`
}

// Validate validates this memory persistent memory namespace config result list
func (m *MemoryPersistentMemoryNamespaceConfigResultList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MemoryPersistentMemoryNamespaceConfigResultList) validateResults(formats strfmt.Registry) error {

	if swag.IsZero(m.Results) { // not required
		return nil
	}

	for i := 0; i < len(m.Results); i++ {
		if swag.IsZero(m.Results[i]) { // not required
			continue
		}

		if m.Results[i] != nil {
			if err := m.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MemoryPersistentMemoryNamespaceConfigResultList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MemoryPersistentMemoryNamespaceConfigResultList) UnmarshalBinary(b []byte) error {
	var res MemoryPersistentMemoryNamespaceConfigResultList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}