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

// FirmwareHTTPServer Firmware:Http Server
//
// An external HTTP file server.
//
// swagger:model firmwareHttpServer
type FirmwareHTTPServer struct {
	MoBaseComplexType

	FirmwareHTTPServerAO1P1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *FirmwareHTTPServer) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseComplexType
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseComplexType = aO0

	// AO1
	var aO1 FirmwareHTTPServerAO1P1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.FirmwareHTTPServerAO1P1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m FirmwareHTTPServer) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseComplexType)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.FirmwareHTTPServerAO1P1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this firmware Http server
func (m *FirmwareHTTPServer) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseComplexType
	if err := m.MoBaseComplexType.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with FirmwareHTTPServerAO1P1
	if err := m.FirmwareHTTPServerAO1P1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *FirmwareHTTPServer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FirmwareHTTPServer) UnmarshalBinary(b []byte) error {
	var res FirmwareHTTPServer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FirmwareHTTPServerAO1P1 firmware HTTP server a o1 p1
// swagger:model FirmwareHTTPServerAO1P1
type FirmwareHTTPServerAO1P1 struct {

	// HTTP/HTTPS link to the image. Accepted formats HTTP[s]://server-hostname/share/image or HTTP[s]://serverip/share/image.
	//
	LocationLink string `json:"LocationLink,omitempty"`

	// Mount option as configured on the HTTP server. Empty if nothing is configured.
	//
	MountOptions string `json:"MountOptions,omitempty"`

	// firmware HTTP server a o1 p1
	FirmwareHTTPServerAO1P1 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *FirmwareHTTPServerAO1P1) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// HTTP/HTTPS link to the image. Accepted formats HTTP[s]://server-hostname/share/image or HTTP[s]://serverip/share/image.
		//
		LocationLink string `json:"LocationLink,omitempty"`

		// Mount option as configured on the HTTP server. Empty if nothing is configured.
		//
		MountOptions string `json:"MountOptions,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv FirmwareHTTPServerAO1P1

	rcv.LocationLink = stage1.LocationLink

	rcv.MountOptions = stage1.MountOptions

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "LocationLink")

	delete(stage2, "MountOptions")

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
		m.FirmwareHTTPServerAO1P1 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m FirmwareHTTPServerAO1P1) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// HTTP/HTTPS link to the image. Accepted formats HTTP[s]://server-hostname/share/image or HTTP[s]://serverip/share/image.
		//
		LocationLink string `json:"LocationLink,omitempty"`

		// Mount option as configured on the HTTP server. Empty if nothing is configured.
		//
		MountOptions string `json:"MountOptions,omitempty"`
	}

	stage1.LocationLink = m.LocationLink

	stage1.MountOptions = m.MountOptions

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.FirmwareHTTPServerAO1P1) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.FirmwareHTTPServerAO1P1)
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

// Validate validates this firmware HTTP server a o1 p1
func (m *FirmwareHTTPServerAO1P1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FirmwareHTTPServerAO1P1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FirmwareHTTPServerAO1P1) UnmarshalBinary(b []byte) error {
	var res FirmwareHTTPServerAO1P1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
