# WorkflowSshCmdAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | Pointer to **string** | SSH command to execute on the remote server. | [optional] 
**CommandType** | Pointer to **string** | SSH command type to execute on the remote server. | [optional] [default to "NonInteractiveCmd"]
**ExpectPrompts** | Pointer to [**[]ConnectorExpectPrompt**](connector.ExpectPrompt.md) |  | [optional] 
**ShellPrompt** | Pointer to **string** | Regex of the remote server&#39;s shell prompt. | [optional] 
**ShellPromptTimeout** | Pointer to **int64** | Expect timeout value in seconds for the shell prompt. | [optional] 

## Methods

### NewWorkflowSshCmdAllOf

`func NewWorkflowSshCmdAllOf() *WorkflowSshCmdAllOf`

NewWorkflowSshCmdAllOf instantiates a new WorkflowSshCmdAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowSshCmdAllOfWithDefaults

`func NewWorkflowSshCmdAllOfWithDefaults() *WorkflowSshCmdAllOf`

NewWorkflowSshCmdAllOfWithDefaults instantiates a new WorkflowSshCmdAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *WorkflowSshCmdAllOf) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *WorkflowSshCmdAllOf) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *WorkflowSshCmdAllOf) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *WorkflowSshCmdAllOf) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetCommandType

`func (o *WorkflowSshCmdAllOf) GetCommandType() string`

GetCommandType returns the CommandType field if non-nil, zero value otherwise.

### GetCommandTypeOk

`func (o *WorkflowSshCmdAllOf) GetCommandTypeOk() (*string, bool)`

GetCommandTypeOk returns a tuple with the CommandType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandType

`func (o *WorkflowSshCmdAllOf) SetCommandType(v string)`

SetCommandType sets CommandType field to given value.

### HasCommandType

`func (o *WorkflowSshCmdAllOf) HasCommandType() bool`

HasCommandType returns a boolean if a field has been set.

### GetExpectPrompts

`func (o *WorkflowSshCmdAllOf) GetExpectPrompts() []ConnectorExpectPrompt`

GetExpectPrompts returns the ExpectPrompts field if non-nil, zero value otherwise.

### GetExpectPromptsOk

`func (o *WorkflowSshCmdAllOf) GetExpectPromptsOk() (*[]ConnectorExpectPrompt, bool)`

GetExpectPromptsOk returns a tuple with the ExpectPrompts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectPrompts

`func (o *WorkflowSshCmdAllOf) SetExpectPrompts(v []ConnectorExpectPrompt)`

SetExpectPrompts sets ExpectPrompts field to given value.

### HasExpectPrompts

`func (o *WorkflowSshCmdAllOf) HasExpectPrompts() bool`

HasExpectPrompts returns a boolean if a field has been set.

### GetShellPrompt

`func (o *WorkflowSshCmdAllOf) GetShellPrompt() string`

GetShellPrompt returns the ShellPrompt field if non-nil, zero value otherwise.

### GetShellPromptOk

`func (o *WorkflowSshCmdAllOf) GetShellPromptOk() (*string, bool)`

GetShellPromptOk returns a tuple with the ShellPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShellPrompt

`func (o *WorkflowSshCmdAllOf) SetShellPrompt(v string)`

SetShellPrompt sets ShellPrompt field to given value.

### HasShellPrompt

`func (o *WorkflowSshCmdAllOf) HasShellPrompt() bool`

HasShellPrompt returns a boolean if a field has been set.

### GetShellPromptTimeout

`func (o *WorkflowSshCmdAllOf) GetShellPromptTimeout() int64`

GetShellPromptTimeout returns the ShellPromptTimeout field if non-nil, zero value otherwise.

### GetShellPromptTimeoutOk

`func (o *WorkflowSshCmdAllOf) GetShellPromptTimeoutOk() (*int64, bool)`

GetShellPromptTimeoutOk returns a tuple with the ShellPromptTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShellPromptTimeout

`func (o *WorkflowSshCmdAllOf) SetShellPromptTimeout(v int64)`

SetShellPromptTimeout sets ShellPromptTimeout field to given value.

### HasShellPromptTimeout

`func (o *WorkflowSshCmdAllOf) HasShellPromptTimeout() bool`

HasShellPromptTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


