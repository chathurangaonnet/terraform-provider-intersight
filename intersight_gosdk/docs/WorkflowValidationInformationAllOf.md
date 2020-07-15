# WorkflowValidationInformationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | The current validation state of this workflow. The possible states are Valid, Invalid, NotValidated (default). | [optional] [readonly] [default to "NotValidated"]
**ValidationError** | Pointer to [**[]WorkflowValidationError**](workflow.ValidationError.md) |  | [optional] 

## Methods

### NewWorkflowValidationInformationAllOf

`func NewWorkflowValidationInformationAllOf() *WorkflowValidationInformationAllOf`

NewWorkflowValidationInformationAllOf instantiates a new WorkflowValidationInformationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowValidationInformationAllOfWithDefaults

`func NewWorkflowValidationInformationAllOfWithDefaults() *WorkflowValidationInformationAllOf`

NewWorkflowValidationInformationAllOfWithDefaults instantiates a new WorkflowValidationInformationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *WorkflowValidationInformationAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *WorkflowValidationInformationAllOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *WorkflowValidationInformationAllOf) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *WorkflowValidationInformationAllOf) HasState() bool`

HasState returns a boolean if a field has been set.

### GetValidationError

`func (o *WorkflowValidationInformationAllOf) GetValidationError() []WorkflowValidationError`

GetValidationError returns the ValidationError field if non-nil, zero value otherwise.

### GetValidationErrorOk

`func (o *WorkflowValidationInformationAllOf) GetValidationErrorOk() (*[]WorkflowValidationError, bool)`

GetValidationErrorOk returns a tuple with the ValidationError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationError

`func (o *WorkflowValidationInformationAllOf) SetValidationError(v []WorkflowValidationError)`

SetValidationError sets ValidationError field to given value.

### HasValidationError

`func (o *WorkflowValidationInformationAllOf) HasValidationError() bool`

HasValidationError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


