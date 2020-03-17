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

// PolicyConfigContext Policy:Config Context
//
// Configuration related state and results.
//
// swagger:model policyConfigContext
type PolicyConfigContext struct {
	MoBaseComplexType

	PolicyConfigContextAO1P1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PolicyConfigContext) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseComplexType
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseComplexType = aO0

	// AO1
	var aO1 PolicyConfigContextAO1P1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.PolicyConfigContextAO1P1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PolicyConfigContext) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseComplexType)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.PolicyConfigContextAO1P1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this policy config context
func (m *PolicyConfigContext) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseComplexType
	if err := m.MoBaseComplexType.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with PolicyConfigContextAO1P1
	if err := m.PolicyConfigContextAO1P1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PolicyConfigContext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicyConfigContext) UnmarshalBinary(b []byte) error {
	var res PolicyConfigContext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PolicyConfigContextAO1P1 policy config context a o1 p1
//
// swagger:model PolicyConfigContextAO1P1
type PolicyConfigContextAO1P1 struct {

	// Indicates a profile's configuration deploying state. Values -- Assigned, Not-assigned, Associated, Pending-changes, Validating, Configuring, Failed.
	// Read Only: true
	ConfigState string `json:"ConfigState,omitempty"`

	// System action to trigger the appropriate workflow. Values -- No_op, ConfigChange, Deploy, Unbind.
	ControlAction string `json:"ControlAction,omitempty"`

	// Indicates a profile's error state. Values -- Validation-error (Static validation error), Pre-config-error (Runtime validation error), Config-error (Runtime configuration error).
	ErrorState string `json:"ErrorState,omitempty"`

	// Combined state (configState, and operational state of the associated physical resource) to indicate the current state of the profile. Values -- n/a, Power-off, Pending-changes, Configuring, Ok, Failed.
	// Read Only: true
	OperState string `json:"OperState,omitempty"`

	// policy config context a o1 p1
	PolicyConfigContextAO1P1 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *PolicyConfigContextAO1P1) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// Indicates a profile's configuration deploying state. Values -- Assigned, Not-assigned, Associated, Pending-changes, Validating, Configuring, Failed.
		// Read Only: true
		ConfigState string `json:"ConfigState,omitempty"`

		// System action to trigger the appropriate workflow. Values -- No_op, ConfigChange, Deploy, Unbind.
		ControlAction string `json:"ControlAction,omitempty"`

		// Indicates a profile's error state. Values -- Validation-error (Static validation error), Pre-config-error (Runtime validation error), Config-error (Runtime configuration error).
		ErrorState string `json:"ErrorState,omitempty"`

		// Combined state (configState, and operational state of the associated physical resource) to indicate the current state of the profile. Values -- n/a, Power-off, Pending-changes, Configuring, Ok, Failed.
		// Read Only: true
		OperState string `json:"OperState,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv PolicyConfigContextAO1P1

	rcv.ConfigState = stage1.ConfigState
	rcv.ControlAction = stage1.ControlAction
	rcv.ErrorState = stage1.ErrorState
	rcv.OperState = stage1.OperState
	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "ConfigState")
	delete(stage2, "ControlAction")
	delete(stage2, "ErrorState")
	delete(stage2, "OperState")
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
		m.PolicyConfigContextAO1P1 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m PolicyConfigContextAO1P1) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// Indicates a profile's configuration deploying state. Values -- Assigned, Not-assigned, Associated, Pending-changes, Validating, Configuring, Failed.
		// Read Only: true
		ConfigState string `json:"ConfigState,omitempty"`

		// System action to trigger the appropriate workflow. Values -- No_op, ConfigChange, Deploy, Unbind.
		ControlAction string `json:"ControlAction,omitempty"`

		// Indicates a profile's error state. Values -- Validation-error (Static validation error), Pre-config-error (Runtime validation error), Config-error (Runtime configuration error).
		ErrorState string `json:"ErrorState,omitempty"`

		// Combined state (configState, and operational state of the associated physical resource) to indicate the current state of the profile. Values -- n/a, Power-off, Pending-changes, Configuring, Ok, Failed.
		// Read Only: true
		OperState string `json:"OperState,omitempty"`
	}

	stage1.ConfigState = m.ConfigState
	stage1.ControlAction = m.ControlAction
	stage1.ErrorState = m.ErrorState
	stage1.OperState = m.OperState

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.PolicyConfigContextAO1P1) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.PolicyConfigContextAO1P1)
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

// Validate validates this policy config context a o1 p1
func (m *PolicyConfigContextAO1P1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PolicyConfigContextAO1P1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicyConfigContextAO1P1) UnmarshalBinary(b []byte) error {
	var res PolicyConfigContextAO1P1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
