// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SoftwarerepositoryCifsServer CIFS
//
// An external file repository accessible through the CIFS protocol.
//
// swagger:model softwarerepositoryCifsServer
type SoftwarerepositoryCifsServer struct {
	SoftwarerepositoryFileServer

	// is password set
	IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`

	// For CIFS, leave the field blank or enter one or more comma seperated options from the following. For Example, " " , " soft " , " soft , nounix " .
	// * soft.
	// * nounix.
	// * noserviceino.
	// * guest.
	// * USERNAME=VALUE.
	// * PASSWORD=VALUE.
	// * sec=VALUE (VALUE could be None, Ntlm, Ntlmi, Ntlmssp, Ntlmsspi, Ntlmv2, Ntlmv2i).
	//
	//
	MountOption string `json:"MountOption,omitempty"`

	// Password configured on the file server.
	//
	Password string `json:"Password,omitempty"`

	// Filename of the image in the CIFS server. For example:ucs-c220m5-huu-3.1.2c.iso.
	//
	RemoteFile string `json:"RemoteFile,omitempty"`

	// Hostname or IP Address of the CIFS server.
	//
	RemoteIP string `json:"RemoteIp,omitempty"`

	// Remote directory where the image is present. For example:/share/subfolder.
	//
	RemoteShare string `json:"RemoteShare,omitempty"`

	// Username configured on the CIFS server.
	//
	Username string `json:"Username,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *SoftwarerepositoryCifsServer) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 SoftwarerepositoryFileServer
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.SoftwarerepositoryFileServer = aO0

	// AO1
	var dataAO1 struct {
		IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`

		MountOption string `json:"MountOption,omitempty"`

		Password string `json:"Password,omitempty"`

		RemoteFile string `json:"RemoteFile,omitempty"`

		RemoteIP string `json:"RemoteIp,omitempty"`

		RemoteShare string `json:"RemoteShare,omitempty"`

		Username string `json:"Username,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.IsPasswordSet = dataAO1.IsPasswordSet

	m.MountOption = dataAO1.MountOption

	m.Password = dataAO1.Password

	m.RemoteFile = dataAO1.RemoteFile

	m.RemoteIP = dataAO1.RemoteIP

	m.RemoteShare = dataAO1.RemoteShare

	m.Username = dataAO1.Username

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m SoftwarerepositoryCifsServer) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.SoftwarerepositoryFileServer)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`

		MountOption string `json:"MountOption,omitempty"`

		Password string `json:"Password,omitempty"`

		RemoteFile string `json:"RemoteFile,omitempty"`

		RemoteIP string `json:"RemoteIp,omitempty"`

		RemoteShare string `json:"RemoteShare,omitempty"`

		Username string `json:"Username,omitempty"`
	}

	dataAO1.IsPasswordSet = m.IsPasswordSet

	dataAO1.MountOption = m.MountOption

	dataAO1.Password = m.Password

	dataAO1.RemoteFile = m.RemoteFile

	dataAO1.RemoteIP = m.RemoteIP

	dataAO1.RemoteShare = m.RemoteShare

	dataAO1.Username = m.Username

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this softwarerepository cifs server
func (m *SoftwarerepositoryCifsServer) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with SoftwarerepositoryFileServer
	if err := m.SoftwarerepositoryFileServer.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SoftwarerepositoryCifsServer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SoftwarerepositoryCifsServer) UnmarshalBinary(b []byte) error {
	var res SoftwarerepositoryCifsServer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
