// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AssetHTTPConnection Asset:Http Connection
//
// HttpConnection provides the necessary details for Intersight to connect to and authenticate with a managed target over an HTTP connection.
//
// swagger:model assetHttpConnection
type AssetHTTPConnection struct {
	AssetConnection

	// The credential to be used to authenticate with the managed target.
	Credential *AssetCredential `json:"Credential,omitempty"`

	// Indicates whether a connection to the target should be established using HTTPS.
	IsSecure *bool `json:"IsSecure,omitempty"`

	// The DNS hostname or IP Address, either IPv4 or IPv6, to be used to connect to the managed target.
	ManagementAddress string `json:"ManagementAddress,omitempty"`

	// The port number to be used to to connect to the managed target. Values 1-65535 indicate a port number to be used. A value of 0 is not a valid port number and instead indicates that the default management port, as defined by the documentation of the managed target, should be used to establish a connection.
	Port int64 `json:"Port,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *AssetHTTPConnection) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 AssetConnection
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.AssetConnection = aO0

	// AO1
	var dataAO1 struct {
		Credential *AssetCredential `json:"Credential,omitempty"`

		IsSecure *bool `json:"IsSecure,omitempty"`

		ManagementAddress string `json:"ManagementAddress,omitempty"`

		Port int64 `json:"Port,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Credential = dataAO1.Credential

	m.IsSecure = dataAO1.IsSecure

	m.ManagementAddress = dataAO1.ManagementAddress

	m.Port = dataAO1.Port

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m AssetHTTPConnection) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.AssetConnection)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Credential *AssetCredential `json:"Credential,omitempty"`

		IsSecure *bool `json:"IsSecure,omitempty"`

		ManagementAddress string `json:"ManagementAddress,omitempty"`

		Port int64 `json:"Port,omitempty"`
	}

	dataAO1.Credential = m.Credential

	dataAO1.IsSecure = m.IsSecure

	dataAO1.ManagementAddress = m.ManagementAddress

	dataAO1.Port = m.Port

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this asset Http connection
func (m *AssetHTTPConnection) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with AssetConnection
	if err := m.AssetConnection.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredential(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AssetHTTPConnection) validateCredential(formats strfmt.Registry) error {

	if swag.IsZero(m.Credential) { // not required
		return nil
	}

	if m.Credential != nil {
		if err := m.Credential.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Credential")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AssetHTTPConnection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssetHTTPConnection) UnmarshalBinary(b []byte) error {
	var res AssetHTTPConnection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}