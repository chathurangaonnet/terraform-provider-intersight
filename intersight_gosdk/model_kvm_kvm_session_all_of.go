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

// KvmKvmSessionAllOf Definition of the list of properties defined in 'kvm.KvmSession', excluding properties defined in parent classes.
type KvmKvmSessionAllOf struct {
	// The URL of the KVM MUX to connect to.
	KvmMuxUrl *string                              `json:"KvmMuxUrl,omitempty"`
	Device    *AssetDeviceRegistrationRelationship `json:"Device,omitempty"`
	Server    *ComputeRackUnitRelationship         `json:"Server,omitempty"`
}

// NewKvmKvmSessionAllOf instantiates a new KvmKvmSessionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKvmKvmSessionAllOf() *KvmKvmSessionAllOf {
	this := KvmKvmSessionAllOf{}
	return &this
}

// NewKvmKvmSessionAllOfWithDefaults instantiates a new KvmKvmSessionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKvmKvmSessionAllOfWithDefaults() *KvmKvmSessionAllOf {
	this := KvmKvmSessionAllOf{}
	return &this
}

// GetKvmMuxUrl returns the KvmMuxUrl field value if set, zero value otherwise.
func (o *KvmKvmSessionAllOf) GetKvmMuxUrl() string {
	if o == nil || o.KvmMuxUrl == nil {
		var ret string
		return ret
	}
	return *o.KvmMuxUrl
}

// GetKvmMuxUrlOk returns a tuple with the KvmMuxUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmKvmSessionAllOf) GetKvmMuxUrlOk() (*string, bool) {
	if o == nil || o.KvmMuxUrl == nil {
		return nil, false
	}
	return o.KvmMuxUrl, true
}

// HasKvmMuxUrl returns a boolean if a field has been set.
func (o *KvmKvmSessionAllOf) HasKvmMuxUrl() bool {
	if o != nil && o.KvmMuxUrl != nil {
		return true
	}

	return false
}

// SetKvmMuxUrl gets a reference to the given string and assigns it to the KvmMuxUrl field.
func (o *KvmKvmSessionAllOf) SetKvmMuxUrl(v string) {
	o.KvmMuxUrl = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *KvmKvmSessionAllOf) GetDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.Device == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmKvmSessionAllOf) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *KvmKvmSessionAllOf) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the Device field.
func (o *KvmKvmSessionAllOf) SetDevice(v AssetDeviceRegistrationRelationship) {
	o.Device = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *KvmKvmSessionAllOf) GetServer() ComputeRackUnitRelationship {
	if o == nil || o.Server == nil {
		var ret ComputeRackUnitRelationship
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmKvmSessionAllOf) GetServerOk() (*ComputeRackUnitRelationship, bool) {
	if o == nil || o.Server == nil {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *KvmKvmSessionAllOf) HasServer() bool {
	if o != nil && o.Server != nil {
		return true
	}

	return false
}

// SetServer gets a reference to the given ComputeRackUnitRelationship and assigns it to the Server field.
func (o *KvmKvmSessionAllOf) SetServer(v ComputeRackUnitRelationship) {
	o.Server = &v
}

func (o KvmKvmSessionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.KvmMuxUrl != nil {
		toSerialize["KvmMuxUrl"] = o.KvmMuxUrl
	}
	if o.Device != nil {
		toSerialize["Device"] = o.Device
	}
	if o.Server != nil {
		toSerialize["Server"] = o.Server
	}
	return json.Marshal(toSerialize)
}

type NullableKvmKvmSessionAllOf struct {
	value *KvmKvmSessionAllOf
	isSet bool
}

func (v NullableKvmKvmSessionAllOf) Get() *KvmKvmSessionAllOf {
	return v.value
}

func (v *NullableKvmKvmSessionAllOf) Set(val *KvmKvmSessionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKvmKvmSessionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKvmKvmSessionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKvmKvmSessionAllOf(val *KvmKvmSessionAllOf) *NullableKvmKvmSessionAllOf {
	return &NullableKvmKvmSessionAllOf{value: val, isSet: true}
}

func (v NullableKvmKvmSessionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKvmKvmSessionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
