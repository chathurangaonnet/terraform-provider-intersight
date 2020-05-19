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

// HyperflexNode struct for HyperflexNode
type HyperflexNode struct {
	MoBaseMo
	BuildNumber    *string                         `json:"BuildNumber,omitempty"`
	DisplayVersion *string                         `json:"DisplayVersion,omitempty"`
	HostName       *string                         `json:"HostName,omitempty"`
	Hypervisor     *string                         `json:"Hypervisor,omitempty"`
	Identity       *HyperflexHxUuIdDt              `json:"Identity,omitempty"`
	Ip             *HyperflexHxNetworkAddressDt    `json:"Ip,omitempty"`
	Lockdown       *bool                           `json:"Lockdown,omitempty"`
	ModelNumber    *string                         `json:"ModelNumber,omitempty"`
	Role           *string                         `json:"Role,omitempty"`
	SerialNumber   *string                         `json:"SerialNumber,omitempty"`
	Status         *string                         `json:"Status,omitempty"`
	Version        *string                         `json:"Version,omitempty"`
	Cluster        *HyperflexClusterRelationship   `json:"Cluster,omitempty"`
	ClusterMember  *AssetClusterMemberRelationship `json:"ClusterMember,omitempty"`
	PhysicalServer *ComputePhysicalRelationship    `json:"PhysicalServer,omitempty"`
}

// NewHyperflexNode instantiates a new HyperflexNode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexNode() *HyperflexNode {
	this := HyperflexNode{}
	var role string = "UNKNOWN"
	this.Role = &role
	var status string = "UNKNOWN"
	this.Status = &status
	return &this
}

// NewHyperflexNodeWithDefaults instantiates a new HyperflexNode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexNodeWithDefaults() *HyperflexNode {
	this := HyperflexNode{}
	var role string = "UNKNOWN"
	this.Role = &role
	var status string = "UNKNOWN"
	this.Status = &status
	return &this
}

// GetBuildNumber returns the BuildNumber field value if set, zero value otherwise.
func (o *HyperflexNode) GetBuildNumber() string {
	if o == nil || o.BuildNumber == nil {
		var ret string
		return ret
	}
	return *o.BuildNumber
}

// GetBuildNumberOk returns a tuple with the BuildNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNode) GetBuildNumberOk() (*string, bool) {
	if o == nil || o.BuildNumber == nil {
		return nil, false
	}
	return o.BuildNumber, true
}

// HasBuildNumber returns a boolean if a field has been set.
func (o *HyperflexNode) HasBuildNumber() bool {
	if o != nil && o.BuildNumber != nil {
		return true
	}

	return false
}

// SetBuildNumber gets a reference to the given string and assigns it to the BuildNumber field.
func (o *HyperflexNode) SetBuildNumber(v string) {
	o.BuildNumber = &v
}

// GetDisplayVersion returns the DisplayVersion field value if set, zero value otherwise.
func (o *HyperflexNode) GetDisplayVersion() string {
	if o == nil || o.DisplayVersion == nil {
		var ret string
		return ret
	}
	return *o.DisplayVersion
}

// GetDisplayVersionOk returns a tuple with the DisplayVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNode) GetDisplayVersionOk() (*string, bool) {
	if o == nil || o.DisplayVersion == nil {
		return nil, false
	}
	return o.DisplayVersion, true
}

// HasDisplayVersion returns a boolean if a field has been set.
func (o *HyperflexNode) HasDisplayVersion() bool {
	if o != nil && o.DisplayVersion != nil {
		return true
	}

	return false
}

// SetDisplayVersion gets a reference to the given string and assigns it to the DisplayVersion field.
func (o *HyperflexNode) SetDisplayVersion(v string) {
	o.DisplayVersion = &v
}

// GetHostName returns the HostName field value if set, zero value otherwise.
func (o *HyperflexNode) GetHostName() string {
	if o == nil || o.HostName == nil {
		var ret string
		return ret
	}
	return *o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNode) GetHostNameOk() (*string, bool) {
	if o == nil || o.HostName == nil {
		return nil, false
	}
	return o.HostName, true
}

// HasHostName returns a boolean if a field has been set.
func (o *HyperflexNode) HasHostName() bool {
	if o != nil && o.HostName != nil {
		return true
	}

	return false
}

// SetHostName gets a reference to the given string and assigns it to the HostName field.
func (o *HyperflexNode) SetHostName(v string) {
	o.HostName = &v
}

