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

// VirtualizationBaseCluster Common attributes of a cluster of resources within a datacenter. A cluster is a convenient grouping of resources such as Host, Datastore, etc.
type VirtualizationBaseCluster struct {
	MoBaseMo
	// Identifies the broad type of the underlying hypervisor.
	HypervisorType *string `json:"HypervisorType,omitempty"`
	// The internally generated identity of this cluster. This entity is not manipulated by users. It aids in uniquely identifying the cluster object. In case of VMware, this is a MOR (managed object reference).
	Identity       *string                       `json:"Identity,omitempty"`
	MemoryCapacity *VirtualizationMemoryCapacity `json:"MemoryCapacity,omitempty"`
	// The user-provided name for this cluster to facilitate identification.
	Name              *string                        `json:"Name,omitempty"`
	ProcessorCapacity *VirtualizationComputeCapacity `json:"ProcessorCapacity,omitempty"`
	// Cluster health status as reported by the hypervisor platform.
	Status *string `json:"Status,omitempty"`
	// Total number of CPU cores in this cluster. It is a cumulative number across all hosts in the cluster.
	TotalCores *int64 `json:"TotalCores,omitempty"`
}

// NewVirtualizationBaseCluster instantiates a new VirtualizationBaseCluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationBaseCluster() *VirtualizationBaseCluster {
	this := VirtualizationBaseCluster{}
	var hypervisorType string = "Unknown"
	this.HypervisorType = &hypervisorType
	var status string = "Unknown"
	this.Status = &status
	return &this
}

// NewVirtualizationBaseClusterWithDefaults instantiates a new VirtualizationBaseCluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationBaseClusterWithDefaults() *VirtualizationBaseCluster {
	this := VirtualizationBaseCluster{}
	var hypervisorType string = "Unknown"
	this.HypervisorType = &hypervisorType
	var status string = "Unknown"
	this.Status = &status
	return &this
}

// GetHypervisorType returns the HypervisorType field value if set, zero value otherwise.
func (o *VirtualizationBaseCluster) GetHypervisorType() string {
	if o == nil || o.HypervisorType == nil {
		var ret string
		return ret
	}
	return *o.HypervisorType
}

// GetHypervisorTypeOk returns a tuple with the HypervisorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseCluster) GetHypervisorTypeOk() (*string, bool) {
	if o == nil || o.HypervisorType == nil {
		return nil, false
	}
	return o.HypervisorType, true
}

// HasHypervisorType returns a boolean if a field has been set.
func (o *VirtualizationBaseCluster) HasHypervisorType() bool {
	if o != nil && o.HypervisorType != nil {
		return true
	}

	return false
}

// SetHypervisorType gets a reference to the given string and assigns it to the HypervisorType field.
func (o *VirtualizationBaseCluster) SetHypervisorType(v string) {
	o.HypervisorType = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *VirtualizationBaseCluster) GetIdentity() string {
	if o == nil || o.Identity == nil {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseCluster) GetIdentityOk() (*string, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *VirtualizationBaseCluster) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *VirtualizationBaseCluster) SetIdentity(v string) {
	o.Identity = &v
}

// GetMemoryCapacity returns the MemoryCapacity field value if set, zero value otherwise.
func (o *VirtualizationBaseCluster) GetMemoryCapacity() VirtualizationMemoryCapacity {
	if o == nil || o.MemoryCapacity == nil {
		var ret VirtualizationMemoryCapacity
		return ret
	}
	return *o.MemoryCapacity
}

// GetMemoryCapacityOk returns a tuple with the MemoryCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseCluster) GetMemoryCapacityOk() (*VirtualizationMemoryCapacity, bool) {
	if o == nil || o.MemoryCapacity == nil {
		return nil, false
	}
	return o.MemoryCapacity, true
}

// HasMemoryCapacity returns a boolean if a field has been set.
func (o *VirtualizationBaseCluster) HasMemoryCapacity() bool {
	if o != nil && o.MemoryCapacity != nil {
		return true
	}

	return false
}

