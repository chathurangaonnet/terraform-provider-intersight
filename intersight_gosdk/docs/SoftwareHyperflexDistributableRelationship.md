# SoftwareHyperflexDistributableRelationship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountMoid** | Pointer to **string** | The Account ID for this managed object. | [optional] [readonly] 
**ClassId** | Pointer to **string** | The concrete type of this complex type. Its value must be the same as the &#39;objectType&#39; property. The OpenAPI document references this property as a discriminator value. | [readonly] 
**CreateTime** | Pointer to [**time.Time**](time.Time.md) | The time when this managed object was created. | [optional] [readonly] 
**DomainGroupMoid** | Pointer to **string** | The DomainGroup ID for this managed object. | [optional] [readonly] 
**ModTime** | Pointer to [**time.Time**](time.Time.md) | The time when this managed object was last modified. | [optional] [readonly] 
**Moid** | Pointer to **string** | The unique identifier of this Managed Object instance. | [optional] 
**ObjectType** | Pointer to **string** | The fully-qualified type of this managed object, i.e. the class name. This property is optional. The ObjectType is implied from the URL path. If specified, the value of objectType must match the class name specified in the URL path. | [readonly] 
**Owners** | Pointer to **[]string** |  | [optional] 
**SharedScope** | Pointer to **string** | Intersight provides pre-built workflows, tasks and policies to end users through global catalogs. Objects that are made available through global catalogs are said to have a &#39;shared&#39; ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. | [optional] [readonly] 
**Tags** | Pointer to [**[]MoTag**](mo.Tag.md) |  | [optional] 
**VersionContext** | Pointer to [**MoVersionContext**](mo.VersionContext.md) |  | [optional] 
**Ancestors** | Pointer to [**[]MoBaseMoRelationship**](mo.BaseMo.Relationship.md) | An array of relationships to moBaseMo resources. | [optional] [readonly] 
**Parent** | Pointer to [**MoBaseMoRelationship**](mo.BaseMo.Relationship.md) |  | [optional] 
**PermissionResources** | Pointer to [**[]MoBaseMoRelationship**](mo.BaseMo.Relationship.md) | An array of relationships to moBaseMo resources. | [optional] [readonly] 
**DisplayNames** | Pointer to [**map[string][]string**](array.md) | a map of display names for a resource. | [optional] [readonly] 
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
**BundleType** | Pointer to **string** | The bundle type of the image, as published on cisco.com. | [optional] [readonly] 
**ComponentMeta** | Pointer to [**[]FirmwareComponentMeta**](firmware.ComponentMeta.md) |  | [optional] 
**Guid** | Pointer to **string** | The unique identifier for an image in a Cisco repository. | [optional] [readonly] 
**Mdfid** | Pointer to **string** | The mdfid of the image provided by cisco.com. | [optional] 
**Model** | Pointer to **string** | The endpoint model for which this firmware image is applicable. | [optional] 
**PlatformType** | Pointer to **string** | The platform type of the image. | [optional] [readonly] 
**RecommendedBuild** | Pointer to **string** | The build which is recommended by Cisco. | [optional] 
**ReleaseNotesUrl** | Pointer to **string** | The url for the release notes of this image. | [optional] 
**SoftwareTypeId** | Pointer to **string** | The software type id provided by cisco.com. | [optional] [readonly] 
**SupportedModels** | Pointer to **[]string** |  | [optional] 
**Vendor** | Pointer to **string** | The vendor or publisher of this file. | [optional] 
**DistributableMetas** | Pointer to [**[]FirmwareDistributableMetaRelationship**](firmware.DistributableMeta.Relationship.md) | An array of relationships to firmwareDistributableMeta resources. | [optional] 
**Release** | Pointer to [**SoftwarerepositoryReleaseRelationship**](softwarerepository.Release.Relationship.md) |  | [optional] 
**Catalog** | Pointer to [**SoftwarerepositoryCatalogRelationship**](softwarerepository.Catalog.Relationship.md) |  | [optional] 

