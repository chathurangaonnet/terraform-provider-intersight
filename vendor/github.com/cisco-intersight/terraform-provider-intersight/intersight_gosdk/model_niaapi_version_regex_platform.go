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

// NiaapiVersionRegexPlatform The regular expression to parse the software version strings for specific platform.
type NiaapiVersionRegexPlatform struct {
	MoBaseComplexType
	// All long live version Regex pattern.
	Anyllregex           *string                `json:"Anyllregex,omitempty"`
	Currentlltrain       *NiaapiSoftwareRegex   `json:"Currentlltrain,omitempty"`
	Latestsltrain        *NiaapiSoftwareRegex   `json:"Latestsltrain,omitempty"`
	Sltrain              *[]NiaapiSoftwareRegex `json:"Sltrain,omitempty"`
	Upcominglltrain      *NiaapiSoftwareRegex   `json:"Upcominglltrain,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiaapiVersionRegexPlatform NiaapiVersionRegexPlatform

// NewNiaapiVersionRegexPlatform instantiates a new NiaapiVersionRegexPlatform object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiaapiVersionRegexPlatform() *NiaapiVersionRegexPlatform {
	this := NiaapiVersionRegexPlatform{}
	return &this
}

// NewNiaapiVersionRegexPlatformWithDefaults instantiates a new NiaapiVersionRegexPlatform object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiaapiVersionRegexPlatformWithDefaults() *NiaapiVersionRegexPlatform {
	this := NiaapiVersionRegexPlatform{}
	return &this
}

// GetAnyllregex returns the Anyllregex field value if set, zero value otherwise.
func (o *NiaapiVersionRegexPlatform) GetAnyllregex() string {
	if o == nil || o.Anyllregex == nil {
		var ret string
		return ret
	}
	return *o.Anyllregex
}

// GetAnyllregexOk returns a tuple with the Anyllregex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiVersionRegexPlatform) GetAnyllregexOk() (*string, bool) {
	if o == nil || o.Anyllregex == nil {
		return nil, false
	}
	return o.Anyllregex, true
}

// HasAnyllregex returns a boolean if a field has been set.
func (o *NiaapiVersionRegexPlatform) HasAnyllregex() bool {
	if o != nil && o.Anyllregex != nil {
		return true
	}

	return false
}

// SetAnyllregex gets a reference to the given string and assigns it to the Anyllregex field.
func (o *NiaapiVersionRegexPlatform) SetAnyllregex(v string) {
	o.Anyllregex = &v
}

// GetCurrentlltrain returns the Currentlltrain field value if set, zero value otherwise.
func (o *NiaapiVersionRegexPlatform) GetCurrentlltrain() NiaapiSoftwareRegex {
	if o == nil || o.Currentlltrain == nil {
		var ret NiaapiSoftwareRegex
		return ret
	}
	return *o.Currentlltrain
}

// GetCurrentlltrainOk returns a tuple with the Currentlltrain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiVersionRegexPlatform) GetCurrentlltrainOk() (*NiaapiSoftwareRegex, bool) {
	if o == nil || o.Currentlltrain == nil {
		return nil, false
	}
	return o.Currentlltrain, true
}

// HasCurrentlltrain returns a boolean if a field has been set.
func (o *NiaapiVersionRegexPlatform) HasCurrentlltrain() bool {
	if o != nil && o.Currentlltrain != nil {
		return true
	}

	return false
}

// SetCurrentlltrain gets a reference to the given NiaapiSoftwareRegex and assigns it to the Currentlltrain field.
func (o *NiaapiVersionRegexPlatform) SetCurrentlltrain(v NiaapiSoftwareRegex) {
	o.Currentlltrain = &v
}

// GetLatestsltrain returns the Latestsltrain field value if set, zero value otherwise.
func (o *NiaapiVersionRegexPlatform) GetLatestsltrain() NiaapiSoftwareRegex {
	if o == nil || o.Latestsltrain == nil {
		var ret NiaapiSoftwareRegex
		return ret
	}
	return *o.Latestsltrain
}

// GetLatestsltrainOk returns a tuple with the Latestsltrain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiVersionRegexPlatform) GetLatestsltrainOk() (*NiaapiSoftwareRegex, bool) {
	if o == nil || o.Latestsltrain == nil {
		return nil, false
	}
	return o.Latestsltrain, true
}

// HasLatestsltrain returns a boolean if a field has been set.
func (o *NiaapiVersionRegexPlatform) HasLatestsltrain() bool {
	if o != nil && o.Latestsltrain != nil {
		return true
	}

	return false
}

// SetLatestsltrain gets a reference to the given NiaapiSoftwareRegex and assigns it to the Latestsltrain field.
func (o *NiaapiVersionRegexPlatform) SetLatestsltrain(v NiaapiSoftwareRegex) {
	o.Latestsltrain = &v
}

// GetSltrain returns the Sltrain field value if set, zero value otherwise.
func (o *NiaapiVersionRegexPlatform) GetSltrain() []NiaapiSoftwareRegex {
	if o == nil || o.Sltrain == nil {
		var ret []NiaapiSoftwareRegex
		return ret
	}
	return *o.Sltrain
}

// GetSltrainOk returns a tuple with the Sltrain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiVersionRegexPlatform) GetSltrainOk() (*[]NiaapiSoftwareRegex, bool) {
	if o == nil || o.Sltrain == nil {
		return nil, false
	}
	return o.Sltrain, true
}

// HasSltrain returns a boolean if a field has been set.
func (o *NiaapiVersionRegexPlatform) HasSltrain() bool {
	if o != nil && o.Sltrain != nil {
		return true
	}

	return false
}

// SetSltrain gets a reference to the given []NiaapiSoftwareRegex and assigns it to the Sltrain field.
func (o *NiaapiVersionRegexPlatform) SetSltrain(v []NiaapiSoftwareRegex) {
	o.Sltrain = &v
}

// GetUpcominglltrain returns the Upcominglltrain field value if set, zero value otherwise.
func (o *NiaapiVersionRegexPlatform) GetUpcominglltrain() NiaapiSoftwareRegex {
	if o == nil || o.Upcominglltrain == nil {
		var ret NiaapiSoftwareRegex
		return ret
	}
	return *o.Upcominglltrain
}

// GetUpcominglltrainOk returns a tuple with the Upcominglltrain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiVersionRegexPlatform) GetUpcominglltrainOk() (*NiaapiSoftwareRegex, bool) {
	if o == nil || o.Upcominglltrain == nil {
		return nil, false
	}
	return o.Upcominglltrain, true
}

// HasUpcominglltrain returns a boolean if a field has been set.
func (o *NiaapiVersionRegexPlatform) HasUpcominglltrain() bool {
	if o != nil && o.Upcominglltrain != nil {
		return true
	}

	return false
}

// SetUpcominglltrain gets a reference to the given NiaapiSoftwareRegex and assigns it to the Upcominglltrain field.
func (o *NiaapiVersionRegexPlatform) SetUpcominglltrain(v NiaapiSoftwareRegex) {
	o.Upcominglltrain = &v
}

func (o NiaapiVersionRegexPlatform) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
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

func (o *NiaapiVersionRegexPlatform) UnmarshalJSON(bytes []byte) (err error) {
	varNiaapiVersionRegexPlatform := _NiaapiVersionRegexPlatform{}

	if err = json.Unmarshal(bytes, &varNiaapiVersionRegexPlatform); err == nil {
		*o = NiaapiVersionRegexPlatform(varNiaapiVersionRegexPlatform)
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

type NullableNiaapiVersionRegexPlatform struct {
	value *NiaapiVersionRegexPlatform
	isSet bool
}

func (v NullableNiaapiVersionRegexPlatform) Get() *NiaapiVersionRegexPlatform {
	return v.value
}

func (v *NullableNiaapiVersionRegexPlatform) Set(val *NiaapiVersionRegexPlatform) {
	v.value = val
	v.isSet = true
}

func (v NullableNiaapiVersionRegexPlatform) IsSet() bool {
	return v.isSet
}

func (v *NullableNiaapiVersionRegexPlatform) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiaapiVersionRegexPlatform(val *NiaapiVersionRegexPlatform) *NullableNiaapiVersionRegexPlatform {
	return &NullableNiaapiVersionRegexPlatform{value: val, isSet: true}
}

func (v NullableNiaapiVersionRegexPlatform) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiaapiVersionRegexPlatform) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
