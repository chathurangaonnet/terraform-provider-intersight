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

// IamPermissionReference Iam:Permission Reference
//
// Users can log in through the base URL (https://intersight.com) or account-specific URLs. When the Intersight user logs in through the base URL, Intersight identifies the accounts and permissions within each account which the user has access to. In case multiple permissions are identified, the user and session objects are created in the onboarding-user account, and the session object is updated with account and permission information. Intersight GUI uses this information to show available accounts and permissions for the user to select. PermissionReference type is used to store permission information of an account.
//
// swagger:model iamPermissionReference
type IamPermissionReference struct {
	MoBaseComplexType

	IamPermissionReferenceAO1P1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *IamPermissionReference) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseComplexType
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseComplexType = aO0

	// AO1
	var aO1 IamPermissionReferenceAO1P1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.IamPermissionReferenceAO1P1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m IamPermissionReference) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseComplexType)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.IamPermissionReferenceAO1P1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this iam permission reference
func (m *IamPermissionReference) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseComplexType
	if err := m.MoBaseComplexType.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with IamPermissionReferenceAO1P1
	if err := m.IamPermissionReferenceAO1P1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *IamPermissionReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IamPermissionReference) UnmarshalBinary(b []byte) error {
	var res IamPermissionReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IamPermissionReferenceAO1P1 iam permission reference a o1 p1
// swagger:model IamPermissionReferenceAO1P1
type IamPermissionReferenceAO1P1 struct {

	// MOID of the permission which user has access to.
	//
	// Read Only: true
	PermissionIdentifier string `json:"PermissionIdentifier,omitempty"`

	// Name of the permission which user has access to.
	//
	// Read Only: true
	PermissionName string `json:"PermissionName,omitempty"`

	// iam permission reference a o1 p1
	IamPermissionReferenceAO1P1 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *IamPermissionReferenceAO1P1) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// MOID of the permission which user has access to.
		//
		// Read Only: true
		PermissionIdentifier string `json:"PermissionIdentifier,omitempty"`

		// Name of the permission which user has access to.
		//
		// Read Only: true
		PermissionName string `json:"PermissionName,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv IamPermissionReferenceAO1P1

	rcv.PermissionIdentifier = stage1.PermissionIdentifier

	rcv.PermissionName = stage1.PermissionName

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "PermissionIdentifier")

	delete(stage2, "PermissionName")

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
		m.IamPermissionReferenceAO1P1 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m IamPermissionReferenceAO1P1) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// MOID of the permission which user has access to.
		//
		// Read Only: true
		PermissionIdentifier string `json:"PermissionIdentifier,omitempty"`

		// Name of the permission which user has access to.
		//
		// Read Only: true
		PermissionName string `json:"PermissionName,omitempty"`
	}

	stage1.PermissionIdentifier = m.PermissionIdentifier

	stage1.PermissionName = m.PermissionName

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.IamPermissionReferenceAO1P1) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.IamPermissionReferenceAO1P1)
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

// Validate validates this iam permission reference a o1 p1
func (m *IamPermissionReferenceAO1P1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IamPermissionReferenceAO1P1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IamPermissionReferenceAO1P1) UnmarshalBinary(b []byte) error {
	var res IamPermissionReferenceAO1P1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
