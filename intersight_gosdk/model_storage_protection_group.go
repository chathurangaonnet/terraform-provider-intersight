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

// StorageProtectionGroup A protection group contains list of members that are protected together through snapshots with point-in-time consistency across the members. Members within the protection group have common data protection requirements and also the same snapshot, replication, and retention schedules.
type StorageProtectionGroup struct {
	MoBaseMo
	// Name of the protection Group.
	Name *string `json:"Name,omitempty"`
	// Prefix used for all generated snapshots from the protection group.
	Prefix *string `json:"Prefix,omitempty"`
	// Flag to determine if replication is enabled. Snapshots are replicated to the target array if this flag is set.
	ReplicationEnabled *bool `json:"ReplicationEnabled,omitempty"`
	// Flag to determine if snapshot creation is enabled. Snapshots are created on local array if this flag is set.
	SnapshotEnabled *bool                            `json:"SnapshotEnabled,omitempty"`
	StorageArray    *StorageGenericArrayRelationship `json:"StorageArray,omitempty"`
}

// NewStorageProtectionGroup instantiates a new StorageProtectionGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageProtectionGroup() *StorageProtectionGroup {
	this := StorageProtectionGroup{}
	return &this
}

// NewStorageProtectionGroupWithDefaults instantiates a new StorageProtectionGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageProtectionGroupWithDefaults() *StorageProtectionGroup {
	this := StorageProtectionGroup{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageProtectionGroup) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageProtectionGroup) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageProtectionGroup) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageProtectionGroup) SetName(v string) {
	o.Name = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *StorageProtectionGroup) GetPrefix() string {
	if o == nil || o.Prefix == nil {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageProtectionGroup) GetPrefixOk() (*string, bool) {
	if o == nil || o.Prefix == nil {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *StorageProtectionGroup) HasPrefix() bool {
	if o != nil && o.Prefix != nil {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *StorageProtectionGroup) SetPrefix(v string) {
	o.Prefix = &v
}

// GetReplicationEnabled returns the ReplicationEnabled field value if set, zero value otherwise.
func (o *StorageProtectionGroup) GetReplicationEnabled() bool {
	if o == nil || o.ReplicationEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ReplicationEnabled
}

// GetReplicationEnabledOk returns a tuple with the ReplicationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageProtectionGroup) GetReplicationEnabledOk() (*bool, bool) {
	if o == nil || o.ReplicationEnabled == nil {
		return nil, false
	}
	return o.ReplicationEnabled, true
}

// HasReplicationEnabled returns a boolean if a field has been set.
func (o *StorageProtectionGroup) HasReplicationEnabled() bool {
	if o != nil && o.ReplicationEnabled != nil {
		return true
	}

	return false
}

// SetReplicationEnabled gets a reference to the given bool and assigns it to the ReplicationEnabled field.
func (o *StorageProtectionGroup) SetReplicationEnabled(v bool) {
	o.ReplicationEnabled = &v
}

// GetSnapshotEnabled returns the SnapshotEnabled field value if set, zero value otherwise.
func (o *StorageProtectionGroup) GetSnapshotEnabled() bool {
	if o == nil || o.SnapshotEnabled == nil {
		var ret bool
		return ret
	}
	return *o.SnapshotEnabled
}

// GetSnapshotEnabledOk returns a tuple with the SnapshotEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageProtectionGroup) GetSnapshotEnabledOk() (*bool, bool) {
	if o == nil || o.SnapshotEnabled == nil {
		return nil, false
	}
	return o.SnapshotEnabled, true
}

// HasSnapshotEnabled returns a boolean if a field has been set.
func (o *StorageProtectionGroup) HasSnapshotEnabled() bool {
	if o != nil && o.SnapshotEnabled != nil {
		return true
	}

	return false
}

// SetSnapshotEnabled gets a reference to the given bool and assigns it to the SnapshotEnabled field.
func (o *StorageProtectionGroup) SetSnapshotEnabled(v bool) {
	o.SnapshotEnabled = &v
}

// GetStorageArray returns the StorageArray field value if set, zero value otherwise.
func (o *StorageProtectionGroup) GetStorageArray() StorageGenericArrayRelationship {
	if o == nil || o.StorageArray == nil {
		var ret StorageGenericArrayRelationship
		return ret
	}
	return *o.StorageArray
}

// GetStorageArrayOk returns a tuple with the StorageArray field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageProtectionGroup) GetStorageArrayOk() (*StorageGenericArrayRelationship, bool) {
	if o == nil || o.StorageArray == nil {
		return nil, false
	}
	return o.StorageArray, true
}

// HasStorageArray returns a boolean if a field has been set.
func (o *StorageProtectionGroup) HasStorageArray() bool {
	if o != nil && o.StorageArray != nil {
		return true
	}

	return false
}

// SetStorageArray gets a reference to the given StorageGenericArrayRelationship and assigns it to the StorageArray field.
func (o *StorageProtectionGroup) SetStorageArray(v StorageGenericArrayRelationship) {
	o.StorageArray = &v
}

func (o StorageProtectionGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Prefix != nil {
		toSerialize["Prefix"] = o.Prefix
	}
	if o.ReplicationEnabled != nil {
		toSerialize["ReplicationEnabled"] = o.ReplicationEnabled
	}
	if o.SnapshotEnabled != nil {
		toSerialize["SnapshotEnabled"] = o.SnapshotEnabled
	}
	if o.StorageArray != nil {
		toSerialize["StorageArray"] = o.StorageArray
	}
	return json.Marshal(toSerialize)
}

// AsStorageProtectionGroupRelationship wraps this instance of StorageProtectionGroup in StorageProtectionGroupRelationship
func (s *StorageProtectionGroup) AsStorageProtectionGroupRelationship() StorageProtectionGroupRelationship {
	return StorageProtectionGroupRelationship{StorageProtectionGroupRelationshipInterface: s}
}

type NullableStorageProtectionGroup struct {
	value *StorageProtectionGroup
	isSet bool
}

func (v NullableStorageProtectionGroup) Get() *StorageProtectionGroup {
	return v.value
}

func (v *NullableStorageProtectionGroup) Set(val *StorageProtectionGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageProtectionGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageProtectionGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageProtectionGroup(val *StorageProtectionGroup) *NullableStorageProtectionGroup {
	return &NullableStorageProtectionGroup{value: val, isSet: true}
}

func (v NullableStorageProtectionGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageProtectionGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
