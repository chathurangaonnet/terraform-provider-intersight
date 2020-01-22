// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// WorkflowFailureEndTask Workflow:Failure End Task
//
// A FailureEndTask denotes the failed completion of a workflow.
//
// swagger:model workflowFailureEndTask
type WorkflowFailureEndTask struct {
	WorkflowEndTask

	WorkflowFailureEndTaskAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *WorkflowFailureEndTask) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 WorkflowEndTask
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.WorkflowEndTask = aO0

	// AO1
	var aO1 WorkflowFailureEndTaskAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.WorkflowFailureEndTaskAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m WorkflowFailureEndTask) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.WorkflowEndTask)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.WorkflowFailureEndTaskAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this workflow failure end task
func (m *WorkflowFailureEndTask) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with WorkflowEndTask
	if err := m.WorkflowEndTask.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with WorkflowFailureEndTaskAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowFailureEndTask) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowFailureEndTask) UnmarshalBinary(b []byte) error {
	var res WorkflowFailureEndTask
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// WorkflowFailureEndTaskAllOf1 workflow failure end task all of1
// swagger:model WorkflowFailureEndTaskAllOf1
type WorkflowFailureEndTaskAllOf1 interface{}
