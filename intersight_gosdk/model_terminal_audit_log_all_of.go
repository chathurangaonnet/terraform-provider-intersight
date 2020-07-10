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

// TerminalAuditLogAllOf Definition of the list of properties defined in 'terminal.AuditLog', excluding properties defined in parent classes.
type TerminalAuditLogAllOf struct {
	// The time the terminal was closed. If terminal has not closed, value is zero time.
	EndTime *time.Time `json:"EndTime,omitempty"`
	// The time the terminal session was opened.
	StartTime          *time.Time                         `json:"StartTime,omitempty"`
	DeviceRegistration *AssetDeviceConnectionRelationship `json:"DeviceRegistration,omitempty"`
	User               *IamUserRelationship               `json:"User,omitempty"`
}

// NewTerminalAuditLogAllOf instantiates a new TerminalAuditLogAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerminalAuditLogAllOf() *TerminalAuditLogAllOf {
	this := TerminalAuditLogAllOf{}
	return &this
}

// NewTerminalAuditLogAllOfWithDefaults instantiates a new TerminalAuditLogAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerminalAuditLogAllOfWithDefaults() *TerminalAuditLogAllOf {
	this := TerminalAuditLogAllOf{}
	return &this
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *TerminalAuditLogAllOf) GetEndTime() time.Time {
	if o == nil || o.EndTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalAuditLogAllOf) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *TerminalAuditLogAllOf) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *TerminalAuditLogAllOf) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *TerminalAuditLogAllOf) GetStartTime() time.Time {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalAuditLogAllOf) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *TerminalAuditLogAllOf) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *TerminalAuditLogAllOf) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetDeviceRegistration returns the DeviceRegistration field value if set, zero value otherwise.
func (o *TerminalAuditLogAllOf) GetDeviceRegistration() AssetDeviceConnectionRelationship {
	if o == nil || o.DeviceRegistration == nil {
		var ret AssetDeviceConnectionRelationship
		return ret
	}
	return *o.DeviceRegistration
}

// GetDeviceRegistrationOk returns a tuple with the DeviceRegistration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalAuditLogAllOf) GetDeviceRegistrationOk() (*AssetDeviceConnectionRelationship, bool) {
	if o == nil || o.DeviceRegistration == nil {
		return nil, false
	}
	return o.DeviceRegistration, true
}

// HasDeviceRegistration returns a boolean if a field has been set.
func (o *TerminalAuditLogAllOf) HasDeviceRegistration() bool {
	if o != nil && o.DeviceRegistration != nil {
		return true
	}

	return false
}

// SetDeviceRegistration gets a reference to the given AssetDeviceConnectionRelationship and assigns it to the DeviceRegistration field.
func (o *TerminalAuditLogAllOf) SetDeviceRegistration(v AssetDeviceConnectionRelationship) {
	o.DeviceRegistration = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *TerminalAuditLogAllOf) GetUser() IamUserRelationship {
	if o == nil || o.User == nil {
		var ret IamUserRelationship
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalAuditLogAllOf) GetUserOk() (*IamUserRelationship, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *TerminalAuditLogAllOf) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given IamUserRelationship and assigns it to the User field.
func (o *TerminalAuditLogAllOf) SetUser(v IamUserRelationship) {
	o.User = &v
}

func (o TerminalAuditLogAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EndTime != nil {
		toSerialize["EndTime"] = o.EndTime
	}
	if o.StartTime != nil {
		toSerialize["StartTime"] = o.StartTime
	}
	if o.DeviceRegistration != nil {
		toSerialize["DeviceRegistration"] = o.DeviceRegistration
	}
	if o.User != nil {
		toSerialize["User"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableTerminalAuditLogAllOf struct {
	value *TerminalAuditLogAllOf
	isSet bool
}

func (v NullableTerminalAuditLogAllOf) Get() *TerminalAuditLogAllOf {
	return v.value
}

func (v *NullableTerminalAuditLogAllOf) Set(val *TerminalAuditLogAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminalAuditLogAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminalAuditLogAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminalAuditLogAllOf(val *TerminalAuditLogAllOf) *NullableTerminalAuditLogAllOf {
	return &NullableTerminalAuditLogAllOf{value: val, isSet: true}
}

func (v NullableTerminalAuditLogAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminalAuditLogAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
