# VnicTcpOffloadSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LargeReceive** | Pointer to **bool** | Enables the reassembly of segmented packets in hardware before sending them to the CPU. | [optional] 
**LargeSend** | Pointer to **bool** | Enables the CPU to send large packets to the hardware for segmentation. | [optional] 
**RxChecksum** | Pointer to **bool** | When enabled, the CPU sends all packet checksums to the hardware for validation. | [optional] 
**TxChecksum** | Pointer to **bool** | When enabled, the CPU sends all packets to the hardware so that the checksum can be calculated. | [optional] 

## Methods

### NewVnicTcpOffloadSettings

`func NewVnicTcpOffloadSettings() *VnicTcpOffloadSettings`

NewVnicTcpOffloadSettings instantiates a new VnicTcpOffloadSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicTcpOffloadSettingsWithDefaults

`func NewVnicTcpOffloadSettingsWithDefaults() *VnicTcpOffloadSettings`

NewVnicTcpOffloadSettingsWithDefaults instantiates a new VnicTcpOffloadSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLargeReceive

`func (o *VnicTcpOffloadSettings) GetLargeReceive() bool`

GetLargeReceive returns the LargeReceive field if non-nil, zero value otherwise.

### GetLargeReceiveOk

`func (o *VnicTcpOffloadSettings) GetLargeReceiveOk() (*bool, bool)`

GetLargeReceiveOk returns a tuple with the LargeReceive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargeReceive

`func (o *VnicTcpOffloadSettings) SetLargeReceive(v bool)`

SetLargeReceive sets LargeReceive field to given value.

### HasLargeReceive

`func (o *VnicTcpOffloadSettings) HasLargeReceive() bool`

HasLargeReceive returns a boolean if a field has been set.

### GetLargeSend

`func (o *VnicTcpOffloadSettings) GetLargeSend() bool`

GetLargeSend returns the LargeSend field if non-nil, zero value otherwise.

### GetLargeSendOk

`func (o *VnicTcpOffloadSettings) GetLargeSendOk() (*bool, bool)`

GetLargeSendOk returns a tuple with the LargeSend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargeSend

`func (o *VnicTcpOffloadSettings) SetLargeSend(v bool)`

SetLargeSend sets LargeSend field to given value.

### HasLargeSend

`func (o *VnicTcpOffloadSettings) HasLargeSend() bool`

HasLargeSend returns a boolean if a field has been set.

### GetRxChecksum

`func (o *VnicTcpOffloadSettings) GetRxChecksum() bool`

GetRxChecksum returns the RxChecksum field if non-nil, zero value otherwise.

### GetRxChecksumOk

`func (o *VnicTcpOffloadSettings) GetRxChecksumOk() (*bool, bool)`

GetRxChecksumOk returns a tuple with the RxChecksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxChecksum

`func (o *VnicTcpOffloadSettings) SetRxChecksum(v bool)`

SetRxChecksum sets RxChecksum field to given value.

### HasRxChecksum

`func (o *VnicTcpOffloadSettings) HasRxChecksum() bool`

HasRxChecksum returns a boolean if a field has been set.

### GetTxChecksum

`func (o *VnicTcpOffloadSettings) GetTxChecksum() bool`

GetTxChecksum returns the TxChecksum field if non-nil, zero value otherwise.

### GetTxChecksumOk

`func (o *VnicTcpOffloadSettings) GetTxChecksumOk() (*bool, bool)`

GetTxChecksumOk returns a tuple with the TxChecksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxChecksum

`func (o *VnicTcpOffloadSettings) SetTxChecksum(v bool)`

SetTxChecksum sets TxChecksum field to given value.

### HasTxChecksum

`func (o *VnicTcpOffloadSettings) HasTxChecksum() bool`

HasTxChecksum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


