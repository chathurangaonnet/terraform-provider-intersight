# CondAlarm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acknowledge** | Pointer to **string** | Alarm acknowledgment state. Default value is None. | [optional] [default to "None"]
**AcknowledgeBy** | Pointer to **string** | User who acknowledged the alarm. | [optional] [readonly] 
**AcknowledgeTime** | Pointer to [**time.Time**](time.Time.md) | Time at which the alarm was acknowledged by the user. | [optional] [readonly] 
**AffectedMoId** | Pointer to **string** | MoId of the affected object from the managed system&#39;s point of view. | [optional] [readonly] 
**AffectedMoType** | Pointer to **string** | Managed system affected object type. For example Adaptor, FI, CIMC. | [optional] [readonly] 
**AffectedObject** | Pointer to **string** | A unique key for an alarm instance, consists of a combination of a unique system name and msAffectedObject. | [optional] [readonly] 
**AncestorMoId** | Pointer to **string** | Parent MoId of the fault from managed system. For example, Blade moid for adaptor fault. | [optional] [readonly] 
**AncestorMoType** | Pointer to **string** | Parent MO type of the fault from managed system. For example, Blade for adaptor fault. | [optional] [readonly] 
**Code** | Pointer to **string** | A unique alarm code. For alarms mapped from UCS faults, this will be the same as the UCS fault code. | [optional] [readonly] 
**CreationTime** | Pointer to [**time.Time**](time.Time.md) | The time the alarm was created. | [optional] [readonly] 
**Description** | Pointer to **string** | A longer description of the alarm than the name. The description contains details of the component reporting the issue. | [optional] [readonly] 
**LastTransitionTime** | Pointer to [**time.Time**](time.Time.md) | The time the alarm last had a change in severity. | [optional] [readonly] 
**MsAffectedObject** | Pointer to **string** | A unique key for the alarm from the managed system&#39;s point of view. For example, in the case of UCS, this is the fault&#39;s dn. | [optional] [readonly] 
**Name** | Pointer to **string** | Uniquely identifies the type of alarm. For alarms originating from Intersight, this will be a descriptive name. For alarms that are mapped from faults, the name will be derived from fault properties. For example, alarms mapped from UCS faults will use a prefix of UCS and appended with the fault code. | [optional] [readonly] 
**OrigSeverity** | Pointer to **string** | The original severity when the alarm was first created. | [optional] [readonly] [default to "None"]
**Severity** | Pointer to **string** | The severity of the alarm. Valid values are Critical, Warning, Info, and Cleared. | [optional] [readonly] [default to "None"]
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewCondAlarm

`func NewCondAlarm() *CondAlarm`

NewCondAlarm instantiates a new CondAlarm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCondAlarmWithDefaults

`func NewCondAlarmWithDefaults() *CondAlarm`

NewCondAlarmWithDefaults instantiates a new CondAlarm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcknowledge

`func (o *CondAlarm) GetAcknowledge() string`

GetAcknowledge returns the Acknowledge field if non-nil, zero value otherwise.

### GetAcknowledgeOk

`func (o *CondAlarm) GetAcknowledgeOk() (*string, bool)`

GetAcknowledgeOk returns a tuple with the Acknowledge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledge

`func (o *CondAlarm) SetAcknowledge(v string)`

SetAcknowledge sets Acknowledge field to given value.

### HasAcknowledge

`func (o *CondAlarm) HasAcknowledge() bool`

HasAcknowledge returns a boolean if a field has been set.

### GetAcknowledgeBy

`func (o *CondAlarm) GetAcknowledgeBy() string`

GetAcknowledgeBy returns the AcknowledgeBy field if non-nil, zero value otherwise.

### GetAcknowledgeByOk

`func (o *CondAlarm) GetAcknowledgeByOk() (*string, bool)`

GetAcknowledgeByOk returns a tuple with the AcknowledgeBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgeBy

`func (o *CondAlarm) SetAcknowledgeBy(v string)`

SetAcknowledgeBy sets AcknowledgeBy field to given value.

### HasAcknowledgeBy

`func (o *CondAlarm) HasAcknowledgeBy() bool`

HasAcknowledgeBy returns a boolean if a field has been set.

### GetAcknowledgeTime

`func (o *CondAlarm) GetAcknowledgeTime() time.Time`

GetAcknowledgeTime returns the AcknowledgeTime field if non-nil, zero value otherwise.

### GetAcknowledgeTimeOk

`func (o *CondAlarm) GetAcknowledgeTimeOk() (*time.Time, bool)`

GetAcknowledgeTimeOk returns a tuple with the AcknowledgeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgeTime

`func (o *CondAlarm) SetAcknowledgeTime(v time.Time)`

SetAcknowledgeTime sets AcknowledgeTime field to given value.

### HasAcknowledgeTime

`func (o *CondAlarm) HasAcknowledgeTime() bool`

HasAcknowledgeTime returns a boolean if a field has been set.

### GetAffectedMoId

`func (o *CondAlarm) GetAffectedMoId() string`

GetAffectedMoId returns the AffectedMoId field if non-nil, zero value otherwise.

### GetAffectedMoIdOk

`func (o *CondAlarm) GetAffectedMoIdOk() (*string, bool)`

GetAffectedMoIdOk returns a tuple with the AffectedMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedMoId

`func (o *CondAlarm) SetAffectedMoId(v string)`

SetAffectedMoId sets AffectedMoId field to given value.

