# TamAdvisoryInstanceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedObjectMoid** | Pointer to **string** | Moid of the Intersight MO affected by the alert. Deprecated now and will be removed in subsequent releases. | [optional] 
**AffectedObjectType** | Pointer to **string** | Object type of the Intersight MO affected by the alert. Deprecated now and will be removed in subsequent releases. | [optional] 
**LastStateChangeTime** | Pointer to [**time.Time**](time.Time.md) | Timestamp when a state change was observed on this advisory instnace. | [optional] [readonly] 
**LastVerifiedTime** | Pointer to [**time.Time**](time.Time.md) | Timestamp when this advisory was last evaluated. | [optional] [readonly] 
**State** | Pointer to **string** | Current state of the advisory instance (Active/Cleared/Unknown etc.). | [optional] [default to "unknown"]
**Advisory** | Pointer to [**TamBaseAdvisoryRelationship**](tam.BaseAdvisory.Relationship.md) |  | [optional] 
**AffectedObject** | Pointer to [**MoBaseMoRelationship**](mo.BaseMo.Relationship.md) |  | [optional] 
**DeviceRegistration** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewTamAdvisoryInstanceAllOf

`func NewTamAdvisoryInstanceAllOf() *TamAdvisoryInstanceAllOf`

NewTamAdvisoryInstanceAllOf instantiates a new TamAdvisoryInstanceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTamAdvisoryInstanceAllOfWithDefaults

`func NewTamAdvisoryInstanceAllOfWithDefaults() *TamAdvisoryInstanceAllOf`

NewTamAdvisoryInstanceAllOfWithDefaults instantiates a new TamAdvisoryInstanceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedObjectMoid

`func (o *TamAdvisoryInstanceAllOf) GetAffectedObjectMoid() string`

GetAffectedObjectMoid returns the AffectedObjectMoid field if non-nil, zero value otherwise.

### GetAffectedObjectMoidOk

`func (o *TamAdvisoryInstanceAllOf) GetAffectedObjectMoidOk() (*string, bool)`

GetAffectedObjectMoidOk returns a tuple with the AffectedObjectMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedObjectMoid

`func (o *TamAdvisoryInstanceAllOf) SetAffectedObjectMoid(v string)`

SetAffectedObjectMoid sets AffectedObjectMoid field to given value.

### HasAffectedObjectMoid

`func (o *TamAdvisoryInstanceAllOf) HasAffectedObjectMoid() bool`

HasAffectedObjectMoid returns a boolean if a field has been set.

### GetAffectedObjectType

`func (o *TamAdvisoryInstanceAllOf) GetAffectedObjectType() string`

GetAffectedObjectType returns the AffectedObjectType field if non-nil, zero value otherwise.

### GetAffectedObjectTypeOk

`func (o *TamAdvisoryInstanceAllOf) GetAffectedObjectTypeOk() (*string, bool)`

GetAffectedObjectTypeOk returns a tuple with the AffectedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedObjectType

`func (o *TamAdvisoryInstanceAllOf) SetAffectedObjectType(v string)`

SetAffectedObjectType sets AffectedObjectType field to given value.

### HasAffectedObjectType

`func (o *TamAdvisoryInstanceAllOf) HasAffectedObjectType() bool`

HasAffectedObjectType returns a boolean if a field has been set.

### GetLastStateChangeTime

`func (o *TamAdvisoryInstanceAllOf) GetLastStateChangeTime() time.Time`

GetLastStateChangeTime returns the LastStateChangeTime field if non-nil, zero value otherwise.

### GetLastStateChangeTimeOk

`func (o *TamAdvisoryInstanceAllOf) GetLastStateChangeTimeOk() (*time.Time, bool)`

GetLastStateChangeTimeOk returns a tuple with the LastStateChangeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStateChangeTime

`func (o *TamAdvisoryInstanceAllOf) SetLastStateChangeTime(v time.Time)`

SetLastStateChangeTime sets LastStateChangeTime field to given value.

### HasLastStateChangeTime

`func (o *TamAdvisoryInstanceAllOf) HasLastStateChangeTime() bool`

HasLastStateChangeTime returns a boolean if a field has been set.

### GetLastVerifiedTime

`func (o *TamAdvisoryInstanceAllOf) GetLastVerifiedTime() time.Time`

GetLastVerifiedTime returns the LastVerifiedTime field if non-nil, zero value otherwise.

### GetLastVerifiedTimeOk

`func (o *TamAdvisoryInstanceAllOf) GetLastVerifiedTimeOk() (*time.Time, bool)`

GetLastVerifiedTimeOk returns a tuple with the LastVerifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastVerifiedTime

`func (o *TamAdvisoryInstanceAllOf) SetLastVerifiedTime(v time.Time)`

SetLastVerifiedTime sets LastVerifiedTime field to given value.

### HasLastVerifiedTime

`func (o *TamAdvisoryInstanceAllOf) HasLastVerifiedTime() bool`

HasLastVerifiedTime returns a boolean if a field has been set.

### GetState

`func (o *TamAdvisoryInstanceAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TamAdvisoryInstanceAllOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TamAdvisoryInstanceAllOf) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *TamAdvisoryInstanceAllOf) HasState() bool`

HasState returns a boolean if a field has been set.

### GetAdvisory

`func (o *TamAdvisoryInstanceAllOf) GetAdvisory() TamBaseAdvisoryRelationship`

GetAdvisory returns the Advisory field if non-nil, zero value otherwise.

### GetAdvisoryOk

`func (o *TamAdvisoryInstanceAllOf) GetAdvisoryOk() (*TamBaseAdvisoryRelationship, bool)`

GetAdvisoryOk returns a tuple with the Advisory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvisory

`func (o *TamAdvisoryInstanceAllOf) SetAdvisory(v TamBaseAdvisoryRelationship)`

SetAdvisory sets Advisory field to given value.

### HasAdvisory

`func (o *TamAdvisoryInstanceAllOf) HasAdvisory() bool`

HasAdvisory returns a boolean if a field has been set.

### GetAffectedObject

`func (o *TamAdvisoryInstanceAllOf) GetAffectedObject() MoBaseMoRelationship`

GetAffectedObject returns the AffectedObject field if non-nil, zero value otherwise.

### GetAffectedObjectOk

`func (o *TamAdvisoryInstanceAllOf) GetAffectedObjectOk() (*MoBaseMoRelationship, bool)`

GetAffectedObjectOk returns a tuple with the AffectedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedObject

`func (o *TamAdvisoryInstanceAllOf) SetAffectedObject(v MoBaseMoRelationship)`

SetAffectedObject sets AffectedObject field to given value.

### HasAffectedObject

`func (o *TamAdvisoryInstanceAllOf) HasAffectedObject() bool`

HasAffectedObject returns a boolean if a field has been set.

### GetDeviceRegistration

`func (o *TamAdvisoryInstanceAllOf) GetDeviceRegistration() AssetDeviceRegistrationRelationship`

GetDeviceRegistration returns the DeviceRegistration field if non-nil, zero value otherwise.

### GetDeviceRegistrationOk

`func (o *TamAdvisoryInstanceAllOf) GetDeviceRegistrationOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceRegistrationOk returns a tuple with the DeviceRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRegistration

`func (o *TamAdvisoryInstanceAllOf) SetDeviceRegistration(v AssetDeviceRegistrationRelationship)`

SetDeviceRegistration sets DeviceRegistration field to given value.

### HasDeviceRegistration

`func (o *TamAdvisoryInstanceAllOf) HasDeviceRegistration() bool`

HasDeviceRegistration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