## Methods

### NewSoftwareHyperflexDistributableRelationship

`func NewSoftwareHyperflexDistributableRelationship(classId string, objectType string, ) *SoftwareHyperflexDistributableRelationship`

NewSoftwareHyperflexDistributableRelationship instantiates a new SoftwareHyperflexDistributableRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwareHyperflexDistributableRelationshipWithDefaults

`func NewSoftwareHyperflexDistributableRelationshipWithDefaults() *SoftwareHyperflexDistributableRelationship`

NewSoftwareHyperflexDistributableRelationshipWithDefaults instantiates a new SoftwareHyperflexDistributableRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *SoftwareHyperflexDistributableRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *SoftwareHyperflexDistributableRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *SoftwareHyperflexDistributableRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *SoftwareHyperflexDistributableRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *SoftwareHyperflexDistributableRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SoftwareHyperflexDistributableRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SoftwareHyperflexDistributableRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *SoftwareHyperflexDistributableRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *SoftwareHyperflexDistributableRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *SoftwareHyperflexDistributableRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *SoftwareHyperflexDistributableRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *SoftwareHyperflexDistributableRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *SoftwareHyperflexDistributableRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *SoftwareHyperflexDistributableRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *SoftwareHyperflexDistributableRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *SoftwareHyperflexDistributableRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *SoftwareHyperflexDistributableRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *SoftwareHyperflexDistributableRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *SoftwareHyperflexDistributableRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *SoftwareHyperflexDistributableRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *SoftwareHyperflexDistributableRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *SoftwareHyperflexDistributableRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *SoftwareHyperflexDistributableRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *SoftwareHyperflexDistributableRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SoftwareHyperflexDistributableRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SoftwareHyperflexDistributableRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *SoftwareHyperflexDistributableRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *SoftwareHyperflexDistributableRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *SoftwareHyperflexDistributableRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *SoftwareHyperflexDistributableRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *SoftwareHyperflexDistributableRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *SoftwareHyperflexDistributableRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *SoftwareHyperflexDistributableRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *SoftwareHyperflexDistributableRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *SoftwareHyperflexDistributableRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SoftwareHyperflexDistributableRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SoftwareHyperflexDistributableRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SoftwareHyperflexDistributableRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *SoftwareHyperflexDistributableRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *SoftwareHyperflexDistributableRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *SoftwareHyperflexDistributableRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *SoftwareHyperflexDistributableRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *SoftwareHyperflexDistributableRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *SoftwareHyperflexDistributableRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *SoftwareHyperflexDistributableRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *SoftwareHyperflexDistributableRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *SoftwareHyperflexDistributableRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *SoftwareHyperflexDistributableRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *SoftwareHyperflexDistributableRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *SoftwareHyperflexDistributableRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *SoftwareHyperflexDistributableRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *SoftwareHyperflexDistributableRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *SoftwareHyperflexDistributableRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *SoftwareHyperflexDistributableRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *SoftwareHyperflexDistributableRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *SoftwareHyperflexDistributableRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *SoftwareHyperflexDistributableRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *SoftwareHyperflexDistributableRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *SoftwareHyperflexDistributableRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *SoftwareHyperflexDistributableRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *SoftwareHyperflexDistributableRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *SoftwareHyperflexDistributableRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *SoftwareHyperflexDistributableRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *SoftwareHyperflexDistributableRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDescription

