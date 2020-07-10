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

// HyperflexSoftwareVersionPolicy A policy capturing software versions for different HyperFlex Cluster compatible components ( like HyperFlex Data Platform, Hypervisor, etc. ), that the user wishes to apply on the HyperFlex cluster.
type HyperflexSoftwareVersionPolicy struct {
	PolicyAbstractPolicy
	// Desired HyperFlex Data Platform software version to apply on the HyperFlex cluster.
	HxdpVersion *string `json:"HxdpVersion,omitempty"`
	// Desired  hypervisor version to apply for all the nodes on the HyperFlex cluster.
	HypervisorVersion *string `json:"HypervisorVersion,omitempty"`
	// Desired server firmware version to apply on the HyperFlex Cluster.
	ServerFirmwareVersion *string   `json:"ServerFirmwareVersion,omitempty"`
	UpgradeTypes          *[]string `json:"UpgradeTypes,omitempty"`
	// An array of relationships to hyperflexClusterProfile resources.
	ClusterProfiles           []HyperflexClusterProfileRelationship       `json:"ClusterProfiles,omitempty"`
	HxdpVersionInfo           *SoftwareHyperflexDistributableRelationship `json:"HxdpVersionInfo,omitempty"`
	HypervisorVersionInfo     *SoftwareHyperflexDistributableRelationship `json:"HypervisorVersionInfo,omitempty"`
	Organization              *OrganizationOrganizationRelationship       `json:"Organization,omitempty"`
	ServerFirmwareVersionInfo *FirmwareDistributableRelationship          `json:"ServerFirmwareVersionInfo,omitempty"`
}

// NewHyperflexSoftwareVersionPolicy instantiates a new HyperflexSoftwareVersionPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexSoftwareVersionPolicy() *HyperflexSoftwareVersionPolicy {
	this := HyperflexSoftwareVersionPolicy{}
	return &this
}

// NewHyperflexSoftwareVersionPolicyWithDefaults instantiates a new HyperflexSoftwareVersionPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexSoftwareVersionPolicyWithDefaults() *HyperflexSoftwareVersionPolicy {
	this := HyperflexSoftwareVersionPolicy{}
	return &this
}

// GetHxdpVersion returns the HxdpVersion field value if set, zero value otherwise.
func (o *HyperflexSoftwareVersionPolicy) GetHxdpVersion() string {
	if o == nil || o.HxdpVersion == nil {
		var ret string
		return ret
	}
	return *o.HxdpVersion
}

// GetHxdpVersionOk returns a tuple with the HxdpVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSoftwareVersionPolicy) GetHxdpVersionOk() (*string, bool) {
	if o == nil || o.HxdpVersion == nil {
		return nil, false
	}
	return o.HxdpVersion, true
}

// HasHxdpVersion returns a boolean if a field has been set.
func (o *HyperflexSoftwareVersionPolicy) HasHxdpVersion() bool {
	if o != nil && o.HxdpVersion != nil {
		return true
	}

	return false
}

// SetHxdpVersion gets a reference to the given string and assigns it to the HxdpVersion field.
func (o *HyperflexSoftwareVersionPolicy) SetHxdpVersion(v string) {
	o.HxdpVersion = &v
}

// GetHypervisorVersion returns the HypervisorVersion field value if set, zero value otherwise.
func (o *HyperflexSoftwareVersionPolicy) GetHypervisorVersion() string {
	if o == nil || o.HypervisorVersion == nil {
		var ret string
		return ret
	}
	return *o.HypervisorVersion
}

// GetHypervisorVersionOk returns a tuple with the HypervisorVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSoftwareVersionPolicy) GetHypervisorVersionOk() (*string, bool) {
	if o == nil || o.HypervisorVersion == nil {
		return nil, false
	}
	return o.HypervisorVersion, true
}

// HasHypervisorVersion returns a boolean if a field has been set.
func (o *HyperflexSoftwareVersionPolicy) HasHypervisorVersion() bool {
	if o != nil && o.HypervisorVersion != nil {
		return true
	}

	return false
}

// SetHypervisorVersion gets a reference to the given string and assigns it to the HypervisorVersion field.
func (o *HyperflexSoftwareVersionPolicy) SetHypervisorVersion(v string) {
	o.HypervisorVersion = &v
}

// GetServerFirmwareVersion returns the ServerFirmwareVersion field value if set, zero value otherwise.
func (o *HyperflexSoftwareVersionPolicy) GetServerFirmwareVersion() string {
	if o == nil || o.ServerFirmwareVersion == nil {
		var ret string
		return ret
	}
	return *o.ServerFirmwareVersion
}

