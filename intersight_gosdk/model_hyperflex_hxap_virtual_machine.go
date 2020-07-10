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

// HyperflexHxapVirtualMachine The Virtual machine that runs on a Hyperflex Application platform compute host.
type HyperflexHxapVirtualMachine struct {
	VirtualizationBaseVirtualMachine
	AffinitySelectors *[]InfraMetaData `json:"AffinitySelectors,omitempty"`
	// Denotes age or life time of the VM in nano seconds.
	Age                   *string            `json:"Age,omitempty"`
	AntiAffinitySelectors *[]InfraMetaData   `json:"AntiAffinitySelectors,omitempty"`
	Disks                 *[]HyperflexVmDisk `json:"Disks,omitempty"`
	// Reason of the failure when VM is in failed state.
	FailureReason *string `json:"FailureReason,omitempty"`
	// Graphical console URL of this VM.
	GraphicConsoleUrl *string                 `json:"GraphicConsoleUrl,omitempty"`
	Interfaces        *[]HyperflexVmInterface `json:"Interfaces,omitempty"`
	Labels            *[]InfraMetaData        `json:"Labels,omitempty"`
	// Number network interfaces associated with the virtual machine.
	NetworkCount *int64 `json:"NetworkCount,omitempty"`
	// Denotes the VM start timestamp.
	StartTime *string `json:"StartTime,omitempty"`
	// Status of virtual machine.
	Status *string `json:"Status,omitempty"`
	// Denotes how long this VM has been running in nano seconds.
	UpTime  *string                           `json:"UpTime,omitempty"`
	Cluster *HyperflexHxapClusterRelationship `json:"Cluster,omitempty"`
	Host    *HyperflexHxapHostRelationship    `json:"Host,omitempty"`
}

// NewHyperflexHxapVirtualMachine instantiates a new HyperflexHxapVirtualMachine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexHxapVirtualMachine() *HyperflexHxapVirtualMachine {
	this := HyperflexHxapVirtualMachine{}
	var status string = "Unknown"
	this.Status = &status
	return &this
}

// NewHyperflexHxapVirtualMachineWithDefaults instantiates a new HyperflexHxapVirtualMachine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexHxapVirtualMachineWithDefaults() *HyperflexHxapVirtualMachine {
	this := HyperflexHxapVirtualMachine{}
	var status string = "Unknown"
	this.Status = &status
	return &this
}

// GetAffinitySelectors returns the AffinitySelectors field value if set, zero value otherwise.
func (o *HyperflexHxapVirtualMachine) GetAffinitySelectors() []InfraMetaData {
	if o == nil || o.AffinitySelectors == nil {
		var ret []InfraMetaData
		return ret
	}
	return *o.AffinitySelectors
}

// GetAffinitySelectorsOk returns a tuple with the AffinitySelectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapVirtualMachine) GetAffinitySelectorsOk() (*[]InfraMetaData, bool) {
	if o == nil || o.AffinitySelectors == nil {
		return nil, false
	}
	return o.AffinitySelectors, true
}

// HasAffinitySelectors returns a boolean if a field has been set.
func (o *HyperflexHxapVirtualMachine) HasAffinitySelectors() bool {
	if o != nil && o.AffinitySelectors != nil {
		return true
	}

	return false
}

// SetAffinitySelectors gets a reference to the given []InfraMetaData and assigns it to the AffinitySelectors field.
func (o *HyperflexHxapVirtualMachine) SetAffinitySelectors(v []InfraMetaData) {
	o.AffinitySelectors = &v
}

// GetAge returns the Age field value if set, zero value otherwise.
func (o *HyperflexHxapVirtualMachine) GetAge() string {
	if o == nil || o.Age == nil {
		var ret string
		return ret
	}
	return *o.Age
}

// GetAgeOk returns a tuple with the Age field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapVirtualMachine) GetAgeOk() (*string, bool) {
	if o == nil || o.Age == nil {
		return nil, false
	}
	return o.Age, true
}

// HasAge returns a boolean if a field has been set.
func (o *HyperflexHxapVirtualMachine) HasAge() bool {
	if o != nil && o.Age != nil {
		return true
	}

	return false
}

