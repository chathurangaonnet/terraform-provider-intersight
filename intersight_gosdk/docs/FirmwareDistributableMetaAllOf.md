# FirmwareDistributableMetaAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketName** | Pointer to **string** | The S3 bucket name where the images are present, if source is external cloud store. | [optional] 
**FileType** | Pointer to **string** | The type of distributable image, example huu, scu, driver, os. | [optional] [default to "Distributable"]
**FromVersion** | Pointer to **string** | The version from which user can download images from amazon store, if source is external cloud store. | [optional] 
**Mdfid** | Pointer to **string** | The mdfid of the image provided by cisco.com. | [optional] 
**SoftwareTypeId** | Pointer to **string** | The software type id provided by cisco.com. | [optional] 
**Source** | Pointer to **string** | The image can be downloaded from cisco.com or external cloud store. | [optional] [default to "Cisco"]
**SupportedModels** | Pointer to **[]string** |  | [optional] 
**ToVersion** | Pointer to **string** | The version till which user can download images from amazon store, if source is external cloud store. | [optional] 

## Methods

### NewFirmwareDistributableMetaAllOf

`func NewFirmwareDistributableMetaAllOf() *FirmwareDistributableMetaAllOf`

NewFirmwareDistributableMetaAllOf instantiates a new FirmwareDistributableMetaAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareDistributableMetaAllOfWithDefaults

`func NewFirmwareDistributableMetaAllOfWithDefaults() *FirmwareDistributableMetaAllOf`

NewFirmwareDistributableMetaAllOfWithDefaults instantiates a new FirmwareDistributableMetaAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *FirmwareDistributableMetaAllOf) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *FirmwareDistributableMetaAllOf) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *FirmwareDistributableMetaAllOf) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *FirmwareDistributableMetaAllOf) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.

### GetFileType

`func (o *FirmwareDistributableMetaAllOf) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *FirmwareDistributableMetaAllOf) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *FirmwareDistributableMetaAllOf) SetFileType(v string)`

SetFileType sets FileType field to given value.

### HasFileType

`func (o *FirmwareDistributableMetaAllOf) HasFileType() bool`

HasFileType returns a boolean if a field has been set.

### GetFromVersion

`func (o *FirmwareDistributableMetaAllOf) GetFromVersion() string`

GetFromVersion returns the FromVersion field if non-nil, zero value otherwise.

### GetFromVersionOk

`func (o *FirmwareDistributableMetaAllOf) GetFromVersionOk() (*string, bool)`

GetFromVersionOk returns a tuple with the FromVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromVersion

`func (o *FirmwareDistributableMetaAllOf) SetFromVersion(v string)`

SetFromVersion sets FromVersion field to given value.

### HasFromVersion

`func (o *FirmwareDistributableMetaAllOf) HasFromVersion() bool`

HasFromVersion returns a boolean if a field has been set.

### GetMdfid

`func (o *FirmwareDistributableMetaAllOf) GetMdfid() string`

GetMdfid returns the Mdfid field if non-nil, zero value otherwise.

### GetMdfidOk

`func (o *FirmwareDistributableMetaAllOf) GetMdfidOk() (*string, bool)`

GetMdfidOk returns a tuple with the Mdfid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdfid

`func (o *FirmwareDistributableMetaAllOf) SetMdfid(v string)`

SetMdfid sets Mdfid field to given value.

### HasMdfid

`func (o *FirmwareDistributableMetaAllOf) HasMdfid() bool`

HasMdfid returns a boolean if a field has been set.

### GetSoftwareTypeId

`func (o *FirmwareDistributableMetaAllOf) GetSoftwareTypeId() string`

GetSoftwareTypeId returns the SoftwareTypeId field if non-nil, zero value otherwise.

### GetSoftwareTypeIdOk

`func (o *FirmwareDistributableMetaAllOf) GetSoftwareTypeIdOk() (*string, bool)`

GetSoftwareTypeIdOk returns a tuple with the SoftwareTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareTypeId

`func (o *FirmwareDistributableMetaAllOf) SetSoftwareTypeId(v string)`

SetSoftwareTypeId sets SoftwareTypeId field to given value.

### HasSoftwareTypeId

`func (o *FirmwareDistributableMetaAllOf) HasSoftwareTypeId() bool`

HasSoftwareTypeId returns a boolean if a field has been set.

### GetSource

`func (o *FirmwareDistributableMetaAllOf) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *FirmwareDistributableMetaAllOf) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *FirmwareDistributableMetaAllOf) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *FirmwareDistributableMetaAllOf) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSupportedModels

`func (o *FirmwareDistributableMetaAllOf) GetSupportedModels() []string`

GetSupportedModels returns the SupportedModels field if non-nil, zero value otherwise.

### GetSupportedModelsOk

`func (o *FirmwareDistributableMetaAllOf) GetSupportedModelsOk() (*[]string, bool)`

GetSupportedModelsOk returns a tuple with the SupportedModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedModels

`func (o *FirmwareDistributableMetaAllOf) SetSupportedModels(v []string)`

SetSupportedModels sets SupportedModels field to given value.

### HasSupportedModels

`func (o *FirmwareDistributableMetaAllOf) HasSupportedModels() bool`

HasSupportedModels returns a boolean if a field has been set.

### GetToVersion

`func (o *FirmwareDistributableMetaAllOf) GetToVersion() string`

GetToVersion returns the ToVersion field if non-nil, zero value otherwise.

### GetToVersionOk

`func (o *FirmwareDistributableMetaAllOf) GetToVersionOk() (*string, bool)`

GetToVersionOk returns a tuple with the ToVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToVersion

`func (o *FirmwareDistributableMetaAllOf) SetToVersion(v string)`

SetToVersion sets ToVersion field to given value.

### HasToVersion

`func (o *FirmwareDistributableMetaAllOf) HasToVersion() bool`

HasToVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