`func (o *SoftwareHyperflexDistributableRelationship) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SoftwareHyperflexDistributableRelationship) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SoftwareHyperflexDistributableRelationship) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SoftwareHyperflexDistributableRelationship) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDownloadCount

`func (o *SoftwareHyperflexDistributableRelationship) GetDownloadCount() int64`

GetDownloadCount returns the DownloadCount field if non-nil, zero value otherwise.

### GetDownloadCountOk

`func (o *SoftwareHyperflexDistributableRelationship) GetDownloadCountOk() (*int64, bool)`

GetDownloadCountOk returns a tuple with the DownloadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadCount

`func (o *SoftwareHyperflexDistributableRelationship) SetDownloadCount(v int64)`

SetDownloadCount sets DownloadCount field to given value.

### HasDownloadCount

`func (o *SoftwareHyperflexDistributableRelationship) HasDownloadCount() bool`

HasDownloadCount returns a boolean if a field has been set.

### GetImportAction

`func (o *SoftwareHyperflexDistributableRelationship) GetImportAction() string`

GetImportAction returns the ImportAction field if non-nil, zero value otherwise.

### GetImportActionOk

`func (o *SoftwareHyperflexDistributableRelationship) GetImportActionOk() (*string, bool)`

GetImportActionOk returns a tuple with the ImportAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportAction

`func (o *SoftwareHyperflexDistributableRelationship) SetImportAction(v string)`

SetImportAction sets ImportAction field to given value.

### HasImportAction

`func (o *SoftwareHyperflexDistributableRelationship) HasImportAction() bool`

HasImportAction returns a boolean if a field has been set.

### GetImportState

`func (o *SoftwareHyperflexDistributableRelationship) GetImportState() string`

GetImportState returns the ImportState field if non-nil, zero value otherwise.

### GetImportStateOk

`func (o *SoftwareHyperflexDistributableRelationship) GetImportStateOk() (*string, bool)`

GetImportStateOk returns a tuple with the ImportState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportState

`func (o *SoftwareHyperflexDistributableRelationship) SetImportState(v string)`

SetImportState sets ImportState field to given value.

### HasImportState

`func (o *SoftwareHyperflexDistributableRelationship) HasImportState() bool`

HasImportState returns a boolean if a field has been set.

### GetImportedTime

`func (o *SoftwareHyperflexDistributableRelationship) GetImportedTime() time.Time`

GetImportedTime returns the ImportedTime field if non-nil, zero value otherwise.

### GetImportedTimeOk

`func (o *SoftwareHyperflexDistributableRelationship) GetImportedTimeOk() (*time.Time, bool)`

GetImportedTimeOk returns a tuple with the ImportedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedTime

`func (o *SoftwareHyperflexDistributableRelationship) SetImportedTime(v time.Time)`

SetImportedTime sets ImportedTime field to given value.

### HasImportedTime

`func (o *SoftwareHyperflexDistributableRelationship) HasImportedTime() bool`

HasImportedTime returns a boolean if a field has been set.

### GetLastAccessTime

`func (o *SoftwareHyperflexDistributableRelationship) GetLastAccessTime() time.Time`

GetLastAccessTime returns the LastAccessTime field if non-nil, zero value otherwise.

### GetLastAccessTimeOk

`func (o *SoftwareHyperflexDistributableRelationship) GetLastAccessTimeOk() (*time.Time, bool)`

GetLastAccessTimeOk returns a tuple with the LastAccessTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessTime

`func (o *SoftwareHyperflexDistributableRelationship) SetLastAccessTime(v time.Time)`

SetLastAccessTime sets LastAccessTime field to given value.

### HasLastAccessTime

`func (o *SoftwareHyperflexDistributableRelationship) HasLastAccessTime() bool`

HasLastAccessTime returns a boolean if a field has been set.

### GetMd5sum

`func (o *SoftwareHyperflexDistributableRelationship) GetMd5sum() string`

GetMd5sum returns the Md5sum field if non-nil, zero value otherwise.

### GetMd5sumOk

`func (o *SoftwareHyperflexDistributableRelationship) GetMd5sumOk() (*string, bool)`

GetMd5sumOk returns a tuple with the Md5sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5sum

`func (o *SoftwareHyperflexDistributableRelationship) SetMd5sum(v string)`

SetMd5sum sets Md5sum field to given value.

### HasMd5sum

`func (o *SoftwareHyperflexDistributableRelationship) HasMd5sum() bool`

HasMd5sum returns a boolean if a field has been set.

### GetName

`func (o *SoftwareHyperflexDistributableRelationship) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SoftwareHyperflexDistributableRelationship) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SoftwareHyperflexDistributableRelationship) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SoftwareHyperflexDistributableRelationship) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReleaseDate

