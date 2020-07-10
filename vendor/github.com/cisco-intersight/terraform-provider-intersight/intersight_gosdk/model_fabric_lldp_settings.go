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

// FabricLldpSettings The LLDP settings on an interface on the switch.
type FabricLldpSettings struct {
	MoBaseComplexType
	// Determines if the LLDP frames can be received by an interface on the switch.
	ReceiveEnabled *bool `json:"ReceiveEnabled,omitempty"`
	// Determines if the LLDP frames can be transmitted by an interface on the switch.
	TransmitEnabled      *bool `json:"TransmitEnabled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricLldpSettings FabricLldpSettings

// NewFabricLldpSettings instantiates a new FabricLldpSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricLldpSettings() *FabricLldpSettings {
	this := FabricLldpSettings{}
	return &this
}

// NewFabricLldpSettingsWithDefaults instantiates a new FabricLldpSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricLldpSettingsWithDefaults() *FabricLldpSettings {
	this := FabricLldpSettings{}
	return &this
}

// GetReceiveEnabled returns the ReceiveEnabled field value if set, zero value otherwise.
func (o *FabricLldpSettings) GetReceiveEnabled() bool {
	if o == nil || o.ReceiveEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ReceiveEnabled
}

// GetReceiveEnabledOk returns a tuple with the ReceiveEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricLldpSettings) GetReceiveEnabledOk() (*bool, bool) {
	if o == nil || o.ReceiveEnabled == nil {
		return nil, false
	}
	return o.ReceiveEnabled, true
}

// HasReceiveEnabled returns a boolean if a field has been set.
func (o *FabricLldpSettings) HasReceiveEnabled() bool {
	if o != nil && o.ReceiveEnabled != nil {
		return true
	}

	return false
}

// SetReceiveEnabled gets a reference to the given bool and assigns it to the ReceiveEnabled field.
func (o *FabricLldpSettings) SetReceiveEnabled(v bool) {
	o.ReceiveEnabled = &v
}

// GetTransmitEnabled returns the TransmitEnabled field value if set, zero value otherwise.
func (o *FabricLldpSettings) GetTransmitEnabled() bool {
	if o == nil || o.TransmitEnabled == nil {
		var ret bool
		return ret
	}
	return *o.TransmitEnabled
}

// GetTransmitEnabledOk returns a tuple with the TransmitEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricLldpSettings) GetTransmitEnabledOk() (*bool, bool) {
	if o == nil || o.TransmitEnabled == nil {
		return nil, false
	}
	return o.TransmitEnabled, true
}

// HasTransmitEnabled returns a boolean if a field has been set.
func (o *FabricLldpSettings) HasTransmitEnabled() bool {
	if o != nil && o.TransmitEnabled != nil {
		return true
	}

	return false
}

// SetTransmitEnabled gets a reference to the given bool and assigns it to the TransmitEnabled field.
func (o *FabricLldpSettings) SetTransmitEnabled(v bool) {
	o.TransmitEnabled = &v
}

func (o FabricLldpSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.ReceiveEnabled != nil {
		toSerialize["ReceiveEnabled"] = o.ReceiveEnabled
	}
	if o.TransmitEnabled != nil {
		toSerialize["TransmitEnabled"] = o.TransmitEnabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricLldpSettings) UnmarshalJSON(bytes []byte) (err error) {
	varFabricLldpSettings := _FabricLldpSettings{}

	if err = json.Unmarshal(bytes, &varFabricLldpSettings); err == nil {
		*o = FabricLldpSettings(varFabricLldpSettings)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ReceiveEnabled")
		delete(additionalProperties, "TransmitEnabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFabricLldpSettings struct {
	value *FabricLldpSettings
	isSet bool
}

func (v NullableFabricLldpSettings) Get() *FabricLldpSettings {
	return v.value
}

func (v *NullableFabricLldpSettings) Set(val *FabricLldpSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricLldpSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricLldpSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricLldpSettings(val *FabricLldpSettings) *NullableFabricLldpSettings {
	return &NullableFabricLldpSettings{value: val, isSet: true}
}

func (v NullableFabricLldpSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricLldpSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
