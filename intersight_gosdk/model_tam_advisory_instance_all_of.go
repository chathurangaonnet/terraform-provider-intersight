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

// TamAdvisoryInstanceAllOf Definition of the list of properties defined in 'tam.AdvisoryInstance', excluding properties defined in parent classes.
type TamAdvisoryInstanceAllOf struct {
	// Moid of the Intersight MO affected by the alert. Deprecated now and will be removed in subsequent releases.
	AffectedObjectMoid *string `json:"AffectedObjectMoid,omitempty"`
	// Object type of the Intersight MO affected by the alert. Deprecated now and will be removed in subsequent releases.
	AffectedObjectType *string `json:"AffectedObjectType,omitempty"`
	// Timestamp when a state change was observed on this advisory instnace.
	LastStateChangeTime *time.Time `json:"LastStateChangeTime,omitempty"`
	// Timestamp when this advisory was last evaluated.
	LastVerifiedTime *time.Time `json:"LastVerifiedTime,omitempty"`
	// Current state of the advisory instance (Active/Cleared/Unknown etc.).
	State              *string                              `json:"State,omitempty"`
	Advisory           *TamAdvisoryRelationship             `json:"Advisory,omitempty"`
	AffectedObject     *MoBaseMoRelationship                `json:"AffectedObject,omitempty"`
	DeviceRegistration *AssetDeviceRegistrationRelationship `json:"DeviceRegistration,omitempty"`
}

// NewTamAdvisoryInstanceAllOf instantiates a new TamAdvisoryInstanceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTamAdvisoryInstanceAllOf() *TamAdvisoryInstanceAllOf {
	this := TamAdvisoryInstanceAllOf{}
	var state string = "unknown"
	this.State = &state
	return &this
}

// NewTamAdvisoryInstanceAllOfWithDefaults instantiates a new TamAdvisoryInstanceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTamAdvisoryInstanceAllOfWithDefaults() *TamAdvisoryInstanceAllOf {
	this := TamAdvisoryInstanceAllOf{}
	var state string = "unknown"
	this.State = &state
	return &this
}

// GetAffectedObjectMoid returns the AffectedObjectMoid field value if set, zero value otherwise.
func (o *TamAdvisoryInstanceAllOf) GetAffectedObjectMoid() string {
	if o == nil || o.AffectedObjectMoid == nil {
		var ret string
		return ret
	}
	return *o.AffectedObjectMoid
}

// GetAffectedObjectMoidOk returns a tuple with the AffectedObjectMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAdvisoryInstanceAllOf) GetAffectedObjectMoidOk() (*string, bool) {
	if o == nil || o.AffectedObjectMoid == nil {
		return nil, false
	}
	return o.AffectedObjectMoid, true
}

// HasAffectedObjectMoid returns a boolean if a field has been set.
func (o *TamAdvisoryInstanceAllOf) HasAffectedObjectMoid() bool {
	if o != nil && o.AffectedObjectMoid != nil {
		return true
	}

	return false
}

// SetAffectedObjectMoid gets a reference to the given string and assigns it to the AffectedObjectMoid field.
func (o *TamAdvisoryInstanceAllOf) SetAffectedObjectMoid(v string) {
	o.AffectedObjectMoid = &v
}

// GetAffectedObjectType returns the AffectedObjectType field value if set, zero value otherwise.
func (o *TamAdvisoryInstanceAllOf) GetAffectedObjectType() string {
	if o == nil || o.AffectedObjectType == nil {
		var ret string
		return ret
	}
	return *o.AffectedObjectType
}

// GetAffectedObjectTypeOk returns a tuple with the AffectedObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAdvisoryInstanceAllOf) GetAffectedObjectTypeOk() (*string, bool) {
	if o == nil || o.AffectedObjectType == nil {
		return nil, false
	}
	return o.AffectedObjectType, true
}

// HasAffectedObjectType returns a boolean if a field has been set.
func (o *TamAdvisoryInstanceAllOf) HasAffectedObjectType() bool {
	if o != nil && o.AffectedObjectType != nil {
		return true
	}

	return false
}

// SetAffectedObjectType gets a reference to the given string and assigns it to the AffectedObjectType field.
func (o *TamAdvisoryInstanceAllOf) SetAffectedObjectType(v string) {
	o.AffectedObjectType = &v
}

// GetLastStateChangeTime returns the LastStateChangeTime field value if set, zero value otherwise.
func (o *TamAdvisoryInstanceAllOf) GetLastStateChangeTime() time.Time {
	if o == nil || o.LastStateChangeTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastStateChangeTime
}

// GetLastStateChangeTimeOk returns a tuple with the LastStateChangeTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAdvisoryInstanceAllOf) GetLastStateChangeTimeOk() (*time.Time, bool) {
	if o == nil || o.LastStateChangeTime == nil {
		return nil, false
	}
	return o.LastStateChangeTime, true
}

// HasLastStateChangeTime returns a boolean if a field has been set.
func (o *TamAdvisoryInstanceAllOf) HasLastStateChangeTime() bool {
	if o != nil && o.LastStateChangeTime != nil {
		return true
	}

	return false
}

// SetLastStateChangeTime gets a reference to the given time.Time and assigns it to the LastStateChangeTime field.
func (o *TamAdvisoryInstanceAllOf) SetLastStateChangeTime(v time.Time) {
	o.LastStateChangeTime = &v
}

// GetLastVerifiedTime returns the LastVerifiedTime field value if set, zero value otherwise.
func (o *TamAdvisoryInstanceAllOf) GetLastVerifiedTime() time.Time {
	if o == nil || o.LastVerifiedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastVerifiedTime
}

