# VnicFcInterruptSettingsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | The preferred driver interrupt mode. This can be one of the following:- MSIx — Message Signaled Interrupts (MSI) with the optional extension. MSI   — MSI only. INTx  — PCI INTx interrupts. MSIx is the recommended option. | [optional] [default to "MSIx"]

## Methods

### NewVnicFcInterruptSettingsAllOf

`func NewVnicFcInterruptSettingsAllOf() *VnicFcInterruptSettingsAllOf`

NewVnicFcInterruptSettingsAllOf instantiates a new VnicFcInterruptSettingsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicFcInterruptSettingsAllOfWithDefaults

`func NewVnicFcInterruptSettingsAllOfWithDefaults() *VnicFcInterruptSettingsAllOf`

NewVnicFcInterruptSettingsAllOfWithDefaults instantiates a new VnicFcInterruptSettingsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *VnicFcInterruptSettingsAllOf) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *VnicFcInterruptSettingsAllOf) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *VnicFcInterruptSettingsAllOf) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *VnicFcInterruptSettingsAllOf) HasMode() bool`

HasMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


