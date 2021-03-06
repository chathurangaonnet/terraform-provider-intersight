/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-06-17T02:04:50-07:00.
 *
 * API version: 1.0.9-1867
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
	"time"
)

// HclServiceStatus Status of the service indicatating if the service is up or under maintenance due to data update. Service will not be able serve any requests when the data is being updated. Collection will have only one document.
type HclServiceStatus struct {
	MoBaseMo
	// Version of the last modified exemption file.
	ExemptionFileVersion *string `json:"ExemptionFileVersion,omitempty"`
	// A field to uniquely identify the document with the status.
	Identity *string `json:"Identity,omitempty"`
	// The timestamp of the last modified record in the HCL tool database. Used to query and get updated records.
	LastHclDataModifiedTime *time.Time `json:"LastHclDataModifiedTime,omitempty"`
	// Checksum of the data dump used as the base for delta updates.
	LastImportedDataChecksum *string `json:"LastImportedDataChecksum,omitempty"`
	// Status of the service indicatating if the service is up or under maintenance due to data update.
	Status               *string `json:"Status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HclServiceStatus HclServiceStatus

// NewHclServiceStatus instantiates a new HclServiceStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHclServiceStatus() *HclServiceStatus {
	this := HclServiceStatus{}
	var status string = "Unknown"
	this.Status = &status
	return &this
}

// NewHclServiceStatusWithDefaults instantiates a new HclServiceStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHclServiceStatusWithDefaults() *HclServiceStatus {
	this := HclServiceStatus{}
	var status string = "Unknown"
	this.Status = &status
	return &this
}

// GetExemptionFileVersion returns the ExemptionFileVersion field value if set, zero value otherwise.
func (o *HclServiceStatus) GetExemptionFileVersion() string {
	if o == nil || o.ExemptionFileVersion == nil {
		var ret string
		return ret
	}
	return *o.ExemptionFileVersion
}

// GetExemptionFileVersionOk returns a tuple with the ExemptionFileVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclServiceStatus) GetExemptionFileVersionOk() (*string, bool) {
	if o == nil || o.ExemptionFileVersion == nil {
		return nil, false
	}
	return o.ExemptionFileVersion, true
}

// HasExemptionFileVersion returns a boolean if a field has been set.
func (o *HclServiceStatus) HasExemptionFileVersion() bool {
	if o != nil && o.ExemptionFileVersion != nil {
		return true
	}

	return false
}

// SetExemptionFileVersion gets a reference to the given string and assigns it to the ExemptionFileVersion field.
func (o *HclServiceStatus) SetExemptionFileVersion(v string) {
	o.ExemptionFileVersion = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *HclServiceStatus) GetIdentity() string {
	if o == nil || o.Identity == nil {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclServiceStatus) GetIdentityOk() (*string, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *HclServiceStatus) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *HclServiceStatus) SetIdentity(v string) {
	o.Identity = &v
}

// GetLastHclDataModifiedTime returns the LastHclDataModifiedTime field value if set, zero value otherwise.
func (o *HclServiceStatus) GetLastHclDataModifiedTime() time.Time {
	if o == nil || o.LastHclDataModifiedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastHclDataModifiedTime
}

// GetLastHclDataModifiedTimeOk returns a tuple with the LastHclDataModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclServiceStatus) GetLastHclDataModifiedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastHclDataModifiedTime == nil {
		return nil, false
	}
	return o.LastHclDataModifiedTime, true
}

// HasLastHclDataModifiedTime returns a boolean if a field has been set.
func (o *HclServiceStatus) HasLastHclDataModifiedTime() bool {
	if o != nil && o.LastHclDataModifiedTime != nil {
		return true
	}

	return false
}

// SetLastHclDataModifiedTime gets a reference to the given time.Time and assigns it to the LastHclDataModifiedTime field.
func (o *HclServiceStatus) SetLastHclDataModifiedTime(v time.Time) {
	o.LastHclDataModifiedTime = &v
}

// GetLastImportedDataChecksum returns the LastImportedDataChecksum field value if set, zero value otherwise.
func (o *HclServiceStatus) GetLastImportedDataChecksum() string {
	if o == nil || o.LastImportedDataChecksum == nil {
		var ret string
		return ret
	}
	return *o.LastImportedDataChecksum
}

// GetLastImportedDataChecksumOk returns a tuple with the LastImportedDataChecksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclServiceStatus) GetLastImportedDataChecksumOk() (*string, bool) {
	if o == nil || o.LastImportedDataChecksum == nil {
		return nil, false
	}
	return o.LastImportedDataChecksum, true
}

// HasLastImportedDataChecksum returns a boolean if a field has been set.
func (o *HclServiceStatus) HasLastImportedDataChecksum() bool {
	if o != nil && o.LastImportedDataChecksum != nil {
		return true
	}

	return false
}

// SetLastImportedDataChecksum gets a reference to the given string and assigns it to the LastImportedDataChecksum field.
func (o *HclServiceStatus) SetLastImportedDataChecksum(v string) {
	o.LastImportedDataChecksum = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HclServiceStatus) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclServiceStatus) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HclServiceStatus) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *HclServiceStatus) SetStatus(v string) {
	o.Status = &v
}

func (o HclServiceStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HclServiceStatus) UnmarshalJSON(bytes []byte) (err error) {
	type HclServiceStatusWithoutEmbeddedStruct struct {
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

	varHclServiceStatusWithoutEmbeddedStruct := HclServiceStatusWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHclServiceStatusWithoutEmbeddedStruct)
	if err == nil {
		varHclServiceStatus := _HclServiceStatus{}
		varHclServiceStatus.ExemptionFileVersion = varHclServiceStatusWithoutEmbeddedStruct.ExemptionFileVersion
		varHclServiceStatus.Identity = varHclServiceStatusWithoutEmbeddedStruct.Identity
		varHclServiceStatus.LastHclDataModifiedTime = varHclServiceStatusWithoutEmbeddedStruct.LastHclDataModifiedTime
		varHclServiceStatus.LastImportedDataChecksum = varHclServiceStatusWithoutEmbeddedStruct.LastImportedDataChecksum
		varHclServiceStatus.Status = varHclServiceStatusWithoutEmbeddedStruct.Status
		*o = HclServiceStatus(varHclServiceStatus)
	} else {
		return err
	}

	varHclServiceStatus := _HclServiceStatus{}

	err = json.Unmarshal(bytes, &varHclServiceStatus)
	if err == nil {
		o.MoBaseMo = varHclServiceStatus.MoBaseMo
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ExemptionFileVersion")
		delete(additionalProperties, "Identity")
		delete(additionalProperties, "LastHclDataModifiedTime")
		delete(additionalProperties, "LastImportedDataChecksum")
		delete(additionalProperties, "Status")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHclServiceStatus struct {
	value *HclServiceStatus
	isSet bool
}

func (v NullableHclServiceStatus) Get() *HclServiceStatus {
	return v.value
}

func (v *NullableHclServiceStatus) Set(val *HclServiceStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableHclServiceStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableHclServiceStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHclServiceStatus(val *HclServiceStatus) *NullableHclServiceStatus {
	return &NullableHclServiceStatus{value: val, isSet: true}
}

func (v NullableHclServiceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHclServiceStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
