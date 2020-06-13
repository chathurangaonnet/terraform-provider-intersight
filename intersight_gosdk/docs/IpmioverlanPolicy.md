# IpmioverlanPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | State of the IPMI Over LAN service on the endpoint. | [optional] 
**EncryptionKey** | Pointer to **string** | The encryption key to use for IPMI communication. It should have an even number of hexadecimal characters and not exceed 40 characters. | [optional] 
**IsEncryptionKeySet** | Pointer to **bool** | Indicates whether the value of the &#39;encryptionKey&#39; property has been set. | [optional] [readonly] 
**Privilege** | Pointer to **string** | The highest privilege level that can be assigned to an IPMI session on a server. | [optional] [default to "admin"]
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]PolicyAbstractConfigProfileRelationship**](policy.AbstractConfigProfile.Relationship.md) | An array of relationships to policyAbstractConfigProfile resources. | [optional] 

## Methods

### NewIpmioverlanPolicy

`func NewIpmioverlanPolicy() *IpmioverlanPolicy`

NewIpmioverlanPolicy instantiates a new IpmioverlanPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpmioverlanPolicyWithDefaults

`func NewIpmioverlanPolicyWithDefaults() *IpmioverlanPolicy`

NewIpmioverlanPolicyWithDefaults instantiates a new IpmioverlanPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *IpmioverlanPolicy) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IpmioverlanPolicy) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IpmioverlanPolicy) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *IpmioverlanPolicy) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEncryptionKey

`func (o *IpmioverlanPolicy) GetEncryptionKey() string`

GetEncryptionKey returns the EncryptionKey field if non-nil, zero value otherwise.

### GetEncryptionKeyOk

`func (o *IpmioverlanPolicy) GetEncryptionKeyOk() (*string, bool)`

GetEncryptionKeyOk returns a tuple with the EncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKey

`func (o *IpmioverlanPolicy) SetEncryptionKey(v string)`

SetEncryptionKey sets EncryptionKey field to given value.

### HasEncryptionKey

`func (o *IpmioverlanPolicy) HasEncryptionKey() bool`

HasEncryptionKey returns a boolean if a field has been set.

### GetIsEncryptionKeySet

`func (o *IpmioverlanPolicy) GetIsEncryptionKeySet() bool`

GetIsEncryptionKeySet returns the IsEncryptionKeySet field if non-nil, zero value otherwise.

### GetIsEncryptionKeySetOk

`func (o *IpmioverlanPolicy) GetIsEncryptionKeySetOk() (*bool, bool)`

GetIsEncryptionKeySetOk returns a tuple with the IsEncryptionKeySet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEncryptionKeySet

`func (o *IpmioverlanPolicy) SetIsEncryptionKeySet(v bool)`

SetIsEncryptionKeySet sets IsEncryptionKeySet field to given value.

### HasIsEncryptionKeySet

`func (o *IpmioverlanPolicy) HasIsEncryptionKeySet() bool`

HasIsEncryptionKeySet returns a boolean if a field has been set.

### GetPrivilege

`func (o *IpmioverlanPolicy) GetPrivilege() string`

GetPrivilege returns the Privilege field if non-nil, zero value otherwise.

### GetPrivilegeOk

`func (o *IpmioverlanPolicy) GetPrivilegeOk() (*string, bool)`

GetPrivilegeOk returns a tuple with the Privilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilege

`func (o *IpmioverlanPolicy) SetPrivilege(v string)`

SetPrivilege sets Privilege field to given value.

### HasPrivilege

`func (o *IpmioverlanPolicy) HasPrivilege() bool`

HasPrivilege returns a boolean if a field has been set.

### GetOrganization

`func (o *IpmioverlanPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *IpmioverlanPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *IpmioverlanPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *IpmioverlanPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfiles

`func (o *IpmioverlanPolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *IpmioverlanPolicy) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *IpmioverlanPolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *IpmioverlanPolicy) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *IpmioverlanPolicy) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *IpmioverlanPolicy) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


