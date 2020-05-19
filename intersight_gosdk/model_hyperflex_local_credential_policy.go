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

// HyperflexLocalCredentialPolicy A policy specifying credentials for HyperFlex cluster such as controller VM password, hypervisor username, and password.
type HyperflexLocalCredentialPolicy struct {
	PolicyAbstractPolicy
	// Indicates if Hypervisor password is the factory set default password. For HyperFlex Data Platform versions 3.0 or higher, enable this if the default password was not changed on the Hypervisor. It is required to supply a new custom Hypervisor password that will be applied to the Hypervisor during deployment. For HyperFlex Data Platform versions prior to 3.0 release, this setting has no effect and the default password will be used for initial install. The Hypervisor password should be changed after deployment.
	FactoryHypervisorPassword *bool `json:"FactoryHypervisorPassword,omitempty"`
	// HyperFlex storage controller VM password must contain a minimum of 10 characters, with at least 1 lowercase, 1 uppercase, 1 numeric, and 1 of these -_@#$%^&*! special characters.
	HxdpRootPwd *string `json:"HxdpRootPwd,omitempty"`
	// Hypervisor administrator username must contain only alphanumeric characters. Use the root account for ESXi deployments.
	HypervisorAdmin *string `json:"HypervisorAdmin,omitempty"`
	// The ESXi root password. For HyperFlex Data Platform 3.0 or later, if the factory default password was not manually changed, you must set a new custom password. If the password was manually changed, you must not enable the factory default password property and provide the current hypervisor password. Note - All HyperFlex nodes require the same hypervisor password for installation. For HyperFlex Data Platform prior to 3.0, use the default password \"Cisco123\" for newly manufactured HyperFlex servers. A custom password should only be entered if hypervisor credentials were manually changed prior to deployment.
	HypervisorAdminPwd *string `json:"HypervisorAdminPwd,omitempty"`
	// Indicates whether the value of the 'hxdpRootPwd' property has been set.
	IsHxdpRootPwdSet *bool `json:"IsHxdpRootPwdSet,omitempty"`
	// Indicates whether the value of the 'hypervisorAdminPwd' property has been set.
	IsHypervisorAdminPwdSet *bool `json:"IsHypervisorAdminPwdSet,omitempty"`
	// An array of relationships to hyperflexClusterProfile resources.
	ClusterProfiles *[]HyperflexClusterProfileRelationship `json:"ClusterProfiles,omitempty"`
	Organization    *OrganizationOrganizationRelationship  `json:"Organization,omitempty"`
}

// NewHyperflexLocalCredentialPolicy instantiates a new HyperflexLocalCredentialPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexLocalCredentialPolicy() *HyperflexLocalCredentialPolicy {
	this := HyperflexLocalCredentialPolicy{}
	return &this
}

// NewHyperflexLocalCredentialPolicyWithDefaults instantiates a new HyperflexLocalCredentialPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexLocalCredentialPolicyWithDefaults() *HyperflexLocalCredentialPolicy {
	this := HyperflexLocalCredentialPolicy{}
	return &this
}

// GetFactoryHypervisorPassword returns the FactoryHypervisorPassword field value if set, zero value otherwise.
func (o *HyperflexLocalCredentialPolicy) GetFactoryHypervisorPassword() bool {
	if o == nil || o.FactoryHypervisorPassword == nil {
		var ret bool
		return ret
	}
	return *o.FactoryHypervisorPassword
}

// GetFactoryHypervisorPasswordOk returns a tuple with the FactoryHypervisorPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexLocalCredentialPolicy) GetFactoryHypervisorPasswordOk() (*bool, bool) {
	if o == nil || o.FactoryHypervisorPassword == nil {
		return nil, false
	}
	return o.FactoryHypervisorPassword, true
}

// HasFactoryHypervisorPassword returns a boolean if a field has been set.
func (o *HyperflexLocalCredentialPolicy) HasFactoryHypervisorPassword() bool {
	if o != nil && o.FactoryHypervisorPassword != nil {
		return true
	}

	return false
}

// SetFactoryHypervisorPassword gets a reference to the given bool and assigns it to the FactoryHypervisorPassword field.
func (o *HyperflexLocalCredentialPolicy) SetFactoryHypervisorPassword(v bool) {
	o.FactoryHypervisorPassword = &v
}

// GetHxdpRootPwd returns the HxdpRootPwd field value if set, zero value otherwise.
func (o *HyperflexLocalCredentialPolicy) GetHxdpRootPwd() string {
	if o == nil || o.HxdpRootPwd == nil {
		var ret string
		return ret
	}
	return *o.HxdpRootPwd
}

