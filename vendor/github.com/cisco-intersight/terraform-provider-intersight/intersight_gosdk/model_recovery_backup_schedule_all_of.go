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
	"time"
)

// RecoveryBackupScheduleAllOf Definition of the list of properties defined in 'recovery.BackupSchedule', excluding properties defined in parent classes.
type RecoveryBackupScheduleAllOf struct {
	// The time at which the backup is to be run on a given day. This is used when the frequency unit is daily.
	ExecutionTime *time.Time `json:"ExecutionTime,omitempty"`
	// The frequency at which the backup schedule must run.
	FrequencyUnit *string `json:"FrequencyUnit,omitempty"`
	// The frequency, in hours, at which the backup schedule runs.
	Hours *int32 `json:"Hours,omitempty"`
}

// NewRecoveryBackupScheduleAllOf instantiates a new RecoveryBackupScheduleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoveryBackupScheduleAllOf() *RecoveryBackupScheduleAllOf {
	this := RecoveryBackupScheduleAllOf{}
	var frequencyUnit string = "Daily"
	this.FrequencyUnit = &frequencyUnit
	var hours int32 = 8
	this.Hours = &hours
	return &this
}

// NewRecoveryBackupScheduleAllOfWithDefaults instantiates a new RecoveryBackupScheduleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoveryBackupScheduleAllOfWithDefaults() *RecoveryBackupScheduleAllOf {
	this := RecoveryBackupScheduleAllOf{}
	var frequencyUnit string = "Daily"
	this.FrequencyUnit = &frequencyUnit
	var hours int32 = 8
	this.Hours = &hours
	return &this
}

// GetExecutionTime returns the ExecutionTime field value if set, zero value otherwise.
func (o *RecoveryBackupScheduleAllOf) GetExecutionTime() time.Time {
	if o == nil || o.ExecutionTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ExecutionTime
}

// GetExecutionTimeOk returns a tuple with the ExecutionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryBackupScheduleAllOf) GetExecutionTimeOk() (*time.Time, bool) {
	if o == nil || o.ExecutionTime == nil {
		return nil, false
	}
	return o.ExecutionTime, true
}

// HasExecutionTime returns a boolean if a field has been set.
func (o *RecoveryBackupScheduleAllOf) HasExecutionTime() bool {
	if o != nil && o.ExecutionTime != nil {
		return true
	}

	return false
}

// SetExecutionTime gets a reference to the given time.Time and assigns it to the ExecutionTime field.
func (o *RecoveryBackupScheduleAllOf) SetExecutionTime(v time.Time) {
	o.ExecutionTime = &v
}

// GetFrequencyUnit returns the FrequencyUnit field value if set, zero value otherwise.
func (o *RecoveryBackupScheduleAllOf) GetFrequencyUnit() string {
	if o == nil || o.FrequencyUnit == nil {
		var ret string
		return ret
	}
	return *o.FrequencyUnit
}

// GetFrequencyUnitOk returns a tuple with the FrequencyUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryBackupScheduleAllOf) GetFrequencyUnitOk() (*string, bool) {
	if o == nil || o.FrequencyUnit == nil {
		return nil, false
	}
	return o.FrequencyUnit, true
}

// HasFrequencyUnit returns a boolean if a field has been set.
func (o *RecoveryBackupScheduleAllOf) HasFrequencyUnit() bool {
	if o != nil && o.FrequencyUnit != nil {
		return true
	}

	return false
}

// SetFrequencyUnit gets a reference to the given string and assigns it to the FrequencyUnit field.
func (o *RecoveryBackupScheduleAllOf) SetFrequencyUnit(v string) {
	o.FrequencyUnit = &v
}

// GetHours returns the Hours field value if set, zero value otherwise.
func (o *RecoveryBackupScheduleAllOf) GetHours() int32 {
	if o == nil || o.Hours == nil {
		var ret int32
		return ret
	}
	return *o.Hours
}

// GetHoursOk returns a tuple with the Hours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryBackupScheduleAllOf) GetHoursOk() (*int32, bool) {
	if o == nil || o.Hours == nil {
		return nil, false
	}
	return o.Hours, true
}

// HasHours returns a boolean if a field has been set.
func (o *RecoveryBackupScheduleAllOf) HasHours() bool {
	if o != nil && o.Hours != nil {
		return true
	}

	return false
}

// SetHours gets a reference to the given int32 and assigns it to the Hours field.
func (o *RecoveryBackupScheduleAllOf) SetHours(v int32) {
	o.Hours = &v
}

func (o RecoveryBackupScheduleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExecutionTime != nil {
		toSerialize["ExecutionTime"] = o.ExecutionTime
	}
	if o.FrequencyUnit != nil {
		toSerialize["FrequencyUnit"] = o.FrequencyUnit
	}
	if o.Hours != nil {
		toSerialize["Hours"] = o.Hours
	}
	return json.Marshal(toSerialize)
}

type NullableRecoveryBackupScheduleAllOf struct {
	value *RecoveryBackupScheduleAllOf
	isSet bool
}

func (v NullableRecoveryBackupScheduleAllOf) Get() *RecoveryBackupScheduleAllOf {
	return v.value
}

func (v *NullableRecoveryBackupScheduleAllOf) Set(val *RecoveryBackupScheduleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoveryBackupScheduleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoveryBackupScheduleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoveryBackupScheduleAllOf(val *RecoveryBackupScheduleAllOf) *NullableRecoveryBackupScheduleAllOf {
	return &NullableRecoveryBackupScheduleAllOf{value: val, isSet: true}
}

func (v NullableRecoveryBackupScheduleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoveryBackupScheduleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
