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

// MoMoRefAllOf Definition of the list of properties defined in 'mo.MoRef', excluding properties defined in parent classes.
type MoMoRefAllOf struct {
	// The Moid of the referenced REST resource.
	Moid *string `json:"Moid,omitempty"`
	// An OData $filter expression which describes the REST resource to be referenced. This field may be set instead of 'moid' by clients. 1. If 'moid' is set this field is ignored. 1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the resource matching the filter expression and populates it in the MoRef that is part of the object instance being inserted/updated to fulfill the REST request. An error is returned if the filter matches zero or more than one REST resource. An example filter string is: Serial eq '3AA8B7T11'.
	Selector *string `json:"Selector,omitempty"`
	// A URL to an instance of the 'mo.MoRef' class.
	Link                 *string `json:"link,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MoMoRefAllOf MoMoRefAllOf

// NewMoMoRefAllOf instantiates a new MoMoRefAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoMoRefAllOf() *MoMoRefAllOf {
	this := MoMoRefAllOf{}
	return &this
}

// NewMoMoRefAllOfWithDefaults instantiates a new MoMoRefAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoMoRefAllOfWithDefaults() *MoMoRefAllOf {
	this := MoMoRefAllOf{}
	return &this
}

// GetMoid returns the Moid field value if set, zero value otherwise.
func (o *MoMoRefAllOf) GetMoid() string {
	if o == nil || o.Moid == nil {
		var ret string
		return ret
	}
	return *o.Moid
}

// GetMoidOk returns a tuple with the Moid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoMoRefAllOf) GetMoidOk() (*string, bool) {
	if o == nil || o.Moid == nil {
		return nil, false
	}
	return o.Moid, true
}

// HasMoid returns a boolean if a field has been set.
func (o *MoMoRefAllOf) HasMoid() bool {
	if o != nil && o.Moid != nil {
		return true
	}

	return false
}

// SetMoid gets a reference to the given string and assigns it to the Moid field.
func (o *MoMoRefAllOf) SetMoid(v string) {
	o.Moid = &v
}

// GetSelector returns the Selector field value if set, zero value otherwise.
func (o *MoMoRefAllOf) GetSelector() string {
	if o == nil || o.Selector == nil {
		var ret string
		return ret
	}
	return *o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoMoRefAllOf) GetSelectorOk() (*string, bool) {
	if o == nil || o.Selector == nil {
		return nil, false
	}
	return o.Selector, true
}

// HasSelector returns a boolean if a field has been set.
func (o *MoMoRefAllOf) HasSelector() bool {
	if o != nil && o.Selector != nil {
		return true
	}

	return false
}

// SetSelector gets a reference to the given string and assigns it to the Selector field.
func (o *MoMoRefAllOf) SetSelector(v string) {
	o.Selector = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *MoMoRefAllOf) GetLink() string {
	if o == nil || o.Link == nil {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoMoRefAllOf) GetLinkOk() (*string, bool) {
	if o == nil || o.Link == nil {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *MoMoRefAllOf) HasLink() bool {
	if o != nil && o.Link != nil {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *MoMoRefAllOf) SetLink(v string) {
	o.Link = &v
}

func (o MoMoRefAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Moid != nil {
		toSerialize["Moid"] = o.Moid
	}
	if o.Selector != nil {
		toSerialize["Selector"] = o.Selector
	}
	if o.Link != nil {
		toSerialize["link"] = o.Link
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MoMoRefAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varMoMoRefAllOf := _MoMoRefAllOf{}

	if err = json.Unmarshal(bytes, &varMoMoRefAllOf); err == nil {
		*o = MoMoRefAllOf(varMoMoRefAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Moid")
		delete(additionalProperties, "Selector")
		delete(additionalProperties, "link")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMoMoRefAllOf struct {
	value *MoMoRefAllOf
	isSet bool
}

func (v NullableMoMoRefAllOf) Get() *MoMoRefAllOf {
	return v.value
}

func (v *NullableMoMoRefAllOf) Set(val *MoMoRefAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMoMoRefAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMoMoRefAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoMoRefAllOf(val *MoMoRefAllOf) *NullableMoMoRefAllOf {
	return &NullableMoMoRefAllOf{value: val, isSet: true}
}

func (v NullableMoMoRefAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoMoRefAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
