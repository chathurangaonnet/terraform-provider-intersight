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
)

// VnicFcInterruptSettingsAllOf Definition of the list of properties defined in 'vnic.FcInterruptSettings', excluding properties defined in parent classes.
type VnicFcInterruptSettingsAllOf struct {
	// The preferred driver interrupt mode. This can be one of the following:- MSIx — Message Signaled Interrupts (MSI) with the optional extension. MSI   — MSI only. INTx  — PCI INTx interrupts. MSIx is the recommended option.
	Mode                 *string `json:"Mode,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicFcInterruptSettingsAllOf VnicFcInterruptSettingsAllOf

// NewVnicFcInterruptSettingsAllOf instantiates a new VnicFcInterruptSettingsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicFcInterruptSettingsAllOf() *VnicFcInterruptSettingsAllOf {
	this := VnicFcInterruptSettingsAllOf{}
	var mode string = "MSIx"
	this.Mode = &mode
	return &this
}

// NewVnicFcInterruptSettingsAllOfWithDefaults instantiates a new VnicFcInterruptSettingsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicFcInterruptSettingsAllOfWithDefaults() *VnicFcInterruptSettingsAllOf {
	this := VnicFcInterruptSettingsAllOf{}
	var mode string = "MSIx"
	this.Mode = &mode
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *VnicFcInterruptSettingsAllOf) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcInterruptSettingsAllOf) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *VnicFcInterruptSettingsAllOf) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *VnicFcInterruptSettingsAllOf) SetMode(v string) {
	o.Mode = &v
}

func (o VnicFcInterruptSettingsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Mode != nil {
		toSerialize["Mode"] = o.Mode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicFcInterruptSettingsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVnicFcInterruptSettingsAllOf := _VnicFcInterruptSettingsAllOf{}

	if err = json.Unmarshal(bytes, &varVnicFcInterruptSettingsAllOf); err == nil {
		*o = VnicFcInterruptSettingsAllOf(varVnicFcInterruptSettingsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Mode")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVnicFcInterruptSettingsAllOf struct {
	value *VnicFcInterruptSettingsAllOf
	isSet bool
}

func (v NullableVnicFcInterruptSettingsAllOf) Get() *VnicFcInterruptSettingsAllOf {
	return v.value
}

func (v *NullableVnicFcInterruptSettingsAllOf) Set(val *VnicFcInterruptSettingsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicFcInterruptSettingsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicFcInterruptSettingsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicFcInterruptSettingsAllOf(val *VnicFcInterruptSettingsAllOf) *NullableVnicFcInterruptSettingsAllOf {
	return &NullableVnicFcInterruptSettingsAllOf{value: val, isSet: true}
}

func (v NullableVnicFcInterruptSettingsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicFcInterruptSettingsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
