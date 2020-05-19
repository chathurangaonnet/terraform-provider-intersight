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

// NiatelemetryNiaLicenseState struct for NiatelemetryNiaLicenseState
type NiatelemetryNiaLicenseState struct {
	MoBaseMo
	// Features activated on device being inventoried
	FeatureActivated *string `json:"FeatureActivated,omitempty"`
	// Licenses activated on device being inventoried
	LicenseActivated *string `json:"LicenseActivated,omitempty"`
	// PID of device being inventoried
	PidType *string `json:"PidType,omitempty"`
	// Serial number of device being inventoried
	Serial *string                               `json:"Serial,omitempty"`
	Device *NiatelemetryNiaInventoryRelationship `json:"Device,omitempty"`
}

// NewNiatelemetryNiaLicenseState instantiates a new NiatelemetryNiaLicenseState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryNiaLicenseState() *NiatelemetryNiaLicenseState {
	this := NiatelemetryNiaLicenseState{}
	return &this
}

// NewNiatelemetryNiaLicenseStateWithDefaults instantiates a new NiatelemetryNiaLicenseState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryNiaLicenseStateWithDefaults() *NiatelemetryNiaLicenseState {
	this := NiatelemetryNiaLicenseState{}
	return &this
}

// GetFeatureActivated returns the FeatureActivated field value if set, zero value otherwise.
func (o *NiatelemetryNiaLicenseState) GetFeatureActivated() string {
	if o == nil || o.FeatureActivated == nil {
		var ret string
		return ret
	}
	return *o.FeatureActivated
}

// GetFeatureActivatedOk returns a tuple with the FeatureActivated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaLicenseState) GetFeatureActivatedOk() (*string, bool) {
	if o == nil || o.FeatureActivated == nil {
		return nil, false
	}
	return o.FeatureActivated, true
}

// HasFeatureActivated returns a boolean if a field has been set.
func (o *NiatelemetryNiaLicenseState) HasFeatureActivated() bool {
	if o != nil && o.FeatureActivated != nil {
		return true
	}

	return false
}

// SetFeatureActivated gets a reference to the given string and assigns it to the FeatureActivated field.
func (o *NiatelemetryNiaLicenseState) SetFeatureActivated(v string) {
	o.FeatureActivated = &v
}

// GetLicenseActivated returns the LicenseActivated field value if set, zero value otherwise.
func (o *NiatelemetryNiaLicenseState) GetLicenseActivated() string {
	if o == nil || o.LicenseActivated == nil {
		var ret string
		return ret
	}
	return *o.LicenseActivated
}

// GetLicenseActivatedOk returns a tuple with the LicenseActivated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaLicenseState) GetLicenseActivatedOk() (*string, bool) {
	if o == nil || o.LicenseActivated == nil {
		return nil, false
	}
	return o.LicenseActivated, true
}

// HasLicenseActivated returns a boolean if a field has been set.
func (o *NiatelemetryNiaLicenseState) HasLicenseActivated() bool {
	if o != nil && o.LicenseActivated != nil {
		return true
	}

	return false
}

// SetLicenseActivated gets a reference to the given string and assigns it to the LicenseActivated field.
func (o *NiatelemetryNiaLicenseState) SetLicenseActivated(v string) {
	o.LicenseActivated = &v
}

// GetPidType returns the PidType field value if set, zero value otherwise.
func (o *NiatelemetryNiaLicenseState) GetPidType() string {
	if o == nil || o.PidType == nil {
		var ret string
		return ret
	}
	return *o.PidType
}

// GetPidTypeOk returns a tuple with the PidType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaLicenseState) GetPidTypeOk() (*string, bool) {
	if o == nil || o.PidType == nil {
		return nil, false
	}
	return o.PidType, true
}

// HasPidType returns a boolean if a field has been set.
func (o *NiatelemetryNiaLicenseState) HasPidType() bool {
	if o != nil && o.PidType != nil {
		return true
	}

	return false
}

// SetPidType gets a reference to the given string and assigns it to the PidType field.
func (o *NiatelemetryNiaLicenseState) SetPidType(v string) {
	o.PidType = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *NiatelemetryNiaLicenseState) GetSerial() string {
	if o == nil || o.Serial == nil {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaLicenseState) GetSerialOk() (*string, bool) {
	if o == nil || o.Serial == nil {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *NiatelemetryNiaLicenseState) HasSerial() bool {
	if o != nil && o.Serial != nil {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *NiatelemetryNiaLicenseState) SetSerial(v string) {
	o.Serial = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *NiatelemetryNiaLicenseState) GetDevice() NiatelemetryNiaInventoryRelationship {
	if o == nil || o.Device == nil {
		var ret NiatelemetryNiaInventoryRelationship
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaLicenseState) GetDeviceOk() (*NiatelemetryNiaInventoryRelationship, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *NiatelemetryNiaLicenseState) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given NiatelemetryNiaInventoryRelationship and assigns it to the Device field.
func (o *NiatelemetryNiaLicenseState) SetDevice(v NiatelemetryNiaInventoryRelationship) {
	o.Device = &v
}

func (o NiatelemetryNiaLicenseState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.FeatureActivated != nil {
		toSerialize["FeatureActivated"] = o.FeatureActivated
	}
	if o.LicenseActivated != nil {
		toSerialize["LicenseActivated"] = o.LicenseActivated
	}
	if o.PidType != nil {
		toSerialize["PidType"] = o.PidType
	}
	if o.Serial != nil {
		toSerialize["Serial"] = o.Serial
	}
	if o.Device != nil {
		toSerialize["Device"] = o.Device
	}
	return json.Marshal(toSerialize)
}

// AsNiatelemetryNiaLicenseStateRelationship wraps this instance of NiatelemetryNiaLicenseState in NiatelemetryNiaLicenseStateRelationship
func (s *NiatelemetryNiaLicenseState) AsNiatelemetryNiaLicenseStateRelationship() NiatelemetryNiaLicenseStateRelationship {
	return NiatelemetryNiaLicenseStateRelationship{NiatelemetryNiaLicenseStateRelationshipInterface: s}
}

type NullableNiatelemetryNiaLicenseState struct {
	value *NiatelemetryNiaLicenseState
	isSet bool
}

func (v NullableNiatelemetryNiaLicenseState) Get() *NiatelemetryNiaLicenseState {
	return v.value
}

func (v *NullableNiatelemetryNiaLicenseState) Set(val *NiatelemetryNiaLicenseState) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryNiaLicenseState) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryNiaLicenseState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryNiaLicenseState(val *NiatelemetryNiaLicenseState) *NullableNiatelemetryNiaLicenseState {
	return &NullableNiatelemetryNiaLicenseState{value: val, isSet: true}
}

func (v NullableNiatelemetryNiaLicenseState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryNiaLicenseState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
