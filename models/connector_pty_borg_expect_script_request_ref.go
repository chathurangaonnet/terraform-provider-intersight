// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ConnectorPtyBorgExpectScriptRequestRef connector pty borg expect script request ref
// swagger:model connectorPtyBorgExpectScriptRequestRef
type ConnectorPtyBorgExpectScriptRequestRef struct {
	MoMoRef
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ConnectorPtyBorgExpectScriptRequestRef) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoMoRef
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoMoRef = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ConnectorPtyBorgExpectScriptRequestRef) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.MoMoRef)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this connector pty borg expect script request ref
func (m *ConnectorPtyBorgExpectScriptRequestRef) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoMoRef
	if err := m.MoMoRef.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ConnectorPtyBorgExpectScriptRequestRef) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConnectorPtyBorgExpectScriptRequestRef) UnmarshalBinary(b []byte) error {
	var res ConnectorPtyBorgExpectScriptRequestRef
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
