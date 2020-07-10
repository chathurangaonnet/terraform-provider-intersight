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

// HclFirmwareAllOf Definition of the list of properties defined in 'hcl.Firmware', excluding properties defined in parent classes.
type HclFirmwareAllOf struct {
	// Protocol for which the driver is provided. E.g.  enic, fnic, lsi_mr3.
	DriverName *string `json:"DriverName,omitempty"`
	// Version of the Driver supported.
	DriverVersion *string `json:"DriverVersion,omitempty"`
	// Error code for the support status.
	ErrorCode *string `json:"ErrorCode,omitempty"`
	// Firmware version of the product/adapter supported.
	FirmwareVersion *string `json:"FirmwareVersion,omitempty"`
	// Identifier of the firmware.
	Id *string `json:"Id,omitempty"`
	// True if the driver is latest recommended driver.
	LatestDriver *bool `json:"LatestDriver,omitempty"`
	// True if the firmware is latest recommended firmware.
	LatestFirmware *bool `json:"LatestFirmware,omitempty"`
}

// NewHclFirmwareAllOf instantiates a new HclFirmwareAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHclFirmwareAllOf() *HclFirmwareAllOf {
	this := HclFirmwareAllOf{}
	var errorCode string = "Success"
	this.ErrorCode = &errorCode
	return &this
}

// NewHclFirmwareAllOfWithDefaults instantiates a new HclFirmwareAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHclFirmwareAllOfWithDefaults() *HclFirmwareAllOf {
	this := HclFirmwareAllOf{}
	var errorCode string = "Success"
	this.ErrorCode = &errorCode
	return &this
}

// GetDriverName returns the DriverName field value if set, zero value otherwise.
func (o *HclFirmwareAllOf) GetDriverName() string {
	if o == nil || o.DriverName == nil {
		var ret string
		return ret
	}
	return *o.DriverName
}

// GetDriverNameOk returns a tuple with the DriverName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclFirmwareAllOf) GetDriverNameOk() (*string, bool) {
	if o == nil || o.DriverName == nil {
		return nil, false
	}
	return o.DriverName, true
}

// HasDriverName returns a boolean if a field has been set.
func (o *HclFirmwareAllOf) HasDriverName() bool {
	if o != nil && o.DriverName != nil {
		return true
	}

	return false
}

// SetDriverName gets a reference to the given string and assigns it to the DriverName field.
func (o *HclFirmwareAllOf) SetDriverName(v string) {
	o.DriverName = &v
}

// GetDriverVersion returns the DriverVersion field value if set, zero value otherwise.
func (o *HclFirmwareAllOf) GetDriverVersion() string {
	if o == nil || o.DriverVersion == nil {
		var ret string
		return ret
	}
	return *o.DriverVersion
}

// GetDriverVersionOk returns a tuple with the DriverVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclFirmwareAllOf) GetDriverVersionOk() (*string, bool) {
	if o == nil || o.DriverVersion == nil {
		return nil, false
	}
	return o.DriverVersion, true
}

// HasDriverVersion returns a boolean if a field has been set.
func (o *HclFirmwareAllOf) HasDriverVersion() bool {
	if o != nil && o.DriverVersion != nil {
		return true
	}

	return false
}

// SetDriverVersion gets a reference to the given string and assigns it to the DriverVersion field.
func (o *HclFirmwareAllOf) SetDriverVersion(v string) {
	o.DriverVersion = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *HclFirmwareAllOf) GetErrorCode() string {
	if o == nil || o.ErrorCode == nil {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclFirmwareAllOf) GetErrorCodeOk() (*string, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *HclFirmwareAllOf) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *HclFirmwareAllOf) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetFirmwareVersion returns the FirmwareVersion field value if set, zero value otherwise.
func (o *HclFirmwareAllOf) GetFirmwareVersion() string {
	if o == nil || o.FirmwareVersion == nil {
		var ret string
		return ret
	}
	return *o.FirmwareVersion
}

// GetFirmwareVersionOk returns a tuple with the FirmwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclFirmwareAllOf) GetFirmwareVersionOk() (*string, bool) {
	if o == nil || o.FirmwareVersion == nil {
		return nil, false
	}
	return o.FirmwareVersion, true
}

