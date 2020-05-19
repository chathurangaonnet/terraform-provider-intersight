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

// StorageDiskGroupPolicy A reusable RAID disk group configuration that can be applied across multiple servers. Also provides options to move JBOD disks in the disk group to Unconfigured Good state before they are used in the disk group.
type StorageDiskGroupPolicy struct {
	PolicyAbstractPolicy
	DedicatedHotSpares *[]StorageLocalDisk `json:"DedicatedHotSpares,omitempty"`
	// The supported RAID level for the disk group.
	RaidLevel  *string             `json:"RaidLevel,omitempty"`
	SpanGroups *[]StorageSpanGroup `json:"SpanGroups,omitempty"`
	// Selected disks in the disk group in JBOD state will be set to Unconfigured Good state before they are used in virtual drive creation.
	UseJbods     *bool                                 `json:"UseJbods,omitempty"`
	Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to storageStoragePolicy resources.
	StoragePolicies *[]StorageStoragePolicyRelationship `json:"StoragePolicies,omitempty"`
}

// NewStorageDiskGroupPolicy instantiates a new StorageDiskGroupPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageDiskGroupPolicy() *StorageDiskGroupPolicy {
	this := StorageDiskGroupPolicy{}
	var raidLevel string = "Raid0"
	this.RaidLevel = &raidLevel
	return &this
}

// NewStorageDiskGroupPolicyWithDefaults instantiates a new StorageDiskGroupPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageDiskGroupPolicyWithDefaults() *StorageDiskGroupPolicy {
	this := StorageDiskGroupPolicy{}
	var raidLevel string = "Raid0"
	this.RaidLevel = &raidLevel
	return &this
}

// GetDedicatedHotSpares returns the DedicatedHotSpares field value if set, zero value otherwise.
func (o *StorageDiskGroupPolicy) GetDedicatedHotSpares() []StorageLocalDisk {
	if o == nil || o.DedicatedHotSpares == nil {
		var ret []StorageLocalDisk
		return ret
	}
	return *o.DedicatedHotSpares
}

// GetDedicatedHotSparesOk returns a tuple with the DedicatedHotSpares field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageDiskGroupPolicy) GetDedicatedHotSparesOk() (*[]StorageLocalDisk, bool) {
	if o == nil || o.DedicatedHotSpares == nil {
		return nil, false
	}
	return o.DedicatedHotSpares, true
}

// HasDedicatedHotSpares returns a boolean if a field has been set.
func (o *StorageDiskGroupPolicy) HasDedicatedHotSpares() bool {
	if o != nil && o.DedicatedHotSpares != nil {
		return true
	}

	return false
}

// SetDedicatedHotSpares gets a reference to the given []StorageLocalDisk and assigns it to the DedicatedHotSpares field.
func (o *StorageDiskGroupPolicy) SetDedicatedHotSpares(v []StorageLocalDisk) {
	o.DedicatedHotSpares = &v
}

// GetRaidLevel returns the RaidLevel field value if set, zero value otherwise.
func (o *StorageDiskGroupPolicy) GetRaidLevel() string {
	if o == nil || o.RaidLevel == nil {
		var ret string
		return ret
	}
	return *o.RaidLevel
}

// GetRaidLevelOk returns a tuple with the RaidLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageDiskGroupPolicy) GetRaidLevelOk() (*string, bool) {
	if o == nil || o.RaidLevel == nil {
		return nil, false
	}
	return o.RaidLevel, true
}

// HasRaidLevel returns a boolean if a field has been set.
func (o *StorageDiskGroupPolicy) HasRaidLevel() bool {
	if o != nil && o.RaidLevel != nil {
		return true
	}

	return false
}

// SetRaidLevel gets a reference to the given string and assigns it to the RaidLevel field.
func (o *StorageDiskGroupPolicy) SetRaidLevel(v string) {
	o.RaidLevel = &v
}

// GetSpanGroups returns the SpanGroups field value if set, zero value otherwise.
func (o *StorageDiskGroupPolicy) GetSpanGroups() []StorageSpanGroup {
	if o == nil || o.SpanGroups == nil {
		var ret []StorageSpanGroup
		return ret
	}
	return *o.SpanGroups
}

// GetSpanGroupsOk returns a tuple with the SpanGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageDiskGroupPolicy) GetSpanGroupsOk() (*[]StorageSpanGroup, bool) {
	if o == nil || o.SpanGroups == nil {
		return nil, false
	}
	return o.SpanGroups, true
}

