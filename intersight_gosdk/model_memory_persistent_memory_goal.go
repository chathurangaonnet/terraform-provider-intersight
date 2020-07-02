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

// MemoryPersistentMemoryGoal A Persistent Memory Goal that needs to be applied on the associated servers through the policy. This would result in the creation of regions and allocation of volatile memory on the server.
type MemoryPersistentMemoryGoal struct {
	MoBaseComplexType
	// Volatile memory percentage.
	MemoryModePercentage *int32 `json:"MemoryModePercentage,omitempty"`
	// Type of the Persistent Memory configuration where the Persistent Memory Modules are combined in an interleaved set or not.
	PersistentMemoryType *string `json:"PersistentMemoryType,omitempty"`
	// CPU Socket ID to which this goal will be applied.
	SocketId *string `json:"SocketId,omitempty"`
}

// NewMemoryPersistentMemoryGoal instantiates a new MemoryPersistentMemoryGoal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemoryPersistentMemoryGoal() *MemoryPersistentMemoryGoal {
	this := MemoryPersistentMemoryGoal{}
	var persistentMemoryType string = "app-direct"
	this.PersistentMemoryType = &persistentMemoryType
	var socketId string = "All Sockets"
	this.SocketId = &socketId
	return &this
}

// NewMemoryPersistentMemoryGoalWithDefaults instantiates a new MemoryPersistentMemoryGoal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemoryPersistentMemoryGoalWithDefaults() *MemoryPersistentMemoryGoal {
	this := MemoryPersistentMemoryGoal{}
	var persistentMemoryType string = "app-direct"
	this.PersistentMemoryType = &persistentMemoryType
	var socketId string = "All Sockets"
	this.SocketId = &socketId
	return &this
}

// GetMemoryModePercentage returns the MemoryModePercentage field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryGoal) GetMemoryModePercentage() int32 {
	if o == nil || o.MemoryModePercentage == nil {
		var ret int32
		return ret
	}
	return *o.MemoryModePercentage
}

// GetMemoryModePercentageOk returns a tuple with the MemoryModePercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryGoal) GetMemoryModePercentageOk() (*int32, bool) {
	if o == nil || o.MemoryModePercentage == nil {
		return nil, false
	}
	return o.MemoryModePercentage, true
}

// HasMemoryModePercentage returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryGoal) HasMemoryModePercentage() bool {
	if o != nil && o.MemoryModePercentage != nil {
		return true
	}

	return false
}

// SetMemoryModePercentage gets a reference to the given int32 and assigns it to the MemoryModePercentage field.
func (o *MemoryPersistentMemoryGoal) SetMemoryModePercentage(v int32) {
	o.MemoryModePercentage = &v
}

// GetPersistentMemoryType returns the PersistentMemoryType field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryGoal) GetPersistentMemoryType() string {
	if o == nil || o.PersistentMemoryType == nil {
		var ret string
		return ret
	}
	return *o.PersistentMemoryType
}

// GetPersistentMemoryTypeOk returns a tuple with the PersistentMemoryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryGoal) GetPersistentMemoryTypeOk() (*string, bool) {
	if o == nil || o.PersistentMemoryType == nil {
		return nil, false
	}
	return o.PersistentMemoryType, true
}

// HasPersistentMemoryType returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryGoal) HasPersistentMemoryType() bool {
	if o != nil && o.PersistentMemoryType != nil {
		return true
	}

	return false
}

// SetPersistentMemoryType gets a reference to the given string and assigns it to the PersistentMemoryType field.
func (o *MemoryPersistentMemoryGoal) SetPersistentMemoryType(v string) {
	o.PersistentMemoryType = &v
}

// GetSocketId returns the SocketId field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryGoal) GetSocketId() string {
	if o == nil || o.SocketId == nil {
		var ret string
		return ret
	}
	return *o.SocketId
}

// GetSocketIdOk returns a tuple with the SocketId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryGoal) GetSocketIdOk() (*string, bool) {
	if o == nil || o.SocketId == nil {
		return nil, false
	}
	return o.SocketId, true
}

// HasSocketId returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryGoal) HasSocketId() bool {
	if o != nil && o.SocketId != nil {
		return true
	}

	return false
}

// SetSocketId gets a reference to the given string and assigns it to the SocketId field.
func (o *MemoryPersistentMemoryGoal) SetSocketId(v string) {
	o.SocketId = &v
}

func (o MemoryPersistentMemoryGoal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.MemoryModePercentage != nil {
		toSerialize["MemoryModePercentage"] = o.MemoryModePercentage
	}
	if o.PersistentMemoryType != nil {
		toSerialize["PersistentMemoryType"] = o.PersistentMemoryType
	}
	if o.SocketId != nil {
		toSerialize["SocketId"] = o.SocketId
	}
	return json.Marshal(toSerialize)
}

type NullableMemoryPersistentMemoryGoal struct {
	value *MemoryPersistentMemoryGoal
	isSet bool
}

func (v NullableMemoryPersistentMemoryGoal) Get() *MemoryPersistentMemoryGoal {
	return v.value
}

func (v *NullableMemoryPersistentMemoryGoal) Set(val *MemoryPersistentMemoryGoal) {
	v.value = val
	v.isSet = true
}

func (v NullableMemoryPersistentMemoryGoal) IsSet() bool {
	return v.isSet
}

func (v *NullableMemoryPersistentMemoryGoal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemoryPersistentMemoryGoal(val *MemoryPersistentMemoryGoal) *NullableMemoryPersistentMemoryGoal {
	return &NullableMemoryPersistentMemoryGoal{value: val, isSet: true}
}

func (v NullableMemoryPersistentMemoryGoal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemoryPersistentMemoryGoal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}