/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-04-12T21:47:47-07:00.
 *
 * API version: 1.0.9-1617
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
)

// WorkflowEnumEntry Captures a single enum entry which has a label and value.
type WorkflowEnumEntry struct {
	MoBaseComplexType
	// Label for the enum value. A user friendly short string to identify the enum value.
	Label *string `json:"Label,omitempty"`
	// Enum value for this enum entry. Value will be passed to the workflow as string type for execution.
	Value *string `json:"Value,omitempty"`
}

// NewWorkflowEnumEntry instantiates a new WorkflowEnumEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowEnumEntry() *WorkflowEnumEntry {
	this := WorkflowEnumEntry{}
	return &this
}

// NewWorkflowEnumEntryWithDefaults instantiates a new WorkflowEnumEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowEnumEntryWithDefaults() *WorkflowEnumEntry {
	this := WorkflowEnumEntry{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WorkflowEnumEntry) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowEnumEntry) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WorkflowEnumEntry) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *WorkflowEnumEntry) SetLabel(v string) {
	o.Label = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *WorkflowEnumEntry) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowEnumEntry) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *WorkflowEnumEntry) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *WorkflowEnumEntry) SetValue(v string) {
	o.Value = &v
}

func (o WorkflowEnumEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.Label != nil {
		toSerialize["Label"] = o.Label
	}
	if o.Value != nil {
		toSerialize["Value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowEnumEntry struct {
	value *WorkflowEnumEntry
	isSet bool
}

func (v NullableWorkflowEnumEntry) Get() *WorkflowEnumEntry {
	return v.value
}

func (v *NullableWorkflowEnumEntry) Set(val *WorkflowEnumEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowEnumEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowEnumEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowEnumEntry(val *WorkflowEnumEntry) *NullableWorkflowEnumEntry {
	return &NullableWorkflowEnumEntry{value: val, isSet: true}
}

func (v NullableWorkflowEnumEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowEnumEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
