# CommIpV6InterfaceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gateway** | Pointer to **string** | The IPv6 address of the default gateway. | [optional] 
**IpAddress** | Pointer to **string** | The IPv6 Address, represented as eight groups of four hexadecimal digits, separated by colons. | [optional] 
**Prefix** | Pointer to **string** | The IPv6 Prefix, represented as eight groups of four hexadecimal digits, separated by colons. | [optional] 

## Methods

### NewCommIpV6InterfaceAllOf

`func NewCommIpV6InterfaceAllOf() *CommIpV6InterfaceAllOf`

NewCommIpV6InterfaceAllOf instantiates a new CommIpV6InterfaceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommIpV6InterfaceAllOfWithDefaults

`func NewCommIpV6InterfaceAllOfWithDefaults() *CommIpV6InterfaceAllOf`

NewCommIpV6InterfaceAllOfWithDefaults instantiates a new CommIpV6InterfaceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGateway

`func (o *CommIpV6InterfaceAllOf) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *CommIpV6InterfaceAllOf) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *CommIpV6InterfaceAllOf) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *CommIpV6InterfaceAllOf) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetIpAddress

`func (o *CommIpV6InterfaceAllOf) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *CommIpV6InterfaceAllOf) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *CommIpV6InterfaceAllOf) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *CommIpV6InterfaceAllOf) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetPrefix

`func (o *CommIpV6InterfaceAllOf) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *CommIpV6InterfaceAllOf) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *CommIpV6InterfaceAllOf) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *CommIpV6InterfaceAllOf) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


