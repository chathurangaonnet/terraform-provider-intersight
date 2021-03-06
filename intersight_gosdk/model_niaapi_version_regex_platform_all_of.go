/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-06-17T02:04:50-07:00.
 *
 * API version: 1.0.9-1867
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
)

// NiaapiVersionRegexPlatformAllOf Definition of the list of properties defined in 'niaapi.VersionRegexPlatform', excluding properties defined in parent classes.
type NiaapiVersionRegexPlatformAllOf struct {
	// All long live version Regex pattern.
	Anyllregex           *string                `json:"Anyllregex,omitempty"`
	Currentlltrain       *NiaapiSoftwareRegex   `json:"Currentlltrain,omitempty"`
	Latestsltrain        *NiaapiSoftwareRegex   `json:"Latestsltrain,omitempty"`
	Sltrain              *[]NiaapiSoftwareRegex `json:"Sltrain,omitempty"`
	Upcominglltrain      *NiaapiSoftwareRegex   `json:"Upcominglltrain,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiaapiVersionRegexPlatformAllOf NiaapiVersionRegexPlatformAllOf

// NewNiaapiVersionRegexPlatformAllOf instantiates a new NiaapiVersionRegexPlatformAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiaapiVersionRegexPlatformAllOf() *NiaapiVersionRegexPlatformAllOf {
	this := NiaapiVersionRegexPlatformAllOf{}
	return &this
}

// NewNiaapiVersionRegexPlatformAllOfWithDefaults instantiates a new NiaapiVersionRegexPlatformAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiaapiVersionRegexPlatformAllOfWithDefaults() *NiaapiVersionRegexPlatformAllOf {
	this := NiaapiVersionRegexPlatformAllOf{}
	return &this
}

// GetAnyllregex returns the Anyllregex field value if set, zero value otherwise.
func (o *NiaapiVersionRegexPlatformAllOf) GetAnyllregex() string {
	if o == nil || o.Anyllregex == nil {
		var ret string
		return ret
	}
	return *o.Anyllregex
}

// GetAnyllregexOk returns a tuple with the Anyllregex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiVersionRegexPlatformAllOf) GetAnyllregexOk() (*string, bool) {
	if o == nil || o.Anyllregex == nil {
		return nil, false
	}
	return o.Anyllregex, true
}

// HasAnyllregex returns a boolean if a field has been set.
func (o *NiaapiVersionRegexPlatformAllOf) HasAnyllregex() bool {
	if o != nil && o.Anyllregex != nil {
		return true
	}

	return false
}

// SetAnyllregex gets a reference to the given string and assigns it to the Anyllregex field.
func (o *NiaapiVersionRegexPlatformAllOf) SetAnyllregex(v string) {
	o.Anyllregex = &v
}

// GetCurrentlltrain returns the Currentlltrain field value if set, zero value otherwise.
func (o *NiaapiVersionRegexPlatformAllOf) GetCurrentlltrain() NiaapiSoftwareRegex {
	if o == nil || o.Currentlltrain == nil {
		var ret NiaapiSoftwareRegex
		return ret
	}
	return *o.Currentlltrain
}

// GetCurrentlltrainOk returns a tuple with the Currentlltrain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiVersionRegexPlatformAllOf) GetCurrentlltrainOk() (*NiaapiSoftwareRegex, bool) {
	if o == nil || o.Currentlltrain == nil {
		return nil, false
	}
	return o.Currentlltrain, true
}

// HasCurrentlltrain returns a boolean if a field has been set.
func (o *NiaapiVersionRegexPlatformAllOf) HasCurrentlltrain() bool {
	if o != nil && o.Currentlltrain != nil {
		return true
	}

	return false
}

// SetCurrentlltrain gets a reference to the given NiaapiSoftwareRegex and assigns it to the Currentlltrain field.
func (o *NiaapiVersionRegexPlatformAllOf) SetCurrentlltrain(v NiaapiSoftwareRegex) {
	o.Currentlltrain = &v
}

// GetLatestsltrain returns the Latestsltrain field value if set, zero value otherwise.
func (o *NiaapiVersionRegexPlatformAllOf) GetLatestsltrain() NiaapiSoftwareRegex {
	if o == nil || o.Latestsltrain == nil {
		var ret NiaapiSoftwareRegex
		return ret
	}
	return *o.Latestsltrain
}

// GetLatestsltrainOk returns a tuple with the Latestsltrain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiVersionRegexPlatformAllOf) GetLatestsltrainOk() (*NiaapiSoftwareRegex, bool) {
	if o == nil || o.Latestsltrain == nil {
		return nil, false
	}
	return o.Latestsltrain, true
}

// HasLatestsltrain returns a boolean if a field has been set.
func (o *NiaapiVersionRegexPlatformAllOf) HasLatestsltrain() bool {
	if o != nil && o.Latestsltrain != nil {
		return true
	}

	return false
}

// SetLatestsltrain gets a reference to the given NiaapiSoftwareRegex and assigns it to the Latestsltrain field.
func (o *NiaapiVersionRegexPlatformAllOf) SetLatestsltrain(v NiaapiSoftwareRegex) {
	o.Latestsltrain = &v
}

// GetSltrain returns the Sltrain field value if set, zero value otherwise.
func (o *NiaapiVersionRegexPlatformAllOf) GetSltrain() []NiaapiSoftwareRegex {
	if o == nil || o.Sltrain == nil {
		var ret []NiaapiSoftwareRegex
		return ret
	}
	return *o.Sltrain
}

// GetSltrainOk returns a tuple with the Sltrain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiVersionRegexPlatformAllOf) GetSltrainOk() (*[]NiaapiSoftwareRegex, bool) {
	if o == nil || o.Sltrain == nil {
		return nil, false
	}
	return o.Sltrain, true
}

// HasSltrain returns a boolean if a field has been set.
func (o *NiaapiVersionRegexPlatformAllOf) HasSltrain() bool {
	if o != nil && o.Sltrain != nil {
		return true
	}

	return false
}

// SetSltrain gets a reference to the given []NiaapiSoftwareRegex and assigns it to the Sltrain field.
func (o *NiaapiVersionRegexPlatformAllOf) SetSltrain(v []NiaapiSoftwareRegex) {
	o.Sltrain = &v
}

// GetUpcominglltrain returns the Upcominglltrain field value if set, zero value otherwise.
func (o *NiaapiVersionRegexPlatformAllOf) GetUpcominglltrain() NiaapiSoftwareRegex {
	if o == nil || o.Upcominglltrain == nil {
		var ret NiaapiSoftwareRegex
		return ret
	}
	return *o.Upcominglltrain
}

// GetUpcominglltrainOk returns a tuple with the Upcominglltrain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiVersionRegexPlatformAllOf) GetUpcominglltrainOk() (*NiaapiSoftwareRegex, bool) {
	if o == nil || o.Upcominglltrain == nil {
		return nil, false
	}
	return o.Upcominglltrain, true
}

// HasUpcominglltrain returns a boolean if a field has been set.
func (o *NiaapiVersionRegexPlatformAllOf) HasUpcominglltrain() bool {
	if o != nil && o.Upcominglltrain != nil {
		return true
	}

	return false
}

// SetUpcominglltrain gets a reference to the given NiaapiSoftwareRegex and assigns it to the Upcominglltrain field.
func (o *NiaapiVersionRegexPlatformAllOf) SetUpcominglltrain(v NiaapiSoftwareRegex) {
	o.Upcominglltrain = &v
}

func (o NiaapiVersionRegexPlatformAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Anyllregex != nil {
		toSerialize["Anyllregex"] = o.Anyllregex
	}
	if o.Currentlltrain != nil {
		toSerialize["Currentlltrain"] = o.Currentlltrain
	}
	if o.Latestsltrain != nil {
		toSerialize["Latestsltrain"] = o.Latestsltrain
	}
	if o.Sltrain != nil {
		toSerialize["Sltrain"] = o.Sltrain
	}
	if o.Upcominglltrain != nil {
		toSerialize["Upcominglltrain"] = o.Upcominglltrain
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiaapiVersionRegexPlatformAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiaapiVersionRegexPlatformAllOf := _NiaapiVersionRegexPlatformAllOf{}

	if err = json.Unmarshal(bytes, &varNiaapiVersionRegexPlatformAllOf); err == nil {
		*o = NiaapiVersionRegexPlatformAllOf(varNiaapiVersionRegexPlatformAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Anyllregex")
		delete(additionalProperties, "Currentlltrain")
		delete(additionalProperties, "Latestsltrain")
		delete(additionalProperties, "Sltrain")
		delete(additionalProperties, "Upcominglltrain")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiaapiVersionRegexPlatformAllOf struct {
	value *NiaapiVersionRegexPlatformAllOf
	isSet bool
}

func (v NullableNiaapiVersionRegexPlatformAllOf) Get() *NiaapiVersionRegexPlatformAllOf {
	return v.value
}

func (v *NullableNiaapiVersionRegexPlatformAllOf) Set(val *NiaapiVersionRegexPlatformAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiaapiVersionRegexPlatformAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiaapiVersionRegexPlatformAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiaapiVersionRegexPlatformAllOf(val *NiaapiVersionRegexPlatformAllOf) *NullableNiaapiVersionRegexPlatformAllOf {
	return &NullableNiaapiVersionRegexPlatformAllOf{value: val, isSet: true}
}

func (v NullableNiaapiVersionRegexPlatformAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiaapiVersionRegexPlatformAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
