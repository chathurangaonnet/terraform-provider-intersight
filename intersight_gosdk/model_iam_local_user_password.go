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

// IamLocalUserPassword LocalUserPassword type is used for changing local user's password. Caller must send old password in Password field and new password in newPassword field. Intersight will verify the old password and sets the new password if everything is OK. This API must not be used for resetting user's password.
type IamLocalUserPassword struct {
	MoBaseMo
	// User-entered passsord to be compared to password for change password function.
	CurrentPassword *string `json:"CurrentPassword,omitempty"`
	// New password that the user's password should be changed to.
	NewPassword *string `json:"NewPassword,omitempty"`
	// User's current valid passsord.
	Password *string              `json:"Password,omitempty"`
	User     *IamUserRelationship `json:"User,omitempty"`
}

// NewIamLocalUserPassword instantiates a new IamLocalUserPassword object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamLocalUserPassword() *IamLocalUserPassword {
	this := IamLocalUserPassword{}
	return &this
}

// NewIamLocalUserPasswordWithDefaults instantiates a new IamLocalUserPassword object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamLocalUserPasswordWithDefaults() *IamLocalUserPassword {
	this := IamLocalUserPassword{}
	return &this
}

// GetCurrentPassword returns the CurrentPassword field value if set, zero value otherwise.
func (o *IamLocalUserPassword) GetCurrentPassword() string {
	if o == nil || o.CurrentPassword == nil {
		var ret string
		return ret
	}
	return *o.CurrentPassword
}

// GetCurrentPasswordOk returns a tuple with the CurrentPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLocalUserPassword) GetCurrentPasswordOk() (*string, bool) {
	if o == nil || o.CurrentPassword == nil {
		return nil, false
	}
	return o.CurrentPassword, true
}

// HasCurrentPassword returns a boolean if a field has been set.
func (o *IamLocalUserPassword) HasCurrentPassword() bool {
	if o != nil && o.CurrentPassword != nil {
		return true
	}

	return false
}

// SetCurrentPassword gets a reference to the given string and assigns it to the CurrentPassword field.
func (o *IamLocalUserPassword) SetCurrentPassword(v string) {
	o.CurrentPassword = &v
}

// GetNewPassword returns the NewPassword field value if set, zero value otherwise.
func (o *IamLocalUserPassword) GetNewPassword() string {
	if o == nil || o.NewPassword == nil {
		var ret string
		return ret
	}
	return *o.NewPassword
}

// GetNewPasswordOk returns a tuple with the NewPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLocalUserPassword) GetNewPasswordOk() (*string, bool) {
	if o == nil || o.NewPassword == nil {
		return nil, false
	}
	return o.NewPassword, true
}

// HasNewPassword returns a boolean if a field has been set.
func (o *IamLocalUserPassword) HasNewPassword() bool {
	if o != nil && o.NewPassword != nil {
		return true
	}

	return false
}

// SetNewPassword gets a reference to the given string and assigns it to the NewPassword field.
func (o *IamLocalUserPassword) SetNewPassword(v string) {
	o.NewPassword = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *IamLocalUserPassword) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLocalUserPassword) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *IamLocalUserPassword) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *IamLocalUserPassword) SetPassword(v string) {
	o.Password = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *IamLocalUserPassword) GetUser() IamUserRelationship {
	if o == nil || o.User == nil {
		var ret IamUserRelationship
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLocalUserPassword) GetUserOk() (*IamUserRelationship, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *IamLocalUserPassword) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given IamUserRelationship and assigns it to the User field.
func (o *IamLocalUserPassword) SetUser(v IamUserRelationship) {
	o.User = &v
}

func (o IamLocalUserPassword) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.CurrentPassword != nil {
		toSerialize["CurrentPassword"] = o.CurrentPassword
	}
	if o.NewPassword != nil {
		toSerialize["NewPassword"] = o.NewPassword
	}
	if o.Password != nil {
		toSerialize["Password"] = o.Password
	}
	if o.User != nil {
		toSerialize["User"] = o.User
	}
	return json.Marshal(toSerialize)
}

// AsIamLocalUserPasswordRelationship wraps this instance of IamLocalUserPassword in IamLocalUserPasswordRelationship
func (s *IamLocalUserPassword) AsIamLocalUserPasswordRelationship() IamLocalUserPasswordRelationship {
	return IamLocalUserPasswordRelationship{IamLocalUserPasswordRelationshipInterface: s}
}

type NullableIamLocalUserPassword struct {
	value *IamLocalUserPassword
	isSet bool
}

func (v NullableIamLocalUserPassword) Get() *IamLocalUserPassword {
	return v.value
}

func (v *NullableIamLocalUserPassword) Set(val *IamLocalUserPassword) {
	v.value = val
	v.isSet = true
}

func (v NullableIamLocalUserPassword) IsSet() bool {
	return v.isSet
}

func (v *NullableIamLocalUserPassword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamLocalUserPassword(val *IamLocalUserPassword) *NullableIamLocalUserPassword {
	return &NullableIamLocalUserPassword{value: val, isSet: true}
}

func (v NullableIamLocalUserPassword) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamLocalUserPassword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