// HasFirmwareVersion returns a boolean if a field has been set.
func (o *HclFirmwareAllOf) HasFirmwareVersion() bool {
	if o != nil && o.FirmwareVersion != nil {
		return true
	}

	return false
}

// SetFirmwareVersion gets a reference to the given string and assigns it to the FirmwareVersion field.
func (o *HclFirmwareAllOf) SetFirmwareVersion(v string) {
	o.FirmwareVersion = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HclFirmwareAllOf) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclFirmwareAllOf) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HclFirmwareAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *HclFirmwareAllOf) SetId(v string) {
	o.Id = &v
}

// GetLatestDriver returns the LatestDriver field value if set, zero value otherwise.
func (o *HclFirmwareAllOf) GetLatestDriver() bool {
	if o == nil || o.LatestDriver == nil {
		var ret bool
		return ret
	}
	return *o.LatestDriver
}

// GetLatestDriverOk returns a tuple with the LatestDriver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclFirmwareAllOf) GetLatestDriverOk() (*bool, bool) {
	if o == nil || o.LatestDriver == nil {
		return nil, false
	}
	return o.LatestDriver, true
}

// HasLatestDriver returns a boolean if a field has been set.
func (o *HclFirmwareAllOf) HasLatestDriver() bool {
	if o != nil && o.LatestDriver != nil {
		return true
	}

	return false
}

// SetLatestDriver gets a reference to the given bool and assigns it to the LatestDriver field.
func (o *HclFirmwareAllOf) SetLatestDriver(v bool) {
	o.LatestDriver = &v
}

// GetLatestFirmware returns the LatestFirmware field value if set, zero value otherwise.
func (o *HclFirmwareAllOf) GetLatestFirmware() bool {
	if o == nil || o.LatestFirmware == nil {
		var ret bool
		return ret
	}
	return *o.LatestFirmware
}

// GetLatestFirmwareOk returns a tuple with the LatestFirmware field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclFirmwareAllOf) GetLatestFirmwareOk() (*bool, bool) {
	if o == nil || o.LatestFirmware == nil {
		return nil, false
	}
	return o.LatestFirmware, true
}

// HasLatestFirmware returns a boolean if a field has been set.
func (o *HclFirmwareAllOf) HasLatestFirmware() bool {
	if o != nil && o.LatestFirmware != nil {
		return true
	}

	return false
}

// SetLatestFirmware gets a reference to the given bool and assigns it to the LatestFirmware field.
func (o *HclFirmwareAllOf) SetLatestFirmware(v bool) {
	o.LatestFirmware = &v
}

func (o HclFirmwareAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DriverName != nil {
		toSerialize["DriverName"] = o.DriverName
	}
	if o.DriverVersion != nil {
		toSerialize["DriverVersion"] = o.DriverVersion
	}
	if o.ErrorCode != nil {
		toSerialize["ErrorCode"] = o.ErrorCode
	}
	if o.FirmwareVersion != nil {
		toSerialize["FirmwareVersion"] = o.FirmwareVersion
	}
	if o.Id != nil {
		toSerialize["Id"] = o.Id
	}
	if o.LatestDriver != nil {
		toSerialize["LatestDriver"] = o.LatestDriver
	}
	if o.LatestFirmware != nil {
		toSerialize["LatestFirmware"] = o.LatestFirmware
	}
	return json.Marshal(toSerialize)
}

type NullableHclFirmwareAllOf struct {
	value *HclFirmwareAllOf
	isSet bool
}

func (v NullableHclFirmwareAllOf) Get() *HclFirmwareAllOf {
	return v.value
}

func (v *NullableHclFirmwareAllOf) Set(val *HclFirmwareAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHclFirmwareAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHclFirmwareAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHclFirmwareAllOf(val *HclFirmwareAllOf) *NullableHclFirmwareAllOf {
	return &NullableHclFirmwareAllOf{value: val, isSet: true}
}

func (v NullableHclFirmwareAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHclFirmwareAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
