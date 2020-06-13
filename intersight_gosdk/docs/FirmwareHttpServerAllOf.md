# FirmwareHttpServerAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationLink** | Pointer to **string** | HTTP/HTTPS link to the image. Accepted formats HTTP[s]://server-hostname/share/image or HTTP[s]://serverip/share/image. For a successful upgrade, the remote share server must have network connectivity with the CIMC of the selected devices. | [optional] 
**MountOptions** | Pointer to **string** | Mount option as configured on the HTTP server. Empty if nothing is configured. | [optional] 

## Methods

### NewFirmwareHttpServerAllOf

`func NewFirmwareHttpServerAllOf() *FirmwareHttpServerAllOf`

NewFirmwareHttpServerAllOf instantiates a new FirmwareHttpServerAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareHttpServerAllOfWithDefaults

`func NewFirmwareHttpServerAllOfWithDefaults() *FirmwareHttpServerAllOf`

NewFirmwareHttpServerAllOfWithDefaults instantiates a new FirmwareHttpServerAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocationLink

`func (o *FirmwareHttpServerAllOf) GetLocationLink() string`

GetLocationLink returns the LocationLink field if non-nil, zero value otherwise.

### GetLocationLinkOk

`func (o *FirmwareHttpServerAllOf) GetLocationLinkOk() (*string, bool)`

GetLocationLinkOk returns a tuple with the LocationLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationLink

`func (o *FirmwareHttpServerAllOf) SetLocationLink(v string)`

SetLocationLink sets LocationLink field to given value.

### HasLocationLink

`func (o *FirmwareHttpServerAllOf) HasLocationLink() bool`

HasLocationLink returns a boolean if a field has been set.

### GetMountOptions

`func (o *FirmwareHttpServerAllOf) GetMountOptions() string`

GetMountOptions returns the MountOptions field if non-nil, zero value otherwise.

### GetMountOptionsOk

`func (o *FirmwareHttpServerAllOf) GetMountOptionsOk() (*string, bool)`

GetMountOptionsOk returns a tuple with the MountOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountOptions

`func (o *FirmwareHttpServerAllOf) SetMountOptions(v string)`

SetMountOptions sets MountOptions field to given value.

### HasMountOptions

`func (o *FirmwareHttpServerAllOf) HasMountOptions() bool`

HasMountOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


