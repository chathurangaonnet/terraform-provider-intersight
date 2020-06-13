# WorkflowCliCommandAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | Pointer to **string** | The command to run on the device connector. | [optional] 
**EndPrompt** | Pointer to **string** | The regex string that identifies the end of the command response. | [optional] 
**ExpectPrompts** | Pointer to [**[]WorkflowExpectPrompt**](workflow.ExpectPrompt.md) |  | [optional] 
**SkipStatusCheck** | Pointer to **bool** | Skips the execution status check of the terminal command. One use case for this is while exiting the terminal session from esxi host. | [optional] 
**TerminalEnd** | Pointer to **bool** | If this flag is set, it marks the end of the terminal session where the previous commands were executed. | [optional] 
**TerminalStart** | Pointer to **bool** | If this flag is set, the execution of this command initiates a terminal session in which the subsequent CLI commands are executed until a command with terminalEnd flag is encountered or the end of the batch. | [optional] 
**Type** | Pointer to **string** | The type of the command - can be interactive or non-interactive. | [optional] [default to "NonInteractive"]

## Methods

### NewWorkflowCliCommandAllOf

`func NewWorkflowCliCommandAllOf() *WorkflowCliCommandAllOf`

NewWorkflowCliCommandAllOf instantiates a new WorkflowCliCommandAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowCliCommandAllOfWithDefaults

`func NewWorkflowCliCommandAllOfWithDefaults() *WorkflowCliCommandAllOf`

NewWorkflowCliCommandAllOfWithDefaults instantiates a new WorkflowCliCommandAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *WorkflowCliCommandAllOf) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *WorkflowCliCommandAllOf) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *WorkflowCliCommandAllOf) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *WorkflowCliCommandAllOf) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetEndPrompt

`func (o *WorkflowCliCommandAllOf) GetEndPrompt() string`

GetEndPrompt returns the EndPrompt field if non-nil, zero value otherwise.

### GetEndPromptOk

`func (o *WorkflowCliCommandAllOf) GetEndPromptOk() (*string, bool)`

GetEndPromptOk returns a tuple with the EndPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPrompt

`func (o *WorkflowCliCommandAllOf) SetEndPrompt(v string)`

SetEndPrompt sets EndPrompt field to given value.

### HasEndPrompt

`func (o *WorkflowCliCommandAllOf) HasEndPrompt() bool`

HasEndPrompt returns a boolean if a field has been set.

### GetExpectPrompts

`func (o *WorkflowCliCommandAllOf) GetExpectPrompts() []WorkflowExpectPrompt`

GetExpectPrompts returns the ExpectPrompts field if non-nil, zero value otherwise.

### GetExpectPromptsOk

`func (o *WorkflowCliCommandAllOf) GetExpectPromptsOk() (*[]WorkflowExpectPrompt, bool)`

GetExpectPromptsOk returns a tuple with the ExpectPrompts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectPrompts

`func (o *WorkflowCliCommandAllOf) SetExpectPrompts(v []WorkflowExpectPrompt)`

SetExpectPrompts sets ExpectPrompts field to given value.

### HasExpectPrompts

`func (o *WorkflowCliCommandAllOf) HasExpectPrompts() bool`

HasExpectPrompts returns a boolean if a field has been set.

### GetSkipStatusCheck

`func (o *WorkflowCliCommandAllOf) GetSkipStatusCheck() bool`

GetSkipStatusCheck returns the SkipStatusCheck field if non-nil, zero value otherwise.

### GetSkipStatusCheckOk

`func (o *WorkflowCliCommandAllOf) GetSkipStatusCheckOk() (*bool, bool)`

GetSkipStatusCheckOk returns a tuple with the SkipStatusCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipStatusCheck

`func (o *WorkflowCliCommandAllOf) SetSkipStatusCheck(v bool)`

SetSkipStatusCheck sets SkipStatusCheck field to given value.

### HasSkipStatusCheck

`func (o *WorkflowCliCommandAllOf) HasSkipStatusCheck() bool`

HasSkipStatusCheck returns a boolean if a field has been set.

### GetTerminalEnd

`func (o *WorkflowCliCommandAllOf) GetTerminalEnd() bool`

GetTerminalEnd returns the TerminalEnd field if non-nil, zero value otherwise.

### GetTerminalEndOk

`func (o *WorkflowCliCommandAllOf) GetTerminalEndOk() (*bool, bool)`

GetTerminalEndOk returns a tuple with the TerminalEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalEnd

`func (o *WorkflowCliCommandAllOf) SetTerminalEnd(v bool)`

SetTerminalEnd sets TerminalEnd field to given value.

### HasTerminalEnd

`func (o *WorkflowCliCommandAllOf) HasTerminalEnd() bool`

HasTerminalEnd returns a boolean if a field has been set.

### GetTerminalStart

`func (o *WorkflowCliCommandAllOf) GetTerminalStart() bool`

GetTerminalStart returns the TerminalStart field if non-nil, zero value otherwise.

### GetTerminalStartOk

`func (o *WorkflowCliCommandAllOf) GetTerminalStartOk() (*bool, bool)`

GetTerminalStartOk returns a tuple with the TerminalStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalStart

`func (o *WorkflowCliCommandAllOf) SetTerminalStart(v bool)`

SetTerminalStart sets TerminalStart field to given value.

### HasTerminalStart

`func (o *WorkflowCliCommandAllOf) HasTerminalStart() bool`

HasTerminalStart returns a boolean if a field has been set.

### GetType

`func (o *WorkflowCliCommandAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkflowCliCommandAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkflowCliCommandAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WorkflowCliCommandAllOf) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


