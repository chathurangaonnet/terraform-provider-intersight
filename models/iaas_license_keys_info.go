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

// IaasLicenseKeysInfo Iaas:License Keys Info
//
// License list with the utilization info for UCSD.
//
// swagger:model iaasLicenseKeysInfo
type IaasLicenseKeysInfo struct {
	IaasLicenseKeysInfoAO0P0
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *IaasLicenseKeysInfo) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 IaasLicenseKeysInfoAO0P0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.IaasLicenseKeysInfoAO0P0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m IaasLicenseKeysInfo) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.IaasLicenseKeysInfoAO0P0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this iaas license keys info
func (m *IaasLicenseKeysInfo) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with IaasLicenseKeysInfoAO0P0
	if err := m.IaasLicenseKeysInfoAO0P0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *IaasLicenseKeysInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IaasLicenseKeysInfo) UnmarshalBinary(b []byte) error {
	var res IaasLicenseKeysInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IaasLicenseKeysInfoAO0P0 iaas license keys info a o0 p0
// swagger:model IaasLicenseKeysInfoAO0P0
type IaasLicenseKeysInfoAO0P0 struct {

	// Number of licenses available for the UCSD PID (Product ID).
	//
	// Read Only: true
	Count int64 `json:"Count,omitempty"`

	// Expiration date for the license.
	//
	// Read Only: true
	ExpirationDate string `json:"ExpirationDate,omitempty"`

	// UCS Director Unique license ID.
	//
	// Read Only: true
	LicenseID string `json:"LicenseId,omitempty"`

	// The concrete type of this complex type.
	//
	// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
	// ObjectType is optional.
	// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
	// are heterogeneous, i.e. the array can contain nested documents of different types.
	//
	//
	ObjectType string `json:"ObjectType,omitempty"`

	// PID (Product ID) for UCSD License.
	//
	// Read Only: true
	Pid string `json:"Pid,omitempty"`

	// iaas license keys info a o0 p0
	IaasLicenseKeysInfoAO0P0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *IaasLicenseKeysInfoAO0P0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// Number of licenses available for the UCSD PID (Product ID).
		//
		// Read Only: true
		Count int64 `json:"Count,omitempty"`

		// Expiration date for the license.
		//
		// Read Only: true
		ExpirationDate string `json:"ExpirationDate,omitempty"`

		// UCS Director Unique license ID.
		//
		// Read Only: true
		LicenseID string `json:"LicenseId,omitempty"`

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// PID (Product ID) for UCSD License.
		//
		// Read Only: true
		Pid string `json:"Pid,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv IaasLicenseKeysInfoAO0P0

	rcv.Count = stage1.Count

	rcv.ExpirationDate = stage1.ExpirationDate

	rcv.LicenseID = stage1.LicenseID

	rcv.ObjectType = stage1.ObjectType

	rcv.Pid = stage1.Pid

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "Count")

	delete(stage2, "ExpirationDate")

	delete(stage2, "LicenseId")

	delete(stage2, "ObjectType")

	delete(stage2, "Pid")

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
		m.IaasLicenseKeysInfoAO0P0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m IaasLicenseKeysInfoAO0P0) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// Number of licenses available for the UCSD PID (Product ID).
		//
		// Read Only: true
		Count int64 `json:"Count,omitempty"`

		// Expiration date for the license.
		//
		// Read Only: true
		ExpirationDate string `json:"ExpirationDate,omitempty"`

		// UCS Director Unique license ID.
		//
		// Read Only: true
		LicenseID string `json:"LicenseId,omitempty"`

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// PID (Product ID) for UCSD License.
		//
		// Read Only: true
		Pid string `json:"Pid,omitempty"`
	}

	stage1.Count = m.Count

	stage1.ExpirationDate = m.ExpirationDate

	stage1.LicenseID = m.LicenseID

	stage1.ObjectType = m.ObjectType

	stage1.Pid = m.Pid

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.IaasLicenseKeysInfoAO0P0) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.IaasLicenseKeysInfoAO0P0)
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

// Validate validates this iaas license keys info a o0 p0
func (m *IaasLicenseKeysInfoAO0P0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IaasLicenseKeysInfoAO0P0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IaasLicenseKeysInfoAO0P0) UnmarshalBinary(b []byte) error {
	var res IaasLicenseKeysInfoAO0P0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}