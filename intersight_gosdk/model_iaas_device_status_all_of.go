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

// IaasDeviceStatusAllOf Definition of the list of properties defined in 'iaas.DeviceStatus', excluding properties defined in parent classes.
type IaasDeviceStatusAllOf struct {
	// The UCSD infra account name. Account Name is created when UCSD admin adds any new infra account (Physical/Virtual/Compute/Network) to be managed by UCSD.
	AccountName *string `json:"AccountName,omitempty"`
	// The UCSD Infra Account type.
	AccountType *string `json:"AccountType,omitempty"`
	// Describes if the device is claimed in Intersight or not.
	ClaimStatus *string `json:"ClaimStatus,omitempty"`
	// Describes about the connection status between the UCSD and the actual end device.
	ConnectionStatus *string `json:"ConnectionStatus,omitempty"`
	// Describes about the device model.
	DeviceModel *string `json:"DeviceModel,omitempty"`
	// Describes about the device vendor/manufacturer of the device.
	DeviceVendor *string `json:"DeviceVendor,omitempty"`
	// Describes about the current firmware version running on the device.
	DeviceVersion *string `json:"DeviceVersion,omitempty"`
	// The IPAddress of the device.
	IpAddress *string `json:"IpAddress,omitempty"`
	// Describes about the pod to which this device belongs to in UCSD.
	Pod *string `json:"Pod,omitempty"`
	// Describes about the podType of Pod to which this device belongs to in UCSD.
	PodType *string                   `json:"PodType,omitempty"`
	Guid    *IaasUcsdInfoRelationship `json:"Guid,omitempty"`
}

// NewIaasDeviceStatusAllOf instantiates a new IaasDeviceStatusAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIaasDeviceStatusAllOf() *IaasDeviceStatusAllOf {
	this := IaasDeviceStatusAllOf{}
	var claimStatus string = "Unknown"
	this.ClaimStatus = &claimStatus
	return &this
}

