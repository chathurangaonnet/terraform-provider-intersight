# TamBaseAdvisory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Brief description of the advisory details. | [optional] 
**Name** | Pointer to **string** | A user defined name for the Intersight Advisory. | [optional] 
**Severity** | Pointer to [**TamSeverity**](tam.Severity.md) |  | [optional] 
**State** | Pointer to **string** | Current state of the advisory. | [optional] [default to "ready"]

## Methods

### NewTamBaseAdvisory

`func NewTamBaseAdvisory() *TamBaseAdvisory`

NewTamBaseAdvisory instantiates a new TamBaseAdvisory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTamBaseAdvisoryWithDefaults

`func NewTamBaseAdvisoryWithDefaults() *TamBaseAdvisory`

NewTamBaseAdvisoryWithDefaults instantiates a new TamBaseAdvisory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *TamBaseAdvisory) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TamBaseAdvisory) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TamBaseAdvisory) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TamBaseAdvisory) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *TamBaseAdvisory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TamBaseAdvisory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TamBaseAdvisory) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TamBaseAdvisory) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSeverity

`func (o *TamBaseAdvisory) GetSeverity() TamSeverity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *TamBaseAdvisory) GetSeverityOk() (*TamSeverity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *TamBaseAdvisory) SetSeverity(v TamSeverity)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *TamBaseAdvisory) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetState

`func (o *TamBaseAdvisory) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TamBaseAdvisory) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TamBaseAdvisory) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *TamBaseAdvisory) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


