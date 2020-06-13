# AssetServiceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Options** | Pointer to [**AssetServiceOptions**](asset.ServiceOptions.md) |  | [optional] 
**Status** | Pointer to **string** | Status indicates if the respective Service can establish a connection and authenticate with the managed target. Status does not include information about the functional health of the target. | [optional] [default to ""]
**StatusErrorReason** | Pointer to **string** | When &#39;Status&#39; is not Connected, statusErrorReason provides further details about why the device is not connected with Intersight. | [optional] 

## Methods

### NewAssetServiceAllOf

`func NewAssetServiceAllOf() *AssetServiceAllOf`

NewAssetServiceAllOf instantiates a new AssetServiceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetServiceAllOfWithDefaults

`func NewAssetServiceAllOfWithDefaults() *AssetServiceAllOf`

NewAssetServiceAllOfWithDefaults instantiates a new AssetServiceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptions

`func (o *AssetServiceAllOf) GetOptions() AssetServiceOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *AssetServiceAllOf) GetOptionsOk() (*AssetServiceOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *AssetServiceAllOf) SetOptions(v AssetServiceOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *AssetServiceAllOf) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetStatus

`func (o *AssetServiceAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AssetServiceAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AssetServiceAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AssetServiceAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusErrorReason

`func (o *AssetServiceAllOf) GetStatusErrorReason() string`

GetStatusErrorReason returns the StatusErrorReason field if non-nil, zero value otherwise.

### GetStatusErrorReasonOk

`func (o *AssetServiceAllOf) GetStatusErrorReasonOk() (*string, bool)`

GetStatusErrorReasonOk returns a tuple with the StatusErrorReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusErrorReason

`func (o *AssetServiceAllOf) SetStatusErrorReason(v string)`

SetStatusErrorReason sets StatusErrorReason field to given value.

### HasStatusErrorReason

`func (o *AssetServiceAllOf) HasStatusErrorReason() bool`

HasStatusErrorReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


