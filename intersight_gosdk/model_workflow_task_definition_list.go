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

// WorkflowTaskDefinitionList This resource list is returned as a response to a HTTP GET request that does not include a specific resource identifier.
type WorkflowTaskDefinitionList struct {
	MoBaseResponse
	// The total number of 'workflow.TaskDefinition' resources matching the request, accross all pages. The 'Count' attribute is included when the HTTP GET request includes the '$inlinecount' parameter.
	Count *int32 `json:"Count,omitempty"`
	// The array of 'workflow.TaskDefinition' resources matching the request.
	Results *[]WorkflowTaskDefinition `json:"Results,omitempty"`
}

// NewWorkflowTaskDefinitionList instantiates a new WorkflowTaskDefinitionList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowTaskDefinitionList() *WorkflowTaskDefinitionList {
	this := WorkflowTaskDefinitionList{}
	return &this
}

// NewWorkflowTaskDefinitionListWithDefaults instantiates a new WorkflowTaskDefinitionList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowTaskDefinitionListWithDefaults() *WorkflowTaskDefinitionList {
	this := WorkflowTaskDefinitionList{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *WorkflowTaskDefinitionList) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskDefinitionList) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *WorkflowTaskDefinitionList) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *WorkflowTaskDefinitionList) SetCount(v int32) {
	o.Count = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *WorkflowTaskDefinitionList) GetResults() []WorkflowTaskDefinition {
	if o == nil || o.Results == nil {
		var ret []WorkflowTaskDefinition
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskDefinitionList) GetResultsOk() (*[]WorkflowTaskDefinition, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *WorkflowTaskDefinitionList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []WorkflowTaskDefinition and assigns it to the Results field.
func (o *WorkflowTaskDefinitionList) SetResults(v []WorkflowTaskDefinition) {
	o.Results = &v
}

func (o WorkflowTaskDefinitionList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseResponse, errMoBaseResponse := json.Marshal(o.MoBaseResponse)
	if errMoBaseResponse != nil {
		return []byte{}, errMoBaseResponse
	}
	errMoBaseResponse = json.Unmarshal([]byte(serializedMoBaseResponse), &toSerialize)
	if errMoBaseResponse != nil {
		return []byte{}, errMoBaseResponse
	}
	if o.Count != nil {
		toSerialize["Count"] = o.Count
	}
	if o.Results != nil {
		toSerialize["Results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

// AsWorkflowTaskDefinitionResponse wraps this instance of WorkflowTaskDefinitionList in WorkflowTaskDefinitionResponse
func (s *WorkflowTaskDefinitionList) AsWorkflowTaskDefinitionResponse() WorkflowTaskDefinitionResponse {
	return WorkflowTaskDefinitionResponse{WorkflowTaskDefinitionResponseInterface: s}
}

type NullableWorkflowTaskDefinitionList struct {
	value *WorkflowTaskDefinitionList
	isSet bool
}

func (v NullableWorkflowTaskDefinitionList) Get() *WorkflowTaskDefinitionList {
	return v.value
}

func (v *NullableWorkflowTaskDefinitionList) Set(val *WorkflowTaskDefinitionList) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTaskDefinitionList) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTaskDefinitionList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTaskDefinitionList(val *WorkflowTaskDefinitionList) *NullableWorkflowTaskDefinitionList {
	return &NullableWorkflowTaskDefinitionList{value: val, isSet: true}
}

func (v NullableWorkflowTaskDefinitionList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTaskDefinitionList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
