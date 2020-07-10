/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-06-19T15:15:17Z.
 *
 * API version: 1.0.9-1903
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
)

// HyperflexAppSettingConstraint A constraint that can be used to qualify an application setting.
type HyperflexAppSettingConstraint struct {
	MoBaseComplexType
	// The supported HyperFlex Data Platform version in regex format.
	HxdpVersion *string `json:"HxdpVersion,omitempty"`
	// The hypervisor type for the HyperFlex cluster.
	HypervisorType *string `json:"HypervisorType,omitempty"`
	// The supported management platform for the HyperFlex Cluster.
	MgmtPlatform *string `json:"MgmtPlatform,omitempty"`
	// The supported server models in regex format.
	ServerModel *string `json:"ServerModel,omitempty"`
}

// NewHyperflexAppSettingConstraint instantiates a new HyperflexAppSettingConstraint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexAppSettingConstraint() *HyperflexAppSettingConstraint {
	this := HyperflexAppSettingConstraint{}
	var hypervisorType string = "ESXi"
	this.HypervisorType = &hypervisorType
	var mgmtPlatform string = "FI"
	this.MgmtPlatform = &mgmtPlatform
	return &this
}

// NewHyperflexAppSettingConstraintWithDefaults instantiates a new HyperflexAppSettingConstraint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexAppSettingConstraintWithDefaults() *HyperflexAppSettingConstraint {
	this := HyperflexAppSettingConstraint{}
	var hypervisorType string = "ESXi"
	this.HypervisorType = &hypervisorType
	var mgmtPlatform string = "FI"
	this.MgmtPlatform = &mgmtPlatform
	return &this
}

// GetHxdpVersion returns the HxdpVersion field value if set, zero value otherwise.
func (o *HyperflexAppSettingConstraint) GetHxdpVersion() string {
	if o == nil || o.HxdpVersion == nil {
		var ret string
		return ret
	}
	return *o.HxdpVersion
}

// GetHxdpVersionOk returns a tuple with the HxdpVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAppSettingConstraint) GetHxdpVersionOk() (*string, bool) {
	if o == nil || o.HxdpVersion == nil {
		return nil, false
	}
	return o.HxdpVersion, true
}

// HasHxdpVersion returns a boolean if a field has been set.
func (o *HyperflexAppSettingConstraint) HasHxdpVersion() bool {
	if o != nil && o.HxdpVersion != nil {
		return true
	}

	return false
}

// SetHxdpVersion gets a reference to the given string and assigns it to the HxdpVersion field.
func (o *HyperflexAppSettingConstraint) SetHxdpVersion(v string) {
	o.HxdpVersion = &v
}

// GetHypervisorType returns the HypervisorType field value if set, zero value otherwise.
func (o *HyperflexAppSettingConstraint) GetHypervisorType() string {
	if o == nil || o.HypervisorType == nil {
		var ret string
		return ret
	}
	return *o.HypervisorType
}

// GetHypervisorTypeOk returns a tuple with the HypervisorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAppSettingConstraint) GetHypervisorTypeOk() (*string, bool) {
	if o == nil || o.HypervisorType == nil {
		return nil, false
	}
	return o.HypervisorType, true
}

// HasHypervisorType returns a boolean if a field has been set.
func (o *HyperflexAppSettingConstraint) HasHypervisorType() bool {
	if o != nil && o.HypervisorType != nil {
		return true
	}

	return false
}

// SetHypervisorType gets a reference to the given string and assigns it to the HypervisorType field.
func (o *HyperflexAppSettingConstraint) SetHypervisorType(v string) {
	o.HypervisorType = &v
}

// GetMgmtPlatform returns the MgmtPlatform field value if set, zero value otherwise.
func (o *HyperflexAppSettingConstraint) GetMgmtPlatform() string {
	if o == nil || o.MgmtPlatform == nil {
		var ret string
		return ret
	}
	return *o.MgmtPlatform
}

// GetMgmtPlatformOk returns a tuple with the MgmtPlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAppSettingConstraint) GetMgmtPlatformOk() (*string, bool) {
	if o == nil || o.MgmtPlatform == nil {
		return nil, false
	}
	return o.MgmtPlatform, true
}

// HasMgmtPlatform returns a boolean if a field has been set.
func (o *HyperflexAppSettingConstraint) HasMgmtPlatform() bool {
	if o != nil && o.MgmtPlatform != nil {
		return true
	}

	return false
}

// SetMgmtPlatform gets a reference to the given string and assigns it to the MgmtPlatform field.
func (o *HyperflexAppSettingConstraint) SetMgmtPlatform(v string) {
	o.MgmtPlatform = &v
}

// GetServerModel returns the ServerModel field value if set, zero value otherwise.
func (o *HyperflexAppSettingConstraint) GetServerModel() string {
	if o == nil || o.ServerModel == nil {
		var ret string
		return ret
	}
	return *o.ServerModel
}

// GetServerModelOk returns a tuple with the ServerModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAppSettingConstraint) GetServerModelOk() (*string, bool) {
	if o == nil || o.ServerModel == nil {
		return nil, false
	}
	return o.ServerModel, true
}

// HasServerModel returns a boolean if a field has been set.
func (o *HyperflexAppSettingConstraint) HasServerModel() bool {
	if o != nil && o.ServerModel != nil {
		return true
	}

	return false
}

// SetServerModel gets a reference to the given string and assigns it to the ServerModel field.
func (o *HyperflexAppSettingConstraint) SetServerModel(v string) {
	o.ServerModel = &v
}

func (o HyperflexAppSettingConstraint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.HxdpVersion != nil {
		toSerialize["HxdpVersion"] = o.HxdpVersion
	}
	if o.HypervisorType != nil {
		toSerialize["HypervisorType"] = o.HypervisorType
	}
	if o.MgmtPlatform != nil {
		toSerialize["MgmtPlatform"] = o.MgmtPlatform
	}
	if o.ServerModel != nil {
		toSerialize["ServerModel"] = o.ServerModel
	}
	return json.Marshal(toSerialize)
}

type NullableHyperflexAppSettingConstraint struct {
	value *HyperflexAppSettingConstraint
	isSet bool
}

func (v NullableHyperflexAppSettingConstraint) Get() *HyperflexAppSettingConstraint {
	return v.value
}

func (v *NullableHyperflexAppSettingConstraint) Set(val *HyperflexAppSettingConstraint) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexAppSettingConstraint) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexAppSettingConstraint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexAppSettingConstraint(val *HyperflexAppSettingConstraint) *NullableHyperflexAppSettingConstraint {
	return &NullableHyperflexAppSettingConstraint{value: val, isSet: true}
}

func (v NullableHyperflexAppSettingConstraint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexAppSettingConstraint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}