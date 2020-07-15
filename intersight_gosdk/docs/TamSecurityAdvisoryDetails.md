# TamSecurityAdvisoryDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseScore** | Pointer to **float32** | CVSS version 3 base score for the security Advisory. | [optional] 
**CveIds** | Pointer to **[]string** |  | [optional] 
**EnvironmentalScore** | Pointer to **float32** | CVSS version 3 environmental score for the security Advisory. | [optional] 
**Status** | Pointer to **string** | Cisco assigned status of the published advisory. Depends on whether the investigation is complete or on-going. | [optional] [default to "interim"]
**TemporalScore** | Pointer to **float32** | CVSS version 3 temporal score for the security Advisory. | [optional] 

## Methods

### NewTamSecurityAdvisoryDetails

`func NewTamSecurityAdvisoryDetails() *TamSecurityAdvisoryDetails`

NewTamSecurityAdvisoryDetails instantiates a new TamSecurityAdvisoryDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTamSecurityAdvisoryDetailsWithDefaults

`func NewTamSecurityAdvisoryDetailsWithDefaults() *TamSecurityAdvisoryDetails`

NewTamSecurityAdvisoryDetailsWithDefaults instantiates a new TamSecurityAdvisoryDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseScore

`func (o *TamSecurityAdvisoryDetails) GetBaseScore() float32`

GetBaseScore returns the BaseScore field if non-nil, zero value otherwise.

### GetBaseScoreOk

`func (o *TamSecurityAdvisoryDetails) GetBaseScoreOk() (*float32, bool)`

GetBaseScoreOk returns a tuple with the BaseScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseScore

`func (o *TamSecurityAdvisoryDetails) SetBaseScore(v float32)`

SetBaseScore sets BaseScore field to given value.

### HasBaseScore

`func (o *TamSecurityAdvisoryDetails) HasBaseScore() bool`

HasBaseScore returns a boolean if a field has been set.

### GetCveIds

`func (o *TamSecurityAdvisoryDetails) GetCveIds() []string`

GetCveIds returns the CveIds field if non-nil, zero value otherwise.

### GetCveIdsOk

`func (o *TamSecurityAdvisoryDetails) GetCveIdsOk() (*[]string, bool)`

GetCveIdsOk returns a tuple with the CveIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCveIds

`func (o *TamSecurityAdvisoryDetails) SetCveIds(v []string)`

SetCveIds sets CveIds field to given value.

### HasCveIds

`func (o *TamSecurityAdvisoryDetails) HasCveIds() bool`

HasCveIds returns a boolean if a field has been set.

### GetEnvironmentalScore

`func (o *TamSecurityAdvisoryDetails) GetEnvironmentalScore() float32`

GetEnvironmentalScore returns the EnvironmentalScore field if non-nil, zero value otherwise.

### GetEnvironmentalScoreOk

`func (o *TamSecurityAdvisoryDetails) GetEnvironmentalScoreOk() (*float32, bool)`

GetEnvironmentalScoreOk returns a tuple with the EnvironmentalScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentalScore

`func (o *TamSecurityAdvisoryDetails) SetEnvironmentalScore(v float32)`

SetEnvironmentalScore sets EnvironmentalScore field to given value.

### HasEnvironmentalScore

`func (o *TamSecurityAdvisoryDetails) HasEnvironmentalScore() bool`

HasEnvironmentalScore returns a boolean if a field has been set.

### GetStatus

`func (o *TamSecurityAdvisoryDetails) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TamSecurityAdvisoryDetails) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TamSecurityAdvisoryDetails) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TamSecurityAdvisoryDetails) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTemporalScore

`func (o *TamSecurityAdvisoryDetails) GetTemporalScore() float32`

GetTemporalScore returns the TemporalScore field if non-nil, zero value otherwise.

### GetTemporalScoreOk

`func (o *TamSecurityAdvisoryDetails) GetTemporalScoreOk() (*float32, bool)`

GetTemporalScoreOk returns a tuple with the TemporalScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporalScore

`func (o *TamSecurityAdvisoryDetails) SetTemporalScore(v float32)`

SetTemporalScore sets TemporalScore field to given value.

### HasTemporalScore

`func (o *TamSecurityAdvisoryDetails) HasTemporalScore() bool`

HasTemporalScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


