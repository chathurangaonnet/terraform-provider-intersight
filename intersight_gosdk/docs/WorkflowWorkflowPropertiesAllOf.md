# WorkflowWorkflowPropertiesAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalMeta** | Pointer to **bool** | When set to false the workflow is owned by the system and used for internal services. Such workflows cannot be directly used by external entities. | [optional] 
**Retryable** | Pointer to **bool** | When true, this workflow can be retried if has not been modified for more than a period of 2 weeks. | [optional] 
**SupportStatus** | Pointer to **string** | Supported status of the definition. | [optional] [default to "Supported"]

## Methods

### NewWorkflowWorkflowPropertiesAllOf

`func NewWorkflowWorkflowPropertiesAllOf() *WorkflowWorkflowPropertiesAllOf`

NewWorkflowWorkflowPropertiesAllOf instantiates a new WorkflowWorkflowPropertiesAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowWorkflowPropertiesAllOfWithDefaults

`func NewWorkflowWorkflowPropertiesAllOfWithDefaults() *WorkflowWorkflowPropertiesAllOf`

NewWorkflowWorkflowPropertiesAllOfWithDefaults instantiates a new WorkflowWorkflowPropertiesAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalMeta

`func (o *WorkflowWorkflowPropertiesAllOf) GetExternalMeta() bool`

GetExternalMeta returns the ExternalMeta field if non-nil, zero value otherwise.

### GetExternalMetaOk

`func (o *WorkflowWorkflowPropertiesAllOf) GetExternalMetaOk() (*bool, bool)`

GetExternalMetaOk returns a tuple with the ExternalMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalMeta

`func (o *WorkflowWorkflowPropertiesAllOf) SetExternalMeta(v bool)`

SetExternalMeta sets ExternalMeta field to given value.

### HasExternalMeta

`func (o *WorkflowWorkflowPropertiesAllOf) HasExternalMeta() bool`

HasExternalMeta returns a boolean if a field has been set.

### GetRetryable

`func (o *WorkflowWorkflowPropertiesAllOf) GetRetryable() bool`

GetRetryable returns the Retryable field if non-nil, zero value otherwise.

### GetRetryableOk

`func (o *WorkflowWorkflowPropertiesAllOf) GetRetryableOk() (*bool, bool)`

GetRetryableOk returns a tuple with the Retryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryable

`func (o *WorkflowWorkflowPropertiesAllOf) SetRetryable(v bool)`

SetRetryable sets Retryable field to given value.

### HasRetryable

`func (o *WorkflowWorkflowPropertiesAllOf) HasRetryable() bool`

HasRetryable returns a boolean if a field has been set.

### GetSupportStatus

`func (o *WorkflowWorkflowPropertiesAllOf) GetSupportStatus() string`

GetSupportStatus returns the SupportStatus field if non-nil, zero value otherwise.

### GetSupportStatusOk

`func (o *WorkflowWorkflowPropertiesAllOf) GetSupportStatusOk() (*string, bool)`

GetSupportStatusOk returns a tuple with the SupportStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportStatus

`func (o *WorkflowWorkflowPropertiesAllOf) SetSupportStatus(v string)`

SetSupportStatus sets SupportStatus field to given value.

### HasSupportStatus

`func (o *WorkflowWorkflowPropertiesAllOf) HasSupportStatus() bool`

HasSupportStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


