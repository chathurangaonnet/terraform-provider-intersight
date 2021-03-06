# EquipmentIoCardIdentityAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IoCardMoid** | Pointer to **string** | MO Reference to equipmentIoCard MO in inventory service. | [optional] 
**ModuleId** | Pointer to **int64** | IOM/MUX Module ID connected to the FI. | [optional] 
**NetworkElementMoid** | Pointer to **string** | MO Reference to networkElement MO in inventory service. | [optional] 
**SwitchId** | Pointer to **int64** | Switch ID to which IOM is connected, ID can be either 1 or 2, equalent to A or B. | [optional] 

## Methods

### NewEquipmentIoCardIdentityAllOf

`func NewEquipmentIoCardIdentityAllOf() *EquipmentIoCardIdentityAllOf`

NewEquipmentIoCardIdentityAllOf instantiates a new EquipmentIoCardIdentityAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentIoCardIdentityAllOfWithDefaults

`func NewEquipmentIoCardIdentityAllOfWithDefaults() *EquipmentIoCardIdentityAllOf`

NewEquipmentIoCardIdentityAllOfWithDefaults instantiates a new EquipmentIoCardIdentityAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIoCardMoid

`func (o *EquipmentIoCardIdentityAllOf) GetIoCardMoid() string`

GetIoCardMoid returns the IoCardMoid field if non-nil, zero value otherwise.

### GetIoCardMoidOk

`func (o *EquipmentIoCardIdentityAllOf) GetIoCardMoidOk() (*string, bool)`

GetIoCardMoidOk returns a tuple with the IoCardMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoCardMoid

`func (o *EquipmentIoCardIdentityAllOf) SetIoCardMoid(v string)`

SetIoCardMoid sets IoCardMoid field to given value.

### HasIoCardMoid

`func (o *EquipmentIoCardIdentityAllOf) HasIoCardMoid() bool`

HasIoCardMoid returns a boolean if a field has been set.

### GetModuleId

`func (o *EquipmentIoCardIdentityAllOf) GetModuleId() int64`

GetModuleId returns the ModuleId field if non-nil, zero value otherwise.

### GetModuleIdOk

`func (o *EquipmentIoCardIdentityAllOf) GetModuleIdOk() (*int64, bool)`

GetModuleIdOk returns a tuple with the ModuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleId

`func (o *EquipmentIoCardIdentityAllOf) SetModuleId(v int64)`

SetModuleId sets ModuleId field to given value.

### HasModuleId

`func (o *EquipmentIoCardIdentityAllOf) HasModuleId() bool`

HasModuleId returns a boolean if a field has been set.

### GetNetworkElementMoid

`func (o *EquipmentIoCardIdentityAllOf) GetNetworkElementMoid() string`

GetNetworkElementMoid returns the NetworkElementMoid field if non-nil, zero value otherwise.

### GetNetworkElementMoidOk

`func (o *EquipmentIoCardIdentityAllOf) GetNetworkElementMoidOk() (*string, bool)`

GetNetworkElementMoidOk returns a tuple with the NetworkElementMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElementMoid

`func (o *EquipmentIoCardIdentityAllOf) SetNetworkElementMoid(v string)`

SetNetworkElementMoid sets NetworkElementMoid field to given value.

### HasNetworkElementMoid

`func (o *EquipmentIoCardIdentityAllOf) HasNetworkElementMoid() bool`

HasNetworkElementMoid returns a boolean if a field has been set.

### GetSwitchId

`func (o *EquipmentIoCardIdentityAllOf) GetSwitchId() int64`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *EquipmentIoCardIdentityAllOf) GetSwitchIdOk() (*int64, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *EquipmentIoCardIdentityAllOf) SetSwitchId(v int64)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *EquipmentIoCardIdentityAllOf) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


