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

// OsIpv4ConfigurationAllOf Definition of the list of properties defined in 'os.Ipv4Configuration', excluding properties defined in parent classes.
type OsIpv4ConfigurationAllOf struct {
	IpV4Config           *CommIpV4Interface `json:"IpV4Config,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OsIpv4ConfigurationAllOf OsIpv4ConfigurationAllOf

// NewOsIpv4ConfigurationAllOf instantiates a new OsIpv4ConfigurationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsIpv4ConfigurationAllOf() *OsIpv4ConfigurationAllOf {
	this := OsIpv4ConfigurationAllOf{}
	return &this
}

// NewOsIpv4ConfigurationAllOfWithDefaults instantiates a new OsIpv4ConfigurationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsIpv4ConfigurationAllOfWithDefaults() *OsIpv4ConfigurationAllOf {
	this := OsIpv4ConfigurationAllOf{}
	return &this
}

// GetIpV4Config returns the IpV4Config field value if set, zero value otherwise.
func (o *OsIpv4ConfigurationAllOf) GetIpV4Config() CommIpV4Interface {
	if o == nil || o.IpV4Config == nil {
		var ret CommIpV4Interface
		return ret
	}
	return *o.IpV4Config
}

// GetIpV4ConfigOk returns a tuple with the IpV4Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsIpv4ConfigurationAllOf) GetIpV4ConfigOk() (*CommIpV4Interface, bool) {
	if o == nil || o.IpV4Config == nil {
		return nil, false
	}
	return o.IpV4Config, true
}

// HasIpV4Config returns a boolean if a field has been set.
func (o *OsIpv4ConfigurationAllOf) HasIpV4Config() bool {
	if o != nil && o.IpV4Config != nil {
		return true
	}

	return false
}

// SetIpV4Config gets a reference to the given CommIpV4Interface and assigns it to the IpV4Config field.
func (o *OsIpv4ConfigurationAllOf) SetIpV4Config(v CommIpV4Interface) {
	o.IpV4Config = &v
}

func (o OsIpv4ConfigurationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IpV4Config != nil {
		toSerialize["IpV4Config"] = o.IpV4Config
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OsIpv4ConfigurationAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varOsIpv4ConfigurationAllOf := _OsIpv4ConfigurationAllOf{}

	if err = json.Unmarshal(bytes, &varOsIpv4ConfigurationAllOf); err == nil {
		*o = OsIpv4ConfigurationAllOf(varOsIpv4ConfigurationAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "IpV4Config")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOsIpv4ConfigurationAllOf struct {
	value *OsIpv4ConfigurationAllOf
	isSet bool
}

func (v NullableOsIpv4ConfigurationAllOf) Get() *OsIpv4ConfigurationAllOf {
	return v.value
}

func (v *NullableOsIpv4ConfigurationAllOf) Set(val *OsIpv4ConfigurationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOsIpv4ConfigurationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOsIpv4ConfigurationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsIpv4ConfigurationAllOf(val *OsIpv4ConfigurationAllOf) *NullableOsIpv4ConfigurationAllOf {
	return &NullableOsIpv4ConfigurationAllOf{value: val, isSet: true}
}

func (v NullableOsIpv4ConfigurationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsIpv4ConfigurationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
