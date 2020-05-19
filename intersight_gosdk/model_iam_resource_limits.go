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

// IamResourceLimits The resource limits used to limit resources such as User accounts.
type IamResourceLimits struct {
	MoBaseMo
	// The maximum number of users allowed in an account. The default value is 200.
	PerAccountUserLimit *int64                  `json:"PerAccountUserLimit,omitempty"`
	Account             *IamAccountRelationship `json:"Account,omitempty"`
}

// NewIamResourceLimits instantiates a new IamResourceLimits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamResourceLimits() *IamResourceLimits {
	this := IamResourceLimits{}
	return &this
}

// NewIamResourceLimitsWithDefaults instantiates a new IamResourceLimits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamResourceLimitsWithDefaults() *IamResourceLimits {
	this := IamResourceLimits{}
	return &this
}

// GetPerAccountUserLimit returns the PerAccountUserLimit field value if set, zero value otherwise.
func (o *IamResourceLimits) GetPerAccountUserLimit() int64 {
	if o == nil || o.PerAccountUserLimit == nil {
		var ret int64
		return ret
	}
	return *o.PerAccountUserLimit
}

// GetPerAccountUserLimitOk returns a tuple with the PerAccountUserLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamResourceLimits) GetPerAccountUserLimitOk() (*int64, bool) {
	if o == nil || o.PerAccountUserLimit == nil {
		return nil, false
	}
	return o.PerAccountUserLimit, true
}

// HasPerAccountUserLimit returns a boolean if a field has been set.
func (o *IamResourceLimits) HasPerAccountUserLimit() bool {
	if o != nil && o.PerAccountUserLimit != nil {
		return true
	}

	return false
}

// SetPerAccountUserLimit gets a reference to the given int64 and assigns it to the PerAccountUserLimit field.
func (o *IamResourceLimits) SetPerAccountUserLimit(v int64) {
	o.PerAccountUserLimit = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *IamResourceLimits) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamResourceLimits) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *IamResourceLimits) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *IamResourceLimits) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o IamResourceLimits) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.PerAccountUserLimit != nil {
		toSerialize["PerAccountUserLimit"] = o.PerAccountUserLimit
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	return json.Marshal(toSerialize)
}

// AsIamResourceLimitsRelationship wraps this instance of IamResourceLimits in IamResourceLimitsRelationship
func (s *IamResourceLimits) AsIamResourceLimitsRelationship() IamResourceLimitsRelationship {
	return IamResourceLimitsRelationship{IamResourceLimitsRelationshipInterface: s}
}

type NullableIamResourceLimits struct {
	value *IamResourceLimits
	isSet bool
}

func (v NullableIamResourceLimits) Get() *IamResourceLimits {
	return v.value
}

func (v *NullableIamResourceLimits) Set(val *IamResourceLimits) {
	v.value = val
	v.isSet = true
}

func (v NullableIamResourceLimits) IsSet() bool {
	return v.isSet
}

func (v *NullableIamResourceLimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamResourceLimits(val *IamResourceLimits) *NullableIamResourceLimits {
	return &NullableIamResourceLimits{value: val, isSet: true}
}

func (v NullableIamResourceLimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamResourceLimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