`func (o *SoftwareHyperflexDistributableRelationship) GetReleaseDate() time.Time`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *SoftwareHyperflexDistributableRelationship) GetReleaseDateOk() (*time.Time, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *SoftwareHyperflexDistributableRelationship) SetReleaseDate(v time.Time)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *SoftwareHyperflexDistributableRelationship) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### GetSha512sum

`func (o *SoftwareHyperflexDistributableRelationship) GetSha512sum() string`

GetSha512sum returns the Sha512sum field if non-nil, zero value otherwise.

### GetSha512sumOk

`func (o *SoftwareHyperflexDistributableRelationship) GetSha512sumOk() (*string, bool)`

GetSha512sumOk returns a tuple with the Sha512sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha512sum

`func (o *SoftwareHyperflexDistributableRelationship) SetSha512sum(v string)`

SetSha512sum sets Sha512sum field to given value.

### HasSha512sum

`func (o *SoftwareHyperflexDistributableRelationship) HasSha512sum() bool`

HasSha512sum returns a boolean if a field has been set.

### GetSize

`func (o *SoftwareHyperflexDistributableRelationship) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *SoftwareHyperflexDistributableRelationship) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *SoftwareHyperflexDistributableRelationship) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *SoftwareHyperflexDistributableRelationship) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSoftwareAdvisoryUrl

`func (o *SoftwareHyperflexDistributableRelationship) GetSoftwareAdvisoryUrl() string`

GetSoftwareAdvisoryUrl returns the SoftwareAdvisoryUrl field if non-nil, zero value otherwise.

### GetSoftwareAdvisoryUrlOk

`func (o *SoftwareHyperflexDistributableRelationship) GetSoftwareAdvisoryUrlOk() (*string, bool)`

GetSoftwareAdvisoryUrlOk returns a tuple with the SoftwareAdvisoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareAdvisoryUrl

`func (o *SoftwareHyperflexDistributableRelationship) SetSoftwareAdvisoryUrl(v string)`

SetSoftwareAdvisoryUrl sets SoftwareAdvisoryUrl field to given value.

### HasSoftwareAdvisoryUrl

`func (o *SoftwareHyperflexDistributableRelationship) HasSoftwareAdvisoryUrl() bool`

HasSoftwareAdvisoryUrl returns a boolean if a field has been set.

### GetSource

`func (o *SoftwareHyperflexDistributableRelationship) GetSource() SoftwarerepositoryFileServer`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SoftwareHyperflexDistributableRelationship) GetSourceOk() (*SoftwarerepositoryFileServer, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SoftwareHyperflexDistributableRelationship) SetSource(v SoftwarerepositoryFileServer)`

SetSource sets Source field to given value.

### HasSource

`func (o *SoftwareHyperflexDistributableRelationship) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetVersion

`func (o *SoftwareHyperflexDistributableRelationship) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SoftwareHyperflexDistributableRelationship) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SoftwareHyperflexDistributableRelationship) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SoftwareHyperflexDistributableRelationship) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetBundleType

`func (o *SoftwareHyperflexDistributableRelationship) GetBundleType() string`

GetBundleType returns the BundleType field if non-nil, zero value otherwise.

### GetBundleTypeOk

`func (o *SoftwareHyperflexDistributableRelationship) GetBundleTypeOk() (*string, bool)`

GetBundleTypeOk returns a tuple with the BundleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleType

`func (o *SoftwareHyperflexDistributableRelationship) SetBundleType(v string)`

SetBundleType sets BundleType field to given value.

### HasBundleType

`func (o *SoftwareHyperflexDistributableRelationship) HasBundleType() bool`

HasBundleType returns a boolean if a field has been set.

### GetComponentMeta

`func (o *SoftwareHyperflexDistributableRelationship) GetComponentMeta() []FirmwareComponentMeta`

