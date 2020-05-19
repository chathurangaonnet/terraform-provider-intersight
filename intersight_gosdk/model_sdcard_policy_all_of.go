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

// SdcardPolicyAllOf Definition of the list of properties defined in 'sdcard.Policy', excluding properties defined in parent classes.
type SdcardPolicyAllOf struct {
	Partitions   *[]SdcardPartition                    `json:"Partitions,omitempty"`
	Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to policyAbstractConfigProfile resources.
	Profiles *[]PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
}

// NewSdcardPolicyAllOf instantiates a new SdcardPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSdcardPolicyAllOf() *SdcardPolicyAllOf {
	this := SdcardPolicyAllOf{}
	return &this
}

// NewSdcardPolicyAllOfWithDefaults instantiates a new SdcardPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSdcardPolicyAllOfWithDefaults() *SdcardPolicyAllOf {
	this := SdcardPolicyAllOf{}
	return &this
}

// GetPartitions returns the Partitions field value if set, zero value otherwise.
func (o *SdcardPolicyAllOf) GetPartitions() []SdcardPartition {
	if o == nil || o.Partitions == nil {
		var ret []SdcardPartition
		return ret
	}
	return *o.Partitions
}

// GetPartitionsOk returns a tuple with the Partitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdcardPolicyAllOf) GetPartitionsOk() (*[]SdcardPartition, bool) {
	if o == nil || o.Partitions == nil {
		return nil, false
	}
	return o.Partitions, true
}

// HasPartitions returns a boolean if a field has been set.
func (o *SdcardPolicyAllOf) HasPartitions() bool {
	if o != nil && o.Partitions != nil {
		return true
	}

	return false
}

// SetPartitions gets a reference to the given []SdcardPartition and assigns it to the Partitions field.
func (o *SdcardPolicyAllOf) SetPartitions(v []SdcardPartition) {
	o.Partitions = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *SdcardPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdcardPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *SdcardPolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *SdcardPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise.
func (o *SdcardPolicyAllOf) GetProfiles() []PolicyAbstractConfigProfileRelationship {
	if o == nil || o.Profiles == nil {
		var ret []PolicyAbstractConfigProfileRelationship
		return ret
	}
	return *o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdcardPolicyAllOf) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *SdcardPolicyAllOf) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []PolicyAbstractConfigProfileRelationship and assigns it to the Profiles field.
func (o *SdcardPolicyAllOf) SetProfiles(v []PolicyAbstractConfigProfileRelationship) {
	o.Profiles = &v
}

func (o SdcardPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Partitions != nil {
		toSerialize["Partitions"] = o.Partitions
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Profiles != nil {
		toSerialize["Profiles"] = o.Profiles
	}
	return json.Marshal(toSerialize)
}

type NullableSdcardPolicyAllOf struct {
	value *SdcardPolicyAllOf
	isSet bool
}

func (v NullableSdcardPolicyAllOf) Get() *SdcardPolicyAllOf {
	return v.value
}

func (v *NullableSdcardPolicyAllOf) Set(val *SdcardPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSdcardPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSdcardPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSdcardPolicyAllOf(val *SdcardPolicyAllOf) *NullableSdcardPolicyAllOf {
	return &NullableSdcardPolicyAllOf{value: val, isSet: true}
}

func (v NullableSdcardPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSdcardPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
