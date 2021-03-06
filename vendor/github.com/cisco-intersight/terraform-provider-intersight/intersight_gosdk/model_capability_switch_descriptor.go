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

// CapabilitySwitchDescriptor Descriptor that uniquely identifies a Fabric interconnect.
type CapabilitySwitchDescriptor struct {
	CapabilityHardwareDescriptor
	// The total expected memory for this hardware.
	ExpectedMemory *int64 `json:"ExpectedMemory,omitempty"`
	// Revision for the fabric interconnect.
	Revision             *string `json:"Revision,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilitySwitchDescriptor CapabilitySwitchDescriptor

// NewCapabilitySwitchDescriptor instantiates a new CapabilitySwitchDescriptor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilitySwitchDescriptor() *CapabilitySwitchDescriptor {
	this := CapabilitySwitchDescriptor{}
	return &this
}

// NewCapabilitySwitchDescriptorWithDefaults instantiates a new CapabilitySwitchDescriptor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilitySwitchDescriptorWithDefaults() *CapabilitySwitchDescriptor {
	this := CapabilitySwitchDescriptor{}
	return &this
}

// GetExpectedMemory returns the ExpectedMemory field value if set, zero value otherwise.
func (o *CapabilitySwitchDescriptor) GetExpectedMemory() int64 {
	if o == nil || o.ExpectedMemory == nil {
		var ret int64
		return ret
	}
	return *o.ExpectedMemory
}

// GetExpectedMemoryOk returns a tuple with the ExpectedMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchDescriptor) GetExpectedMemoryOk() (*int64, bool) {
	if o == nil || o.ExpectedMemory == nil {
		return nil, false
	}
	return o.ExpectedMemory, true
}

// HasExpectedMemory returns a boolean if a field has been set.
func (o *CapabilitySwitchDescriptor) HasExpectedMemory() bool {
	if o != nil && o.ExpectedMemory != nil {
		return true
	}

	return false
}

// SetExpectedMemory gets a reference to the given int64 and assigns it to the ExpectedMemory field.
func (o *CapabilitySwitchDescriptor) SetExpectedMemory(v int64) {
	o.ExpectedMemory = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *CapabilitySwitchDescriptor) GetRevision() string {
	if o == nil || o.Revision == nil {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchDescriptor) GetRevisionOk() (*string, bool) {
	if o == nil || o.Revision == nil {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *CapabilitySwitchDescriptor) HasRevision() bool {
	if o != nil && o.Revision != nil {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *CapabilitySwitchDescriptor) SetRevision(v string) {
	o.Revision = &v
}

func (o CapabilitySwitchDescriptor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCapabilityHardwareDescriptor, errCapabilityHardwareDescriptor := json.Marshal(o.CapabilityHardwareDescriptor)
	if errCapabilityHardwareDescriptor != nil {
		return []byte{}, errCapabilityHardwareDescriptor
	}
	errCapabilityHardwareDescriptor = json.Unmarshal([]byte(serializedCapabilityHardwareDescriptor), &toSerialize)
	if errCapabilityHardwareDescriptor != nil {
		return []byte{}, errCapabilityHardwareDescriptor
	}
	if o.ExpectedMemory != nil {
		toSerialize["ExpectedMemory"] = o.ExpectedMemory
	}
	if o.Revision != nil {
		toSerialize["Revision"] = o.Revision
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CapabilitySwitchDescriptor) UnmarshalJSON(bytes []byte) (err error) {
	type CapabilitySwitchDescriptorWithoutEmbeddedStruct struct {
		// The total expected memory for this hardware.
		ExpectedMemory *int64 `json:"ExpectedMemory,omitempty"`
		// Revision for the fabric interconnect.
		Revision *string `json:"Revision,omitempty"`
	}

	varCapabilitySwitchDescriptorWithoutEmbeddedStruct := CapabilitySwitchDescriptorWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCapabilitySwitchDescriptorWithoutEmbeddedStruct)
	if err == nil {
		varCapabilitySwitchDescriptor := _CapabilitySwitchDescriptor{}
		varCapabilitySwitchDescriptor.ExpectedMemory = varCapabilitySwitchDescriptorWithoutEmbeddedStruct.ExpectedMemory
		varCapabilitySwitchDescriptor.Revision = varCapabilitySwitchDescriptorWithoutEmbeddedStruct.Revision
		*o = CapabilitySwitchDescriptor(varCapabilitySwitchDescriptor)
	} else {
		return err
	}

	varCapabilitySwitchDescriptor := _CapabilitySwitchDescriptor{}

	err = json.Unmarshal(bytes, &varCapabilitySwitchDescriptor)
	if err == nil {
		o.CapabilityHardwareDescriptor = varCapabilitySwitchDescriptor.CapabilityHardwareDescriptor
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ExpectedMemory")
		delete(additionalProperties, "Revision")

		// remove fields from embedded structs
		reflectCapabilityHardwareDescriptor := reflect.ValueOf(o.CapabilityHardwareDescriptor)
		for i := 0; i < reflectCapabilityHardwareDescriptor.Type().NumField(); i++ {
			t := reflectCapabilityHardwareDescriptor.Type().Field(i)

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

type NullableCapabilitySwitchDescriptor struct {
	value *CapabilitySwitchDescriptor
	isSet bool
}

func (v NullableCapabilitySwitchDescriptor) Get() *CapabilitySwitchDescriptor {
	return v.value
}

func (v *NullableCapabilitySwitchDescriptor) Set(val *CapabilitySwitchDescriptor) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilitySwitchDescriptor) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilitySwitchDescriptor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilitySwitchDescriptor(val *CapabilitySwitchDescriptor) *NullableCapabilitySwitchDescriptor {
	return &NullableCapabilitySwitchDescriptor{value: val, isSet: true}
}

func (v NullableCapabilitySwitchDescriptor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilitySwitchDescriptor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
