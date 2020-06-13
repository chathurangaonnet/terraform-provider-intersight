# SyslogRemoteClientBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enables/disables remote logging for the endpoint If enabled, log messages will be sent to the syslog server mentioned in the Hostname/IP Address field. | [optional] 
**Hostname** | Pointer to **string** | Hostname or IP Address of the syslog server where log should be stored. | [optional] 
**MinSeverity** | Pointer to **string** | Lowest level of messages to be included in the remote log. | [optional] [default to "warning"]
**Port** | Pointer to **int64** | Port number used for logging on syslog server. | [optional] 
**Protocol** | Pointer to **string** | Transport layer protocol for transmission of log messages to syslog server. | [optional] [default to "udp"]

## Methods

### NewSyslogRemoteClientBase

`func NewSyslogRemoteClientBase() *SyslogRemoteClientBase`

NewSyslogRemoteClientBase instantiates a new SyslogRemoteClientBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyslogRemoteClientBaseWithDefaults

`func NewSyslogRemoteClientBaseWithDefaults() *SyslogRemoteClientBase`

NewSyslogRemoteClientBaseWithDefaults instantiates a new SyslogRemoteClientBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SyslogRemoteClientBase) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SyslogRemoteClientBase) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SyslogRemoteClientBase) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SyslogRemoteClientBase) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHostname

`func (o *SyslogRemoteClientBase) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *SyslogRemoteClientBase) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *SyslogRemoteClientBase) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *SyslogRemoteClientBase) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetMinSeverity

`func (o *SyslogRemoteClientBase) GetMinSeverity() string`

GetMinSeverity returns the MinSeverity field if non-nil, zero value otherwise.

### GetMinSeverityOk

`func (o *SyslogRemoteClientBase) GetMinSeverityOk() (*string, bool)`

GetMinSeverityOk returns a tuple with the MinSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSeverity

`func (o *SyslogRemoteClientBase) SetMinSeverity(v string)`

SetMinSeverity sets MinSeverity field to given value.

### HasMinSeverity

`func (o *SyslogRemoteClientBase) HasMinSeverity() bool`

HasMinSeverity returns a boolean if a field has been set.

### GetPort

`func (o *SyslogRemoteClientBase) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SyslogRemoteClientBase) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SyslogRemoteClientBase) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *SyslogRemoteClientBase) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *SyslogRemoteClientBase) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *SyslogRemoteClientBase) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *SyslogRemoteClientBase) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *SyslogRemoteClientBase) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