// GetHypervisor returns the Hypervisor field value if set, zero value otherwise.
func (o *HyperflexNode) GetHypervisor() string {
	if o == nil || o.Hypervisor == nil {
		var ret string
		return ret
	}
	return *o.Hypervisor
}

// GetHypervisorOk returns a tuple with the Hypervisor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNode) GetHypervisorOk() (*string, bool) {
	if o == nil || o.Hypervisor == nil {
		return nil, false
	}
	return o.Hypervisor, true
}

// HasHypervisor returns a boolean if a field has been set.
func (o *HyperflexNode) HasHypervisor() bool {
	if o != nil && o.Hypervisor != nil {
		return true
	}

	return false
}

// SetHypervisor gets a reference to the given string and assigns it to the Hypervisor field.
func (o *HyperflexNode) SetHypervisor(v string) {
	o.Hypervisor = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *HyperflexNode) GetIdentity() HyperflexHxUuIdDt {
	if o == nil || o.Identity == nil {
		var ret HyperflexHxUuIdDt
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNode) GetIdentityOk() (*HyperflexHxUuIdDt, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *HyperflexNode) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given HyperflexHxUuIdDt and assigns it to the Identity field.
func (o *HyperflexNode) SetIdentity(v HyperflexHxUuIdDt) {
	o.Identity = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *HyperflexNode) GetIp() HyperflexHxNetworkAddressDt {
	if o == nil || o.Ip == nil {
		var ret HyperflexHxNetworkAddressDt
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNode) GetIpOk() (*HyperflexHxNetworkAddressDt, bool) {
	if o == nil || o.Ip == nil {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *HyperflexNode) HasIp() bool {
	if o != nil && o.Ip != nil {
		return true
	}

	return false
}

// SetIp gets a reference to the given HyperflexHxNetworkAddressDt and assigns it to the Ip field.
func (o *HyperflexNode) SetIp(v HyperflexHxNetworkAddressDt) {
	o.Ip = &v
}

// GetLockdown returns the Lockdown field value if set, zero value otherwise.
func (o *HyperflexNode) GetLockdown() bool {
	if o == nil || o.Lockdown == nil {
		var ret bool
		return ret
	}
	return *o.Lockdown
}

// GetLockdownOk returns a tuple with the Lockdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNode) GetLockdownOk() (*bool, bool) {
	if o == nil || o.Lockdown == nil {
		return nil, false
	}
	return o.Lockdown, true
}

// HasLockdown returns a boolean if a field has been set.
func (o *HyperflexNode) HasLockdown() bool {
	if o != nil && o.Lockdown != nil {
		return true
	}

	return false
}

// SetLockdown gets a reference to the given bool and assigns it to the Lockdown field.
func (o *HyperflexNode) SetLockdown(v bool) {
	o.Lockdown = &v
}

// GetModelNumber returns the ModelNumber field value if set, zero value otherwise.
func (o *HyperflexNode) GetModelNumber() string {
	if o == nil || o.ModelNumber == nil {
		var ret string
		return ret
	}
	return *o.ModelNumber
}

// GetModelNumberOk returns a tuple with the ModelNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNode) GetModelNumberOk() (*string, bool) {
	if o == nil || o.ModelNumber == nil {
		return nil, false
	}
	return o.ModelNumber, true
}

// HasModelNumber returns a boolean if a field has been set.
func (o *HyperflexNode) HasModelNumber() bool {
	if o != nil && o.ModelNumber != nil {
		return true
	}

	return false
}

// SetModelNumber gets a reference to the given string and assigns it to the ModelNumber field.
func (o *HyperflexNode) SetModelNumber(v string) {
	o.ModelNumber = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *HyperflexNode) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNode) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *HyperflexNode) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *HyperflexNode) SetRole(v string) {
	o.Role = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *HyperflexNode) GetSerialNumber() string {
	if o == nil || o.SerialNumber == nil {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNode) GetSerialNumberOk() (*string, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *HyperflexNode) HasSerialNumber() bool {
	if o != nil && o.SerialNumber != nil {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *HyperflexNode) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HyperflexNode) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNode) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HyperflexNode) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *HyperflexNode) SetStatus(v string) {
	o.Status = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *HyperflexNode) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNode) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *HyperflexNode) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *HyperflexNode) SetVersion(v string) {
	o.Version = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *HyperflexNode) GetCluster() HyperflexClusterRelationship {
	if o == nil || o.Cluster == nil {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNode) GetClusterOk() (*HyperflexClusterRelationship, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *HyperflexNode) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given HyperflexClusterRelationship and assigns it to the Cluster field.
func (o *HyperflexNode) SetCluster(v HyperflexClusterRelationship) {
	o.Cluster = &v
}

// GetClusterMember returns the ClusterMember field value if set, zero value otherwise.
func (o *HyperflexNode) GetClusterMember() AssetClusterMemberRelationship {
	if o == nil || o.ClusterMember == nil {
		var ret AssetClusterMemberRelationship
		return ret
	}
	return *o.ClusterMember
}

// GetClusterMemberOk returns a tuple with the ClusterMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNode) GetClusterMemberOk() (*AssetClusterMemberRelationship, bool) {
	if o == nil || o.ClusterMember == nil {
		return nil, false
	}
	return o.ClusterMember, true
}

