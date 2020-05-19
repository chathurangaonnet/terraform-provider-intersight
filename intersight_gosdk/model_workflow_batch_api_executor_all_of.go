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

// WorkflowBatchApiExecutorAllOf Definition of the list of properties defined in 'workflow.BatchApiExecutor', excluding properties defined in parent classes.
type WorkflowBatchApiExecutorAllOf struct {
	Batch       *[]WorkflowApi           `json:"Batch,omitempty"`
	Constraints *WorkflowTaskConstraints `json:"Constraints,omitempty"`
	// A detailed description about the batch APIs.
	Description *string `json:"Description,omitempty"`
	// Name for the batch API task.
	Name *string `json:"Name,omitempty"`
	// All the possible outcomes of this task are captured here. Outcomes property is a collection property of type workflow.Outcome objects. The outcomes can be mapped to the message to be shown. The outcomes are evaluated in the order they are given. At the end of the outcomes list, an catchall success/fail outcome can be added with condition as 'true'. This is an optional property and if not specified the task will be marked as success.
	Outcomes *map[string]interface{} `json:"Outcomes,omitempty"`
	// Intersight Orchestrator allows the extraction of required values from API responses using the API response grammar. These extracted values can be mapped to task output parameters defined in task definition. The mapping of API output parameters to the task output parameters is provided as JSON in this property.
	Output *map[string]interface{} `json:"Output,omitempty"`
	// The skip expression, if provided, allows the batch API executor to skip the task execution when the given expression evaluates to true. The expression is given as such a golang template that has to be evaluated to a final content true/false. The expression is an optional and in case not provided, the API will always be executed.
	SkipOnCondition *string                             `json:"SkipOnCondition,omitempty"`
	TaskDefinition  *WorkflowTaskDefinitionRelationship `json:"TaskDefinition,omitempty"`
}

// NewWorkflowBatchApiExecutorAllOf instantiates a new WorkflowBatchApiExecutorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowBatchApiExecutorAllOf() *WorkflowBatchApiExecutorAllOf {
	this := WorkflowBatchApiExecutorAllOf{}
	return &this
}

// NewWorkflowBatchApiExecutorAllOfWithDefaults instantiates a new WorkflowBatchApiExecutorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowBatchApiExecutorAllOfWithDefaults() *WorkflowBatchApiExecutorAllOf {
	this := WorkflowBatchApiExecutorAllOf{}
	return &this
}

// GetBatch returns the Batch field value if set, zero value otherwise.
func (o *WorkflowBatchApiExecutorAllOf) GetBatch() []WorkflowApi {
	if o == nil || o.Batch == nil {
		var ret []WorkflowApi
		return ret
	}
	return *o.Batch
}

// GetBatchOk returns a tuple with the Batch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowBatchApiExecutorAllOf) GetBatchOk() (*[]WorkflowApi, bool) {
	if o == nil || o.Batch == nil {
		return nil, false
	}
	return o.Batch, true
}

// HasBatch returns a boolean if a field has been set.
func (o *WorkflowBatchApiExecutorAllOf) HasBatch() bool {
	if o != nil && o.Batch != nil {
		return true
	}

	return false
}

// SetBatch gets a reference to the given []WorkflowApi and assigns it to the Batch field.
func (o *WorkflowBatchApiExecutorAllOf) SetBatch(v []WorkflowApi) {
	o.Batch = &v
}

// GetConstraints returns the Constraints field value if set, zero value otherwise.
func (o *WorkflowBatchApiExecutorAllOf) GetConstraints() WorkflowTaskConstraints {
	if o == nil || o.Constraints == nil {
		var ret WorkflowTaskConstraints
		return ret
	}
	return *o.Constraints
}

// GetConstraintsOk returns a tuple with the Constraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowBatchApiExecutorAllOf) GetConstraintsOk() (*WorkflowTaskConstraints, bool) {
	if o == nil || o.Constraints == nil {
		return nil, false
	}
	return o.Constraints, true
}

// HasConstraints returns a boolean if a field has been set.
func (o *WorkflowBatchApiExecutorAllOf) HasConstraints() bool {
	if o != nil && o.Constraints != nil {
		return true
	}

	return false
}

// SetConstraints gets a reference to the given WorkflowTaskConstraints and assigns it to the Constraints field.
func (o *WorkflowBatchApiExecutorAllOf) SetConstraints(v WorkflowTaskConstraints) {
	o.Constraints = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowBatchApiExecutorAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowBatchApiExecutorAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowBatchApiExecutorAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowBatchApiExecutorAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowBatchApiExecutorAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowBatchApiExecutorAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowBatchApiExecutorAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowBatchApiExecutorAllOf) SetName(v string) {
	o.Name = &v
}

// GetOutcomes returns the Outcomes field value if set, zero value otherwise.
func (o *WorkflowBatchApiExecutorAllOf) GetOutcomes() map[string]interface{} {
	if o == nil || o.Outcomes == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Outcomes
}

// GetOutcomesOk returns a tuple with the Outcomes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowBatchApiExecutorAllOf) GetOutcomesOk() (*map[string]interface{}, bool) {
	if o == nil || o.Outcomes == nil {
		return nil, false
	}
	return o.Outcomes, true
}

