# OnpremUpgradePhase

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

### NewOnpremUpgradePhase

`func NewOnpremUpgradePhase() *OnpremUpgradePhase`

NewOnpremUpgradePhase instantiates a new OnpremUpgradePhase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnpremUpgradePhaseWithDefaults

`func NewOnpremUpgradePhaseWithDefaults() *OnpremUpgradePhase`

NewOnpremUpgradePhaseWithDefaults instantiates a new OnpremUpgradePhase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElapsedTime

`func (o *OnpremUpgradePhase) GetElapsedTime() int64`

GetElapsedTime returns the ElapsedTime field if non-nil, zero value otherwise.

### GetElapsedTimeOk

`func (o *OnpremUpgradePhase) GetElapsedTimeOk() (*int64, bool)`

GetElapsedTimeOk returns a tuple with the ElapsedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsedTime

`func (o *OnpremUpgradePhase) SetElapsedTime(v int64)`

SetElapsedTime sets ElapsedTime field to given value.

### HasElapsedTime

`func (o *OnpremUpgradePhase) HasElapsedTime() bool`

HasElapsedTime returns a boolean if a field has been set.

### GetEndTime

`func (o *OnpremUpgradePhase) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *OnpremUpgradePhase) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *OnpremUpgradePhase) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *OnpremUpgradePhase) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetFailed

`func (o *OnpremUpgradePhase) GetFailed() bool`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *OnpremUpgradePhase) GetFailedOk() (*bool, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *OnpremUpgradePhase) SetFailed(v bool)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *OnpremUpgradePhase) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetMessage

`func (o *OnpremUpgradePhase) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *OnpremUpgradePhase) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *OnpremUpgradePhase) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *OnpremUpgradePhase) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetName

`func (o *OnpremUpgradePhase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OnpremUpgradePhase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OnpremUpgradePhase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OnpremUpgradePhase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStartTime

`func (o *OnpremUpgradePhase) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *OnpremUpgradePhase) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *OnpremUpgradePhase) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *OnpremUpgradePhase) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


