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

// HyperflexVcenterConfigPolicy A policy specifying vCenter configuration.
type HyperflexVcenterConfigPolicy struct {
	PolicyAbstractPolicy
	// The vCenter datacenter name.
	DataCenter *string `json:"DataCenter,omitempty"`
	// The vCenter server FQDN or IP.
	Hostname *string `json:"Hostname,omitempty"`
	// Indicates whether the value of the 'password' property has been set.
	IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
	// The password for authenticating with vCenter. Follow the corresponding password policy governed by vCenter.
	Password *string `json:"Password,omitempty"`
	// Overrides the default vCenter Single Sign-On URL. Do not specify unless instructed by Cisco TAC.
	SsoUrl *string `json:"SsoUrl,omitempty"`
	// The vCenter username (e.g. administrator@vsphere.local).
	Username *string `json:"Username,omitempty"`
	// An array of relationships to hyperflexClusterProfile resources.
	ClusterProfiles *[]HyperflexClusterProfileRelationship `json:"ClusterProfiles,omitempty"`
	Organization    *OrganizationOrganizationRelationship  `json:"Organization,omitempty"`
}

// NewHyperflexVcenterConfigPolicy instantiates a new HyperflexVcenterConfigPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexVcenterConfigPolicy() *HyperflexVcenterConfigPolicy {
	this := HyperflexVcenterConfigPolicy{}
	return &this
}

// NewHyperflexVcenterConfigPolicyWithDefaults instantiates a new HyperflexVcenterConfigPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexVcenterConfigPolicyWithDefaults() *HyperflexVcenterConfigPolicy {
	this := HyperflexVcenterConfigPolicy{}
	return &this
}

// GetDataCenter returns the DataCenter field value if set, zero value otherwise.
func (o *HyperflexVcenterConfigPolicy) GetDataCenter() string {
	if o == nil || o.DataCenter == nil {
		var ret string
		return ret
	}
	return *o.DataCenter
}

// GetDataCenterOk returns a tuple with the DataCenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVcenterConfigPolicy) GetDataCenterOk() (*string, bool) {
	if o == nil || o.DataCenter == nil {
		return nil, false
	}
	return o.DataCenter, true
}

// HasDataCenter returns a boolean if a field has been set.
func (o *HyperflexVcenterConfigPolicy) HasDataCenter() bool {
	if o != nil && o.DataCenter != nil {
		return true
	}

	return false
}

// SetDataCenter gets a reference to the given string and assigns it to the DataCenter field.
func (o *HyperflexVcenterConfigPolicy) SetDataCenter(v string) {
	o.DataCenter = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *HyperflexVcenterConfigPolicy) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVcenterConfigPolicy) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *HyperflexVcenterConfigPolicy) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *HyperflexVcenterConfigPolicy) SetHostname(v string) {
	o.Hostname = &v
}

// GetIsPasswordSet returns the IsPasswordSet field value if set, zero value otherwise.
func (o *HyperflexVcenterConfigPolicy) GetIsPasswordSet() bool {
	if o == nil || o.IsPasswordSet == nil {
		var ret bool
		return ret
	}
	return *o.IsPasswordSet
}

// GetIsPasswordSetOk returns a tuple with the IsPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVcenterConfigPolicy) GetIsPasswordSetOk() (*bool, bool) {
	if o == nil || o.IsPasswordSet == nil {
		return nil, false
	}
	return o.IsPasswordSet, true
}

// HasIsPasswordSet returns a boolean if a field has been set.
func (o *HyperflexVcenterConfigPolicy) HasIsPasswordSet() bool {
	if o != nil && o.IsPasswordSet != nil {
		return true
	}

	return false
}

// SetIsPasswordSet gets a reference to the given bool and assigns it to the IsPasswordSet field.
func (o *HyperflexVcenterConfigPolicy) SetIsPasswordSet(v bool) {
	o.IsPasswordSet = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *HyperflexVcenterConfigPolicy) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVcenterConfigPolicy) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *HyperflexVcenterConfigPolicy) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *HyperflexVcenterConfigPolicy) SetPassword(v string) {
	o.Password = &v
}

// GetSsoUrl returns the SsoUrl field value if set, zero value otherwise.
func (o *HyperflexVcenterConfigPolicy) GetSsoUrl() string {
	if o == nil || o.SsoUrl == nil {
		var ret string
		return ret
	}
	return *o.SsoUrl
}

