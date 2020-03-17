// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HyperflexFeatureLimitExternal Hyperflex:Feature Limit External
//
// The HyperFlex feature limits that are available to end users.
//
// swagger:model hyperflexFeatureLimitExternal
type HyperflexFeatureLimitExternal struct {
	MoBaseMo

	// A collection of references to the [hyperflex.AppCatalog](mo://hyperflex.AppCatalog) Managed Object.
	// When this managed object is deleted, the referenced [hyperflex.AppCatalog](mo://hyperflex.AppCatalog) MO unsets its reference to this deleted MO.
	AppCatalog *HyperflexAppCatalogRef `json:"AppCatalog,omitempty"`

	// The HyperFlex feature limits which are made available to users.
	FeatureLimitEntries []*HyperflexFeatureLimitEntry `json:"FeatureLimitEntries"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *HyperflexFeatureLimitExternal) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		AppCatalog *HyperflexAppCatalogRef `json:"AppCatalog,omitempty"`

		FeatureLimitEntries []*HyperflexFeatureLimitEntry `json:"FeatureLimitEntries"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.AppCatalog = dataAO1.AppCatalog

	m.FeatureLimitEntries = dataAO1.FeatureLimitEntries

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m HyperflexFeatureLimitExternal) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		AppCatalog *HyperflexAppCatalogRef `json:"AppCatalog,omitempty"`

		FeatureLimitEntries []*HyperflexFeatureLimitEntry `json:"FeatureLimitEntries"`
	}

	dataAO1.AppCatalog = m.AppCatalog

	dataAO1.FeatureLimitEntries = m.FeatureLimitEntries

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this hyperflex feature limit external
func (m *HyperflexFeatureLimitExternal) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAppCatalog(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeatureLimitEntries(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HyperflexFeatureLimitExternal) validateAppCatalog(formats strfmt.Registry) error {

	if swag.IsZero(m.AppCatalog) { // not required
		return nil
	}

	if m.AppCatalog != nil {
		if err := m.AppCatalog.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("AppCatalog")
			}
			return err
		}
	}

	return nil
}

func (m *HyperflexFeatureLimitExternal) validateFeatureLimitEntries(formats strfmt.Registry) error {

	if swag.IsZero(m.FeatureLimitEntries) { // not required
		return nil
	}

	for i := 0; i < len(m.FeatureLimitEntries); i++ {
		if swag.IsZero(m.FeatureLimitEntries[i]) { // not required
			continue
		}

		if m.FeatureLimitEntries[i] != nil {
			if err := m.FeatureLimitEntries[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("FeatureLimitEntries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HyperflexFeatureLimitExternal) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HyperflexFeatureLimitExternal) UnmarshalBinary(b []byte) error {
	var res HyperflexFeatureLimitExternal
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
