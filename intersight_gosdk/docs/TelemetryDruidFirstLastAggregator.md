# TelemetryDruidFirstLastAggregator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The aggregator type. | 
**Name** | Pointer to **string** | Output name for the first/last value. | 
**FieldName** | Pointer to **string** | Name of the metric column. | 

## Methods

### NewTelemetryDruidFirstLastAggregator

`func NewTelemetryDruidFirstLastAggregator(type_ string, name string, fieldName string, ) *TelemetryDruidFirstLastAggregator`

NewTelemetryDruidFirstLastAggregator instantiates a new TelemetryDruidFirstLastAggregator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidFirstLastAggregatorWithDefaults

`func NewTelemetryDruidFirstLastAggregatorWithDefaults() *TelemetryDruidFirstLastAggregator`

NewTelemetryDruidFirstLastAggregatorWithDefaults instantiates a new TelemetryDruidFirstLastAggregator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidFirstLastAggregator) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidFirstLastAggregator) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidFirstLastAggregator) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *TelemetryDruidFirstLastAggregator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TelemetryDruidFirstLastAggregator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TelemetryDruidFirstLastAggregator) SetName(v string)`

SetName sets Name field to given value.


### GetFieldName

`func (o *TelemetryDruidFirstLastAggregator) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *TelemetryDruidFirstLastAggregator) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *TelemetryDruidFirstLastAggregator) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