// SetMemoryCapacity gets a reference to the given VirtualizationMemoryCapacity and assigns it to the MemoryCapacity field.
func (o *VirtualizationBaseCluster) SetMemoryCapacity(v VirtualizationMemoryCapacity) {
	o.MemoryCapacity = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VirtualizationBaseCluster) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseCluster) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VirtualizationBaseCluster) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VirtualizationBaseCluster) SetName(v string) {
	o.Name = &v
}

// GetProcessorCapacity returns the ProcessorCapacity field value if set, zero value otherwise.
func (o *VirtualizationBaseCluster) GetProcessorCapacity() VirtualizationComputeCapacity {
	if o == nil || o.ProcessorCapacity == nil {
		var ret VirtualizationComputeCapacity
		return ret
	}
	return *o.ProcessorCapacity
}

// GetProcessorCapacityOk returns a tuple with the ProcessorCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseCluster) GetProcessorCapacityOk() (*VirtualizationComputeCapacity, bool) {
	if o == nil || o.ProcessorCapacity == nil {
		return nil, false
	}
	return o.ProcessorCapacity, true
}

// HasProcessorCapacity returns a boolean if a field has been set.
func (o *VirtualizationBaseCluster) HasProcessorCapacity() bool {
	if o != nil && o.ProcessorCapacity != nil {
		return true
	}

	return false
}

// SetProcessorCapacity gets a reference to the given VirtualizationComputeCapacity and assigns it to the ProcessorCapacity field.
func (o *VirtualizationBaseCluster) SetProcessorCapacity(v VirtualizationComputeCapacity) {
	o.ProcessorCapacity = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *VirtualizationBaseCluster) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseCluster) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *VirtualizationBaseCluster) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *VirtualizationBaseCluster) SetStatus(v string) {
	o.Status = &v
}

// GetTotalCores returns the TotalCores field value if set, zero value otherwise.
func (o *VirtualizationBaseCluster) GetTotalCores() int64 {
	if o == nil || o.TotalCores == nil {
		var ret int64
		return ret
	}
	return *o.TotalCores
}

// GetTotalCoresOk returns a tuple with the TotalCores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseCluster) GetTotalCoresOk() (*int64, bool) {
	if o == nil || o.TotalCores == nil {
		return nil, false
	}
	return o.TotalCores, true
}

// HasTotalCores returns a boolean if a field has been set.
func (o *VirtualizationBaseCluster) HasTotalCores() bool {
	if o != nil && o.TotalCores != nil {
		return true
	}

	return false
}

// SetTotalCores gets a reference to the given int64 and assigns it to the TotalCores field.
func (o *VirtualizationBaseCluster) SetTotalCores(v int64) {
	o.TotalCores = &v
}

func (o VirtualizationBaseCluster) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.HypervisorType != nil {
		toSerialize["HypervisorType"] = o.HypervisorType
	}
	if o.Identity != nil {
		toSerialize["Identity"] = o.Identity
	}
	if o.MemoryCapacity != nil {
		toSerialize["MemoryCapacity"] = o.MemoryCapacity
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.ProcessorCapacity != nil {
		toSerialize["ProcessorCapacity"] = o.ProcessorCapacity
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.TotalCores != nil {
		toSerialize["TotalCores"] = o.TotalCores
	}
	return json.Marshal(toSerialize)
}

type NullableVirtualizationBaseCluster struct {
	value *VirtualizationBaseCluster
	isSet bool
}

func (v NullableVirtualizationBaseCluster) Get() *VirtualizationBaseCluster {
	return v.value
}

func (v *NullableVirtualizationBaseCluster) Set(val *VirtualizationBaseCluster) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationBaseCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationBaseCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationBaseCluster(val *VirtualizationBaseCluster) *NullableVirtualizationBaseCluster {
	return &NullableVirtualizationBaseCluster{value: val, isSet: true}
}

func (v NullableVirtualizationBaseCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationBaseCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
