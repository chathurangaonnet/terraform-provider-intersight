/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-06-17T02:04:50-07:00.
 *
 * API version: 1.0.9-1867
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// VirtualizationBaseVirtualMachine Common attributes of a virtual machine managed by a  hypervisor. Serves as a base class for all concrete virtual machine types. A virtual machine (VM) is what a user and applications interact with. A VM usually runs a guest OS and applications run on the guest OS.
type VirtualizationBaseVirtualMachine struct {
	VirtualizationBaseSourceDevice
	Capacity  *InfraHardwareInfo       `json:"Capacity,omitempty"`
	GuestInfo *VirtualizationGuestInfo `json:"GuestInfo,omitempty"`
	// Type of hypervisor where the virtual machine is hosted, for example VMware ESXi.
	HypervisorType *string `json:"HypervisorType,omitempty"`
	// The internally generated identity of this VM. This entity is not manipulated by users. It aids in uniquely identifying the virtual machine object. For VMware, this is MOR (managed object reference).
	Identity       *string                       `json:"Identity,omitempty"`
	IpAddress      *[]string                     `json:"IpAddress,omitempty"`
	MemoryCapacity *VirtualizationMemoryCapacity `json:"MemoryCapacity,omitempty"`
	// User-provided name to identify the virtual machine.
	Name *string `json:"Name,omitempty"`
	// Power state of the virtual machine.
	PowerState        *string                        `json:"PowerState,omitempty"`
	ProcessorCapacity *VirtualizationComputeCapacity `json:"ProcessorCapacity,omitempty"`
	// The uuid of this virtual machine. The uuid is internally generated and not user specified.
	Uuid                 *string `json:"Uuid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationBaseVirtualMachine VirtualizationBaseVirtualMachine

// NewVirtualizationBaseVirtualMachine instantiates a new VirtualizationBaseVirtualMachine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationBaseVirtualMachine() *VirtualizationBaseVirtualMachine {
	this := VirtualizationBaseVirtualMachine{}
	var powerState string = "Unknown"
	this.PowerState = &powerState
	return &this
}

// NewVirtualizationBaseVirtualMachineWithDefaults instantiates a new VirtualizationBaseVirtualMachine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationBaseVirtualMachineWithDefaults() *VirtualizationBaseVirtualMachine {
	this := VirtualizationBaseVirtualMachine{}
	var powerState string = "Unknown"
	this.PowerState = &powerState
	return &this
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualMachine) GetCapacity() InfraHardwareInfo {
	if o == nil || o.Capacity == nil {
		var ret InfraHardwareInfo
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachine) GetCapacityOk() (*InfraHardwareInfo, bool) {
	if o == nil || o.Capacity == nil {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualMachine) HasCapacity() bool {
	if o != nil && o.Capacity != nil {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given InfraHardwareInfo and assigns it to the Capacity field.
func (o *VirtualizationBaseVirtualMachine) SetCapacity(v InfraHardwareInfo) {
	o.Capacity = &v
}

// GetGuestInfo returns the GuestInfo field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualMachine) GetGuestInfo() VirtualizationGuestInfo {
	if o == nil || o.GuestInfo == nil {
		var ret VirtualizationGuestInfo
		return ret
	}
	return *o.GuestInfo
}

// GetGuestInfoOk returns a tuple with the GuestInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachine) GetGuestInfoOk() (*VirtualizationGuestInfo, bool) {
	if o == nil || o.GuestInfo == nil {
		return nil, false
	}
	return o.GuestInfo, true
}

// HasGuestInfo returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualMachine) HasGuestInfo() bool {
	if o != nil && o.GuestInfo != nil {
		return true
	}

	return false
}

// SetGuestInfo gets a reference to the given VirtualizationGuestInfo and assigns it to the GuestInfo field.
func (o *VirtualizationBaseVirtualMachine) SetGuestInfo(v VirtualizationGuestInfo) {
	o.GuestInfo = &v
}

// GetHypervisorType returns the HypervisorType field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualMachine) GetHypervisorType() string {
	if o == nil || o.HypervisorType == nil {
		var ret string
		return ret
	}
	return *o.HypervisorType
}

// GetHypervisorTypeOk returns a tuple with the HypervisorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachine) GetHypervisorTypeOk() (*string, bool) {
	if o == nil || o.HypervisorType == nil {
		return nil, false
	}
	return o.HypervisorType, true
}

// HasHypervisorType returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualMachine) HasHypervisorType() bool {
	if o != nil && o.HypervisorType != nil {
		return true
	}

	return false
}

// SetHypervisorType gets a reference to the given string and assigns it to the HypervisorType field.
func (o *VirtualizationBaseVirtualMachine) SetHypervisorType(v string) {
	o.HypervisorType = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualMachine) GetIdentity() string {
	if o == nil || o.Identity == nil {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachine) GetIdentityOk() (*string, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualMachine) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *VirtualizationBaseVirtualMachine) SetIdentity(v string) {
	o.Identity = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualMachine) GetIpAddress() []string {
	if o == nil || o.IpAddress == nil {
		var ret []string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachine) GetIpAddressOk() (*[]string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualMachine) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given []string and assigns it to the IpAddress field.
func (o *VirtualizationBaseVirtualMachine) SetIpAddress(v []string) {
	o.IpAddress = &v
}

// GetMemoryCapacity returns the MemoryCapacity field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualMachine) GetMemoryCapacity() VirtualizationMemoryCapacity {
	if o == nil || o.MemoryCapacity == nil {
		var ret VirtualizationMemoryCapacity
		return ret
	}
	return *o.MemoryCapacity
}

// GetMemoryCapacityOk returns a tuple with the MemoryCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachine) GetMemoryCapacityOk() (*VirtualizationMemoryCapacity, bool) {
	if o == nil || o.MemoryCapacity == nil {
		return nil, false
	}
	return o.MemoryCapacity, true
}

// HasMemoryCapacity returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualMachine) HasMemoryCapacity() bool {
	if o != nil && o.MemoryCapacity != nil {
		return true
	}

	return false
}

// SetMemoryCapacity gets a reference to the given VirtualizationMemoryCapacity and assigns it to the MemoryCapacity field.
func (o *VirtualizationBaseVirtualMachine) SetMemoryCapacity(v VirtualizationMemoryCapacity) {
	o.MemoryCapacity = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualMachine) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachine) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualMachine) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VirtualizationBaseVirtualMachine) SetName(v string) {
	o.Name = &v
}

// GetPowerState returns the PowerState field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualMachine) GetPowerState() string {
	if o == nil || o.PowerState == nil {
		var ret string
		return ret
	}
	return *o.PowerState
}

// GetPowerStateOk returns a tuple with the PowerState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachine) GetPowerStateOk() (*string, bool) {
	if o == nil || o.PowerState == nil {
		return nil, false
	}
	return o.PowerState, true
}

// HasPowerState returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualMachine) HasPowerState() bool {
	if o != nil && o.PowerState != nil {
		return true
	}

	return false
}

// SetPowerState gets a reference to the given string and assigns it to the PowerState field.
func (o *VirtualizationBaseVirtualMachine) SetPowerState(v string) {
	o.PowerState = &v
}

// GetProcessorCapacity returns the ProcessorCapacity field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualMachine) GetProcessorCapacity() VirtualizationComputeCapacity {
	if o == nil || o.ProcessorCapacity == nil {
		var ret VirtualizationComputeCapacity
		return ret
	}
	return *o.ProcessorCapacity
}

// GetProcessorCapacityOk returns a tuple with the ProcessorCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachine) GetProcessorCapacityOk() (*VirtualizationComputeCapacity, bool) {
	if o == nil || o.ProcessorCapacity == nil {
		return nil, false
	}
	return o.ProcessorCapacity, true
}

// HasProcessorCapacity returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualMachine) HasProcessorCapacity() bool {
	if o != nil && o.ProcessorCapacity != nil {
		return true
	}

	return false
}

// SetProcessorCapacity gets a reference to the given VirtualizationComputeCapacity and assigns it to the ProcessorCapacity field.
func (o *VirtualizationBaseVirtualMachine) SetProcessorCapacity(v VirtualizationComputeCapacity) {
	o.ProcessorCapacity = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualMachine) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachine) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualMachine) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *VirtualizationBaseVirtualMachine) SetUuid(v string) {
	o.Uuid = &v
}

func (o VirtualizationBaseVirtualMachine) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseSourceDevice, errVirtualizationBaseSourceDevice := json.Marshal(o.VirtualizationBaseSourceDevice)
	if errVirtualizationBaseSourceDevice != nil {
		return []byte{}, errVirtualizationBaseSourceDevice
	}
	errVirtualizationBaseSourceDevice = json.Unmarshal([]byte(serializedVirtualizationBaseSourceDevice), &toSerialize)
	if errVirtualizationBaseSourceDevice != nil {
		return []byte{}, errVirtualizationBaseSourceDevice
	}
	if o.Capacity != nil {
		toSerialize["Capacity"] = o.Capacity
	}
	if o.GuestInfo != nil {
		toSerialize["GuestInfo"] = o.GuestInfo
	}
	if o.HypervisorType != nil {
		toSerialize["HypervisorType"] = o.HypervisorType
	}
	if o.Identity != nil {
		toSerialize["Identity"] = o.Identity
	}
	if o.IpAddress != nil {
		toSerialize["IpAddress"] = o.IpAddress
	}
	if o.MemoryCapacity != nil {
		toSerialize["MemoryCapacity"] = o.MemoryCapacity
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.PowerState != nil {
		toSerialize["PowerState"] = o.PowerState
	}
	if o.ProcessorCapacity != nil {
		toSerialize["ProcessorCapacity"] = o.ProcessorCapacity
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationBaseVirtualMachine) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationBaseVirtualMachineWithoutEmbeddedStruct struct {
		Capacity  *InfraHardwareInfo       `json:"Capacity,omitempty"`
		GuestInfo *VirtualizationGuestInfo `json:"GuestInfo,omitempty"`
		// Type of hypervisor where the virtual machine is hosted, for example VMware ESXi.
		HypervisorType *string `json:"HypervisorType,omitempty"`
		// The internally generated identity of this VM. This entity is not manipulated by users. It aids in uniquely identifying the virtual machine object. For VMware, this is MOR (managed object reference).
		Identity       *string                       `json:"Identity,omitempty"`
		IpAddress      *[]string                     `json:"IpAddress,omitempty"`
		MemoryCapacity *VirtualizationMemoryCapacity `json:"MemoryCapacity,omitempty"`
		// User-provided name to identify the virtual machine.
		Name *string `json:"Name,omitempty"`
		// Power state of the virtual machine.
		PowerState        *string                        `json:"PowerState,omitempty"`
		ProcessorCapacity *VirtualizationComputeCapacity `json:"ProcessorCapacity,omitempty"`
		// The uuid of this virtual machine. The uuid is internally generated and not user specified.
		Uuid *string `json:"Uuid,omitempty"`
	}

	varVirtualizationBaseVirtualMachineWithoutEmbeddedStruct := VirtualizationBaseVirtualMachineWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationBaseVirtualMachineWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationBaseVirtualMachine := _VirtualizationBaseVirtualMachine{}
		varVirtualizationBaseVirtualMachine.Capacity = varVirtualizationBaseVirtualMachineWithoutEmbeddedStruct.Capacity
		varVirtualizationBaseVirtualMachine.GuestInfo = varVirtualizationBaseVirtualMachineWithoutEmbeddedStruct.GuestInfo
		varVirtualizationBaseVirtualMachine.HypervisorType = varVirtualizationBaseVirtualMachineWithoutEmbeddedStruct.HypervisorType
		varVirtualizationBaseVirtualMachine.Identity = varVirtualizationBaseVirtualMachineWithoutEmbeddedStruct.Identity
		varVirtualizationBaseVirtualMachine.IpAddress = varVirtualizationBaseVirtualMachineWithoutEmbeddedStruct.IpAddress
		varVirtualizationBaseVirtualMachine.MemoryCapacity = varVirtualizationBaseVirtualMachineWithoutEmbeddedStruct.MemoryCapacity
		varVirtualizationBaseVirtualMachine.Name = varVirtualizationBaseVirtualMachineWithoutEmbeddedStruct.Name
		varVirtualizationBaseVirtualMachine.PowerState = varVirtualizationBaseVirtualMachineWithoutEmbeddedStruct.PowerState
		varVirtualizationBaseVirtualMachine.ProcessorCapacity = varVirtualizationBaseVirtualMachineWithoutEmbeddedStruct.ProcessorCapacity
		varVirtualizationBaseVirtualMachine.Uuid = varVirtualizationBaseVirtualMachineWithoutEmbeddedStruct.Uuid
		*o = VirtualizationBaseVirtualMachine(varVirtualizationBaseVirtualMachine)
	} else {
		return err
	}

	varVirtualizationBaseVirtualMachine := _VirtualizationBaseVirtualMachine{}

	err = json.Unmarshal(bytes, &varVirtualizationBaseVirtualMachine)
	if err == nil {
		o.VirtualizationBaseSourceDevice = varVirtualizationBaseVirtualMachine.VirtualizationBaseSourceDevice
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Capacity")
		delete(additionalProperties, "GuestInfo")
		delete(additionalProperties, "HypervisorType")
		delete(additionalProperties, "Identity")
		delete(additionalProperties, "IpAddress")
		delete(additionalProperties, "MemoryCapacity")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "PowerState")
		delete(additionalProperties, "ProcessorCapacity")
		delete(additionalProperties, "Uuid")

		// remove fields from embedded structs
		reflectVirtualizationBaseSourceDevice := reflect.ValueOf(o.VirtualizationBaseSourceDevice)
		for i := 0; i < reflectVirtualizationBaseSourceDevice.Type().NumField(); i++ {
			t := reflectVirtualizationBaseSourceDevice.Type().Field(i)

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

type NullableVirtualizationBaseVirtualMachine struct {
	value *VirtualizationBaseVirtualMachine
	isSet bool
}

func (v NullableVirtualizationBaseVirtualMachine) Get() *VirtualizationBaseVirtualMachine {
	return v.value
}

func (v *NullableVirtualizationBaseVirtualMachine) Set(val *VirtualizationBaseVirtualMachine) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationBaseVirtualMachine) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationBaseVirtualMachine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationBaseVirtualMachine(val *VirtualizationBaseVirtualMachine) *NullableVirtualizationBaseVirtualMachine {
	return &NullableVirtualizationBaseVirtualMachine{value: val, isSet: true}
}

func (v NullableVirtualizationBaseVirtualMachine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationBaseVirtualMachine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
