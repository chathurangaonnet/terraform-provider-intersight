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

// HclServiceStatusAllOf Definition of the list of properties defined in 'hcl.ServiceStatus', excluding properties defined in parent classes.
type HclServiceStatusAllOf struct {
	// Version of the last modified exemption file.
	ExemptionFileVersion *string `json:"ExemptionFileVersion,omitempty"`
	// A field to uniquely identify the document with the status.
	Identity *string `json:"Identity,omitempty"`
	// The timestamp of the last modified record in the HCL tool database. Used to query and get updated records.
	LastHclDataModifiedTime *time.Time `json:"LastHclDataModifiedTime,omitempty"`
	// Checksum of the data dump used as the base for delta updates.
	LastImportedDataChecksum *string `json:"LastImportedDataChecksum,omitempty"`
	// Status of the service indicatating if the service is up or under maintenance due to data update.
	Status *string `json:"Status,omitempty"`
}

// NewHclServiceStatusAllOf instantiates a new HclServiceStatusAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHclServiceStatusAllOf() *HclServiceStatusAllOf {
	this := HclServiceStatusAllOf{}
	var status string = "Unknown"
	this.Status = &status
	return &this
}

// NewHclServiceStatusAllOfWithDefaults instantiates a new HclServiceStatusAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHclServiceStatusAllOfWithDefaults() *HclServiceStatusAllOf {
	this := HclServiceStatusAllOf{}
	var status string = "Unknown"
	this.Status = &status
	return &this
}

// GetExemptionFileVersion returns the ExemptionFileVersion field value if set, zero value otherwise.
func (o *HclServiceStatusAllOf) GetExemptionFileVersion() string {
	if o == nil || o.ExemptionFileVersion == nil {
		var ret string
		return ret
	}
	return *o.ExemptionFileVersion
}

// GetExemptionFileVersionOk returns a tuple with the ExemptionFileVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclServiceStatusAllOf) GetExemptionFileVersionOk() (*string, bool) {
	if o == nil || o.ExemptionFileVersion == nil {
		return nil, false
	}
	return o.ExemptionFileVersion, true
}

// HasExemptionFileVersion returns a boolean if a field has been set.
func (o *HclServiceStatusAllOf) HasExemptionFileVersion() bool {
	if o != nil && o.ExemptionFileVersion != nil {
		return true
	}

	return false
}

// SetExemptionFileVersion gets a reference to the given string and assigns it to the ExemptionFileVersion field.
func (o *HclServiceStatusAllOf) SetExemptionFileVersion(v string) {
	o.ExemptionFileVersion = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *HclServiceStatusAllOf) GetIdentity() string {
	if o == nil || o.Identity == nil {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclServiceStatusAllOf) GetIdentityOk() (*string, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *HclServiceStatusAllOf) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *HclServiceStatusAllOf) SetIdentity(v string) {
	o.Identity = &v
}

// GetLastHclDataModifiedTime returns the LastHclDataModifiedTime field value if set, zero value otherwise.
func (o *HclServiceStatusAllOf) GetLastHclDataModifiedTime() time.Time {
	if o == nil || o.LastHclDataModifiedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastHclDataModifiedTime
}

// GetLastHclDataModifiedTimeOk returns a tuple with the LastHclDataModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclServiceStatusAllOf) GetLastHclDataModifiedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastHclDataModifiedTime == nil {
		return nil, false
	}
	return o.LastHclDataModifiedTime, true
}

// HasLastHclDataModifiedTime returns a boolean if a field has been set.
func (o *HclServiceStatusAllOf) HasLastHclDataModifiedTime() bool {
	if o != nil && o.LastHclDataModifiedTime != nil {
		return true
	}

	return false
}

// SetLastHclDataModifiedTime gets a reference to the given time.Time and assigns it to the LastHclDataModifiedTime field.
func (o *HclServiceStatusAllOf) SetLastHclDataModifiedTime(v time.Time) {
	o.LastHclDataModifiedTime = &v
}

// GetLastImportedDataChecksum returns the LastImportedDataChecksum field value if set, zero value otherwise.
func (o *HclServiceStatusAllOf) GetLastImportedDataChecksum() string {
	if o == nil || o.LastImportedDataChecksum == nil {
		var ret string
		return ret
	}
	return *o.LastImportedDataChecksum
}

// GetLastImportedDataChecksumOk returns a tuple with the LastImportedDataChecksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclServiceStatusAllOf) GetLastImportedDataChecksumOk() (*string, bool) {
	if o == nil || o.LastImportedDataChecksum == nil {
		return nil, false
	}
	return o.LastImportedDataChecksum, true
}

// HasLastImportedDataChecksum returns a boolean if a field has been set.
func (o *HclServiceStatusAllOf) HasLastImportedDataChecksum() bool {
	if o != nil && o.LastImportedDataChecksum != nil {
		return true
	}

	return false
}

// SetLastImportedDataChecksum gets a reference to the given string and assigns it to the LastImportedDataChecksum field.
func (o *HclServiceStatusAllOf) SetLastImportedDataChecksum(v string) {
	o.LastImportedDataChecksum = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HclServiceStatusAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclServiceStatusAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HclServiceStatusAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *HclServiceStatusAllOf) SetStatus(v string) {
	o.Status = &v
}

func (o HclServiceStatusAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExemptionFileVersion != nil {
		toSerialize["ExemptionFileVersion"] = o.ExemptionFileVersion
	}
	if o.Identity != nil {
		toSerialize["Identity"] = o.Identity
	}
	if o.LastHclDataModifiedTime != nil {
		toSerialize["LastHclDataModifiedTime"] = o.LastHclDataModifiedTime
	}
	if o.LastImportedDataChecksum != nil {
		toSerialize["LastImportedDataChecksum"] = o.LastImportedDataChecksum
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableHclServiceStatusAllOf struct {
	value *HclServiceStatusAllOf
	isSet bool
}

func (v NullableHclServiceStatusAllOf) Get() *HclServiceStatusAllOf {
	return v.value
}

func (v *NullableHclServiceStatusAllOf) Set(val *HclServiceStatusAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHclServiceStatusAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHclServiceStatusAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHclServiceStatusAllOf(val *HclServiceStatusAllOf) *NullableHclServiceStatusAllOf {
	return &NullableHclServiceStatusAllOf{value: val, isSet: true}
}

func (v NullableHclServiceStatusAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHclServiceStatusAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
