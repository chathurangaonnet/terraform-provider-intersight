# AssetManagedDeviceStatusAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudPort** | Pointer to **int64** | Port used for the connection to the Cloud by the Device Connector for the Managed Device. | [optional] 
**ConnectionFailureReason** | Pointer to **string** | Maintains the reason for the failure of connection to the Device in case of connection failure. | [optional] 
**ConnectionStatus** | Pointer to **string** | Maintains the status of the connection to the Device. | [optional] [default to "Unknown"]
**ErrorCode** | Pointer to **int64** | Maintains code related to error from Device Connector, if any. | [optional] 
**ErrorReason** | Pointer to **string** | Maintains the reason for the error from Device Connector, if any. | [optional] 
**ProcessId** | Pointer to **int64** | Maintains the process pid of the Device Connector for the Managed Device. | [optional] 
**ServerPort** | Pointer to **int64** | Port used for receiving requests from Intersight Assist by the Device Connector for the Managed Device. | [optional] 
**State** | Pointer to **string** | Maintains the state of the Managed Device, such as Start Success, Start Failure, etc. See ManagedDeviceState for device connection states. | [optional] [default to "New"]

## Methods

### NewAssetManagedDeviceStatusAllOf

`func NewAssetManagedDeviceStatusAllOf() *AssetManagedDeviceStatusAllOf`

NewAssetManagedDeviceStatusAllOf instantiates a new AssetManagedDeviceStatusAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetManagedDeviceStatusAllOfWithDefaults

`func NewAssetManagedDeviceStatusAllOfWithDefaults() *AssetManagedDeviceStatusAllOf`

NewAssetManagedDeviceStatusAllOfWithDefaults instantiates a new AssetManagedDeviceStatusAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudPort

`func (o *AssetManagedDeviceStatusAllOf) GetCloudPort() int64`

GetCloudPort returns the CloudPort field if non-nil, zero value otherwise.

### GetCloudPortOk

`func (o *AssetManagedDeviceStatusAllOf) GetCloudPortOk() (*int64, bool)`

GetCloudPortOk returns a tuple with the CloudPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudPort

`func (o *AssetManagedDeviceStatusAllOf) SetCloudPort(v int64)`

SetCloudPort sets CloudPort field to given value.

### HasCloudPort

`func (o *AssetManagedDeviceStatusAllOf) HasCloudPort() bool`

HasCloudPort returns a boolean if a field has been set.

### GetConnectionFailureReason

`func (o *AssetManagedDeviceStatusAllOf) GetConnectionFailureReason() string`

GetConnectionFailureReason returns the ConnectionFailureReason field if non-nil, zero value otherwise.

### GetConnectionFailureReasonOk

`func (o *AssetManagedDeviceStatusAllOf) GetConnectionFailureReasonOk() (*string, bool)`

GetConnectionFailureReasonOk returns a tuple with the ConnectionFailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionFailureReason

`func (o *AssetManagedDeviceStatusAllOf) SetConnectionFailureReason(v string)`

SetConnectionFailureReason sets ConnectionFailureReason field to given value.

### HasConnectionFailureReason

`func (o *AssetManagedDeviceStatusAllOf) HasConnectionFailureReason() bool`

HasConnectionFailureReason returns a boolean if a field has been set.

### GetConnectionStatus

`func (o *AssetManagedDeviceStatusAllOf) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *AssetManagedDeviceStatusAllOf) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *AssetManagedDeviceStatusAllOf) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *AssetManagedDeviceStatusAllOf) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### GetErrorCode

`func (o *AssetManagedDeviceStatusAllOf) GetErrorCode() int64`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *AssetManagedDeviceStatusAllOf) GetErrorCodeOk() (*int64, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *AssetManagedDeviceStatusAllOf) SetErrorCode(v int64)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *AssetManagedDeviceStatusAllOf) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorReason

`func (o *AssetManagedDeviceStatusAllOf) GetErrorReason() string`

GetErrorReason returns the ErrorReason field if non-nil, zero value otherwise.

### GetErrorReasonOk

`func (o *AssetManagedDeviceStatusAllOf) GetErrorReasonOk() (*string, bool)`

GetErrorReasonOk returns a tuple with the ErrorReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorReason

`func (o *AssetManagedDeviceStatusAllOf) SetErrorReason(v string)`

SetErrorReason sets ErrorReason field to given value.

### HasErrorReason

`func (o *AssetManagedDeviceStatusAllOf) HasErrorReason() bool`

HasErrorReason returns a boolean if a field has been set.

### GetProcessId

`func (o *AssetManagedDeviceStatusAllOf) GetProcessId() int64`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *AssetManagedDeviceStatusAllOf) GetProcessIdOk() (*int64, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *AssetManagedDeviceStatusAllOf) SetProcessId(v int64)`

SetProcessId sets ProcessId field to given value.

### HasProcessId

`func (o *AssetManagedDeviceStatusAllOf) HasProcessId() bool`

HasProcessId returns a boolean if a field has been set.

### GetServerPort

`func (o *AssetManagedDeviceStatusAllOf) GetServerPort() int64`

GetServerPort returns the ServerPort field if non-nil, zero value otherwise.

### GetServerPortOk

`func (o *AssetManagedDeviceStatusAllOf) GetServerPortOk() (*int64, bool)`

GetServerPortOk returns a tuple with the ServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPort

`func (o *AssetManagedDeviceStatusAllOf) SetServerPort(v int64)`

SetServerPort sets ServerPort field to given value.

### HasServerPort

`func (o *AssetManagedDeviceStatusAllOf) HasServerPort() bool`

HasServerPort returns a boolean if a field has been set.

### GetState

`func (o *AssetManagedDeviceStatusAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AssetManagedDeviceStatusAllOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AssetManagedDeviceStatusAllOf) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AssetManagedDeviceStatusAllOf) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


