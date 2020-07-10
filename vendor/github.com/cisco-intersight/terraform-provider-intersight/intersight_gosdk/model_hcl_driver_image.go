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

// HclDriverImage Collection used to store the driver ISO urls for each server based on how it is managed.
type HclDriverImage struct {
	MoBaseMo
	// URL of the driver ISO images.
	DriverIsoUrl *string `json:"DriverIsoUrl,omitempty"`
	// Type of the UCS version indicating whether it is a UCSM release vesion or a IMC release.
	ManagementType *string `json:"ManagementType,omitempty"`
	// Three part ID representing the server model as returned by UCSM/CIMC XML APIs.
	ServerPid *string `json:"ServerPid,omitempty"`
}

// NewHclDriverImage instantiates a new HclDriverImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHclDriverImage() *HclDriverImage {
	this := HclDriverImage{}
	var managementType string = "UCSM"
	this.ManagementType = &managementType
	return &this
}

// NewHclDriverImageWithDefaults instantiates a new HclDriverImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHclDriverImageWithDefaults() *HclDriverImage {
	this := HclDriverImage{}
	var managementType string = "UCSM"
	this.ManagementType = &managementType
	return &this
}

// GetDriverIsoUrl returns the DriverIsoUrl field value if set, zero value otherwise.
func (o *HclDriverImage) GetDriverIsoUrl() string {
	if o == nil || o.DriverIsoUrl == nil {
		var ret string
		return ret
	}
	return *o.DriverIsoUrl
}

// GetDriverIsoUrlOk returns a tuple with the DriverIsoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclDriverImage) GetDriverIsoUrlOk() (*string, bool) {
	if o == nil || o.DriverIsoUrl == nil {
		return nil, false
	}
	return o.DriverIsoUrl, true
}

// HasDriverIsoUrl returns a boolean if a field has been set.
func (o *HclDriverImage) HasDriverIsoUrl() bool {
	if o != nil && o.DriverIsoUrl != nil {
		return true
	}

	return false
}

// SetDriverIsoUrl gets a reference to the given string and assigns it to the DriverIsoUrl field.
func (o *HclDriverImage) SetDriverIsoUrl(v string) {
	o.DriverIsoUrl = &v
}

// GetManagementType returns the ManagementType field value if set, zero value otherwise.
func (o *HclDriverImage) GetManagementType() string {
	if o == nil || o.ManagementType == nil {
		var ret string
		return ret
	}
	return *o.ManagementType
}

// GetManagementTypeOk returns a tuple with the ManagementType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclDriverImage) GetManagementTypeOk() (*string, bool) {
	if o == nil || o.ManagementType == nil {
		return nil, false
	}
	return o.ManagementType, true
}

// HasManagementType returns a boolean if a field has been set.
func (o *HclDriverImage) HasManagementType() bool {
	if o != nil && o.ManagementType != nil {
		return true
	}

	return false
}

// SetManagementType gets a reference to the given string and assigns it to the ManagementType field.
func (o *HclDriverImage) SetManagementType(v string) {
	o.ManagementType = &v
}

// GetServerPid returns the ServerPid field value if set, zero value otherwise.
func (o *HclDriverImage) GetServerPid() string {
	if o == nil || o.ServerPid == nil {
		var ret string
		return ret
	}
	return *o.ServerPid
}

// GetServerPidOk returns a tuple with the ServerPid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclDriverImage) GetServerPidOk() (*string, bool) {
	if o == nil || o.ServerPid == nil {
		return nil, false
	}
	return o.ServerPid, true
}

// HasServerPid returns a boolean if a field has been set.
func (o *HclDriverImage) HasServerPid() bool {
	if o != nil && o.ServerPid != nil {
		return true
	}

	return false
}

// SetServerPid gets a reference to the given string and assigns it to the ServerPid field.
func (o *HclDriverImage) SetServerPid(v string) {
	o.ServerPid = &v
}

func (o HclDriverImage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.DriverIsoUrl != nil {
		toSerialize["DriverIsoUrl"] = o.DriverIsoUrl
	}
	if o.ManagementType != nil {
		toSerialize["ManagementType"] = o.ManagementType
	}
	if o.ServerPid != nil {
		toSerialize["ServerPid"] = o.ServerPid
	}
	return json.Marshal(toSerialize)
}

type NullableHclDriverImage struct {
	value *HclDriverImage
	isSet bool
}

func (v NullableHclDriverImage) Get() *HclDriverImage {
	return v.value
}

func (v *NullableHclDriverImage) Set(val *HclDriverImage) {
	v.value = val
	v.isSet = true
}

func (v NullableHclDriverImage) IsSet() bool {
	return v.isSet
}

func (v *NullableHclDriverImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHclDriverImage(val *HclDriverImage) *NullableHclDriverImage {
	return &NullableHclDriverImage{value: val, isSet: true}
}

func (v NullableHclDriverImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHclDriverImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
