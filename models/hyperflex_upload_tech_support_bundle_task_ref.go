// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HyperflexUploadTechSupportBundleTaskRef hyperflex upload tech support bundle task ref
//
// swagger:model hyperflexUploadTechSupportBundleTaskRef
type HyperflexUploadTechSupportBundleTaskRef struct {
	MoMoRef
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *HyperflexUploadTechSupportBundleTaskRef) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoMoRef
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoMoRef = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m HyperflexUploadTechSupportBundleTaskRef) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.MoMoRef)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this hyperflex upload tech support bundle task ref
func (m *HyperflexUploadTechSupportBundleTaskRef) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoMoRef
	if err := m.MoMoRef.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *HyperflexUploadTechSupportBundleTaskRef) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HyperflexUploadTechSupportBundleTaskRef) UnmarshalBinary(b []byte) error {
	var res HyperflexUploadTechSupportBundleTaskRef
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}