// SetAge gets a reference to the given string and assigns it to the Age field.
func (o *HyperflexHxapVirtualMachine) SetAge(v string) {
	o.Age = &v
}

// GetAntiAffinitySelectors returns the AntiAffinitySelectors field value if set, zero value otherwise.
func (o *HyperflexHxapVirtualMachine) GetAntiAffinitySelectors() []InfraMetaData {
	if o == nil || o.AntiAffinitySelectors == nil {
		var ret []InfraMetaData
		return ret
	}
	return *o.AntiAffinitySelectors
}

// GetAntiAffinitySelectorsOk returns a tuple with the AntiAffinitySelectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapVirtualMachine) GetAntiAffinitySelectorsOk() (*[]InfraMetaData, bool) {
	if o == nil || o.AntiAffinitySelectors == nil {
		return nil, false
	}
	return o.AntiAffinitySelectors, true
}

// HasAntiAffinitySelectors returns a boolean if a field has been set.
func (o *HyperflexHxapVirtualMachine) HasAntiAffinitySelectors() bool {
	if o != nil && o.AntiAffinitySelectors != nil {
		return true
	}

	return false
}

// SetAntiAffinitySelectors gets a reference to the given []InfraMetaData and assigns it to the AntiAffinitySelectors field.
func (o *HyperflexHxapVirtualMachine) SetAntiAffinitySelectors(v []InfraMetaData) {
	o.AntiAffinitySelectors = &v
}

// GetDisks returns the Disks field value if set, zero value otherwise.
func (o *HyperflexHxapVirtualMachine) GetDisks() []HyperflexVmDisk {
	if o == nil || o.Disks == nil {
		var ret []HyperflexVmDisk
		return ret
	}
	return *o.Disks
}

// GetDisksOk returns a tuple with the Disks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapVirtualMachine) GetDisksOk() (*[]HyperflexVmDisk, bool) {
	if o == nil || o.Disks == nil {
		return nil, false
	}
	return o.Disks, true
}

// HasDisks returns a boolean if a field has been set.
func (o *HyperflexHxapVirtualMachine) HasDisks() bool {
	if o != nil && o.Disks != nil {
		return true
	}

	return false
}

// SetDisks gets a reference to the given []HyperflexVmDisk and assigns it to the Disks field.
func (o *HyperflexHxapVirtualMachine) SetDisks(v []HyperflexVmDisk) {
	o.Disks = &v
}

// GetFailureReason returns the FailureReason field value if set, zero value otherwise.
func (o *HyperflexHxapVirtualMachine) GetFailureReason() string {
	if o == nil || o.FailureReason == nil {
		var ret string
		return ret
	}
	return *o.FailureReason
}

// GetFailureReasonOk returns a tuple with the FailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapVirtualMachine) GetFailureReasonOk() (*string, bool) {
	if o == nil || o.FailureReason == nil {
		return nil, false
	}
	return o.FailureReason, true
}

// HasFailureReason returns a boolean if a field has been set.
func (o *HyperflexHxapVirtualMachine) HasFailureReason() bool {
	if o != nil && o.FailureReason != nil {
		return true
	}

	return false
}

// SetFailureReason gets a reference to the given string and assigns it to the FailureReason field.
func (o *HyperflexHxapVirtualMachine) SetFailureReason(v string) {
	o.FailureReason = &v
}

// GetGraphicConsoleUrl returns the GraphicConsoleUrl field value if set, zero value otherwise.
func (o *HyperflexHxapVirtualMachine) GetGraphicConsoleUrl() string {
	if o == nil || o.GraphicConsoleUrl == nil {
		var ret string
		return ret
	}
	return *o.GraphicConsoleUrl
}

// GetGraphicConsoleUrlOk returns a tuple with the GraphicConsoleUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapVirtualMachine) GetGraphicConsoleUrlOk() (*string, bool) {
	if o == nil || o.GraphicConsoleUrl == nil {
		return nil, false
	}
	return o.GraphicConsoleUrl, true
}

