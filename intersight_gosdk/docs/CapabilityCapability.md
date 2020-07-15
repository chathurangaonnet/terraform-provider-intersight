# CapabilityCapability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | An unique identifer for a capability descriptor. | [optional] 
**Section** | Pointer to [**CapabilitySectionRelationship**](capability.Section.Relationship.md) |  | [optional] 

## Methods

### NewCapabilityCapability

`func NewCapabilityCapability() *CapabilityCapability`

NewCapabilityCapability instantiates a new CapabilityCapability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityCapabilityWithDefaults

`func NewCapabilityCapabilityWithDefaults() *CapabilityCapability`

NewCapabilityCapabilityWithDefaults instantiates a new CapabilityCapability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CapabilityCapability) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CapabilityCapability) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CapabilityCapability) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CapabilityCapability) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSection

`func (o *CapabilityCapability) GetSection() CapabilitySectionRelationship`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *CapabilityCapability) GetSectionOk() (*CapabilitySectionRelationship, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *CapabilityCapability) SetSection(v CapabilitySectionRelationship)`

SetSection sets Section field to given value.

### HasSection

`func (o *CapabilityCapability) HasSection() bool`

HasSection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


