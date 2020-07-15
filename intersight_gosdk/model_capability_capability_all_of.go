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
)

// CapabilityCapabilityAllOf Definition of the list of properties defined in 'capability.Capability', excluding properties defined in parent classes.
type CapabilityCapabilityAllOf struct {
	// An unique identifer for a capability descriptor.
	Name                 *string                        `json:"Name,omitempty"`
	Section              *CapabilitySectionRelationship `json:"Section,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilityCapabilityAllOf CapabilityCapabilityAllOf

// NewCapabilityCapabilityAllOf instantiates a new CapabilityCapabilityAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityCapabilityAllOf() *CapabilityCapabilityAllOf {
	this := CapabilityCapabilityAllOf{}
	return &this
}

// NewCapabilityCapabilityAllOfWithDefaults instantiates a new CapabilityCapabilityAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityCapabilityAllOfWithDefaults() *CapabilityCapabilityAllOf {
	this := CapabilityCapabilityAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CapabilityCapabilityAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityCapabilityAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CapabilityCapabilityAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CapabilityCapabilityAllOf) SetName(v string) {
	o.Name = &v
}

// GetSection returns the Section field value if set, zero value otherwise.
func (o *CapabilityCapabilityAllOf) GetSection() CapabilitySectionRelationship {
	if o == nil || o.Section == nil {
		var ret CapabilitySectionRelationship
		return ret
	}
	return *o.Section
}

// GetSectionOk returns a tuple with the Section field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityCapabilityAllOf) GetSectionOk() (*CapabilitySectionRelationship, bool) {
	if o == nil || o.Section == nil {
		return nil, false
	}
	return o.Section, true
}

// HasSection returns a boolean if a field has been set.
func (o *CapabilityCapabilityAllOf) HasSection() bool {
	if o != nil && o.Section != nil {
		return true
	}

	return false
}

// SetSection gets a reference to the given CapabilitySectionRelationship and assigns it to the Section field.
func (o *CapabilityCapabilityAllOf) SetSection(v CapabilitySectionRelationship) {
	o.Section = &v
}

func (o CapabilityCapabilityAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Section != nil {
		toSerialize["Section"] = o.Section
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CapabilityCapabilityAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCapabilityCapabilityAllOf := _CapabilityCapabilityAllOf{}

	if err = json.Unmarshal(bytes, &varCapabilityCapabilityAllOf); err == nil {
		*o = CapabilityCapabilityAllOf(varCapabilityCapabilityAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Section")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCapabilityCapabilityAllOf struct {
	value *CapabilityCapabilityAllOf
	isSet bool
}

func (v NullableCapabilityCapabilityAllOf) Get() *CapabilityCapabilityAllOf {
	return v.value
}

func (v *NullableCapabilityCapabilityAllOf) Set(val *CapabilityCapabilityAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityCapabilityAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityCapabilityAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityCapabilityAllOf(val *CapabilityCapabilityAllOf) *NullableCapabilityCapabilityAllOf {
	return &NullableCapabilityCapabilityAllOf{value: val, isSet: true}
}

func (v NullableCapabilityCapabilityAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityCapabilityAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
