// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IamEndPointRole Iam:End Point Role
//
// The role defined in the end point which can be assigned to a user.
//
// swagger:model iamEndPointRole
type IamEndPointRole struct {
	MoBaseMo

	// A collection of references to the [iam.Account](mo://iam.Account) Managed Object.
	// When this managed object is deleted, the referenced [iam.Account](mo://iam.Account) MO unsets its reference to this deleted MO.
	// Read Only: true
	Account *IamAccountRef `json:"Account,omitempty"`

	// Privileges assigned to this end point role. These privileges are assigned to users using end point roles to perform operations such as GUI/CLI cross launch.
	// Read Only: true
	EndPointPrivileges []*IamEndPointPrivilegeRef `json:"EndPointPrivileges"`

	// The name of the end point role.
	// Read Only: true
	Name string `json:"Name,omitempty"`

	// User specified tags to roles like as ep-admin or ep-readonly.
	// Read Only: true
	RoleType string `json:"RoleType,omitempty"`

	// A collection of references to the [iam.System](mo://iam.System) Managed Object.
	// When this managed object is deleted, the referenced [iam.System](mo://iam.System) MO unsets its reference to this deleted MO.
	// Read Only: true
	System *IamSystemRef `json:"System,omitempty"`

	// The type of the end point like Cisco UCS Fabric Interconnect or Cisco IMC.
	// Read Only: true
	// Enum: [ APIC DCNM UCSFI UCSFIISM IMC IMCM4 IMCM5 UCSIOM HX HXTriton UCSD IntersightAppliance PureStorageFlashArray VmwareVcenter ServiceEngine IMCBlade]
	Type string `json:"Type,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *IamEndPointRole) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		Account *IamAccountRef `json:"Account,omitempty"`

		EndPointPrivileges []*IamEndPointPrivilegeRef `json:"EndPointPrivileges"`

		Name string `json:"Name,omitempty"`

		RoleType string `json:"RoleType,omitempty"`

		System *IamSystemRef `json:"System,omitempty"`

		Type string `json:"Type,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Account = dataAO1.Account

	m.EndPointPrivileges = dataAO1.EndPointPrivileges

	m.Name = dataAO1.Name

	m.RoleType = dataAO1.RoleType

	m.System = dataAO1.System

	m.Type = dataAO1.Type

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m IamEndPointRole) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Account *IamAccountRef `json:"Account,omitempty"`

		EndPointPrivileges []*IamEndPointPrivilegeRef `json:"EndPointPrivileges"`

		Name string `json:"Name,omitempty"`

		RoleType string `json:"RoleType,omitempty"`

		System *IamSystemRef `json:"System,omitempty"`

		Type string `json:"Type,omitempty"`
	}

	dataAO1.Account = m.Account

	dataAO1.EndPointPrivileges = m.EndPointPrivileges

	dataAO1.Name = m.Name

	dataAO1.RoleType = m.RoleType

	dataAO1.System = m.System

	dataAO1.Type = m.Type

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this iam end point role
func (m *IamEndPointRole) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndPointPrivileges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSystem(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IamEndPointRole) validateAccount(formats strfmt.Registry) error {

	if swag.IsZero(m.Account) { // not required
		return nil
	}

	if m.Account != nil {
		if err := m.Account.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Account")
			}
			return err
		}
	}

	return nil
}

func (m *IamEndPointRole) validateEndPointPrivileges(formats strfmt.Registry) error {

	if swag.IsZero(m.EndPointPrivileges) { // not required
		return nil
	}

	for i := 0; i < len(m.EndPointPrivileges); i++ {
		if swag.IsZero(m.EndPointPrivileges[i]) { // not required
			continue
		}

		if m.EndPointPrivileges[i] != nil {
			if err := m.EndPointPrivileges[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("EndPointPrivileges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IamEndPointRole) validateSystem(formats strfmt.Registry) error {

	if swag.IsZero(m.System) { // not required
		return nil
	}

	if m.System != nil {
		if err := m.System.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("System")
			}
			return err
		}
	}

	return nil
}

var iamEndPointRoleTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["","APIC","DCNM","UCSFI","UCSFIISM","IMC","IMCM4","IMCM5","UCSIOM","HX","HXTriton","UCSD","IntersightAppliance","PureStorageFlashArray","VmwareVcenter","ServiceEngine","IMCBlade"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		iamEndPointRoleTypeTypePropEnum = append(iamEndPointRoleTypeTypePropEnum, v)
	}
}

// property enum
func (m *IamEndPointRole) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, iamEndPointRoleTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IamEndPointRole) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("Type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IamEndPointRole) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IamEndPointRole) UnmarshalBinary(b []byte) error {
	var res IamEndPointRole
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
