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

// IamAccount Iam:Account
//
// The Intersight Account used to access Intersight.
//
// swagger:model iamAccount
type IamAccount struct {
	MoBaseMo

	// The domain Groups are configured in an account for scaling purpose. Currently, only onboarding-device account has multiple domain groups and other accounts have only one domain group per account.
	//
	// Read Only: true
	DomainGroups []*IamDomainGroupRef `json:"DomainGroups"`

	// User defined end point roles. These roles are assigned to Intersight users to perform end point operations such as GUI/CLI cross launch.
	//
	// Read Only: true
	EndPointRoles []*IamEndPointRoleRef `json:"EndPointRoles"`

	// System created IdPs configured for authentication in an account. By default Cisco IdP is created upon account creation.
	//
	// Read Only: true
	Idpreferences []*IamIdpReferenceRef `json:"Idpreferences"`

	// IdPs configured for authentication in an account. IdP object handles the third-party IdP details.
	//
	// Read Only: true
	Idps []*IamIdpRef `json:"Idps"`

	// Name of the Intersight account. By default, name is same as the MoID of the account.
	//
	Name string `json:"Name,omitempty"`

	// System defined permissions within an account. Permission provides a way to assign roles to a user or user group to perform operations on object hierarchy.
	//
	// Read Only: true
	Permissions []*IamPermissionRef `json:"Permissions"`

	// User defined privilege sets. Privilege set is a collection of privileges. Privilege sets are assigned to a user using roles.
	//
	// Read Only: true
	PrivilegeSets []*IamPrivilegeSetRef `json:"PrivilegeSets"`

	// Privileges are assigned to a user using privilege sets and roles. Privileges define user permissions and the actions a user can perform in Intersight.
	//
	// Read Only: true
	Privileges []*IamPrivilegeRef `json:"Privileges"`

	// User and user group related configuration limits.
	//
	// Read Only: true
	ResourceLimits *IamResourceLimitsRef `json:"ResourceLimits,omitempty"`

	// User defined roles created within an account. Role is a collection of privilege sets. Roles are assigned to user using permission object.
	//
	// Read Only: true
	Roles []*IamRoleRef `json:"Roles"`

	// Session related configuration limits.
	//
	// Read Only: true
	SessionLimits *IamSessionLimitsRef `json:"SessionLimits,omitempty"`

	// Status of the account. To activate the Intersight account, claim a device to the account.
	//
	// Read Only: true
	Status string `json:"Status,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *IamAccount) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		DomainGroups []*IamDomainGroupRef `json:"DomainGroups"`

		EndPointRoles []*IamEndPointRoleRef `json:"EndPointRoles"`

		Idpreferences []*IamIdpReferenceRef `json:"Idpreferences"`

		Idps []*IamIdpRef `json:"Idps"`

		Name string `json:"Name,omitempty"`

		Permissions []*IamPermissionRef `json:"Permissions"`

		PrivilegeSets []*IamPrivilegeSetRef `json:"PrivilegeSets"`

		Privileges []*IamPrivilegeRef `json:"Privileges"`

		ResourceLimits *IamResourceLimitsRef `json:"ResourceLimits,omitempty"`

		Roles []*IamRoleRef `json:"Roles"`

		SessionLimits *IamSessionLimitsRef `json:"SessionLimits,omitempty"`

		Status string `json:"Status,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.DomainGroups = dataAO1.DomainGroups

	m.EndPointRoles = dataAO1.EndPointRoles

	m.Idpreferences = dataAO1.Idpreferences

	m.Idps = dataAO1.Idps

	m.Name = dataAO1.Name

	m.Permissions = dataAO1.Permissions

	m.PrivilegeSets = dataAO1.PrivilegeSets

	m.Privileges = dataAO1.Privileges

	m.ResourceLimits = dataAO1.ResourceLimits

	m.Roles = dataAO1.Roles

	m.SessionLimits = dataAO1.SessionLimits

	m.Status = dataAO1.Status

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m IamAccount) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		DomainGroups []*IamDomainGroupRef `json:"DomainGroups"`

		EndPointRoles []*IamEndPointRoleRef `json:"EndPointRoles"`

		Idpreferences []*IamIdpReferenceRef `json:"Idpreferences"`

		Idps []*IamIdpRef `json:"Idps"`

		Name string `json:"Name,omitempty"`

		Permissions []*IamPermissionRef `json:"Permissions"`

		PrivilegeSets []*IamPrivilegeSetRef `json:"PrivilegeSets"`

		Privileges []*IamPrivilegeRef `json:"Privileges"`

		ResourceLimits *IamResourceLimitsRef `json:"ResourceLimits,omitempty"`

		Roles []*IamRoleRef `json:"Roles"`

		SessionLimits *IamSessionLimitsRef `json:"SessionLimits,omitempty"`

		Status string `json:"Status,omitempty"`
	}

	dataAO1.DomainGroups = m.DomainGroups

	dataAO1.EndPointRoles = m.EndPointRoles

	dataAO1.Idpreferences = m.Idpreferences

	dataAO1.Idps = m.Idps

	dataAO1.Name = m.Name

	dataAO1.Permissions = m.Permissions

	dataAO1.PrivilegeSets = m.PrivilegeSets

	dataAO1.Privileges = m.Privileges

	dataAO1.ResourceLimits = m.ResourceLimits

	dataAO1.Roles = m.Roles

	dataAO1.SessionLimits = m.SessionLimits

	dataAO1.Status = m.Status

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this iam account
func (m *IamAccount) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDomainGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndPointRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdpreferences(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePermissions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivilegeSets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivileges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceLimits(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSessionLimits(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IamAccount) validateDomainGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.DomainGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.DomainGroups); i++ {
		if swag.IsZero(m.DomainGroups[i]) { // not required
			continue
		}

		if m.DomainGroups[i] != nil {
			if err := m.DomainGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("DomainGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IamAccount) validateEndPointRoles(formats strfmt.Registry) error {

	if swag.IsZero(m.EndPointRoles) { // not required
		return nil
	}

	for i := 0; i < len(m.EndPointRoles); i++ {
		if swag.IsZero(m.EndPointRoles[i]) { // not required
			continue
		}

		if m.EndPointRoles[i] != nil {
			if err := m.EndPointRoles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("EndPointRoles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IamAccount) validateIdpreferences(formats strfmt.Registry) error {

	if swag.IsZero(m.Idpreferences) { // not required
		return nil
	}

	for i := 0; i < len(m.Idpreferences); i++ {
		if swag.IsZero(m.Idpreferences[i]) { // not required
			continue
		}

		if m.Idpreferences[i] != nil {
			if err := m.Idpreferences[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Idpreferences" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IamAccount) validateIdps(formats strfmt.Registry) error {

	if swag.IsZero(m.Idps) { // not required
		return nil
	}

	for i := 0; i < len(m.Idps); i++ {
		if swag.IsZero(m.Idps[i]) { // not required
			continue
		}

		if m.Idps[i] != nil {
			if err := m.Idps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Idps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IamAccount) validatePermissions(formats strfmt.Registry) error {

	if swag.IsZero(m.Permissions) { // not required
		return nil
	}

	for i := 0; i < len(m.Permissions); i++ {
		if swag.IsZero(m.Permissions[i]) { // not required
			continue
		}

		if m.Permissions[i] != nil {
			if err := m.Permissions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Permissions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IamAccount) validatePrivilegeSets(formats strfmt.Registry) error {

	if swag.IsZero(m.PrivilegeSets) { // not required
		return nil
	}

	for i := 0; i < len(m.PrivilegeSets); i++ {
		if swag.IsZero(m.PrivilegeSets[i]) { // not required
			continue
		}

		if m.PrivilegeSets[i] != nil {
			if err := m.PrivilegeSets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("PrivilegeSets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IamAccount) validatePrivileges(formats strfmt.Registry) error {

	if swag.IsZero(m.Privileges) { // not required
		return nil
	}

	for i := 0; i < len(m.Privileges); i++ {
		if swag.IsZero(m.Privileges[i]) { // not required
			continue
		}

		if m.Privileges[i] != nil {
			if err := m.Privileges[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Privileges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IamAccount) validateResourceLimits(formats strfmt.Registry) error {

	if swag.IsZero(m.ResourceLimits) { // not required
		return nil
	}

	if m.ResourceLimits != nil {
		if err := m.ResourceLimits.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ResourceLimits")
			}
			return err
		}
	}

	return nil
}

func (m *IamAccount) validateRoles(formats strfmt.Registry) error {

	if swag.IsZero(m.Roles) { // not required
		return nil
	}

	for i := 0; i < len(m.Roles); i++ {
		if swag.IsZero(m.Roles[i]) { // not required
			continue
		}

		if m.Roles[i] != nil {
			if err := m.Roles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Roles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IamAccount) validateSessionLimits(formats strfmt.Registry) error {

	if swag.IsZero(m.SessionLimits) { // not required
		return nil
	}

	if m.SessionLimits != nil {
		if err := m.SessionLimits.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SessionLimits")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IamAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IamAccount) UnmarshalBinary(b []byte) error {
	var res IamAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}