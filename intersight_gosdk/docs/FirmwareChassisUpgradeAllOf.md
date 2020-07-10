# FirmwareChassisUpgradeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeComponentList** | Pointer to **[]string** |  | [optional] 
**Chassis** | Pointer to [**EquipmentChassisRelationship**](equipment.Chassis.Relationship.md) |  | [optional] 
**Device** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewFirmwareChassisUpgradeAllOf

`func NewFirmwareChassisUpgradeAllOf() *FirmwareChassisUpgradeAllOf`

NewFirmwareChassisUpgradeAllOf instantiates a new FirmwareChassisUpgradeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareChassisUpgradeAllOfWithDefaults

`func NewFirmwareChassisUpgradeAllOfWithDefaults() *FirmwareChassisUpgradeAllOf`

NewFirmwareChassisUpgradeAllOfWithDefaults instantiates a new FirmwareChassisUpgradeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludeComponentList

`func (o *FirmwareChassisUpgradeAllOf) GetExcludeComponentList() []string`

GetExcludeComponentList returns the ExcludeComponentList field if non-nil, zero value otherwise.

### GetExcludeComponentListOk

`func (o *FirmwareChassisUpgradeAllOf) GetExcludeComponentListOk() (*[]string, bool)`

GetExcludeComponentListOk returns a tuple with the ExcludeComponentList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeComponentList

`func (o *FirmwareChassisUpgradeAllOf) SetExcludeComponentList(v []string)`

SetExcludeComponentList sets ExcludeComponentList field to given value.

### HasExcludeComponentList

`func (o *FirmwareChassisUpgradeAllOf) HasExcludeComponentList() bool`

HasExcludeComponentList returns a boolean if a field has been set.

### GetChassis

`func (o *FirmwareChassisUpgradeAllOf) GetChassis() EquipmentChassisRelationship`

GetChassis returns the Chassis field if non-nil, zero value otherwise.

### GetChassisOk

`func (o *FirmwareChassisUpgradeAllOf) GetChassisOk() (*EquipmentChassisRelationship, bool)`

GetChassisOk returns a tuple with the Chassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassis

`func (o *FirmwareChassisUpgradeAllOf) SetChassis(v EquipmentChassisRelationship)`

SetChassis sets Chassis field to given value.

### HasChassis

`func (o *FirmwareChassisUpgradeAllOf) HasChassis() bool`

HasChassis returns a boolean if a field has been set.

### GetDevice

`func (o *FirmwareChassisUpgradeAllOf) GetDevice() AssetDeviceRegistrationRelationship`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *FirmwareChassisUpgradeAllOf) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *FirmwareChassisUpgradeAllOf) SetDevice(v AssetDeviceRegistrationRelationship)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *FirmwareChassisUpgradeAllOf) HasDevice() bool`

HasDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


