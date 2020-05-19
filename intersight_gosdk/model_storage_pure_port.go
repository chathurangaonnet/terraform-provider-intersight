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

// StoragePurePort Port entity in Pure FlashArray.
type StoragePurePort struct {
	StoragePhysicalPort
	// Name of the port to which this port has failed over.
	Failover *string `json:"Failover,omitempty"`
	// Ip address of iSCSI portal configured on the port.
	Portal           *string                              `json:"Portal,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
}

// NewStoragePurePort instantiates a new StoragePurePort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoragePurePort() *StoragePurePort {
	this := StoragePurePort{}
	return &this
}

// NewStoragePurePortWithDefaults instantiates a new StoragePurePort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoragePurePortWithDefaults() *StoragePurePort {
	this := StoragePurePort{}
	return &this
}

// GetFailover returns the Failover field value if set, zero value otherwise.
func (o *StoragePurePort) GetFailover() string {
	if o == nil || o.Failover == nil {
		var ret string
		return ret
	}
	return *o.Failover
}

// GetFailoverOk returns a tuple with the Failover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePurePort) GetFailoverOk() (*string, bool) {
	if o == nil || o.Failover == nil {
		return nil, false
	}
	return o.Failover, true
}

// HasFailover returns a boolean if a field has been set.
func (o *StoragePurePort) HasFailover() bool {
	if o != nil && o.Failover != nil {
		return true
	}

	return false
}

// SetFailover gets a reference to the given string and assigns it to the Failover field.
func (o *StoragePurePort) SetFailover(v string) {
	o.Failover = &v
}

// GetPortal returns the Portal field value if set, zero value otherwise.
func (o *StoragePurePort) GetPortal() string {
	if o == nil || o.Portal == nil {
		var ret string
		return ret
	}
	return *o.Portal
}

// GetPortalOk returns a tuple with the Portal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePurePort) GetPortalOk() (*string, bool) {
	if o == nil || o.Portal == nil {
		return nil, false
	}
	return o.Portal, true
}

// HasPortal returns a boolean if a field has been set.
func (o *StoragePurePort) HasPortal() bool {
	if o != nil && o.Portal != nil {
		return true
	}

	return false
}

// SetPortal gets a reference to the given string and assigns it to the Portal field.
func (o *StoragePurePort) SetPortal(v string) {
	o.Portal = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StoragePurePort) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePurePort) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StoragePurePort) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StoragePurePort) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o StoragePurePort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedStoragePhysicalPort, errStoragePhysicalPort := json.Marshal(o.StoragePhysicalPort)
	if errStoragePhysicalPort != nil {
		return []byte{}, errStoragePhysicalPort
	}
	errStoragePhysicalPort = json.Unmarshal([]byte(serializedStoragePhysicalPort), &toSerialize)
	if errStoragePhysicalPort != nil {
		return []byte{}, errStoragePhysicalPort
	}
	if o.Failover != nil {
		toSerialize["Failover"] = o.Failover
	}
	if o.Portal != nil {
		toSerialize["Portal"] = o.Portal
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	return json.Marshal(toSerialize)
}

type NullableStoragePurePort struct {
	value *StoragePurePort
	isSet bool
}

func (v NullableStoragePurePort) Get() *StoragePurePort {
	return v.value
}

func (v *NullableStoragePurePort) Set(val *StoragePurePort) {
	v.value = val
	v.isSet = true
}

func (v NullableStoragePurePort) IsSet() bool {
	return v.isSet
}

func (v *NullableStoragePurePort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoragePurePort(val *StoragePurePort) *NullableStoragePurePort {
	return &NullableStoragePurePort{value: val, isSet: true}
}

func (v NullableStoragePurePort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoragePurePort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
