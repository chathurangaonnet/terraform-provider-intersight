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

// TestcryptTokenApiAllOf Definition of the list of properties defined in 'testcrypt.TokenApi', excluding properties defined in parent classes.
type TestcryptTokenApiAllOf struct {
	// The session identifier to be used to create/read/update or delete Encryption token. This could be a user's (web) session id or api key id, etc.
	SessionId *string                 `json:"SessionId,omitempty"`
	Account   *IamAccountRelationship `json:"Account,omitempty"`
}

// NewTestcryptTokenApiAllOf instantiates a new TestcryptTokenApiAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestcryptTokenApiAllOf() *TestcryptTokenApiAllOf {
	this := TestcryptTokenApiAllOf{}
	return &this
}

// NewTestcryptTokenApiAllOfWithDefaults instantiates a new TestcryptTokenApiAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestcryptTokenApiAllOfWithDefaults() *TestcryptTokenApiAllOf {
	this := TestcryptTokenApiAllOf{}
	return &this
}

// GetSessionId returns the SessionId field value if set, zero value otherwise.
func (o *TestcryptTokenApiAllOf) GetSessionId() string {
	if o == nil || o.SessionId == nil {
		var ret string
		return ret
	}
	return *o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestcryptTokenApiAllOf) GetSessionIdOk() (*string, bool) {
	if o == nil || o.SessionId == nil {
		return nil, false
	}
	return o.SessionId, true
}

// HasSessionId returns a boolean if a field has been set.
func (o *TestcryptTokenApiAllOf) HasSessionId() bool {
	if o != nil && o.SessionId != nil {
		return true
	}

	return false
}

// SetSessionId gets a reference to the given string and assigns it to the SessionId field.
func (o *TestcryptTokenApiAllOf) SetSessionId(v string) {
	o.SessionId = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *TestcryptTokenApiAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestcryptTokenApiAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *TestcryptTokenApiAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *TestcryptTokenApiAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o TestcryptTokenApiAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SessionId != nil {
		toSerialize["SessionId"] = o.SessionId
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	return json.Marshal(toSerialize)
}

type NullableTestcryptTokenApiAllOf struct {
	value *TestcryptTokenApiAllOf
	isSet bool
}

func (v NullableTestcryptTokenApiAllOf) Get() *TestcryptTokenApiAllOf {
	return v.value
}

func (v *NullableTestcryptTokenApiAllOf) Set(val *TestcryptTokenApiAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTestcryptTokenApiAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTestcryptTokenApiAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestcryptTokenApiAllOf(val *TestcryptTokenApiAllOf) *NullableTestcryptTokenApiAllOf {
	return &NullableTestcryptTokenApiAllOf{value: val, isSet: true}
}

func (v NullableTestcryptTokenApiAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestcryptTokenApiAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
