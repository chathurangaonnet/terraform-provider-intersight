# SoftwarerepositoryLocalMachineAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownloadUrl** | Pointer to **string** | When the import action in the file MO is updated with &#39;GeneratePreSignedDownloadUrl&#39;, Intersight returns a pre-signed URL in this property as part of the patch response. The user is expected to subsequently download the file using this URL. | [optional] [readonly] 
**PartSize** | Pointer to **int64** | The chunk size (in bytes) for each part of the file to be uploaded. | [optional] 
**UploadId** | Pointer to **string** | When the import action in file MO is updated with &#39;GeneratePreSignedUploadUrl&#39;, Intersight shall return a upload Id in this property as part of the PATCH response. | [optional] 
**UploadUrl** | Pointer to **string** | When a file MO is created with &#39;LocalMachine&#39; as the source, Intersight returns a pre-signed URL in this property as part of the POST response. The user is expected to subsequently upload the file content using this URL. Once the upload is completed, the user is expected to patch the Uploader object&#39;s transfer to a success state. | [optional] [readonly] 
**UploadUrls** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSoftwarerepositoryLocalMachineAllOf

`func NewSoftwarerepositoryLocalMachineAllOf() *SoftwarerepositoryLocalMachineAllOf`

NewSoftwarerepositoryLocalMachineAllOf instantiates a new SoftwarerepositoryLocalMachineAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwarerepositoryLocalMachineAllOfWithDefaults

`func NewSoftwarerepositoryLocalMachineAllOfWithDefaults() *SoftwarerepositoryLocalMachineAllOf`

NewSoftwarerepositoryLocalMachineAllOfWithDefaults instantiates a new SoftwarerepositoryLocalMachineAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownloadUrl

`func (o *SoftwarerepositoryLocalMachineAllOf) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *SoftwarerepositoryLocalMachineAllOf) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *SoftwarerepositoryLocalMachineAllOf) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *SoftwarerepositoryLocalMachineAllOf) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.

### GetPartSize

`func (o *SoftwarerepositoryLocalMachineAllOf) GetPartSize() int64`

GetPartSize returns the PartSize field if non-nil, zero value otherwise.

### GetPartSizeOk

`func (o *SoftwarerepositoryLocalMachineAllOf) GetPartSizeOk() (*int64, bool)`

GetPartSizeOk returns a tuple with the PartSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartSize

`func (o *SoftwarerepositoryLocalMachineAllOf) SetPartSize(v int64)`

SetPartSize sets PartSize field to given value.

### HasPartSize

`func (o *SoftwarerepositoryLocalMachineAllOf) HasPartSize() bool`

HasPartSize returns a boolean if a field has been set.

### GetUploadId

`func (o *SoftwarerepositoryLocalMachineAllOf) GetUploadId() string`

GetUploadId returns the UploadId field if non-nil, zero value otherwise.

### GetUploadIdOk

`func (o *SoftwarerepositoryLocalMachineAllOf) GetUploadIdOk() (*string, bool)`

GetUploadIdOk returns a tuple with the UploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadId

`func (o *SoftwarerepositoryLocalMachineAllOf) SetUploadId(v string)`

SetUploadId sets UploadId field to given value.

### HasUploadId

`func (o *SoftwarerepositoryLocalMachineAllOf) HasUploadId() bool`

HasUploadId returns a boolean if a field has been set.

### GetUploadUrl

`func (o *SoftwarerepositoryLocalMachineAllOf) GetUploadUrl() string`

GetUploadUrl returns the UploadUrl field if non-nil, zero value otherwise.

### GetUploadUrlOk

`func (o *SoftwarerepositoryLocalMachineAllOf) GetUploadUrlOk() (*string, bool)`

GetUploadUrlOk returns a tuple with the UploadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadUrl

`func (o *SoftwarerepositoryLocalMachineAllOf) SetUploadUrl(v string)`

SetUploadUrl sets UploadUrl field to given value.

### HasUploadUrl

`func (o *SoftwarerepositoryLocalMachineAllOf) HasUploadUrl() bool`

HasUploadUrl returns a boolean if a field has been set.

### GetUploadUrls

`func (o *SoftwarerepositoryLocalMachineAllOf) GetUploadUrls() []string`

GetUploadUrls returns the UploadUrls field if non-nil, zero value otherwise.

### GetUploadUrlsOk

`func (o *SoftwarerepositoryLocalMachineAllOf) GetUploadUrlsOk() (*[]string, bool)`

GetUploadUrlsOk returns a tuple with the UploadUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadUrls

`func (o *SoftwarerepositoryLocalMachineAllOf) SetUploadUrls(v []string)`

SetUploadUrls sets UploadUrls field to given value.

### HasUploadUrls

`func (o *SoftwarerepositoryLocalMachineAllOf) HasUploadUrls() bool`

HasUploadUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


