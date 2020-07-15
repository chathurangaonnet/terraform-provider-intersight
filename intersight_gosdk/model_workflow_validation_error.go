/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-07-11T05:47:33Z.
 *
 * API version: 1.0.9-1999
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// WorkflowValidationError Used to show validation errors for the workflow or for the task within the workflow.
type WorkflowValidationError struct {
	MoBaseComplexType
	// Description of the error.} type: string
	ErrorLog interface{} `json:"ErrorLog,omitempty"`
	// When populated this refers to the input or output field within the workflow or task.
	Field *string `json:"Field,omitempty"`
	// The task name on which the error is found, when empty the error applies to the top level workflow.
	TaskName *string `json:"TaskName,omitempty"`
	// When populated this refers to the transition connection that has a problem. When this field has value OnSuccess it means the transition connection OnSuccess for the task has an issue.
	TransitionName       *string `json:"TransitionName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowValidationError WorkflowValidationError

// NewWorkflowValidationError instantiates a new WorkflowValidationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowValidationError() *WorkflowValidationError {
	this := WorkflowValidationError{}
	return &this
}

// NewWorkflowValidationErrorWithDefaults instantiates a new WorkflowValidationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowValidationErrorWithDefaults() *WorkflowValidationError {
	this := WorkflowValidationError{}
	return &this
}

// GetErrorLog returns the ErrorLog field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowValidationError) GetErrorLog() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ErrorLog
}

// GetErrorLogOk returns a tuple with the ErrorLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowValidationError) GetErrorLogOk() (*interface{}, bool) {
	if o == nil || o.ErrorLog == nil {
		return nil, false
	}
	return &o.ErrorLog, true
}

// HasErrorLog returns a boolean if a field has been set.
func (o *WorkflowValidationError) HasErrorLog() bool {
	if o != nil && o.ErrorLog != nil {
		return true
	}

	return false
}

// SetErrorLog gets a reference to the given interface{} and assigns it to the ErrorLog field.
func (o *WorkflowValidationError) SetErrorLog(v interface{}) {
	o.ErrorLog = v
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *WorkflowValidationError) GetField() string {
	if o == nil || o.Field == nil {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowValidationError) GetFieldOk() (*string, bool) {
	if o == nil || o.Field == nil {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *WorkflowValidationError) HasField() bool {
	if o != nil && o.Field != nil {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *WorkflowValidationError) SetField(v string) {
	o.Field = &v
}

// GetTaskName returns the TaskName field value if set, zero value otherwise.
func (o *WorkflowValidationError) GetTaskName() string {
	if o == nil || o.TaskName == nil {
		var ret string
		return ret
	}
	return *o.TaskName
}

// GetTaskNameOk returns a tuple with the TaskName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowValidationError) GetTaskNameOk() (*string, bool) {
	if o == nil || o.TaskName == nil {
		return nil, false
	}
	return o.TaskName, true
}

// HasTaskName returns a boolean if a field has been set.
func (o *WorkflowValidationError) HasTaskName() bool {
	if o != nil && o.TaskName != nil {
		return true
	}

	return false
}

// SetTaskName gets a reference to the given string and assigns it to the TaskName field.
func (o *WorkflowValidationError) SetTaskName(v string) {
	o.TaskName = &v
}

// GetTransitionName returns the TransitionName field value if set, zero value otherwise.
func (o *WorkflowValidationError) GetTransitionName() string {
	if o == nil || o.TransitionName == nil {
		var ret string
		return ret
	}
	return *o.TransitionName
}

// GetTransitionNameOk returns a tuple with the TransitionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowValidationError) GetTransitionNameOk() (*string, bool) {
	if o == nil || o.TransitionName == nil {
		return nil, false
	}
	return o.TransitionName, true
}

// HasTransitionName returns a boolean if a field has been set.
func (o *WorkflowValidationError) HasTransitionName() bool {
	if o != nil && o.TransitionName != nil {
		return true
	}

	return false
}

// SetTransitionName gets a reference to the given string and assigns it to the TransitionName field.
func (o *WorkflowValidationError) SetTransitionName(v string) {
	o.TransitionName = &v
}

func (o WorkflowValidationError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.ErrorLog != nil {
		toSerialize["ErrorLog"] = o.ErrorLog
	}
	if o.Field != nil {
		toSerialize["Field"] = o.Field
	}
	if o.TaskName != nil {
		toSerialize["TaskName"] = o.TaskName
	}
	if o.TransitionName != nil {
		toSerialize["TransitionName"] = o.TransitionName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowValidationError) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowValidationErrorWithoutEmbeddedStruct struct {
		// Description of the error.} type: string
		ErrorLog interface{} `json:"ErrorLog,omitempty"`
		// When populated this refers to the input or output field within the workflow or task.
		Field *string `json:"Field,omitempty"`
		// The task name on which the error is found, when empty the error applies to the top level workflow.
		TaskName *string `json:"TaskName,omitempty"`
		// When populated this refers to the transition connection that has a problem. When this field has value OnSuccess it means the transition connection OnSuccess for the task has an issue.
		TransitionName *string `json:"TransitionName,omitempty"`
	}

	varWorkflowValidationErrorWithoutEmbeddedStruct := WorkflowValidationErrorWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowValidationErrorWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowValidationError := _WorkflowValidationError{}
		varWorkflowValidationError.ErrorLog = varWorkflowValidationErrorWithoutEmbeddedStruct.ErrorLog
		varWorkflowValidationError.Field = varWorkflowValidationErrorWithoutEmbeddedStruct.Field
		varWorkflowValidationError.TaskName = varWorkflowValidationErrorWithoutEmbeddedStruct.TaskName
		varWorkflowValidationError.TransitionName = varWorkflowValidationErrorWithoutEmbeddedStruct.TransitionName
		*o = WorkflowValidationError(varWorkflowValidationError)
	} else {
		return err
	}

	varWorkflowValidationError := _WorkflowValidationError{}

	err = json.Unmarshal(bytes, &varWorkflowValidationError)
	if err == nil {
		o.MoBaseComplexType = varWorkflowValidationError.MoBaseComplexType
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ErrorLog")
		delete(additionalProperties, "Field")
		delete(additionalProperties, "TaskName")
		delete(additionalProperties, "TransitionName")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowValidationError struct {
	value *WorkflowValidationError
	isSet bool
}

func (v NullableWorkflowValidationError) Get() *WorkflowValidationError {
	return v.value
}

func (v *NullableWorkflowValidationError) Set(val *WorkflowValidationError) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowValidationError) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowValidationError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowValidationError(val *WorkflowValidationError) *NullableWorkflowValidationError {
	return &NullableWorkflowValidationError{value: val, isSet: true}
}

func (v NullableWorkflowValidationError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowValidationError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
