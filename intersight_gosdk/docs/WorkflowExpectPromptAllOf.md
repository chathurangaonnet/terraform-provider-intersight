# WorkflowExpectPromptAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expect** | Pointer to **string** | The regex of the expect prompt of the interactive command. | [optional] 
**Send** | Pointer to **string** | The answer string to the expect prompt. | [optional] 

## Methods

### NewWorkflowExpectPromptAllOf

`func NewWorkflowExpectPromptAllOf() *WorkflowExpectPromptAllOf`

NewWorkflowExpectPromptAllOf instantiates a new WorkflowExpectPromptAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowExpectPromptAllOfWithDefaults

`func NewWorkflowExpectPromptAllOfWithDefaults() *WorkflowExpectPromptAllOf`

NewWorkflowExpectPromptAllOfWithDefaults instantiates a new WorkflowExpectPromptAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpect

`func (o *WorkflowExpectPromptAllOf) GetExpect() string`

GetExpect returns the Expect field if non-nil, zero value otherwise.

### GetExpectOk

`func (o *WorkflowExpectPromptAllOf) GetExpectOk() (*string, bool)`

GetExpectOk returns a tuple with the Expect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpect

`func (o *WorkflowExpectPromptAllOf) SetExpect(v string)`

SetExpect sets Expect field to given value.

### HasExpect

`func (o *WorkflowExpectPromptAllOf) HasExpect() bool`

HasExpect returns a boolean if a field has been set.

### GetSend

`func (o *WorkflowExpectPromptAllOf) GetSend() string`

GetSend returns the Send field if non-nil, zero value otherwise.

### GetSendOk

`func (o *WorkflowExpectPromptAllOf) GetSendOk() (*string, bool)`

GetSendOk returns a tuple with the Send field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSend

`func (o *WorkflowExpectPromptAllOf) SetSend(v string)`

SetSend sets Send field to given value.

### HasSend

`func (o *WorkflowExpectPromptAllOf) HasSend() bool`

HasSend returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


