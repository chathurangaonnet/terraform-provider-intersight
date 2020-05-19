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

// AdapterPortChannelSettings Port Channel setting for this adapter.
type AdapterPortChannelSettings struct {
	MoBaseComplexType
	// When Port Channel is enabled, two vNICs and two vHBAs are available for use on the adapter card. When disabled, four vNICs and four vHBAs are available for use on the adapter card. Disabling port channel reboots the server. Port Channel is supported only for Cisco VIC 1455/1457 adapters.
	Enabled *bool `json:"Enabled,omitempty"`
}

// NewAdapterPortChannelSettings instantiates a new AdapterPortChannelSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdapterPortChannelSettings() *AdapterPortChannelSettings {
	this := AdapterPortChannelSettings{}
	return &this
}

// NewAdapterPortChannelSettingsWithDefaults instantiates a new AdapterPortChannelSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdapterPortChannelSettingsWithDefaults() *AdapterPortChannelSettings {
	this := AdapterPortChannelSettings{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AdapterPortChannelSettings) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterPortChannelSettings) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AdapterPortChannelSettings) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AdapterPortChannelSettings) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o AdapterPortChannelSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableAdapterPortChannelSettings struct {
	value *AdapterPortChannelSettings
	isSet bool
}

func (v NullableAdapterPortChannelSettings) Get() *AdapterPortChannelSettings {
	return v.value
}

func (v *NullableAdapterPortChannelSettings) Set(val *AdapterPortChannelSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAdapterPortChannelSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAdapterPortChannelSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdapterPortChannelSettings(val *AdapterPortChannelSettings) *NullableAdapterPortChannelSettings {
	return &NullableAdapterPortChannelSettings{value: val, isSet: true}
}

func (v NullableAdapterPortChannelSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdapterPortChannelSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
