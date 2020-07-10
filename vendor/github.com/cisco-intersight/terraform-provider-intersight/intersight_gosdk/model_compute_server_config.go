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

// ComputeServerConfig Configurable properties of servers common to both a server and a server profile. The user will apply these properties on a server directly when the server is not associated with a server profile and through a server profile when the server is attached with a server profile.
type ComputeServerConfig struct {
	MoBaseComplexType
	// User defined asset tag of the server.
	AssetTag *string `json:"AssetTag,omitempty"`
	// User defined description of the server.
	UserLabel            *string `json:"UserLabel,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ComputeServerConfig ComputeServerConfig

// NewComputeServerConfig instantiates a new ComputeServerConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeServerConfig() *ComputeServerConfig {
	this := ComputeServerConfig{}
	return &this
}

// NewComputeServerConfigWithDefaults instantiates a new ComputeServerConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeServerConfigWithDefaults() *ComputeServerConfig {
	this := ComputeServerConfig{}
	return &this
}

// GetAssetTag returns the AssetTag field value if set, zero value otherwise.
func (o *ComputeServerConfig) GetAssetTag() string {
	if o == nil || o.AssetTag == nil {
		var ret string
		return ret
	}
	return *o.AssetTag
}

// GetAssetTagOk returns a tuple with the AssetTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeServerConfig) GetAssetTagOk() (*string, bool) {
	if o == nil || o.AssetTag == nil {
		return nil, false
	}
	return o.AssetTag, true
}

// HasAssetTag returns a boolean if a field has been set.
func (o *ComputeServerConfig) HasAssetTag() bool {
	if o != nil && o.AssetTag != nil {
		return true
	}

	return false
}

// SetAssetTag gets a reference to the given string and assigns it to the AssetTag field.
func (o *ComputeServerConfig) SetAssetTag(v string) {
	o.AssetTag = &v
}

// GetUserLabel returns the UserLabel field value if set, zero value otherwise.
func (o *ComputeServerConfig) GetUserLabel() string {
	if o == nil || o.UserLabel == nil {
		var ret string
		return ret
	}
	return *o.UserLabel
}

// GetUserLabelOk returns a tuple with the UserLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeServerConfig) GetUserLabelOk() (*string, bool) {
	if o == nil || o.UserLabel == nil {
		return nil, false
	}
	return o.UserLabel, true
}

// HasUserLabel returns a boolean if a field has been set.
func (o *ComputeServerConfig) HasUserLabel() bool {
	if o != nil && o.UserLabel != nil {
		return true
	}

	return false
}

// SetUserLabel gets a reference to the given string and assigns it to the UserLabel field.
func (o *ComputeServerConfig) SetUserLabel(v string) {
	o.UserLabel = &v
}

func (o ComputeServerConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.AssetTag != nil {
		toSerialize["AssetTag"] = o.AssetTag
	}
	if o.UserLabel != nil {
		toSerialize["UserLabel"] = o.UserLabel
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ComputeServerConfig) UnmarshalJSON(bytes []byte) (err error) {
	varComputeServerConfig := _ComputeServerConfig{}

	if err = json.Unmarshal(bytes, &varComputeServerConfig); err == nil {
		*o = ComputeServerConfig(varComputeServerConfig)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "AssetTag")
		delete(additionalProperties, "UserLabel")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableComputeServerConfig struct {
	value *ComputeServerConfig
	isSet bool
}

func (v NullableComputeServerConfig) Get() *ComputeServerConfig {
	return v.value
}

func (v *NullableComputeServerConfig) Set(val *ComputeServerConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableComputeServerConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableComputeServerConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputeServerConfig(val *ComputeServerConfig) *NullableComputeServerConfig {
	return &NullableComputeServerConfig{value: val, isSet: true}
}

func (v NullableComputeServerConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputeServerConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