// GetServerFirmwareVersionOk returns a tuple with the ServerFirmwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSoftwareVersionPolicy) GetServerFirmwareVersionOk() (*string, bool) {
	if o == nil || o.ServerFirmwareVersion == nil {
		return nil, false
	}
	return o.ServerFirmwareVersion, true
}

// HasServerFirmwareVersion returns a boolean if a field has been set.
func (o *HyperflexSoftwareVersionPolicy) HasServerFirmwareVersion() bool {
	if o != nil && o.ServerFirmwareVersion != nil {
		return true
	}

	return false
}

// SetServerFirmwareVersion gets a reference to the given string and assigns it to the ServerFirmwareVersion field.
func (o *HyperflexSoftwareVersionPolicy) SetServerFirmwareVersion(v string) {
	o.ServerFirmwareVersion = &v
}

// GetUpgradeTypes returns the UpgradeTypes field value if set, zero value otherwise.
func (o *HyperflexSoftwareVersionPolicy) GetUpgradeTypes() []string {
	if o == nil || o.UpgradeTypes == nil {
		var ret []string
		return ret
	}
	return *o.UpgradeTypes
}

// GetUpgradeTypesOk returns a tuple with the UpgradeTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSoftwareVersionPolicy) GetUpgradeTypesOk() (*[]string, bool) {
	if o == nil || o.UpgradeTypes == nil {
		return nil, false
	}
	return o.UpgradeTypes, true
}

// HasUpgradeTypes returns a boolean if a field has been set.
func (o *HyperflexSoftwareVersionPolicy) HasUpgradeTypes() bool {
	if o != nil && o.UpgradeTypes != nil {
		return true
	}

	return false
}

// SetUpgradeTypes gets a reference to the given []string and assigns it to the UpgradeTypes field.
func (o *HyperflexSoftwareVersionPolicy) SetUpgradeTypes(v []string) {
	o.UpgradeTypes = &v
}

// GetClusterProfiles returns the ClusterProfiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexSoftwareVersionPolicy) GetClusterProfiles() []HyperflexClusterProfileRelationship {
	if o == nil {
		var ret []HyperflexClusterProfileRelationship
		return ret
	}
	return o.ClusterProfiles
}

// GetClusterProfilesOk returns a tuple with the ClusterProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexSoftwareVersionPolicy) GetClusterProfilesOk() (*[]HyperflexClusterProfileRelationship, bool) {
	if o == nil || o.ClusterProfiles == nil {
		return nil, false
	}
	return &o.ClusterProfiles, true
}

// HasClusterProfiles returns a boolean if a field has been set.
func (o *HyperflexSoftwareVersionPolicy) HasClusterProfiles() bool {
	if o != nil && o.ClusterProfiles != nil {
		return true
	}

	return false
}

// SetClusterProfiles gets a reference to the given []HyperflexClusterProfileRelationship and assigns it to the ClusterProfiles field.
func (o *HyperflexSoftwareVersionPolicy) SetClusterProfiles(v []HyperflexClusterProfileRelationship) {
	o.ClusterProfiles = v
}

// GetHxdpVersionInfo returns the HxdpVersionInfo field value if set, zero value otherwise.
func (o *HyperflexSoftwareVersionPolicy) GetHxdpVersionInfo() SoftwareHyperflexDistributableRelationship {
	if o == nil || o.HxdpVersionInfo == nil {
		var ret SoftwareHyperflexDistributableRelationship
		return ret
	}
	return *o.HxdpVersionInfo
}

// GetHxdpVersionInfoOk returns a tuple with the HxdpVersionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSoftwareVersionPolicy) GetHxdpVersionInfoOk() (*SoftwareHyperflexDistributableRelationship, bool) {
	if o == nil || o.HxdpVersionInfo == nil {
		return nil, false
	}
	return o.HxdpVersionInfo, true
}

// HasHxdpVersionInfo returns a boolean if a field has been set.
func (o *HyperflexSoftwareVersionPolicy) HasHxdpVersionInfo() bool {
	if o != nil && o.HxdpVersionInfo != nil {
		return true
	}

	return false
}

// SetHxdpVersionInfo gets a reference to the given SoftwareHyperflexDistributableRelationship and assigns it to the HxdpVersionInfo field.
func (o *HyperflexSoftwareVersionPolicy) SetHxdpVersionInfo(v SoftwareHyperflexDistributableRelationship) {
	o.HxdpVersionInfo = &v
}

// GetHypervisorVersionInfo returns the HypervisorVersionInfo field value if set, zero value otherwise.
func (o *HyperflexSoftwareVersionPolicy) GetHypervisorVersionInfo() SoftwareHyperflexDistributableRelationship {
	if o == nil || o.HypervisorVersionInfo == nil {
		var ret SoftwareHyperflexDistributableRelationship
		return ret
	}
	return *o.HypervisorVersionInfo
}

