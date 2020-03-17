// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IamSessionLimits Iam:Session Limits
//
// The session related configuration limits.
//
// swagger:model iamSessionLimits
type IamSessionLimits struct {
	MoBaseMo

	// A collection of references to the [iam.Account](mo://iam.Account) Managed Object.
	// When this managed object is deleted, the referenced [iam.Account](mo://iam.Account) MO unsets its reference to this deleted MO.
	// Read Only: true
	Account *IamAccountRef `json:"Account,omitempty"`

	// The idle timeout interval for the web session in seconds. The default value is 1800 seconds. When a session is not refreshed for this duration, backend will mark the session as idle and remove the session.
	// Read Only: true
	IdleTimeOut int64 `json:"IdleTimeOut,omitempty"`

	// The maximum number of sessions allowed in an account. The default value is 128.
	// Read Only: true
	MaximumLimit int64 `json:"MaximumLimit,omitempty"`

	// The maximum number of sessions allowed per user. Default value is 32.
	// Read Only: true
	PerUserLimit int64 `json:"PerUserLimit,omitempty"`

	// The session expiry duration in seconds. The default value is 57600 seconds or 16 hours.
	// Read Only: true
	SessionTimeOut int64 `json:"SessionTimeOut,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *IamSessionLimits) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		Account *IamAccountRef `json:"Account,omitempty"`

		IdleTimeOut int64 `json:"IdleTimeOut,omitempty"`

		MaximumLimit int64 `json:"MaximumLimit,omitempty"`

		PerUserLimit int64 `json:"PerUserLimit,omitempty"`

		SessionTimeOut int64 `json:"SessionTimeOut,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Account = dataAO1.Account

	m.IdleTimeOut = dataAO1.IdleTimeOut

	m.MaximumLimit = dataAO1.MaximumLimit

	m.PerUserLimit = dataAO1.PerUserLimit

	m.SessionTimeOut = dataAO1.SessionTimeOut

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m IamSessionLimits) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Account *IamAccountRef `json:"Account,omitempty"`

		IdleTimeOut int64 `json:"IdleTimeOut,omitempty"`

		MaximumLimit int64 `json:"MaximumLimit,omitempty"`

		PerUserLimit int64 `json:"PerUserLimit,omitempty"`

		SessionTimeOut int64 `json:"SessionTimeOut,omitempty"`
	}

	dataAO1.Account = m.Account

	dataAO1.IdleTimeOut = m.IdleTimeOut

	dataAO1.MaximumLimit = m.MaximumLimit

	dataAO1.PerUserLimit = m.PerUserLimit

	dataAO1.SessionTimeOut = m.SessionTimeOut

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this iam session limits
func (m *IamSessionLimits) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IamSessionLimits) validateAccount(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *IamSessionLimits) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IamSessionLimits) UnmarshalBinary(b []byte) error {
	var res IamSessionLimits
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
