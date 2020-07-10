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

// FabricSwitchClusterProfile This specifies the configuration policies for a cluster of switches.
type FabricSwitchClusterProfile struct {
	PolicyAbstractProfile
	ConfigContext *PolicyConfigContext                  `json:"ConfigContext,omitempty"`
	Organization  *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to fabricSwitchProfile resources.
	SwitchProfiles []FabricSwitchProfileRelationship `json:"SwitchProfiles,omitempty"`
}

// NewFabricSwitchClusterProfile instantiates a new FabricSwitchClusterProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricSwitchClusterProfile() *FabricSwitchClusterProfile {
	this := FabricSwitchClusterProfile{}
	return &this
}

// NewFabricSwitchClusterProfileWithDefaults instantiates a new FabricSwitchClusterProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricSwitchClusterProfileWithDefaults() *FabricSwitchClusterProfile {
	this := FabricSwitchClusterProfile{}
	return &this
}

// GetConfigContext returns the ConfigContext field value if set, zero value otherwise.
func (o *FabricSwitchClusterProfile) GetConfigContext() PolicyConfigContext {
	if o == nil || o.ConfigContext == nil {
		var ret PolicyConfigContext
		return ret
	}
	return *o.ConfigContext
}

// GetConfigContextOk returns a tuple with the ConfigContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricSwitchClusterProfile) GetConfigContextOk() (*PolicyConfigContext, bool) {
	if o == nil || o.ConfigContext == nil {
		return nil, false
	}
	return o.ConfigContext, true
}

// HasConfigContext returns a boolean if a field has been set.
func (o *FabricSwitchClusterProfile) HasConfigContext() bool {
	if o != nil && o.ConfigContext != nil {
		return true
	}

	return false
}

// SetConfigContext gets a reference to the given PolicyConfigContext and assigns it to the ConfigContext field.
func (o *FabricSwitchClusterProfile) SetConfigContext(v PolicyConfigContext) {
	o.ConfigContext = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *FabricSwitchClusterProfile) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricSwitchClusterProfile) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *FabricSwitchClusterProfile) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *FabricSwitchClusterProfile) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetSwitchProfiles returns the SwitchProfiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricSwitchClusterProfile) GetSwitchProfiles() []FabricSwitchProfileRelationship {
	if o == nil {
		var ret []FabricSwitchProfileRelationship
		return ret
	}
	return o.SwitchProfiles
}

// GetSwitchProfilesOk returns a tuple with the SwitchProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricSwitchClusterProfile) GetSwitchProfilesOk() (*[]FabricSwitchProfileRelationship, bool) {
	if o == nil || o.SwitchProfiles == nil {
		return nil, false
	}
	return &o.SwitchProfiles, true
}

// HasSwitchProfiles returns a boolean if a field has been set.
func (o *FabricSwitchClusterProfile) HasSwitchProfiles() bool {
	if o != nil && o.SwitchProfiles != nil {
		return true
	}

	return false
}

// SetSwitchProfiles gets a reference to the given []FabricSwitchProfileRelationship and assigns it to the SwitchProfiles field.
func (o *FabricSwitchClusterProfile) SetSwitchProfiles(v []FabricSwitchProfileRelationship) {
	o.SwitchProfiles = v
}

func (o FabricSwitchClusterProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractProfile, errPolicyAbstractProfile := json.Marshal(o.PolicyAbstractProfile)
	if errPolicyAbstractProfile != nil {
		return []byte{}, errPolicyAbstractProfile
	}
	errPolicyAbstractProfile = json.Unmarshal([]byte(serializedPolicyAbstractProfile), &toSerialize)
	if errPolicyAbstractProfile != nil {
		return []byte{}, errPolicyAbstractProfile
	}
	if o.ConfigContext != nil {
		toSerialize["ConfigContext"] = o.ConfigContext
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.SwitchProfiles != nil {
		toSerialize["SwitchProfiles"] = o.SwitchProfiles
	}
	return json.Marshal(toSerialize)
}

type NullableFabricSwitchClusterProfile struct {
	value *FabricSwitchClusterProfile
	isSet bool
}

func (v NullableFabricSwitchClusterProfile) Get() *FabricSwitchClusterProfile {
	return v.value
}

func (v *NullableFabricSwitchClusterProfile) Set(val *FabricSwitchClusterProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricSwitchClusterProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricSwitchClusterProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricSwitchClusterProfile(val *FabricSwitchClusterProfile) *NullableFabricSwitchClusterProfile {
	return &NullableFabricSwitchClusterProfile{value: val, isSet: true}
}

func (v NullableFabricSwitchClusterProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricSwitchClusterProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
