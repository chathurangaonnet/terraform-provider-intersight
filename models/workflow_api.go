// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WorkflowAPI API
//
// Intersight Orchestrator supports generic API workflow tasks that can execute
// an API given the request body and response parser specification.
//
// API type models a single API request within a batch of requests that get
// executed within a single workflow task.
//
// swagger:model workflowApi
type WorkflowAPI struct {
	MoBaseComplexType

	WorkflowAPIAO1P1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *WorkflowAPI) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseComplexType
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseComplexType = aO0

	// AO1
	var aO1 WorkflowAPIAO1P1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.WorkflowAPIAO1P1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m WorkflowAPI) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseComplexType)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.WorkflowAPIAO1P1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this workflow Api
func (m *WorkflowAPI) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseComplexType
	if err := m.MoBaseComplexType.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with WorkflowAPIAO1P1
	if err := m.WorkflowAPIAO1P1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowAPI) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowAPI) UnmarshalBinary(b []byte) error {
	var res WorkflowAPI
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// WorkflowAPIAO1P1 workflow API a o1 p1
// swagger:model WorkflowAPIAO1P1
type WorkflowAPIAO1P1 struct {

	// The optional request body that is sent as part of this API request.
	//
	// The request body can contain a golang template that can be populated with task input
	// parameters and previous API output parameters.
	//
	//
	Body string `json:"Body,omitempty"`

	// Intersight Orchestrator, with the support of response parser specification,
	// can extract the values from API responses and map them to task output parameters.
	// The value extraction is supported for response content types XML and JSON.
	//
	// The type of the content that gets passed as payload and response in this
	// API.
	//
	//
	// Enum: [json xml text]
	ContentType *string `json:"ContentType,omitempty"`

	// A reference name for this API request within the batch API request.
	//
	// This name shall be used to map the API output parameters to subsequent
	// API input parameters within a batch API task.
	//
	//
	Name string `json:"Name,omitempty"`

	// All the possible outcomes of this API are captured here. Outcomes property
	// is a collection property of type workflow.Outcome objects.
	//
	// The outcomes can be mapped to the message to be shown. The outcomes are
	// evaluated in the order they are given. At the end of the outcomes list,
	// an catchall success/fail outcome can be added with condition as 'true'.
	//
	// This is an optional
	// property and if not specified the task will be marked as success.
	//
	//
	Outcomes interface{} `json:"Outcomes,omitempty"`

	// The optional grammar specification for parsing the response to extract the
	// required values.
	//
	// The specification should have extraction specification specified for
	// all the API output parameters.
	//
	//
	ResponseSpec *ContentGrammar `json:"ResponseSpec,omitempty"`

	// The skip expression, if provided, allows the batch API executor to skip the
	// api execution when the given expression evaluates to true.
	//
	// The expression is given as such a golang template that has to be
	// evaluated to a final content true/false. The expression is an optional and in
	// case not provided, the API will always be executed.
	//
	//
	SkipOnCondition string `json:"SkipOnCondition,omitempty"`

	// The duration in seconds by which the API response is expected from the API target.
	//
	// If the end point does not respond for the API request within this timeout
	// duration, the task will be marked as failed.
	//
	//
	Timeout int64 `json:"Timeout,omitempty"`

	// workflow API a o1 p1
	WorkflowAPIAO1P1 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *WorkflowAPIAO1P1) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// The optional request body that is sent as part of this API request.
		//
		// The request body can contain a golang template that can be populated with task input
		// parameters and previous API output parameters.
		//
		//
		Body string `json:"Body,omitempty"`

		// Intersight Orchestrator, with the support of response parser specification,
		// can extract the values from API responses and map them to task output parameters.
		// The value extraction is supported for response content types XML and JSON.
		//
		// The type of the content that gets passed as payload and response in this
		// API.
		//
		//
		// Enum: [json xml text]
		ContentType *string `json:"ContentType,omitempty"`

		// A reference name for this API request within the batch API request.
		//
		// This name shall be used to map the API output parameters to subsequent
		// API input parameters within a batch API task.
		//
		//
		Name string `json:"Name,omitempty"`

		// All the possible outcomes of this API are captured here. Outcomes property
		// is a collection property of type workflow.Outcome objects.
		//
		// The outcomes can be mapped to the message to be shown. The outcomes are
		// evaluated in the order they are given. At the end of the outcomes list,
		// an catchall success/fail outcome can be added with condition as 'true'.
		//
		// This is an optional
		// property and if not specified the task will be marked as success.
		//
		//
		Outcomes interface{} `json:"Outcomes,omitempty"`

		// The optional grammar specification for parsing the response to extract the
		// required values.
		//
		// The specification should have extraction specification specified for
		// all the API output parameters.
		//
		//
		ResponseSpec *ContentGrammar `json:"ResponseSpec,omitempty"`

		// The skip expression, if provided, allows the batch API executor to skip the
		// api execution when the given expression evaluates to true.
		//
		// The expression is given as such a golang template that has to be
		// evaluated to a final content true/false. The expression is an optional and in
		// case not provided, the API will always be executed.
		//
		//
		SkipOnCondition string `json:"SkipOnCondition,omitempty"`

		// The duration in seconds by which the API response is expected from the API target.
		//
		// If the end point does not respond for the API request within this timeout
		// duration, the task will be marked as failed.
		//
		//
		Timeout int64 `json:"Timeout,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv WorkflowAPIAO1P1

	rcv.Body = stage1.Body

	rcv.ContentType = stage1.ContentType

	rcv.Name = stage1.Name

	rcv.Outcomes = stage1.Outcomes

	rcv.ResponseSpec = stage1.ResponseSpec

	rcv.SkipOnCondition = stage1.SkipOnCondition

	rcv.Timeout = stage1.Timeout

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "Body")

	delete(stage2, "ContentType")

	delete(stage2, "Name")

	delete(stage2, "Outcomes")

	delete(stage2, "ResponseSpec")

	delete(stage2, "SkipOnCondition")

	delete(stage2, "Timeout")

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
		m.WorkflowAPIAO1P1 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m WorkflowAPIAO1P1) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// The optional request body that is sent as part of this API request.
		//
		// The request body can contain a golang template that can be populated with task input
		// parameters and previous API output parameters.
		//
		//
		Body string `json:"Body,omitempty"`

		// Intersight Orchestrator, with the support of response parser specification,
		// can extract the values from API responses and map them to task output parameters.
		// The value extraction is supported for response content types XML and JSON.
		//
		// The type of the content that gets passed as payload and response in this
		// API.
		//
		//
		// Enum: [json xml text]
		ContentType *string `json:"ContentType,omitempty"`

		// A reference name for this API request within the batch API request.
		//
		// This name shall be used to map the API output parameters to subsequent
		// API input parameters within a batch API task.
		//
		//
		Name string `json:"Name,omitempty"`

		// All the possible outcomes of this API are captured here. Outcomes property
		// is a collection property of type workflow.Outcome objects.
		//
		// The outcomes can be mapped to the message to be shown. The outcomes are
		// evaluated in the order they are given. At the end of the outcomes list,
		// an catchall success/fail outcome can be added with condition as 'true'.
		//
		// This is an optional
		// property and if not specified the task will be marked as success.
		//
		//
		Outcomes interface{} `json:"Outcomes,omitempty"`

		// The optional grammar specification for parsing the response to extract the
		// required values.
		//
		// The specification should have extraction specification specified for
		// all the API output parameters.
		//
		//
		ResponseSpec *ContentGrammar `json:"ResponseSpec,omitempty"`

		// The skip expression, if provided, allows the batch API executor to skip the
		// api execution when the given expression evaluates to true.
		//
		// The expression is given as such a golang template that has to be
		// evaluated to a final content true/false. The expression is an optional and in
		// case not provided, the API will always be executed.
		//
		//
		SkipOnCondition string `json:"SkipOnCondition,omitempty"`

		// The duration in seconds by which the API response is expected from the API target.
		//
		// If the end point does not respond for the API request within this timeout
		// duration, the task will be marked as failed.
		//
		//
		Timeout int64 `json:"Timeout,omitempty"`
	}

	stage1.Body = m.Body

	stage1.ContentType = m.ContentType

	stage1.Name = m.Name

	stage1.Outcomes = m.Outcomes

	stage1.ResponseSpec = m.ResponseSpec

	stage1.SkipOnCondition = m.SkipOnCondition

	stage1.Timeout = m.Timeout

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.WorkflowAPIAO1P1) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.WorkflowAPIAO1P1)
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

// Validate validates this workflow API a o1 p1
func (m *WorkflowAPIAO1P1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponseSpec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var workflowApiAO1P1TypeContentTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["json","xml","text"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		workflowApiAO1P1TypeContentTypePropEnum = append(workflowApiAO1P1TypeContentTypePropEnum, v)
	}
}

const (

	// WorkflowAPIAO1P1ContentTypeJSON captures enum value "json"
	WorkflowAPIAO1P1ContentTypeJSON string = "json"

	// WorkflowAPIAO1P1ContentTypeXML captures enum value "xml"
	WorkflowAPIAO1P1ContentTypeXML string = "xml"

	// WorkflowAPIAO1P1ContentTypeText captures enum value "text"
	WorkflowAPIAO1P1ContentTypeText string = "text"
)

// prop value enum
func (m *WorkflowAPIAO1P1) validateContentTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, workflowApiAO1P1TypeContentTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WorkflowAPIAO1P1) validateContentType(formats strfmt.Registry) error {

	if swag.IsZero(m.ContentType) { // not required
		return nil
	}

	// value enum
	if err := m.validateContentTypeEnum("ContentType", "body", *m.ContentType); err != nil {
		return err
	}

	return nil
}

func (m *WorkflowAPIAO1P1) validateResponseSpec(formats strfmt.Registry) error {

	if swag.IsZero(m.ResponseSpec) { // not required
		return nil
	}

	if m.ResponseSpec != nil {
		if err := m.ResponseSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ResponseSpec")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowAPIAO1P1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowAPIAO1P1) UnmarshalBinary(b []byte) error {
	var res WorkflowAPIAO1P1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
