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

// IamPermissionToRolesAllOf Definition of the list of properties defined in 'iam.PermissionToRoles', excluding properties defined in parent classes.
type IamPermissionToRolesAllOf struct {
	Permission *CmrfCmRf   `json:"Permission,omitempty"`
	Roles      *[]CmrfCmRf `json:"Roles,omitempty"`
}

// NewIamPermissionToRolesAllOf instantiates a new IamPermissionToRolesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamPermissionToRolesAllOf() *IamPermissionToRolesAllOf {
	this := IamPermissionToRolesAllOf{}
	return &this
}

// NewIamPermissionToRolesAllOfWithDefaults instantiates a new IamPermissionToRolesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamPermissionToRolesAllOfWithDefaults() *IamPermissionToRolesAllOf {
	this := IamPermissionToRolesAllOf{}
	return &this
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *IamPermissionToRolesAllOf) GetPermission() CmrfCmRf {
	if o == nil || o.Permission == nil {
		var ret CmrfCmRf
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamPermissionToRolesAllOf) GetPermissionOk() (*CmrfCmRf, bool) {
	if o == nil || o.Permission == nil {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *IamPermissionToRolesAllOf) HasPermission() bool {
	if o != nil && o.Permission != nil {
		return true
	}

	return false
}

// SetPermission gets a reference to the given CmrfCmRf and assigns it to the Permission field.
func (o *IamPermissionToRolesAllOf) SetPermission(v CmrfCmRf) {
	o.Permission = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *IamPermissionToRolesAllOf) GetRoles() []CmrfCmRf {
	if o == nil || o.Roles == nil {
		var ret []CmrfCmRf
		return ret
	}
	return *o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamPermissionToRolesAllOf) GetRolesOk() (*[]CmrfCmRf, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *IamPermissionToRolesAllOf) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []CmrfCmRf and assigns it to the Roles field.
func (o *IamPermissionToRolesAllOf) SetRoles(v []CmrfCmRf) {
	o.Roles = &v
}

func (o IamPermissionToRolesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Permission != nil {
		toSerialize["Permission"] = o.Permission
	}
	if o.Roles != nil {
		toSerialize["Roles"] = o.Roles
	}
	return json.Marshal(toSerialize)
}

type NullableIamPermissionToRolesAllOf struct {
	value *IamPermissionToRolesAllOf
	isSet bool
}

func (v NullableIamPermissionToRolesAllOf) Get() *IamPermissionToRolesAllOf {
	return v.value
}

func (v *NullableIamPermissionToRolesAllOf) Set(val *IamPermissionToRolesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIamPermissionToRolesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIamPermissionToRolesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamPermissionToRolesAllOf(val *IamPermissionToRolesAllOf) *NullableIamPermissionToRolesAllOf {
	return &NullableIamPermissionToRolesAllOf{value: val, isSet: true}
}

func (v NullableIamPermissionToRolesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamPermissionToRolesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
