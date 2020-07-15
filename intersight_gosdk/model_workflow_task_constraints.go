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

// WorkflowTaskConstraints Captures the constraints for a task. Currently, it hold only targetDataType constraints such as Vendor, ObjectType that are properties of a target device.
type WorkflowTaskConstraints struct {
	MoBaseComplexType
	// List of property constraints that helps to narrow down task implementations based on target device input.
	TargetDataType       interface{} `json:"TargetDataType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowTaskConstraints WorkflowTaskConstraints

// NewWorkflowTaskConstraints instantiates a new WorkflowTaskConstraints object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowTaskConstraints() *WorkflowTaskConstraints {
	this := WorkflowTaskConstraints{}
	return &this
}

// NewWorkflowTaskConstraintsWithDefaults instantiates a new WorkflowTaskConstraints object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowTaskConstraintsWithDefaults() *WorkflowTaskConstraints {
	this := WorkflowTaskConstraints{}
	return &this
}

// GetTargetDataType returns the TargetDataType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowTaskConstraints) GetTargetDataType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.TargetDataType
}

// GetTargetDataTypeOk returns a tuple with the TargetDataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowTaskConstraints) GetTargetDataTypeOk() (*interface{}, bool) {
	if o == nil || o.TargetDataType == nil {
		return nil, false
	}
	return &o.TargetDataType, true
}

// HasTargetDataType returns a boolean if a field has been set.
func (o *WorkflowTaskConstraints) HasTargetDataType() bool {
	if o != nil && o.TargetDataType != nil {
		return true
	}

	return false
}

// SetTargetDataType gets a reference to the given interface{} and assigns it to the TargetDataType field.
func (o *WorkflowTaskConstraints) SetTargetDataType(v interface{}) {
	o.TargetDataType = v
}

func (o WorkflowTaskConstraints) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.TargetDataType != nil {
		toSerialize["TargetDataType"] = o.TargetDataType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowTaskConstraints) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowTaskConstraintsWithoutEmbeddedStruct struct {
		// List of property constraints that helps to narrow down task implementations based on target device input.
		TargetDataType interface{} `json:"TargetDataType,omitempty"`
	}

	varWorkflowTaskConstraintsWithoutEmbeddedStruct := WorkflowTaskConstraintsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowTaskConstraintsWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowTaskConstraints := _WorkflowTaskConstraints{}
		varWorkflowTaskConstraints.TargetDataType = varWorkflowTaskConstraintsWithoutEmbeddedStruct.TargetDataType
		*o = WorkflowTaskConstraints(varWorkflowTaskConstraints)
	} else {
		return err
	}

	varWorkflowTaskConstraints := _WorkflowTaskConstraints{}

	err = json.Unmarshal(bytes, &varWorkflowTaskConstraints)
	if err == nil {
		o.MoBaseComplexType = varWorkflowTaskConstraints.MoBaseComplexType
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "TargetDataType")

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

type NullableWorkflowTaskConstraints struct {
	value *WorkflowTaskConstraints
	isSet bool
}

func (v NullableWorkflowTaskConstraints) Get() *WorkflowTaskConstraints {
	return v.value
}

func (v *NullableWorkflowTaskConstraints) Set(val *WorkflowTaskConstraints) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTaskConstraints) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTaskConstraints) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTaskConstraints(val *WorkflowTaskConstraints) *NullableWorkflowTaskConstraints {
	return &NullableWorkflowTaskConstraints{value: val, isSet: true}
}

func (v NullableWorkflowTaskConstraints) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTaskConstraints) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
