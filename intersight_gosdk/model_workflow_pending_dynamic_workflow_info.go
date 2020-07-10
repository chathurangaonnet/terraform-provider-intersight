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

// WorkflowPendingDynamicWorkflowInfo Information for a pending Dynamic Workflow Instance before it is run.  Validation needs to be done on the dynamic workflow tasks before starting.  After it begins, it will be tracked with regular WorkflowInstance.
type WorkflowPendingDynamicWorkflowInfo struct {
	MoBaseMo
	// The inputs of the workflow.
	Input interface{} `json:"Input,omitempty"`
	// A name for the pending dynamic workflow.
	Name            *string   `json:"Name,omitempty"`
	PendingServices *[]string `json:"PendingServices,omitempty"`
	// The src is workflow owner service.
	Src *string `json:"Src,omitempty"`
	// The current status of the PendingDynamicWorkflowInfo.
	Status *string `json:"Status,omitempty"`
	// When set to true workflow engine will wait for a duplicate to finish before starting a new one.
	WaitOnDuplicate         *bool                                    `json:"WaitOnDuplicate,omitempty"`
	WorkflowActionTaskLists *[]WorkflowDynamicWorkflowActionTaskList `json:"WorkflowActionTaskLists,omitempty"`
	// The workflow's workflow context which contains initiator and target information.
	WorkflowCtx interface{} `json:"WorkflowCtx,omitempty"`
	// This key contains workflow, initiator and target name. Workflow engine uses the key to do workflow dedup.
	WorkflowKey *string `json:"WorkflowKey,omitempty"`
	// The metadata of the workflow.
	WorkflowMeta interface{}                       `json:"WorkflowMeta,omitempty"`
	WorkflowInfo *WorkflowWorkflowInfoRelationship `json:"WorkflowInfo,omitempty"`
}

// NewWorkflowPendingDynamicWorkflowInfo instantiates a new WorkflowPendingDynamicWorkflowInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowPendingDynamicWorkflowInfo() *WorkflowPendingDynamicWorkflowInfo {
	this := WorkflowPendingDynamicWorkflowInfo{}
	var status string = "GatheringTasks"
	this.Status = &status
	return &this
}

// NewWorkflowPendingDynamicWorkflowInfoWithDefaults instantiates a new WorkflowPendingDynamicWorkflowInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowPendingDynamicWorkflowInfoWithDefaults() *WorkflowPendingDynamicWorkflowInfo {
	this := WorkflowPendingDynamicWorkflowInfo{}
	var status string = "GatheringTasks"
	this.Status = &status
	return &this
}

// GetInput returns the Input field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowPendingDynamicWorkflowInfo) GetInput() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowPendingDynamicWorkflowInfo) GetInputOk() (*interface{}, bool) {
	if o == nil || o.Input == nil {
		return nil, false
	}
	return &o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *WorkflowPendingDynamicWorkflowInfo) HasInput() bool {
	if o != nil && o.Input != nil {
		return true
	}

	return false
}

// SetInput gets a reference to the given interface{} and assigns it to the Input field.
func (o *WorkflowPendingDynamicWorkflowInfo) SetInput(v interface{}) {
	o.Input = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowPendingDynamicWorkflowInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowPendingDynamicWorkflowInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowPendingDynamicWorkflowInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowPendingDynamicWorkflowInfo) SetName(v string) {
	o.Name = &v
}

// GetPendingServices returns the PendingServices field value if set, zero value otherwise.
func (o *WorkflowPendingDynamicWorkflowInfo) GetPendingServices() []string {
	if o == nil || o.PendingServices == nil {
		var ret []string
		return ret
	}
	return *o.PendingServices
}

// GetPendingServicesOk returns a tuple with the PendingServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowPendingDynamicWorkflowInfo) GetPendingServicesOk() (*[]string, bool) {
	if o == nil || o.PendingServices == nil {
		return nil, false
	}
	return o.PendingServices, true
}

// HasPendingServices returns a boolean if a field has been set.
func (o *WorkflowPendingDynamicWorkflowInfo) HasPendingServices() bool {
	if o != nil && o.PendingServices != nil {
		return true
	}

	return false
}

// SetPendingServices gets a reference to the given []string and assigns it to the PendingServices field.
func (o *WorkflowPendingDynamicWorkflowInfo) SetPendingServices(v []string) {
	o.PendingServices = &v
}

// GetSrc returns the Src field value if set, zero value otherwise.
func (o *WorkflowPendingDynamicWorkflowInfo) GetSrc() string {
	if o == nil || o.Src == nil {
		var ret string
		return ret
	}
	return *o.Src
}

// GetSrcOk returns a tuple with the Src field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowPendingDynamicWorkflowInfo) GetSrcOk() (*string, bool) {
	if o == nil || o.Src == nil {
		return nil, false
	}
	return o.Src, true
}

