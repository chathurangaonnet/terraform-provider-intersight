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

// IamPermissionReference Users can log in through the base URL (https://intersight.com) or account-specific URLs. When the Intersight user logs in through the base URL, Intersight identifies the accounts and permissions within each account which the user has access to. In case multiple permissions are identified, the user and session objects are created in the onboarding-user account, and the session object is updated with account and permission information. Intersight GUI uses this information to show available accounts and permissions for the user to select. PermissionReference type is used to store permission information of an account.
type IamPermissionReference struct {
	MoBaseComplexType
	// MOID of the permission which user has access to.
	PermissionIdentifier *string `json:"PermissionIdentifier,omitempty"`
	// Name of the permission which user has access to.
	PermissionName       *string `json:"PermissionName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamPermissionReference IamPermissionReference

// NewIamPermissionReference instantiates a new IamPermissionReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamPermissionReference() *IamPermissionReference {
	this := IamPermissionReference{}
	return &this
}

// NewIamPermissionReferenceWithDefaults instantiates a new IamPermissionReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamPermissionReferenceWithDefaults() *IamPermissionReference {
	this := IamPermissionReference{}
	return &this
}

// GetPermissionIdentifier returns the PermissionIdentifier field value if set, zero value otherwise.
func (o *IamPermissionReference) GetPermissionIdentifier() string {
	if o == nil || o.PermissionIdentifier == nil {
		var ret string
		return ret
	}
	return *o.PermissionIdentifier
}

// GetPermissionIdentifierOk returns a tuple with the PermissionIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamPermissionReference) GetPermissionIdentifierOk() (*string, bool) {
	if o == nil || o.PermissionIdentifier == nil {
		return nil, false
	}
	return o.PermissionIdentifier, true
}

// HasPermissionIdentifier returns a boolean if a field has been set.
func (o *IamPermissionReference) HasPermissionIdentifier() bool {
	if o != nil && o.PermissionIdentifier != nil {
		return true
	}

	return false
}

// SetPermissionIdentifier gets a reference to the given string and assigns it to the PermissionIdentifier field.
func (o *IamPermissionReference) SetPermissionIdentifier(v string) {
	o.PermissionIdentifier = &v
}

// GetPermissionName returns the PermissionName field value if set, zero value otherwise.
func (o *IamPermissionReference) GetPermissionName() string {
	if o == nil || o.PermissionName == nil {
		var ret string
		return ret
	}
	return *o.PermissionName
}

// GetPermissionNameOk returns a tuple with the PermissionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamPermissionReference) GetPermissionNameOk() (*string, bool) {
	if o == nil || o.PermissionName == nil {
		return nil, false
	}
	return o.PermissionName, true
}

// HasPermissionName returns a boolean if a field has been set.
func (o *IamPermissionReference) HasPermissionName() bool {
	if o != nil && o.PermissionName != nil {
		return true
	}

	return false
}

// SetPermissionName gets a reference to the given string and assigns it to the PermissionName field.
func (o *IamPermissionReference) SetPermissionName(v string) {
	o.PermissionName = &v
}

func (o IamPermissionReference) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.PermissionIdentifier != nil {
		toSerialize["PermissionIdentifier"] = o.PermissionIdentifier
	}
	if o.PermissionName != nil {
		toSerialize["PermissionName"] = o.PermissionName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamPermissionReference) UnmarshalJSON(bytes []byte) (err error) {
	varIamPermissionReference := _IamPermissionReference{}

	if err = json.Unmarshal(bytes, &varIamPermissionReference); err == nil {
		*o = IamPermissionReference(varIamPermissionReference)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "PermissionIdentifier")
		delete(additionalProperties, "PermissionName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIamPermissionReference struct {
	value *IamPermissionReference
	isSet bool
}

func (v NullableIamPermissionReference) Get() *IamPermissionReference {
	return v.value
}

func (v *NullableIamPermissionReference) Set(val *IamPermissionReference) {
	v.value = val
	v.isSet = true
}

func (v NullableIamPermissionReference) IsSet() bool {
	return v.isSet
}

func (v *NullableIamPermissionReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamPermissionReference(val *IamPermissionReference) *NullableIamPermissionReference {
	return &NullableIamPermissionReference{value: val, isSet: true}
}

func (v NullableIamPermissionReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamPermissionReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
