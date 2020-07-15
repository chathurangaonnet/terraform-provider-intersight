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

// ComputeRackUnitIdentityAllOf Definition of the list of properties defined in 'compute.RackUnitIdentity', excluding properties defined in parent classes.
type ComputeRackUnitIdentityAllOf struct {
	// Serial Identifier of an adapter participating in SWM.
	AdapterSerial        *string `json:"AdapterSerial,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ComputeRackUnitIdentityAllOf ComputeRackUnitIdentityAllOf

// NewComputeRackUnitIdentityAllOf instantiates a new ComputeRackUnitIdentityAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeRackUnitIdentityAllOf() *ComputeRackUnitIdentityAllOf {
	this := ComputeRackUnitIdentityAllOf{}
	return &this
}

// NewComputeRackUnitIdentityAllOfWithDefaults instantiates a new ComputeRackUnitIdentityAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeRackUnitIdentityAllOfWithDefaults() *ComputeRackUnitIdentityAllOf {
	this := ComputeRackUnitIdentityAllOf{}
	return &this
}

// GetAdapterSerial returns the AdapterSerial field value if set, zero value otherwise.
func (o *ComputeRackUnitIdentityAllOf) GetAdapterSerial() string {
	if o == nil || o.AdapterSerial == nil {
		var ret string
		return ret
	}
	return *o.AdapterSerial
}

// GetAdapterSerialOk returns a tuple with the AdapterSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeRackUnitIdentityAllOf) GetAdapterSerialOk() (*string, bool) {
	if o == nil || o.AdapterSerial == nil {
		return nil, false
	}
	return o.AdapterSerial, true
}

// HasAdapterSerial returns a boolean if a field has been set.
func (o *ComputeRackUnitIdentityAllOf) HasAdapterSerial() bool {
	if o != nil && o.AdapterSerial != nil {
		return true
	}

	return false
}

// SetAdapterSerial gets a reference to the given string and assigns it to the AdapterSerial field.
func (o *ComputeRackUnitIdentityAllOf) SetAdapterSerial(v string) {
	o.AdapterSerial = &v
}

func (o ComputeRackUnitIdentityAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdapterSerial != nil {
		toSerialize["AdapterSerial"] = o.AdapterSerial
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ComputeRackUnitIdentityAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varComputeRackUnitIdentityAllOf := _ComputeRackUnitIdentityAllOf{}

	if err = json.Unmarshal(bytes, &varComputeRackUnitIdentityAllOf); err == nil {
		*o = ComputeRackUnitIdentityAllOf(varComputeRackUnitIdentityAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "AdapterSerial")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableComputeRackUnitIdentityAllOf struct {
	value *ComputeRackUnitIdentityAllOf
	isSet bool
}

func (v NullableComputeRackUnitIdentityAllOf) Get() *ComputeRackUnitIdentityAllOf {
	return v.value
}

func (v *NullableComputeRackUnitIdentityAllOf) Set(val *ComputeRackUnitIdentityAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableComputeRackUnitIdentityAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableComputeRackUnitIdentityAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputeRackUnitIdentityAllOf(val *ComputeRackUnitIdentityAllOf) *NullableComputeRackUnitIdentityAllOf {
	return &NullableComputeRackUnitIdentityAllOf{value: val, isSet: true}
}

func (v NullableComputeRackUnitIdentityAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputeRackUnitIdentityAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