// HasSpanGroups returns a boolean if a field has been set.
func (o *StorageDiskGroupPolicy) HasSpanGroups() bool {
	if o != nil && o.SpanGroups != nil {
		return true
	}

	return false
}

// SetSpanGroups gets a reference to the given []StorageSpanGroup and assigns it to the SpanGroups field.
func (o *StorageDiskGroupPolicy) SetSpanGroups(v []StorageSpanGroup) {
	o.SpanGroups = &v
}

// GetUseJbods returns the UseJbods field value if set, zero value otherwise.
func (o *StorageDiskGroupPolicy) GetUseJbods() bool {
	if o == nil || o.UseJbods == nil {
		var ret bool
		return ret
	}
	return *o.UseJbods
}

// GetUseJbodsOk returns a tuple with the UseJbods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageDiskGroupPolicy) GetUseJbodsOk() (*bool, bool) {
	if o == nil || o.UseJbods == nil {
		return nil, false
	}
	return o.UseJbods, true
}

// HasUseJbods returns a boolean if a field has been set.
func (o *StorageDiskGroupPolicy) HasUseJbods() bool {
	if o != nil && o.UseJbods != nil {
		return true
	}

	return false
}

// SetUseJbods gets a reference to the given bool and assigns it to the UseJbods field.
func (o *StorageDiskGroupPolicy) SetUseJbods(v bool) {
	o.UseJbods = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *StorageDiskGroupPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageDiskGroupPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *StorageDiskGroupPolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *StorageDiskGroupPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetStoragePolicies returns the StoragePolicies field value if set, zero value otherwise.
func (o *StorageDiskGroupPolicy) GetStoragePolicies() []StorageStoragePolicyRelationship {
	if o == nil || o.StoragePolicies == nil {
		var ret []StorageStoragePolicyRelationship
		return ret
	}
	return *o.StoragePolicies
}

// GetStoragePoliciesOk returns a tuple with the StoragePolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageDiskGroupPolicy) GetStoragePoliciesOk() (*[]StorageStoragePolicyRelationship, bool) {
	if o == nil || o.StoragePolicies == nil {
		return nil, false
	}
	return o.StoragePolicies, true
}

// HasStoragePolicies returns a boolean if a field has been set.
func (o *StorageDiskGroupPolicy) HasStoragePolicies() bool {
	if o != nil && o.StoragePolicies != nil {
		return true
	}

	return false
}

// SetStoragePolicies gets a reference to the given []StorageStoragePolicyRelationship and assigns it to the StoragePolicies field.
func (o *StorageDiskGroupPolicy) SetStoragePolicies(v []StorageStoragePolicyRelationship) {
	o.StoragePolicies = &v
}

func (o StorageDiskGroupPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	if o.DedicatedHotSpares != nil {
		toSerialize["DedicatedHotSpares"] = o.DedicatedHotSpares
	}
	if o.RaidLevel != nil {
		toSerialize["RaidLevel"] = o.RaidLevel
	}
	if o.SpanGroups != nil {
		toSerialize["SpanGroups"] = o.SpanGroups
	}
	if o.UseJbods != nil {
		toSerialize["UseJbods"] = o.UseJbods
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.StoragePolicies != nil {
		toSerialize["StoragePolicies"] = o.StoragePolicies
	}
	return json.Marshal(toSerialize)
}

// AsStorageDiskGroupPolicyRelationship wraps this instance of StorageDiskGroupPolicy in StorageDiskGroupPolicyRelationship
func (s *StorageDiskGroupPolicy) AsStorageDiskGroupPolicyRelationship() StorageDiskGroupPolicyRelationship {
	return StorageDiskGroupPolicyRelationship{StorageDiskGroupPolicyRelationshipInterface: s}
}

type NullableStorageDiskGroupPolicy struct {
	value *StorageDiskGroupPolicy
	isSet bool
}

func (v NullableStorageDiskGroupPolicy) Get() *StorageDiskGroupPolicy {
	return v.value
}

func (v *NullableStorageDiskGroupPolicy) Set(val *StorageDiskGroupPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageDiskGroupPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageDiskGroupPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageDiskGroupPolicy(val *StorageDiskGroupPolicy) *NullableStorageDiskGroupPolicy {
	return &NullableStorageDiskGroupPolicy{value: val, isSet: true}
}

func (v NullableStorageDiskGroupPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageDiskGroupPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
