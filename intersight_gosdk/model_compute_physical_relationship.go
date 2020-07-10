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

// ComputePhysicalRelationship - A relationship to the 'compute.Physical' resource, or the expanded 'compute.Physical' resource, or the 'null' value.
type ComputePhysicalRelationship struct {
	ComputePhysical *ComputePhysical
	MoMoRef         *MoMoRef
}

// ComputePhysicalAsComputePhysicalRelationship is a convenience function that returns ComputePhysical wrapped in ComputePhysicalRelationship
func ComputePhysicalAsComputePhysicalRelationship(v *ComputePhysical) ComputePhysicalRelationship {
	return ComputePhysicalRelationship{ComputePhysical: v}
}

// MoMoRefAsComputePhysicalRelationship is a convenience function that returns MoMoRef wrapped in ComputePhysicalRelationship
func MoMoRefAsComputePhysicalRelationship(v *MoMoRef) ComputePhysicalRelationship {
	return ComputePhysicalRelationship{MoMoRef: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ComputePhysicalRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'compute.Physical'
	if jsonDict["ClassId"] == "compute.Physical" {
		// try to unmarshal JSON data into ComputePhysical
		err = json.Unmarshal(data, &dst.ComputePhysical)
		if err == nil {
			jsonComputePhysical, _ := json.Marshal(dst.ComputePhysical)
			if string(jsonComputePhysical) == "{}" { // empty struct
				dst.ComputePhysical = nil
			} else {
				return nil // data stored in dst.ComputePhysical, return on the first match
			}
		} else {
			dst.ComputePhysical = nil
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

	// try to unmarshal data into ComputePhysical
	err = json.Unmarshal(data, &dst.ComputePhysical)
	if err == nil {
		jsonComputePhysical, _ := json.Marshal(dst.ComputePhysical)
		if string(jsonComputePhysical) == "{}" { // empty struct
			dst.ComputePhysical = nil
		} else {
			match++
		}
	} else {
		dst.ComputePhysical = nil
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
		dst.ComputePhysical = nil
		dst.MoMoRef = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(ComputePhysicalRelationship)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(ComputePhysicalRelationship)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ComputePhysicalRelationship) MarshalJSON() ([]byte, error) {
	if src.ComputePhysical != nil {
		return json.Marshal(&src.ComputePhysical)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ComputePhysicalRelationship) GetActualInstance() interface{} {
	if obj.ComputePhysical != nil {
		return obj.ComputePhysical
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableComputePhysicalRelationship struct {
	value *ComputePhysicalRelationship
	isSet bool
}

func (v NullableComputePhysicalRelationship) Get() *ComputePhysicalRelationship {
	return v.value
}

func (v *NullableComputePhysicalRelationship) Set(val *ComputePhysicalRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableComputePhysicalRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableComputePhysicalRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputePhysicalRelationship(val *ComputePhysicalRelationship) *NullableComputePhysicalRelationship {
	return &NullableComputePhysicalRelationship{value: val, isSet: true}
}

func (v NullableComputePhysicalRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputePhysicalRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
