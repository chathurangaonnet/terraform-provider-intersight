# SyslogRemoteClientBaseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enables/disables remote logging for the endpoint If enabled, log messages will be sent to the syslog server mentioned in the Hostname/IP Address field. | [optional] 
**Hostname** | Pointer to **string** | Hostname or IP Address of the syslog server where log should be stored. | [optional] 
**MinSeverity** | Pointer to **string** | Lowest level of messages to be included in the remote log. | [optional] [default to "warning"]
**Port** | Pointer to **int64** | Port number used for logging on syslog server. | [optional] 
**Protocol** | Pointer to **string** | Transport layer protocol for transmission of log messages to syslog server. | [optional] [default to "udp"]

## Methods

### NewSyslogRemoteClientBaseAllOf

`func NewSyslogRemoteClientBaseAllOf() *SyslogRemoteClientBaseAllOf`

NewSyslogRemoteClientBaseAllOf instantiates a new SyslogRemoteClientBaseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyslogRemoteClientBaseAllOfWithDefaults

`func NewSyslogRemoteClientBaseAllOfWithDefaults() *SyslogRemoteClientBaseAllOf`

NewSyslogRemoteClientBaseAllOfWithDefaults instantiates a new SyslogRemoteClientBaseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SyslogRemoteClientBaseAllOf) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SyslogRemoteClientBaseAllOf) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SyslogRemoteClientBaseAllOf) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SyslogRemoteClientBaseAllOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHostname

`func (o *SyslogRemoteClientBaseAllOf) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *SyslogRemoteClientBaseAllOf) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *SyslogRemoteClientBaseAllOf) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *SyslogRemoteClientBaseAllOf) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetMinSeverity

`func (o *SyslogRemoteClientBaseAllOf) GetMinSeverity() string`

GetMinSeverity returns the MinSeverity field if non-nil, zero value otherwise.

### GetMinSeverityOk

`func (o *SyslogRemoteClientBaseAllOf) GetMinSeverityOk() (*string, bool)`

GetMinSeverityOk returns a tuple with the MinSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSeverity

`func (o *SyslogRemoteClientBaseAllOf) SetMinSeverity(v string)`

SetMinSeverity sets MinSeverity field to given value.

### HasMinSeverity

`func (o *SyslogRemoteClientBaseAllOf) HasMinSeverity() bool`

HasMinSeverity returns a boolean if a field has been set.

### GetPort

`func (o *SyslogRemoteClientBaseAllOf) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SyslogRemoteClientBaseAllOf) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SyslogRemoteClientBaseAllOf) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *SyslogRemoteClientBaseAllOf) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *SyslogRemoteClientBaseAllOf) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *SyslogRemoteClientBaseAllOf) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *SyslogRemoteClientBaseAllOf) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *SyslogRemoteClientBaseAllOf) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