// GetHxdpRootPwdOk returns a tuple with the HxdpRootPwd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexLocalCredentialPolicy) GetHxdpRootPwdOk() (*string, bool) {
	if o == nil || o.HxdpRootPwd == nil {
		return nil, false
	}
	return o.HxdpRootPwd, true
}

// HasHxdpRootPwd returns a boolean if a field has been set.
func (o *HyperflexLocalCredentialPolicy) HasHxdpRootPwd() bool {
	if o != nil && o.HxdpRootPwd != nil {
		return true
	}

	return false
}

// SetHxdpRootPwd gets a reference to the given string and assigns it to the HxdpRootPwd field.
func (o *HyperflexLocalCredentialPolicy) SetHxdpRootPwd(v string) {
	o.HxdpRootPwd = &v
}

// GetHypervisorAdmin returns the HypervisorAdmin field value if set, zero value otherwise.
func (o *HyperflexLocalCredentialPolicy) GetHypervisorAdmin() string {
	if o == nil || o.HypervisorAdmin == nil {
		var ret string
		return ret
	}
	return *o.HypervisorAdmin
}

// GetHypervisorAdminOk returns a tuple with the HypervisorAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexLocalCredentialPolicy) GetHypervisorAdminOk() (*string, bool) {
	if o == nil || o.HypervisorAdmin == nil {
		return nil, false
	}
	return o.HypervisorAdmin, true
}

// HasHypervisorAdmin returns a boolean if a field has been set.
func (o *HyperflexLocalCredentialPolicy) HasHypervisorAdmin() bool {
	if o != nil && o.HypervisorAdmin != nil {
		return true
	}

	return false
}

// SetHypervisorAdmin gets a reference to the given string and assigns it to the HypervisorAdmin field.
func (o *HyperflexLocalCredentialPolicy) SetHypervisorAdmin(v string) {
	o.HypervisorAdmin = &v
}

// GetHypervisorAdminPwd returns the HypervisorAdminPwd field value if set, zero value otherwise.
func (o *HyperflexLocalCredentialPolicy) GetHypervisorAdminPwd() string {
	if o == nil || o.HypervisorAdminPwd == nil {
		var ret string
		return ret
	}
	return *o.HypervisorAdminPwd
}

// GetHypervisorAdminPwdOk returns a tuple with the HypervisorAdminPwd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexLocalCredentialPolicy) GetHypervisorAdminPwdOk() (*string, bool) {
	if o == nil || o.HypervisorAdminPwd == nil {
		return nil, false
	}
	return o.HypervisorAdminPwd, true
}

// HasHypervisorAdminPwd returns a boolean if a field has been set.
func (o *HyperflexLocalCredentialPolicy) HasHypervisorAdminPwd() bool {
	if o != nil && o.HypervisorAdminPwd != nil {
		return true
	}

	return false
}

// SetHypervisorAdminPwd gets a reference to the given string and assigns it to the HypervisorAdminPwd field.
func (o *HyperflexLocalCredentialPolicy) SetHypervisorAdminPwd(v string) {
	o.HypervisorAdminPwd = &v
}

// GetIsHxdpRootPwdSet returns the IsHxdpRootPwdSet field value if set, zero value otherwise.
func (o *HyperflexLocalCredentialPolicy) GetIsHxdpRootPwdSet() bool {
	if o == nil || o.IsHxdpRootPwdSet == nil {
		var ret bool
		return ret
	}
	return *o.IsHxdpRootPwdSet
}

// GetIsHxdpRootPwdSetOk returns a tuple with the IsHxdpRootPwdSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexLocalCredentialPolicy) GetIsHxdpRootPwdSetOk() (*bool, bool) {
	if o == nil || o.IsHxdpRootPwdSet == nil {
		return nil, false
	}
	return o.IsHxdpRootPwdSet, true
}

// HasIsHxdpRootPwdSet returns a boolean if a field has been set.
func (o *HyperflexLocalCredentialPolicy) HasIsHxdpRootPwdSet() bool {
	if o != nil && o.IsHxdpRootPwdSet != nil {
		return true
	}

	return false
}

// SetIsHxdpRootPwdSet gets a reference to the given bool and assigns it to the IsHxdpRootPwdSet field.
func (o *HyperflexLocalCredentialPolicy) SetIsHxdpRootPwdSet(v bool) {
	o.IsHxdpRootPwdSet = &v
}

// GetIsHypervisorAdminPwdSet returns the IsHypervisorAdminPwdSet field value if set, zero value otherwise.
func (o *HyperflexLocalCredentialPolicy) GetIsHypervisorAdminPwdSet() bool {
	if o == nil || o.IsHypervisorAdminPwdSet == nil {
		var ret bool
		return ret
	}
	return *o.IsHypervisorAdminPwdSet
}

