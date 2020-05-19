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

// IamEndPointRole The role defined in the end point which can be assigned to a user.
type IamEndPointRole struct {
	MoBaseMo
	// The name of the end point role.
	Name *string `json:"Name,omitempty"`
	// User specified tags to roles like as ep-admin or ep-readonly.
	RoleType *string `json:"RoleType,omitempty"`
	// The type of the end point like Cisco UCS Fabric Interconnect or Cisco IMC.
	Type    *string                 `json:"Type,omitempty"`
	Account *IamAccountRelationship `json:"Account,omitempty"`
	// An array of relationships to iamEndPointPrivilege resources.
	EndPointPrivileges *[]IamEndPointPrivilegeRelationship `json:"EndPointPrivileges,omitempty"`
	System             *IamSystemRelationship              `json:"System,omitempty"`
}

// NewIamEndPointRole instantiates a new IamEndPointRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamEndPointRole() *IamEndPointRole {
	this := IamEndPointRole{}
	var type_ string = ""
	this.Type = &type_
	return &this
}

// NewIamEndPointRoleWithDefaults instantiates a new IamEndPointRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamEndPointRoleWithDefaults() *IamEndPointRole {
	this := IamEndPointRole{}
	var type_ string = ""
	this.Type = &type_
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IamEndPointRole) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamEndPointRole) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IamEndPointRole) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IamEndPointRole) SetName(v string) {
	o.Name = &v
}

// GetRoleType returns the RoleType field value if set, zero value otherwise.
func (o *IamEndPointRole) GetRoleType() string {
	if o == nil || o.RoleType == nil {
		var ret string
		return ret
	}
	return *o.RoleType
}

// GetRoleTypeOk returns a tuple with the RoleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamEndPointRole) GetRoleTypeOk() (*string, bool) {
	if o == nil || o.RoleType == nil {
		return nil, false
	}
	return o.RoleType, true
}

// HasRoleType returns a boolean if a field has been set.
func (o *IamEndPointRole) HasRoleType() bool {
	if o != nil && o.RoleType != nil {
		return true
	}

	return false
}

// SetRoleType gets a reference to the given string and assigns it to the RoleType field.
func (o *IamEndPointRole) SetRoleType(v string) {
	o.RoleType = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IamEndPointRole) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamEndPointRole) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IamEndPointRole) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IamEndPointRole) SetType(v string) {
	o.Type = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *IamEndPointRole) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamEndPointRole) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *IamEndPointRole) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *IamEndPointRole) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

// GetEndPointPrivileges returns the EndPointPrivileges field value if set, zero value otherwise.
func (o *IamEndPointRole) GetEndPointPrivileges() []IamEndPointPrivilegeRelationship {
	if o == nil || o.EndPointPrivileges == nil {
		var ret []IamEndPointPrivilegeRelationship
		return ret
	}
	return *o.EndPointPrivileges
}

// GetEndPointPrivilegesOk returns a tuple with the EndPointPrivileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamEndPointRole) GetEndPointPrivilegesOk() (*[]IamEndPointPrivilegeRelationship, bool) {
	if o == nil || o.EndPointPrivileges == nil {
		return nil, false
	}
	return o.EndPointPrivileges, true
}

// HasEndPointPrivileges returns a boolean if a field has been set.
func (o *IamEndPointRole) HasEndPointPrivileges() bool {
	if o != nil && o.EndPointPrivileges != nil {
		return true
	}

	return false
}

// SetEndPointPrivileges gets a reference to the given []IamEndPointPrivilegeRelationship and assigns it to the EndPointPrivileges field.
func (o *IamEndPointRole) SetEndPointPrivileges(v []IamEndPointPrivilegeRelationship) {
	o.EndPointPrivileges = &v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *IamEndPointRole) GetSystem() IamSystemRelationship {
	if o == nil || o.System == nil {
		var ret IamSystemRelationship
		return ret
	}
	return *o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamEndPointRole) GetSystemOk() (*IamSystemRelationship, bool) {
	if o == nil || o.System == nil {
		return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *IamEndPointRole) HasSystem() bool {
	if o != nil && o.System != nil {
		return true
	}

	return false
}

// SetSystem gets a reference to the given IamSystemRelationship and assigns it to the System field.
func (o *IamEndPointRole) SetSystem(v IamSystemRelationship) {
	o.System = &v
}

func (o IamEndPointRole) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.RoleType != nil {
		toSerialize["RoleType"] = o.RoleType
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.EndPointPrivileges != nil {
		toSerialize["EndPointPrivileges"] = o.EndPointPrivileges
	}
	if o.System != nil {
		toSerialize["System"] = o.System
	}
	return json.Marshal(toSerialize)
}

// AsIamEndPointRoleRelationship wraps this instance of IamEndPointRole in IamEndPointRoleRelationship
func (s *IamEndPointRole) AsIamEndPointRoleRelationship() IamEndPointRoleRelationship {
	return IamEndPointRoleRelationship{IamEndPointRoleRelationshipInterface: s}
}

type NullableIamEndPointRole struct {
	value *IamEndPointRole
	isSet bool
}

func (v NullableIamEndPointRole) Get() *IamEndPointRole {
	return v.value
}

func (v *NullableIamEndPointRole) Set(val *IamEndPointRole) {
	v.value = val
	v.isSet = true
}

func (v NullableIamEndPointRole) IsSet() bool {
	return v.isSet
}

func (v *NullableIamEndPointRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamEndPointRole(val *IamEndPointRole) *NullableIamEndPointRole {
	return &NullableIamEndPointRole{value: val, isSet: true}
}

func (v NullableIamEndPointRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamEndPointRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
