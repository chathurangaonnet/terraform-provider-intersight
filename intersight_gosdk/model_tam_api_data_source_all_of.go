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

// TamApiDataSourceAllOf Definition of the list of properties defined in 'tam.ApiDataSource', excluding properties defined in parent classes.
type TamApiDataSourceAllOf struct {
	// Type of Intersight managed object used as data source.
	MoType               *string          `json:"MoType,omitempty"`
	Queries              *[]TamQueryEntry `json:"Queries,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TamApiDataSourceAllOf TamApiDataSourceAllOf

// NewTamApiDataSourceAllOf instantiates a new TamApiDataSourceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTamApiDataSourceAllOf() *TamApiDataSourceAllOf {
	this := TamApiDataSourceAllOf{}
	return &this
}

// NewTamApiDataSourceAllOfWithDefaults instantiates a new TamApiDataSourceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTamApiDataSourceAllOfWithDefaults() *TamApiDataSourceAllOf {
	this := TamApiDataSourceAllOf{}
	return &this
}

// GetMoType returns the MoType field value if set, zero value otherwise.
func (o *TamApiDataSourceAllOf) GetMoType() string {
	if o == nil || o.MoType == nil {
		var ret string
		return ret
	}
	return *o.MoType
}

// GetMoTypeOk returns a tuple with the MoType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamApiDataSourceAllOf) GetMoTypeOk() (*string, bool) {
	if o == nil || o.MoType == nil {
		return nil, false
	}
	return o.MoType, true
}

// HasMoType returns a boolean if a field has been set.
func (o *TamApiDataSourceAllOf) HasMoType() bool {
	if o != nil && o.MoType != nil {
		return true
	}

	return false
}

// SetMoType gets a reference to the given string and assigns it to the MoType field.
func (o *TamApiDataSourceAllOf) SetMoType(v string) {
	o.MoType = &v
}

// GetQueries returns the Queries field value if set, zero value otherwise.
func (o *TamApiDataSourceAllOf) GetQueries() []TamQueryEntry {
	if o == nil || o.Queries == nil {
		var ret []TamQueryEntry
		return ret
	}
	return *o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamApiDataSourceAllOf) GetQueriesOk() (*[]TamQueryEntry, bool) {
	if o == nil || o.Queries == nil {
		return nil, false
	}
	return o.Queries, true
}

// HasQueries returns a boolean if a field has been set.
func (o *TamApiDataSourceAllOf) HasQueries() bool {
	if o != nil && o.Queries != nil {
		return true
	}

	return false
}

// SetQueries gets a reference to the given []TamQueryEntry and assigns it to the Queries field.
func (o *TamApiDataSourceAllOf) SetQueries(v []TamQueryEntry) {
	o.Queries = &v
}

func (o TamApiDataSourceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MoType != nil {
		toSerialize["MoType"] = o.MoType
	}
	if o.Queries != nil {
		toSerialize["Queries"] = o.Queries
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TamApiDataSourceAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTamApiDataSourceAllOf := _TamApiDataSourceAllOf{}

	if err = json.Unmarshal(bytes, &varTamApiDataSourceAllOf); err == nil {
		*o = TamApiDataSourceAllOf(varTamApiDataSourceAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "MoType")
		delete(additionalProperties, "Queries")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTamApiDataSourceAllOf struct {
	value *TamApiDataSourceAllOf
	isSet bool
}

func (v NullableTamApiDataSourceAllOf) Get() *TamApiDataSourceAllOf {
	return v.value
}

func (v *NullableTamApiDataSourceAllOf) Set(val *TamApiDataSourceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTamApiDataSourceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTamApiDataSourceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTamApiDataSourceAllOf(val *TamApiDataSourceAllOf) *NullableTamApiDataSourceAllOf {
	return &NullableTamApiDataSourceAllOf{value: val, isSet: true}
}

func (v NullableTamApiDataSourceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTamApiDataSourceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
