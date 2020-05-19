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

// StoragePhysicalPortAllOf Definition of the list of properties defined in 'storage.PhysicalPort', excluding properties defined in parent classes.
type StoragePhysicalPortAllOf struct {
	// ISCSI qualified name applicable for ethernet port. Example - 'iqn.2008-05.com.storage:fnm00151300643-514f0c50141faf05'.
	Iqn *string `json:"Iqn,omitempty"`
	// Name of the physical port available in storage array.
	Name *string `json:"Name,omitempty"`
	// Operational speed of physical port measured in Gbps.
	Speed *int64 `json:"Speed,omitempty"`
	// Status of storage array port.
	Status *string `json:"Status,omitempty"`
	// Port type - possible values are FC, FCoE and iSCSI.
	Type *string `json:"Type,omitempty"`
	// World wide name of FC port. It is a combination of WWNN and WWPN represented in 128 bit hexadecimal formatted with colon. Example: '51:4f:0c:50:14:1f:af:01:51:4f:0c:51:14:1f:af:01'.
	Wwn *string `json:"Wwn,omitempty"`
	// World wide node name of FC port. Represented in 64 bit hex digits, formatted with colon. Example - '51:4f:0c:50:14:1f:af:01'.
	Wwnn *string `json:"Wwnn,omitempty"`
	// World wide port name of FC port. Represented in 64 bit hex digits, formatted with colon. Example - '51:4f:0c:51:14:1f:af:01'.
	Wwpn         *string                             `json:"Wwpn,omitempty"`
	Controller   *StorageArrayControllerRelationship `json:"Controller,omitempty"`
	StorageArray *StorageGenericArrayRelationship    `json:"StorageArray,omitempty"`
}

// NewStoragePhysicalPortAllOf instantiates a new StoragePhysicalPortAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoragePhysicalPortAllOf() *StoragePhysicalPortAllOf {
	this := StoragePhysicalPortAllOf{}
	var status string = "Unknown"
	this.Status = &status
	var type_ string = "FC"
	this.Type = &type_
	return &this
}

// NewStoragePhysicalPortAllOfWithDefaults instantiates a new StoragePhysicalPortAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoragePhysicalPortAllOfWithDefaults() *StoragePhysicalPortAllOf {
	this := StoragePhysicalPortAllOf{}
	var status string = "Unknown"
	this.Status = &status
	var type_ string = "FC"
	this.Type = &type_
	return &this
}

// GetIqn returns the Iqn field value if set, zero value otherwise.
func (o *StoragePhysicalPortAllOf) GetIqn() string {
	if o == nil || o.Iqn == nil {
		var ret string
		return ret
	}
	return *o.Iqn
}

// GetIqnOk returns a tuple with the Iqn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePhysicalPortAllOf) GetIqnOk() (*string, bool) {
	if o == nil || o.Iqn == nil {
		return nil, false
	}
	return o.Iqn, true
}

// HasIqn returns a boolean if a field has been set.
func (o *StoragePhysicalPortAllOf) HasIqn() bool {
	if o != nil && o.Iqn != nil {
		return true
	}

	return false
}

