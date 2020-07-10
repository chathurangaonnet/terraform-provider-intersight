/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-07-08T07:46:03Z.
 *
 * API version: 0.0.1-15983
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
)

// WorkflowCliCommandAllOf Definition of the list of properties defined in 'workflow.CliCommand', excluding properties defined in parent classes.
type WorkflowCliCommandAllOf struct {
	// The command to run on the device connector.
	Command *string `json:"Command,omitempty"`
	// The regex string that identifies the end of the command response.
	EndPrompt     *string                 `json:"EndPrompt,omitempty"`
	ExpectPrompts *[]WorkflowExpectPrompt `json:"ExpectPrompts,omitempty"`
	// Skips the execution status check of the terminal command. One use case for this is while exiting the terminal session from esxi host.
	SkipStatusCheck *bool `json:"SkipStatusCheck,omitempty"`
	// If this flag is set, it marks the end of the terminal session where the previous commands were executed.
	TerminalEnd *bool `json:"TerminalEnd,omitempty"`
	// If this flag is set, the execution of this command initiates a terminal session in which the subsequent CLI commands are executed until a command with terminalEnd flag is encountered or the end of the batch.
	TerminalStart *bool `json:"TerminalStart,omitempty"`
	// The type of the command - can be interactive or non-interactive.
	Type *string `json:"Type,omitempty"`
}

// NewWorkflowCliCommandAllOf instantiates a new WorkflowCliCommandAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowCliCommandAllOf() *WorkflowCliCommandAllOf {
	this := WorkflowCliCommandAllOf{}
	var type_ string = "NonInteractive"
	this.Type = &type_
	return &this
}

// NewWorkflowCliCommandAllOfWithDefaults instantiates a new WorkflowCliCommandAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowCliCommandAllOfWithDefaults() *WorkflowCliCommandAllOf {
	this := WorkflowCliCommandAllOf{}
	var type_ string = "NonInteractive"
	this.Type = &type_
	return &this
}

// GetCommand returns the Command field value if set, zero value otherwise.
func (o *WorkflowCliCommandAllOf) GetCommand() string {
	if o == nil || o.Command == nil {
		var ret string
		return ret
	}
	return *o.Command
}

// GetCommandOk returns a tuple with the Command field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCliCommandAllOf) GetCommandOk() (*string, bool) {
	if o == nil || o.Command == nil {
		return nil, false
	}
	return o.Command, true
}

// HasCommand returns a boolean if a field has been set.
func (o *WorkflowCliCommandAllOf) HasCommand() bool {
	if o != nil && o.Command != nil {
		return true
	}

	return false
}

// SetCommand gets a reference to the given string and assigns it to the Command field.
func (o *WorkflowCliCommandAllOf) SetCommand(v string) {
	o.Command = &v
}

// GetEndPrompt returns the EndPrompt field value if set, zero value otherwise.
func (o *WorkflowCliCommandAllOf) GetEndPrompt() string {
	if o == nil || o.EndPrompt == nil {
		var ret string
		return ret
	}
	return *o.EndPrompt
}

// GetEndPromptOk returns a tuple with the EndPrompt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCliCommandAllOf) GetEndPromptOk() (*string, bool) {
	if o == nil || o.EndPrompt == nil {
		return nil, false
	}
	return o.EndPrompt, true
}

// HasEndPrompt returns a boolean if a field has been set.
func (o *WorkflowCliCommandAllOf) HasEndPrompt() bool {
	if o != nil && o.EndPrompt != nil {
		return true
	}

	return false
}

// SetEndPrompt gets a reference to the given string and assigns it to the EndPrompt field.
func (o *WorkflowCliCommandAllOf) SetEndPrompt(v string) {
	o.EndPrompt = &v
}

// GetExpectPrompts returns the ExpectPrompts field value if set, zero value otherwise.
func (o *WorkflowCliCommandAllOf) GetExpectPrompts() []WorkflowExpectPrompt {
	if o == nil || o.ExpectPrompts == nil {
		var ret []WorkflowExpectPrompt
		return ret
	}
	return *o.ExpectPrompts
}

// GetExpectPromptsOk returns a tuple with the ExpectPrompts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCliCommandAllOf) GetExpectPromptsOk() (*[]WorkflowExpectPrompt, bool) {
	if o == nil || o.ExpectPrompts == nil {
		return nil, false
	}
	return o.ExpectPrompts, true
}

// HasExpectPrompts returns a boolean if a field has been set.
func (o *WorkflowCliCommandAllOf) HasExpectPrompts() bool {
	if o != nil && o.ExpectPrompts != nil {
		return true
	}

	return false
}