GetComponentMeta returns the ComponentMeta field if non-nil, zero value otherwise.

### GetComponentMetaOk

`func (o *SoftwareHyperflexDistributableRelationship) GetComponentMetaOk() (*[]FirmwareComponentMeta, bool)`

GetComponentMetaOk returns a tuple with the ComponentMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentMeta

`func (o *SoftwareHyperflexDistributableRelationship) SetComponentMeta(v []FirmwareComponentMeta)`

SetComponentMeta sets ComponentMeta field to given value.

### HasComponentMeta

`func (o *SoftwareHyperflexDistributableRelationship) HasComponentMeta() bool`

HasComponentMeta returns a boolean if a field has been set.

### GetGuid

`func (o *SoftwareHyperflexDistributableRelationship) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *SoftwareHyperflexDistributableRelationship) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *SoftwareHyperflexDistributableRelationship) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *SoftwareHyperflexDistributableRelationship) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetMdfid

`func (o *SoftwareHyperflexDistributableRelationship) GetMdfid() string`

GetMdfid returns the Mdfid field if non-nil, zero value otherwise.

### GetMdfidOk

`func (o *SoftwareHyperflexDistributableRelationship) GetMdfidOk() (*string, bool)`

GetMdfidOk returns a tuple with the Mdfid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdfid

`func (o *SoftwareHyperflexDistributableRelationship) SetMdfid(v string)`

SetMdfid sets Mdfid field to given value.

### HasMdfid

`func (o *SoftwareHyperflexDistributableRelationship) HasMdfid() bool`

HasMdfid returns a boolean if a field has been set.

### GetModel

`func (o *SoftwareHyperflexDistributableRelationship) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *SoftwareHyperflexDistributableRelationship) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *SoftwareHyperflexDistributableRelationship) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *SoftwareHyperflexDistributableRelationship) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetPlatformType

`func (o *SoftwareHyperflexDistributableRelationship) GetPlatformType() string`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *SoftwareHyperflexDistributableRelationship) GetPlatformTypeOk() (*string, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *SoftwareHyperflexDistributableRelationship) SetPlatformType(v string)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *SoftwareHyperflexDistributableRelationship) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### GetRecommendedBuild

`func (o *SoftwareHyperflexDistributableRelationship) GetRecommendedBuild() string`

GetRecommendedBuild returns the RecommendedBuild field if non-nil, zero value otherwise.

### GetRecommendedBuildOk

`func (o *SoftwareHyperflexDistributableRelationship) GetRecommendedBuildOk() (*string, bool)`

GetRecommendedBuildOk returns a tuple with the RecommendedBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendedBuild

`func (o *SoftwareHyperflexDistributableRelationship) SetRecommendedBuild(v string)`

SetRecommendedBuild sets RecommendedBuild field to given value.

### HasRecommendedBuild

`func (o *SoftwareHyperflexDistributableRelationship) HasRecommendedBuild() bool`

HasRecommendedBuild returns a boolean if a field has been set.

### GetReleaseNotesUrl

`func (o *SoftwareHyperflexDistributableRelationship) GetReleaseNotesUrl() string`

GetReleaseNotesUrl returns the ReleaseNotesUrl field if non-nil, zero value otherwise.

### GetReleaseNotesUrlOk

`func (o *SoftwareHyperflexDistributableRelationship) GetReleaseNotesUrlOk() (*string, bool)`

GetReleaseNotesUrlOk returns a tuple with the ReleaseNotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseNotesUrl

`func (o *SoftwareHyperflexDistributableRelationship) SetReleaseNotesUrl(v string)`

SetReleaseNotesUrl sets ReleaseNotesUrl field to given value.

### HasReleaseNotesUrl

`func (o *SoftwareHyperflexDistributableRelationship) HasReleaseNotesUrl() bool`

HasReleaseNotesUrl returns a boolean if a field has been set.

### GetSoftwareTypeId

