# IamLdapDnsParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchDomain** | Pointer to **string** | Domain name that acts as a source for a DNS query. | [optional] 
**SearchForest** | Pointer to **string** | Forest name that acts as a source for a DNS query. | [optional] 
**Source** | Pointer to **string** | Source of the domain name used for the DNS SRV request. | [optional] [default to "Extracted"]

## Methods

### NewIamLdapDnsParameters

`func NewIamLdapDnsParameters() *IamLdapDnsParameters`

NewIamLdapDnsParameters instantiates a new IamLdapDnsParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamLdapDnsParametersWithDefaults

`func NewIamLdapDnsParametersWithDefaults() *IamLdapDnsParameters`

NewIamLdapDnsParametersWithDefaults instantiates a new IamLdapDnsParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchDomain

`func (o *IamLdapDnsParameters) GetSearchDomain() string`

GetSearchDomain returns the SearchDomain field if non-nil, zero value otherwise.

### GetSearchDomainOk

`func (o *IamLdapDnsParameters) GetSearchDomainOk() (*string, bool)`

GetSearchDomainOk returns a tuple with the SearchDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchDomain

`func (o *IamLdapDnsParameters) SetSearchDomain(v string)`

SetSearchDomain sets SearchDomain field to given value.

### HasSearchDomain

`func (o *IamLdapDnsParameters) HasSearchDomain() bool`

HasSearchDomain returns a boolean if a field has been set.

### GetSearchForest

`func (o *IamLdapDnsParameters) GetSearchForest() string`

GetSearchForest returns the SearchForest field if non-nil, zero value otherwise.

### GetSearchForestOk

`func (o *IamLdapDnsParameters) GetSearchForestOk() (*string, bool)`

GetSearchForestOk returns a tuple with the SearchForest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchForest

`func (o *IamLdapDnsParameters) SetSearchForest(v string)`

SetSearchForest sets SearchForest field to given value.

### HasSearchForest

`func (o *IamLdapDnsParameters) HasSearchForest() bool`

HasSearchForest returns a boolean if a field has been set.

### GetSource

`func (o *IamLdapDnsParameters) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *IamLdapDnsParameters) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *IamLdapDnsParameters) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *IamLdapDnsParameters) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


