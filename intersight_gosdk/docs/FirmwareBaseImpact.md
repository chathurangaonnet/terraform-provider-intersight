# FirmwareBaseImpact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComputationError** | Pointer to **string** | Details for the error that occurred during the reboot validation analysis. | [optional] 
**ComputationStatusDetail** | Pointer to **string** | The computation status of the estimate operation for a component. | [optional] [default to "Inprogress"]
**DomainName** | Pointer to **string** | The endpoint type or name. | [optional] 
**EndPoint** | Pointer to **string** | A reference to a REST resource, uniquely identified by object type and MOID. | [optional] 
**FirmwareVersion** | Pointer to **string** | The current firmware version of the component. | [optional] 
**ImpactType** | Pointer to **string** | The impact type of the endpoint, whether a reboot is required or not. | [optional] [default to "NoReboot"]
**Model** | Pointer to **string** | The model details of the component. | [optional] 
**TargetFirmwareVersion** | Pointer to **string** | The target firmware version of the component. | [optional] 
**VersionImpact** | Pointer to **string** | The change of version impact for the endpoint. | [optional] [default to "None"]

## Methods

### NewFirmwareBaseImpact

`func NewFirmwareBaseImpact() *FirmwareBaseImpact`

NewFirmwareBaseImpact instantiates a new FirmwareBaseImpact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareBaseImpactWithDefaults

`func NewFirmwareBaseImpactWithDefaults() *FirmwareBaseImpact`

NewFirmwareBaseImpactWithDefaults instantiates a new FirmwareBaseImpact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComputationError

`func (o *FirmwareBaseImpact) GetComputationError() string`

GetComputationError returns the ComputationError field if non-nil, zero value otherwise.

### GetComputationErrorOk

`func (o *FirmwareBaseImpact) GetComputationErrorOk() (*string, bool)`

GetComputationErrorOk returns a tuple with the ComputationError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputationError

`func (o *FirmwareBaseImpact) SetComputationError(v string)`

SetComputationError sets ComputationError field to given value.

### HasComputationError

`func (o *FirmwareBaseImpact) HasComputationError() bool`

HasComputationError returns a boolean if a field has been set.

### GetComputationStatusDetail

`func (o *FirmwareBaseImpact) GetComputationStatusDetail() string`

GetComputationStatusDetail returns the ComputationStatusDetail field if non-nil, zero value otherwise.

### GetComputationStatusDetailOk

`func (o *FirmwareBaseImpact) GetComputationStatusDetailOk() (*string, bool)`

GetComputationStatusDetailOk returns a tuple with the ComputationStatusDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputationStatusDetail

`func (o *FirmwareBaseImpact) SetComputationStatusDetail(v string)`

SetComputationStatusDetail sets ComputationStatusDetail field to given value.

### HasComputationStatusDetail

`func (o *FirmwareBaseImpact) HasComputationStatusDetail() bool`

HasComputationStatusDetail returns a boolean if a field has been set.

### GetDomainName

`func (o *FirmwareBaseImpact) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *FirmwareBaseImpact) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *FirmwareBaseImpact) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *FirmwareBaseImpact) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetEndPoint

`func (o *FirmwareBaseImpact) GetEndPoint() string`

GetEndPoint returns the EndPoint field if non-nil, zero value otherwise.

### GetEndPointOk

`func (o *FirmwareBaseImpact) GetEndPointOk() (*string, bool)`

GetEndPointOk returns a tuple with the EndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPoint

`func (o *FirmwareBaseImpact) SetEndPoint(v string)`

SetEndPoint sets EndPoint field to given value.

### HasEndPoint

`func (o *FirmwareBaseImpact) HasEndPoint() bool`

HasEndPoint returns a boolean if a field has been set.

### GetFirmwareVersion

`func (o *FirmwareBaseImpact) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *FirmwareBaseImpact) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *FirmwareBaseImpact) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *FirmwareBaseImpact) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### GetImpactType

`func (o *FirmwareBaseImpact) GetImpactType() string`

GetImpactType returns the ImpactType field if non-nil, zero value otherwise.

### GetImpactTypeOk

`func (o *FirmwareBaseImpact) GetImpactTypeOk() (*string, bool)`

GetImpactTypeOk returns a tuple with the ImpactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactType

`func (o *FirmwareBaseImpact) SetImpactType(v string)`

SetImpactType sets ImpactType field to given value.

### HasImpactType

`func (o *FirmwareBaseImpact) HasImpactType() bool`

HasImpactType returns a boolean if a field has been set.

### GetModel

`func (o *FirmwareBaseImpact) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *FirmwareBaseImpact) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *FirmwareBaseImpact) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *FirmwareBaseImpact) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetTargetFirmwareVersion

`func (o *FirmwareBaseImpact) GetTargetFirmwareVersion() string`

GetTargetFirmwareVersion returns the TargetFirmwareVersion field if non-nil, zero value otherwise.

### GetTargetFirmwareVersionOk

`func (o *FirmwareBaseImpact) GetTargetFirmwareVersionOk() (*string, bool)`

GetTargetFirmwareVersionOk returns a tuple with the TargetFirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetFirmwareVersion

`func (o *FirmwareBaseImpact) SetTargetFirmwareVersion(v string)`

SetTargetFirmwareVersion sets TargetFirmwareVersion field to given value.

### HasTargetFirmwareVersion

`func (o *FirmwareBaseImpact) HasTargetFirmwareVersion() bool`

HasTargetFirmwareVersion returns a boolean if a field has been set.

### GetVersionImpact

`func (o *FirmwareBaseImpact) GetVersionImpact() string`

GetVersionImpact returns the VersionImpact field if non-nil, zero value otherwise.

### GetVersionImpactOk

`func (o *FirmwareBaseImpact) GetVersionImpactOk() (*string, bool)`

GetVersionImpactOk returns a tuple with the VersionImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionImpact

`func (o *FirmwareBaseImpact) SetVersionImpact(v string)`

SetVersionImpact sets VersionImpact field to given value.

### HasVersionImpact

`func (o *FirmwareBaseImpact) HasVersionImpact() bool`

HasVersionImpact returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


