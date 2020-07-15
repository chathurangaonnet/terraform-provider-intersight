# HclProductAllOf

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

### NewHclProductAllOf

`func NewHclProductAllOf() *HclProductAllOf`

NewHclProductAllOf instantiates a new HclProductAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHclProductAllOfWithDefaults

`func NewHclProductAllOfWithDefaults() *HclProductAllOf`

NewHclProductAllOfWithDefaults instantiates a new HclProductAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDriverNames

`func (o *HclProductAllOf) GetDriverNames() []string`

GetDriverNames returns the DriverNames field if non-nil, zero value otherwise.

### GetDriverNamesOk

`func (o *HclProductAllOf) GetDriverNamesOk() (*[]string, bool)`

GetDriverNamesOk returns a tuple with the DriverNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverNames

`func (o *HclProductAllOf) SetDriverNames(v []string)`

SetDriverNames sets DriverNames field to given value.

### HasDriverNames

`func (o *HclProductAllOf) HasDriverNames() bool`

HasDriverNames returns a boolean if a field has been set.

### GetErrorCode

`func (o *HclProductAllOf) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *HclProductAllOf) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *HclProductAllOf) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *HclProductAllOf) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetFirmwares

`func (o *HclProductAllOf) GetFirmwares() []HclFirmware`

GetFirmwares returns the Firmwares field if non-nil, zero value otherwise.

### GetFirmwaresOk

`func (o *HclProductAllOf) GetFirmwaresOk() (*[]HclFirmware, bool)`

GetFirmwaresOk returns a tuple with the Firmwares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwares

`func (o *HclProductAllOf) SetFirmwares(v []HclFirmware)`

SetFirmwares sets Firmwares field to given value.

### HasFirmwares

`func (o *HclProductAllOf) HasFirmwares() bool`

HasFirmwares returns a boolean if a field has been set.

### GetId

`func (o *HclProductAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HclProductAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HclProductAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HclProductAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModel

`func (o *HclProductAllOf) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *HclProductAllOf) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *HclProductAllOf) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *HclProductAllOf) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetRevision

`func (o *HclProductAllOf) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *HclProductAllOf) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *HclProductAllOf) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *HclProductAllOf) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetType

`func (o *HclProductAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HclProductAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HclProductAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HclProductAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVendor

`func (o *HclProductAllOf) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *HclProductAllOf) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *HclProductAllOf) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *HclProductAllOf) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


