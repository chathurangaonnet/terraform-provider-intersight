/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-04-12T21:47:47-07:00.
 *
 * API version: 1.0.9-1617
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
)

// IamClientMeta Contains meta information of the client. Currently it is used only within OAuth clients.
type IamClientMeta struct {
	MoBaseComplexType
	// Parsed device model from raw User-Agent.
	DeviceModel *string `json:"DeviceModel,omitempty"`
	// The value of the \"User-Agent\" HTTP header, as sent by the HTTP client when it initiate a session to Intersight. This can be used to identify the client operating system, browser type and browser version. Example - Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36 It is set when User successfully passed OAuth login flow and receives Access Token.
	UserAgent *string `json:"UserAgent,omitempty"`
}

// NewIamClientMeta instantiates a new IamClientMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamClientMeta() *IamClientMeta {
	this := IamClientMeta{}
	return &this
}

// NewIamClientMetaWithDefaults instantiates a new IamClientMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamClientMetaWithDefaults() *IamClientMeta {
	this := IamClientMeta{}
	return &this
}

// GetDeviceModel returns the DeviceModel field value if set, zero value otherwise.
func (o *IamClientMeta) GetDeviceModel() string {
	if o == nil || o.DeviceModel == nil {
		var ret string
		return ret
	}
	return *o.DeviceModel
}

// GetDeviceModelOk returns a tuple with the DeviceModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamClientMeta) GetDeviceModelOk() (*string, bool) {
	if o == nil || o.DeviceModel == nil {
		return nil, false
	}
	return o.DeviceModel, true
}

// HasDeviceModel returns a boolean if a field has been set.
func (o *IamClientMeta) HasDeviceModel() bool {
	if o != nil && o.DeviceModel != nil {
		return true
	}

	return false
}

// SetDeviceModel gets a reference to the given string and assigns it to the DeviceModel field.
func (o *IamClientMeta) SetDeviceModel(v string) {
	o.DeviceModel = &v
}

// GetUserAgent returns the UserAgent field value if set, zero value otherwise.
func (o *IamClientMeta) GetUserAgent() string {
	if o == nil || o.UserAgent == nil {
		var ret string
		return ret
	}
	return *o.UserAgent
}

// GetUserAgentOk returns a tuple with the UserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamClientMeta) GetUserAgentOk() (*string, bool) {
	if o == nil || o.UserAgent == nil {
		return nil, false
	}
	return o.UserAgent, true
}

// HasUserAgent returns a boolean if a field has been set.
func (o *IamClientMeta) HasUserAgent() bool {
	if o != nil && o.UserAgent != nil {
		return true
	}

	return false
}

// SetUserAgent gets a reference to the given string and assigns it to the UserAgent field.
func (o *IamClientMeta) SetUserAgent(v string) {
	o.UserAgent = &v
}

func (o IamClientMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.DeviceModel != nil {
		toSerialize["DeviceModel"] = o.DeviceModel
	}
	if o.UserAgent != nil {
		toSerialize["UserAgent"] = o.UserAgent
	}
	return json.Marshal(toSerialize)
}

type NullableIamClientMeta struct {
	value *IamClientMeta
	isSet bool
}

func (v NullableIamClientMeta) Get() *IamClientMeta {
	return v.value
}

func (v *NullableIamClientMeta) Set(val *IamClientMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableIamClientMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableIamClientMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamClientMeta(val *IamClientMeta) *NullableIamClientMeta {
	return &NullableIamClientMeta{value: val, isSet: true}
}

func (v NullableIamClientMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamClientMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
