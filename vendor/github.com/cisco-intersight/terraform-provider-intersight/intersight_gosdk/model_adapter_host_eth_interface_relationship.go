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
	"fmt"
)

// AdapterHostEthInterfaceRelationship - A relationship to the 'adapter.HostEthInterface' resource, or the expanded 'adapter.HostEthInterface' resource, or the 'null' value.
type AdapterHostEthInterfaceRelationship struct {
	AdapterHostEthInterface *AdapterHostEthInterface
	MoMoRef                 *MoMoRef
}

// AdapterHostEthInterfaceAsAdapterHostEthInterfaceRelationship is a convenience function that returns AdapterHostEthInterface wrapped in AdapterHostEthInterfaceRelationship
func AdapterHostEthInterfaceAsAdapterHostEthInterfaceRelationship(v *AdapterHostEthInterface) AdapterHostEthInterfaceRelationship {
	return AdapterHostEthInterfaceRelationship{AdapterHostEthInterface: v}
}

// MoMoRefAsAdapterHostEthInterfaceRelationship is a convenience function that returns MoMoRef wrapped in AdapterHostEthInterfaceRelationship
func MoMoRefAsAdapterHostEthInterfaceRelationship(v *MoMoRef) AdapterHostEthInterfaceRelationship {
	return AdapterHostEthInterfaceRelationship{MoMoRef: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AdapterHostEthInterfaceRelationship) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discrimintor lookup.")
	}

	// check if the discriminator value is 'adapter.HostEthInterface'
	if jsonDict["ClassId"] == "adapter.HostEthInterface" {
		// try to unmarshal JSON data into AdapterHostEthInterface
		err = json.Unmarshal(data, &dst.AdapterHostEthInterface)
		if err == nil {
			jsonAdapterHostEthInterface, _ := json.Marshal(dst.AdapterHostEthInterface)
			if string(jsonAdapterHostEthInterface) == "{}" { // empty struct
				dst.AdapterHostEthInterface = nil
			} else {
				return nil // data stored in dst.AdapterHostEthInterface, return on the first match
			}
		} else {
			dst.AdapterHostEthInterface = nil
		}
	}

	// check if the discriminator value is 'mo.MoRef'
	if jsonDict["ClassId"] == "mo.MoRef" {
		// try to unmarshal JSON data into MoMoRef
		err = json.Unmarshal(data, &dst.MoMoRef)
		if err == nil {
			jsonMoMoRef, _ := json.Marshal(dst.MoMoRef)
			if string(jsonMoMoRef) == "{}" { // empty struct
				dst.MoMoRef = nil
			} else {
				return nil // data stored in dst.MoMoRef, return on the first match
			}
		} else {
			dst.MoMoRef = nil
		}
	}

	// try to unmarshal data into AdapterHostEthInterface
	err = json.Unmarshal(data, &dst.AdapterHostEthInterface)
	if err == nil {
		jsonAdapterHostEthInterface, _ := json.Marshal(dst.AdapterHostEthInterface)
		if string(jsonAdapterHostEthInterface) == "{}" { // empty struct
			dst.AdapterHostEthInterface = nil
		} else {
			match++
		}
	} else {
		dst.AdapterHostEthInterface = nil
	}

	// try to unmarshal data into MoMoRef
	err = json.Unmarshal(data, &dst.MoMoRef)
	if err == nil {
		jsonMoMoRef, _ := json.Marshal(dst.MoMoRef)
		if string(jsonMoMoRef) == "{}" { // empty struct
			dst.MoMoRef = nil
		} else {
			match++
		}
	} else {
		dst.MoMoRef = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AdapterHostEthInterface = nil
		dst.MoMoRef = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(AdapterHostEthInterfaceRelationship)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(AdapterHostEthInterfaceRelationship)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AdapterHostEthInterfaceRelationship) MarshalJSON() ([]byte, error) {
	if src.AdapterHostEthInterface != nil {
		return json.Marshal(&src.AdapterHostEthInterface)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AdapterHostEthInterfaceRelationship) GetActualInstance() interface{} {
	if obj.AdapterHostEthInterface != nil {
		return obj.AdapterHostEthInterface
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableAdapterHostEthInterfaceRelationship struct {
	value *AdapterHostEthInterfaceRelationship
	isSet bool
}

func (v NullableAdapterHostEthInterfaceRelationship) Get() *AdapterHostEthInterfaceRelationship {
	return v.value
}

func (v *NullableAdapterHostEthInterfaceRelationship) Set(val *AdapterHostEthInterfaceRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableAdapterHostEthInterfaceRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableAdapterHostEthInterfaceRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdapterHostEthInterfaceRelationship(val *AdapterHostEthInterfaceRelationship) *NullableAdapterHostEthInterfaceRelationship {
	return &NullableAdapterHostEthInterfaceRelationship{value: val, isSet: true}
}

func (v NullableAdapterHostEthInterfaceRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdapterHostEthInterfaceRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
