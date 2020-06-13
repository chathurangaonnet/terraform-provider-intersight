# SoftwarerepositoryFileAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | User provided description about the file. Cisco provided description for image inventoried from a Cisco repository. | [optional] 
**DownloadCount** | Pointer to **int64** | The number of times this file has been downloaded from the local repository. It is used by the repository monitoring process to determine the files that are to be evicted from the cache. | [optional] [readonly] 
**ImportAction** | Pointer to **string** | The action to be performed on the imported file. If &#39;PreCache&#39; is set, the image will be cached in Appliance. Applicable in Intersight appliance deployment. If &#39;Evict&#39; is set, the cached file will be removed. Applicable in Intersight appliance deployment. If &#39;GeneratePreSignedUploadUrl&#39; is set, generates pre signed URL (s) for the file to be imported into the repository. Applicable for local machine source. The URL (s) will be populated under LocalMachine file server. If &#39;CompleteImportProcess&#39; is set, the ImportState is marked as &#39;Imported&#39;. Applicable for local machine source. If &#39;Cancel&#39; is set, the ImportState is marked as &#39;Failed&#39;. Applicable for local machine source. | [optional] [default to "None"]
**ImportState** | Pointer to **string** | The state  of this file in the repository or Appliance. The importState is updated during the import operation and as part of the repository monitoring process. | [optional] [readonly] [default to "ReadyForImport"]
**ImportedTime** | Pointer to [**time.Time**](time.Time.md) | The time at which this image or file was imported/cached into the repositry. if the &#39;ImportState&#39; is &#39;Imported&#39;, the time at which this image or file was imported. if the &#39;ImportState&#39; is &#39;Cached&#39;, the time at which this image or file was cached. | [optional] [readonly] 
**LastAccessTime** | Pointer to [**time.Time**](time.Time.md) | The time at which this file was last downloaded from the local repository. It is used by the repository monitoring process to determine the files that are to be evicted from the cache. | [optional] [readonly] 
**Md5sum** | Pointer to **string** | The md5sum checksum of the file. This information is available for all Cisco distributed images and files imported to the local repository. | [optional] 
**Name** | Pointer to **string** | The name of the file. It is populated as part of the image import operation. | [optional] 
**ReleaseDate** | Pointer to [**time.Time**](time.Time.md) | The date on which the file was released or distributed by its vendor. | [optional] 
**Sha512sum** | Pointer to **string** | The sha512sum of the file. This information is available for all Cisco distributed images and files imported to the local repository. | [optional] 
**Size** | Pointer to **int64** | The size (in bytes) of the file. This information is available for all Cisco distributed images and files imported to the local repository. | [optional] 
**SoftwareAdvisoryUrl** | Pointer to **string** | The software advisory, if any, provided by the vendor for this file. | [optional] 
**Source** | Pointer to [**SoftwarerepositoryFileServer**](softwarerepository.FileServer.md) |  | [optional] 
**Version** | Pointer to **string** | Vendor provided version for the file. | [optional] 

## Methods

### NewSoftwarerepositoryFileAllOf

`func NewSoftwarerepositoryFileAllOf() *SoftwarerepositoryFileAllOf`

NewSoftwarerepositoryFileAllOf instantiates a new SoftwarerepositoryFileAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwarerepositoryFileAllOfWithDefaults

`func NewSoftwarerepositoryFileAllOfWithDefaults() *SoftwarerepositoryFileAllOf`

NewSoftwarerepositoryFileAllOfWithDefaults instantiates a new SoftwarerepositoryFileAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SoftwarerepositoryFileAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SoftwarerepositoryFileAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SoftwarerepositoryFileAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SoftwarerepositoryFileAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDownloadCount

`func (o *SoftwarerepositoryFileAllOf) GetDownloadCount() int64`

GetDownloadCount returns the DownloadCount field if non-nil, zero value otherwise.

### GetDownloadCountOk

`func (o *SoftwarerepositoryFileAllOf) GetDownloadCountOk() (*int64, bool)`

GetDownloadCountOk returns a tuple with the DownloadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadCount

`func (o *SoftwarerepositoryFileAllOf) SetDownloadCount(v int64)`

SetDownloadCount sets DownloadCount field to given value.

### HasDownloadCount

`func (o *SoftwarerepositoryFileAllOf) HasDownloadCount() bool`

HasDownloadCount returns a boolean if a field has been set.

### GetImportAction

`func (o *SoftwarerepositoryFileAllOf) GetImportAction() string`

GetImportAction returns the ImportAction field if non-nil, zero value otherwise.

### GetImportActionOk

`func (o *SoftwarerepositoryFileAllOf) GetImportActionOk() (*string, bool)`

GetImportActionOk returns a tuple with the ImportAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportAction

`func (o *SoftwarerepositoryFileAllOf) SetImportAction(v string)`

SetImportAction sets ImportAction field to given value.

### HasImportAction

`func (o *SoftwarerepositoryFileAllOf) HasImportAction() bool`

HasImportAction returns a boolean if a field has been set.

### GetImportState

`func (o *SoftwarerepositoryFileAllOf) GetImportState() string`

GetImportState returns the ImportState field if non-nil, zero value otherwise.

### GetImportStateOk

`func (o *SoftwarerepositoryFileAllOf) GetImportStateOk() (*string, bool)`

GetImportStateOk returns a tuple with the ImportState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportState

`func (o *SoftwarerepositoryFileAllOf) SetImportState(v string)`

