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

// IamSsoSessionAttributes Iam:Sso Session Attributes
//
// Session attributes required to maintain states of SP and IdP.
//
// swagger:model iamSsoSessionAttributes
type IamSsoSessionAttributes struct {
	MoBaseComplexType

	IamSsoSessionAttributesAO1P1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *IamSsoSessionAttributes) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseComplexType
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseComplexType = aO0

	// AO1
	var aO1 IamSsoSessionAttributesAO1P1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.IamSsoSessionAttributesAO1P1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m IamSsoSessionAttributes) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseComplexType)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.IamSsoSessionAttributesAO1P1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this iam sso session attributes
func (m *IamSsoSessionAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseComplexType
	if err := m.MoBaseComplexType.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with IamSsoSessionAttributesAO1P1
	if err := m.IamSsoSessionAttributesAO1P1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *IamSsoSessionAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IamSsoSessionAttributes) UnmarshalBinary(b []byte) error {
	var res IamSsoSessionAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IamSsoSessionAttributesAO1P1 iam sso session attributes a o1 p1
// swagger:model IamSsoSessionAttributesAO1P1
type IamSsoSessionAttributesAO1P1 struct {

	// SAML SessionNotOnOrAfter attribute sent by IdP in the assertion. IdP uses this to control for how long SP session maybe. SP does not issue SLO if the session is not valid.
	//
	// Read Only: true
	IdpSessionExpiration string `json:"IdpSessionExpiration,omitempty"`

	// SAML SessionIndex attribute sent by IdP in the assertion. This has to be sent back to IdP in LogoutRequest.
	//
	// Read Only: true
	IdpSessionIndex string `json:"IdpSessionIndex,omitempty"`

	// iam sso session attributes a o1 p1
	IamSsoSessionAttributesAO1P1 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *IamSsoSessionAttributesAO1P1) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// SAML SessionNotOnOrAfter attribute sent by IdP in the assertion. IdP uses this to control for how long SP session maybe. SP does not issue SLO if the session is not valid.
		//
		// Read Only: true
		IdpSessionExpiration string `json:"IdpSessionExpiration,omitempty"`

		// SAML SessionIndex attribute sent by IdP in the assertion. This has to be sent back to IdP in LogoutRequest.
		//
		// Read Only: true
		IdpSessionIndex string `json:"IdpSessionIndex,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv IamSsoSessionAttributesAO1P1

	rcv.IdpSessionExpiration = stage1.IdpSessionExpiration

	rcv.IdpSessionIndex = stage1.IdpSessionIndex

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "IdpSessionExpiration")

	delete(stage2, "IdpSessionIndex")

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
		m.IamSsoSessionAttributesAO1P1 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m IamSsoSessionAttributesAO1P1) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// SAML SessionNotOnOrAfter attribute sent by IdP in the assertion. IdP uses this to control for how long SP session maybe. SP does not issue SLO if the session is not valid.
		//
		// Read Only: true
		IdpSessionExpiration string `json:"IdpSessionExpiration,omitempty"`

		// SAML SessionIndex attribute sent by IdP in the assertion. This has to be sent back to IdP in LogoutRequest.
		//
		// Read Only: true
		IdpSessionIndex string `json:"IdpSessionIndex,omitempty"`
	}

	stage1.IdpSessionExpiration = m.IdpSessionExpiration

	stage1.IdpSessionIndex = m.IdpSessionIndex

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.IamSsoSessionAttributesAO1P1) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.IamSsoSessionAttributesAO1P1)
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

// Validate validates this iam sso session attributes a o1 p1
func (m *IamSsoSessionAttributesAO1P1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IamSsoSessionAttributesAO1P1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IamSsoSessionAttributesAO1P1) UnmarshalBinary(b []byte) error {
	var res IamSsoSessionAttributesAO1P1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
