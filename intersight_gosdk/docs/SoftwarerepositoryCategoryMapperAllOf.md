# SoftwarerepositoryCategoryMapperAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | The category of the model series. | [optional] 
**FileType** | Pointer to **string** | The type of distributable image, example huu, scu, driver, os. | [optional] [default to "Distributable"]
**MdfId** | Pointer to **string** | Cisco software repository image category identifier. | [optional] 
**RegexPattern** | Pointer to **string** | The regex that all images of this category follow. | [optional] 
**Source** | Pointer to **string** | The image can be downloaded from cisco.com or external cloud store. | [optional] [default to "Cisco"]
**SupportedModels** | Pointer to **[]string** |  | [optional] 
**SwId** | Pointer to **string** | The software type id provided by cisco.com. | [optional] 
**TagTypes** | Pointer to **[]string** |  | [optional] 
**Version** | Pointer to **string** | The version from which user can download images from amazon store, if source is external cloud store. | [optional] 

## Methods

### NewSoftwarerepositoryCategoryMapperAllOf

`func NewSoftwarerepositoryCategoryMapperAllOf() *SoftwarerepositoryCategoryMapperAllOf`

NewSoftwarerepositoryCategoryMapperAllOf instantiates a new SoftwarerepositoryCategoryMapperAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwarerepositoryCategoryMapperAllOfWithDefaults

`func NewSoftwarerepositoryCategoryMapperAllOfWithDefaults() *SoftwarerepositoryCategoryMapperAllOf`

NewSoftwarerepositoryCategoryMapperAllOfWithDefaults instantiates a new SoftwarerepositoryCategoryMapperAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *SoftwarerepositoryCategoryMapperAllOf) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *SoftwarerepositoryCategoryMapperAllOf) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *SoftwarerepositoryCategoryMapperAllOf) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *SoftwarerepositoryCategoryMapperAllOf) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetFileType

`func (o *SoftwarerepositoryCategoryMapperAllOf) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *SoftwarerepositoryCategoryMapperAllOf) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *SoftwarerepositoryCategoryMapperAllOf) SetFileType(v string)`

SetFileType sets FileType field to given value.

### HasFileType

`func (o *SoftwarerepositoryCategoryMapperAllOf) HasFileType() bool`

HasFileType returns a boolean if a field has been set.

### GetMdfId

`func (o *SoftwarerepositoryCategoryMapperAllOf) GetMdfId() string`

GetMdfId returns the MdfId field if non-nil, zero value otherwise.

### GetMdfIdOk

`func (o *SoftwarerepositoryCategoryMapperAllOf) GetMdfIdOk() (*string, bool)`

GetMdfIdOk returns a tuple with the MdfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdfId

`func (o *SoftwarerepositoryCategoryMapperAllOf) SetMdfId(v string)`

SetMdfId sets MdfId field to given value.

### HasMdfId

`func (o *SoftwarerepositoryCategoryMapperAllOf) HasMdfId() bool`

HasMdfId returns a boolean if a field has been set.

### GetRegexPattern

`func (o *SoftwarerepositoryCategoryMapperAllOf) GetRegexPattern() string`

GetRegexPattern returns the RegexPattern field if non-nil, zero value otherwise.

### GetRegexPatternOk

`func (o *SoftwarerepositoryCategoryMapperAllOf) GetRegexPatternOk() (*string, bool)`

GetRegexPatternOk returns a tuple with the RegexPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegexPattern

`func (o *SoftwarerepositoryCategoryMapperAllOf) SetRegexPattern(v string)`

SetRegexPattern sets RegexPattern field to given value.

### HasRegexPattern

`func (o *SoftwarerepositoryCategoryMapperAllOf) HasRegexPattern() bool`

HasRegexPattern returns a boolean if a field has been set.

### GetSource

`func (o *SoftwarerepositoryCategoryMapperAllOf) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SoftwarerepositoryCategoryMapperAllOf) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SoftwarerepositoryCategoryMapperAllOf) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *SoftwarerepositoryCategoryMapperAllOf) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSupportedModels

`func (o *SoftwarerepositoryCategoryMapperAllOf) GetSupportedModels() []string`

GetSupportedModels returns the SupportedModels field if non-nil, zero value otherwise.

### GetSupportedModelsOk

`func (o *SoftwarerepositoryCategoryMapperAllOf) GetSupportedModelsOk() (*[]string, bool)`

GetSupportedModelsOk returns a tuple with the SupportedModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedModels

`func (o *SoftwarerepositoryCategoryMapperAllOf) SetSupportedModels(v []string)`

SetSupportedModels sets SupportedModels field to given value.

### HasSupportedModels

`func (o *SoftwarerepositoryCategoryMapperAllOf) HasSupportedModels() bool`

HasSupportedModels returns a boolean if a field has been set.

### GetSwId

`func (o *SoftwarerepositoryCategoryMapperAllOf) GetSwId() string`

GetSwId returns the SwId field if non-nil, zero value otherwise.

### GetSwIdOk

`func (o *SoftwarerepositoryCategoryMapperAllOf) GetSwIdOk() (*string, bool)`

GetSwIdOk returns a tuple with the SwId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwId

`func (o *SoftwarerepositoryCategoryMapperAllOf) SetSwId(v string)`

SetSwId sets SwId field to given value.

### HasSwId

`func (o *SoftwarerepositoryCategoryMapperAllOf) HasSwId() bool`

HasSwId returns a boolean if a field has been set.

### GetTagTypes

`func (o *SoftwarerepositoryCategoryMapperAllOf) GetTagTypes() []string`

GetTagTypes returns the TagTypes field if non-nil, zero value otherwise.

### GetTagTypesOk

`func (o *SoftwarerepositoryCategoryMapperAllOf) GetTagTypesOk() (*[]string, bool)`

GetTagTypesOk returns a tuple with the TagTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagTypes

`func (o *SoftwarerepositoryCategoryMapperAllOf) SetTagTypes(v []string)`

SetTagTypes sets TagTypes field to given value.

### HasTagTypes

`func (o *SoftwarerepositoryCategoryMapperAllOf) HasTagTypes() bool`

HasTagTypes returns a boolean if a field has been set.

### GetVersion

`func (o *SoftwarerepositoryCategoryMapperAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SoftwarerepositoryCategoryMapperAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SoftwarerepositoryCategoryMapperAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SoftwarerepositoryCategoryMapperAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


