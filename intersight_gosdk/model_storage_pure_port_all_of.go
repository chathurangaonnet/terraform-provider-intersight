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

// StoragePurePortAllOf Definition of the list of properties defined in 'storage.PurePort', excluding properties defined in parent classes.
type StoragePurePortAllOf struct {
	// Name of the port to which this port has failed over.
	Failover *string `json:"Failover,omitempty"`
	// Ip address of iSCSI portal configured on the port.
	Portal           *string                              `json:"Portal,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
}

// NewStoragePurePortAllOf instantiates a new StoragePurePortAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoragePurePortAllOf() *StoragePurePortAllOf {
	this := StoragePurePortAllOf{}
	return &this
}

// NewStoragePurePortAllOfWithDefaults instantiates a new StoragePurePortAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoragePurePortAllOfWithDefaults() *StoragePurePortAllOf {
	this := StoragePurePortAllOf{}
	return &this
}

// GetFailover returns the Failover field value if set, zero value otherwise.
func (o *StoragePurePortAllOf) GetFailover() string {
	if o == nil || o.Failover == nil {
		var ret string
		return ret
	}
	return *o.Failover
}

// GetFailoverOk returns a tuple with the Failover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePurePortAllOf) GetFailoverOk() (*string, bool) {
	if o == nil || o.Failover == nil {
		return nil, false
	}
	return o.Failover, true
}

// HasFailover returns a boolean if a field has been set.
func (o *StoragePurePortAllOf) HasFailover() bool {
	if o != nil && o.Failover != nil {
		return true
	}

	return false
}

// SetFailover gets a reference to the given string and assigns it to the Failover field.
func (o *StoragePurePortAllOf) SetFailover(v string) {
	o.Failover = &v
}

// GetPortal returns the Portal field value if set, zero value otherwise.
func (o *StoragePurePortAllOf) GetPortal() string {
	if o == nil || o.Portal == nil {
		var ret string
		return ret
	}
	return *o.Portal
}

// GetPortalOk returns a tuple with the Portal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePurePortAllOf) GetPortalOk() (*string, bool) {
	if o == nil || o.Portal == nil {
		return nil, false
	}
	return o.Portal, true
}

// HasPortal returns a boolean if a field has been set.
func (o *StoragePurePortAllOf) HasPortal() bool {
	if o != nil && o.Portal != nil {
		return true
	}

	return false
}

// SetPortal gets a reference to the given string and assigns it to the Portal field.
func (o *StoragePurePortAllOf) SetPortal(v string) {
	o.Portal = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StoragePurePortAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePurePortAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StoragePurePortAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StoragePurePortAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o StoragePurePortAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

type NullableStoragePurePortAllOf struct {
	value *StoragePurePortAllOf
	isSet bool
}

func (v NullableStoragePurePortAllOf) Get() *StoragePurePortAllOf {
	return v.value
}

func (v *NullableStoragePurePortAllOf) Set(val *StoragePurePortAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStoragePurePortAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStoragePurePortAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoragePurePortAllOf(val *StoragePurePortAllOf) *NullableStoragePurePortAllOf {
	return &NullableStoragePurePortAllOf{value: val, isSet: true}
}

func (v NullableStoragePurePortAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoragePurePortAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