// HasOutcomes returns a boolean if a field has been set.
func (o *WorkflowBatchApiExecutorAllOf) HasOutcomes() bool {
	if o != nil && o.Outcomes != nil {
		return true
	}

	return false
}

// SetOutcomes gets a reference to the given map[string]interface{} and assigns it to the Outcomes field.
func (o *WorkflowBatchApiExecutorAllOf) SetOutcomes(v map[string]interface{}) {
	o.Outcomes = &v
}

// GetOutput returns the Output field value if set, zero value otherwise.
func (o *WorkflowBatchApiExecutorAllOf) GetOutput() map[string]interface{} {
	if o == nil || o.Output == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Output
}

// GetOutputOk returns a tuple with the Output field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowBatchApiExecutorAllOf) GetOutputOk() (*map[string]interface{}, bool) {
	if o == nil || o.Output == nil {
		return nil, false
	}
	return o.Output, true
}

// HasOutput returns a boolean if a field has been set.
func (o *WorkflowBatchApiExecutorAllOf) HasOutput() bool {
	if o != nil && o.Output != nil {
		return true
	}

	return false
}

// SetOutput gets a reference to the given map[string]interface{} and assigns it to the Output field.
func (o *WorkflowBatchApiExecutorAllOf) SetOutput(v map[string]interface{}) {
	o.Output = &v
}

// GetSkipOnCondition returns the SkipOnCondition field value if set, zero value otherwise.
func (o *WorkflowBatchApiExecutorAllOf) GetSkipOnCondition() string {
	if o == nil || o.SkipOnCondition == nil {
		var ret string
		return ret
	}
	return *o.SkipOnCondition
}

// GetSkipOnConditionOk returns a tuple with the SkipOnCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowBatchApiExecutorAllOf) GetSkipOnConditionOk() (*string, bool) {
	if o == nil || o.SkipOnCondition == nil {
		return nil, false
	}
	return o.SkipOnCondition, true
}

// HasSkipOnCondition returns a boolean if a field has been set.
func (o *WorkflowBatchApiExecutorAllOf) HasSkipOnCondition() bool {
	if o != nil && o.SkipOnCondition != nil {
		return true
	}

	return false
}

// SetSkipOnCondition gets a reference to the given string and assigns it to the SkipOnCondition field.
func (o *WorkflowBatchApiExecutorAllOf) SetSkipOnCondition(v string) {
	o.SkipOnCondition = &v
}

// GetTaskDefinition returns the TaskDefinition field value if set, zero value otherwise.
func (o *WorkflowBatchApiExecutorAllOf) GetTaskDefinition() WorkflowTaskDefinitionRelationship {
	if o == nil || o.TaskDefinition == nil {
		var ret WorkflowTaskDefinitionRelationship
		return ret
	}
	return *o.TaskDefinition
}

// GetTaskDefinitionOk returns a tuple with the TaskDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowBatchApiExecutorAllOf) GetTaskDefinitionOk() (*WorkflowTaskDefinitionRelationship, bool) {
	if o == nil || o.TaskDefinition == nil {
		return nil, false
	}
	return o.TaskDefinition, true
}

// HasTaskDefinition returns a boolean if a field has been set.
func (o *WorkflowBatchApiExecutorAllOf) HasTaskDefinition() bool {
	if o != nil && o.TaskDefinition != nil {
		return true
	}

	return false
}

// SetTaskDefinition gets a reference to the given WorkflowTaskDefinitionRelationship and assigns it to the TaskDefinition field.
func (o *WorkflowBatchApiExecutorAllOf) SetTaskDefinition(v WorkflowTaskDefinitionRelationship) {
	o.TaskDefinition = &v
}

func (o WorkflowBatchApiExecutorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Batch != nil {
		toSerialize["Batch"] = o.Batch
	}
	if o.Constraints != nil {
		toSerialize["Constraints"] = o.Constraints
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Outcomes != nil {
		toSerialize["Outcomes"] = o.Outcomes
	}
	if o.Output != nil {
		toSerialize["Output"] = o.Output
	}
	if o.SkipOnCondition != nil {
		toSerialize["SkipOnCondition"] = o.SkipOnCondition
	}
	if o.TaskDefinition != nil {
		toSerialize["TaskDefinition"] = o.TaskDefinition
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowBatchApiExecutorAllOf struct {
	value *WorkflowBatchApiExecutorAllOf
	isSet bool
}

func (v NullableWorkflowBatchApiExecutorAllOf) Get() *WorkflowBatchApiExecutorAllOf {
	return v.value
}

func (v *NullableWorkflowBatchApiExecutorAllOf) Set(val *WorkflowBatchApiExecutorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowBatchApiExecutorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowBatchApiExecutorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowBatchApiExecutorAllOf(val *WorkflowBatchApiExecutorAllOf) *NullableWorkflowBatchApiExecutorAllOf {
	return &NullableWorkflowBatchApiExecutorAllOf{value: val, isSet: true}
}

func (v NullableWorkflowBatchApiExecutorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowBatchApiExecutorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
