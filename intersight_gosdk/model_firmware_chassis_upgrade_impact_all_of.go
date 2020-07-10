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

// FirmwareChassisUpgradeImpactAllOf Definition of the list of properties defined in 'firmware.ChassisUpgradeImpact', excluding properties defined in parent classes.
type FirmwareChassisUpgradeImpactAllOf struct {
	ImpactDetail *[]FirmwareComponentImpact `json:"ImpactDetail,omitempty"`
	// Name of the chassis that will be affected by the upgrade.
	Name *string `json:"Name,omitempty"`
	// Details for the chassis that will be impacted by the upgrade.
	UserLabel *string `json:"UserLabel,omitempty"`
}

// NewFirmwareChassisUpgradeImpactAllOf instantiates a new FirmwareChassisUpgradeImpactAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareChassisUpgradeImpactAllOf() *FirmwareChassisUpgradeImpactAllOf {
	this := FirmwareChassisUpgradeImpactAllOf{}
	return &this
}

// NewFirmwareChassisUpgradeImpactAllOfWithDefaults instantiates a new FirmwareChassisUpgradeImpactAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareChassisUpgradeImpactAllOfWithDefaults() *FirmwareChassisUpgradeImpactAllOf {
	this := FirmwareChassisUpgradeImpactAllOf{}
	return &this
}

// GetImpactDetail returns the ImpactDetail field value if set, zero value otherwise.
func (o *FirmwareChassisUpgradeImpactAllOf) GetImpactDetail() []FirmwareComponentImpact {
	if o == nil || o.ImpactDetail == nil {
		var ret []FirmwareComponentImpact
		return ret
	}
	return *o.ImpactDetail
}

// GetImpactDetailOk returns a tuple with the ImpactDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareChassisUpgradeImpactAllOf) GetImpactDetailOk() (*[]FirmwareComponentImpact, bool) {
	if o == nil || o.ImpactDetail == nil {
		return nil, false
	}
	return o.ImpactDetail, true
}

// HasImpactDetail returns a boolean if a field has been set.
func (o *FirmwareChassisUpgradeImpactAllOf) HasImpactDetail() bool {
	if o != nil && o.ImpactDetail != nil {
		return true
	}

	return false
}

// SetImpactDetail gets a reference to the given []FirmwareComponentImpact and assigns it to the ImpactDetail field.
func (o *FirmwareChassisUpgradeImpactAllOf) SetImpactDetail(v []FirmwareComponentImpact) {
	o.ImpactDetail = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FirmwareChassisUpgradeImpactAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareChassisUpgradeImpactAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FirmwareChassisUpgradeImpactAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FirmwareChassisUpgradeImpactAllOf) SetName(v string) {
	o.Name = &v
}

// GetUserLabel returns the UserLabel field value if set, zero value otherwise.
func (o *FirmwareChassisUpgradeImpactAllOf) GetUserLabel() string {
	if o == nil || o.UserLabel == nil {
		var ret string
		return ret
	}
	return *o.UserLabel
}

// GetUserLabelOk returns a tuple with the UserLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareChassisUpgradeImpactAllOf) GetUserLabelOk() (*string, bool) {
	if o == nil || o.UserLabel == nil {
		return nil, false
	}
	return o.UserLabel, true
}

// HasUserLabel returns a boolean if a field has been set.
func (o *FirmwareChassisUpgradeImpactAllOf) HasUserLabel() bool {
	if o != nil && o.UserLabel != nil {
		return true
	}

	return false
}

// SetUserLabel gets a reference to the given string and assigns it to the UserLabel field.
func (o *FirmwareChassisUpgradeImpactAllOf) SetUserLabel(v string) {
	o.UserLabel = &v
}

func (o FirmwareChassisUpgradeImpactAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ImpactDetail != nil {
		toSerialize["ImpactDetail"] = o.ImpactDetail
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.UserLabel != nil {
		toSerialize["UserLabel"] = o.UserLabel
	}
	return json.Marshal(toSerialize)
}

type NullableFirmwareChassisUpgradeImpactAllOf struct {
	value *FirmwareChassisUpgradeImpactAllOf
	isSet bool
}

func (v NullableFirmwareChassisUpgradeImpactAllOf) Get() *FirmwareChassisUpgradeImpactAllOf {
	return v.value
}

func (v *NullableFirmwareChassisUpgradeImpactAllOf) Set(val *FirmwareChassisUpgradeImpactAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareChassisUpgradeImpactAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareChassisUpgradeImpactAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareChassisUpgradeImpactAllOf(val *FirmwareChassisUpgradeImpactAllOf) *NullableFirmwareChassisUpgradeImpactAllOf {
	return &NullableFirmwareChassisUpgradeImpactAllOf{value: val, isSet: true}
}

func (v NullableFirmwareChassisUpgradeImpactAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareChassisUpgradeImpactAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