### HasAffectedMoId

`func (o *CondAlarm) HasAffectedMoId() bool`

HasAffectedMoId returns a boolean if a field has been set.

### GetAffectedMoType

`func (o *CondAlarm) GetAffectedMoType() string`

GetAffectedMoType returns the AffectedMoType field if non-nil, zero value otherwise.

### GetAffectedMoTypeOk

`func (o *CondAlarm) GetAffectedMoTypeOk() (*string, bool)`

GetAffectedMoTypeOk returns a tuple with the AffectedMoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedMoType

`func (o *CondAlarm) SetAffectedMoType(v string)`

SetAffectedMoType sets AffectedMoType field to given value.

### HasAffectedMoType

`func (o *CondAlarm) HasAffectedMoType() bool`

HasAffectedMoType returns a boolean if a field has been set.

### GetAffectedObject

`func (o *CondAlarm) GetAffectedObject() string`

GetAffectedObject returns the AffectedObject field if non-nil, zero value otherwise.

### GetAffectedObjectOk

`func (o *CondAlarm) GetAffectedObjectOk() (*string, bool)`

GetAffectedObjectOk returns a tuple with the AffectedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedObject

`func (o *CondAlarm) SetAffectedObject(v string)`

SetAffectedObject sets AffectedObject field to given value.

### HasAffectedObject

`func (o *CondAlarm) HasAffectedObject() bool`

HasAffectedObject returns a boolean if a field has been set.

### GetAncestorMoId

`func (o *CondAlarm) GetAncestorMoId() string`

GetAncestorMoId returns the AncestorMoId field if non-nil, zero value otherwise.

### GetAncestorMoIdOk

`func (o *CondAlarm) GetAncestorMoIdOk() (*string, bool)`

GetAncestorMoIdOk returns a tuple with the AncestorMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestorMoId

`func (o *CondAlarm) SetAncestorMoId(v string)`

SetAncestorMoId sets AncestorMoId field to given value.

### HasAncestorMoId

`func (o *CondAlarm) HasAncestorMoId() bool`

HasAncestorMoId returns a boolean if a field has been set.

### GetAncestorMoType

`func (o *CondAlarm) GetAncestorMoType() string`

GetAncestorMoType returns the AncestorMoType field if non-nil, zero value otherwise.

### GetAncestorMoTypeOk

`func (o *CondAlarm) GetAncestorMoTypeOk() (*string, bool)`

GetAncestorMoTypeOk returns a tuple with the AncestorMoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestorMoType

`func (o *CondAlarm) SetAncestorMoType(v string)`

SetAncestorMoType sets AncestorMoType field to given value.

### HasAncestorMoType

`func (o *CondAlarm) HasAncestorMoType() bool`

HasAncestorMoType returns a boolean if a field has been set.

### GetCode

`func (o *CondAlarm) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CondAlarm) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CondAlarm) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CondAlarm) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCreationTime

`func (o *CondAlarm) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *CondAlarm) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *CondAlarm) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *CondAlarm) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetDescription

`func (o *CondAlarm) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CondAlarm) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CondAlarm) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CondAlarm) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLastTransitionTime

`func (o *CondAlarm) GetLastTransitionTime() time.Time`

GetLastTransitionTime returns the LastTransitionTime field if non-nil, zero value otherwise.

### GetLastTransitionTimeOk

`func (o *CondAlarm) GetLastTransitionTimeOk() (*time.Time, bool)`

GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransitionTime

`func (o *CondAlarm) SetLastTransitionTime(v time.Time)`

SetLastTransitionTime sets LastTransitionTime field to given value.

### HasLastTransitionTime

`func (o *CondAlarm) HasLastTransitionTime() bool`

HasLastTransitionTime returns a boolean if a field has been set.

### GetMsAffectedObject

`func (o *CondAlarm) GetMsAffectedObject() string`

GetMsAffectedObject returns the MsAffectedObject field if non-nil, zero value otherwise.

### GetMsAffectedObjectOk

`func (o *CondAlarm) GetMsAffectedObjectOk() (*string, bool)`

GetMsAffectedObjectOk returns a tuple with the MsAffectedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAffectedObject

`func (o *CondAlarm) SetMsAffectedObject(v string)`

SetMsAffectedObject sets MsAffectedObject field to given value.

### HasMsAffectedObject

`func (o *CondAlarm) HasMsAffectedObject() bool`

HasMsAffectedObject returns a boolean if a field has been set.

### GetName

`func (o *CondAlarm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CondAlarm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CondAlarm) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CondAlarm) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrigSeverity

`func (o *CondAlarm) GetOrigSeverity() string`

GetOrigSeverity returns the OrigSeverity field if non-nil, zero value otherwise.

### GetOrigSeverityOk

`func (o *CondAlarm) GetOrigSeverityOk() (*string, bool)`

GetOrigSeverityOk returns a tuple with the OrigSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigSeverity

`func (o *CondAlarm) SetOrigSeverity(v string)`

SetOrigSeverity sets OrigSeverity field to given value.

### HasOrigSeverity

`func (o *CondAlarm) HasOrigSeverity() bool`

HasOrigSeverity returns a boolean if a field has been set.

### GetSeverity

`func (o *CondAlarm) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *CondAlarm) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *CondAlarm) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *CondAlarm) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *CondAlarm) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *CondAlarm) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *CondAlarm) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *CondAlarm) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


