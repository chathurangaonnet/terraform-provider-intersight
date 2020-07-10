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

// HclConstraintAllOf Definition of the list of properties defined in 'hcl.Constraint', excluding properties defined in parent classes.
type HclConstraintAllOf struct {
	// Name or key of the applicable compatibility constraint.
	ConstraintName *string `json:"ConstraintName,omitempty"`
	// Value of the applicable compatibility constraint. Could either be a string value or a regex.
	ConstraintValue *string `json:"ConstraintValue,omitempty"`
}

// NewHclConstraintAllOf instantiates a new HclConstraintAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHclConstraintAllOf() *HclConstraintAllOf {
	this := HclConstraintAllOf{}
	return &this
}

// NewHclConstraintAllOfWithDefaults instantiates a new HclConstraintAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHclConstraintAllOfWithDefaults() *HclConstraintAllOf {
	this := HclConstraintAllOf{}
	return &this
}

// GetConstraintName returns the ConstraintName field value if set, zero value otherwise.
func (o *HclConstraintAllOf) GetConstraintName() string {
	if o == nil || o.ConstraintName == nil {
		var ret string
		return ret
	}
	return *o.ConstraintName
}

// GetConstraintNameOk returns a tuple with the ConstraintName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclConstraintAllOf) GetConstraintNameOk() (*string, bool) {
	if o == nil || o.ConstraintName == nil {
		return nil, false
	}
	return o.ConstraintName, true
}

// HasConstraintName returns a boolean if a field has been set.
func (o *HclConstraintAllOf) HasConstraintName() bool {
	if o != nil && o.ConstraintName != nil {
		return true
	}

	return false
}

// SetConstraintName gets a reference to the given string and assigns it to the ConstraintName field.
func (o *HclConstraintAllOf) SetConstraintName(v string) {
	o.ConstraintName = &v
}

// GetConstraintValue returns the ConstraintValue field value if set, zero value otherwise.
func (o *HclConstraintAllOf) GetConstraintValue() string {
	if o == nil || o.ConstraintValue == nil {
		var ret string
		return ret
	}
	return *o.ConstraintValue
}

// GetConstraintValueOk returns a tuple with the ConstraintValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclConstraintAllOf) GetConstraintValueOk() (*string, bool) {
	if o == nil || o.ConstraintValue == nil {
		return nil, false
	}
	return o.ConstraintValue, true
}

// HasConstraintValue returns a boolean if a field has been set.
func (o *HclConstraintAllOf) HasConstraintValue() bool {
	if o != nil && o.ConstraintValue != nil {
		return true
	}

	return false
}

// SetConstraintValue gets a reference to the given string and assigns it to the ConstraintValue field.
func (o *HclConstraintAllOf) SetConstraintValue(v string) {
	o.ConstraintValue = &v
}

func (o HclConstraintAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ConstraintName != nil {
		toSerialize["ConstraintName"] = o.ConstraintName
	}
	if o.ConstraintValue != nil {
		toSerialize["ConstraintValue"] = o.ConstraintValue
	}
	return json.Marshal(toSerialize)
}

type NullableHclConstraintAllOf struct {
	value *HclConstraintAllOf
	isSet bool
}

func (v NullableHclConstraintAllOf) Get() *HclConstraintAllOf {
	return v.value
}

func (v *NullableHclConstraintAllOf) Set(val *HclConstraintAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHclConstraintAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHclConstraintAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHclConstraintAllOf(val *HclConstraintAllOf) *NullableHclConstraintAllOf {
	return &NullableHclConstraintAllOf{value: val, isSet: true}
}

func (v NullableHclConstraintAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHclConstraintAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