// HasGraphicConsoleUrl returns a boolean if a field has been set.
func (o *HyperflexHxapVirtualMachine) HasGraphicConsoleUrl() bool {
	if o != nil && o.GraphicConsoleUrl != nil {
		return true
	}

	return false
}

// SetGraphicConsoleUrl gets a reference to the given string and assigns it to the GraphicConsoleUrl field.
func (o *HyperflexHxapVirtualMachine) SetGraphicConsoleUrl(v string) {
	o.GraphicConsoleUrl = &v
}

// GetInterfaces returns the Interfaces field value if set, zero value otherwise.
func (o *HyperflexHxapVirtualMachine) GetInterfaces() []HyperflexVmInterface {
	if o == nil || o.Interfaces == nil {
		var ret []HyperflexVmInterface
		return ret
	}
	return *o.Interfaces
}

// GetInterfacesOk returns a tuple with the Interfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapVirtualMachine) GetInterfacesOk() (*[]HyperflexVmInterface, bool) {
	if o == nil || o.Interfaces == nil {
		return nil, false
	}
	return o.Interfaces, true
}

// HasInterfaces returns a boolean if a field has been set.
func (o *HyperflexHxapVirtualMachine) HasInterfaces() bool {
	if o != nil && o.Interfaces != nil {
		return true
	}

	return false
}

// SetInterfaces gets a reference to the given []HyperflexVmInterface and assigns it to the Interfaces field.
func (o *HyperflexHxapVirtualMachine) SetInterfaces(v []HyperflexVmInterface) {
	o.Interfaces = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *HyperflexHxapVirtualMachine) GetLabels() []InfraMetaData {
	if o == nil || o.Labels == nil {
		var ret []InfraMetaData
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapVirtualMachine) GetLabelsOk() (*[]InfraMetaData, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *HyperflexHxapVirtualMachine) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []InfraMetaData and assigns it to the Labels field.
func (o *HyperflexHxapVirtualMachine) SetLabels(v []InfraMetaData) {
	o.Labels = &v
}

// GetNetworkCount returns the NetworkCount field value if set, zero value otherwise.
func (o *HyperflexHxapVirtualMachine) GetNetworkCount() int64 {
	if o == nil || o.NetworkCount == nil {
		var ret int64
		return ret
	}
	return *o.NetworkCount
}

// GetNetworkCountOk returns a tuple with the NetworkCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapVirtualMachine) GetNetworkCountOk() (*int64, bool) {
	if o == nil || o.NetworkCount == nil {
		return nil, false
	}
	return o.NetworkCount, true
}

// HasNetworkCount returns a boolean if a field has been set.
func (o *HyperflexHxapVirtualMachine) HasNetworkCount() bool {
	if o != nil && o.NetworkCount != nil {
		return true
	}

	return false
}

// SetNetworkCount gets a reference to the given int64 and assigns it to the NetworkCount field.
func (o *HyperflexHxapVirtualMachine) SetNetworkCount(v int64) {
	o.NetworkCount = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *HyperflexHxapVirtualMachine) GetStartTime() string {
	if o == nil || o.StartTime == nil {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapVirtualMachine) GetStartTimeOk() (*string, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *HyperflexHxapVirtualMachine) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *HyperflexHxapVirtualMachine) SetStartTime(v string) {
	o.StartTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HyperflexHxapVirtualMachine) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapVirtualMachine) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HyperflexHxapVirtualMachine) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *HyperflexHxapVirtualMachine) SetStatus(v string) {
	o.Status = &v
}

// GetUpTime returns the UpTime field value if set, zero value otherwise.
func (o *HyperflexHxapVirtualMachine) GetUpTime() string {
	if o == nil || o.UpTime == nil {
		var ret string
		return ret
	}
	return *o.UpTime
}

// GetUpTimeOk returns a tuple with the UpTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapVirtualMachine) GetUpTimeOk() (*string, bool) {
	if o == nil || o.UpTime == nil {
		return nil, false
	}
	return o.UpTime, true
}

