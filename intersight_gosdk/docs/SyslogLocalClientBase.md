# SyslogLocalClientBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinSeverity** | Pointer to **string** | Lowest level of messages to be included in the local log. | [optional] [default to "warning"]

## Methods

### NewSyslogLocalClientBase

`func NewSyslogLocalClientBase() *SyslogLocalClientBase`

NewSyslogLocalClientBase instantiates a new SyslogLocalClientBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyslogLocalClientBaseWithDefaults

`func NewSyslogLocalClientBaseWithDefaults() *SyslogLocalClientBase`

NewSyslogLocalClientBaseWithDefaults instantiates a new SyslogLocalClientBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinSeverity

`func (o *SyslogLocalClientBase) GetMinSeverity() string`

GetMinSeverity returns the MinSeverity field if non-nil, zero value otherwise.

### GetMinSeverityOk

`func (o *SyslogLocalClientBase) GetMinSeverityOk() (*string, bool)`

GetMinSeverityOk returns a tuple with the MinSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSeverity

`func (o *SyslogLocalClientBase) SetMinSeverity(v string)`

SetMinSeverity sets MinSeverity field to given value.

### HasMinSeverity

`func (o *SyslogLocalClientBase) HasMinSeverity() bool`

HasMinSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


