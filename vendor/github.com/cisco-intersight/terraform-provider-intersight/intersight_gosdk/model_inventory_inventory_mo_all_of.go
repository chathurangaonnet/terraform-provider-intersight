/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-06-19T15:15:17Z.
 *
 * API version: 1.0.9-1903
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
)

// InventoryInventoryMoAllOf Definition of the list of properties defined in 'inventory.InventoryMo', excluding properties defined in parent classes.
type InventoryInventoryMoAllOf struct {
	// The UCS DN of the MO for which the latest inventory to be fetched. If this property is empty and moId property has the Moid of the MO to be updated, the Moid will be used. If this property is empty and moId is also empty, all the MOs of the given moType will be updated.
	MoDn *string `json:"MoDn,omitempty"`
	// The MO id of an MO for which the latest inventory to be fetched. If this property is empty and moDn property has the UCS DN of the MO to be updated, the DN will be used. If this property is empty and moDn is also empty, all the MOs of the given moType will be updated.
	MoId *string `json:"MoId,omitempty"`
	// The type of the MO for which the latest inventory to be fetched.
	MoType *string `json:"MoType,omitempty"`
}

// NewInventoryInventoryMoAllOf instantiates a new InventoryInventoryMoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryInventoryMoAllOf() *InventoryInventoryMoAllOf {
	this := InventoryInventoryMoAllOf{}
	return &this
}

// NewInventoryInventoryMoAllOfWithDefaults instantiates a new InventoryInventoryMoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryInventoryMoAllOfWithDefaults() *InventoryInventoryMoAllOf {
	this := InventoryInventoryMoAllOf{}
	return &this
}

// GetMoDn returns the MoDn field value if set, zero value otherwise.
func (o *InventoryInventoryMoAllOf) GetMoDn() string {
	if o == nil || o.MoDn == nil {
		var ret string
		return ret
	}
	return *o.MoDn
}

// GetMoDnOk returns a tuple with the MoDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryInventoryMoAllOf) GetMoDnOk() (*string, bool) {
	if o == nil || o.MoDn == nil {
		return nil, false
	}
	return o.MoDn, true
}

// HasMoDn returns a boolean if a field has been set.
func (o *InventoryInventoryMoAllOf) HasMoDn() bool {
	if o != nil && o.MoDn != nil {
		return true
	}

	return false
}

// SetMoDn gets a reference to the given string and assigns it to the MoDn field.
func (o *InventoryInventoryMoAllOf) SetMoDn(v string) {
	o.MoDn = &v
}

// GetMoId returns the MoId field value if set, zero value otherwise.
func (o *InventoryInventoryMoAllOf) GetMoId() string {
	if o == nil || o.MoId == nil {
		var ret string
		return ret
	}
	return *o.MoId
}

// GetMoIdOk returns a tuple with the MoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryInventoryMoAllOf) GetMoIdOk() (*string, bool) {
	if o == nil || o.MoId == nil {
		return nil, false
	}
	return o.MoId, true
}

// HasMoId returns a boolean if a field has been set.
func (o *InventoryInventoryMoAllOf) HasMoId() bool {
	if o != nil && o.MoId != nil {
		return true
	}

	return false
}

// SetMoId gets a reference to the given string and assigns it to the MoId field.
func (o *InventoryInventoryMoAllOf) SetMoId(v string) {
	o.MoId = &v
}

// GetMoType returns the MoType field value if set, zero value otherwise.
func (o *InventoryInventoryMoAllOf) GetMoType() string {
	if o == nil || o.MoType == nil {
		var ret string
		return ret
	}
	return *o.MoType
}

// GetMoTypeOk returns a tuple with the MoType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryInventoryMoAllOf) GetMoTypeOk() (*string, bool) {
	if o == nil || o.MoType == nil {
		return nil, false
	}
	return o.MoType, true
}

// HasMoType returns a boolean if a field has been set.
func (o *InventoryInventoryMoAllOf) HasMoType() bool {
	if o != nil && o.MoType != nil {
		return true
	}

	return false
}

// SetMoType gets a reference to the given string and assigns it to the MoType field.
func (o *InventoryInventoryMoAllOf) SetMoType(v string) {
	o.MoType = &v
}

func (o InventoryInventoryMoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MoDn != nil {
		toSerialize["MoDn"] = o.MoDn
	}
	if o.MoId != nil {
		toSerialize["MoId"] = o.MoId
	}
	if o.MoType != nil {
		toSerialize["MoType"] = o.MoType
	}
	return json.Marshal(toSerialize)
}

type NullableInventoryInventoryMoAllOf struct {
	value *InventoryInventoryMoAllOf
	isSet bool
}

func (v NullableInventoryInventoryMoAllOf) Get() *InventoryInventoryMoAllOf {
	return v.value
}

func (v *NullableInventoryInventoryMoAllOf) Set(val *InventoryInventoryMoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryInventoryMoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryInventoryMoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryInventoryMoAllOf(val *InventoryInventoryMoAllOf) *NullableInventoryInventoryMoAllOf {
	return &NullableInventoryInventoryMoAllOf{value: val, isSet: true}
}

func (v NullableInventoryInventoryMoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryInventoryMoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}