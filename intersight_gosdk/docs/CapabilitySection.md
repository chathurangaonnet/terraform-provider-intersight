# CapabilitySection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Administrative action to initialize/load the catalog section from a particular source. | [optional] [readonly] [default to "None"]
**CatalogName** | Pointer to **string** | The catalog name reference. | [optional] 
**Name** | Pointer to **string** | A unique name for the section inside a catalog. | [optional] 
**Source** | Pointer to **string** | The configured source for this section of the catalog. | [optional] [readonly] [default to "Local"]
**Version** | Pointer to **string** | Version of the section inside a catalog. | [optional] 

## Methods

### NewCapabilitySection

`func NewCapabilitySection() *CapabilitySection`

NewCapabilitySection instantiates a new CapabilitySection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilitySectionWithDefaults

`func NewCapabilitySectionWithDefaults() *CapabilitySection`

NewCapabilitySectionWithDefaults instantiates a new CapabilitySection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *CapabilitySection) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CapabilitySection) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CapabilitySection) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *CapabilitySection) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCatalogName

`func (o *CapabilitySection) GetCatalogName() string`

GetCatalogName returns the CatalogName field if non-nil, zero value otherwise.

### GetCatalogNameOk

`func (o *CapabilitySection) GetCatalogNameOk() (*string, bool)`

GetCatalogNameOk returns a tuple with the CatalogName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogName

`func (o *CapabilitySection) SetCatalogName(v string)`

SetCatalogName sets CatalogName field to given value.

### HasCatalogName

`func (o *CapabilitySection) HasCatalogName() bool`

HasCatalogName returns a boolean if a field has been set.

### GetName

`func (o *CapabilitySection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CapabilitySection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CapabilitySection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CapabilitySection) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSource

`func (o *CapabilitySection) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CapabilitySection) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CapabilitySection) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CapabilitySection) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetVersion

`func (o *CapabilitySection) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CapabilitySection) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CapabilitySection) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CapabilitySection) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