// GetHypervisorVersionInfoOk returns a tuple with the HypervisorVersionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSoftwareVersionPolicy) GetHypervisorVersionInfoOk() (*SoftwareHyperflexDistributableRelationship, bool) {
	if o == nil || o.HypervisorVersionInfo == nil {
		return nil, false
	}
	return o.HypervisorVersionInfo, true
}

// HasHypervisorVersionInfo returns a boolean if a field has been set.
func (o *HyperflexSoftwareVersionPolicy) HasHypervisorVersionInfo() bool {
	if o != nil && o.HypervisorVersionInfo != nil {
		return true
	}

	return false
}

// SetHypervisorVersionInfo gets a reference to the given SoftwareHyperflexDistributableRelationship and assigns it to the HypervisorVersionInfo field.
func (o *HyperflexSoftwareVersionPolicy) SetHypervisorVersionInfo(v SoftwareHyperflexDistributableRelationship) {
	o.HypervisorVersionInfo = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *HyperflexSoftwareVersionPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSoftwareVersionPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *HyperflexSoftwareVersionPolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *HyperflexSoftwareVersionPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetServerFirmwareVersionInfo returns the ServerFirmwareVersionInfo field value if set, zero value otherwise.
func (o *HyperflexSoftwareVersionPolicy) GetServerFirmwareVersionInfo() FirmwareDistributableRelationship {
	if o == nil || o.ServerFirmwareVersionInfo == nil {
		var ret FirmwareDistributableRelationship
		return ret
	}
	return *o.ServerFirmwareVersionInfo
}

// GetServerFirmwareVersionInfoOk returns a tuple with the ServerFirmwareVersionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSoftwareVersionPolicy) GetServerFirmwareVersionInfoOk() (*FirmwareDistributableRelationship, bool) {
	if o == nil || o.ServerFirmwareVersionInfo == nil {
		return nil, false
	}
	return o.ServerFirmwareVersionInfo, true
}

// HasServerFirmwareVersionInfo returns a boolean if a field has been set.
func (o *HyperflexSoftwareVersionPolicy) HasServerFirmwareVersionInfo() bool {
	if o != nil && o.ServerFirmwareVersionInfo != nil {
		return true
	}

	return false
}

// SetServerFirmwareVersionInfo gets a reference to the given FirmwareDistributableRelationship and assigns it to the ServerFirmwareVersionInfo field.
func (o *HyperflexSoftwareVersionPolicy) SetServerFirmwareVersionInfo(v FirmwareDistributableRelationship) {
	o.ServerFirmwareVersionInfo = &v
}

func (o HyperflexSoftwareVersionPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	if o.HxdpVersion != nil {
		toSerialize["HxdpVersion"] = o.HxdpVersion
	}
	if o.HypervisorVersion != nil {
		toSerialize["HypervisorVersion"] = o.HypervisorVersion
	}
	if o.ServerFirmwareVersion != nil {
		toSerialize["ServerFirmwareVersion"] = o.ServerFirmwareVersion
	}
	if o.UpgradeTypes != nil {
		toSerialize["UpgradeTypes"] = o.UpgradeTypes
	}
	if o.ClusterProfiles != nil {
		toSerialize["ClusterProfiles"] = o.ClusterProfiles
	}
	if o.HxdpVersionInfo != nil {
		toSerialize["HxdpVersionInfo"] = o.HxdpVersionInfo
	}
	if o.HypervisorVersionInfo != nil {
		toSerialize["HypervisorVersionInfo"] = o.HypervisorVersionInfo
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.ServerFirmwareVersionInfo != nil {
		toSerialize["ServerFirmwareVersionInfo"] = o.ServerFirmwareVersionInfo
	}
	return json.Marshal(toSerialize)
}

type NullableHyperflexSoftwareVersionPolicy struct {
	value *HyperflexSoftwareVersionPolicy
	isSet bool
}

func (v NullableHyperflexSoftwareVersionPolicy) Get() *HyperflexSoftwareVersionPolicy {
	return v.value
}

func (v *NullableHyperflexSoftwareVersionPolicy) Set(val *HyperflexSoftwareVersionPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexSoftwareVersionPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexSoftwareVersionPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexSoftwareVersionPolicy(val *HyperflexSoftwareVersionPolicy) *NullableHyperflexSoftwareVersionPolicy {
	return &NullableHyperflexSoftwareVersionPolicy{value: val, isSet: true}
}

func (v NullableHyperflexSoftwareVersionPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexSoftwareVersionPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
