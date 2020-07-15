# SolPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaudRate** | Pointer to **int32** | Baud Rate used for Serial Over LAN communication. | [optional] [default to 9600]
**ComPort** | Pointer to **string** | Serial port through which the system routes Serial Over LAN communication. This field is available only on some Cisco UCS C-Series servers. If it is unavailable, the server uses COM port 0 by default. | [optional] [default to "com0"]
**Enabled** | Pointer to **bool** | State of Serial Over LAN service on the endpoint. | [optional] 
**SshPort** | Pointer to **int64** | SSH port used to access Serial Over LAN directly. Enables bypassing Cisco IMC shell to provide direct access to Serial Over LAN. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]PolicyAbstractConfigProfileRelationship**](policy.AbstractConfigProfile.Relationship.md) | An array of relationships to policyAbstractConfigProfile resources. | [optional] 

## Methods

### NewSolPolicyAllOf

`func NewSolPolicyAllOf() *SolPolicyAllOf`

NewSolPolicyAllOf instantiates a new SolPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSolPolicyAllOfWithDefaults

`func NewSolPolicyAllOfWithDefaults() *SolPolicyAllOf`

NewSolPolicyAllOfWithDefaults instantiates a new SolPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaudRate

`func (o *SolPolicyAllOf) GetBaudRate() int32`

GetBaudRate returns the BaudRate field if non-nil, zero value otherwise.

### GetBaudRateOk

`func (o *SolPolicyAllOf) GetBaudRateOk() (*int32, bool)`

GetBaudRateOk returns a tuple with the BaudRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaudRate

`func (o *SolPolicyAllOf) SetBaudRate(v int32)`

SetBaudRate sets BaudRate field to given value.

### HasBaudRate

`func (o *SolPolicyAllOf) HasBaudRate() bool`

HasBaudRate returns a boolean if a field has been set.

### GetComPort

`func (o *SolPolicyAllOf) GetComPort() string`

GetComPort returns the ComPort field if non-nil, zero value otherwise.

### GetComPortOk

`func (o *SolPolicyAllOf) GetComPortOk() (*string, bool)`

GetComPortOk returns a tuple with the ComPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComPort

`func (o *SolPolicyAllOf) SetComPort(v string)`

SetComPort sets ComPort field to given value.

### HasComPort

`func (o *SolPolicyAllOf) HasComPort() bool`

HasComPort returns a boolean if a field has been set.

### GetEnabled

`func (o *SolPolicyAllOf) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SolPolicyAllOf) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SolPolicyAllOf) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SolPolicyAllOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSshPort

`func (o *SolPolicyAllOf) GetSshPort() int64`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *SolPolicyAllOf) GetSshPortOk() (*int64, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *SolPolicyAllOf) SetSshPort(v int64)`

SetSshPort sets SshPort field to given value.

### HasSshPort

`func (o *SolPolicyAllOf) HasSshPort() bool`

HasSshPort returns a boolean if a field has been set.

### GetOrganization

`func (o *SolPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *SolPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *SolPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *SolPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfiles

`func (o *SolPolicyAllOf) GetProfiles() []PolicyAbstractConfigProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *SolPolicyAllOf) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *SolPolicyAllOf) SetProfiles(v []PolicyAbstractConfigProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *SolPolicyAllOf) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *SolPolicyAllOf) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *SolPolicyAllOf) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


