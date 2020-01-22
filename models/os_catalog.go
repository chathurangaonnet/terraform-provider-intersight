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

// OsCatalog Os:Catalog
//
// A catalog of operating systems related objects such as configuration files and post install scripts. Each user account will have a local OS catalog where account users can store their private configuration files and post install scripts.
// Cisco provides validated answer files and post install scripts to Intersight users via shared catalogs. Intersight users will be able to read, use these files and scripts in OS installation within their account context. The shared catalogs will be managed entirely by Cisco. Contributions to shared catalogs will need to be provided to Cisco who will publish them at their own discretion.
//
// swagger:model osCatalog
type OsCatalog struct {
	MoBaseMo

	// This captures the associated Configuration files.
	//
	ConfigurationFiles []*OsConfigurationFileRef `json:"ConfigurationFiles"`

	// The catalog name. There will be one for system and one for each user account.
	//
	Name string `json:"Name,omitempty"`

	// Relationship to the Organization that owns the Managed Object.
	//
	Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *OsCatalog) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		ConfigurationFiles []*OsConfigurationFileRef `json:"ConfigurationFiles"`

		Name string `json:"Name,omitempty"`

		Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ConfigurationFiles = dataAO1.ConfigurationFiles

	m.Name = dataAO1.Name

	m.Organization = dataAO1.Organization

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m OsCatalog) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		ConfigurationFiles []*OsConfigurationFileRef `json:"ConfigurationFiles"`

		Name string `json:"Name,omitempty"`

		Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`
	}

	dataAO1.ConfigurationFiles = m.ConfigurationFiles

	dataAO1.Name = m.Name

	dataAO1.Organization = m.Organization

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this os catalog
func (m *OsCatalog) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfigurationFiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OsCatalog) validateConfigurationFiles(formats strfmt.Registry) error {

	if swag.IsZero(m.ConfigurationFiles) { // not required
		return nil
	}

	for i := 0; i < len(m.ConfigurationFiles); i++ {
		if swag.IsZero(m.ConfigurationFiles[i]) { // not required
			continue
		}

		if m.ConfigurationFiles[i] != nil {
			if err := m.ConfigurationFiles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ConfigurationFiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OsCatalog) validateOrganization(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *OsCatalog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OsCatalog) UnmarshalBinary(b []byte) error {
	var res OsCatalog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