SetImportState sets ImportState field to given value.

### HasImportState

`func (o *SoftwarerepositoryFileAllOf) HasImportState() bool`

HasImportState returns a boolean if a field has been set.

### GetImportedTime

`func (o *SoftwarerepositoryFileAllOf) GetImportedTime() time.Time`

GetImportedTime returns the ImportedTime field if non-nil, zero value otherwise.

### GetImportedTimeOk

`func (o *SoftwarerepositoryFileAllOf) GetImportedTimeOk() (*time.Time, bool)`

GetImportedTimeOk returns a tuple with the ImportedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedTime

`func (o *SoftwarerepositoryFileAllOf) SetImportedTime(v time.Time)`

SetImportedTime sets ImportedTime field to given value.

### HasImportedTime

`func (o *SoftwarerepositoryFileAllOf) HasImportedTime() bool`

HasImportedTime returns a boolean if a field has been set.

### GetLastAccessTime

`func (o *SoftwarerepositoryFileAllOf) GetLastAccessTime() time.Time`

GetLastAccessTime returns the LastAccessTime field if non-nil, zero value otherwise.

### GetLastAccessTimeOk

`func (o *SoftwarerepositoryFileAllOf) GetLastAccessTimeOk() (*time.Time, bool)`

GetLastAccessTimeOk returns a tuple with the LastAccessTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessTime

`func (o *SoftwarerepositoryFileAllOf) SetLastAccessTime(v time.Time)`

SetLastAccessTime sets LastAccessTime field to given value.

### HasLastAccessTime

`func (o *SoftwarerepositoryFileAllOf) HasLastAccessTime() bool`

HasLastAccessTime returns a boolean if a field has been set.

### GetMd5sum

`func (o *SoftwarerepositoryFileAllOf) GetMd5sum() string`

GetMd5sum returns the Md5sum field if non-nil, zero value otherwise.

### GetMd5sumOk

`func (o *SoftwarerepositoryFileAllOf) GetMd5sumOk() (*string, bool)`

GetMd5sumOk returns a tuple with the Md5sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5sum

`func (o *SoftwarerepositoryFileAllOf) SetMd5sum(v string)`

SetMd5sum sets Md5sum field to given value.

### HasMd5sum

`func (o *SoftwarerepositoryFileAllOf) HasMd5sum() bool`

HasMd5sum returns a boolean if a field has been set.

### GetName

`func (o *SoftwarerepositoryFileAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SoftwarerepositoryFileAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SoftwarerepositoryFileAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SoftwarerepositoryFileAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReleaseDate

`func (o *SoftwarerepositoryFileAllOf) GetReleaseDate() time.Time`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *SoftwarerepositoryFileAllOf) GetReleaseDateOk() (*time.Time, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *SoftwarerepositoryFileAllOf) SetReleaseDate(v time.Time)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *SoftwarerepositoryFileAllOf) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### GetSha512sum

`func (o *SoftwarerepositoryFileAllOf) GetSha512sum() string`

GetSha512sum returns the Sha512sum field if non-nil, zero value otherwise.

### GetSha512sumOk

`func (o *SoftwarerepositoryFileAllOf) GetSha512sumOk() (*string, bool)`

GetSha512sumOk returns a tuple with the Sha512sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha512sum

`func (o *SoftwarerepositoryFileAllOf) SetSha512sum(v string)`

SetSha512sum sets Sha512sum field to given value.

### HasSha512sum

`func (o *SoftwarerepositoryFileAllOf) HasSha512sum() bool`

HasSha512sum returns a boolean if a field has been set.

### GetSize

`func (o *SoftwarerepositoryFileAllOf) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *SoftwarerepositoryFileAllOf) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *SoftwarerepositoryFileAllOf) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *SoftwarerepositoryFileAllOf) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSoftwareAdvisoryUrl

`func (o *SoftwarerepositoryFileAllOf) GetSoftwareAdvisoryUrl() string`

GetSoftwareAdvisoryUrl returns the SoftwareAdvisoryUrl field if non-nil, zero value otherwise.

### GetSoftwareAdvisoryUrlOk

`func (o *SoftwarerepositoryFileAllOf) GetSoftwareAdvisoryUrlOk() (*string, bool)`

GetSoftwareAdvisoryUrlOk returns a tuple with the SoftwareAdvisoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareAdvisoryUrl

`func (o *SoftwarerepositoryFileAllOf) SetSoftwareAdvisoryUrl(v string)`

SetSoftwareAdvisoryUrl sets SoftwareAdvisoryUrl field to given value.

### HasSoftwareAdvisoryUrl

`func (o *SoftwarerepositoryFileAllOf) HasSoftwareAdvisoryUrl() bool`

HasSoftwareAdvisoryUrl returns a boolean if a field has been set.

### GetSource

`func (o *SoftwarerepositoryFileAllOf) GetSource() SoftwarerepositoryFileServer`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SoftwarerepositoryFileAllOf) GetSourceOk() (*SoftwarerepositoryFileServer, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SoftwarerepositoryFileAllOf) SetSource(v SoftwarerepositoryFileServer)`

SetSource sets Source field to given value.

### HasSource

`func (o *SoftwarerepositoryFileAllOf) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetVersion

`func (o *SoftwarerepositoryFileAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SoftwarerepositoryFileAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SoftwarerepositoryFileAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SoftwarerepositoryFileAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


