# SnmpTrap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | Pointer to **string** | Address to which the SNMP trap information is sent. | [optional] 
**Enabled** | Pointer to **bool** | Enables/disables the trap on the server If enabled, trap is active on the server. | [optional] 
**Port** | Pointer to **int64** | Port used by the server to communicate with trap destination. Enter a value between 1-65535. | [optional] 
**Type** | Pointer to **string** | Type of trap which decides whether to receive a notification when a trap is received at the destination. | [optional] [default to "Trap"]
**User** | Pointer to **string** | SNMP user for the trap. Applicable only to SNMPv3. | [optional] 
**Version** | Pointer to **string** | SNMP version used for the trap. | [optional] [default to "V3"]

## Methods

### NewSnmpTrap

`func NewSnmpTrap() *SnmpTrap`

NewSnmpTrap instantiates a new SnmpTrap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnmpTrapWithDefaults

`func NewSnmpTrapWithDefaults() *SnmpTrap`

NewSnmpTrapWithDefaults instantiates a new SnmpTrap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *SnmpTrap) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *SnmpTrap) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *SnmpTrap) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *SnmpTrap) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetEnabled

`func (o *SnmpTrap) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SnmpTrap) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SnmpTrap) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SnmpTrap) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPort

`func (o *SnmpTrap) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SnmpTrap) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SnmpTrap) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *SnmpTrap) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetType

`func (o *SnmpTrap) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SnmpTrap) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SnmpTrap) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SnmpTrap) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUser

`func (o *SnmpTrap) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SnmpTrap) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SnmpTrap) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *SnmpTrap) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetVersion

`func (o *SnmpTrap) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SnmpTrap) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SnmpTrap) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SnmpTrap) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