// GetSsoUrlOk returns a tuple with the SsoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVcenterConfigPolicy) GetSsoUrlOk() (*string, bool) {
	if o == nil || o.SsoUrl == nil {
		return nil, false
	}
	return o.SsoUrl, true
}

// HasSsoUrl returns a boolean if a field has been set.
func (o *HyperflexVcenterConfigPolicy) HasSsoUrl() bool {
	if o != nil && o.SsoUrl != nil {
		return true
	}

	return false
}

// SetSsoUrl gets a reference to the given string and assigns it to the SsoUrl field.
func (o *HyperflexVcenterConfigPolicy) SetSsoUrl(v string) {
	o.SsoUrl = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *HyperflexVcenterConfigPolicy) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVcenterConfigPolicy) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *HyperflexVcenterConfigPolicy) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *HyperflexVcenterConfigPolicy) SetUsername(v string) {
	o.Username = &v
}

// GetClusterProfiles returns the ClusterProfiles field value if set, zero value otherwise.
func (o *HyperflexVcenterConfigPolicy) GetClusterProfiles() []HyperflexClusterProfileRelationship {
	if o == nil || o.ClusterProfiles == nil {
		var ret []HyperflexClusterProfileRelationship
		return ret
	}
	return *o.ClusterProfiles
}

// GetClusterProfilesOk returns a tuple with the ClusterProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVcenterConfigPolicy) GetClusterProfilesOk() (*[]HyperflexClusterProfileRelationship, bool) {
	if o == nil || o.ClusterProfiles == nil {
		return nil, false
	}
	return o.ClusterProfiles, true
}

// HasClusterProfiles returns a boolean if a field has been set.
func (o *HyperflexVcenterConfigPolicy) HasClusterProfiles() bool {
	if o != nil && o.ClusterProfiles != nil {
		return true
	}

	return false
}

// SetClusterProfiles gets a reference to the given []HyperflexClusterProfileRelationship and assigns it to the ClusterProfiles field.
func (o *HyperflexVcenterConfigPolicy) SetClusterProfiles(v []HyperflexClusterProfileRelationship) {
	o.ClusterProfiles = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *HyperflexVcenterConfigPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVcenterConfigPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *HyperflexVcenterConfigPolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *HyperflexVcenterConfigPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o HyperflexVcenterConfigPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	if o.DataCenter != nil {
		toSerialize["DataCenter"] = o.DataCenter
	}
	if o.Hostname != nil {
		toSerialize["Hostname"] = o.Hostname
	}
	if o.IsPasswordSet != nil {
		toSerialize["IsPasswordSet"] = o.IsPasswordSet
	}
	if o.Password != nil {
		toSerialize["Password"] = o.Password
	}
	if o.SsoUrl != nil {
		toSerialize["SsoUrl"] = o.SsoUrl
	}
	if o.Username != nil {
		toSerialize["Username"] = o.Username
	}
	if o.ClusterProfiles != nil {
		toSerialize["ClusterProfiles"] = o.ClusterProfiles
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	return json.Marshal(toSerialize)
}

// AsHyperflexVcenterConfigPolicyRelationship wraps this instance of HyperflexVcenterConfigPolicy in HyperflexVcenterConfigPolicyRelationship
func (s *HyperflexVcenterConfigPolicy) AsHyperflexVcenterConfigPolicyRelationship() HyperflexVcenterConfigPolicyRelationship {
	return HyperflexVcenterConfigPolicyRelationship{HyperflexVcenterConfigPolicyRelationshipInterface: s}
}

type NullableHyperflexVcenterConfigPolicy struct {
	value *HyperflexVcenterConfigPolicy
	isSet bool
}

func (v NullableHyperflexVcenterConfigPolicy) Get() *HyperflexVcenterConfigPolicy {
	return v.value
}

func (v *NullableHyperflexVcenterConfigPolicy) Set(val *HyperflexVcenterConfigPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexVcenterConfigPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexVcenterConfigPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexVcenterConfigPolicy(val *HyperflexVcenterConfigPolicy) *NullableHyperflexVcenterConfigPolicy {
	return &NullableHyperflexVcenterConfigPolicy{value: val, isSet: true}
}

func (v NullableHyperflexVcenterConfigPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexVcenterConfigPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
