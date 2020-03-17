// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IamLdapDNSParameters Iam:Ldap Dns Parameters
//
// DNS settings used to access LDAP Providers.
//
// swagger:model iamLdapDnsParameters
type IamLdapDNSParameters struct {
	MoBaseComplexType

	IamLdapDNSParametersAO1P1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *IamLdapDNSParameters) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseComplexType
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseComplexType = aO0

	// AO1
	var aO1 IamLdapDNSParametersAO1P1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.IamLdapDNSParametersAO1P1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m IamLdapDNSParameters) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseComplexType)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.IamLdapDNSParametersAO1P1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this iam ldap Dns parameters
func (m *IamLdapDNSParameters) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseComplexType
	if err := m.MoBaseComplexType.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with IamLdapDNSParametersAO1P1
	if err := m.IamLdapDNSParametersAO1P1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *IamLdapDNSParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IamLdapDNSParameters) UnmarshalBinary(b []byte) error {
	var res IamLdapDNSParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IamLdapDNSParametersAO1P1 iam ldap DNS parameters a o1 p1
//
// swagger:model IamLdapDNSParametersAO1P1
type IamLdapDNSParametersAO1P1 struct {

	// Domain name that acts as a source for a DNS query.
	SearchDomain string `json:"SearchDomain,omitempty"`

	// Forest name that acts as a source for a DNS query.
	SearchForest string `json:"SearchForest,omitempty"`

	// Source of the domain name used for the DNS SRV request.
	// Enum: [Extracted Configured ConfiguredExtracted]
	Source *string `json:"Source,omitempty"`

	// iam ldap DNS parameters a o1 p1
	IamLdapDNSParametersAO1P1 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *IamLdapDNSParametersAO1P1) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// Domain name that acts as a source for a DNS query.
		SearchDomain string `json:"SearchDomain,omitempty"`

		// Forest name that acts as a source for a DNS query.
		SearchForest string `json:"SearchForest,omitempty"`

		// Source of the domain name used for the DNS SRV request.
		// Enum: [Extracted Configured ConfiguredExtracted]
		Source *string `json:"Source,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv IamLdapDNSParametersAO1P1

	rcv.SearchDomain = stage1.SearchDomain
	rcv.SearchForest = stage1.SearchForest
	rcv.Source = stage1.Source
	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "SearchDomain")
	delete(stage2, "SearchForest")
	delete(stage2, "Source")
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
		m.IamLdapDNSParametersAO1P1 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m IamLdapDNSParametersAO1P1) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// Domain name that acts as a source for a DNS query.
		SearchDomain string `json:"SearchDomain,omitempty"`

		// Forest name that acts as a source for a DNS query.
		SearchForest string `json:"SearchForest,omitempty"`

		// Source of the domain name used for the DNS SRV request.
		// Enum: [Extracted Configured ConfiguredExtracted]
		Source *string `json:"Source,omitempty"`
	}

	stage1.SearchDomain = m.SearchDomain
	stage1.SearchForest = m.SearchForest
	stage1.Source = m.Source

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.IamLdapDNSParametersAO1P1) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.IamLdapDNSParametersAO1P1)
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

// Validate validates this iam ldap DNS parameters a o1 p1
func (m *IamLdapDNSParametersAO1P1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var iamLdapDnsParametersAO1P1TypeSourcePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Extracted","Configured","ConfiguredExtracted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		iamLdapDnsParametersAO1P1TypeSourcePropEnum = append(iamLdapDnsParametersAO1P1TypeSourcePropEnum, v)
	}
}

const (

	// IamLdapDNSParametersAO1P1SourceExtracted captures enum value "Extracted"
	IamLdapDNSParametersAO1P1SourceExtracted string = "Extracted"

	// IamLdapDNSParametersAO1P1SourceConfigured captures enum value "Configured"
	IamLdapDNSParametersAO1P1SourceConfigured string = "Configured"

	// IamLdapDNSParametersAO1P1SourceConfiguredExtracted captures enum value "ConfiguredExtracted"
	IamLdapDNSParametersAO1P1SourceConfiguredExtracted string = "ConfiguredExtracted"
)

// prop value enum
func (m *IamLdapDNSParametersAO1P1) validateSourceEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, iamLdapDnsParametersAO1P1TypeSourcePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IamLdapDNSParametersAO1P1) validateSource(formats strfmt.Registry) error {

	if swag.IsZero(m.Source) { // not required
		return nil
	}

	// value enum
	if err := m.validateSourceEnum("Source", "body", *m.Source); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IamLdapDNSParametersAO1P1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IamLdapDNSParametersAO1P1) UnmarshalBinary(b []byte) error {
	var res IamLdapDNSParametersAO1P1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
