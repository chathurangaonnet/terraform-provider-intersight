# HclServiceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExemptionFileVersion** | Pointer to **string** | Version of the last modified exemption file. | [optional] 
**Identity** | Pointer to **string** | A field to uniquely identify the document with the status. | [optional] 
**LastHclDataModifiedTime** | Pointer to [**time.Time**](time.Time.md) | The timestamp of the last modified record in the HCL tool database. Used to query and get updated records. | [optional] 
**LastImportedDataChecksum** | Pointer to **string** | Checksum of the data dump used as the base for delta updates. | [optional] 
**Status** | Pointer to **string** | Status of the service indicatating if the service is up or under maintenance due to data update. | [optional] [default to "Unknown"]

## Methods

### NewHclServiceStatus

`func NewHclServiceStatus() *HclServiceStatus`

NewHclServiceStatus instantiates a new HclServiceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHclServiceStatusWithDefaults

`func NewHclServiceStatusWithDefaults() *HclServiceStatus`

NewHclServiceStatusWithDefaults instantiates a new HclServiceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExemptionFileVersion

`func (o *HclServiceStatus) GetExemptionFileVersion() string`

GetExemptionFileVersion returns the ExemptionFileVersion field if non-nil, zero value otherwise.

### GetExemptionFileVersionOk

`func (o *HclServiceStatus) GetExemptionFileVersionOk() (*string, bool)`

GetExemptionFileVersionOk returns a tuple with the ExemptionFileVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptionFileVersion

`func (o *HclServiceStatus) SetExemptionFileVersion(v string)`

SetExemptionFileVersion sets ExemptionFileVersion field to given value.

### HasExemptionFileVersion

`func (o *HclServiceStatus) HasExemptionFileVersion() bool`

HasExemptionFileVersion returns a boolean if a field has been set.

### GetIdentity

`func (o *HclServiceStatus) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *HclServiceStatus) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *HclServiceStatus) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *HclServiceStatus) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetLastHclDataModifiedTime

`func (o *HclServiceStatus) GetLastHclDataModifiedTime() time.Time`

GetLastHclDataModifiedTime returns the LastHclDataModifiedTime field if non-nil, zero value otherwise.

### GetLastHclDataModifiedTimeOk

`func (o *HclServiceStatus) GetLastHclDataModifiedTimeOk() (*time.Time, bool)`

GetLastHclDataModifiedTimeOk returns a tuple with the LastHclDataModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHclDataModifiedTime

`func (o *HclServiceStatus) SetLastHclDataModifiedTime(v time.Time)`

SetLastHclDataModifiedTime sets LastHclDataModifiedTime field to given value.

### HasLastHclDataModifiedTime

`func (o *HclServiceStatus) HasLastHclDataModifiedTime() bool`

HasLastHclDataModifiedTime returns a boolean if a field has been set.

### GetLastImportedDataChecksum

`func (o *HclServiceStatus) GetLastImportedDataChecksum() string`

GetLastImportedDataChecksum returns the LastImportedDataChecksum field if non-nil, zero value otherwise.

### GetLastImportedDataChecksumOk

`func (o *HclServiceStatus) GetLastImportedDataChecksumOk() (*string, bool)`

GetLastImportedDataChecksumOk returns a tuple with the LastImportedDataChecksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastImportedDataChecksum

`func (o *HclServiceStatus) SetLastImportedDataChecksum(v string)`

SetLastImportedDataChecksum sets LastImportedDataChecksum field to given value.

### HasLastImportedDataChecksum

`func (o *HclServiceStatus) HasLastImportedDataChecksum() bool`

HasLastImportedDataChecksum returns a boolean if a field has been set.

### GetStatus

`func (o *HclServiceStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HclServiceStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HclServiceStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HclServiceStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


