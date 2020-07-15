# VnicCdnAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | Pointer to **string** | Source of the CDN. It can either be user specified or be the same as the vNIC name. | [optional] [default to "vnic"]
**Value** | Pointer to **string** | The CDN value entered in case of user defined mode. | [optional] 

## Methods

### NewVnicCdnAllOf

`func NewVnicCdnAllOf() *VnicCdnAllOf`

NewVnicCdnAllOf instantiates a new VnicCdnAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicCdnAllOfWithDefaults

`func NewVnicCdnAllOfWithDefaults() *VnicCdnAllOf`

NewVnicCdnAllOfWithDefaults instantiates a new VnicCdnAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *VnicCdnAllOf) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *VnicCdnAllOf) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *VnicCdnAllOf) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *VnicCdnAllOf) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetValue

`func (o *VnicCdnAllOf) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *VnicCdnAllOf) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *VnicCdnAllOf) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *VnicCdnAllOf) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