// HasUpTime returns a boolean if a field has been set.
func (o *HyperflexHxapVirtualMachine) HasUpTime() bool {
	if o != nil && o.UpTime != nil {
		return true
	}

	return false
}

// SetUpTime gets a reference to the given string and assigns it to the UpTime field.
func (o *HyperflexHxapVirtualMachine) SetUpTime(v string) {
	o.UpTime = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *HyperflexHxapVirtualMachine) GetCluster() HyperflexHxapClusterRelationship {
	if o == nil || o.Cluster == nil {
		var ret HyperflexHxapClusterRelationship
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapVirtualMachine) GetClusterOk() (*HyperflexHxapClusterRelationship, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *HyperflexHxapVirtualMachine) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given HyperflexHxapClusterRelationship and assigns it to the Cluster field.
func (o *HyperflexHxapVirtualMachine) SetCluster(v HyperflexHxapClusterRelationship) {
	o.Cluster = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *HyperflexHxapVirtualMachine) GetHost() HyperflexHxapHostRelationship {
	if o == nil || o.Host == nil {
		var ret HyperflexHxapHostRelationship
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapVirtualMachine) GetHostOk() (*HyperflexHxapHostRelationship, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *HyperflexHxapVirtualMachine) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given HyperflexHxapHostRelationship and assigns it to the Host field.
func (o *HyperflexHxapVirtualMachine) SetHost(v HyperflexHxapHostRelationship) {
	o.Host = &v
}

func (o HyperflexHxapVirtualMachine) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseVirtualMachine, errVirtualizationBaseVirtualMachine := json.Marshal(o.VirtualizationBaseVirtualMachine)
	if errVirtualizationBaseVirtualMachine != nil {
		return []byte{}, errVirtualizationBaseVirtualMachine
	}
	errVirtualizationBaseVirtualMachine = json.Unmarshal([]byte(serializedVirtualizationBaseVirtualMachine), &toSerialize)
	if errVirtualizationBaseVirtualMachine != nil {
		return []byte{}, errVirtualizationBaseVirtualMachine
	}
	if o.AffinitySelectors != nil {
		toSerialize["AffinitySelectors"] = o.AffinitySelectors
	}
	if o.Age != nil {
		toSerialize["Age"] = o.Age
	}
	if o.AntiAffinitySelectors != nil {
		toSerialize["AntiAffinitySelectors"] = o.AntiAffinitySelectors
	}
	if o.Disks != nil {
		toSerialize["Disks"] = o.Disks
	}
	if o.FailureReason != nil {
		toSerialize["FailureReason"] = o.FailureReason
	}
	if o.GraphicConsoleUrl != nil {
		toSerialize["GraphicConsoleUrl"] = o.GraphicConsoleUrl
	}
	if o.Interfaces != nil {
		toSerialize["Interfaces"] = o.Interfaces
	}
	if o.Labels != nil {
		toSerialize["Labels"] = o.Labels
	}
	if o.NetworkCount != nil {
		toSerialize["NetworkCount"] = o.NetworkCount
	}
	if o.StartTime != nil {
		toSerialize["StartTime"] = o.StartTime
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.UpTime != nil {
		toSerialize["UpTime"] = o.UpTime
	}
	if o.Cluster != nil {
		toSerialize["Cluster"] = o.Cluster
	}
	if o.Host != nil {
		toSerialize["Host"] = o.Host
	}
	return json.Marshal(toSerialize)
}

type NullableHyperflexHxapVirtualMachine struct {
	value *HyperflexHxapVirtualMachine
	isSet bool
}

func (v NullableHyperflexHxapVirtualMachine) Get() *HyperflexHxapVirtualMachine {
	return v.value
}

func (v *NullableHyperflexHxapVirtualMachine) Set(val *HyperflexHxapVirtualMachine) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHxapVirtualMachine) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHxapVirtualMachine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHxapVirtualMachine(val *HyperflexHxapVirtualMachine) *NullableHyperflexHxapVirtualMachine {
	return &NullableHyperflexHxapVirtualMachine{value: val, isSet: true}
}

func (v NullableHyperflexHxapVirtualMachine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHxapVirtualMachine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
