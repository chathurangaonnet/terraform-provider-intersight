# AssetManagedDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credential** | Pointer to [**CommCredential**](comm.Credential.md) |  | [optional] 
**DeviceType** | Pointer to **string** | Type of the Device such as VMware, Pure Storage supported by Intersight Assist. | [optional] [default to ""]
**IgnoreCert** | Pointer to **bool** | Ignore Certificates with protocol https for connecting to the Managed Device. It is not used for other protocols. | [optional] 
**IsEnabled** | Pointer to **bool** | Device is Enabled/Disabled. | [optional] 
**ManagementAddress** | Pointer to **string** | Management address of the device. It can be IPv4, IPv6 or FQDN. | [optional] 
**Port** | Pointer to **int64** | Port to use for connecting to the Managed Device. Port is optional. If not specified, default port for protocol is used. | [optional] 
**Protocol** | Pointer to **string** | Protocol to use for connecting to the Managed Device. | [optional] [default to "https"]
**Status** | Pointer to [**AssetManagedDeviceStatus**](asset.ManagedDeviceStatus.md) |  | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 
**DeviceConnectorManager** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**WorkflowInfo** | Pointer to [**WorkflowWorkflowInfoRelationship**](workflow.WorkflowInfo.Relationship.md) |  | [optional] 

## Methods

### NewAssetManagedDevice

`func NewAssetManagedDevice() *AssetManagedDevice`

NewAssetManagedDevice instantiates a new AssetManagedDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetManagedDeviceWithDefaults

`func NewAssetManagedDeviceWithDefaults() *AssetManagedDevice`

NewAssetManagedDeviceWithDefaults instantiates a new AssetManagedDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredential

`func (o *AssetManagedDevice) GetCredential() CommCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *AssetManagedDevice) GetCredentialOk() (*CommCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *AssetManagedDevice) SetCredential(v CommCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *AssetManagedDevice) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetDeviceType

`func (o *AssetManagedDevice) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *AssetManagedDevice) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *AssetManagedDevice) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *AssetManagedDevice) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetIgnoreCert

`func (o *AssetManagedDevice) GetIgnoreCert() bool`

GetIgnoreCert returns the IgnoreCert field if non-nil, zero value otherwise.

### GetIgnoreCertOk

`func (o *AssetManagedDevice) GetIgnoreCertOk() (*bool, bool)`

GetIgnoreCertOk returns a tuple with the IgnoreCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreCert

`func (o *AssetManagedDevice) SetIgnoreCert(v bool)`

SetIgnoreCert sets IgnoreCert field to given value.

### HasIgnoreCert

`func (o *AssetManagedDevice) HasIgnoreCert() bool`

HasIgnoreCert returns a boolean if a field has been set.

### GetIsEnabled

`func (o *AssetManagedDevice) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *AssetManagedDevice) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *AssetManagedDevice) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *AssetManagedDevice) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetManagementAddress

`func (o *AssetManagedDevice) GetManagementAddress() string`

GetManagementAddress returns the ManagementAddress field if non-nil, zero value otherwise.

### GetManagementAddressOk

`func (o *AssetManagedDevice) GetManagementAddressOk() (*string, bool)`

GetManagementAddressOk returns a tuple with the ManagementAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementAddress

`func (o *AssetManagedDevice) SetManagementAddress(v string)`

SetManagementAddress sets ManagementAddress field to given value.

### HasManagementAddress

`func (o *AssetManagedDevice) HasManagementAddress() bool`

HasManagementAddress returns a boolean if a field has been set.

### GetPort

`func (o *AssetManagedDevice) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *AssetManagedDevice) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *AssetManagedDevice) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *AssetManagedDevice) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *AssetManagedDevice) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *AssetManagedDevice) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *AssetManagedDevice) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *AssetManagedDevice) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetStatus

`func (o *AssetManagedDevice) GetStatus() AssetManagedDeviceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AssetManagedDevice) GetStatusOk() (*AssetManagedDeviceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AssetManagedDevice) SetStatus(v AssetManagedDeviceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AssetManagedDevice) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAccount

`func (o *AssetManagedDevice) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AssetManagedDevice) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AssetManagedDevice) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AssetManagedDevice) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetDeviceConnectorManager

`func (o *AssetManagedDevice) GetDeviceConnectorManager() AssetDeviceRegistrationRelationship`

GetDeviceConnectorManager returns the DeviceConnectorManager field if non-nil, zero value otherwise.

### GetDeviceConnectorManagerOk

`func (o *AssetManagedDevice) GetDeviceConnectorManagerOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceConnectorManagerOk returns a tuple with the DeviceConnectorManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceConnectorManager

`func (o *AssetManagedDevice) SetDeviceConnectorManager(v AssetDeviceRegistrationRelationship)`

SetDeviceConnectorManager sets DeviceConnectorManager field to given value.

### HasDeviceConnectorManager

`func (o *AssetManagedDevice) HasDeviceConnectorManager() bool`

HasDeviceConnectorManager returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *AssetManagedDevice) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *AssetManagedDevice) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *AssetManagedDevice) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *AssetManagedDevice) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetWorkflowInfo

`func (o *AssetManagedDevice) GetWorkflowInfo() WorkflowWorkflowInfoRelationship`

GetWorkflowInfo returns the WorkflowInfo field if non-nil, zero value otherwise.

### GetWorkflowInfoOk

`func (o *AssetManagedDevice) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetWorkflowInfoOk returns a tuple with the WorkflowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowInfo

`func (o *AssetManagedDevice) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship)`

SetWorkflowInfo sets WorkflowInfo field to given value.

### HasWorkflowInfo

`func (o *AssetManagedDevice) HasWorkflowInfo() bool`

HasWorkflowInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


