// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VirtualizationGuestInfo Virtualization:Guest Info
//
// Captures the common details of the guest personality that runs in a VM.
//
// swagger:model virtualizationGuestInfo
type VirtualizationGuestInfo struct {
	MoBaseComplexType

	VirtualizationGuestInfoAO1P1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *VirtualizationGuestInfo) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseComplexType
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseComplexType = aO0

	// AO1
	var aO1 VirtualizationGuestInfoAO1P1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.VirtualizationGuestInfoAO1P1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m VirtualizationGuestInfo) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseComplexType)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.VirtualizationGuestInfoAO1P1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this virtualization guest info
func (m *VirtualizationGuestInfo) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseComplexType
	if err := m.MoBaseComplexType.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with VirtualizationGuestInfoAO1P1
	if err := m.VirtualizationGuestInfoAO1P1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *VirtualizationGuestInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualizationGuestInfo) UnmarshalBinary(b []byte) error {
	var res VirtualizationGuestInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VirtualizationGuestInfoAO1P1 virtualization guest info a o1 p1
//
// swagger:model VirtualizationGuestInfoAO1P1
type VirtualizationGuestInfoAO1P1 struct {

	// Name provided to the host OS (example, ubuntu6410, test-gateway, etc.).
	Hostname string `json:"Hostname,omitempty"`

	// Primary IP address of the guest os.
	IPAddress string `json:"IpAddress,omitempty"`

	// The name of the guest running on this VM. This may not be the same as the hostname.
	Name string `json:"Name,omitempty"`

	// The name of the guest OS running on this VM (Cent OS 4/5/6/7).
	OperatingSystem string `json:"OperatingSystem,omitempty"`

	// virtualization guest info a o1 p1
	VirtualizationGuestInfoAO1P1 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *VirtualizationGuestInfoAO1P1) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// Name provided to the host OS (example, ubuntu6410, test-gateway, etc.).
		Hostname string `json:"Hostname,omitempty"`

		// Primary IP address of the guest os.
		IPAddress string `json:"IpAddress,omitempty"`

		// The name of the guest running on this VM. This may not be the same as the hostname.
		Name string `json:"Name,omitempty"`

		// The name of the guest OS running on this VM (Cent OS 4/5/6/7).
		OperatingSystem string `json:"OperatingSystem,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv VirtualizationGuestInfoAO1P1

	rcv.Hostname = stage1.Hostname
	rcv.IPAddress = stage1.IPAddress
	rcv.Name = stage1.Name
	rcv.OperatingSystem = stage1.OperatingSystem
	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "Hostname")
	delete(stage2, "IpAddress")
	delete(stage2, "Name")
	delete(stage2, "OperatingSystem")
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
		m.VirtualizationGuestInfoAO1P1 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m VirtualizationGuestInfoAO1P1) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// Name provided to the host OS (example, ubuntu6410, test-gateway, etc.).
		Hostname string `json:"Hostname,omitempty"`

		// Primary IP address of the guest os.
		IPAddress string `json:"IpAddress,omitempty"`

		// The name of the guest running on this VM. This may not be the same as the hostname.
		Name string `json:"Name,omitempty"`

		// The name of the guest OS running on this VM (Cent OS 4/5/6/7).
		OperatingSystem string `json:"OperatingSystem,omitempty"`
	}

	stage1.Hostname = m.Hostname
	stage1.IPAddress = m.IPAddress
	stage1.Name = m.Name
	stage1.OperatingSystem = m.OperatingSystem

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.VirtualizationGuestInfoAO1P1) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.VirtualizationGuestInfoAO1P1)
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

// Validate validates this virtualization guest info a o1 p1
func (m *VirtualizationGuestInfoAO1P1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VirtualizationGuestInfoAO1P1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualizationGuestInfoAO1P1) UnmarshalBinary(b []byte) error {
	var res VirtualizationGuestInfoAO1P1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}