// SetIqn gets a reference to the given string and assigns it to the Iqn field.
func (o *StoragePhysicalPortAllOf) SetIqn(v string) {
	o.Iqn = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StoragePhysicalPortAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePhysicalPortAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StoragePhysicalPortAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StoragePhysicalPortAllOf) SetName(v string) {
	o.Name = &v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *StoragePhysicalPortAllOf) GetSpeed() int64 {
	if o == nil || o.Speed == nil {
		var ret int64
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePhysicalPortAllOf) GetSpeedOk() (*int64, bool) {
	if o == nil || o.Speed == nil {
		return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *StoragePhysicalPortAllOf) HasSpeed() bool {
	if o != nil && o.Speed != nil {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given int64 and assigns it to the Speed field.
func (o *StoragePhysicalPortAllOf) SetSpeed(v int64) {
	o.Speed = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *StoragePhysicalPortAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePhysicalPortAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *StoragePhysicalPortAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *StoragePhysicalPortAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *StoragePhysicalPortAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePhysicalPortAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *StoragePhysicalPortAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *StoragePhysicalPortAllOf) SetType(v string) {
	o.Type = &v
}

// GetWwn returns the Wwn field value if set, zero value otherwise.
func (o *StoragePhysicalPortAllOf) GetWwn() string {
	if o == nil || o.Wwn == nil {
		var ret string
		return ret
	}
	return *o.Wwn
}

// GetWwnOk returns a tuple with the Wwn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePhysicalPortAllOf) GetWwnOk() (*string, bool) {
	if o == nil || o.Wwn == nil {
		return nil, false
	}
	return o.Wwn, true
}

// HasWwn returns a boolean if a field has been set.
func (o *StoragePhysicalPortAllOf) HasWwn() bool {
	if o != nil && o.Wwn != nil {
		return true
	}

	return false
}

// SetWwn gets a reference to the given string and assigns it to the Wwn field.
func (o *StoragePhysicalPortAllOf) SetWwn(v string) {
	o.Wwn = &v
}

// GetWwnn returns the Wwnn field value if set, zero value otherwise.
func (o *StoragePhysicalPortAllOf) GetWwnn() string {
	if o == nil || o.Wwnn == nil {
		var ret string
		return ret
	}
	return *o.Wwnn
}

// GetWwnnOk returns a tuple with the Wwnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePhysicalPortAllOf) GetWwnnOk() (*string, bool) {
	if o == nil || o.Wwnn == nil {
		return nil, false
	}
	return o.Wwnn, true
}

// HasWwnn returns a boolean if a field has been set.
func (o *StoragePhysicalPortAllOf) HasWwnn() bool {
	if o != nil && o.Wwnn != nil {
		return true
	}

	return false
}

// SetWwnn gets a reference to the given string and assigns it to the Wwnn field.
func (o *StoragePhysicalPortAllOf) SetWwnn(v string) {
	o.Wwnn = &v
}

// GetWwpn returns the Wwpn field value if set, zero value otherwise.
func (o *StoragePhysicalPortAllOf) GetWwpn() string {
	if o == nil || o.Wwpn == nil {
		var ret string
		return ret
	}
	return *o.Wwpn
}

// GetWwpnOk returns a tuple with the Wwpn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePhysicalPortAllOf) GetWwpnOk() (*string, bool) {
	if o == nil || o.Wwpn == nil {
		return nil, false
	}
	return o.Wwpn, true
}

// HasWwpn returns a boolean if a field has been set.
func (o *StoragePhysicalPortAllOf) HasWwpn() bool {
	if o != nil && o.Wwpn != nil {
		return true
	}

	return false
}

// SetWwpn gets a reference to the given string and assigns it to the Wwpn field.
func (o *StoragePhysicalPortAllOf) SetWwpn(v string) {
	o.Wwpn = &v
}

// GetController returns the Controller field value if set, zero value otherwise.
func (o *StoragePhysicalPortAllOf) GetController() StorageArrayControllerRelationship {
	if o == nil || o.Controller == nil {
		var ret StorageArrayControllerRelationship
		return ret
	}
	return *o.Controller
}

// GetControllerOk returns a tuple with the Controller field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePhysicalPortAllOf) GetControllerOk() (*StorageArrayControllerRelationship, bool) {
	if o == nil || o.Controller == nil {
		return nil, false
	}
	return o.Controller, true
}

// HasController returns a boolean if a field has been set.
func (o *StoragePhysicalPortAllOf) HasController() bool {
	if o != nil && o.Controller != nil {
		return true
	}

	return false
}

// SetController gets a reference to the given StorageArrayControllerRelationship and assigns it to the Controller field.
func (o *StoragePhysicalPortAllOf) SetController(v StorageArrayControllerRelationship) {
	o.Controller = &v
}

// GetStorageArray returns the StorageArray field value if set, zero value otherwise.
func (o *StoragePhysicalPortAllOf) GetStorageArray() StorageGenericArrayRelationship {
	if o == nil || o.StorageArray == nil {
		var ret StorageGenericArrayRelationship
		return ret
	}
	return *o.StorageArray
}

// GetStorageArrayOk returns a tuple with the StorageArray field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePhysicalPortAllOf) GetStorageArrayOk() (*StorageGenericArrayRelationship, bool) {
	if o == nil || o.StorageArray == nil {
		return nil, false
	}
	return o.StorageArray, true
}

// HasStorageArray returns a boolean if a field has been set.
func (o *StoragePhysicalPortAllOf) HasStorageArray() bool {
	if o != nil && o.StorageArray != nil {
		return true
	}

	return false
}

// SetStorageArray gets a reference to the given StorageGenericArrayRelationship and assigns it to the StorageArray field.
func (o *StoragePhysicalPortAllOf) SetStorageArray(v StorageGenericArrayRelationship) {
	o.StorageArray = &v
}

func (o StoragePhysicalPortAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Iqn != nil {
		toSerialize["Iqn"] = o.Iqn
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Speed != nil {
		toSerialize["Speed"] = o.Speed
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.Wwn != nil {
		toSerialize["Wwn"] = o.Wwn
	}
	if o.Wwnn != nil {
		toSerialize["Wwnn"] = o.Wwnn
	}
	if o.Wwpn != nil {
		toSerialize["Wwpn"] = o.Wwpn
	}
	if o.Controller != nil {
		toSerialize["Controller"] = o.Controller
	}
	if o.StorageArray != nil {
		toSerialize["StorageArray"] = o.StorageArray
	}
	return json.Marshal(toSerialize)
}

type NullableStoragePhysicalPortAllOf struct {
	value *StoragePhysicalPortAllOf
	isSet bool
}

func (v NullableStoragePhysicalPortAllOf) Get() *StoragePhysicalPortAllOf {
	return v.value
}

func (v *NullableStoragePhysicalPortAllOf) Set(val *StoragePhysicalPortAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStoragePhysicalPortAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStoragePhysicalPortAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoragePhysicalPortAllOf(val *StoragePhysicalPortAllOf) *NullableStoragePhysicalPortAllOf {
	return &NullableStoragePhysicalPortAllOf{value: val, isSet: true}
}

func (v NullableStoragePhysicalPortAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoragePhysicalPortAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