// NewIaasDeviceStatusAllOfWithDefaults instantiates a new IaasDeviceStatusAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIaasDeviceStatusAllOfWithDefaults() *IaasDeviceStatusAllOf {
	this := IaasDeviceStatusAllOf{}
	var claimStatus string = "Unknown"
	this.ClaimStatus = &claimStatus
	return &this
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *IaasDeviceStatusAllOf) GetAccountName() string {
	if o == nil || o.AccountName == nil {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasDeviceStatusAllOf) GetAccountNameOk() (*string, bool) {
	if o == nil || o.AccountName == nil {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *IaasDeviceStatusAllOf) HasAccountName() bool {
	if o != nil && o.AccountName != nil {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *IaasDeviceStatusAllOf) SetAccountName(v string) {
	o.AccountName = &v
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *IaasDeviceStatusAllOf) GetAccountType() string {
	if o == nil || o.AccountType == nil {
		var ret string
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasDeviceStatusAllOf) GetAccountTypeOk() (*string, bool) {
	if o == nil || o.AccountType == nil {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *IaasDeviceStatusAllOf) HasAccountType() bool {
	if o != nil && o.AccountType != nil {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given string and assigns it to the AccountType field.
func (o *IaasDeviceStatusAllOf) SetAccountType(v string) {
	o.AccountType = &v
}

// GetClaimStatus returns the ClaimStatus field value if set, zero value otherwise.
func (o *IaasDeviceStatusAllOf) GetClaimStatus() string {
	if o == nil || o.ClaimStatus == nil {
		var ret string
		return ret
	}
	return *o.ClaimStatus
}

// GetClaimStatusOk returns a tuple with the ClaimStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasDeviceStatusAllOf) GetClaimStatusOk() (*string, bool) {
	if o == nil || o.ClaimStatus == nil {
		return nil, false
	}
	return o.ClaimStatus, true
}

// HasClaimStatus returns a boolean if a field has been set.
func (o *IaasDeviceStatusAllOf) HasClaimStatus() bool {
	if o != nil && o.ClaimStatus != nil {
		return true
	}

	return false
}

// SetClaimStatus gets a reference to the given string and assigns it to the ClaimStatus field.
func (o *IaasDeviceStatusAllOf) SetClaimStatus(v string) {
	o.ClaimStatus = &v
}

// GetConnectionStatus returns the ConnectionStatus field value if set, zero value otherwise.
func (o *IaasDeviceStatusAllOf) GetConnectionStatus() string {
	if o == nil || o.ConnectionStatus == nil {
		var ret string
		return ret
	}
	return *o.ConnectionStatus
}

// GetConnectionStatusOk returns a tuple with the ConnectionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasDeviceStatusAllOf) GetConnectionStatusOk() (*string, bool) {
	if o == nil || o.ConnectionStatus == nil {
		return nil, false
	}
	return o.ConnectionStatus, true
}

// HasConnectionStatus returns a boolean if a field has been set.
func (o *IaasDeviceStatusAllOf) HasConnectionStatus() bool {
	if o != nil && o.ConnectionStatus != nil {
		return true
	}

	return false
}

// SetConnectionStatus gets a reference to the given string and assigns it to the ConnectionStatus field.
func (o *IaasDeviceStatusAllOf) SetConnectionStatus(v string) {
	o.ConnectionStatus = &v
}

// GetDeviceModel returns the DeviceModel field value if set, zero value otherwise.
func (o *IaasDeviceStatusAllOf) GetDeviceModel() string {
	if o == nil || o.DeviceModel == nil {
		var ret string
		return ret
	}
	return *o.DeviceModel
}

// GetDeviceModelOk returns a tuple with the DeviceModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasDeviceStatusAllOf) GetDeviceModelOk() (*string, bool) {
	if o == nil || o.DeviceModel == nil {
		return nil, false
	}
	return o.DeviceModel, true
}

// HasDeviceModel returns a boolean if a field has been set.
func (o *IaasDeviceStatusAllOf) HasDeviceModel() bool {
	if o != nil && o.DeviceModel != nil {
		return true
	}

	return false
}

// SetDeviceModel gets a reference to the given string and assigns it to the DeviceModel field.
func (o *IaasDeviceStatusAllOf) SetDeviceModel(v string) {
	o.DeviceModel = &v
}

// GetDeviceVendor returns the DeviceVendor field value if set, zero value otherwise.
func (o *IaasDeviceStatusAllOf) GetDeviceVendor() string {
	if o == nil || o.DeviceVendor == nil {
		var ret string
		return ret
	}
	return *o.DeviceVendor
}

// GetDeviceVendorOk returns a tuple with the DeviceVendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasDeviceStatusAllOf) GetDeviceVendorOk() (*string, bool) {
	if o == nil || o.DeviceVendor == nil {
		return nil, false
	}
	return o.DeviceVendor, true
}

// HasDeviceVendor returns a boolean if a field has been set.
func (o *IaasDeviceStatusAllOf) HasDeviceVendor() bool {
	if o != nil && o.DeviceVendor != nil {
		return true
	}

	return false
}

// SetDeviceVendor gets a reference to the given string and assigns it to the DeviceVendor field.
func (o *IaasDeviceStatusAllOf) SetDeviceVendor(v string) {
	o.DeviceVendor = &v
}

// GetDeviceVersion returns the DeviceVersion field value if set, zero value otherwise.
func (o *IaasDeviceStatusAllOf) GetDeviceVersion() string {
	if o == nil || o.DeviceVersion == nil {
		var ret string
		return ret
	}
	return *o.DeviceVersion
}

// GetDeviceVersionOk returns a tuple with the DeviceVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasDeviceStatusAllOf) GetDeviceVersionOk() (*string, bool) {
	if o == nil || o.DeviceVersion == nil {
		return nil, false
	}
	return o.DeviceVersion, true
}

// HasDeviceVersion returns a boolean if a field has been set.
func (o *IaasDeviceStatusAllOf) HasDeviceVersion() bool {
	if o != nil && o.DeviceVersion != nil {
		return true
	}

	return false
}

// SetDeviceVersion gets a reference to the given string and assigns it to the DeviceVersion field.
func (o *IaasDeviceStatusAllOf) SetDeviceVersion(v string) {
	o.DeviceVersion = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *IaasDeviceStatusAllOf) GetIpAddress() string {
	if o == nil || o.IpAddress == nil {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasDeviceStatusAllOf) GetIpAddressOk() (*string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *IaasDeviceStatusAllOf) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *IaasDeviceStatusAllOf) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetPod returns the Pod field value if set, zero value otherwise.
func (o *IaasDeviceStatusAllOf) GetPod() string {
	if o == nil || o.Pod == nil {
		var ret string
		return ret
	}
	return *o.Pod
}

// GetPodOk returns a tuple with the Pod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasDeviceStatusAllOf) GetPodOk() (*string, bool) {
	if o == nil || o.Pod == nil {
		return nil, false
	}
	return o.Pod, true
}

// HasPod returns a boolean if a field has been set.
func (o *IaasDeviceStatusAllOf) HasPod() bool {
	if o != nil && o.Pod != nil {
		return true
	}

	return false
}

// SetPod gets a reference to the given string and assigns it to the Pod field.
func (o *IaasDeviceStatusAllOf) SetPod(v string) {
	o.Pod = &v
}

// GetPodType returns the PodType field value if set, zero value otherwise.
func (o *IaasDeviceStatusAllOf) GetPodType() string {
	if o == nil || o.PodType == nil {
		var ret string
		return ret
	}
	return *o.PodType
}

// GetPodTypeOk returns a tuple with the PodType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasDeviceStatusAllOf) GetPodTypeOk() (*string, bool) {
	if o == nil || o.PodType == nil {
		return nil, false
	}
	return o.PodType, true
}

