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

// StorageArrayController Common attributes for a Storage array controller. It can be a hardware or software unit which manages the physical storage disks available in the array.
type StorageArrayController struct {
	EquipmentBase
	// Storage array controller name.
	Name *string `json:"Name,omitempty"`
	// Controller running mode, Primary or Secondary.
	OperationalMode *string `json:"OperationalMode,omitempty"`
	// Status of the storage controller.
	Status *string `json:"Status,omitempty"`
	// Software version running on a storage controller.
	Version      *string                          `json:"Version,omitempty"`
	StorageArray *StorageGenericArrayRelationship `json:"StorageArray,omitempty"`
}

// NewStorageArrayController instantiates a new StorageArrayController object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageArrayController() *StorageArrayController {
	this := StorageArrayController{}
	var operationalMode string = "Unknown"
	this.OperationalMode = &operationalMode
	var status string = "Unknown"
	this.Status = &status
	return &this
}

// NewStorageArrayControllerWithDefaults instantiates a new StorageArrayController object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageArrayControllerWithDefaults() *StorageArrayController {
	this := StorageArrayController{}
	var operationalMode string = "Unknown"
	this.OperationalMode = &operationalMode
	var status string = "Unknown"
	this.Status = &status
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageArrayController) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageArrayController) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageArrayController) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageArrayController) SetName(v string) {
	o.Name = &v
}

// GetOperationalMode returns the OperationalMode field value if set, zero value otherwise.
func (o *StorageArrayController) GetOperationalMode() string {
	if o == nil || o.OperationalMode == nil {
		var ret string
		return ret
	}
	return *o.OperationalMode
}

// GetOperationalModeOk returns a tuple with the OperationalMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageArrayController) GetOperationalModeOk() (*string, bool) {
	if o == nil || o.OperationalMode == nil {
		return nil, false
	}
	return o.OperationalMode, true
}

// HasOperationalMode returns a boolean if a field has been set.
func (o *StorageArrayController) HasOperationalMode() bool {
	if o != nil && o.OperationalMode != nil {
		return true
	}

	return false
}

// SetOperationalMode gets a reference to the given string and assigns it to the OperationalMode field.
func (o *StorageArrayController) SetOperationalMode(v string) {
	o.OperationalMode = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *StorageArrayController) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageArrayController) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *StorageArrayController) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *StorageArrayController) SetStatus(v string) {
	o.Status = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *StorageArrayController) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageArrayController) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *StorageArrayController) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *StorageArrayController) SetVersion(v string) {
	o.Version = &v
}

// GetStorageArray returns the StorageArray field value if set, zero value otherwise.
func (o *StorageArrayController) GetStorageArray() StorageGenericArrayRelationship {
	if o == nil || o.StorageArray == nil {
		var ret StorageGenericArrayRelationship
		return ret
	}
	return *o.StorageArray
}

// GetStorageArrayOk returns a tuple with the StorageArray field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageArrayController) GetStorageArrayOk() (*StorageGenericArrayRelationship, bool) {
	if o == nil || o.StorageArray == nil {
		return nil, false
	}
	return o.StorageArray, true
}

// HasStorageArray returns a boolean if a field has been set.
func (o *StorageArrayController) HasStorageArray() bool {
	if o != nil && o.StorageArray != nil {
		return true
	}

	return false
}

// SetStorageArray gets a reference to the given StorageGenericArrayRelationship and assigns it to the StorageArray field.
func (o *StorageArrayController) SetStorageArray(v StorageGenericArrayRelationship) {
	o.StorageArray = &v
}

func (o StorageArrayController) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentBase, errEquipmentBase := json.Marshal(o.EquipmentBase)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	errEquipmentBase = json.Unmarshal([]byte(serializedEquipmentBase), &toSerialize)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.OperationalMode != nil {
		toSerialize["OperationalMode"] = o.OperationalMode
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.StorageArray != nil {
		toSerialize["StorageArray"] = o.StorageArray
	}
	return json.Marshal(toSerialize)
}

// AsStorageArrayControllerRelationship wraps this instance of StorageArrayController in StorageArrayControllerRelationship
func (s *StorageArrayController) AsStorageArrayControllerRelationship() StorageArrayControllerRelationship {
	return StorageArrayControllerRelationship{StorageArrayControllerRelationshipInterface: s}
}

type NullableStorageArrayController struct {
	value *StorageArrayController
	isSet bool
}

func (v NullableStorageArrayController) Get() *StorageArrayController {
	return v.value
}

func (v *NullableStorageArrayController) Set(val *StorageArrayController) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageArrayController) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageArrayController) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageArrayController(val *StorageArrayController) *NullableStorageArrayController {
	return &NullableStorageArrayController{value: val, isSet: true}
}

func (v NullableStorageArrayController) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageArrayController) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