// GetLastVerifiedTimeOk returns a tuple with the LastVerifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAdvisoryInstanceAllOf) GetLastVerifiedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastVerifiedTime == nil {
		return nil, false
	}
	return o.LastVerifiedTime, true
}

// HasLastVerifiedTime returns a boolean if a field has been set.
func (o *TamAdvisoryInstanceAllOf) HasLastVerifiedTime() bool {
	if o != nil && o.LastVerifiedTime != nil {
		return true
	}

	return false
}

// SetLastVerifiedTime gets a reference to the given time.Time and assigns it to the LastVerifiedTime field.
func (o *TamAdvisoryInstanceAllOf) SetLastVerifiedTime(v time.Time) {
	o.LastVerifiedTime = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *TamAdvisoryInstanceAllOf) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAdvisoryInstanceAllOf) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *TamAdvisoryInstanceAllOf) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *TamAdvisoryInstanceAllOf) SetState(v string) {
	o.State = &v
}

// GetAdvisory returns the Advisory field value if set, zero value otherwise.
func (o *TamAdvisoryInstanceAllOf) GetAdvisory() TamAdvisoryRelationship {
	if o == nil || o.Advisory == nil {
		var ret TamAdvisoryRelationship
		return ret
	}
	return *o.Advisory
}

// GetAdvisoryOk returns a tuple with the Advisory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAdvisoryInstanceAllOf) GetAdvisoryOk() (*TamAdvisoryRelationship, bool) {
	if o == nil || o.Advisory == nil {
		return nil, false
	}
	return o.Advisory, true
}

// HasAdvisory returns a boolean if a field has been set.
func (o *TamAdvisoryInstanceAllOf) HasAdvisory() bool {
	if o != nil && o.Advisory != nil {
		return true
	}

	return false
}

// SetAdvisory gets a reference to the given TamAdvisoryRelationship and assigns it to the Advisory field.
func (o *TamAdvisoryInstanceAllOf) SetAdvisory(v TamAdvisoryRelationship) {
	o.Advisory = &v
}

// GetAffectedObject returns the AffectedObject field value if set, zero value otherwise.
func (o *TamAdvisoryInstanceAllOf) GetAffectedObject() MoBaseMoRelationship {
	if o == nil || o.AffectedObject == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.AffectedObject
}

// GetAffectedObjectOk returns a tuple with the AffectedObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAdvisoryInstanceAllOf) GetAffectedObjectOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.AffectedObject == nil {
		return nil, false
	}
	return o.AffectedObject, true
}

// HasAffectedObject returns a boolean if a field has been set.
func (o *TamAdvisoryInstanceAllOf) HasAffectedObject() bool {
	if o != nil && o.AffectedObject != nil {
		return true
	}

	return false
}

// SetAffectedObject gets a reference to the given MoBaseMoRelationship and assigns it to the AffectedObject field.
func (o *TamAdvisoryInstanceAllOf) SetAffectedObject(v MoBaseMoRelationship) {
	o.AffectedObject = &v
}

// GetDeviceRegistration returns the DeviceRegistration field value if set, zero value otherwise.
func (o *TamAdvisoryInstanceAllOf) GetDeviceRegistration() AssetDeviceRegistrationRelationship {
	if o == nil || o.DeviceRegistration == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.DeviceRegistration
}

// GetDeviceRegistrationOk returns a tuple with the DeviceRegistration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAdvisoryInstanceAllOf) GetDeviceRegistrationOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.DeviceRegistration == nil {
		return nil, false
	}
	return o.DeviceRegistration, true
}

// HasDeviceRegistration returns a boolean if a field has been set.
func (o *TamAdvisoryInstanceAllOf) HasDeviceRegistration() bool {
	if o != nil && o.DeviceRegistration != nil {
		return true
	}

	return false
}

// SetDeviceRegistration gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the DeviceRegistration field.
func (o *TamAdvisoryInstanceAllOf) SetDeviceRegistration(v AssetDeviceRegistrationRelationship) {
	o.DeviceRegistration = &v
}

func (o TamAdvisoryInstanceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AffectedObjectMoid != nil {
		toSerialize["AffectedObjectMoid"] = o.AffectedObjectMoid
	}
	if o.AffectedObjectType != nil {
		toSerialize["AffectedObjectType"] = o.AffectedObjectType
	}
	if o.LastStateChangeTime != nil {
		toSerialize["LastStateChangeTime"] = o.LastStateChangeTime
	}
	if o.LastVerifiedTime != nil {
		toSerialize["LastVerifiedTime"] = o.LastVerifiedTime
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.Advisory != nil {
		toSerialize["Advisory"] = o.Advisory
	}
	if o.AffectedObject != nil {
		toSerialize["AffectedObject"] = o.AffectedObject
	}
	if o.DeviceRegistration != nil {
		toSerialize["DeviceRegistration"] = o.DeviceRegistration
	}
	return json.Marshal(toSerialize)
}

type NullableTamAdvisoryInstanceAllOf struct {
	value *TamAdvisoryInstanceAllOf
	isSet bool
}

func (v NullableTamAdvisoryInstanceAllOf) Get() *TamAdvisoryInstanceAllOf {
	return v.value
}

func (v *NullableTamAdvisoryInstanceAllOf) Set(val *TamAdvisoryInstanceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTamAdvisoryInstanceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTamAdvisoryInstanceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTamAdvisoryInstanceAllOf(val *TamAdvisoryInstanceAllOf) *NullableTamAdvisoryInstanceAllOf {
	return &NullableTamAdvisoryInstanceAllOf{value: val, isSet: true}
}

func (v NullableTamAdvisoryInstanceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTamAdvisoryInstanceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