// HasPodType returns a boolean if a field has been set.
func (o *IaasDeviceStatusAllOf) HasPodType() bool {
	if o != nil && o.PodType != nil {
		return true
	}

	return false
}

// SetPodType gets a reference to the given string and assigns it to the PodType field.
func (o *IaasDeviceStatusAllOf) SetPodType(v string) {
	o.PodType = &v
}

// GetGuid returns the Guid field value if set, zero value otherwise.
func (o *IaasDeviceStatusAllOf) GetGuid() IaasUcsdInfoRelationship {
	if o == nil || o.Guid == nil {
		var ret IaasUcsdInfoRelationship
		return ret
	}
	return *o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasDeviceStatusAllOf) GetGuidOk() (*IaasUcsdInfoRelationship, bool) {
	if o == nil || o.Guid == nil {
		return nil, false
	}
	return o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *IaasDeviceStatusAllOf) HasGuid() bool {
	if o != nil && o.Guid != nil {
		return true
	}

	return false
}

// SetGuid gets a reference to the given IaasUcsdInfoRelationship and assigns it to the Guid field.
func (o *IaasDeviceStatusAllOf) SetGuid(v IaasUcsdInfoRelationship) {
	o.Guid = &v
}

func (o IaasDeviceStatusAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountName != nil {
		toSerialize["AccountName"] = o.AccountName
	}
	if o.AccountType != nil {
		toSerialize["AccountType"] = o.AccountType
	}
	if o.ClaimStatus != nil {
		toSerialize["ClaimStatus"] = o.ClaimStatus
	}
	if o.ConnectionStatus != nil {
		toSerialize["ConnectionStatus"] = o.ConnectionStatus
	}
	if o.DeviceModel != nil {
		toSerialize["DeviceModel"] = o.DeviceModel
	}
	if o.DeviceVendor != nil {
		toSerialize["DeviceVendor"] = o.DeviceVendor
	}
	if o.DeviceVersion != nil {
		toSerialize["DeviceVersion"] = o.DeviceVersion
	}
	if o.IpAddress != nil {
		toSerialize["IpAddress"] = o.IpAddress
	}
	if o.Pod != nil {
		toSerialize["Pod"] = o.Pod
	}
	if o.PodType != nil {
		toSerialize["PodType"] = o.PodType
	}
	if o.Guid != nil {
		toSerialize["Guid"] = o.Guid
	}
	return json.Marshal(toSerialize)
}

type NullableIaasDeviceStatusAllOf struct {
	value *IaasDeviceStatusAllOf
	isSet bool
}

func (v NullableIaasDeviceStatusAllOf) Get() *IaasDeviceStatusAllOf {
	return v.value
}

func (v *NullableIaasDeviceStatusAllOf) Set(val *IaasDeviceStatusAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIaasDeviceStatusAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIaasDeviceStatusAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIaasDeviceStatusAllOf(val *IaasDeviceStatusAllOf) *NullableIaasDeviceStatusAllOf {
	return &NullableIaasDeviceStatusAllOf{value: val, isSet: true}
}

func (v NullableIaasDeviceStatusAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIaasDeviceStatusAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
