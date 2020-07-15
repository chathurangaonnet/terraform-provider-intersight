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

// AssetManagedDeviceStatus Maintains the Managed Device Status.
type AssetManagedDeviceStatus struct {
	MoBaseComplexType
	// Port used for the connection to the Cloud by the Device Connector for the Managed Device.
	CloudPort *int64 `json:"CloudPort,omitempty"`
	// Maintains the reason for the failure of connection to the Device in case of connection failure.
	ConnectionFailureReason *string `json:"ConnectionFailureReason,omitempty"`
	// Maintains the status of the connection to the Device.
	ConnectionStatus *string `json:"ConnectionStatus,omitempty"`
	// Maintains code related to error from Device Connector, if any.
	ErrorCode *int64 `json:"ErrorCode,omitempty"`
	// Maintains the reason for the error from Device Connector, if any.
	ErrorReason *string `json:"ErrorReason,omitempty"`
	// Maintains the process pid of the Device Connector for the Managed Device.
	ProcessId *int64 `json:"ProcessId,omitempty"`
	// Port used for receiving requests from Intersight Assist by the Device Connector for the Managed Device.
	ServerPort *int64 `json:"ServerPort,omitempty"`
	// Maintains the state of the Managed Device, such as Start Success, Start Failure, etc. See ManagedDeviceState for device connection states.
	State                *string `json:"State,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetManagedDeviceStatus AssetManagedDeviceStatus

// NewAssetManagedDeviceStatus instantiates a new AssetManagedDeviceStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetManagedDeviceStatus() *AssetManagedDeviceStatus {
	this := AssetManagedDeviceStatus{}
	var connectionStatus string = "Unknown"
	this.ConnectionStatus = &connectionStatus
	var state string = "New"
	this.State = &state
	return &this
}

// NewAssetManagedDeviceStatusWithDefaults instantiates a new AssetManagedDeviceStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetManagedDeviceStatusWithDefaults() *AssetManagedDeviceStatus {
	this := AssetManagedDeviceStatus{}
	var connectionStatus string = "Unknown"
	this.ConnectionStatus = &connectionStatus
	var state string = "New"
	this.State = &state
	return &this
}

// GetCloudPort returns the CloudPort field value if set, zero value otherwise.
func (o *AssetManagedDeviceStatus) GetCloudPort() int64 {
	if o == nil || o.CloudPort == nil {
		var ret int64
		return ret
	}
	return *o.CloudPort
}

// GetCloudPortOk returns a tuple with the CloudPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetManagedDeviceStatus) GetCloudPortOk() (*int64, bool) {
	if o == nil || o.CloudPort == nil {
		return nil, false
	}
	return o.CloudPort, true
}

// HasCloudPort returns a boolean if a field has been set.
func (o *AssetManagedDeviceStatus) HasCloudPort() bool {
	if o != nil && o.CloudPort != nil {
		return true
	}

	return false
}

// SetCloudPort gets a reference to the given int64 and assigns it to the CloudPort field.
func (o *AssetManagedDeviceStatus) SetCloudPort(v int64) {
	o.CloudPort = &v
}

// GetConnectionFailureReason returns the ConnectionFailureReason field value if set, zero value otherwise.
func (o *AssetManagedDeviceStatus) GetConnectionFailureReason() string {
	if o == nil || o.ConnectionFailureReason == nil {
		var ret string
		return ret
	}
	return *o.ConnectionFailureReason
}

// GetConnectionFailureReasonOk returns a tuple with the ConnectionFailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetManagedDeviceStatus) GetConnectionFailureReasonOk() (*string, bool) {
	if o == nil || o.ConnectionFailureReason == nil {
		return nil, false
	}
	return o.ConnectionFailureReason, true
}

// HasConnectionFailureReason returns a boolean if a field has been set.
func (o *AssetManagedDeviceStatus) HasConnectionFailureReason() bool {
	if o != nil && o.ConnectionFailureReason != nil {
		return true
	}

	return false
}

// SetConnectionFailureReason gets a reference to the given string and assigns it to the ConnectionFailureReason field.
func (o *AssetManagedDeviceStatus) SetConnectionFailureReason(v string) {
	o.ConnectionFailureReason = &v
}

// GetConnectionStatus returns the ConnectionStatus field value if set, zero value otherwise.
func (o *AssetManagedDeviceStatus) GetConnectionStatus() string {
	if o == nil || o.ConnectionStatus == nil {
		var ret string
		return ret
	}
	return *o.ConnectionStatus
}

// GetConnectionStatusOk returns a tuple with the ConnectionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetManagedDeviceStatus) GetConnectionStatusOk() (*string, bool) {
	if o == nil || o.ConnectionStatus == nil {
		return nil, false
	}
	return o.ConnectionStatus, true
}

// HasConnectionStatus returns a boolean if a field has been set.
func (o *AssetManagedDeviceStatus) HasConnectionStatus() bool {
	if o != nil && o.ConnectionStatus != nil {
		return true
	}

	return false
}

// SetConnectionStatus gets a reference to the given string and assigns it to the ConnectionStatus field.
func (o *AssetManagedDeviceStatus) SetConnectionStatus(v string) {
	o.ConnectionStatus = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *AssetManagedDeviceStatus) GetErrorCode() int64 {
	if o == nil || o.ErrorCode == nil {
		var ret int64
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetManagedDeviceStatus) GetErrorCodeOk() (*int64, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *AssetManagedDeviceStatus) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given int64 and assigns it to the ErrorCode field.
func (o *AssetManagedDeviceStatus) SetErrorCode(v int64) {
	o.ErrorCode = &v
}

// GetErrorReason returns the ErrorReason field value if set, zero value otherwise.
func (o *AssetManagedDeviceStatus) GetErrorReason() string {
	if o == nil || o.ErrorReason == nil {
		var ret string
		return ret
	}
	return *o.ErrorReason
}

// GetErrorReasonOk returns a tuple with the ErrorReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetManagedDeviceStatus) GetErrorReasonOk() (*string, bool) {
	if o == nil || o.ErrorReason == nil {
		return nil, false
	}
	return o.ErrorReason, true
}

// HasErrorReason returns a boolean if a field has been set.
func (o *AssetManagedDeviceStatus) HasErrorReason() bool {
	if o != nil && o.ErrorReason != nil {
		return true
	}

	return false
}

// SetErrorReason gets a reference to the given string and assigns it to the ErrorReason field.
func (o *AssetManagedDeviceStatus) SetErrorReason(v string) {
	o.ErrorReason = &v
}

// GetProcessId returns the ProcessId field value if set, zero value otherwise.
func (o *AssetManagedDeviceStatus) GetProcessId() int64 {
	if o == nil || o.ProcessId == nil {
		var ret int64
		return ret
	}
	return *o.ProcessId
}

// GetProcessIdOk returns a tuple with the ProcessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetManagedDeviceStatus) GetProcessIdOk() (*int64, bool) {
	if o == nil || o.ProcessId == nil {
		return nil, false
	}
	return o.ProcessId, true
}

// HasProcessId returns a boolean if a field has been set.
func (o *AssetManagedDeviceStatus) HasProcessId() bool {
	if o != nil && o.ProcessId != nil {
		return true
	}

	return false
}

// SetProcessId gets a reference to the given int64 and assigns it to the ProcessId field.
func (o *AssetManagedDeviceStatus) SetProcessId(v int64) {
	o.ProcessId = &v
}

// GetServerPort returns the ServerPort field value if set, zero value otherwise.
func (o *AssetManagedDeviceStatus) GetServerPort() int64 {
	if o == nil || o.ServerPort == nil {
		var ret int64
		return ret
	}
	return *o.ServerPort
}

// GetServerPortOk returns a tuple with the ServerPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetManagedDeviceStatus) GetServerPortOk() (*int64, bool) {
	if o == nil || o.ServerPort == nil {
		return nil, false
	}
	return o.ServerPort, true
}

// HasServerPort returns a boolean if a field has been set.
func (o *AssetManagedDeviceStatus) HasServerPort() bool {
	if o != nil && o.ServerPort != nil {
		return true
	}

	return false
}

// SetServerPort gets a reference to the given int64 and assigns it to the ServerPort field.
func (o *AssetManagedDeviceStatus) SetServerPort(v int64) {
	o.ServerPort = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *AssetManagedDeviceStatus) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetManagedDeviceStatus) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *AssetManagedDeviceStatus) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *AssetManagedDeviceStatus) SetState(v string) {
	o.State = &v
}

func (o AssetManagedDeviceStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.CloudPort != nil {
		toSerialize["CloudPort"] = o.CloudPort
	}
	if o.ConnectionFailureReason != nil {
		toSerialize["ConnectionFailureReason"] = o.ConnectionFailureReason
	}
	if o.ConnectionStatus != nil {
		toSerialize["ConnectionStatus"] = o.ConnectionStatus
	}
	if o.ErrorCode != nil {
		toSerialize["ErrorCode"] = o.ErrorCode
	}
	if o.ErrorReason != nil {
		toSerialize["ErrorReason"] = o.ErrorReason
	}
	if o.ProcessId != nil {
		toSerialize["ProcessId"] = o.ProcessId
	}
	if o.ServerPort != nil {
		toSerialize["ServerPort"] = o.ServerPort
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetManagedDeviceStatus) UnmarshalJSON(bytes []byte) (err error) {
	type AssetManagedDeviceStatusWithoutEmbeddedStruct struct {
		// Port used for the connection to the Cloud by the Device Connector for the Managed Device.
		CloudPort *int64 `json:"CloudPort,omitempty"`
		// Maintains the reason for the failure of connection to the Device in case of connection failure.
		ConnectionFailureReason *string `json:"ConnectionFailureReason,omitempty"`
		// Maintains the status of the connection to the Device.
		ConnectionStatus *string `json:"ConnectionStatus,omitempty"`
		// Maintains code related to error from Device Connector, if any.
		ErrorCode *int64 `json:"ErrorCode,omitempty"`
		// Maintains the reason for the error from Device Connector, if any.
		ErrorReason *string `json:"ErrorReason,omitempty"`
		// Maintains the process pid of the Device Connector for the Managed Device.
		ProcessId *int64 `json:"ProcessId,omitempty"`
		// Port used for receiving requests from Intersight Assist by the Device Connector for the Managed Device.
		ServerPort *int64 `json:"ServerPort,omitempty"`
		// Maintains the state of the Managed Device, such as Start Success, Start Failure, etc. See ManagedDeviceState for device connection states.
		State *string `json:"State,omitempty"`
	}

	varAssetManagedDeviceStatusWithoutEmbeddedStruct := AssetManagedDeviceStatusWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAssetManagedDeviceStatusWithoutEmbeddedStruct)
	if err == nil {
		varAssetManagedDeviceStatus := _AssetManagedDeviceStatus{}
		varAssetManagedDeviceStatus.CloudPort = varAssetManagedDeviceStatusWithoutEmbeddedStruct.CloudPort
		varAssetManagedDeviceStatus.ConnectionFailureReason = varAssetManagedDeviceStatusWithoutEmbeddedStruct.ConnectionFailureReason
		varAssetManagedDeviceStatus.ConnectionStatus = varAssetManagedDeviceStatusWithoutEmbeddedStruct.ConnectionStatus
		varAssetManagedDeviceStatus.ErrorCode = varAssetManagedDeviceStatusWithoutEmbeddedStruct.ErrorCode
		varAssetManagedDeviceStatus.ErrorReason = varAssetManagedDeviceStatusWithoutEmbeddedStruct.ErrorReason
		varAssetManagedDeviceStatus.ProcessId = varAssetManagedDeviceStatusWithoutEmbeddedStruct.ProcessId
		varAssetManagedDeviceStatus.ServerPort = varAssetManagedDeviceStatusWithoutEmbeddedStruct.ServerPort
		varAssetManagedDeviceStatus.State = varAssetManagedDeviceStatusWithoutEmbeddedStruct.State
		*o = AssetManagedDeviceStatus(varAssetManagedDeviceStatus)
	} else {
		return err
	}

	varAssetManagedDeviceStatus := _AssetManagedDeviceStatus{}

	err = json.Unmarshal(bytes, &varAssetManagedDeviceStatus)
	if err == nil {
		o.MoBaseComplexType = varAssetManagedDeviceStatus.MoBaseComplexType
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "CloudPort")
		delete(additionalProperties, "ConnectionFailureReason")
		delete(additionalProperties, "ConnectionStatus")
		delete(additionalProperties, "ErrorCode")
		delete(additionalProperties, "ErrorReason")
		delete(additionalProperties, "ProcessId")
		delete(additionalProperties, "ServerPort")
		delete(additionalProperties, "State")

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

type NullableAssetManagedDeviceStatus struct {
	value *AssetManagedDeviceStatus
	isSet bool
}

func (v NullableAssetManagedDeviceStatus) Get() *AssetManagedDeviceStatus {
	return v.value
}

func (v *NullableAssetManagedDeviceStatus) Set(val *AssetManagedDeviceStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetManagedDeviceStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetManagedDeviceStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetManagedDeviceStatus(val *AssetManagedDeviceStatus) *NullableAssetManagedDeviceStatus {
	return &NullableAssetManagedDeviceStatus{value: val, isSet: true}
}

func (v NullableAssetManagedDeviceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetManagedDeviceStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
