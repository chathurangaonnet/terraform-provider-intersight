# OnpremUpgradePhaseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ElapsedTime** | Pointer to **int64** | Elapsed time in seconds to complete the current upgrade phase. | [optional] [readonly] 
**EndTime** | Pointer to [**time.Time**](time.Time.md) | End date of the software upgrade phase. | [optional] [readonly] 
**Failed** | Pointer to **bool** | Indicates if the upgrade phase has failed or not. | [optional] [readonly] 
**Message** | Pointer to **string** | Status message set during the upgrade phase. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the upgrade phase. | [optional] [readonly] [default to "init"]
**StartTime** | Pointer to [**time.Time**](time.Time.md) | Start date of the software upgrade phase. | [optional] [readonly] 

## Methods

### NewOnpremUpgradePhaseAllOf

`func NewOnpremUpgradePhaseAllOf() *OnpremUpgradePhaseAllOf`

NewOnpremUpgradePhaseAllOf instantiates a new OnpremUpgradePhaseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnpremUpgradePhaseAllOfWithDefaults

`func NewOnpremUpgradePhaseAllOfWithDefaults() *OnpremUpgradePhaseAllOf`

NewOnpremUpgradePhaseAllOfWithDefaults instantiates a new OnpremUpgradePhaseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElapsedTime

`func (o *OnpremUpgradePhaseAllOf) GetElapsedTime() int64`

GetElapsedTime returns the ElapsedTime field if non-nil, zero value otherwise.

### GetElapsedTimeOk

`func (o *OnpremUpgradePhaseAllOf) GetElapsedTimeOk() (*int64, bool)`

GetElapsedTimeOk returns a tuple with the ElapsedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsedTime

`func (o *OnpremUpgradePhaseAllOf) SetElapsedTime(v int64)`

SetElapsedTime sets ElapsedTime field to given value.

### HasElapsedTime

`func (o *OnpremUpgradePhaseAllOf) HasElapsedTime() bool`

HasElapsedTime returns a boolean if a field has been set.

### GetEndTime

`func (o *OnpremUpgradePhaseAllOf) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *OnpremUpgradePhaseAllOf) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *OnpremUpgradePhaseAllOf) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *OnpremUpgradePhaseAllOf) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetFailed

`func (o *OnpremUpgradePhaseAllOf) GetFailed() bool`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *OnpremUpgradePhaseAllOf) GetFailedOk() (*bool, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *OnpremUpgradePhaseAllOf) SetFailed(v bool)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *OnpremUpgradePhaseAllOf) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetMessage

`func (o *OnpremUpgradePhaseAllOf) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *OnpremUpgradePhaseAllOf) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *OnpremUpgradePhaseAllOf) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *OnpremUpgradePhaseAllOf) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetName

`func (o *OnpremUpgradePhaseAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OnpremUpgradePhaseAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OnpremUpgradePhaseAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OnpremUpgradePhaseAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStartTime

`func (o *OnpremUpgradePhaseAllOf) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *OnpremUpgradePhaseAllOf) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *OnpremUpgradePhaseAllOf) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *OnpremUpgradePhaseAllOf) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


