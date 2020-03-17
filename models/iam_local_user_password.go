// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IamLocalUserPassword Iam:Local User Password
//
// LocalUserPassword type is used for changing local user's password. Caller must send old password in Password field and new password in newPassword field. Intersight will verify the old password and sets the new password if everything is OK. This API must not be used for resetting user's password.
//
// swagger:model iamLocalUserPassword
type IamLocalUserPassword struct {
	MoBaseMo

	// User-entered passsord to be compared to password for change password function.
	CurrentPassword string `json:"CurrentPassword,omitempty"`

	// New password that the user's password should be changed to.
	NewPassword string `json:"NewPassword,omitempty"`

	// User's current valid passsord.
	// Format: byte
	Password strfmt.Base64 `json:"Password,omitempty"`

	// The backreference of the user password to it's user object.
	// Read Only: true
	User *IamUserRef `json:"User,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *IamLocalUserPassword) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		CurrentPassword string `json:"CurrentPassword,omitempty"`

		NewPassword string `json:"NewPassword,omitempty"`

		Password strfmt.Base64 `json:"Password,omitempty"`

		User *IamUserRef `json:"User,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.CurrentPassword = dataAO1.CurrentPassword

	m.NewPassword = dataAO1.NewPassword

	m.Password = dataAO1.Password

	m.User = dataAO1.User

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m IamLocalUserPassword) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		CurrentPassword string `json:"CurrentPassword,omitempty"`

		NewPassword string `json:"NewPassword,omitempty"`

		Password strfmt.Base64 `json:"Password,omitempty"`

		User *IamUserRef `json:"User,omitempty"`
	}

	dataAO1.CurrentPassword = m.CurrentPassword

	dataAO1.NewPassword = m.NewPassword

	dataAO1.Password = m.Password

	dataAO1.User = m.User

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this iam local user password
func (m *IamLocalUserPassword) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IamLocalUserPassword) validateUser(formats strfmt.Registry) error {

	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("User")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IamLocalUserPassword) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IamLocalUserPassword) UnmarshalBinary(b []byte) error {
	var res IamLocalUserPassword
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
