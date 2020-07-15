/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-07-11T05:47:33Z.
 *
 * API version: 1.0.9-1999
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// InventoryEpInfo The runtime information about a UEM (Unified Event Mechanism) supported endpoint.
type InventoryEpInfo struct {
	MoBaseComplexType
	// Connections status of UEM endpoint.
	ConnectionStatus *string `json:"ConnectionStatus,omitempty"`
	// Type of UEM endpoint (BMC, CMC or VIC).
	EpType *string `json:"EpType,omitempty"`
	// The IP address of the UEM endpoint.
	Ip *string `json:"Ip,omitempty"`
	// The MAC address of the UEM endpoint.
	MacAddress *string `json:"MacAddress,omitempty"`
	// The member identity of the UEM connection being inventoried.
	MemberIdentity *string `json:"MemberIdentity,omitempty"`
	// The model of the UEM endpoint.
	Model *string `json:"Model,omitempty"`
	// The serial number of the UEM endpoint.
	Serial *string `json:"Serial,omitempty"`
	// The device id of the server which this EP is a part of.
	ServerRegistrationId *string `json:"ServerRegistrationId,omitempty"`
	// The vendor of the UEM endpoint.
	Vendor               *string `json:"Vendor,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InventoryEpInfo InventoryEpInfo

// NewInventoryEpInfo instantiates a new InventoryEpInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryEpInfo() *InventoryEpInfo {
	this := InventoryEpInfo{}
	var connectionStatus string = "Unknown"
	this.ConnectionStatus = &connectionStatus
	return &this
}

// NewInventoryEpInfoWithDefaults instantiates a new InventoryEpInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryEpInfoWithDefaults() *InventoryEpInfo {
	this := InventoryEpInfo{}
	var connectionStatus string = "Unknown"
	this.ConnectionStatus = &connectionStatus
	return &this
}

// GetConnectionStatus returns the ConnectionStatus field value if set, zero value otherwise.
func (o *InventoryEpInfo) GetConnectionStatus() string {
	if o == nil || o.ConnectionStatus == nil {
		var ret string
		return ret
	}
	return *o.ConnectionStatus
}

// GetConnectionStatusOk returns a tuple with the ConnectionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryEpInfo) GetConnectionStatusOk() (*string, bool) {
	if o == nil || o.ConnectionStatus == nil {
		return nil, false
	}
	return o.ConnectionStatus, true
}

// HasConnectionStatus returns a boolean if a field has been set.
func (o *InventoryEpInfo) HasConnectionStatus() bool {
	if o != nil && o.ConnectionStatus != nil {
		return true
	}

	return false
}

// SetConnectionStatus gets a reference to the given string and assigns it to the ConnectionStatus field.
func (o *InventoryEpInfo) SetConnectionStatus(v string) {
	o.ConnectionStatus = &v
}

// GetEpType returns the EpType field value if set, zero value otherwise.
func (o *InventoryEpInfo) GetEpType() string {
	if o == nil || o.EpType == nil {
		var ret string
		return ret
	}
	return *o.EpType
}

// GetEpTypeOk returns a tuple with the EpType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryEpInfo) GetEpTypeOk() (*string, bool) {
	if o == nil || o.EpType == nil {
		return nil, false
	}
	return o.EpType, true
}

// HasEpType returns a boolean if a field has been set.
func (o *InventoryEpInfo) HasEpType() bool {
	if o != nil && o.EpType != nil {
		return true
	}

	return false
}

// SetEpType gets a reference to the given string and assigns it to the EpType field.
func (o *InventoryEpInfo) SetEpType(v string) {
	o.EpType = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *InventoryEpInfo) GetIp() string {
	if o == nil || o.Ip == nil {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryEpInfo) GetIpOk() (*string, bool) {
	if o == nil || o.Ip == nil {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *InventoryEpInfo) HasIp() bool {
	if o != nil && o.Ip != nil {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *InventoryEpInfo) SetIp(v string) {
	o.Ip = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *InventoryEpInfo) GetMacAddress() string {
	if o == nil || o.MacAddress == nil {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryEpInfo) GetMacAddressOk() (*string, bool) {
	if o == nil || o.MacAddress == nil {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *InventoryEpInfo) HasMacAddress() bool {
	if o != nil && o.MacAddress != nil {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *InventoryEpInfo) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetMemberIdentity returns the MemberIdentity field value if set, zero value otherwise.
func (o *InventoryEpInfo) GetMemberIdentity() string {
	if o == nil || o.MemberIdentity == nil {
		var ret string
		return ret
	}
	return *o.MemberIdentity
}

// GetMemberIdentityOk returns a tuple with the MemberIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryEpInfo) GetMemberIdentityOk() (*string, bool) {
	if o == nil || o.MemberIdentity == nil {
		return nil, false
	}
	return o.MemberIdentity, true
}

// HasMemberIdentity returns a boolean if a field has been set.
func (o *InventoryEpInfo) HasMemberIdentity() bool {
	if o != nil && o.MemberIdentity != nil {
		return true
	}

	return false
}

// SetMemberIdentity gets a reference to the given string and assigns it to the MemberIdentity field.
func (o *InventoryEpInfo) SetMemberIdentity(v string) {
	o.MemberIdentity = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *InventoryEpInfo) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryEpInfo) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *InventoryEpInfo) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *InventoryEpInfo) SetModel(v string) {
	o.Model = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InventoryEpInfo) GetSerial() string {
	if o == nil || o.Serial == nil {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryEpInfo) GetSerialOk() (*string, bool) {
	if o == nil || o.Serial == nil {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InventoryEpInfo) HasSerial() bool {
	if o != nil && o.Serial != nil {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InventoryEpInfo) SetSerial(v string) {
	o.Serial = &v
}

// GetServerRegistrationId returns the ServerRegistrationId field value if set, zero value otherwise.
func (o *InventoryEpInfo) GetServerRegistrationId() string {
	if o == nil || o.ServerRegistrationId == nil {
		var ret string
		return ret
	}
	return *o.ServerRegistrationId
}

// GetServerRegistrationIdOk returns a tuple with the ServerRegistrationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryEpInfo) GetServerRegistrationIdOk() (*string, bool) {
	if o == nil || o.ServerRegistrationId == nil {
		return nil, false
	}
	return o.ServerRegistrationId, true
}

// HasServerRegistrationId returns a boolean if a field has been set.
func (o *InventoryEpInfo) HasServerRegistrationId() bool {
	if o != nil && o.ServerRegistrationId != nil {
		return true
	}

	return false
}

// SetServerRegistrationId gets a reference to the given string and assigns it to the ServerRegistrationId field.
func (o *InventoryEpInfo) SetServerRegistrationId(v string) {
	o.ServerRegistrationId = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *InventoryEpInfo) GetVendor() string {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryEpInfo) GetVendorOk() (*string, bool) {
	if o == nil || o.Vendor == nil {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *InventoryEpInfo) HasVendor() bool {
	if o != nil && o.Vendor != nil {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *InventoryEpInfo) SetVendor(v string) {
	o.Vendor = &v
}

func (o InventoryEpInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.ConnectionStatus != nil {
		toSerialize["ConnectionStatus"] = o.ConnectionStatus
	}
	if o.EpType != nil {
		toSerialize["EpType"] = o.EpType
	}
	if o.Ip != nil {
		toSerialize["Ip"] = o.Ip
	}
	if o.MacAddress != nil {
		toSerialize["MacAddress"] = o.MacAddress
	}
	if o.MemberIdentity != nil {
		toSerialize["MemberIdentity"] = o.MemberIdentity
	}
	if o.Model != nil {
		toSerialize["Model"] = o.Model
	}
	if o.Serial != nil {
		toSerialize["Serial"] = o.Serial
	}
	if o.ServerRegistrationId != nil {
		toSerialize["ServerRegistrationId"] = o.ServerRegistrationId
	}
	if o.Vendor != nil {
		toSerialize["Vendor"] = o.Vendor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InventoryEpInfo) UnmarshalJSON(bytes []byte) (err error) {
	type InventoryEpInfoWithoutEmbeddedStruct struct {
		// Connections status of UEM endpoint.
		ConnectionStatus *string `json:"ConnectionStatus,omitempty"`
		// Type of UEM endpoint (BMC, CMC or VIC).
		EpType *string `json:"EpType,omitempty"`
		// The IP address of the UEM endpoint.
		Ip *string `json:"Ip,omitempty"`
		// The MAC address of the UEM endpoint.
		MacAddress *string `json:"MacAddress,omitempty"`
		// The member identity of the UEM connection being inventoried.
		MemberIdentity *string `json:"MemberIdentity,omitempty"`
		// The model of the UEM endpoint.
		Model *string `json:"Model,omitempty"`
		// The serial number of the UEM endpoint.
		Serial *string `json:"Serial,omitempty"`
		// The device id of the server which this EP is a part of.
		ServerRegistrationId *string `json:"ServerRegistrationId,omitempty"`
		// The vendor of the UEM endpoint.
		Vendor *string `json:"Vendor,omitempty"`
	}

	varInventoryEpInfoWithoutEmbeddedStruct := InventoryEpInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varInventoryEpInfoWithoutEmbeddedStruct)
	if err == nil {
		varInventoryEpInfo := _InventoryEpInfo{}
		varInventoryEpInfo.ConnectionStatus = varInventoryEpInfoWithoutEmbeddedStruct.ConnectionStatus
		varInventoryEpInfo.EpType = varInventoryEpInfoWithoutEmbeddedStruct.EpType
		varInventoryEpInfo.Ip = varInventoryEpInfoWithoutEmbeddedStruct.Ip
		varInventoryEpInfo.MacAddress = varInventoryEpInfoWithoutEmbeddedStruct.MacAddress
		varInventoryEpInfo.MemberIdentity = varInventoryEpInfoWithoutEmbeddedStruct.MemberIdentity
		varInventoryEpInfo.Model = varInventoryEpInfoWithoutEmbeddedStruct.Model
		varInventoryEpInfo.Serial = varInventoryEpInfoWithoutEmbeddedStruct.Serial
		varInventoryEpInfo.ServerRegistrationId = varInventoryEpInfoWithoutEmbeddedStruct.ServerRegistrationId
		varInventoryEpInfo.Vendor = varInventoryEpInfoWithoutEmbeddedStruct.Vendor
		*o = InventoryEpInfo(varInventoryEpInfo)
	} else {
		return err
	}

	varInventoryEpInfo := _InventoryEpInfo{}

	err = json.Unmarshal(bytes, &varInventoryEpInfo)
	if err == nil {
		o.MoBaseComplexType = varInventoryEpInfo.MoBaseComplexType
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ConnectionStatus")
		delete(additionalProperties, "EpType")
		delete(additionalProperties, "Ip")
		delete(additionalProperties, "MacAddress")
		delete(additionalProperties, "MemberIdentity")
		delete(additionalProperties, "Model")
		delete(additionalProperties, "Serial")
		delete(additionalProperties, "ServerRegistrationId")
		delete(additionalProperties, "Vendor")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableInventoryEpInfo struct {
	value *InventoryEpInfo
	isSet bool
}

func (v NullableInventoryEpInfo) Get() *InventoryEpInfo {
	return v.value
}

func (v *NullableInventoryEpInfo) Set(val *InventoryEpInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryEpInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryEpInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryEpInfo(val *InventoryEpInfo) *NullableInventoryEpInfo {
	return &NullableInventoryEpInfo{value: val, isSet: true}
}

func (v NullableInventoryEpInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryEpInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
