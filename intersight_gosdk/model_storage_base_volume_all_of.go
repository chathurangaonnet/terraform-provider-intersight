/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-07-11T05:47:33Z.
 *
 * API version: 1.0.9-1999
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
)

// StorageBaseVolumeAllOf Definition of the list of properties defined in 'storage.BaseVolume', excluding properties defined in parent classes.
type StorageBaseVolumeAllOf struct {
	// Short description about the volume.
	Description *string `json:"Description,omitempty"`
	// NAA id of volume. It is a significant number to identify corresponding lun path in hypervisor.
	NaaId *string `json:"NaaId,omitempty"`
	// Named entity of the volume.
	Name *string `json:"Name,omitempty"`
	// User provisioned volume size. It is the size exposed to host.
	Size                 *int64               `json:"Size,omitempty"`
	StorageUtilization   *StorageBaseCapacity `json:"StorageUtilization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageBaseVolumeAllOf StorageBaseVolumeAllOf

// NewStorageBaseVolumeAllOf instantiates a new StorageBaseVolumeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageBaseVolumeAllOf() *StorageBaseVolumeAllOf {
	this := StorageBaseVolumeAllOf{}
	return &this
}

// NewStorageBaseVolumeAllOfWithDefaults instantiates a new StorageBaseVolumeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageBaseVolumeAllOfWithDefaults() *StorageBaseVolumeAllOf {
	this := StorageBaseVolumeAllOf{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *StorageBaseVolumeAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseVolumeAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *StorageBaseVolumeAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *StorageBaseVolumeAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetNaaId returns the NaaId field value if set, zero value otherwise.
func (o *StorageBaseVolumeAllOf) GetNaaId() string {
	if o == nil || o.NaaId == nil {
		var ret string
		return ret
	}
	return *o.NaaId
}

// GetNaaIdOk returns a tuple with the NaaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseVolumeAllOf) GetNaaIdOk() (*string, bool) {
	if o == nil || o.NaaId == nil {
		return nil, false
	}
	return o.NaaId, true
}

// HasNaaId returns a boolean if a field has been set.
func (o *StorageBaseVolumeAllOf) HasNaaId() bool {
	if o != nil && o.NaaId != nil {
		return true
	}

	return false
}

// SetNaaId gets a reference to the given string and assigns it to the NaaId field.
func (o *StorageBaseVolumeAllOf) SetNaaId(v string) {
	o.NaaId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageBaseVolumeAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseVolumeAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageBaseVolumeAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageBaseVolumeAllOf) SetName(v string) {
	o.Name = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *StorageBaseVolumeAllOf) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseVolumeAllOf) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *StorageBaseVolumeAllOf) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *StorageBaseVolumeAllOf) SetSize(v int64) {
	o.Size = &v
}

// GetStorageUtilization returns the StorageUtilization field value if set, zero value otherwise.
func (o *StorageBaseVolumeAllOf) GetStorageUtilization() StorageBaseCapacity {
	if o == nil || o.StorageUtilization == nil {
		var ret StorageBaseCapacity
		return ret
	}
	return *o.StorageUtilization
}

// GetStorageUtilizationOk returns a tuple with the StorageUtilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseVolumeAllOf) GetStorageUtilizationOk() (*StorageBaseCapacity, bool) {
	if o == nil || o.StorageUtilization == nil {
		return nil, false
	}
	return o.StorageUtilization, true
}

// HasStorageUtilization returns a boolean if a field has been set.
func (o *StorageBaseVolumeAllOf) HasStorageUtilization() bool {
	if o != nil && o.StorageUtilization != nil {
		return true
	}

	return false
}

// SetStorageUtilization gets a reference to the given StorageBaseCapacity and assigns it to the StorageUtilization field.
func (o *StorageBaseVolumeAllOf) SetStorageUtilization(v StorageBaseCapacity) {
	o.StorageUtilization = &v
}

func (o StorageBaseVolumeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.NaaId != nil {
		toSerialize["NaaId"] = o.NaaId
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Size != nil {
		toSerialize["Size"] = o.Size
	}
	if o.StorageUtilization != nil {
		toSerialize["StorageUtilization"] = o.StorageUtilization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageBaseVolumeAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageBaseVolumeAllOf := _StorageBaseVolumeAllOf{}

	if err = json.Unmarshal(bytes, &varStorageBaseVolumeAllOf); err == nil {
		*o = StorageBaseVolumeAllOf(varStorageBaseVolumeAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Description")
		delete(additionalProperties, "NaaId")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Size")
		delete(additionalProperties, "StorageUtilization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageBaseVolumeAllOf struct {
	value *StorageBaseVolumeAllOf
	isSet bool
}

func (v NullableStorageBaseVolumeAllOf) Get() *StorageBaseVolumeAllOf {
	return v.value
}

func (v *NullableStorageBaseVolumeAllOf) Set(val *StorageBaseVolumeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageBaseVolumeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageBaseVolumeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageBaseVolumeAllOf(val *StorageBaseVolumeAllOf) *NullableStorageBaseVolumeAllOf {
	return &NullableStorageBaseVolumeAllOf{value: val, isSet: true}
}

func (v NullableStorageBaseVolumeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageBaseVolumeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
