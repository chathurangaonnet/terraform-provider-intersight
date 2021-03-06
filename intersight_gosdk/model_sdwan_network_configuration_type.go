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

// SdwanNetworkConfigurationType struct for SdwanNetworkConfigurationType
type SdwanNetworkConfigurationType struct {
	MoBaseComplexType
	// Type of the Port group being added.
	NetworkType *string `json:"NetworkType,omitempty"`
	// Name of the Port Group to create.
	PortGroup *string `json:"PortGroup,omitempty"`
	// VLAN to be added to the Port Group.
	Vlan                 *int64 `json:"Vlan,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SdwanNetworkConfigurationType SdwanNetworkConfigurationType

// NewSdwanNetworkConfigurationType instantiates a new SdwanNetworkConfigurationType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSdwanNetworkConfigurationType() *SdwanNetworkConfigurationType {
	this := SdwanNetworkConfigurationType{}
	var networkType string = "WAN"
	this.NetworkType = &networkType
	return &this
}

// NewSdwanNetworkConfigurationTypeWithDefaults instantiates a new SdwanNetworkConfigurationType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSdwanNetworkConfigurationTypeWithDefaults() *SdwanNetworkConfigurationType {
	this := SdwanNetworkConfigurationType{}
	var networkType string = "WAN"
	this.NetworkType = &networkType
	return &this
}

// GetNetworkType returns the NetworkType field value if set, zero value otherwise.
func (o *SdwanNetworkConfigurationType) GetNetworkType() string {
	if o == nil || o.NetworkType == nil {
		var ret string
		return ret
	}
	return *o.NetworkType
}

// GetNetworkTypeOk returns a tuple with the NetworkType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdwanNetworkConfigurationType) GetNetworkTypeOk() (*string, bool) {
	if o == nil || o.NetworkType == nil {
		return nil, false
	}
	return o.NetworkType, true
}

// HasNetworkType returns a boolean if a field has been set.
func (o *SdwanNetworkConfigurationType) HasNetworkType() bool {
	if o != nil && o.NetworkType != nil {
		return true
	}

	return false
}

// SetNetworkType gets a reference to the given string and assigns it to the NetworkType field.
func (o *SdwanNetworkConfigurationType) SetNetworkType(v string) {
	o.NetworkType = &v
}

// GetPortGroup returns the PortGroup field value if set, zero value otherwise.
func (o *SdwanNetworkConfigurationType) GetPortGroup() string {
	if o == nil || o.PortGroup == nil {
		var ret string
		return ret
	}
	return *o.PortGroup
}

// GetPortGroupOk returns a tuple with the PortGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdwanNetworkConfigurationType) GetPortGroupOk() (*string, bool) {
	if o == nil || o.PortGroup == nil {
		return nil, false
	}
	return o.PortGroup, true
}

// HasPortGroup returns a boolean if a field has been set.
func (o *SdwanNetworkConfigurationType) HasPortGroup() bool {
	if o != nil && o.PortGroup != nil {
		return true
	}

	return false
}

// SetPortGroup gets a reference to the given string and assigns it to the PortGroup field.
func (o *SdwanNetworkConfigurationType) SetPortGroup(v string) {
	o.PortGroup = &v
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *SdwanNetworkConfigurationType) GetVlan() int64 {
	if o == nil || o.Vlan == nil {
		var ret int64
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdwanNetworkConfigurationType) GetVlanOk() (*int64, bool) {
	if o == nil || o.Vlan == nil {
		return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *SdwanNetworkConfigurationType) HasVlan() bool {
	if o != nil && o.Vlan != nil {
		return true
	}

	return false
}

// SetVlan gets a reference to the given int64 and assigns it to the Vlan field.
func (o *SdwanNetworkConfigurationType) SetVlan(v int64) {
	o.Vlan = &v
}

func (o SdwanNetworkConfigurationType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.NetworkType != nil {
		toSerialize["NetworkType"] = o.NetworkType
	}
	if o.PortGroup != nil {
		toSerialize["PortGroup"] = o.PortGroup
	}
	if o.Vlan != nil {
		toSerialize["Vlan"] = o.Vlan
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SdwanNetworkConfigurationType) UnmarshalJSON(bytes []byte) (err error) {
	type SdwanNetworkConfigurationTypeWithoutEmbeddedStruct struct {
		// Type of the Port group being added.
		NetworkType *string `json:"NetworkType,omitempty"`
		// Name of the Port Group to create.
		PortGroup *string `json:"PortGroup,omitempty"`
		// VLAN to be added to the Port Group.
		Vlan *int64 `json:"Vlan,omitempty"`
	}

	varSdwanNetworkConfigurationTypeWithoutEmbeddedStruct := SdwanNetworkConfigurationTypeWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varSdwanNetworkConfigurationTypeWithoutEmbeddedStruct)
	if err == nil {
		varSdwanNetworkConfigurationType := _SdwanNetworkConfigurationType{}
		varSdwanNetworkConfigurationType.NetworkType = varSdwanNetworkConfigurationTypeWithoutEmbeddedStruct.NetworkType
		varSdwanNetworkConfigurationType.PortGroup = varSdwanNetworkConfigurationTypeWithoutEmbeddedStruct.PortGroup
		varSdwanNetworkConfigurationType.Vlan = varSdwanNetworkConfigurationTypeWithoutEmbeddedStruct.Vlan
		*o = SdwanNetworkConfigurationType(varSdwanNetworkConfigurationType)
	} else {
		return err
	}

	varSdwanNetworkConfigurationType := _SdwanNetworkConfigurationType{}

	err = json.Unmarshal(bytes, &varSdwanNetworkConfigurationType)
	if err == nil {
		o.MoBaseComplexType = varSdwanNetworkConfigurationType.MoBaseComplexType
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "NetworkType")
		delete(additionalProperties, "PortGroup")
		delete(additionalProperties, "Vlan")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableSdwanNetworkConfigurationType struct {
	value *SdwanNetworkConfigurationType
	isSet bool
}

func (v NullableSdwanNetworkConfigurationType) Get() *SdwanNetworkConfigurationType {
	return v.value
}

func (v *NullableSdwanNetworkConfigurationType) Set(val *SdwanNetworkConfigurationType) {
	v.value = val
	v.isSet = true
}

func (v NullableSdwanNetworkConfigurationType) IsSet() bool {
	return v.isSet
}

func (v *NullableSdwanNetworkConfigurationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSdwanNetworkConfigurationType(val *SdwanNetworkConfigurationType) *NullableSdwanNetworkConfigurationType {
	return &NullableSdwanNetworkConfigurationType{value: val, isSet: true}
}

func (v NullableSdwanNetworkConfigurationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSdwanNetworkConfigurationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
