# RecoveryBackupScheduleAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExecutionTime** | Pointer to [**time.Time**](time.Time.md) | The time at which the backup is to be run on a given day. This is used when the frequency unit is daily. | [optional] 
**FrequencyUnit** | Pointer to **string** | The frequency at which the backup schedule must run. | [optional] [default to "Daily"]
**Hours** | Pointer to **int32** | The frequency, in hours, at which the backup schedule runs. | [optional] [default to 8]

## Methods

### NewRecoveryBackupScheduleAllOf

`func NewRecoveryBackupScheduleAllOf() *RecoveryBackupScheduleAllOf`

NewRecoveryBackupScheduleAllOf instantiates a new RecoveryBackupScheduleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryBackupScheduleAllOfWithDefaults

`func NewRecoveryBackupScheduleAllOfWithDefaults() *RecoveryBackupScheduleAllOf`

NewRecoveryBackupScheduleAllOfWithDefaults instantiates a new RecoveryBackupScheduleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutionTime

`func (o *RecoveryBackupScheduleAllOf) GetExecutionTime() time.Time`

GetExecutionTime returns the ExecutionTime field if non-nil, zero value otherwise.

### GetExecutionTimeOk

`func (o *RecoveryBackupScheduleAllOf) GetExecutionTimeOk() (*time.Time, bool)`

GetExecutionTimeOk returns a tuple with the ExecutionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionTime

`func (o *RecoveryBackupScheduleAllOf) SetExecutionTime(v time.Time)`

SetExecutionTime sets ExecutionTime field to given value.

### HasExecutionTime

`func (o *RecoveryBackupScheduleAllOf) HasExecutionTime() bool`

HasExecutionTime returns a boolean if a field has been set.

### GetFrequencyUnit

`func (o *RecoveryBackupScheduleAllOf) GetFrequencyUnit() string`

GetFrequencyUnit returns the FrequencyUnit field if non-nil, zero value otherwise.

### GetFrequencyUnitOk

`func (o *RecoveryBackupScheduleAllOf) GetFrequencyUnitOk() (*string, bool)`

GetFrequencyUnitOk returns a tuple with the FrequencyUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyUnit

`func (o *RecoveryBackupScheduleAllOf) SetFrequencyUnit(v string)`

SetFrequencyUnit sets FrequencyUnit field to given value.

### HasFrequencyUnit

`func (o *RecoveryBackupScheduleAllOf) HasFrequencyUnit() bool`

HasFrequencyUnit returns a boolean if a field has been set.

### GetHours

`func (o *RecoveryBackupScheduleAllOf) GetHours() int32`

GetHours returns the Hours field if non-nil, zero value otherwise.

### GetHoursOk

`func (o *RecoveryBackupScheduleAllOf) GetHoursOk() (*int32, bool)`

GetHoursOk returns a tuple with the Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHours

`func (o *RecoveryBackupScheduleAllOf) SetHours(v int32)`

SetHours sets Hours field to given value.

### HasHours

`func (o *RecoveryBackupScheduleAllOf) HasHours() bool`

HasHours returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


