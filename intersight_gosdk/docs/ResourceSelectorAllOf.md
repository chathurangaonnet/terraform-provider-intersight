# ResourceSelectorAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Selector** | Pointer to **string** | ODATA filter to select resources. The group selector may include URLs of individual resource, or OData query with filters that match multiple queries. The URLs must be relative (i.e. do not include the host). | [optional] 

## Methods

### NewResourceSelectorAllOf

`func NewResourceSelectorAllOf() *ResourceSelectorAllOf`

NewResourceSelectorAllOf instantiates a new ResourceSelectorAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceSelectorAllOfWithDefaults

`func NewResourceSelectorAllOfWithDefaults() *ResourceSelectorAllOf`

NewResourceSelectorAllOfWithDefaults instantiates a new ResourceSelectorAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelector

`func (o *ResourceSelectorAllOf) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *ResourceSelectorAllOf) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *ResourceSelectorAllOf) SetSelector(v string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *ResourceSelectorAllOf) HasSelector() bool`

HasSelector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


