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
	"reflect"
	"strings"
)

// InventoryDeviceInfo Information pertaining to a Registered Device in starship which is pertinent to Inventory Microservice.
type InventoryDeviceInfo struct {
	PolicyinventoryAbstractDeviceInfo
	AdditionalProperties map[string]interface{}
}

type _InventoryDeviceInfo InventoryDeviceInfo

// NewInventoryDeviceInfo instantiates a new InventoryDeviceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryDeviceInfo() *InventoryDeviceInfo {
	this := InventoryDeviceInfo{}
	return &this
}

// NewInventoryDeviceInfoWithDefaults instantiates a new InventoryDeviceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryDeviceInfoWithDefaults() *InventoryDeviceInfo {
	this := InventoryDeviceInfo{}
	return &this
}

func (o InventoryDeviceInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyinventoryAbstractDeviceInfo, errPolicyinventoryAbstractDeviceInfo := json.Marshal(o.PolicyinventoryAbstractDeviceInfo)
	if errPolicyinventoryAbstractDeviceInfo != nil {
		return []byte{}, errPolicyinventoryAbstractDeviceInfo
	}
	errPolicyinventoryAbstractDeviceInfo = json.Unmarshal([]byte(serializedPolicyinventoryAbstractDeviceInfo), &toSerialize)
	if errPolicyinventoryAbstractDeviceInfo != nil {
		return []byte{}, errPolicyinventoryAbstractDeviceInfo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InventoryDeviceInfo) UnmarshalJSON(bytes []byte) (err error) {
	type InventoryDeviceInfoWithoutEmbeddedStruct struct {
	}

	varInventoryDeviceInfoWithoutEmbeddedStruct := InventoryDeviceInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varInventoryDeviceInfoWithoutEmbeddedStruct)
	if err == nil {
		varInventoryDeviceInfo := _InventoryDeviceInfo{}
		*o = InventoryDeviceInfo(varInventoryDeviceInfo)
	} else {
		return err
	}

	varInventoryDeviceInfo := _InventoryDeviceInfo{}

	err = json.Unmarshal(bytes, &varInventoryDeviceInfo)
	if err == nil {
		o.PolicyinventoryAbstractDeviceInfo = varInventoryDeviceInfo.PolicyinventoryAbstractDeviceInfo
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {

		// remove fields from embedded structs
		reflectPolicyinventoryAbstractDeviceInfo := reflect.ValueOf(o.PolicyinventoryAbstractDeviceInfo)
		for i := 0; i < reflectPolicyinventoryAbstractDeviceInfo.Type().NumField(); i++ {
			t := reflectPolicyinventoryAbstractDeviceInfo.Type().Field(i)

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

type NullableInventoryDeviceInfo struct {
	value *InventoryDeviceInfo
	isSet bool
}

func (v NullableInventoryDeviceInfo) Get() *InventoryDeviceInfo {
	return v.value
}

func (v *NullableInventoryDeviceInfo) Set(val *InventoryDeviceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryDeviceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryDeviceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryDeviceInfo(val *InventoryDeviceInfo) *NullableInventoryDeviceInfo {
	return &NullableInventoryDeviceInfo{value: val, isSet: true}
}

func (v NullableInventoryDeviceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryDeviceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
