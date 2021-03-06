# TechsupportmanagementNiaParamAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionLevel** | Pointer to **int32** | CollectionLevel controls the size / depth of the tech support files collected. | [optional] [default to 1]
**Filename** | Pointer to **string** | Filename specifies an individual filename to collect from the endpoint. | [optional] 
**ForceFresh** | Pointer to **bool** | ForceFresh controls whether to return pre-collected files or force the collection of new files. | [optional] 
**Pids** | Pointer to **[]string** |  | [optional] 
**SerialNumbers** | Pointer to **[]string** |  | [optional] 

## Methods

### NewTechsupportmanagementNiaParamAllOf

`func NewTechsupportmanagementNiaParamAllOf() *TechsupportmanagementNiaParamAllOf`

NewTechsupportmanagementNiaParamAllOf instantiates a new TechsupportmanagementNiaParamAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTechsupportmanagementNiaParamAllOfWithDefaults

`func NewTechsupportmanagementNiaParamAllOfWithDefaults() *TechsupportmanagementNiaParamAllOf`

NewTechsupportmanagementNiaParamAllOfWithDefaults instantiates a new TechsupportmanagementNiaParamAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionLevel

`func (o *TechsupportmanagementNiaParamAllOf) GetCollectionLevel() int32`

GetCollectionLevel returns the CollectionLevel field if non-nil, zero value otherwise.

### GetCollectionLevelOk

`func (o *TechsupportmanagementNiaParamAllOf) GetCollectionLevelOk() (*int32, bool)`

GetCollectionLevelOk returns a tuple with the CollectionLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionLevel

`func (o *TechsupportmanagementNiaParamAllOf) SetCollectionLevel(v int32)`

SetCollectionLevel sets CollectionLevel field to given value.

### HasCollectionLevel

`func (o *TechsupportmanagementNiaParamAllOf) HasCollectionLevel() bool`

HasCollectionLevel returns a boolean if a field has been set.

### GetFilename

`func (o *TechsupportmanagementNiaParamAllOf) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *TechsupportmanagementNiaParamAllOf) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *TechsupportmanagementNiaParamAllOf) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *TechsupportmanagementNiaParamAllOf) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetForceFresh

`func (o *TechsupportmanagementNiaParamAllOf) GetForceFresh() bool`

GetForceFresh returns the ForceFresh field if non-nil, zero value otherwise.

### GetForceFreshOk

`func (o *TechsupportmanagementNiaParamAllOf) GetForceFreshOk() (*bool, bool)`

GetForceFreshOk returns a tuple with the ForceFresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceFresh

`func (o *TechsupportmanagementNiaParamAllOf) SetForceFresh(v bool)`

SetForceFresh sets ForceFresh field to given value.

### HasForceFresh

`func (o *TechsupportmanagementNiaParamAllOf) HasForceFresh() bool`

HasForceFresh returns a boolean if a field has been set.

### GetPids

`func (o *TechsupportmanagementNiaParamAllOf) GetPids() []string`

GetPids returns the Pids field if non-nil, zero value otherwise.

### GetPidsOk

`func (o *TechsupportmanagementNiaParamAllOf) GetPidsOk() (*[]string, bool)`

GetPidsOk returns a tuple with the Pids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPids

`func (o *TechsupportmanagementNiaParamAllOf) SetPids(v []string)`

SetPids sets Pids field to given value.

### HasPids

`func (o *TechsupportmanagementNiaParamAllOf) HasPids() bool`

HasPids returns a boolean if a field has been set.

### GetSerialNumbers

`func (o *TechsupportmanagementNiaParamAllOf) GetSerialNumbers() []string`

GetSerialNumbers returns the SerialNumbers field if non-nil, zero value otherwise.

### GetSerialNumbersOk

`func (o *TechsupportmanagementNiaParamAllOf) GetSerialNumbersOk() (*[]string, bool)`

GetSerialNumbersOk returns a tuple with the SerialNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumbers

`func (o *TechsupportmanagementNiaParamAllOf) SetSerialNumbers(v []string)`

SetSerialNumbers sets SerialNumbers field to given value.

### HasSerialNumbers

`func (o *TechsupportmanagementNiaParamAllOf) HasSerialNumbers() bool`

HasSerialNumbers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


