# StorageBasePhysicalPortAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Iqn** | Pointer to **string** | ISCSI qualified name applicable for ethernet port. Example - &#39;iqn.2008-05.com.storage:fnm00151300643-514f0c50141faf05&#39;. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the physical port available in storage array. | [optional] [readonly] 
**Speed** | Pointer to **int64** | Operational speed of physical port measured in Gbps. | [optional] [readonly] 
**Status** | Pointer to **string** | Status of storage array port. | [optional] [readonly] [default to "Unknown"]
**Type** | Pointer to **string** | Port type - possible values are FC, FCoE and iSCSI. | [optional] [readonly] [default to "FC"]
**Wwn** | Pointer to **string** | World wide name of FC port. It is a combination of WWNN and WWPN represented in 128 bit hexadecimal formatted with colon. Example: &#39;51:4f:0c:50:14:1f:af:01:51:4f:0c:51:14:1f:af:01&#39;. | [optional] [readonly] 
**Wwnn** | Pointer to **string** | World wide node name of FC port. Represented in 64 bit hex digits, formatted with colon. Example - &#39;51:4f:0c:50:14:1f:af:01&#39;. | [optional] [readonly] 
**Wwpn** | Pointer to **string** | World wide port name of FC port. Represented in 64 bit hex digits, formatted with colon. Example - &#39;51:4f:0c:51:14:1f:af:01&#39;. | [optional] [readonly] 

## Methods

### NewStorageBasePhysicalPortAllOf

`func NewStorageBasePhysicalPortAllOf() *StorageBasePhysicalPortAllOf`

NewStorageBasePhysicalPortAllOf instantiates a new StorageBasePhysicalPortAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageBasePhysicalPortAllOfWithDefaults

`func NewStorageBasePhysicalPortAllOfWithDefaults() *StorageBasePhysicalPortAllOf`

NewStorageBasePhysicalPortAllOfWithDefaults instantiates a new StorageBasePhysicalPortAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIqn

`func (o *StorageBasePhysicalPortAllOf) GetIqn() string`

GetIqn returns the Iqn field if non-nil, zero value otherwise.

### GetIqnOk

`func (o *StorageBasePhysicalPortAllOf) GetIqnOk() (*string, bool)`

GetIqnOk returns a tuple with the Iqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIqn

`func (o *StorageBasePhysicalPortAllOf) SetIqn(v string)`

SetIqn sets Iqn field to given value.

### HasIqn

`func (o *StorageBasePhysicalPortAllOf) HasIqn() bool`

HasIqn returns a boolean if a field has been set.

### GetName

`func (o *StorageBasePhysicalPortAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageBasePhysicalPortAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageBasePhysicalPortAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageBasePhysicalPortAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSpeed

`func (o *StorageBasePhysicalPortAllOf) GetSpeed() int64`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *StorageBasePhysicalPortAllOf) GetSpeedOk() (*int64, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *StorageBasePhysicalPortAllOf) SetSpeed(v int64)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *StorageBasePhysicalPortAllOf) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetStatus

`func (o *StorageBasePhysicalPortAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StorageBasePhysicalPortAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StorageBasePhysicalPortAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StorageBasePhysicalPortAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *StorageBasePhysicalPortAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageBasePhysicalPortAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageBasePhysicalPortAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StorageBasePhysicalPortAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWwn

`func (o *StorageBasePhysicalPortAllOf) GetWwn() string`

GetWwn returns the Wwn field if non-nil, zero value otherwise.

### GetWwnOk

`func (o *StorageBasePhysicalPortAllOf) GetWwnOk() (*string, bool)`

GetWwnOk returns a tuple with the Wwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwn

`func (o *StorageBasePhysicalPortAllOf) SetWwn(v string)`

SetWwn sets Wwn field to given value.

### HasWwn

`func (o *StorageBasePhysicalPortAllOf) HasWwn() bool`

HasWwn returns a boolean if a field has been set.

### GetWwnn

`func (o *StorageBasePhysicalPortAllOf) GetWwnn() string`

GetWwnn returns the Wwnn field if non-nil, zero value otherwise.

### GetWwnnOk

`func (o *StorageBasePhysicalPortAllOf) GetWwnnOk() (*string, bool)`

GetWwnnOk returns a tuple with the Wwnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwnn

`func (o *StorageBasePhysicalPortAllOf) SetWwnn(v string)`

SetWwnn sets Wwnn field to given value.

### HasWwnn

`func (o *StorageBasePhysicalPortAllOf) HasWwnn() bool`

HasWwnn returns a boolean if a field has been set.

### GetWwpn

`func (o *StorageBasePhysicalPortAllOf) GetWwpn() string`

GetWwpn returns the Wwpn field if non-nil, zero value otherwise.

### GetWwpnOk

`func (o *StorageBasePhysicalPortAllOf) GetWwpnOk() (*string, bool)`

GetWwpnOk returns a tuple with the Wwpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwpn

`func (o *StorageBasePhysicalPortAllOf) SetWwpn(v string)`

SetWwpn sets Wwpn field to given value.

### HasWwpn

`func (o *StorageBasePhysicalPortAllOf) HasWwpn() bool`

HasWwpn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


