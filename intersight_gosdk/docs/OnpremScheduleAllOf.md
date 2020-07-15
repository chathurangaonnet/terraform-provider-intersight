# OnpremScheduleAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DayOfMonth** | Pointer to **int64** | Schedule a task on a specific day of the month. Valid values are 1 through 31. If monthOfYear is specified, then dayOfMonth value must be valid for that month. DayOfMonth may not be set when dayOfWeek is specfied. | [optional] 
**DayOfWeek** | Pointer to **int64** | Schedule a task on a specific day of the week. Valid values are 1 through 7, with 1 being Sunday. DayOfWeek may not be specfied when dayOfMonth is specified. | [optional] 
**MonthOfYear** | Pointer to **int64** | Schedule a task on a specific month of the year. Valid values are 1 through 12, with 1 being January. | [optional] 
**RepeatInterval** | Pointer to **int64** | Schedule a task to run periodically at an interval. Default unit of the RepeatInterval is in minutes. If the RepeateInterval value is set, then all other properties are ignored by the scheduler. RepeateInterval constraints are enforced by the services that use the schedule. Each service has pre-configured service specific properties for enforcing minimum and maximum values of the RepeatInterval. | [optional] 
**TimeOfDay** | Pointer to **int64** | Time of the day in seconds. TimeOfDay is required for all schedule configurations, except when the RepeateInterval field is specified. | [optional] 
**TimeZone** | Pointer to **string** | Timezone to use for the schedule calculation. If a timezone value is not speficied, then the schedule calculation will be based on UTC. | [optional] [default to "Pacific/Niue"]
**WeekOfMonth** | Pointer to **int64** | Schedule a task on a specific week of the month. Valid values are 1 through 5. Value of 5 means last week of the month. WeekOfMonth may not be set when dayOfMonth is specified. | [optional] 

## Methods

### NewOnpremScheduleAllOf

`func NewOnpremScheduleAllOf() *OnpremScheduleAllOf`

NewOnpremScheduleAllOf instantiates a new OnpremScheduleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnpremScheduleAllOfWithDefaults

`func NewOnpremScheduleAllOfWithDefaults() *OnpremScheduleAllOf`

NewOnpremScheduleAllOfWithDefaults instantiates a new OnpremScheduleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDayOfMonth

`func (o *OnpremScheduleAllOf) GetDayOfMonth() int64`

GetDayOfMonth returns the DayOfMonth field if non-nil, zero value otherwise.

### GetDayOfMonthOk

`func (o *OnpremScheduleAllOf) GetDayOfMonthOk() (*int64, bool)`

GetDayOfMonthOk returns a tuple with the DayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfMonth

`func (o *OnpremScheduleAllOf) SetDayOfMonth(v int64)`

SetDayOfMonth sets DayOfMonth field to given value.

### HasDayOfMonth

`func (o *OnpremScheduleAllOf) HasDayOfMonth() bool`

HasDayOfMonth returns a boolean if a field has been set.

### GetDayOfWeek

`func (o *OnpremScheduleAllOf) GetDayOfWeek() int64`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *OnpremScheduleAllOf) GetDayOfWeekOk() (*int64, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *OnpremScheduleAllOf) SetDayOfWeek(v int64)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *OnpremScheduleAllOf) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### GetMonthOfYear

`func (o *OnpremScheduleAllOf) GetMonthOfYear() int64`

GetMonthOfYear returns the MonthOfYear field if non-nil, zero value otherwise.

### GetMonthOfYearOk

`func (o *OnpremScheduleAllOf) GetMonthOfYearOk() (*int64, bool)`

GetMonthOfYearOk returns a tuple with the MonthOfYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthOfYear

`func (o *OnpremScheduleAllOf) SetMonthOfYear(v int64)`

SetMonthOfYear sets MonthOfYear field to given value.

### HasMonthOfYear

`func (o *OnpremScheduleAllOf) HasMonthOfYear() bool`

HasMonthOfYear returns a boolean if a field has been set.

### GetRepeatInterval

`func (o *OnpremScheduleAllOf) GetRepeatInterval() int64`

GetRepeatInterval returns the RepeatInterval field if non-nil, zero value otherwise.

### GetRepeatIntervalOk

`func (o *OnpremScheduleAllOf) GetRepeatIntervalOk() (*int64, bool)`

GetRepeatIntervalOk returns a tuple with the RepeatInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatInterval

`func (o *OnpremScheduleAllOf) SetRepeatInterval(v int64)`

SetRepeatInterval sets RepeatInterval field to given value.

### HasRepeatInterval

`func (o *OnpremScheduleAllOf) HasRepeatInterval() bool`

HasRepeatInterval returns a boolean if a field has been set.

### GetTimeOfDay

`func (o *OnpremScheduleAllOf) GetTimeOfDay() int64`

GetTimeOfDay returns the TimeOfDay field if non-nil, zero value otherwise.

### GetTimeOfDayOk

`func (o *OnpremScheduleAllOf) GetTimeOfDayOk() (*int64, bool)`

GetTimeOfDayOk returns a tuple with the TimeOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOfDay

`func (o *OnpremScheduleAllOf) SetTimeOfDay(v int64)`

SetTimeOfDay sets TimeOfDay field to given value.

### HasTimeOfDay

`func (o *OnpremScheduleAllOf) HasTimeOfDay() bool`

HasTimeOfDay returns a boolean if a field has been set.

### GetTimeZone

`func (o *OnpremScheduleAllOf) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *OnpremScheduleAllOf) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *OnpremScheduleAllOf) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *OnpremScheduleAllOf) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetWeekOfMonth

`func (o *OnpremScheduleAllOf) GetWeekOfMonth() int64`

GetWeekOfMonth returns the WeekOfMonth field if non-nil, zero value otherwise.

### GetWeekOfMonthOk

`func (o *OnpremScheduleAllOf) GetWeekOfMonthOk() (*int64, bool)`

GetWeekOfMonthOk returns a tuple with the WeekOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekOfMonth

`func (o *OnpremScheduleAllOf) SetWeekOfMonth(v int64)`

SetWeekOfMonth sets WeekOfMonth field to given value.

### HasWeekOfMonth

`func (o *OnpremScheduleAllOf) HasWeekOfMonth() bool`

HasWeekOfMonth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


