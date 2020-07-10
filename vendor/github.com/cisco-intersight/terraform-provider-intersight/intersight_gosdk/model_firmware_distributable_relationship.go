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

// FirmwareDistributableRelationship - A relationship to the 'firmware.Distributable' resource, or the expanded 'firmware.Distributable' resource, or the 'null' value.
type FirmwareDistributableRelationship struct {
	FirmwareDistributable *FirmwareDistributable
	MoMoRef               *MoMoRef
}

// FirmwareDistributableAsFirmwareDistributableRelationship is a convenience function that returns FirmwareDistributable wrapped in FirmwareDistributableRelationship
func FirmwareDistributableAsFirmwareDistributableRelationship(v *FirmwareDistributable) FirmwareDistributableRelationship {
	return FirmwareDistributableRelationship{FirmwareDistributable: v}
}

// MoMoRefAsFirmwareDistributableRelationship is a convenience function that returns MoMoRef wrapped in FirmwareDistributableRelationship
func MoMoRefAsFirmwareDistributableRelationship(v *MoMoRef) FirmwareDistributableRelationship {
	return FirmwareDistributableRelationship{MoMoRef: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *FirmwareDistributableRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'firmware.Distributable'
	if jsonDict["ClassId"] == "firmware.Distributable" {
		// try to unmarshal JSON data into FirmwareDistributable
		err = json.Unmarshal(data, &dst.FirmwareDistributable)
		if err == nil {
			jsonFirmwareDistributable, _ := json.Marshal(dst.FirmwareDistributable)
			if string(jsonFirmwareDistributable) == "{}" { // empty struct
				dst.FirmwareDistributable = nil
			} else {
				return nil // data stored in dst.FirmwareDistributable, return on the first match
			}
		} else {
			dst.FirmwareDistributable = nil
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

	// try to unmarshal data into FirmwareDistributable
	err = json.Unmarshal(data, &dst.FirmwareDistributable)
	if err == nil {
		jsonFirmwareDistributable, _ := json.Marshal(dst.FirmwareDistributable)
		if string(jsonFirmwareDistributable) == "{}" { // empty struct
			dst.FirmwareDistributable = nil
		} else {
			match++
		}
	} else {
		dst.FirmwareDistributable = nil
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
		dst.FirmwareDistributable = nil
		dst.MoMoRef = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(FirmwareDistributableRelationship)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(FirmwareDistributableRelationship)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src FirmwareDistributableRelationship) MarshalJSON() ([]byte, error) {
	if src.FirmwareDistributable != nil {
		return json.Marshal(&src.FirmwareDistributable)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *FirmwareDistributableRelationship) GetActualInstance() interface{} {
	if obj.FirmwareDistributable != nil {
		return obj.FirmwareDistributable
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableFirmwareDistributableRelationship struct {
	value *FirmwareDistributableRelationship
	isSet bool
}

func (v NullableFirmwareDistributableRelationship) Get() *FirmwareDistributableRelationship {
	return v.value
}

func (v *NullableFirmwareDistributableRelationship) Set(val *FirmwareDistributableRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareDistributableRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareDistributableRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareDistributableRelationship(val *FirmwareDistributableRelationship) *NullableFirmwareDistributableRelationship {
	return &NullableFirmwareDistributableRelationship{value: val, isSet: true}
}

func (v NullableFirmwareDistributableRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareDistributableRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
