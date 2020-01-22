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

// HyperflexExtFcStoragePolicy External FC Storage
//
// A policy specifying external storage connectivity information via Fabric attached FC storage.
//
// swagger:model hyperflexExtFcStoragePolicy
type HyperflexExtFcStoragePolicy struct {
	PolicyAbstractPolicy

	// Enables or disables external FC storage configuration.
	//
	AdminState *bool `json:"AdminState,omitempty"`

	// List of cluster profiles using this policy.
	//
	ClusterProfiles []*HyperflexClusterProfileRef `json:"ClusterProfiles"`

	// VSAN for the primary Fabric Interconnect external FC storage traffic.
	//
	ExtaTraffic *HyperflexNamedVsan `json:"ExtaTraffic,omitempty"`

	// VSAN for the secondary Fabric Interconnect external FC storage traffic.
	//
	ExtbTraffic *HyperflexNamedVsan `json:"ExtbTraffic,omitempty"`

	// Relationship to the Organization that owns the Managed Object.
	//
	Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`

	// The range of WWxN addresses to use for the FC storage configuration.
	//
	WwxnPrefixRange *HyperflexWwxnPrefixRange `json:"WwxnPrefixRange,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *HyperflexExtFcStoragePolicy) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 PolicyAbstractPolicy
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.PolicyAbstractPolicy = aO0

	// AO1
	var dataAO1 struct {
		AdminState *bool `json:"AdminState,omitempty"`

		ClusterProfiles []*HyperflexClusterProfileRef `json:"ClusterProfiles"`

		ExtaTraffic *HyperflexNamedVsan `json:"ExtaTraffic,omitempty"`

		ExtbTraffic *HyperflexNamedVsan `json:"ExtbTraffic,omitempty"`

		Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`

		WwxnPrefixRange *HyperflexWwxnPrefixRange `json:"WwxnPrefixRange,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.AdminState = dataAO1.AdminState

	m.ClusterProfiles = dataAO1.ClusterProfiles

	m.ExtaTraffic = dataAO1.ExtaTraffic

	m.ExtbTraffic = dataAO1.ExtbTraffic

	m.Organization = dataAO1.Organization

	m.WwxnPrefixRange = dataAO1.WwxnPrefixRange

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m HyperflexExtFcStoragePolicy) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.PolicyAbstractPolicy)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		AdminState *bool `json:"AdminState,omitempty"`

		ClusterProfiles []*HyperflexClusterProfileRef `json:"ClusterProfiles"`

		ExtaTraffic *HyperflexNamedVsan `json:"ExtaTraffic,omitempty"`

		ExtbTraffic *HyperflexNamedVsan `json:"ExtbTraffic,omitempty"`

		Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`

		WwxnPrefixRange *HyperflexWwxnPrefixRange `json:"WwxnPrefixRange,omitempty"`
	}

	dataAO1.AdminState = m.AdminState

	dataAO1.ClusterProfiles = m.ClusterProfiles

	dataAO1.ExtaTraffic = m.ExtaTraffic

	dataAO1.ExtbTraffic = m.ExtbTraffic

	dataAO1.Organization = m.Organization

	dataAO1.WwxnPrefixRange = m.WwxnPrefixRange

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this hyperflex ext fc storage policy
func (m *HyperflexExtFcStoragePolicy) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with PolicyAbstractPolicy
	if err := m.PolicyAbstractPolicy.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterProfiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtaTraffic(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtbTraffic(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWwxnPrefixRange(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HyperflexExtFcStoragePolicy) validateClusterProfiles(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterProfiles) { // not required
		return nil
	}

	for i := 0; i < len(m.ClusterProfiles); i++ {
		if swag.IsZero(m.ClusterProfiles[i]) { // not required
			continue
		}

		if m.ClusterProfiles[i] != nil {
			if err := m.ClusterProfiles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ClusterProfiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HyperflexExtFcStoragePolicy) validateExtaTraffic(formats strfmt.Registry) error {

	if swag.IsZero(m.ExtaTraffic) { // not required
		return nil
	}

	if m.ExtaTraffic != nil {
		if err := m.ExtaTraffic.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ExtaTraffic")
			}
			return err
		}
	}

	return nil
}

func (m *HyperflexExtFcStoragePolicy) validateExtbTraffic(formats strfmt.Registry) error {

	if swag.IsZero(m.ExtbTraffic) { // not required
		return nil
	}

	if m.ExtbTraffic != nil {
		if err := m.ExtbTraffic.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ExtbTraffic")
			}
			return err
		}
	}

	return nil
}

func (m *HyperflexExtFcStoragePolicy) validateOrganization(formats strfmt.Registry) error {

	if swag.IsZero(m.Organization) { // not required
		return nil
	}

	if m.Organization != nil {
		if err := m.Organization.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Organization")
			}
			return err
		}
	}

	return nil
}

func (m *HyperflexExtFcStoragePolicy) validateWwxnPrefixRange(formats strfmt.Registry) error {

	if swag.IsZero(m.WwxnPrefixRange) { // not required
		return nil
	}

	if m.WwxnPrefixRange != nil {
		if err := m.WwxnPrefixRange.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("WwxnPrefixRange")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HyperflexExtFcStoragePolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HyperflexExtFcStoragePolicy) UnmarshalBinary(b []byte) error {
	var res HyperflexExtFcStoragePolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
