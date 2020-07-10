# FabricEthNetworkControlPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CdpEnabled** | Pointer to **bool** | Enables the CDP on an interface. | [optional] 
**ForgeMac** | Pointer to **string** | Determines if the MAC forging is allowed or denied on an interface. | [optional] [default to "allow"]
**LldpSettings** | Pointer to [**FabricLldpSettings**](fabric.LldpSettings.md) |  | [optional] 
**MacRegistrationMode** | Pointer to **string** | Determines the MAC addresses that have to be registered with the switch. | [optional] [default to "nativeVlanOnly"]
**UplinkFailAction** | Pointer to **string** | Determines the state of the virtual interface (vethernet / vfc) on the switch when a suitable uplink is not pinned. | [optional] [default to "linkDown"]
**NetworkPolicy** | Pointer to [**[]VnicEthNetworkPolicyRelationship**](vnic.EthNetworkPolicy.Relationship.md) | An array of relationships to vnicEthNetworkPolicy resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 

## Methods

### NewFabricEthNetworkControlPolicy

`func NewFabricEthNetworkControlPolicy() *FabricEthNetworkControlPolicy`

NewFabricEthNetworkControlPolicy instantiates a new FabricEthNetworkControlPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricEthNetworkControlPolicyWithDefaults

`func NewFabricEthNetworkControlPolicyWithDefaults() *FabricEthNetworkControlPolicy`

NewFabricEthNetworkControlPolicyWithDefaults instantiates a new FabricEthNetworkControlPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCdpEnabled

`func (o *FabricEthNetworkControlPolicy) GetCdpEnabled() bool`

GetCdpEnabled returns the CdpEnabled field if non-nil, zero value otherwise.

### GetCdpEnabledOk

`func (o *FabricEthNetworkControlPolicy) GetCdpEnabledOk() (*bool, bool)`

GetCdpEnabledOk returns a tuple with the CdpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdpEnabled

`func (o *FabricEthNetworkControlPolicy) SetCdpEnabled(v bool)`

SetCdpEnabled sets CdpEnabled field to given value.

### HasCdpEnabled

`func (o *FabricEthNetworkControlPolicy) HasCdpEnabled() bool`

HasCdpEnabled returns a boolean if a field has been set.

### GetForgeMac

`func (o *FabricEthNetworkControlPolicy) GetForgeMac() string`

GetForgeMac returns the ForgeMac field if non-nil, zero value otherwise.

### GetForgeMacOk

`func (o *FabricEthNetworkControlPolicy) GetForgeMacOk() (*string, bool)`

GetForgeMacOk returns a tuple with the ForgeMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForgeMac

`func (o *FabricEthNetworkControlPolicy) SetForgeMac(v string)`

SetForgeMac sets ForgeMac field to given value.

### HasForgeMac

`func (o *FabricEthNetworkControlPolicy) HasForgeMac() bool`

HasForgeMac returns a boolean if a field has been set.

### GetLldpSettings

`func (o *FabricEthNetworkControlPolicy) GetLldpSettings() FabricLldpSettings`

GetLldpSettings returns the LldpSettings field if non-nil, zero value otherwise.

### GetLldpSettingsOk

`func (o *FabricEthNetworkControlPolicy) GetLldpSettingsOk() (*FabricLldpSettings, bool)`

GetLldpSettingsOk returns a tuple with the LldpSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldpSettings

`func (o *FabricEthNetworkControlPolicy) SetLldpSettings(v FabricLldpSettings)`

SetLldpSettings sets LldpSettings field to given value.

### HasLldpSettings

`func (o *FabricEthNetworkControlPolicy) HasLldpSettings() bool`

HasLldpSettings returns a boolean if a field has been set.

### GetMacRegistrationMode

`func (o *FabricEthNetworkControlPolicy) GetMacRegistrationMode() string`

GetMacRegistrationMode returns the MacRegistrationMode field if non-nil, zero value otherwise.

### GetMacRegistrationModeOk

`func (o *FabricEthNetworkControlPolicy) GetMacRegistrationModeOk() (*string, bool)`

GetMacRegistrationModeOk returns a tuple with the MacRegistrationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacRegistrationMode

`func (o *FabricEthNetworkControlPolicy) SetMacRegistrationMode(v string)`

SetMacRegistrationMode sets MacRegistrationMode field to given value.

### HasMacRegistrationMode

`func (o *FabricEthNetworkControlPolicy) HasMacRegistrationMode() bool`

HasMacRegistrationMode returns a boolean if a field has been set.

### GetUplinkFailAction

`func (o *FabricEthNetworkControlPolicy) GetUplinkFailAction() string`

GetUplinkFailAction returns the UplinkFailAction field if non-nil, zero value otherwise.

### GetUplinkFailActionOk

`func (o *FabricEthNetworkControlPolicy) GetUplinkFailActionOk() (*string, bool)`

GetUplinkFailActionOk returns a tuple with the UplinkFailAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkFailAction

`func (o *FabricEthNetworkControlPolicy) SetUplinkFailAction(v string)`

SetUplinkFailAction sets UplinkFailAction field to given value.

### HasUplinkFailAction

`func (o *FabricEthNetworkControlPolicy) HasUplinkFailAction() bool`

HasUplinkFailAction returns a boolean if a field has been set.

### GetNetworkPolicy

`func (o *FabricEthNetworkControlPolicy) GetNetworkPolicy() []VnicEthNetworkPolicyRelationship`

GetNetworkPolicy returns the NetworkPolicy field if non-nil, zero value otherwise.

### GetNetworkPolicyOk

`func (o *FabricEthNetworkControlPolicy) GetNetworkPolicyOk() (*[]VnicEthNetworkPolicyRelationship, bool)`

GetNetworkPolicyOk returns a tuple with the NetworkPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPolicy

`func (o *FabricEthNetworkControlPolicy) SetNetworkPolicy(v []VnicEthNetworkPolicyRelationship)`

SetNetworkPolicy sets NetworkPolicy field to given value.

### HasNetworkPolicy

`func (o *FabricEthNetworkControlPolicy) HasNetworkPolicy() bool`

HasNetworkPolicy returns a boolean if a field has been set.

### SetNetworkPolicyNil

`func (o *FabricEthNetworkControlPolicy) SetNetworkPolicyNil(b bool)`

 SetNetworkPolicyNil sets the value for NetworkPolicy to be an explicit nil

### UnsetNetworkPolicy
`func (o *FabricEthNetworkControlPolicy) UnsetNetworkPolicy()`

UnsetNetworkPolicy ensures that no value is present for NetworkPolicy, not even an explicit nil
### GetOrganization

`func (o *FabricEthNetworkControlPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *FabricEthNetworkControlPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *FabricEthNetworkControlPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *FabricEthNetworkControlPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