// GetIsHypervisorAdminPwdSetOk returns a tuple with the IsHypervisorAdminPwdSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexLocalCredentialPolicy) GetIsHypervisorAdminPwdSetOk() (*bool, bool) {
	if o == nil || o.IsHypervisorAdminPwdSet == nil {
		return nil, false
	}
	return o.IsHypervisorAdminPwdSet, true
}

// HasIsHypervisorAdminPwdSet returns a boolean if a field has been set.
func (o *HyperflexLocalCredentialPolicy) HasIsHypervisorAdminPwdSet() bool {
	if o != nil && o.IsHypervisorAdminPwdSet != nil {
		return true
	}

	return false
}

// SetIsHypervisorAdminPwdSet gets a reference to the given bool and assigns it to the IsHypervisorAdminPwdSet field.
func (o *HyperflexLocalCredentialPolicy) SetIsHypervisorAdminPwdSet(v bool) {
	o.IsHypervisorAdminPwdSet = &v
}

// GetClusterProfiles returns the ClusterProfiles field value if set, zero value otherwise.
func (o *HyperflexLocalCredentialPolicy) GetClusterProfiles() []HyperflexClusterProfileRelationship {
	if o == nil || o.ClusterProfiles == nil {
		var ret []HyperflexClusterProfileRelationship
		return ret
	}
	return *o.ClusterProfiles
}

// GetClusterProfilesOk returns a tuple with the ClusterProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexLocalCredentialPolicy) GetClusterProfilesOk() (*[]HyperflexClusterProfileRelationship, bool) {
	if o == nil || o.ClusterProfiles == nil {
		return nil, false
	}
	return o.ClusterProfiles, true
}

// HasClusterProfiles returns a boolean if a field has been set.
func (o *HyperflexLocalCredentialPolicy) HasClusterProfiles() bool {
	if o != nil && o.ClusterProfiles != nil {
		return true
	}

	return false
}

// SetClusterProfiles gets a reference to the given []HyperflexClusterProfileRelationship and assigns it to the ClusterProfiles field.
func (o *HyperflexLocalCredentialPolicy) SetClusterProfiles(v []HyperflexClusterProfileRelationship) {
	o.ClusterProfiles = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *HyperflexLocalCredentialPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexLocalCredentialPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *HyperflexLocalCredentialPolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *HyperflexLocalCredentialPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o HyperflexLocalCredentialPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	if o.FactoryHypervisorPassword != nil {
		toSerialize["FactoryHypervisorPassword"] = o.FactoryHypervisorPassword
	}
	if o.HxdpRootPwd != nil {
		toSerialize["HxdpRootPwd"] = o.HxdpRootPwd
	}
	if o.HypervisorAdmin != nil {
		toSerialize["HypervisorAdmin"] = o.HypervisorAdmin
	}
	if o.HypervisorAdminPwd != nil {
		toSerialize["HypervisorAdminPwd"] = o.HypervisorAdminPwd
	}
	if o.IsHxdpRootPwdSet != nil {
		toSerialize["IsHxdpRootPwdSet"] = o.IsHxdpRootPwdSet
	}
	if o.IsHypervisorAdminPwdSet != nil {
		toSerialize["IsHypervisorAdminPwdSet"] = o.IsHypervisorAdminPwdSet
	}
	if o.ClusterProfiles != nil {
		toSerialize["ClusterProfiles"] = o.ClusterProfiles
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	return json.Marshal(toSerialize)
}

// AsHyperflexLocalCredentialPolicyRelationship wraps this instance of HyperflexLocalCredentialPolicy in HyperflexLocalCredentialPolicyRelationship
func (s *HyperflexLocalCredentialPolicy) AsHyperflexLocalCredentialPolicyRelationship() HyperflexLocalCredentialPolicyRelationship {
	return HyperflexLocalCredentialPolicyRelationship{HyperflexLocalCredentialPolicyRelationshipInterface: s}
}

type NullableHyperflexLocalCredentialPolicy struct {
	value *HyperflexLocalCredentialPolicy
	isSet bool
}

func (v NullableHyperflexLocalCredentialPolicy) Get() *HyperflexLocalCredentialPolicy {
	return v.value
}

func (v *NullableHyperflexLocalCredentialPolicy) Set(val *HyperflexLocalCredentialPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexLocalCredentialPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexLocalCredentialPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexLocalCredentialPolicy(val *HyperflexLocalCredentialPolicy) *NullableHyperflexLocalCredentialPolicy {
	return &NullableHyperflexLocalCredentialPolicy{value: val, isSet: true}
}

func (v NullableHyperflexLocalCredentialPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexLocalCredentialPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
