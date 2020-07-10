/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-06-19T15:15:17Z.
 *
 * API version: 1.0.9-1903
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
)

// ApplianceDeviceClaim DeviceClaim managed object represents a user initiated claim request for claiming an endpoint device. There can be many DeviceClaim managed object for a given endpoint device when users claim and unclaim devices repeatedly. Claiming an endpoint device is a multi-step operation. The Intersight Appliance starts a workflow with multiple tasks to process the device claim request. The status of the device claim operation can be obtained from the claim workflow.
type ApplianceDeviceClaim struct {
	MoBaseMo
	// Device identifier of the endpoint device.
	DeviceId *string `json:"DeviceId,omitempty"`
	// Hostname or IP address of the endpoint device the user wants to claim.
	Hostname *string `json:"Hostname,omitempty"`
	// Indicates whether the value of the 'password' property has been set.
	IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
	// Message set by the device claim process.
	Message *string `json:"Message,omitempty"`
	// Password to be used to login to the endpoint device.
	Password *string `json:"Password,omitempty"`
	// Platform type of the endpoint device.
	PlatformType *string `json:"PlatformType,omitempty"`
	// User defined claim request identifier set by the UI. The RequestId field is not a mandatory. The Intersight Appliance will assign a unique value automatically if the field is not set.
	RequestId *string `json:"RequestId,omitempty"`
	// Device security token of the endpoint device.
	SecurityToken *string `json:"SecurityToken,omitempty"`
	// Status of the device claim process.
	Status *string `json:"Status,omitempty"`
	// Username to log in to the endpoint device.
	Username *string                 `json:"Username,omitempty"`
	Account  *IamAccountRelationship `json:"Account,omitempty"`
}

// NewApplianceDeviceClaim instantiates a new ApplianceDeviceClaim object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceDeviceClaim() *ApplianceDeviceClaim {
	this := ApplianceDeviceClaim{}
	var platformType string = ""
	this.PlatformType = &platformType
	var status string = "started"
	this.Status = &status
	return &this
}

// NewApplianceDeviceClaimWithDefaults instantiates a new ApplianceDeviceClaim object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceDeviceClaimWithDefaults() *ApplianceDeviceClaim {
	this := ApplianceDeviceClaim{}
	var platformType string = ""
	this.PlatformType = &platformType
	var status string = "started"
	this.Status = &status
	return &this
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *ApplianceDeviceClaim) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceClaim) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *ApplianceDeviceClaim) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *ApplianceDeviceClaim) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *ApplianceDeviceClaim) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceClaim) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *ApplianceDeviceClaim) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *ApplianceDeviceClaim) SetHostname(v string) {
	o.Hostname = &v
}

// GetIsPasswordSet returns the IsPasswordSet field value if set, zero value otherwise.
func (o *ApplianceDeviceClaim) GetIsPasswordSet() bool {
	if o == nil || o.IsPasswordSet == nil {
		var ret bool
		return ret
	}
	return *o.IsPasswordSet
}

// GetIsPasswordSetOk returns a tuple with the IsPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceClaim) GetIsPasswordSetOk() (*bool, bool) {
	if o == nil || o.IsPasswordSet == nil {
		return nil, false
	}
	return o.IsPasswordSet, true
}

// HasIsPasswordSet returns a boolean if a field has been set.
func (o *ApplianceDeviceClaim) HasIsPasswordSet() bool {
	if o != nil && o.IsPasswordSet != nil {
		return true
	}

	return false
}

// SetIsPasswordSet gets a reference to the given bool and assigns it to the IsPasswordSet field.
func (o *ApplianceDeviceClaim) SetIsPasswordSet(v bool) {
	o.IsPasswordSet = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ApplianceDeviceClaim) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceClaim) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ApplianceDeviceClaim) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ApplianceDeviceClaim) SetMessage(v string) {
	o.Message = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ApplianceDeviceClaim) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceClaim) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ApplianceDeviceClaim) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *ApplianceDeviceClaim) SetPassword(v string) {
	o.Password = &v
}

