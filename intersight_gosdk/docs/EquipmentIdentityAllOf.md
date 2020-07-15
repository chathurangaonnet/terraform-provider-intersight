# EquipmentIdentityAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminAction** | Pointer to **string** | Updated by UI/API to trigger specific chassis action type. | [optional] [default to "None"]
**AdminActionState** | Pointer to **string** | The state of Maintenance Action performed. This will have three states. Applying - Action is in progress. Applied - Action is completed and applied. Failed - Action has failed. | [optional] [default to "None"]
**Identifier** | Pointer to **int64** | Numeric Identifier assigned by the management system to the equipment. | [optional] 
**Lifecycle** | Pointer to **string** | The equipment&#39;s lifecycle status. | [optional] [default to "None"]
**Model** | Pointer to **string** | The vendor provided model name for the equipment. | [optional] 
**Serial** | Pointer to **string** | The serial number of the equipment. | [optional] 
**Vendor** | Pointer to **string** | The manufacturer of the equipment. | [optional] 
**DeviceRegistration** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewEquipmentIdentityAllOf

`func NewEquipmentIdentityAllOf() *EquipmentIdentityAllOf`

NewEquipmentIdentityAllOf instantiates a new EquipmentIdentityAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentIdentityAllOfWithDefaults

`func NewEquipmentIdentityAllOfWithDefaults() *EquipmentIdentityAllOf`

NewEquipmentIdentityAllOfWithDefaults instantiates a new EquipmentIdentityAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminAction

`func (o *EquipmentIdentityAllOf) GetAdminAction() string`

GetAdminAction returns the AdminAction field if non-nil, zero value otherwise.

### GetAdminActionOk

`func (o *EquipmentIdentityAllOf) GetAdminActionOk() (*string, bool)`

GetAdminActionOk returns a tuple with the AdminAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminAction

`func (o *EquipmentIdentityAllOf) SetAdminAction(v string)`

SetAdminAction sets AdminAction field to given value.

### HasAdminAction

`func (o *EquipmentIdentityAllOf) HasAdminAction() bool`

HasAdminAction returns a boolean if a field has been set.

### GetAdminActionState

`func (o *EquipmentIdentityAllOf) GetAdminActionState() string`

GetAdminActionState returns the AdminActionState field if non-nil, zero value otherwise.

### GetAdminActionStateOk

`func (o *EquipmentIdentityAllOf) GetAdminActionStateOk() (*string, bool)`

GetAdminActionStateOk returns a tuple with the AdminActionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminActionState

`func (o *EquipmentIdentityAllOf) SetAdminActionState(v string)`

SetAdminActionState sets AdminActionState field to given value.

### HasAdminActionState

`func (o *EquipmentIdentityAllOf) HasAdminActionState() bool`

HasAdminActionState returns a boolean if a field has been set.

### GetIdentifier

`func (o *EquipmentIdentityAllOf) GetIdentifier() int64`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *EquipmentIdentityAllOf) GetIdentifierOk() (*int64, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *EquipmentIdentityAllOf) SetIdentifier(v int64)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *EquipmentIdentityAllOf) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetLifecycle

`func (o *EquipmentIdentityAllOf) GetLifecycle() string`

GetLifecycle returns the Lifecycle field if non-nil, zero value otherwise.

### GetLifecycleOk

`func (o *EquipmentIdentityAllOf) GetLifecycleOk() (*string, bool)`

GetLifecycleOk returns a tuple with the Lifecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycle

`func (o *EquipmentIdentityAllOf) SetLifecycle(v string)`

SetLifecycle sets Lifecycle field to given value.

### HasLifecycle

`func (o *EquipmentIdentityAllOf) HasLifecycle() bool`

HasLifecycle returns a boolean if a field has been set.

### GetModel

`func (o *EquipmentIdentityAllOf) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *EquipmentIdentityAllOf) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *EquipmentIdentityAllOf) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *EquipmentIdentityAllOf) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetSerial

`func (o *EquipmentIdentityAllOf) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *EquipmentIdentityAllOf) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *EquipmentIdentityAllOf) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *EquipmentIdentityAllOf) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetVendor

`func (o *EquipmentIdentityAllOf) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *EquipmentIdentityAllOf) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *EquipmentIdentityAllOf) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *EquipmentIdentityAllOf) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetDeviceRegistration

`func (o *EquipmentIdentityAllOf) GetDeviceRegistration() AssetDeviceRegistrationRelationship`

GetDeviceRegistration returns the DeviceRegistration field if non-nil, zero value otherwise.

### GetDeviceRegistrationOk

`func (o *EquipmentIdentityAllOf) GetDeviceRegistrationOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceRegistrationOk returns a tuple with the DeviceRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRegistration

`func (o *EquipmentIdentityAllOf) SetDeviceRegistration(v AssetDeviceRegistrationRelationship)`

SetDeviceRegistration sets DeviceRegistration field to given value.

### HasDeviceRegistration

`func (o *EquipmentIdentityAllOf) HasDeviceRegistration() bool`

HasDeviceRegistration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


