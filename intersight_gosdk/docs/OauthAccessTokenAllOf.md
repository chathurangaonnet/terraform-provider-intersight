# OauthAccessTokenAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiType** | Pointer to **string** | Oauth api type | [optional] [default to "Unknown"]
**ExpiryTime** | Pointer to [**time.Time**](time.Time.md) | Access token expiry time | [optional] [readonly] 
**Status** | Pointer to **string** | Access token status | [optional] [default to "Inactive"]
**Token** | Pointer to **string** | Access token | [optional] [readonly] 

## Methods

### NewOauthAccessTokenAllOf

`func NewOauthAccessTokenAllOf() *OauthAccessTokenAllOf`

NewOauthAccessTokenAllOf instantiates a new OauthAccessTokenAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOauthAccessTokenAllOfWithDefaults

`func NewOauthAccessTokenAllOfWithDefaults() *OauthAccessTokenAllOf`

NewOauthAccessTokenAllOfWithDefaults instantiates a new OauthAccessTokenAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiType

`func (o *OauthAccessTokenAllOf) GetApiType() string`

GetApiType returns the ApiType field if non-nil, zero value otherwise.

### GetApiTypeOk

`func (o *OauthAccessTokenAllOf) GetApiTypeOk() (*string, bool)`

GetApiTypeOk returns a tuple with the ApiType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiType

`func (o *OauthAccessTokenAllOf) SetApiType(v string)`

SetApiType sets ApiType field to given value.

### HasApiType

`func (o *OauthAccessTokenAllOf) HasApiType() bool`

HasApiType returns a boolean if a field has been set.

### GetExpiryTime

`func (o *OauthAccessTokenAllOf) GetExpiryTime() time.Time`

GetExpiryTime returns the ExpiryTime field if non-nil, zero value otherwise.

### GetExpiryTimeOk

`func (o *OauthAccessTokenAllOf) GetExpiryTimeOk() (*time.Time, bool)`

GetExpiryTimeOk returns a tuple with the ExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTime

`func (o *OauthAccessTokenAllOf) SetExpiryTime(v time.Time)`

SetExpiryTime sets ExpiryTime field to given value.

### HasExpiryTime

`func (o *OauthAccessTokenAllOf) HasExpiryTime() bool`

HasExpiryTime returns a boolean if a field has been set.

### GetStatus

`func (o *OauthAccessTokenAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OauthAccessTokenAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OauthAccessTokenAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OauthAccessTokenAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetToken

`func (o *OauthAccessTokenAllOf) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *OauthAccessTokenAllOf) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *OauthAccessTokenAllOf) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *OauthAccessTokenAllOf) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