// HasClusterMember returns a boolean if a field has been set.
func (o *HyperflexNode) HasClusterMember() bool {
	if o != nil && o.ClusterMember != nil {
		return true
	}

	return false
}

// SetClusterMember gets a reference to the given AssetClusterMemberRelationship and assigns it to the ClusterMember field.
func (o *HyperflexNode) SetClusterMember(v AssetClusterMemberRelationship) {
	o.ClusterMember = &v
}

// GetPhysicalServer returns the PhysicalServer field value if set, zero value otherwise.
func (o *HyperflexNode) GetPhysicalServer() ComputePhysicalRelationship {
	if o == nil || o.PhysicalServer == nil {
		var ret ComputePhysicalRelationship
		return ret
	}
	return *o.PhysicalServer
}

// GetPhysicalServerOk returns a tuple with the PhysicalServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNode) GetPhysicalServerOk() (*ComputePhysicalRelationship, bool) {
	if o == nil || o.PhysicalServer == nil {
		return nil, false
	}
	return o.PhysicalServer, true
}

// HasPhysicalServer returns a boolean if a field has been set.
func (o *HyperflexNode) HasPhysicalServer() bool {
	if o != nil && o.PhysicalServer != nil {
		return true
	}

	return false
}

// SetPhysicalServer gets a reference to the given ComputePhysicalRelationship and assigns it to the PhysicalServer field.
func (o *HyperflexNode) SetPhysicalServer(v ComputePhysicalRelationship) {
	o.PhysicalServer = &v
}

func (o HyperflexNode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.BuildNumber != nil {
		toSerialize["BuildNumber"] = o.BuildNumber
	}
	if o.DisplayVersion != nil {
		toSerialize["DisplayVersion"] = o.DisplayVersion
	}
	if o.HostName != nil {
		toSerialize["HostName"] = o.HostName
	}
	if o.Hypervisor != nil {
		toSerialize["Hypervisor"] = o.Hypervisor
	}
	if o.Identity != nil {
		toSerialize["Identity"] = o.Identity
	}
	if o.Ip != nil {
		toSerialize["Ip"] = o.Ip
	}
	if o.Lockdown != nil {
		toSerialize["Lockdown"] = o.Lockdown
	}
	if o.ModelNumber != nil {
		toSerialize["ModelNumber"] = o.ModelNumber
	}
	if o.Role != nil {
		toSerialize["Role"] = o.Role
	}
	if o.SerialNumber != nil {
		toSerialize["SerialNumber"] = o.SerialNumber
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.Cluster != nil {
		toSerialize["Cluster"] = o.Cluster
	}
	if o.ClusterMember != nil {
		toSerialize["ClusterMember"] = o.ClusterMember
	}
	if o.PhysicalServer != nil {
		toSerialize["PhysicalServer"] = o.PhysicalServer
	}
	return json.Marshal(toSerialize)
}

// AsHyperflexNodeRelationship wraps this instance of HyperflexNode in HyperflexNodeRelationship
func (s *HyperflexNode) AsHyperflexNodeRelationship() HyperflexNodeRelationship {
	return HyperflexNodeRelationship{HyperflexNodeRelationshipInterface: s}
}

type NullableHyperflexNode struct {
	value *HyperflexNode
	isSet bool
}

func (v NullableHyperflexNode) Get() *HyperflexNode {
	return v.value
}

func (v *NullableHyperflexNode) Set(val *HyperflexNode) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexNode) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexNode(val *HyperflexNode) *NullableHyperflexNode {
	return &NullableHyperflexNode{value: val, isSet: true}
}

func (v NullableHyperflexNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
