# HyperflexExtFcStoragePolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminState** | Pointer to **bool** | Enables or disables external FC storage configuration. | [optional] 
**ExtaTraffic** | Pointer to [**HyperflexNamedVsan**](hyperflex.NamedVsan.md) |  | [optional] 
**ExtbTraffic** | Pointer to [**HyperflexNamedVsan**](hyperflex.NamedVsan.md) |  | [optional] 
**WwxnPrefixRange** | Pointer to [**HyperflexWwxnPrefixRange**](hyperflex.WwxnPrefixRange.md) |  | [optional] 
**ClusterProfiles** | Pointer to [**[]HyperflexClusterProfileRelationship**](hyperflex.ClusterProfile.Relationship.md) | An array of relationships to hyperflexClusterProfile resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 

## Methods

### NewHyperflexExtFcStoragePolicyAllOf

`func NewHyperflexExtFcStoragePolicyAllOf() *HyperflexExtFcStoragePolicyAllOf`

NewHyperflexExtFcStoragePolicyAllOf instantiates a new HyperflexExtFcStoragePolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexExtFcStoragePolicyAllOfWithDefaults

`func NewHyperflexExtFcStoragePolicyAllOfWithDefaults() *HyperflexExtFcStoragePolicyAllOf`

NewHyperflexExtFcStoragePolicyAllOfWithDefaults instantiates a new HyperflexExtFcStoragePolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminState

`func (o *HyperflexExtFcStoragePolicyAllOf) GetAdminState() bool`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *HyperflexExtFcStoragePolicyAllOf) GetAdminStateOk() (*bool, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *HyperflexExtFcStoragePolicyAllOf) SetAdminState(v bool)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *HyperflexExtFcStoragePolicyAllOf) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetExtaTraffic

`func (o *HyperflexExtFcStoragePolicyAllOf) GetExtaTraffic() HyperflexNamedVsan`

GetExtaTraffic returns the ExtaTraffic field if non-nil, zero value otherwise.

### GetExtaTrafficOk

`func (o *HyperflexExtFcStoragePolicyAllOf) GetExtaTrafficOk() (*HyperflexNamedVsan, bool)`

GetExtaTrafficOk returns a tuple with the ExtaTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtaTraffic

`func (o *HyperflexExtFcStoragePolicyAllOf) SetExtaTraffic(v HyperflexNamedVsan)`

SetExtaTraffic sets ExtaTraffic field to given value.

### HasExtaTraffic

`func (o *HyperflexExtFcStoragePolicyAllOf) HasExtaTraffic() bool`

HasExtaTraffic returns a boolean if a field has been set.

### GetExtbTraffic

`func (o *HyperflexExtFcStoragePolicyAllOf) GetExtbTraffic() HyperflexNamedVsan`

GetExtbTraffic returns the ExtbTraffic field if non-nil, zero value otherwise.

### GetExtbTrafficOk

`func (o *HyperflexExtFcStoragePolicyAllOf) GetExtbTrafficOk() (*HyperflexNamedVsan, bool)`

GetExtbTrafficOk returns a tuple with the ExtbTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtbTraffic

`func (o *HyperflexExtFcStoragePolicyAllOf) SetExtbTraffic(v HyperflexNamedVsan)`

SetExtbTraffic sets ExtbTraffic field to given value.

### HasExtbTraffic

`func (o *HyperflexExtFcStoragePolicyAllOf) HasExtbTraffic() bool`

HasExtbTraffic returns a boolean if a field has been set.

### GetWwxnPrefixRange

`func (o *HyperflexExtFcStoragePolicyAllOf) GetWwxnPrefixRange() HyperflexWwxnPrefixRange`

GetWwxnPrefixRange returns the WwxnPrefixRange field if non-nil, zero value otherwise.

### GetWwxnPrefixRangeOk

`func (o *HyperflexExtFcStoragePolicyAllOf) GetWwxnPrefixRangeOk() (*HyperflexWwxnPrefixRange, bool)`

GetWwxnPrefixRangeOk returns a tuple with the WwxnPrefixRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwxnPrefixRange

`func (o *HyperflexExtFcStoragePolicyAllOf) SetWwxnPrefixRange(v HyperflexWwxnPrefixRange)`

SetWwxnPrefixRange sets WwxnPrefixRange field to given value.

### HasWwxnPrefixRange

`func (o *HyperflexExtFcStoragePolicyAllOf) HasWwxnPrefixRange() bool`

HasWwxnPrefixRange returns a boolean if a field has been set.

### GetClusterProfiles

`func (o *HyperflexExtFcStoragePolicyAllOf) GetClusterProfiles() []HyperflexClusterProfileRelationship`

GetClusterProfiles returns the ClusterProfiles field if non-nil, zero value otherwise.

### GetClusterProfilesOk

`func (o *HyperflexExtFcStoragePolicyAllOf) GetClusterProfilesOk() (*[]HyperflexClusterProfileRelationship, bool)`

GetClusterProfilesOk returns a tuple with the ClusterProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterProfiles

`func (o *HyperflexExtFcStoragePolicyAllOf) SetClusterProfiles(v []HyperflexClusterProfileRelationship)`

SetClusterProfiles sets ClusterProfiles field to given value.

### HasClusterProfiles

`func (o *HyperflexExtFcStoragePolicyAllOf) HasClusterProfiles() bool`

HasClusterProfiles returns a boolean if a field has been set.

### SetClusterProfilesNil

`func (o *HyperflexExtFcStoragePolicyAllOf) SetClusterProfilesNil(b bool)`

 SetClusterProfilesNil sets the value for ClusterProfiles to be an explicit nil

### UnsetClusterProfiles
`func (o *HyperflexExtFcStoragePolicyAllOf) UnsetClusterProfiles()`

UnsetClusterProfiles ensures that no value is present for ClusterProfiles, not even an explicit nil
### GetOrganization

`func (o *HyperflexExtFcStoragePolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *HyperflexExtFcStoragePolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *HyperflexExtFcStoragePolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *HyperflexExtFcStoragePolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


