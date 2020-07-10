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

// IamQualifierRelationship - A relationship to the 'iam.Qualifier' resource, or the expanded 'iam.Qualifier' resource, or the 'null' value.
type IamQualifierRelationship struct {
	IamQualifier *IamQualifier
	MoMoRef      *MoMoRef
}

// IamQualifierAsIamQualifierRelationship is a convenience function that returns IamQualifier wrapped in IamQualifierRelationship
func IamQualifierAsIamQualifierRelationship(v *IamQualifier) IamQualifierRelationship {
	return IamQualifierRelationship{IamQualifier: v}
}

// MoMoRefAsIamQualifierRelationship is a convenience function that returns MoMoRef wrapped in IamQualifierRelationship
func MoMoRefAsIamQualifierRelationship(v *MoMoRef) IamQualifierRelationship {
	return IamQualifierRelationship{MoMoRef: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *IamQualifierRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'iam.Qualifier'
	if jsonDict["ClassId"] == "iam.Qualifier" {
		// try to unmarshal JSON data into IamQualifier
		err = json.Unmarshal(data, &dst.IamQualifier)
		if err == nil {
			jsonIamQualifier, _ := json.Marshal(dst.IamQualifier)
			if string(jsonIamQualifier) == "{}" { // empty struct
				dst.IamQualifier = nil
			} else {
				return nil // data stored in dst.IamQualifier, return on the first match
			}
		} else {
			dst.IamQualifier = nil
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

	// try to unmarshal data into IamQualifier
	err = json.Unmarshal(data, &dst.IamQualifier)
	if err == nil {
		jsonIamQualifier, _ := json.Marshal(dst.IamQualifier)
		if string(jsonIamQualifier) == "{}" { // empty struct
			dst.IamQualifier = nil
		} else {
			match++
		}
	} else {
		dst.IamQualifier = nil
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
		dst.IamQualifier = nil
		dst.MoMoRef = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(IamQualifierRelationship)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(IamQualifierRelationship)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IamQualifierRelationship) MarshalJSON() ([]byte, error) {
	if src.IamQualifier != nil {
		return json.Marshal(&src.IamQualifier)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *IamQualifierRelationship) GetActualInstance() interface{} {
	if obj.IamQualifier != nil {
		return obj.IamQualifier
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableIamQualifierRelationship struct {
	value *IamQualifierRelationship
	isSet bool
}

func (v NullableIamQualifierRelationship) Get() *IamQualifierRelationship {
	return v.value
}

func (v *NullableIamQualifierRelationship) Set(val *IamQualifierRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableIamQualifierRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableIamQualifierRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamQualifierRelationship(val *IamQualifierRelationship) *NullableIamQualifierRelationship {
	return &NullableIamQualifierRelationship{value: val, isSet: true}
}

func (v NullableIamQualifierRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamQualifierRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
