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
	"time"
)

// StorageSnapshot Reference marker for volume. It is a point-in-time copy of the volume.
type StorageSnapshot struct {
	MoBaseMo
	// Exact date and time at which snapshot was created.
	CreatedTime *time.Time `json:"CreatedTime,omitempty"`
	// Name of the snapshot which represents point-in-time copy of volume.
	Name *string `json:"Name,omitempty"`
	// Name of the protection group to which the snapshot belongs. Value is empty, if the snapshot is created directly on volume.
	ProtectionGroupName *string `json:"ProtectionGroupName,omitempty"`
	// Snapshot size represented in bytes.
	Size *int64 `json:"Size,omitempty"`
	// Source object on which the snapshot is created. It is the name of the originating volume.
	Source                  *string                                     `json:"Source,omitempty"`
	ProtectionGroupSnapshot *StorageProtectionGroupSnapshotRelationship `json:"ProtectionGroupSnapshot,omitempty"`
	StorageArray            *StorageGenericArrayRelationship            `json:"StorageArray,omitempty"`
	Volume                  *StorageVolumeRelationship                  `json:"Volume,omitempty"`
}

// NewStorageSnapshot instantiates a new StorageSnapshot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageSnapshot() *StorageSnapshot {
	this := StorageSnapshot{}
	return &this
}

// NewStorageSnapshotWithDefaults instantiates a new StorageSnapshot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageSnapshotWithDefaults() *StorageSnapshot {
	this := StorageSnapshot{}
	return &this
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *StorageSnapshot) GetCreatedTime() time.Time {
	if o == nil || o.CreatedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSnapshot) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil || o.CreatedTime == nil {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *StorageSnapshot) HasCreatedTime() bool {
	if o != nil && o.CreatedTime != nil {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given time.Time and assigns it to the CreatedTime field.
func (o *StorageSnapshot) SetCreatedTime(v time.Time) {
	o.CreatedTime = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageSnapshot) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSnapshot) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageSnapshot) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageSnapshot) SetName(v string) {
	o.Name = &v
}

// GetProtectionGroupName returns the ProtectionGroupName field value if set, zero value otherwise.
func (o *StorageSnapshot) GetProtectionGroupName() string {
	if o == nil || o.ProtectionGroupName == nil {
		var ret string
		return ret
	}
	return *o.ProtectionGroupName
}

// GetProtectionGroupNameOk returns a tuple with the ProtectionGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSnapshot) GetProtectionGroupNameOk() (*string, bool) {
	if o == nil || o.ProtectionGroupName == nil {
		return nil, false
	}
	return o.ProtectionGroupName, true
}

// HasProtectionGroupName returns a boolean if a field has been set.
func (o *StorageSnapshot) HasProtectionGroupName() bool {
	if o != nil && o.ProtectionGroupName != nil {
		return true
	}

	return false
}

// SetProtectionGroupName gets a reference to the given string and assigns it to the ProtectionGroupName field.
func (o *StorageSnapshot) SetProtectionGroupName(v string) {
	o.ProtectionGroupName = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *StorageSnapshot) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSnapshot) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *StorageSnapshot) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *StorageSnapshot) SetSize(v int64) {
	o.Size = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *StorageSnapshot) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSnapshot) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *StorageSnapshot) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *StorageSnapshot) SetSource(v string) {
	o.Source = &v
}

// GetProtectionGroupSnapshot returns the ProtectionGroupSnapshot field value if set, zero value otherwise.
func (o *StorageSnapshot) GetProtectionGroupSnapshot() StorageProtectionGroupSnapshotRelationship {
	if o == nil || o.ProtectionGroupSnapshot == nil {
		var ret StorageProtectionGroupSnapshotRelationship
		return ret
	}
	return *o.ProtectionGroupSnapshot
}

// GetProtectionGroupSnapshotOk returns a tuple with the ProtectionGroupSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSnapshot) GetProtectionGroupSnapshotOk() (*StorageProtectionGroupSnapshotRelationship, bool) {
	if o == nil || o.ProtectionGroupSnapshot == nil {
		return nil, false
	}
	return o.ProtectionGroupSnapshot, true
}

// HasProtectionGroupSnapshot returns a boolean if a field has been set.
func (o *StorageSnapshot) HasProtectionGroupSnapshot() bool {
	if o != nil && o.ProtectionGroupSnapshot != nil {
		return true
	}

	return false
}

// SetProtectionGroupSnapshot gets a reference to the given StorageProtectionGroupSnapshotRelationship and assigns it to the ProtectionGroupSnapshot field.
func (o *StorageSnapshot) SetProtectionGroupSnapshot(v StorageProtectionGroupSnapshotRelationship) {
	o.ProtectionGroupSnapshot = &v
}

// GetStorageArray returns the StorageArray field value if set, zero value otherwise.
func (o *StorageSnapshot) GetStorageArray() StorageGenericArrayRelationship {
	if o == nil || o.StorageArray == nil {
		var ret StorageGenericArrayRelationship
		return ret
	}
	return *o.StorageArray
}

// GetStorageArrayOk returns a tuple with the StorageArray field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSnapshot) GetStorageArrayOk() (*StorageGenericArrayRelationship, bool) {
	if o == nil || o.StorageArray == nil {
		return nil, false
	}
	return o.StorageArray, true
}

// HasStorageArray returns a boolean if a field has been set.
func (o *StorageSnapshot) HasStorageArray() bool {
	if o != nil && o.StorageArray != nil {
		return true
	}

	return false
}

// SetStorageArray gets a reference to the given StorageGenericArrayRelationship and assigns it to the StorageArray field.
func (o *StorageSnapshot) SetStorageArray(v StorageGenericArrayRelationship) {
	o.StorageArray = &v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *StorageSnapshot) GetVolume() StorageVolumeRelationship {
	if o == nil || o.Volume == nil {
		var ret StorageVolumeRelationship
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSnapshot) GetVolumeOk() (*StorageVolumeRelationship, bool) {
	if o == nil || o.Volume == nil {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *StorageSnapshot) HasVolume() bool {
	if o != nil && o.Volume != nil {
		return true
	}

	return false
}

// SetVolume gets a reference to the given StorageVolumeRelationship and assigns it to the Volume field.
func (o *StorageSnapshot) SetVolume(v StorageVolumeRelationship) {
	o.Volume = &v
}

func (o StorageSnapshot) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.CreatedTime != nil {
		toSerialize["CreatedTime"] = o.CreatedTime
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.ProtectionGroupName != nil {
		toSerialize["ProtectionGroupName"] = o.ProtectionGroupName
	}
	if o.Size != nil {
		toSerialize["Size"] = o.Size
	}
	if o.Source != nil {
		toSerialize["Source"] = o.Source
	}
	if o.ProtectionGroupSnapshot != nil {
		toSerialize["ProtectionGroupSnapshot"] = o.ProtectionGroupSnapshot
	}
	if o.StorageArray != nil {
		toSerialize["StorageArray"] = o.StorageArray
	}
	if o.Volume != nil {
		toSerialize["Volume"] = o.Volume
	}
	return json.Marshal(toSerialize)
}

type NullableStorageSnapshot struct {
	value *StorageSnapshot
	isSet bool
}

func (v NullableStorageSnapshot) Get() *StorageSnapshot {
	return v.value
}

func (v *NullableStorageSnapshot) Set(val *StorageSnapshot) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageSnapshot) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageSnapshot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageSnapshot(val *StorageSnapshot) *NullableStorageSnapshot {
	return &NullableStorageSnapshot{value: val, isSet: true}
}

func (v NullableStorageSnapshot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageSnapshot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