// HasSrc returns a boolean if a field has been set.
func (o *WorkflowPendingDynamicWorkflowInfo) HasSrc() bool {
	if o != nil && o.Src != nil {
		return true
	}

	return false
}

// SetSrc gets a reference to the given string and assigns it to the Src field.
func (o *WorkflowPendingDynamicWorkflowInfo) SetSrc(v string) {
	o.Src = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WorkflowPendingDynamicWorkflowInfo) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowPendingDynamicWorkflowInfo) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WorkflowPendingDynamicWorkflowInfo) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *WorkflowPendingDynamicWorkflowInfo) SetStatus(v string) {
	o.Status = &v
}

// GetWaitOnDuplicate returns the WaitOnDuplicate field value if set, zero value otherwise.
func (o *WorkflowPendingDynamicWorkflowInfo) GetWaitOnDuplicate() bool {
	if o == nil || o.WaitOnDuplicate == nil {
		var ret bool
		return ret
	}
	return *o.WaitOnDuplicate
}

// GetWaitOnDuplicateOk returns a tuple with the WaitOnDuplicate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowPendingDynamicWorkflowInfo) GetWaitOnDuplicateOk() (*bool, bool) {
	if o == nil || o.WaitOnDuplicate == nil {
		return nil, false
	}
	return o.WaitOnDuplicate, true
}

// HasWaitOnDuplicate returns a boolean if a field has been set.
func (o *WorkflowPendingDynamicWorkflowInfo) HasWaitOnDuplicate() bool {
	if o != nil && o.WaitOnDuplicate != nil {
		return true
	}

	return false
}

// SetWaitOnDuplicate gets a reference to the given bool and assigns it to the WaitOnDuplicate field.
func (o *WorkflowPendingDynamicWorkflowInfo) SetWaitOnDuplicate(v bool) {
	o.WaitOnDuplicate = &v
}

// GetWorkflowActionTaskLists returns the WorkflowActionTaskLists field value if set, zero value otherwise.
func (o *WorkflowPendingDynamicWorkflowInfo) GetWorkflowActionTaskLists() []WorkflowDynamicWorkflowActionTaskList {
	if o == nil || o.WorkflowActionTaskLists == nil {
		var ret []WorkflowDynamicWorkflowActionTaskList
		return ret
	}
	return *o.WorkflowActionTaskLists
}

// GetWorkflowActionTaskListsOk returns a tuple with the WorkflowActionTaskLists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowPendingDynamicWorkflowInfo) GetWorkflowActionTaskListsOk() (*[]WorkflowDynamicWorkflowActionTaskList, bool) {
	if o == nil || o.WorkflowActionTaskLists == nil {
		return nil, false
	}
	return o.WorkflowActionTaskLists, true
}

// HasWorkflowActionTaskLists returns a boolean if a field has been set.
func (o *WorkflowPendingDynamicWorkflowInfo) HasWorkflowActionTaskLists() bool {
	if o != nil && o.WorkflowActionTaskLists != nil {
		return true
	}

	return false
}

// SetWorkflowActionTaskLists gets a reference to the given []WorkflowDynamicWorkflowActionTaskList and assigns it to the WorkflowActionTaskLists field.
func (o *WorkflowPendingDynamicWorkflowInfo) SetWorkflowActionTaskLists(v []WorkflowDynamicWorkflowActionTaskList) {
	o.WorkflowActionTaskLists = &v
}

// GetWorkflowCtx returns the WorkflowCtx field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowPendingDynamicWorkflowInfo) GetWorkflowCtx() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.WorkflowCtx
}

// GetWorkflowCtxOk returns a tuple with the WorkflowCtx field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowPendingDynamicWorkflowInfo) GetWorkflowCtxOk() (*interface{}, bool) {
	if o == nil || o.WorkflowCtx == nil {
		return nil, false
	}
	return &o.WorkflowCtx, true
}

// HasWorkflowCtx returns a boolean if a field has been set.
func (o *WorkflowPendingDynamicWorkflowInfo) HasWorkflowCtx() bool {
	if o != nil && o.WorkflowCtx != nil {
		return true
	}

	return false
}

// SetWorkflowCtx gets a reference to the given interface{} and assigns it to the WorkflowCtx field.
func (o *WorkflowPendingDynamicWorkflowInfo) SetWorkflowCtx(v interface{}) {
	o.WorkflowCtx = v
}

// GetWorkflowKey returns the WorkflowKey field value if set, zero value otherwise.
func (o *WorkflowPendingDynamicWorkflowInfo) GetWorkflowKey() string {
	if o == nil || o.WorkflowKey == nil {
		var ret string
		return ret
	}
	return *o.WorkflowKey
}