// GetPlatformType returns the PlatformType field value if set, zero value otherwise.
func (o *ApplianceDeviceClaim) GetPlatformType() string {
	if o == nil || o.PlatformType == nil {
		var ret string
		return ret
	}
	return *o.PlatformType
}

// GetPlatformTypeOk returns a tuple with the PlatformType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceClaim) GetPlatformTypeOk() (*string, bool) {
	if o == nil || o.PlatformType == nil {
		return nil, false
	}
	return o.PlatformType, true
}

// HasPlatformType returns a boolean if a field has been set.
func (o *ApplianceDeviceClaim) HasPlatformType() bool {
	if o != nil && o.PlatformType != nil {
		return true
	}

	return false
}

// SetPlatformType gets a reference to the given string and assigns it to the PlatformType field.
func (o *ApplianceDeviceClaim) SetPlatformType(v string) {
	o.PlatformType = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *ApplianceDeviceClaim) GetRequestId() string {
	if o == nil || o.RequestId == nil {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceClaim) GetRequestIdOk() (*string, bool) {
	if o == nil || o.RequestId == nil {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *ApplianceDeviceClaim) HasRequestId() bool {
	if o != nil && o.RequestId != nil {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *ApplianceDeviceClaim) SetRequestId(v string) {
	o.RequestId = &v
}

// GetSecurityToken returns the SecurityToken field value if set, zero value otherwise.
func (o *ApplianceDeviceClaim) GetSecurityToken() string {
	if o == nil || o.SecurityToken == nil {
		var ret string
		return ret
	}
	return *o.SecurityToken
}

// GetSecurityTokenOk returns a tuple with the SecurityToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceClaim) GetSecurityTokenOk() (*string, bool) {
	if o == nil || o.SecurityToken == nil {
		return nil, false
	}
	return o.SecurityToken, true
}

// HasSecurityToken returns a boolean if a field has been set.
func (o *ApplianceDeviceClaim) HasSecurityToken() bool {
	if o != nil && o.SecurityToken != nil {
		return true
	}

	return false
}

// SetSecurityToken gets a reference to the given string and assigns it to the SecurityToken field.
func (o *ApplianceDeviceClaim) SetSecurityToken(v string) {
	o.SecurityToken = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApplianceDeviceClaim) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceClaim) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApplianceDeviceClaim) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ApplianceDeviceClaim) SetStatus(v string) {
	o.Status = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ApplianceDeviceClaim) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceClaim) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ApplianceDeviceClaim) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ApplianceDeviceClaim) SetUsername(v string) {
	o.Username = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *ApplianceDeviceClaim) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceClaim) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *ApplianceDeviceClaim) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *ApplianceDeviceClaim) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o ApplianceDeviceClaim) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.DeviceId != nil {
		toSerialize["DeviceId"] = o.DeviceId
	}
	if o.Hostname != nil {
		toSerialize["Hostname"] = o.Hostname
	}
	if o.IsPasswordSet != nil {
		toSerialize["IsPasswordSet"] = o.IsPasswordSet
	}
	if o.Message != nil {
		toSerialize["Message"] = o.Message
	}
	if o.Password != nil {
		toSerialize["Password"] = o.Password
	}
	if o.PlatformType != nil {
		toSerialize["PlatformType"] = o.PlatformType
	}
	if o.RequestId != nil {
		toSerialize["RequestId"] = o.RequestId
	}
	if o.SecurityToken != nil {
		toSerialize["SecurityToken"] = o.SecurityToken
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.Username != nil {
		toSerialize["Username"] = o.Username
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	return json.Marshal(toSerialize)
}

type NullableApplianceDeviceClaim struct {
	value *ApplianceDeviceClaim
	isSet bool
}

func (v NullableApplianceDeviceClaim) Get() *ApplianceDeviceClaim {
	return v.value
}

func (v *NullableApplianceDeviceClaim) Set(val *ApplianceDeviceClaim) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceDeviceClaim) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceDeviceClaim) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceDeviceClaim(val *ApplianceDeviceClaim) *NullableApplianceDeviceClaim {
	return &NullableApplianceDeviceClaim{value: val, isSet: true}
}

func (v NullableApplianceDeviceClaim) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceDeviceClaim) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}