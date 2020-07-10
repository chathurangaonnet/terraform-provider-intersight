# HyperflexHxapHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Age** | Pointer to **string** | Denotes age or life time of the Host in nano seconds. | [optional] 
**ClusterUuid** | Pointer to **string** | The UUID of the cluster to which this Host belongs to. | [optional] 
**FailureReason** | Pointer to **string** | Reason of the failure when host is in failed state. | [optional] 
**HwPowerState** | Pointer to **string** | Is the host Powered-up or Powered-down. | [optional] [default to "Unknown"]
**InternalIpAddress** | Pointer to **string** | Internal IP Address of the Host. | [optional] 
**ManagementIpAddress** | Pointer to **string** | Management IP Address of the Host. | [optional] 
**MasterRole** | Pointer to **bool** | Is the role of this host is master in the cluster? true or false. | [optional] 
**Version** | Pointer to **string** | Product version of the Host. | [optional] 
**Cluster** | Pointer to [**HyperflexHxapClusterRelationship**](hyperflex.HxapCluster.Relationship.md) |  | [optional] 
**ClusterMember** | Pointer to [**AssetClusterMemberRelationship**](asset.ClusterMember.Relationship.md) |  | [optional] 
**PhysicalServer** | Pointer to [**ComputePhysicalRelationship**](compute.Physical.Relationship.md) |  | [optional] 

## Methods

### NewHyperflexHxapHost

`func NewHyperflexHxapHost() *HyperflexHxapHost`

NewHyperflexHxapHost instantiates a new HyperflexHxapHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHxapHostWithDefaults

`func NewHyperflexHxapHostWithDefaults() *HyperflexHxapHost`

NewHyperflexHxapHostWithDefaults instantiates a new HyperflexHxapHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAge

`func (o *HyperflexHxapHost) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *HyperflexHxapHost) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *HyperflexHxapHost) SetAge(v string)`

SetAge sets Age field to given value.

### HasAge

`func (o *HyperflexHxapHost) HasAge() bool`

HasAge returns a boolean if a field has been set.

### GetClusterUuid

`func (o *HyperflexHxapHost) GetClusterUuid() string`

GetClusterUuid returns the ClusterUuid field if non-nil, zero value otherwise.

### GetClusterUuidOk

`func (o *HyperflexHxapHost) GetClusterUuidOk() (*string, bool)`

GetClusterUuidOk returns a tuple with the ClusterUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterUuid

`func (o *HyperflexHxapHost) SetClusterUuid(v string)`

SetClusterUuid sets ClusterUuid field to given value.

### HasClusterUuid

`func (o *HyperflexHxapHost) HasClusterUuid() bool`

HasClusterUuid returns a boolean if a field has been set.

### GetFailureReason

`func (o *HyperflexHxapHost) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *HyperflexHxapHost) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *HyperflexHxapHost) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *HyperflexHxapHost) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetHwPowerState

`func (o *HyperflexHxapHost) GetHwPowerState() string`

GetHwPowerState returns the HwPowerState field if non-nil, zero value otherwise.

### GetHwPowerStateOk

`func (o *HyperflexHxapHost) GetHwPowerStateOk() (*string, bool)`

GetHwPowerStateOk returns a tuple with the HwPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwPowerState

`func (o *HyperflexHxapHost) SetHwPowerState(v string)`

SetHwPowerState sets HwPowerState field to given value.

### HasHwPowerState

`func (o *HyperflexHxapHost) HasHwPowerState() bool`

HasHwPowerState returns a boolean if a field has been set.

### GetInternalIpAddress

`func (o *HyperflexHxapHost) GetInternalIpAddress() string`

GetInternalIpAddress returns the InternalIpAddress field if non-nil, zero value otherwise.

### GetInternalIpAddressOk

`func (o *HyperflexHxapHost) GetInternalIpAddressOk() (*string, bool)`

GetInternalIpAddressOk returns a tuple with the InternalIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIpAddress

`func (o *HyperflexHxapHost) SetInternalIpAddress(v string)`

SetInternalIpAddress sets InternalIpAddress field to given value.

### HasInternalIpAddress

`func (o *HyperflexHxapHost) HasInternalIpAddress() bool`

HasInternalIpAddress returns a boolean if a field has been set.

### GetManagementIpAddress

`func (o *HyperflexHxapHost) GetManagementIpAddress() string`

GetManagementIpAddress returns the ManagementIpAddress field if non-nil, zero value otherwise.

### GetManagementIpAddressOk

`func (o *HyperflexHxapHost) GetManagementIpAddressOk() (*string, bool)`

GetManagementIpAddressOk returns a tuple with the ManagementIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementIpAddress

`func (o *HyperflexHxapHost) SetManagementIpAddress(v string)`

SetManagementIpAddress sets ManagementIpAddress field to given value.

### HasManagementIpAddress

`func (o *HyperflexHxapHost) HasManagementIpAddress() bool`

HasManagementIpAddress returns a boolean if a field has been set.

### GetMasterRole

`func (o *HyperflexHxapHost) GetMasterRole() bool`

GetMasterRole returns the MasterRole field if non-nil, zero value otherwise.

### GetMasterRoleOk

`func (o *HyperflexHxapHost) GetMasterRoleOk() (*bool, bool)`

GetMasterRoleOk returns a tuple with the MasterRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterRole

`func (o *HyperflexHxapHost) SetMasterRole(v bool)`

SetMasterRole sets MasterRole field to given value.

### HasMasterRole

`func (o *HyperflexHxapHost) HasMasterRole() bool`

HasMasterRole returns a boolean if a field has been set.

### GetVersion

`func (o *HyperflexHxapHost) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HyperflexHxapHost) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HyperflexHxapHost) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HyperflexHxapHost) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCluster

`func (o *HyperflexHxapHost) GetCluster() HyperflexHxapClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HyperflexHxapHost) GetClusterOk() (*HyperflexHxapClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HyperflexHxapHost) SetCluster(v HyperflexHxapClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HyperflexHxapHost) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetClusterMember

`func (o *HyperflexHxapHost) GetClusterMember() AssetClusterMemberRelationship`

GetClusterMember returns the ClusterMember field if non-nil, zero value otherwise.

### GetClusterMemberOk

`func (o *HyperflexHxapHost) GetClusterMemberOk() (*AssetClusterMemberRelationship, bool)`

GetClusterMemberOk returns a tuple with the ClusterMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterMember

`func (o *HyperflexHxapHost) SetClusterMember(v AssetClusterMemberRelationship)`

SetClusterMember sets ClusterMember field to given value.

### HasClusterMember

`func (o *HyperflexHxapHost) HasClusterMember() bool`

HasClusterMember returns a boolean if a field has been set.

### GetPhysicalServer

`func (o *HyperflexHxapHost) GetPhysicalServer() ComputePhysicalRelationship`

GetPhysicalServer returns the PhysicalServer field if non-nil, zero value otherwise.

### GetPhysicalServerOk

`func (o *HyperflexHxapHost) GetPhysicalServerOk() (*ComputePhysicalRelationship, bool)`

GetPhysicalServerOk returns a tuple with the PhysicalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalServer

`func (o *HyperflexHxapHost) SetPhysicalServer(v ComputePhysicalRelationship)`

SetPhysicalServer sets PhysicalServer field to given value.

### HasPhysicalServer

`func (o *HyperflexHxapHost) HasPhysicalServer() bool`

HasPhysicalServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


