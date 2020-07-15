# LicenseSmartlicenseToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** | Smart license registration token. | [optional] 
**AccountLicenseData** | Pointer to [**LicenseAccountLicenseDataRelationship**](license.AccountLicenseData.Relationship.md) |  | [optional] 

## Methods

### NewLicenseSmartlicenseToken

`func NewLicenseSmartlicenseToken() *LicenseSmartlicenseToken`

NewLicenseSmartlicenseToken instantiates a new LicenseSmartlicenseToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseSmartlicenseTokenWithDefaults

`func NewLicenseSmartlicenseTokenWithDefaults() *LicenseSmartlicenseToken`

NewLicenseSmartlicenseTokenWithDefaults instantiates a new LicenseSmartlicenseToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *LicenseSmartlicenseToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *LicenseSmartlicenseToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *LicenseSmartlicenseToken) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *LicenseSmartlicenseToken) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetAccountLicenseData

`func (o *LicenseSmartlicenseToken) GetAccountLicenseData() LicenseAccountLicenseDataRelationship`

GetAccountLicenseData returns the AccountLicenseData field if non-nil, zero value otherwise.

### GetAccountLicenseDataOk

`func (o *LicenseSmartlicenseToken) GetAccountLicenseDataOk() (*LicenseAccountLicenseDataRelationship, bool)`

GetAccountLicenseDataOk returns a tuple with the AccountLicenseData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLicenseData

`func (o *LicenseSmartlicenseToken) SetAccountLicenseData(v LicenseAccountLicenseDataRelationship)`

SetAccountLicenseData sets AccountLicenseData field to given value.

### HasAccountLicenseData

`func (o *LicenseSmartlicenseToken) HasAccountLicenseData() bool`

HasAccountLicenseData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


