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

// HyperflexSysConfigPolicyAllOf Definition of the list of properties defined in 'hyperflex.SysConfigPolicy', excluding properties defined in parent classes.
type HyperflexSysConfigPolicyAllOf struct {
	// The DNS Search Domain Name. This setting applies to HyperFlex Data Platform 3.0 or later only.
	DnsDomainName *string   `json:"DnsDomainName,omitempty"`
	DnsServers    *[]string `json:"DnsServers,omitempty"`
	NtpServers    *[]string `json:"NtpServers,omitempty"`
	// The timezone of the HyperFlex cluster's system clock.
	Timezone *string `json:"Timezone,omitempty"`
	// An array of relationships to hyperflexClusterProfile resources.
	ClusterProfiles      []HyperflexClusterProfileRelationship `json:"ClusterProfiles,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexSysConfigPolicyAllOf HyperflexSysConfigPolicyAllOf

// NewHyperflexSysConfigPolicyAllOf instantiates a new HyperflexSysConfigPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexSysConfigPolicyAllOf() *HyperflexSysConfigPolicyAllOf {
	this := HyperflexSysConfigPolicyAllOf{}
	var timezone string = "Pacific/Niue"
	this.Timezone = &timezone
	return &this
}

// NewHyperflexSysConfigPolicyAllOfWithDefaults instantiates a new HyperflexSysConfigPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexSysConfigPolicyAllOfWithDefaults() *HyperflexSysConfigPolicyAllOf {
	this := HyperflexSysConfigPolicyAllOf{}
	var timezone string = "Pacific/Niue"
	this.Timezone = &timezone
	return &this
}

// GetDnsDomainName returns the DnsDomainName field value if set, zero value otherwise.
func (o *HyperflexSysConfigPolicyAllOf) GetDnsDomainName() string {
	if o == nil || o.DnsDomainName == nil {
		var ret string
		return ret
	}
	return *o.DnsDomainName
}

// GetDnsDomainNameOk returns a tuple with the DnsDomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSysConfigPolicyAllOf) GetDnsDomainNameOk() (*string, bool) {
	if o == nil || o.DnsDomainName == nil {
		return nil, false
	}
	return o.DnsDomainName, true
}

// HasDnsDomainName returns a boolean if a field has been set.
func (o *HyperflexSysConfigPolicyAllOf) HasDnsDomainName() bool {
	if o != nil && o.DnsDomainName != nil {
		return true
	}

	return false
}

// SetDnsDomainName gets a reference to the given string and assigns it to the DnsDomainName field.
func (o *HyperflexSysConfigPolicyAllOf) SetDnsDomainName(v string) {
	o.DnsDomainName = &v
}

// GetDnsServers returns the DnsServers field value if set, zero value otherwise.
func (o *HyperflexSysConfigPolicyAllOf) GetDnsServers() []string {
	if o == nil || o.DnsServers == nil {
		var ret []string
		return ret
	}
	return *o.DnsServers
}

// GetDnsServersOk returns a tuple with the DnsServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSysConfigPolicyAllOf) GetDnsServersOk() (*[]string, bool) {
	if o == nil || o.DnsServers == nil {
		return nil, false
	}
	return o.DnsServers, true
}

// HasDnsServers returns a boolean if a field has been set.
func (o *HyperflexSysConfigPolicyAllOf) HasDnsServers() bool {
	if o != nil && o.DnsServers != nil {
		return true
	}

	return false
}

// SetDnsServers gets a reference to the given []string and assigns it to the DnsServers field.
func (o *HyperflexSysConfigPolicyAllOf) SetDnsServers(v []string) {
	o.DnsServers = &v
}

// GetNtpServers returns the NtpServers field value if set, zero value otherwise.
func (o *HyperflexSysConfigPolicyAllOf) GetNtpServers() []string {
	if o == nil || o.NtpServers == nil {
		var ret []string
		return ret
	}
	return *o.NtpServers
}

// GetNtpServersOk returns a tuple with the NtpServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSysConfigPolicyAllOf) GetNtpServersOk() (*[]string, bool) {
	if o == nil || o.NtpServers == nil {
		return nil, false
	}
	return o.NtpServers, true
}

// HasNtpServers returns a boolean if a field has been set.
func (o *HyperflexSysConfigPolicyAllOf) HasNtpServers() bool {
	if o != nil && o.NtpServers != nil {
		return true
	}

	return false
}

// SetNtpServers gets a reference to the given []string and assigns it to the NtpServers field.
func (o *HyperflexSysConfigPolicyAllOf) SetNtpServers(v []string) {
	o.NtpServers = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *HyperflexSysConfigPolicyAllOf) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSysConfigPolicyAllOf) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *HyperflexSysConfigPolicyAllOf) HasTimezone() bool {
	if o != nil && o.Timezone != nil {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *HyperflexSysConfigPolicyAllOf) SetTimezone(v string) {
	o.Timezone = &v
}

// GetClusterProfiles returns the ClusterProfiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexSysConfigPolicyAllOf) GetClusterProfiles() []HyperflexClusterProfileRelationship {
	if o == nil {
		var ret []HyperflexClusterProfileRelationship
		return ret
	}
	return o.ClusterProfiles
}

// GetClusterProfilesOk returns a tuple with the ClusterProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexSysConfigPolicyAllOf) GetClusterProfilesOk() (*[]HyperflexClusterProfileRelationship, bool) {
	if o == nil || o.ClusterProfiles == nil {
		return nil, false
	}
	return &o.ClusterProfiles, true
}

// HasClusterProfiles returns a boolean if a field has been set.
func (o *HyperflexSysConfigPolicyAllOf) HasClusterProfiles() bool {
	if o != nil && o.ClusterProfiles != nil {
		return true
	}

	return false
}

// SetClusterProfiles gets a reference to the given []HyperflexClusterProfileRelationship and assigns it to the ClusterProfiles field.
func (o *HyperflexSysConfigPolicyAllOf) SetClusterProfiles(v []HyperflexClusterProfileRelationship) {
	o.ClusterProfiles = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *HyperflexSysConfigPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSysConfigPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *HyperflexSysConfigPolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *HyperflexSysConfigPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o HyperflexSysConfigPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DnsDomainName != nil {
		toSerialize["DnsDomainName"] = o.DnsDomainName
	}
	if o.DnsServers != nil {
		toSerialize["DnsServers"] = o.DnsServers
	}
	if o.NtpServers != nil {
		toSerialize["NtpServers"] = o.NtpServers
	}
	if o.Timezone != nil {
		toSerialize["Timezone"] = o.Timezone
	}
	if o.ClusterProfiles != nil {
		toSerialize["ClusterProfiles"] = o.ClusterProfiles
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexSysConfigPolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexSysConfigPolicyAllOf := _HyperflexSysConfigPolicyAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexSysConfigPolicyAllOf); err == nil {
		*o = HyperflexSysConfigPolicyAllOf(varHyperflexSysConfigPolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "DnsDomainName")
		delete(additionalProperties, "DnsServers")
		delete(additionalProperties, "NtpServers")
		delete(additionalProperties, "Timezone")
		delete(additionalProperties, "ClusterProfiles")
		delete(additionalProperties, "Organization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexSysConfigPolicyAllOf struct {
	value *HyperflexSysConfigPolicyAllOf
	isSet bool
}

func (v NullableHyperflexSysConfigPolicyAllOf) Get() *HyperflexSysConfigPolicyAllOf {
	return v.value
}

func (v *NullableHyperflexSysConfigPolicyAllOf) Set(val *HyperflexSysConfigPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexSysConfigPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexSysConfigPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexSysConfigPolicyAllOf(val *HyperflexSysConfigPolicyAllOf) *NullableHyperflexSysConfigPolicyAllOf {
	return &NullableHyperflexSysConfigPolicyAllOf{value: val, isSet: true}
}

func (v NullableHyperflexSysConfigPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexSysConfigPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