`func (o *SoftwareHyperflexDistributableRelationship) GetSoftwareTypeId() string`

GetSoftwareTypeId returns the SoftwareTypeId field if non-nil, zero value otherwise.

### GetSoftwareTypeIdOk

`func (o *SoftwareHyperflexDistributableRelationship) GetSoftwareTypeIdOk() (*string, bool)`

GetSoftwareTypeIdOk returns a tuple with the SoftwareTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareTypeId

`func (o *SoftwareHyperflexDistributableRelationship) SetSoftwareTypeId(v string)`

SetSoftwareTypeId sets SoftwareTypeId field to given value.

### HasSoftwareTypeId

`func (o *SoftwareHyperflexDistributableRelationship) HasSoftwareTypeId() bool`

HasSoftwareTypeId returns a boolean if a field has been set.

### GetSupportedModels

`func (o *SoftwareHyperflexDistributableRelationship) GetSupportedModels() []string`

GetSupportedModels returns the SupportedModels field if non-nil, zero value otherwise.

### GetSupportedModelsOk

`func (o *SoftwareHyperflexDistributableRelationship) GetSupportedModelsOk() (*[]string, bool)`

GetSupportedModelsOk returns a tuple with the SupportedModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedModels

`func (o *SoftwareHyperflexDistributableRelationship) SetSupportedModels(v []string)`

SetSupportedModels sets SupportedModels field to given value.

### HasSupportedModels

`func (o *SoftwareHyperflexDistributableRelationship) HasSupportedModels() bool`

HasSupportedModels returns a boolean if a field has been set.

### GetVendor

`func (o *SoftwareHyperflexDistributableRelationship) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *SoftwareHyperflexDistributableRelationship) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *SoftwareHyperflexDistributableRelationship) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *SoftwareHyperflexDistributableRelationship) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetDistributableMetas

`func (o *SoftwareHyperflexDistributableRelationship) GetDistributableMetas() []FirmwareDistributableMetaRelationship`

GetDistributableMetas returns the DistributableMetas field if non-nil, zero value otherwise.

### GetDistributableMetasOk

`func (o *SoftwareHyperflexDistributableRelationship) GetDistributableMetasOk() (*[]FirmwareDistributableMetaRelationship, bool)`

GetDistributableMetasOk returns a tuple with the DistributableMetas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributableMetas

`func (o *SoftwareHyperflexDistributableRelationship) SetDistributableMetas(v []FirmwareDistributableMetaRelationship)`

SetDistributableMetas sets DistributableMetas field to given value.

### HasDistributableMetas

`func (o *SoftwareHyperflexDistributableRelationship) HasDistributableMetas() bool`

HasDistributableMetas returns a boolean if a field has been set.

### SetDistributableMetasNil

`func (o *SoftwareHyperflexDistributableRelationship) SetDistributableMetasNil(b bool)`

 SetDistributableMetasNil sets the value for DistributableMetas to be an explicit nil

### UnsetDistributableMetas
`func (o *SoftwareHyperflexDistributableRelationship) UnsetDistributableMetas()`

UnsetDistributableMetas ensures that no value is present for DistributableMetas, not even an explicit nil
### GetRelease

`func (o *SoftwareHyperflexDistributableRelationship) GetRelease() SoftwarerepositoryReleaseRelationship`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *SoftwareHyperflexDistributableRelationship) GetReleaseOk() (*SoftwarerepositoryReleaseRelationship, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *SoftwareHyperflexDistributableRelationship) SetRelease(v SoftwarerepositoryReleaseRelationship)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *SoftwareHyperflexDistributableRelationship) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetCatalog

`func (o *SoftwareHyperflexDistributableRelationship) GetCatalog() SoftwarerepositoryCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *SoftwareHyperflexDistributableRelationship) GetCatalogOk() (*SoftwarerepositoryCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *SoftwareHyperflexDistributableRelationship) SetCatalog(v SoftwarerepositoryCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *SoftwareHyperflexDistributableRelationship) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


