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

// OsOsSupport OsSupport is used to validate the support for an Operating System's version.
type OsOsSupport struct {
	MoBaseMo
	// The version of the Operating System to be validated. The version should be as per HCL.
	OsVersion *string `json:"OsVersion,omitempty"`
}

// NewOsOsSupport instantiates a new OsOsSupport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsOsSupport() *OsOsSupport {
	this := OsOsSupport{}
	return &this
}

// NewOsOsSupportWithDefaults instantiates a new OsOsSupport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsOsSupportWithDefaults() *OsOsSupport {
	this := OsOsSupport{}
	return &this
}

// GetOsVersion returns the OsVersion field value if set, zero value otherwise.
func (o *OsOsSupport) GetOsVersion() string {
	if o == nil || o.OsVersion == nil {
		var ret string
		return ret
	}
	return *o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsOsSupport) GetOsVersionOk() (*string, bool) {
	if o == nil || o.OsVersion == nil {
		return nil, false
	}
	return o.OsVersion, true
}

// HasOsVersion returns a boolean if a field has been set.
func (o *OsOsSupport) HasOsVersion() bool {
	if o != nil && o.OsVersion != nil {
		return true
	}

	return false
}

// SetOsVersion gets a reference to the given string and assigns it to the OsVersion field.
func (o *OsOsSupport) SetOsVersion(v string) {
	o.OsVersion = &v
}

func (o OsOsSupport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.OsVersion != nil {
		toSerialize["OsVersion"] = o.OsVersion
	}
	return json.Marshal(toSerialize)
}

type NullableOsOsSupport struct {
	value *OsOsSupport
	isSet bool
}

func (v NullableOsOsSupport) Get() *OsOsSupport {
	return v.value
}

func (v *NullableOsOsSupport) Set(val *OsOsSupport) {
	v.value = val
	v.isSet = true
}

func (v NullableOsOsSupport) IsSet() bool {
	return v.isSet
}

func (v *NullableOsOsSupport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsOsSupport(val *OsOsSupport) *NullableOsOsSupport {
	return &NullableOsOsSupport{value: val, isSet: true}
}

func (v NullableOsOsSupport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsOsSupport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
