/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-06-19T15:15:17Z.
 *
 * API version: 1.0.9-1903
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
)

// WorkflowWorkflowInfoProperties Properties for a workflowInfo.
type WorkflowWorkflowInfoProperties struct {
	MoBaseComplexType
	// When true, this workflow can be retried if has not been modified for more than a period of 2 weeks.
	Retryable *bool `json:"Retryable,omitempty"`
}

// NewWorkflowWorkflowInfoProperties instantiates a new WorkflowWorkflowInfoProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowWorkflowInfoProperties() *WorkflowWorkflowInfoProperties {
	this := WorkflowWorkflowInfoProperties{}
	return &this
}

// NewWorkflowWorkflowInfoPropertiesWithDefaults instantiates a new WorkflowWorkflowInfoProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowWorkflowInfoPropertiesWithDefaults() *WorkflowWorkflowInfoProperties {
	this := WorkflowWorkflowInfoProperties{}
	return &this
}

// GetRetryable returns the Retryable field value if set, zero value otherwise.
func (o *WorkflowWorkflowInfoProperties) GetRetryable() bool {
	if o == nil || o.Retryable == nil {
		var ret bool
		return ret
	}
	return *o.Retryable
}

// GetRetryableOk returns a tuple with the Retryable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowInfoProperties) GetRetryableOk() (*bool, bool) {
	if o == nil || o.Retryable == nil {
		return nil, false
	}
	return o.Retryable, true
}

// HasRetryable returns a boolean if a field has been set.
func (o *WorkflowWorkflowInfoProperties) HasRetryable() bool {
	if o != nil && o.Retryable != nil {
		return true
	}

	return false
}

// SetRetryable gets a reference to the given bool and assigns it to the Retryable field.
func (o *WorkflowWorkflowInfoProperties) SetRetryable(v bool) {
	o.Retryable = &v
}

func (o WorkflowWorkflowInfoProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.Retryable != nil {
		toSerialize["Retryable"] = o.Retryable
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowWorkflowInfoProperties struct {
	value *WorkflowWorkflowInfoProperties
	isSet bool
}

func (v NullableWorkflowWorkflowInfoProperties) Get() *WorkflowWorkflowInfoProperties {
	return v.value
}

func (v *NullableWorkflowWorkflowInfoProperties) Set(val *WorkflowWorkflowInfoProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowWorkflowInfoProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowWorkflowInfoProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowWorkflowInfoProperties(val *WorkflowWorkflowInfoProperties) *NullableWorkflowWorkflowInfoProperties {
	return &NullableWorkflowWorkflowInfoProperties{value: val, isSet: true}
}

func (v NullableWorkflowWorkflowInfoProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowWorkflowInfoProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}