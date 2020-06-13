# AssetClusterMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Leadership** | Pointer to **string** | The current leadershipstate of this member. Updated by the device connector on failover or leadership change. If a member is elected as Primary within the cluster this connection will be the same as the DeviceRegistration connection. E.g a message addressed to the DeviceRegistration will be forwarded to the ClusterMember connection. | [optional] [readonly] [default to "Unknown"]
**LockedLeader** | Pointer to **bool** | Devices lock their leadership on failure to heartbeat with peers. Value acts as a third party tie breaker in election between nodes. Intersight enforces that only one cluster member exists with this value set to true. | [optional] 
**MemberIdentity** | Pointer to **string** | The unique identity of the member within the cluster. The identity is retrieved from the platform and reported by the device connector at connection time. | [optional] [readonly] 
**ParentClusterMemberIdentity** | Pointer to **string** | The member idenity of the cluster member through which this device is connected if applicable. | [optional] [readonly] 
**Sudi** | Pointer to [**AssetSudiInfo**](asset.SudiInfo.md) |  | [optional] 
**Device** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewAssetClusterMember

`func NewAssetClusterMember() *AssetClusterMember`

NewAssetClusterMember instantiates a new AssetClusterMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetClusterMemberWithDefaults

`func NewAssetClusterMemberWithDefaults() *AssetClusterMember`

NewAssetClusterMemberWithDefaults instantiates a new AssetClusterMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLeadership

`func (o *AssetClusterMember) GetLeadership() string`

GetLeadership returns the Leadership field if non-nil, zero value otherwise.

### GetLeadershipOk

`func (o *AssetClusterMember) GetLeadershipOk() (*string, bool)`

GetLeadershipOk returns a tuple with the Leadership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadership

`func (o *AssetClusterMember) SetLeadership(v string)`

SetLeadership sets Leadership field to given value.

### HasLeadership

`func (o *AssetClusterMember) HasLeadership() bool`

HasLeadership returns a boolean if a field has been set.

### GetLockedLeader

`func (o *AssetClusterMember) GetLockedLeader() bool`

GetLockedLeader returns the LockedLeader field if non-nil, zero value otherwise.

### GetLockedLeaderOk

`func (o *AssetClusterMember) GetLockedLeaderOk() (*bool, bool)`

GetLockedLeaderOk returns a tuple with the LockedLeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedLeader

`func (o *AssetClusterMember) SetLockedLeader(v bool)`

SetLockedLeader sets LockedLeader field to given value.

### HasLockedLeader

`func (o *AssetClusterMember) HasLockedLeader() bool`

HasLockedLeader returns a boolean if a field has been set.

### GetMemberIdentity

`func (o *AssetClusterMember) GetMemberIdentity() string`

GetMemberIdentity returns the MemberIdentity field if non-nil, zero value otherwise.

### GetMemberIdentityOk

`func (o *AssetClusterMember) GetMemberIdentityOk() (*string, bool)`

GetMemberIdentityOk returns a tuple with the MemberIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberIdentity

`func (o *AssetClusterMember) SetMemberIdentity(v string)`

SetMemberIdentity sets MemberIdentity field to given value.

### HasMemberIdentity

`func (o *AssetClusterMember) HasMemberIdentity() bool`

HasMemberIdentity returns a boolean if a field has been set.

### GetParentClusterMemberIdentity

`func (o *AssetClusterMember) GetParentClusterMemberIdentity() string`

GetParentClusterMemberIdentity returns the ParentClusterMemberIdentity field if non-nil, zero value otherwise.

### GetParentClusterMemberIdentityOk

`func (o *AssetClusterMember) GetParentClusterMemberIdentityOk() (*string, bool)`

GetParentClusterMemberIdentityOk returns a tuple with the ParentClusterMemberIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentClusterMemberIdentity

`func (o *AssetClusterMember) SetParentClusterMemberIdentity(v string)`

SetParentClusterMemberIdentity sets ParentClusterMemberIdentity field to given value.

### HasParentClusterMemberIdentity

`func (o *AssetClusterMember) HasParentClusterMemberIdentity() bool`

HasParentClusterMemberIdentity returns a boolean if a field has been set.

### GetSudi

`func (o *AssetClusterMember) GetSudi() AssetSudiInfo`

GetSudi returns the Sudi field if non-nil, zero value otherwise.

### GetSudiOk

`func (o *AssetClusterMember) GetSudiOk() (*AssetSudiInfo, bool)`

GetSudiOk returns a tuple with the Sudi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSudi

`func (o *AssetClusterMember) SetSudi(v AssetSudiInfo)`

SetSudi sets Sudi field to given value.

### HasSudi

`func (o *AssetClusterMember) HasSudi() bool`

HasSudi returns a boolean if a field has been set.

### GetDevice

`func (o *AssetClusterMember) GetDevice() AssetDeviceRegistrationRelationship`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *AssetClusterMember) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *AssetClusterMember) SetDevice(v AssetDeviceRegistrationRelationship)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *AssetClusterMember) HasDevice() bool`

HasDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


