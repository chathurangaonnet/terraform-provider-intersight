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

// IamEndPointUserPolicy Local User
//
// Enables creation of local users on endpoints.
//
// swagger:model iamEndPointUserPolicy
type IamEndPointUserPolicy struct {
	PolicyAbstractPolicy

	// Relationship to the collection of Endpoint user roles.
	EndPointUserRoles []*IamEndPointUserRoleRef `json:"EndPointUserRoles"`

	// Relationship to the Organization that owns the Managed Object.
	Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`

	// Set password properties for endpoint users.
	PasswordProperties *IamEndPointPasswordProperties `json:"PasswordProperties,omitempty"`

	// Relationship to the server profile object.
	Profiles []*PolicyAbstractConfigProfileRef `json:"Profiles"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *IamEndPointUserPolicy) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 PolicyAbstractPolicy
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.PolicyAbstractPolicy = aO0

	// AO1
	var dataAO1 struct {
		EndPointUserRoles []*IamEndPointUserRoleRef `json:"EndPointUserRoles"`

		Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`

		PasswordProperties *IamEndPointPasswordProperties `json:"PasswordProperties,omitempty"`

		Profiles []*PolicyAbstractConfigProfileRef `json:"Profiles"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.EndPointUserRoles = dataAO1.EndPointUserRoles

	m.Organization = dataAO1.Organization

	m.PasswordProperties = dataAO1.PasswordProperties

	m.Profiles = dataAO1.Profiles

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m IamEndPointUserPolicy) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.PolicyAbstractPolicy)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		EndPointUserRoles []*IamEndPointUserRoleRef `json:"EndPointUserRoles"`

		Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`

		PasswordProperties *IamEndPointPasswordProperties `json:"PasswordProperties,omitempty"`

		Profiles []*PolicyAbstractConfigProfileRef `json:"Profiles"`
	}

	dataAO1.EndPointUserRoles = m.EndPointUserRoles

	dataAO1.Organization = m.Organization

	dataAO1.PasswordProperties = m.PasswordProperties

	dataAO1.Profiles = m.Profiles

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this iam end point user policy
func (m *IamEndPointUserPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with PolicyAbstractPolicy
	if err := m.PolicyAbstractPolicy.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndPointUserRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePasswordProperties(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProfiles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IamEndPointUserPolicy) validateEndPointUserRoles(formats strfmt.Registry) error {

	if swag.IsZero(m.EndPointUserRoles) { // not required
		return nil
	}

	for i := 0; i < len(m.EndPointUserRoles); i++ {
		if swag.IsZero(m.EndPointUserRoles[i]) { // not required
			continue
		}

		if m.EndPointUserRoles[i] != nil {
			if err := m.EndPointUserRoles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("EndPointUserRoles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IamEndPointUserPolicy) validateOrganization(formats strfmt.Registry) error {

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

func (m *IamEndPointUserPolicy) validatePasswordProperties(formats strfmt.Registry) error {

	if swag.IsZero(m.PasswordProperties) { // not required
		return nil
	}

	if m.PasswordProperties != nil {
		if err := m.PasswordProperties.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PasswordProperties")
			}
			return err
		}
	}

	return nil
}

func (m *IamEndPointUserPolicy) validateProfiles(formats strfmt.Registry) error {

	if swag.IsZero(m.Profiles) { // not required
		return nil
	}

	for i := 0; i < len(m.Profiles); i++ {
		if swag.IsZero(m.Profiles[i]) { // not required
			continue
		}

		if m.Profiles[i] != nil {
			if err := m.Profiles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Profiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IamEndPointUserPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IamEndPointUserPolicy) UnmarshalBinary(b []byte) error {
	var res IamEndPointUserPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
