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

// VirtualizationVirtualMachineDisk Virtual machine disk configuration data.
type VirtualizationVirtualMachineDisk struct {
	MoBaseComplexType
	// Disk bus name given for a virtual machine.
	Bus *string `json:"Bus,omitempty"`
	// Virtual machine network bridge name.
	Name *string `json:"Name,omitempty"`
	// Priority order of the disk.
	Order *int64 `json:"Order,omitempty"`
	// Disk type hdd or cdrom for a virtual machine.
	Type        *string                          `json:"Type,omitempty"`
	VirtualDisk *VirtualizationVirtualDiskConfig `json:"VirtualDisk,omitempty"`
	// Name of the existing virtual disk to be attached to the Virtual Machine.
	VirtualDiskReference *string `json:"VirtualDiskReference,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVirtualMachineDisk VirtualizationVirtualMachineDisk

// NewVirtualizationVirtualMachineDisk instantiates a new VirtualizationVirtualMachineDisk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVirtualMachineDisk() *VirtualizationVirtualMachineDisk {
	this := VirtualizationVirtualMachineDisk{}
	var bus string = "virtio"
	this.Bus = &bus
	var type_ string = "hdd"
	this.Type = &type_
	return &this
}

// NewVirtualizationVirtualMachineDiskWithDefaults instantiates a new VirtualizationVirtualMachineDisk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVirtualMachineDiskWithDefaults() *VirtualizationVirtualMachineDisk {
	this := VirtualizationVirtualMachineDisk{}
	var bus string = "virtio"
	this.Bus = &bus
	var type_ string = "hdd"
	this.Type = &type_
	return &this
}

// GetBus returns the Bus field value if set, zero value otherwise.
func (o *VirtualizationVirtualMachineDisk) GetBus() string {
	if o == nil || o.Bus == nil {
		var ret string
		return ret
	}
	return *o.Bus
}

// GetBusOk returns a tuple with the Bus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualMachineDisk) GetBusOk() (*string, bool) {
	if o == nil || o.Bus == nil {
		return nil, false
	}
	return o.Bus, true
}

// HasBus returns a boolean if a field has been set.
func (o *VirtualizationVirtualMachineDisk) HasBus() bool {
	if o != nil && o.Bus != nil {
		return true
	}

	return false
}

// SetBus gets a reference to the given string and assigns it to the Bus field.
func (o *VirtualizationVirtualMachineDisk) SetBus(v string) {
	o.Bus = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VirtualizationVirtualMachineDisk) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualMachineDisk) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VirtualizationVirtualMachineDisk) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VirtualizationVirtualMachineDisk) SetName(v string) {
	o.Name = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *VirtualizationVirtualMachineDisk) GetOrder() int64 {
	if o == nil || o.Order == nil {
		var ret int64
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualMachineDisk) GetOrderOk() (*int64, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *VirtualizationVirtualMachineDisk) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int64 and assigns it to the Order field.
func (o *VirtualizationVirtualMachineDisk) SetOrder(v int64) {
	o.Order = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VirtualizationVirtualMachineDisk) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualMachineDisk) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VirtualizationVirtualMachineDisk) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *VirtualizationVirtualMachineDisk) SetType(v string) {
	o.Type = &v
}

// GetVirtualDisk returns the VirtualDisk field value if set, zero value otherwise.
func (o *VirtualizationVirtualMachineDisk) GetVirtualDisk() VirtualizationVirtualDiskConfig {
	if o == nil || o.VirtualDisk == nil {
		var ret VirtualizationVirtualDiskConfig
		return ret
	}
	return *o.VirtualDisk
}

// GetVirtualDiskOk returns a tuple with the VirtualDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualMachineDisk) GetVirtualDiskOk() (*VirtualizationVirtualDiskConfig, bool) {
	if o == nil || o.VirtualDisk == nil {
		return nil, false
	}
	return o.VirtualDisk, true
}

// HasVirtualDisk returns a boolean if a field has been set.
func (o *VirtualizationVirtualMachineDisk) HasVirtualDisk() bool {
	if o != nil && o.VirtualDisk != nil {
		return true
	}

	return false
}

// SetVirtualDisk gets a reference to the given VirtualizationVirtualDiskConfig and assigns it to the VirtualDisk field.
func (o *VirtualizationVirtualMachineDisk) SetVirtualDisk(v VirtualizationVirtualDiskConfig) {
	o.VirtualDisk = &v
}

// GetVirtualDiskReference returns the VirtualDiskReference field value if set, zero value otherwise.
func (o *VirtualizationVirtualMachineDisk) GetVirtualDiskReference() string {
	if o == nil || o.VirtualDiskReference == nil {
		var ret string
		return ret
	}
	return *o.VirtualDiskReference
}

// GetVirtualDiskReferenceOk returns a tuple with the VirtualDiskReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualMachineDisk) GetVirtualDiskReferenceOk() (*string, bool) {
	if o == nil || o.VirtualDiskReference == nil {
		return nil, false
	}
	return o.VirtualDiskReference, true
}

// HasVirtualDiskReference returns a boolean if a field has been set.
func (o *VirtualizationVirtualMachineDisk) HasVirtualDiskReference() bool {
	if o != nil && o.VirtualDiskReference != nil {
		return true
	}

	return false
}

// SetVirtualDiskReference gets a reference to the given string and assigns it to the VirtualDiskReference field.
func (o *VirtualizationVirtualMachineDisk) SetVirtualDiskReference(v string) {
	o.VirtualDiskReference = &v
}

func (o VirtualizationVirtualMachineDisk) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.Bus != nil {
		toSerialize["Bus"] = o.Bus
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Order != nil {
		toSerialize["Order"] = o.Order
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.VirtualDisk != nil {
		toSerialize["VirtualDisk"] = o.VirtualDisk
	}
	if o.VirtualDiskReference != nil {
		toSerialize["VirtualDiskReference"] = o.VirtualDiskReference
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationVirtualMachineDisk) UnmarshalJSON(bytes []byte) (err error) {
	varVirtualizationVirtualMachineDisk := _VirtualizationVirtualMachineDisk{}

	if err = json.Unmarshal(bytes, &varVirtualizationVirtualMachineDisk); err == nil {
		*o = VirtualizationVirtualMachineDisk(varVirtualizationVirtualMachineDisk)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Bus")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Order")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "VirtualDisk")
		delete(additionalProperties, "VirtualDiskReference")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualizationVirtualMachineDisk struct {
	value *VirtualizationVirtualMachineDisk
	isSet bool
}

func (v NullableVirtualizationVirtualMachineDisk) Get() *VirtualizationVirtualMachineDisk {
	return v.value
}

func (v *NullableVirtualizationVirtualMachineDisk) Set(val *VirtualizationVirtualMachineDisk) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVirtualMachineDisk) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVirtualMachineDisk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVirtualMachineDisk(val *VirtualizationVirtualMachineDisk) *NullableVirtualizationVirtualMachineDisk {
	return &NullableVirtualizationVirtualMachineDisk{value: val, isSet: true}
}

func (v NullableVirtualizationVirtualMachineDisk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVirtualMachineDisk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
