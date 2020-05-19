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

// BiosUnit Bios Unit.
type BiosUnit struct {
	EquipmentBase
	InitSeq          *string                              `json:"InitSeq,omitempty"`
	InitTs           *string                              `json:"InitTs,omitempty"`
	ComputeBlade     *ComputeBladeRelationship            `json:"ComputeBlade,omitempty"`
	ComputeRackUnit  *ComputeRackUnitRelationship         `json:"ComputeRackUnit,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	// An array of relationships to firmwareRunningFirmware resources.
	RunningFirmware *[]FirmwareRunningFirmwareRelationship `json:"RunningFirmware,omitempty"`
}

// NewBiosUnit instantiates a new BiosUnit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBiosUnit() *BiosUnit {
	this := BiosUnit{}
	return &this
}

// NewBiosUnitWithDefaults instantiates a new BiosUnit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBiosUnitWithDefaults() *BiosUnit {
	this := BiosUnit{}
	return &this
}

// GetInitSeq returns the InitSeq field value if set, zero value otherwise.
func (o *BiosUnit) GetInitSeq() string {
	if o == nil || o.InitSeq == nil {
		var ret string
		return ret
	}
	return *o.InitSeq
}

// GetInitSeqOk returns a tuple with the InitSeq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosUnit) GetInitSeqOk() (*string, bool) {
	if o == nil || o.InitSeq == nil {
		return nil, false
	}
	return o.InitSeq, true
}

// HasInitSeq returns a boolean if a field has been set.
func (o *BiosUnit) HasInitSeq() bool {
	if o != nil && o.InitSeq != nil {
		return true
	}

	return false
}

// SetInitSeq gets a reference to the given string and assigns it to the InitSeq field.
func (o *BiosUnit) SetInitSeq(v string) {
	o.InitSeq = &v
}

// GetInitTs returns the InitTs field value if set, zero value otherwise.
func (o *BiosUnit) GetInitTs() string {
	if o == nil || o.InitTs == nil {
		var ret string
		return ret
	}
	return *o.InitTs
}

// GetInitTsOk returns a tuple with the InitTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosUnit) GetInitTsOk() (*string, bool) {
	if o == nil || o.InitTs == nil {
		return nil, false
	}
	return o.InitTs, true
}

// HasInitTs returns a boolean if a field has been set.
func (o *BiosUnit) HasInitTs() bool {
	if o != nil && o.InitTs != nil {
		return true
	}

	return false
}

// SetInitTs gets a reference to the given string and assigns it to the InitTs field.
func (o *BiosUnit) SetInitTs(v string) {
	o.InitTs = &v
}

// GetComputeBlade returns the ComputeBlade field value if set, zero value otherwise.
func (o *BiosUnit) GetComputeBlade() ComputeBladeRelationship {
	if o == nil || o.ComputeBlade == nil {
		var ret ComputeBladeRelationship
		return ret
	}
	return *o.ComputeBlade
}

// GetComputeBladeOk returns a tuple with the ComputeBlade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosUnit) GetComputeBladeOk() (*ComputeBladeRelationship, bool) {
	if o == nil || o.ComputeBlade == nil {
		return nil, false
	}
	return o.ComputeBlade, true
}

// HasComputeBlade returns a boolean if a field has been set.
func (o *BiosUnit) HasComputeBlade() bool {
	if o != nil && o.ComputeBlade != nil {
		return true
	}

	return false
}

// SetComputeBlade gets a reference to the given ComputeBladeRelationship and assigns it to the ComputeBlade field.
func (o *BiosUnit) SetComputeBlade(v ComputeBladeRelationship) {
	o.ComputeBlade = &v
}

// GetComputeRackUnit returns the ComputeRackUnit field value if set, zero value otherwise.
func (o *BiosUnit) GetComputeRackUnit() ComputeRackUnitRelationship {
	if o == nil || o.ComputeRackUnit == nil {
		var ret ComputeRackUnitRelationship
		return ret
	}
	return *o.ComputeRackUnit
}

// GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosUnit) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool) {
	if o == nil || o.ComputeRackUnit == nil {
		return nil, false
	}
	return o.ComputeRackUnit, true
}

// HasComputeRackUnit returns a boolean if a field has been set.
func (o *BiosUnit) HasComputeRackUnit() bool {
	if o != nil && o.ComputeRackUnit != nil {
		return true
	}

	return false
}

// SetComputeRackUnit gets a reference to the given ComputeRackUnitRelationship and assigns it to the ComputeRackUnit field.
func (o *BiosUnit) SetComputeRackUnit(v ComputeRackUnitRelationship) {
	o.ComputeRackUnit = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *BiosUnit) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosUnit) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *BiosUnit) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *BiosUnit) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetRunningFirmware returns the RunningFirmware field value if set, zero value otherwise.
func (o *BiosUnit) GetRunningFirmware() []FirmwareRunningFirmwareRelationship {
	if o == nil || o.RunningFirmware == nil {
		var ret []FirmwareRunningFirmwareRelationship
		return ret
	}
	return *o.RunningFirmware
}

// GetRunningFirmwareOk returns a tuple with the RunningFirmware field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosUnit) GetRunningFirmwareOk() (*[]FirmwareRunningFirmwareRelationship, bool) {
	if o == nil || o.RunningFirmware == nil {
		return nil, false
	}
	return o.RunningFirmware, true
}

// HasRunningFirmware returns a boolean if a field has been set.
func (o *BiosUnit) HasRunningFirmware() bool {
	if o != nil && o.RunningFirmware != nil {
		return true
	}

	return false
}

// SetRunningFirmware gets a reference to the given []FirmwareRunningFirmwareRelationship and assigns it to the RunningFirmware field.
func (o *BiosUnit) SetRunningFirmware(v []FirmwareRunningFirmwareRelationship) {
	o.RunningFirmware = &v
}

func (o BiosUnit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentBase, errEquipmentBase := json.Marshal(o.EquipmentBase)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	errEquipmentBase = json.Unmarshal([]byte(serializedEquipmentBase), &toSerialize)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	if o.InitSeq != nil {
		toSerialize["InitSeq"] = o.InitSeq
	}
	if o.InitTs != nil {
		toSerialize["InitTs"] = o.InitTs
	}
	if o.ComputeBlade != nil {
		toSerialize["ComputeBlade"] = o.ComputeBlade
	}
	if o.ComputeRackUnit != nil {
		toSerialize["ComputeRackUnit"] = o.ComputeRackUnit
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.RunningFirmware != nil {
		toSerialize["RunningFirmware"] = o.RunningFirmware
	}
	return json.Marshal(toSerialize)
}

// AsBiosUnitRelationship wraps this instance of BiosUnit in BiosUnitRelationship
func (s *BiosUnit) AsBiosUnitRelationship() BiosUnitRelationship {
	return BiosUnitRelationship{BiosUnitRelationshipInterface: s}
}

type NullableBiosUnit struct {
	value *BiosUnit
	isSet bool
}

func (v NullableBiosUnit) Get() *BiosUnit {
	return v.value
}

func (v *NullableBiosUnit) Set(val *BiosUnit) {
	v.value = val
	v.isSet = true
}

func (v NullableBiosUnit) IsSet() bool {
	return v.isSet
}

func (v *NullableBiosUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBiosUnit(val *BiosUnit) *NullableBiosUnit {
	return &NullableBiosUnit{value: val, isSet: true}
}

func (v NullableBiosUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBiosUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
