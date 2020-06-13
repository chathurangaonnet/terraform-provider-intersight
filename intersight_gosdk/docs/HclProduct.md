# HclProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DriverNames** | Pointer to **[]string** |  | [optional] 
**ErrorCode** | Pointer to **string** | Error code indicating the support status. | [optional] [readonly] [default to "Success"]
**Firmwares** | Pointer to [**[]HclFirmware**](hcl.Firmware.md) |  | [optional] 
**Id** | Pointer to **string** | Identifier of the product. | [optional] 
**Model** | Pointer to **string** | Model/PID of the product/adapter. | [optional] 
**Revision** | Pointer to **string** | Revision of the adapter model. | [optional] 
**Type** | Pointer to **string** | Type of the product/adapter say OCP, PT, GPU. | [optional] 
**Vendor** | Pointer to **string** | Vendor of the product or adapter. | [optional] 

## Methods

### NewHclProduct

`func NewHclProduct() *HclProduct`

NewHclProduct instantiates a new HclProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHclProductWithDefaults

`func NewHclProductWithDefaults() *HclProduct`

NewHclProductWithDefaults instantiates a new HclProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDriverNames

`func (o *HclProduct) GetDriverNames() []string`

GetDriverNames returns the DriverNames field if non-nil, zero value otherwise.

### GetDriverNamesOk

`func (o *HclProduct) GetDriverNamesOk() (*[]string, bool)`

GetDriverNamesOk returns a tuple with the DriverNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverNames

`func (o *HclProduct) SetDriverNames(v []string)`

SetDriverNames sets DriverNames field to given value.

### HasDriverNames

`func (o *HclProduct) HasDriverNames() bool`

HasDriverNames returns a boolean if a field has been set.

### GetErrorCode

`func (o *HclProduct) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *HclProduct) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *HclProduct) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *HclProduct) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetFirmwares

`func (o *HclProduct) GetFirmwares() []HclFirmware`

GetFirmwares returns the Firmwares field if non-nil, zero value otherwise.

### GetFirmwaresOk

`func (o *HclProduct) GetFirmwaresOk() (*[]HclFirmware, bool)`

GetFirmwaresOk returns a tuple with the Firmwares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwares

`func (o *HclProduct) SetFirmwares(v []HclFirmware)`

SetFirmwares sets Firmwares field to given value.

### HasFirmwares

`func (o *HclProduct) HasFirmwares() bool`

HasFirmwares returns a boolean if a field has been set.

### GetId

`func (o *HclProduct) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HclProduct) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HclProduct) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HclProduct) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModel

`func (o *HclProduct) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *HclProduct) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *HclProduct) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *HclProduct) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetRevision

`func (o *HclProduct) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *HclProduct) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *HclProduct) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *HclProduct) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetType

`func (o *HclProduct) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HclProduct) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HclProduct) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HclProduct) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVendor

`func (o *HclProduct) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *HclProduct) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *HclProduct) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *HclProduct) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


