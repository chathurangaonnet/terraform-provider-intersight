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

// ServerConfigResultEntry The profile configuration (deploy, validation) results details information.
type ServerConfigResultEntry struct {
	PolicyAbstractConfigResultEntry
	ConfigResult *ServerConfigResultRelationship `json:"ConfigResult,omitempty"`
}

// NewServerConfigResultEntry instantiates a new ServerConfigResultEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerConfigResultEntry() *ServerConfigResultEntry {
	this := ServerConfigResultEntry{}
	return &this
}

// NewServerConfigResultEntryWithDefaults instantiates a new ServerConfigResultEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerConfigResultEntryWithDefaults() *ServerConfigResultEntry {
	this := ServerConfigResultEntry{}
	return &this
}

// GetConfigResult returns the ConfigResult field value if set, zero value otherwise.
func (o *ServerConfigResultEntry) GetConfigResult() ServerConfigResultRelationship {
	if o == nil || o.ConfigResult == nil {
		var ret ServerConfigResultRelationship
		return ret
	}
	return *o.ConfigResult
}

// GetConfigResultOk returns a tuple with the ConfigResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigResultEntry) GetConfigResultOk() (*ServerConfigResultRelationship, bool) {
	if o == nil || o.ConfigResult == nil {
		return nil, false
	}
	return o.ConfigResult, true
}

// HasConfigResult returns a boolean if a field has been set.
func (o *ServerConfigResultEntry) HasConfigResult() bool {
	if o != nil && o.ConfigResult != nil {
		return true
	}

	return false
}

// SetConfigResult gets a reference to the given ServerConfigResultRelationship and assigns it to the ConfigResult field.
func (o *ServerConfigResultEntry) SetConfigResult(v ServerConfigResultRelationship) {
	o.ConfigResult = &v
}

func (o ServerConfigResultEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractConfigResultEntry, errPolicyAbstractConfigResultEntry := json.Marshal(o.PolicyAbstractConfigResultEntry)
	if errPolicyAbstractConfigResultEntry != nil {
		return []byte{}, errPolicyAbstractConfigResultEntry
	}
	errPolicyAbstractConfigResultEntry = json.Unmarshal([]byte(serializedPolicyAbstractConfigResultEntry), &toSerialize)
	if errPolicyAbstractConfigResultEntry != nil {
		return []byte{}, errPolicyAbstractConfigResultEntry
	}
	if o.ConfigResult != nil {
		toSerialize["ConfigResult"] = o.ConfigResult
	}
	return json.Marshal(toSerialize)
}

type NullableServerConfigResultEntry struct {
	value *ServerConfigResultEntry
	isSet bool
}

func (v NullableServerConfigResultEntry) Get() *ServerConfigResultEntry {
	return v.value
}

func (v *NullableServerConfigResultEntry) Set(val *ServerConfigResultEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableServerConfigResultEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableServerConfigResultEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerConfigResultEntry(val *ServerConfigResultEntry) *NullableServerConfigResultEntry {
	return &NullableServerConfigResultEntry{value: val, isSet: true}
}

func (v NullableServerConfigResultEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerConfigResultEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
