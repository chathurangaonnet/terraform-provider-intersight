// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WorkflowProperties Workflow:Properties
//
// Properties for the task definition like the inputs, outputs, timeout and retry policies. Tasks are the building blocks for workflows.
//
// swagger:model workflowProperties
type WorkflowProperties struct {
	MoBaseComplexType

	WorkflowPropertiesAO1P1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *WorkflowProperties) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseComplexType
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseComplexType = aO0

	// AO1
	var aO1 WorkflowPropertiesAO1P1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.WorkflowPropertiesAO1P1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m WorkflowProperties) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseComplexType)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.WorkflowPropertiesAO1P1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this workflow properties
func (m *WorkflowProperties) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseComplexType
	if err := m.MoBaseComplexType.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with WorkflowPropertiesAO1P1
	if err := m.WorkflowPropertiesAO1P1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowProperties) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowProperties) UnmarshalBinary(b []byte) error {
	var res WorkflowProperties
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// WorkflowPropertiesAO1P1 workflow properties a o1 p1
// swagger:model WorkflowPropertiesAO1P1
type WorkflowPropertiesAO1P1 struct {

	// The schema expected for input parameters for this task.
	//
	InputDefinition []*WorkflowBaseDataType `json:"InputDefinition"`

	// The schema expected for output parameters for this task.
	//
	OutputDefinition []*WorkflowBaseDataType `json:"OutputDefinition"`

	// The number of times a task should be tried before marking as failed.
	//
	RetryCount int64 `json:"RetryCount,omitempty"`

	// The delay in seconds after which the the task is re-tried.
	//
	RetryDelay int64 `json:"RetryDelay,omitempty"`

	// The retry policy for the task.
	//
	// Enum: [Fixed]
	RetryPolicy *string `json:"RetryPolicy,omitempty"`

	// The timeout value in seconds after which task will be marked as timed out. Max allowed value is 7 days.
	//
	Timeout int64 `json:"Timeout,omitempty"`

	// The timeout policy for the task.
	//
	// Enum: [Timeout Retry]
	TimeoutPolicy *string `json:"TimeoutPolicy,omitempty"`

	// workflow properties a o1 p1
	WorkflowPropertiesAO1P1 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *WorkflowPropertiesAO1P1) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// The schema expected for input parameters for this task.
		//
		InputDefinition []*WorkflowBaseDataType `json:"InputDefinition"`

		// The schema expected for output parameters for this task.
		//
		OutputDefinition []*WorkflowBaseDataType `json:"OutputDefinition"`

		// The number of times a task should be tried before marking as failed.
		//
		RetryCount int64 `json:"RetryCount,omitempty"`

		// The delay in seconds after which the the task is re-tried.
		//
		RetryDelay int64 `json:"RetryDelay,omitempty"`

		// The retry policy for the task.
		//
		// Enum: [Fixed]
		RetryPolicy *string `json:"RetryPolicy,omitempty"`

		// The timeout value in seconds after which task will be marked as timed out. Max allowed value is 7 days.
		//
		Timeout int64 `json:"Timeout,omitempty"`

		// The timeout policy for the task.
		//
		// Enum: [Timeout Retry]
		TimeoutPolicy *string `json:"TimeoutPolicy,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv WorkflowPropertiesAO1P1

	rcv.InputDefinition = stage1.InputDefinition

	rcv.OutputDefinition = stage1.OutputDefinition

	rcv.RetryCount = stage1.RetryCount

	rcv.RetryDelay = stage1.RetryDelay

	rcv.RetryPolicy = stage1.RetryPolicy

	rcv.Timeout = stage1.Timeout

	rcv.TimeoutPolicy = stage1.TimeoutPolicy

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "InputDefinition")

	delete(stage2, "OutputDefinition")

	delete(stage2, "RetryCount")

	delete(stage2, "RetryDelay")

	delete(stage2, "RetryPolicy")

	delete(stage2, "Timeout")

	delete(stage2, "TimeoutPolicy")

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
		m.WorkflowPropertiesAO1P1 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m WorkflowPropertiesAO1P1) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// The schema expected for input parameters for this task.
		//
		InputDefinition []*WorkflowBaseDataType `json:"InputDefinition"`

		// The schema expected for output parameters for this task.
		//
		OutputDefinition []*WorkflowBaseDataType `json:"OutputDefinition"`

		// The number of times a task should be tried before marking as failed.
		//
		RetryCount int64 `json:"RetryCount,omitempty"`

		// The delay in seconds after which the the task is re-tried.
		//
		RetryDelay int64 `json:"RetryDelay,omitempty"`

		// The retry policy for the task.
		//
		// Enum: [Fixed]
		RetryPolicy *string `json:"RetryPolicy,omitempty"`

		// The timeout value in seconds after which task will be marked as timed out. Max allowed value is 7 days.
		//
		Timeout int64 `json:"Timeout,omitempty"`

		// The timeout policy for the task.
		//
		// Enum: [Timeout Retry]
		TimeoutPolicy *string `json:"TimeoutPolicy,omitempty"`
	}

	stage1.InputDefinition = m.InputDefinition

	stage1.OutputDefinition = m.OutputDefinition

	stage1.RetryCount = m.RetryCount

	stage1.RetryDelay = m.RetryDelay

	stage1.RetryPolicy = m.RetryPolicy

	stage1.Timeout = m.Timeout

	stage1.TimeoutPolicy = m.TimeoutPolicy

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.WorkflowPropertiesAO1P1) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.WorkflowPropertiesAO1P1)
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

// Validate validates this workflow properties a o1 p1
func (m *WorkflowPropertiesAO1P1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInputDefinition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutputDefinition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRetryPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeoutPolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkflowPropertiesAO1P1) validateInputDefinition(formats strfmt.Registry) error {

	if swag.IsZero(m.InputDefinition) { // not required
		return nil
	}

	for i := 0; i < len(m.InputDefinition); i++ {
		if swag.IsZero(m.InputDefinition[i]) { // not required
			continue
		}

		if m.InputDefinition[i] != nil {
			if err := m.InputDefinition[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("InputDefinition" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WorkflowPropertiesAO1P1) validateOutputDefinition(formats strfmt.Registry) error {

	if swag.IsZero(m.OutputDefinition) { // not required
		return nil
	}

	for i := 0; i < len(m.OutputDefinition); i++ {
		if swag.IsZero(m.OutputDefinition[i]) { // not required
			continue
		}

		if m.OutputDefinition[i] != nil {
			if err := m.OutputDefinition[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("OutputDefinition" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var workflowPropertiesAO1P1TypeRetryPolicyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Fixed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		workflowPropertiesAO1P1TypeRetryPolicyPropEnum = append(workflowPropertiesAO1P1TypeRetryPolicyPropEnum, v)
	}
}

const (

	// WorkflowPropertiesAO1P1RetryPolicyFixed captures enum value "Fixed"
	WorkflowPropertiesAO1P1RetryPolicyFixed string = "Fixed"
)

// prop value enum
func (m *WorkflowPropertiesAO1P1) validateRetryPolicyEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, workflowPropertiesAO1P1TypeRetryPolicyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WorkflowPropertiesAO1P1) validateRetryPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.RetryPolicy) { // not required
		return nil
	}

	// value enum
	if err := m.validateRetryPolicyEnum("RetryPolicy", "body", *m.RetryPolicy); err != nil {
		return err
	}

	return nil
}

var workflowPropertiesAO1P1TypeTimeoutPolicyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Timeout","Retry"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		workflowPropertiesAO1P1TypeTimeoutPolicyPropEnum = append(workflowPropertiesAO1P1TypeTimeoutPolicyPropEnum, v)
	}
}

const (

	// WorkflowPropertiesAO1P1TimeoutPolicyTimeout captures enum value "Timeout"
	WorkflowPropertiesAO1P1TimeoutPolicyTimeout string = "Timeout"

	// WorkflowPropertiesAO1P1TimeoutPolicyRetry captures enum value "Retry"
	WorkflowPropertiesAO1P1TimeoutPolicyRetry string = "Retry"
)

// prop value enum
func (m *WorkflowPropertiesAO1P1) validateTimeoutPolicyEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, workflowPropertiesAO1P1TypeTimeoutPolicyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WorkflowPropertiesAO1P1) validateTimeoutPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.TimeoutPolicy) { // not required
		return nil
	}

	// value enum
	if err := m.validateTimeoutPolicyEnum("TimeoutPolicy", "body", *m.TimeoutPolicy); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowPropertiesAO1P1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowPropertiesAO1P1) UnmarshalBinary(b []byte) error {
	var res WorkflowPropertiesAO1P1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
