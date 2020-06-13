# PkixEcdsaKeySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Curve** | Pointer to **string** | A specific set of Elliptic Curve parameters, as recommended by NIST in FIPS 186-4. | [optional] [default to "P256"]

## Methods

### NewPkixEcdsaKeySpec

`func NewPkixEcdsaKeySpec() *PkixEcdsaKeySpec`

NewPkixEcdsaKeySpec instantiates a new PkixEcdsaKeySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkixEcdsaKeySpecWithDefaults

`func NewPkixEcdsaKeySpecWithDefaults() *PkixEcdsaKeySpec`

NewPkixEcdsaKeySpecWithDefaults instantiates a new PkixEcdsaKeySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurve

`func (o *PkixEcdsaKeySpec) GetCurve() string`

GetCurve returns the Curve field if non-nil, zero value otherwise.

### GetCurveOk

`func (o *PkixEcdsaKeySpec) GetCurveOk() (*string, bool)`

GetCurveOk returns a tuple with the Curve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurve

`func (o *PkixEcdsaKeySpec) SetCurve(v string)`

SetCurve sets Curve field to given value.

### HasCurve

`func (o *PkixEcdsaKeySpec) HasCurve() bool`

HasCurve returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


