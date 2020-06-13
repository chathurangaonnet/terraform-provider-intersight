# FirmwareDirectDownload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpServer** | Pointer to [**FirmwareHttpServer**](firmware.HttpServer.md) |  | [optional] 
**ImageSource** | Pointer to **string** | Source type referring the image to be downloaded from CCO or from a local HTTPS server. | [optional] [default to "cisco"]
**IsPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;password&#39; property has been set. | [optional] [readonly] 
**Password** | Pointer to **string** | Password as configured on the local https server. | [optional] 
**Upgradeoption** | Pointer to **string** | Option to control the upgrade, e.g., sd_upgrade_mount_only - download the image into sd and upgrade wait for the server on-next boot. | [optional] [default to "sd_upgrade_mount_only"]
**Username** | Pointer to **string** | Username as configured on the local https server. | [optional] 

## Methods

### NewFirmwareDirectDownload

`func NewFirmwareDirectDownload() *FirmwareDirectDownload`

NewFirmwareDirectDownload instantiates a new FirmwareDirectDownload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareDirectDownloadWithDefaults

`func NewFirmwareDirectDownloadWithDefaults() *FirmwareDirectDownload`

NewFirmwareDirectDownloadWithDefaults instantiates a new FirmwareDirectDownload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpServer

`func (o *FirmwareDirectDownload) GetHttpServer() FirmwareHttpServer`

GetHttpServer returns the HttpServer field if non-nil, zero value otherwise.

### GetHttpServerOk

`func (o *FirmwareDirectDownload) GetHttpServerOk() (*FirmwareHttpServer, bool)`

GetHttpServerOk returns a tuple with the HttpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpServer

`func (o *FirmwareDirectDownload) SetHttpServer(v FirmwareHttpServer)`

SetHttpServer sets HttpServer field to given value.

### HasHttpServer

`func (o *FirmwareDirectDownload) HasHttpServer() bool`

HasHttpServer returns a boolean if a field has been set.

### GetImageSource

`func (o *FirmwareDirectDownload) GetImageSource() string`

GetImageSource returns the ImageSource field if non-nil, zero value otherwise.

### GetImageSourceOk

`func (o *FirmwareDirectDownload) GetImageSourceOk() (*string, bool)`

GetImageSourceOk returns a tuple with the ImageSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageSource

`func (o *FirmwareDirectDownload) SetImageSource(v string)`

SetImageSource sets ImageSource field to given value.

### HasImageSource

`func (o *FirmwareDirectDownload) HasImageSource() bool`

HasImageSource returns a boolean if a field has been set.

### GetIsPasswordSet

`func (o *FirmwareDirectDownload) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *FirmwareDirectDownload) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *FirmwareDirectDownload) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *FirmwareDirectDownload) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetPassword

`func (o *FirmwareDirectDownload) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *FirmwareDirectDownload) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *FirmwareDirectDownload) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *FirmwareDirectDownload) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUpgradeoption

`func (o *FirmwareDirectDownload) GetUpgradeoption() string`

GetUpgradeoption returns the Upgradeoption field if non-nil, zero value otherwise.

### GetUpgradeoptionOk

`func (o *FirmwareDirectDownload) GetUpgradeoptionOk() (*string, bool)`

GetUpgradeoptionOk returns a tuple with the Upgradeoption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeoption

`func (o *FirmwareDirectDownload) SetUpgradeoption(v string)`

SetUpgradeoption sets Upgradeoption field to given value.

### HasUpgradeoption

`func (o *FirmwareDirectDownload) HasUpgradeoption() bool`

HasUpgradeoption returns a boolean if a field has been set.

### GetUsername

`func (o *FirmwareDirectDownload) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *FirmwareDirectDownload) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *FirmwareDirectDownload) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *FirmwareDirectDownload) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


