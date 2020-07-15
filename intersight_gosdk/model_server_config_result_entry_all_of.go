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

// ServerConfigResultEntryAllOf Definition of the list of properties defined in 'server.ConfigResultEntry', excluding properties defined in parent classes.
type ServerConfigResultEntryAllOf struct {
	ConfigResult         *ServerConfigResultRelationship `json:"ConfigResult,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServerConfigResultEntryAllOf ServerConfigResultEntryAllOf

// NewServerConfigResultEntryAllOf instantiates a new ServerConfigResultEntryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerConfigResultEntryAllOf() *ServerConfigResultEntryAllOf {
	this := ServerConfigResultEntryAllOf{}
	return &this
}

// NewServerConfigResultEntryAllOfWithDefaults instantiates a new ServerConfigResultEntryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerConfigResultEntryAllOfWithDefaults() *ServerConfigResultEntryAllOf {
	this := ServerConfigResultEntryAllOf{}
	return &this
}

// GetConfigResult returns the ConfigResult field value if set, zero value otherwise.
func (o *ServerConfigResultEntryAllOf) GetConfigResult() ServerConfigResultRelationship {
	if o == nil || o.ConfigResult == nil {
		var ret ServerConfigResultRelationship
		return ret
	}
	return *o.ConfigResult
}

// GetConfigResultOk returns a tuple with the ConfigResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigResultEntryAllOf) GetConfigResultOk() (*ServerConfigResultRelationship, bool) {
	if o == nil || o.ConfigResult == nil {
		return nil, false
	}
	return o.ConfigResult, true
}

// HasConfigResult returns a boolean if a field has been set.
func (o *ServerConfigResultEntryAllOf) HasConfigResult() bool {
	if o != nil && o.ConfigResult != nil {
		return true
	}

	return false
}

// SetConfigResult gets a reference to the given ServerConfigResultRelationship and assigns it to the ConfigResult field.
func (o *ServerConfigResultEntryAllOf) SetConfigResult(v ServerConfigResultRelationship) {
	o.ConfigResult = &v
}

func (o ServerConfigResultEntryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ConfigResult != nil {
		toSerialize["ConfigResult"] = o.ConfigResult
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ServerConfigResultEntryAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varServerConfigResultEntryAllOf := _ServerConfigResultEntryAllOf{}

	if err = json.Unmarshal(bytes, &varServerConfigResultEntryAllOf); err == nil {
		*o = ServerConfigResultEntryAllOf(varServerConfigResultEntryAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ConfigResult")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServerConfigResultEntryAllOf struct {
	value *ServerConfigResultEntryAllOf
	isSet bool
}

func (v NullableServerConfigResultEntryAllOf) Get() *ServerConfigResultEntryAllOf {
	return v.value
}

func (v *NullableServerConfigResultEntryAllOf) Set(val *ServerConfigResultEntryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableServerConfigResultEntryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableServerConfigResultEntryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerConfigResultEntryAllOf(val *ServerConfigResultEntryAllOf) *NullableServerConfigResultEntryAllOf {
	return &NullableServerConfigResultEntryAllOf{value: val, isSet: true}
}

func (v NullableServerConfigResultEntryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerConfigResultEntryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
