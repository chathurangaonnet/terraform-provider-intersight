/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-07-11T05:47:33Z.
 *
 * API version: 1.0.9-1999
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
)

// IamEndPointRoleAllOf Definition of the list of properties defined in 'iam.EndPointRole', excluding properties defined in parent classes.
type IamEndPointRoleAllOf struct {
	// The name of the end point role.
	Name *string `json:"Name,omitempty"`
	// User specified tags to roles like as ep-admin or ep-readonly.
	RoleType *string `json:"RoleType,omitempty"`
	// The type of the end point like Cisco UCS Fabric Interconnect or Cisco IMC.
	Type    *string                 `json:"Type,omitempty"`
	Account *IamAccountRelationship `json:"Account,omitempty"`
	// An array of relationships to iamEndPointPrivilege resources.
	EndPointPrivileges   []IamEndPointPrivilegeRelationship `json:"EndPointPrivileges,omitempty"`
	System               *IamSystemRelationship             `json:"System,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamEndPointRoleAllOf IamEndPointRoleAllOf

// NewIamEndPointRoleAllOf instantiates a new IamEndPointRoleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamEndPointRoleAllOf() *IamEndPointRoleAllOf {
	this := IamEndPointRoleAllOf{}
	var type_ string = ""
	this.Type = &type_
	return &this
}

// NewIamEndPointRoleAllOfWithDefaults instantiates a new IamEndPointRoleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamEndPointRoleAllOfWithDefaults() *IamEndPointRoleAllOf {
	this := IamEndPointRoleAllOf{}
	var type_ string = ""
	this.Type = &type_
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IamEndPointRoleAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamEndPointRoleAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IamEndPointRoleAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IamEndPointRoleAllOf) SetName(v string) {
	o.Name = &v
}

// GetRoleType returns the RoleType field value if set, zero value otherwise.
func (o *IamEndPointRoleAllOf) GetRoleType() string {
	if o == nil || o.RoleType == nil {
		var ret string
		return ret
	}
	return *o.RoleType
}

// GetRoleTypeOk returns a tuple with the RoleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamEndPointRoleAllOf) GetRoleTypeOk() (*string, bool) {
	if o == nil || o.RoleType == nil {
		return nil, false
	}
	return o.RoleType, true
}

// HasRoleType returns a boolean if a field has been set.
func (o *IamEndPointRoleAllOf) HasRoleType() bool {
	if o != nil && o.RoleType != nil {
		return true
	}

	return false
}

// SetRoleType gets a reference to the given string and assigns it to the RoleType field.
func (o *IamEndPointRoleAllOf) SetRoleType(v string) {
	o.RoleType = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IamEndPointRoleAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamEndPointRoleAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IamEndPointRoleAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IamEndPointRoleAllOf) SetType(v string) {
	o.Type = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *IamEndPointRoleAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamEndPointRoleAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *IamEndPointRoleAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *IamEndPointRoleAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

// GetEndPointPrivileges returns the EndPointPrivileges field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamEndPointRoleAllOf) GetEndPointPrivileges() []IamEndPointPrivilegeRelationship {
	if o == nil {
		var ret []IamEndPointPrivilegeRelationship
		return ret
	}
	return o.EndPointPrivileges
}

// GetEndPointPrivilegesOk returns a tuple with the EndPointPrivileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamEndPointRoleAllOf) GetEndPointPrivilegesOk() (*[]IamEndPointPrivilegeRelationship, bool) {
	if o == nil || o.EndPointPrivileges == nil {
		return nil, false
	}
	return &o.EndPointPrivileges, true
}

// HasEndPointPrivileges returns a boolean if a field has been set.
func (o *IamEndPointRoleAllOf) HasEndPointPrivileges() bool {
	if o != nil && o.EndPointPrivileges != nil {
		return true
	}

	return false
}

// SetEndPointPrivileges gets a reference to the given []IamEndPointPrivilegeRelationship and assigns it to the EndPointPrivileges field.
func (o *IamEndPointRoleAllOf) SetEndPointPrivileges(v []IamEndPointPrivilegeRelationship) {
	o.EndPointPrivileges = v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *IamEndPointRoleAllOf) GetSystem() IamSystemRelationship {
	if o == nil || o.System == nil {
		var ret IamSystemRelationship
		return ret
	}
	return *o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamEndPointRoleAllOf) GetSystemOk() (*IamSystemRelationship, bool) {
	if o == nil || o.System == nil {
		return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *IamEndPointRoleAllOf) HasSystem() bool {
	if o != nil && o.System != nil {
		return true
	}

	return false
}

// SetSystem gets a reference to the given IamSystemRelationship and assigns it to the System field.
func (o *IamEndPointRoleAllOf) SetSystem(v IamSystemRelationship) {
	o.System = &v
}

func (o IamEndPointRoleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamEndPointRoleAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIamEndPointRoleAllOf := _IamEndPointRoleAllOf{}

	if err = json.Unmarshal(bytes, &varIamEndPointRoleAllOf); err == nil {
		*o = IamEndPointRoleAllOf(varIamEndPointRoleAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Name")
		delete(additionalProperties, "RoleType")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "Account")
		delete(additionalProperties, "EndPointPrivileges")
		delete(additionalProperties, "System")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIamEndPointRoleAllOf struct {
	value *IamEndPointRoleAllOf
	isSet bool
}

func (v NullableIamEndPointRoleAllOf) Get() *IamEndPointRoleAllOf {
	return v.value
}

func (v *NullableIamEndPointRoleAllOf) Set(val *IamEndPointRoleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIamEndPointRoleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIamEndPointRoleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamEndPointRoleAllOf(val *IamEndPointRoleAllOf) *NullableIamEndPointRoleAllOf {
	return &NullableIamEndPointRoleAllOf{value: val, isSet: true}
}

func (v NullableIamEndPointRoleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamEndPointRoleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