// SetExpectPrompts gets a reference to the given []WorkflowExpectPrompt and assigns it to the ExpectPrompts field.
func (o *WorkflowCliCommandAllOf) SetExpectPrompts(v []WorkflowExpectPrompt) {
	o.ExpectPrompts = &v
}

// GetSkipStatusCheck returns the SkipStatusCheck field value if set, zero value otherwise.
func (o *WorkflowCliCommandAllOf) GetSkipStatusCheck() bool {
	if o == nil || o.SkipStatusCheck == nil {
		var ret bool
		return ret
	}
	return *o.SkipStatusCheck
}

// GetSkipStatusCheckOk returns a tuple with the SkipStatusCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCliCommandAllOf) GetSkipStatusCheckOk() (*bool, bool) {
	if o == nil || o.SkipStatusCheck == nil {
		return nil, false
	}
	return o.SkipStatusCheck, true
}

// HasSkipStatusCheck returns a boolean if a field has been set.
func (o *WorkflowCliCommandAllOf) HasSkipStatusCheck() bool {
	if o != nil && o.SkipStatusCheck != nil {
		return true
	}

	return false
}

// SetSkipStatusCheck gets a reference to the given bool and assigns it to the SkipStatusCheck field.
func (o *WorkflowCliCommandAllOf) SetSkipStatusCheck(v bool) {
	o.SkipStatusCheck = &v
}

// GetTerminalEnd returns the TerminalEnd field value if set, zero value otherwise.
func (o *WorkflowCliCommandAllOf) GetTerminalEnd() bool {
	if o == nil || o.TerminalEnd == nil {
		var ret bool
		return ret
	}
	return *o.TerminalEnd
}

// GetTerminalEndOk returns a tuple with the TerminalEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCliCommandAllOf) GetTerminalEndOk() (*bool, bool) {
	if o == nil || o.TerminalEnd == nil {
		return nil, false
	}
	return o.TerminalEnd, true
}

// HasTerminalEnd returns a boolean if a field has been set.
func (o *WorkflowCliCommandAllOf) HasTerminalEnd() bool {
	if o != nil && o.TerminalEnd != nil {
		return true
	}

	return false
}

// SetTerminalEnd gets a reference to the given bool and assigns it to the TerminalEnd field.
func (o *WorkflowCliCommandAllOf) SetTerminalEnd(v bool) {
	o.TerminalEnd = &v
}

// GetTerminalStart returns the TerminalStart field value if set, zero value otherwise.
func (o *WorkflowCliCommandAllOf) GetTerminalStart() bool {
	if o == nil || o.TerminalStart == nil {
		var ret bool
		return ret
	}
	return *o.TerminalStart
}

// GetTerminalStartOk returns a tuple with the TerminalStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCliCommandAllOf) GetTerminalStartOk() (*bool, bool) {
	if o == nil || o.TerminalStart == nil {
		return nil, false
	}
	return o.TerminalStart, true
}

// HasTerminalStart returns a boolean if a field has been set.
func (o *WorkflowCliCommandAllOf) HasTerminalStart() bool {
	if o != nil && o.TerminalStart != nil {
		return true
	}

	return false
}

// SetTerminalStart gets a reference to the given bool and assigns it to the TerminalStart field.
func (o *WorkflowCliCommandAllOf) SetTerminalStart(v bool) {
	o.TerminalStart = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WorkflowCliCommandAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCliCommandAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WorkflowCliCommandAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WorkflowCliCommandAllOf) SetType(v string) {
	o.Type = &v
}

func (o WorkflowCliCommandAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Command != nil {
		toSerialize["Command"] = o.Command
	}
	if o.EndPrompt != nil {
		toSerialize["EndPrompt"] = o.EndPrompt
	}
	if o.ExpectPrompts != nil {
		toSerialize["ExpectPrompts"] = o.ExpectPrompts
	}
	if o.SkipStatusCheck != nil {
		toSerialize["SkipStatusCheck"] = o.SkipStatusCheck
	}
	if o.TerminalEnd != nil {
		toSerialize["TerminalEnd"] = o.TerminalEnd
	}
	if o.TerminalStart != nil {
		toSerialize["TerminalStart"] = o.TerminalStart
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowCliCommandAllOf struct {
	value *WorkflowCliCommandAllOf
	isSet bool
}

func (v NullableWorkflowCliCommandAllOf) Get() *WorkflowCliCommandAllOf {
	return v.value
}

func (v *NullableWorkflowCliCommandAllOf) Set(val *WorkflowCliCommandAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowCliCommandAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowCliCommandAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowCliCommandAllOf(val *WorkflowCliCommandAllOf) *NullableWorkflowCliCommandAllOf {
	return &NullableWorkflowCliCommandAllOf{value: val, isSet: true}
}

func (v NullableWorkflowCliCommandAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowCliCommandAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