// GetWorkflowKeyOk returns a tuple with the WorkflowKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowPendingDynamicWorkflowInfo) GetWorkflowKeyOk() (*string, bool) {
	if o == nil || o.WorkflowKey == nil {
		return nil, false
	}
	return o.WorkflowKey, true
}

// HasWorkflowKey returns a boolean if a field has been set.
func (o *WorkflowPendingDynamicWorkflowInfo) HasWorkflowKey() bool {
	if o != nil && o.WorkflowKey != nil {
		return true
	}

	return false
}

// SetWorkflowKey gets a reference to the given string and assigns it to the WorkflowKey field.
func (o *WorkflowPendingDynamicWorkflowInfo) SetWorkflowKey(v string) {
	o.WorkflowKey = &v
}

// GetWorkflowMeta returns the WorkflowMeta field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowPendingDynamicWorkflowInfo) GetWorkflowMeta() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.WorkflowMeta
}

// GetWorkflowMetaOk returns a tuple with the WorkflowMeta field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowPendingDynamicWorkflowInfo) GetWorkflowMetaOk() (*interface{}, bool) {
	if o == nil || o.WorkflowMeta == nil {
		return nil, false
	}
	return &o.WorkflowMeta, true
}

// HasWorkflowMeta returns a boolean if a field has been set.
func (o *WorkflowPendingDynamicWorkflowInfo) HasWorkflowMeta() bool {
	if o != nil && o.WorkflowMeta != nil {
		return true
	}

	return false
}

// SetWorkflowMeta gets a reference to the given interface{} and assigns it to the WorkflowMeta field.
func (o *WorkflowPendingDynamicWorkflowInfo) SetWorkflowMeta(v interface{}) {
	o.WorkflowMeta = v
}

// GetWorkflowInfo returns the WorkflowInfo field value if set, zero value otherwise.
func (o *WorkflowPendingDynamicWorkflowInfo) GetWorkflowInfo() WorkflowWorkflowInfoRelationship {
	if o == nil || o.WorkflowInfo == nil {
		var ret WorkflowWorkflowInfoRelationship
		return ret
	}
	return *o.WorkflowInfo
}

// GetWorkflowInfoOk returns a tuple with the WorkflowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowPendingDynamicWorkflowInfo) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.WorkflowInfo == nil {
		return nil, false
	}
	return o.WorkflowInfo, true
}

// HasWorkflowInfo returns a boolean if a field has been set.
func (o *WorkflowPendingDynamicWorkflowInfo) HasWorkflowInfo() bool {
	if o != nil && o.WorkflowInfo != nil {
		return true
	}

	return false
}

// SetWorkflowInfo gets a reference to the given WorkflowWorkflowInfoRelationship and assigns it to the WorkflowInfo field.
func (o *WorkflowPendingDynamicWorkflowInfo) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship) {
	o.WorkflowInfo = &v
}

func (o WorkflowPendingDynamicWorkflowInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.Input != nil {
		toSerialize["Input"] = o.Input
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.PendingServices != nil {
		toSerialize["PendingServices"] = o.PendingServices
	}
	if o.Src != nil {
		toSerialize["Src"] = o.Src
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.WaitOnDuplicate != nil {
		toSerialize["WaitOnDuplicate"] = o.WaitOnDuplicate
	}
	if o.WorkflowActionTaskLists != nil {
		toSerialize["WorkflowActionTaskLists"] = o.WorkflowActionTaskLists
	}
	if o.WorkflowCtx != nil {
		toSerialize["WorkflowCtx"] = o.WorkflowCtx
	}
	if o.WorkflowKey != nil {
		toSerialize["WorkflowKey"] = o.WorkflowKey
	}
	if o.WorkflowMeta != nil {
		toSerialize["WorkflowMeta"] = o.WorkflowMeta
	}
	if o.WorkflowInfo != nil {
		toSerialize["WorkflowInfo"] = o.WorkflowInfo
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowPendingDynamicWorkflowInfo struct {
	value *WorkflowPendingDynamicWorkflowInfo
	isSet bool
}

func (v NullableWorkflowPendingDynamicWorkflowInfo) Get() *WorkflowPendingDynamicWorkflowInfo {
	return v.value
}

func (v *NullableWorkflowPendingDynamicWorkflowInfo) Set(val *WorkflowPendingDynamicWorkflowInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowPendingDynamicWorkflowInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowPendingDynamicWorkflowInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowPendingDynamicWorkflowInfo(val *WorkflowPendingDynamicWorkflowInfo) *NullableWorkflowPendingDynamicWorkflowInfo {
	return &NullableWorkflowPendingDynamicWorkflowInfo{value: val, isSet: true}
}

func (v NullableWorkflowPendingDynamicWorkflowInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowPendingDynamicWorkflowInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
