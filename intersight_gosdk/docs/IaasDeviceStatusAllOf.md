# IaasDeviceStatusAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountName** | Pointer to **string** | The UCSD infra account name. Account Name is created when UCSD admin adds any new infra account (Physical/Virtual/Compute/Network) to be managed by UCSD. | [optional] [readonly] 
**AccountType** | Pointer to **string** | The UCSD Infra Account type. | [optional] [readonly] 
**ClaimStatus** | Pointer to **string** | Describes if the device is claimed in Intersight or not. | [optional] [readonly] [default to "Unknown"]
**ConnectionStatus** | Pointer to **string** | Describes about the connection status between the UCSD and the actual end device. | [optional] [readonly] 
**DeviceModel** | Pointer to **string** | Describes about the device model. | [optional] [readonly] 
**DeviceVendor** | Pointer to **string** | Describes about the device vendor/manufacturer of the device. | [optional] [readonly] 
**DeviceVersion** | Pointer to **string** | Describes about the current firmware version running on the device. | [optional] [readonly] 
**IpAddress** | Pointer to **string** | The IPAddress of the device. | [optional] [readonly] 
**Pod** | Pointer to **string** | Describes about the pod to which this device belongs to in UCSD. | [optional] [readonly] 
**PodType** | Pointer to **string** | Describes about the podType of Pod to which this device belongs to in UCSD. | [optional] [readonly] 
**Guid** | Pointer to [**IaasUcsdInfoRelationship**](iaas.UcsdInfo.Relationship.md) |  | [optional] 

## Methods

### NewIaasDeviceStatusAllOf

`func NewIaasDeviceStatusAllOf() *IaasDeviceStatusAllOf`

NewIaasDeviceStatusAllOf instantiates a new IaasDeviceStatusAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIaasDeviceStatusAllOfWithDefaults

`func NewIaasDeviceStatusAllOfWithDefaults() *IaasDeviceStatusAllOf`

NewIaasDeviceStatusAllOfWithDefaults instantiates a new IaasDeviceStatusAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountName

`func (o *IaasDeviceStatusAllOf) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *IaasDeviceStatusAllOf) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *IaasDeviceStatusAllOf) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *IaasDeviceStatusAllOf) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetAccountType

`func (o *IaasDeviceStatusAllOf) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *IaasDeviceStatusAllOf) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *IaasDeviceStatusAllOf) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *IaasDeviceStatusAllOf) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetClaimStatus

`func (o *IaasDeviceStatusAllOf) GetClaimStatus() string`

GetClaimStatus returns the ClaimStatus field if non-nil, zero value otherwise.

### GetClaimStatusOk

`func (o *IaasDeviceStatusAllOf) GetClaimStatusOk() (*string, bool)`

GetClaimStatusOk returns a tuple with the ClaimStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimStatus

`func (o *IaasDeviceStatusAllOf) SetClaimStatus(v string)`

SetClaimStatus sets ClaimStatus field to given value.

### HasClaimStatus

`func (o *IaasDeviceStatusAllOf) HasClaimStatus() bool`

HasClaimStatus returns a boolean if a field has been set.

### GetConnectionStatus

`func (o *IaasDeviceStatusAllOf) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *IaasDeviceStatusAllOf) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *IaasDeviceStatusAllOf) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *IaasDeviceStatusAllOf) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### GetDeviceModel

`func (o *IaasDeviceStatusAllOf) GetDeviceModel() string`

GetDeviceModel returns the DeviceModel field if non-nil, zero value otherwise.

### GetDeviceModelOk

`func (o *IaasDeviceStatusAllOf) GetDeviceModelOk() (*string, bool)`

GetDeviceModelOk returns a tuple with the DeviceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModel

`func (o *IaasDeviceStatusAllOf) SetDeviceModel(v string)`

SetDeviceModel sets DeviceModel field to given value.

### HasDeviceModel

`func (o *IaasDeviceStatusAllOf) HasDeviceModel() bool`

HasDeviceModel returns a boolean if a field has been set.

### GetDeviceVendor

`func (o *IaasDeviceStatusAllOf) GetDeviceVendor() string`

GetDeviceVendor returns the DeviceVendor field if non-nil, zero value otherwise.

### GetDeviceVendorOk

`func (o *IaasDeviceStatusAllOf) GetDeviceVendorOk() (*string, bool)`

GetDeviceVendorOk returns a tuple with the DeviceVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceVendor

`func (o *IaasDeviceStatusAllOf) SetDeviceVendor(v string)`

SetDeviceVendor sets DeviceVendor field to given value.

### HasDeviceVendor

`func (o *IaasDeviceStatusAllOf) HasDeviceVendor() bool`

HasDeviceVendor returns a boolean if a field has been set.

### GetDeviceVersion

`func (o *IaasDeviceStatusAllOf) GetDeviceVersion() string`

GetDeviceVersion returns the DeviceVersion field if non-nil, zero value otherwise.

### GetDeviceVersionOk

`func (o *IaasDeviceStatusAllOf) GetDeviceVersionOk() (*string, bool)`

GetDeviceVersionOk returns a tuple with the DeviceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceVersion

`func (o *IaasDeviceStatusAllOf) SetDeviceVersion(v string)`

SetDeviceVersion sets DeviceVersion field to given value.

### HasDeviceVersion

`func (o *IaasDeviceStatusAllOf) HasDeviceVersion() bool`

HasDeviceVersion returns a boolean if a field has been set.

### GetIpAddress

`func (o *IaasDeviceStatusAllOf) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *IaasDeviceStatusAllOf) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *IaasDeviceStatusAllOf) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *IaasDeviceStatusAllOf) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetPod

`func (o *IaasDeviceStatusAllOf) GetPod() string`

GetPod returns the Pod field if non-nil, zero value otherwise.

### GetPodOk

`func (o *IaasDeviceStatusAllOf) GetPodOk() (*string, bool)`

GetPodOk returns a tuple with the Pod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPod

`func (o *IaasDeviceStatusAllOf) SetPod(v string)`

SetPod sets Pod field to given value.

### HasPod

`func (o *IaasDeviceStatusAllOf) HasPod() bool`

HasPod returns a boolean if a field has been set.

### GetPodType

`func (o *IaasDeviceStatusAllOf) GetPodType() string`

GetPodType returns the PodType field if non-nil, zero value otherwise.

### GetPodTypeOk

`func (o *IaasDeviceStatusAllOf) GetPodTypeOk() (*string, bool)`

GetPodTypeOk returns a tuple with the PodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodType

`func (o *IaasDeviceStatusAllOf) SetPodType(v string)`

SetPodType sets PodType field to given value.

### HasPodType

`func (o *IaasDeviceStatusAllOf) HasPodType() bool`

HasPodType returns a boolean if a field has been set.

### GetGuid

`func (o *IaasDeviceStatusAllOf) GetGuid() IaasUcsdInfoRelationship`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *IaasDeviceStatusAllOf) GetGuidOk() (*IaasUcsdInfoRelationship, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *IaasDeviceStatusAllOf) SetGuid(v IaasUcsdInfoRelationship)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *IaasDeviceStatusAllOf) HasGuid() bool`

HasGuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


