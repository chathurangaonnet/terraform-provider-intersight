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

// HyperflexDiskStatus Status details of virtual disk.
type HyperflexDiskStatus struct {
	MoBaseComplexType
	// Percentage of download completed.
	DownloadPercentage *string `json:"DownloadPercentage,omitempty"`
	// Current state of the virtual disk.
	State *string `json:"State,omitempty"`
	// Identity of the Volume associated with virtual machine disk.
	VolumeHandle         *string `json:"VolumeHandle,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexDiskStatus HyperflexDiskStatus

// NewHyperflexDiskStatus instantiates a new HyperflexDiskStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexDiskStatus() *HyperflexDiskStatus {
	this := HyperflexDiskStatus{}
	var state string = "Unknown"
	this.State = &state
	return &this
}

// NewHyperflexDiskStatusWithDefaults instantiates a new HyperflexDiskStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexDiskStatusWithDefaults() *HyperflexDiskStatus {
	this := HyperflexDiskStatus{}
	var state string = "Unknown"
	this.State = &state
	return &this
}

// GetDownloadPercentage returns the DownloadPercentage field value if set, zero value otherwise.
func (o *HyperflexDiskStatus) GetDownloadPercentage() string {
	if o == nil || o.DownloadPercentage == nil {
		var ret string
		return ret
	}
	return *o.DownloadPercentage
}

// GetDownloadPercentageOk returns a tuple with the DownloadPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexDiskStatus) GetDownloadPercentageOk() (*string, bool) {
	if o == nil || o.DownloadPercentage == nil {
		return nil, false
	}
	return o.DownloadPercentage, true
}

// HasDownloadPercentage returns a boolean if a field has been set.
func (o *HyperflexDiskStatus) HasDownloadPercentage() bool {
	if o != nil && o.DownloadPercentage != nil {
		return true
	}

	return false
}

// SetDownloadPercentage gets a reference to the given string and assigns it to the DownloadPercentage field.
func (o *HyperflexDiskStatus) SetDownloadPercentage(v string) {
	o.DownloadPercentage = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *HyperflexDiskStatus) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexDiskStatus) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *HyperflexDiskStatus) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *HyperflexDiskStatus) SetState(v string) {
	o.State = &v
}

// GetVolumeHandle returns the VolumeHandle field value if set, zero value otherwise.
func (o *HyperflexDiskStatus) GetVolumeHandle() string {
	if o == nil || o.VolumeHandle == nil {
		var ret string
		return ret
	}
	return *o.VolumeHandle
}

// GetVolumeHandleOk returns a tuple with the VolumeHandle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexDiskStatus) GetVolumeHandleOk() (*string, bool) {
	if o == nil || o.VolumeHandle == nil {
		return nil, false
	}
	return o.VolumeHandle, true
}

// HasVolumeHandle returns a boolean if a field has been set.
func (o *HyperflexDiskStatus) HasVolumeHandle() bool {
	if o != nil && o.VolumeHandle != nil {
		return true
	}

	return false
}

// SetVolumeHandle gets a reference to the given string and assigns it to the VolumeHandle field.
func (o *HyperflexDiskStatus) SetVolumeHandle(v string) {
	o.VolumeHandle = &v
}

func (o HyperflexDiskStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.DownloadPercentage != nil {
		toSerialize["DownloadPercentage"] = o.DownloadPercentage
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.VolumeHandle != nil {
		toSerialize["VolumeHandle"] = o.VolumeHandle
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexDiskStatus) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexDiskStatus := _HyperflexDiskStatus{}

	if err = json.Unmarshal(bytes, &varHyperflexDiskStatus); err == nil {
		*o = HyperflexDiskStatus(varHyperflexDiskStatus)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "DownloadPercentage")
		delete(additionalProperties, "State")
		delete(additionalProperties, "VolumeHandle")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexDiskStatus struct {
	value *HyperflexDiskStatus
	isSet bool
}

func (v NullableHyperflexDiskStatus) Get() *HyperflexDiskStatus {
	return v.value
}

func (v *NullableHyperflexDiskStatus) Set(val *HyperflexDiskStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexDiskStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexDiskStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexDiskStatus(val *HyperflexDiskStatus) *NullableHyperflexDiskStatus {
	return &NullableHyperflexDiskStatus{value: val, isSet: true}
}

func (v NullableHyperflexDiskStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexDiskStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}