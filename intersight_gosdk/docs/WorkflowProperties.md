# WorkflowProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalMeta** | Pointer to **bool** | When set to false the task definition can only be used by internal system workflows. When set to true then the task can be included in user defined workflows. | [optional] 
**InputDefinition** | Pointer to [**[]WorkflowBaseDataType**](workflow.BaseDataType.md) |  | [optional] 
**OutputDefinition** | Pointer to [**[]WorkflowBaseDataType**](workflow.BaseDataType.md) |  | [optional] 
**RetryCount** | Pointer to **int64** | The number of times a task should be tried before marking as failed. | [optional] 
**RetryDelay** | Pointer to **int64** | The delay in seconds after which the the task is re-tried. | [optional] 
**RetryPolicy** | Pointer to **string** | The retry policy for the task. | [optional] [default to "Fixed"]
**SupportStatus** | Pointer to **string** | Supported status of the definition. | [optional] [default to "Supported"]
**Timeout** | Pointer to **int64** | The timeout value in seconds after which task will be marked as timed out. Max allowed value is 7 days. | [optional] 
**TimeoutPolicy** | Pointer to **string** | The timeout policy for the task. | [optional] [default to "Timeout"]

## Methods

### NewWorkflowProperties

`func NewWorkflowProperties() *WorkflowProperties`

NewWorkflowProperties instantiates a new WorkflowProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowPropertiesWithDefaults

`func NewWorkflowPropertiesWithDefaults() *WorkflowProperties`

NewWorkflowPropertiesWithDefaults instantiates a new WorkflowProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalMeta

`func (o *WorkflowProperties) GetExternalMeta() bool`

GetExternalMeta returns the ExternalMeta field if non-nil, zero value otherwise.

### GetExternalMetaOk

`func (o *WorkflowProperties) GetExternalMetaOk() (*bool, bool)`

GetExternalMetaOk returns a tuple with the ExternalMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalMeta

`func (o *WorkflowProperties) SetExternalMeta(v bool)`

SetExternalMeta sets ExternalMeta field to given value.

### HasExternalMeta

`func (o *WorkflowProperties) HasExternalMeta() bool`

HasExternalMeta returns a boolean if a field has been set.

### GetInputDefinition

`func (o *WorkflowProperties) GetInputDefinition() []WorkflowBaseDataType`

GetInputDefinition returns the InputDefinition field if non-nil, zero value otherwise.

### GetInputDefinitionOk

`func (o *WorkflowProperties) GetInputDefinitionOk() (*[]WorkflowBaseDataType, bool)`

GetInputDefinitionOk returns a tuple with the InputDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputDefinition

`func (o *WorkflowProperties) SetInputDefinition(v []WorkflowBaseDataType)`

SetInputDefinition sets InputDefinition field to given value.

### HasInputDefinition

`func (o *WorkflowProperties) HasInputDefinition() bool`

HasInputDefinition returns a boolean if a field has been set.

### GetOutputDefinition

`func (o *WorkflowProperties) GetOutputDefinition() []WorkflowBaseDataType`

GetOutputDefinition returns the OutputDefinition field if non-nil, zero value otherwise.

### GetOutputDefinitionOk

`func (o *WorkflowProperties) GetOutputDefinitionOk() (*[]WorkflowBaseDataType, bool)`

GetOutputDefinitionOk returns a tuple with the OutputDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputDefinition

`func (o *WorkflowProperties) SetOutputDefinition(v []WorkflowBaseDataType)`

SetOutputDefinition sets OutputDefinition field to given value.

### HasOutputDefinition

`func (o *WorkflowProperties) HasOutputDefinition() bool`

HasOutputDefinition returns a boolean if a field has been set.

### GetRetryCount

`func (o *WorkflowProperties) GetRetryCount() int64`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *WorkflowProperties) GetRetryCountOk() (*int64, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *WorkflowProperties) SetRetryCount(v int64)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *WorkflowProperties) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetRetryDelay

`func (o *WorkflowProperties) GetRetryDelay() int64`

GetRetryDelay returns the RetryDelay field if non-nil, zero value otherwise.

### GetRetryDelayOk

`func (o *WorkflowProperties) GetRetryDelayOk() (*int64, bool)`

GetRetryDelayOk returns a tuple with the RetryDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDelay

`func (o *WorkflowProperties) SetRetryDelay(v int64)`

SetRetryDelay sets RetryDelay field to given value.

### HasRetryDelay

`func (o *WorkflowProperties) HasRetryDelay() bool`

HasRetryDelay returns a boolean if a field has been set.

### GetRetryPolicy

`func (o *WorkflowProperties) GetRetryPolicy() string`

GetRetryPolicy returns the RetryPolicy field if non-nil, zero value otherwise.

### GetRetryPolicyOk

`func (o *WorkflowProperties) GetRetryPolicyOk() (*string, bool)`

GetRetryPolicyOk returns a tuple with the RetryPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryPolicy

`func (o *WorkflowProperties) SetRetryPolicy(v string)`

SetRetryPolicy sets RetryPolicy field to given value.

### HasRetryPolicy

`func (o *WorkflowProperties) HasRetryPolicy() bool`

HasRetryPolicy returns a boolean if a field has been set.

### GetSupportStatus

`func (o *WorkflowProperties) GetSupportStatus() string`

GetSupportStatus returns the SupportStatus field if non-nil, zero value otherwise.

### GetSupportStatusOk

`func (o *WorkflowProperties) GetSupportStatusOk() (*string, bool)`

GetSupportStatusOk returns a tuple with the SupportStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportStatus

`func (o *WorkflowProperties) SetSupportStatus(v string)`

SetSupportStatus sets SupportStatus field to given value.

### HasSupportStatus

`func (o *WorkflowProperties) HasSupportStatus() bool`

HasSupportStatus returns a boolean if a field has been set.

### GetTimeout

`func (o *WorkflowProperties) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *WorkflowProperties) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *WorkflowProperties) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *WorkflowProperties) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetTimeoutPolicy

`func (o *WorkflowProperties) GetTimeoutPolicy() string`

GetTimeoutPolicy returns the TimeoutPolicy field if non-nil, zero value otherwise.

### GetTimeoutPolicyOk

`func (o *WorkflowProperties) GetTimeoutPolicyOk() (*string, bool)`

GetTimeoutPolicyOk returns a tuple with the TimeoutPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutPolicy

`func (o *WorkflowProperties) SetTimeoutPolicy(v string)`

SetTimeoutPolicy sets TimeoutPolicy field to given value.

### HasTimeoutPolicy

`func (o *WorkflowProperties) HasTimeoutPolicy() bool`

HasTimeoutPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


