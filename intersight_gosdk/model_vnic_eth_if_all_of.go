/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-04-12T21:47:47-07:00.
 *
 * API version: 1.0.9-1617
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
)

// VnicEthIfAllOf Definition of the list of properties defined in 'vnic.EthIf', excluding properties defined in parent classes.
type VnicEthIfAllOf struct {
	Cdn *VnicCdn `json:"Cdn,omitempty"`
	// Name of the virtual ethernet interface.
	Name *string `json:"Name,omitempty"`
	// The order in which the virtual interface is brought up. The order assigned to an interface should be unique for all the Ethernet and Fibre-Channel interfaces on each PCI link on a VIC adapter. The maximum value of PCI order is limited by the number of virtual interfaces (Ethernet and Fibre-Channel) on each PCI link on a VIC adapter. All VIC adapters have a single PCI link except VIC 1385 which has two.
	Order                 *int64                                 `json:"Order,omitempty"`
	Placement             *VnicPlacementSettings                 `json:"Placement,omitempty"`
	UsnicSettings         *VnicUsnicSettings                     `json:"UsnicSettings,omitempty"`
	VmqSettings           *VnicVmqSettings                       `json:"VmqSettings,omitempty"`
	EthAdapterPolicy      *VnicEthAdapterPolicyRelationship      `json:"EthAdapterPolicy,omitempty"`
	EthNetworkPolicy      *VnicEthNetworkPolicyRelationship      `json:"EthNetworkPolicy,omitempty"`
	EthQosPolicy          *VnicEthQosPolicyRelationship          `json:"EthQosPolicy,omitempty"`
	LanConnectivityPolicy *VnicLanConnectivityPolicyRelationship `json:"LanConnectivityPolicy,omitempty"`
	Organization          *OrganizationOrganizationRelationship  `json:"Organization,omitempty"`
}

// NewVnicEthIfAllOf instantiates a new VnicEthIfAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicEthIfAllOf() *VnicEthIfAllOf {
	this := VnicEthIfAllOf{}
	return &this
}

// NewVnicEthIfAllOfWithDefaults instantiates a new VnicEthIfAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicEthIfAllOfWithDefaults() *VnicEthIfAllOf {
	this := VnicEthIfAllOf{}
	return &this
}

// GetCdn returns the Cdn field value if set, zero value otherwise.
func (o *VnicEthIfAllOf) GetCdn() VnicCdn {
	if o == nil || o.Cdn == nil {
		var ret VnicCdn
		return ret
	}
	return *o.Cdn
}

// GetCdnOk returns a tuple with the Cdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthIfAllOf) GetCdnOk() (*VnicCdn, bool) {
	if o == nil || o.Cdn == nil {
		return nil, false
	}
	return o.Cdn, true
}

// HasCdn returns a boolean if a field has been set.
func (o *VnicEthIfAllOf) HasCdn() bool {
	if o != nil && o.Cdn != nil {
		return true
	}

	return false
}

// SetCdn gets a reference to the given VnicCdn and assigns it to the Cdn field.
func (o *VnicEthIfAllOf) SetCdn(v VnicCdn) {
	o.Cdn = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VnicEthIfAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthIfAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VnicEthIfAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VnicEthIfAllOf) SetName(v string) {
	o.Name = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *VnicEthIfAllOf) GetOrder() int64 {
	if o == nil || o.Order == nil {
		var ret int64
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthIfAllOf) GetOrderOk() (*int64, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *VnicEthIfAllOf) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int64 and assigns it to the Order field.
func (o *VnicEthIfAllOf) SetOrder(v int64) {
	o.Order = &v
}

// GetPlacement returns the Placement field value if set, zero value otherwise.
func (o *VnicEthIfAllOf) GetPlacement() VnicPlacementSettings {
	if o == nil || o.Placement == nil {
		var ret VnicPlacementSettings
		return ret
	}
	return *o.Placement
}

// GetPlacementOk returns a tuple with the Placement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthIfAllOf) GetPlacementOk() (*VnicPlacementSettings, bool) {
	if o == nil || o.Placement == nil {
		return nil, false
	}
	return o.Placement, true
}

// HasPlacement returns a boolean if a field has been set.
func (o *VnicEthIfAllOf) HasPlacement() bool {
	if o != nil && o.Placement != nil {
		return true
	}

	return false
}

// SetPlacement gets a reference to the given VnicPlacementSettings and assigns it to the Placement field.
func (o *VnicEthIfAllOf) SetPlacement(v VnicPlacementSettings) {
	o.Placement = &v
}

// GetUsnicSettings returns the UsnicSettings field value if set, zero value otherwise.
func (o *VnicEthIfAllOf) GetUsnicSettings() VnicUsnicSettings {
	if o == nil || o.UsnicSettings == nil {
		var ret VnicUsnicSettings
		return ret
	}
	return *o.UsnicSettings
}

// GetUsnicSettingsOk returns a tuple with the UsnicSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthIfAllOf) GetUsnicSettingsOk() (*VnicUsnicSettings, bool) {
	if o == nil || o.UsnicSettings == nil {
		return nil, false
	}
	return o.UsnicSettings, true
}

// HasUsnicSettings returns a boolean if a field has been set.
func (o *VnicEthIfAllOf) HasUsnicSettings() bool {
	if o != nil && o.UsnicSettings != nil {
		return true
	}

	return false
}

// SetUsnicSettings gets a reference to the given VnicUsnicSettings and assigns it to the UsnicSettings field.
func (o *VnicEthIfAllOf) SetUsnicSettings(v VnicUsnicSettings) {
	o.UsnicSettings = &v
}

// GetVmqSettings returns the VmqSettings field value if set, zero value otherwise.
func (o *VnicEthIfAllOf) GetVmqSettings() VnicVmqSettings {
	if o == nil || o.VmqSettings == nil {
		var ret VnicVmqSettings
		return ret
	}
	return *o.VmqSettings
}

// GetVmqSettingsOk returns a tuple with the VmqSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthIfAllOf) GetVmqSettingsOk() (*VnicVmqSettings, bool) {
	if o == nil || o.VmqSettings == nil {
		return nil, false
	}
	return o.VmqSettings, true
}

// HasVmqSettings returns a boolean if a field has been set.
func (o *VnicEthIfAllOf) HasVmqSettings() bool {
	if o != nil && o.VmqSettings != nil {
		return true
	}

	return false
}

// SetVmqSettings gets a reference to the given VnicVmqSettings and assigns it to the VmqSettings field.
func (o *VnicEthIfAllOf) SetVmqSettings(v VnicVmqSettings) {
	o.VmqSettings = &v
}

// GetEthAdapterPolicy returns the EthAdapterPolicy field value if set, zero value otherwise.
func (o *VnicEthIfAllOf) GetEthAdapterPolicy() VnicEthAdapterPolicyRelationship {
	if o == nil || o.EthAdapterPolicy == nil {
		var ret VnicEthAdapterPolicyRelationship
		return ret
	}
	return *o.EthAdapterPolicy
}

// GetEthAdapterPolicyOk returns a tuple with the EthAdapterPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthIfAllOf) GetEthAdapterPolicyOk() (*VnicEthAdapterPolicyRelationship, bool) {
	if o == nil || o.EthAdapterPolicy == nil {
		return nil, false
	}
	return o.EthAdapterPolicy, true
}

// HasEthAdapterPolicy returns a boolean if a field has been set.
func (o *VnicEthIfAllOf) HasEthAdapterPolicy() bool {
	if o != nil && o.EthAdapterPolicy != nil {
		return true
	}

	return false
}

// SetEthAdapterPolicy gets a reference to the given VnicEthAdapterPolicyRelationship and assigns it to the EthAdapterPolicy field.
func (o *VnicEthIfAllOf) SetEthAdapterPolicy(v VnicEthAdapterPolicyRelationship) {
	o.EthAdapterPolicy = &v
}

// GetEthNetworkPolicy returns the EthNetworkPolicy field value if set, zero value otherwise.
func (o *VnicEthIfAllOf) GetEthNetworkPolicy() VnicEthNetworkPolicyRelationship {
	if o == nil || o.EthNetworkPolicy == nil {
		var ret VnicEthNetworkPolicyRelationship
		return ret
	}
	return *o.EthNetworkPolicy
}

// GetEthNetworkPolicyOk returns a tuple with the EthNetworkPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthIfAllOf) GetEthNetworkPolicyOk() (*VnicEthNetworkPolicyRelationship, bool) {
	if o == nil || o.EthNetworkPolicy == nil {
		return nil, false
	}
	return o.EthNetworkPolicy, true
}

// HasEthNetworkPolicy returns a boolean if a field has been set.
func (o *VnicEthIfAllOf) HasEthNetworkPolicy() bool {
	if o != nil && o.EthNetworkPolicy != nil {
		return true
	}

	return false
}

// SetEthNetworkPolicy gets a reference to the given VnicEthNetworkPolicyRelationship and assigns it to the EthNetworkPolicy field.
func (o *VnicEthIfAllOf) SetEthNetworkPolicy(v VnicEthNetworkPolicyRelationship) {
	o.EthNetworkPolicy = &v
}

// GetEthQosPolicy returns the EthQosPolicy field value if set, zero value otherwise.
func (o *VnicEthIfAllOf) GetEthQosPolicy() VnicEthQosPolicyRelationship {
	if o == nil || o.EthQosPolicy == nil {
		var ret VnicEthQosPolicyRelationship
		return ret
	}
	return *o.EthQosPolicy
}

// GetEthQosPolicyOk returns a tuple with the EthQosPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthIfAllOf) GetEthQosPolicyOk() (*VnicEthQosPolicyRelationship, bool) {
	if o == nil || o.EthQosPolicy == nil {
		return nil, false
	}
	return o.EthQosPolicy, true
}

// HasEthQosPolicy returns a boolean if a field has been set.
func (o *VnicEthIfAllOf) HasEthQosPolicy() bool {
	if o != nil && o.EthQosPolicy != nil {
		return true
	}

	return false
}

// SetEthQosPolicy gets a reference to the given VnicEthQosPolicyRelationship and assigns it to the EthQosPolicy field.
func (o *VnicEthIfAllOf) SetEthQosPolicy(v VnicEthQosPolicyRelationship) {
	o.EthQosPolicy = &v
}

// GetLanConnectivityPolicy returns the LanConnectivityPolicy field value if set, zero value otherwise.
func (o *VnicEthIfAllOf) GetLanConnectivityPolicy() VnicLanConnectivityPolicyRelationship {
	if o == nil || o.LanConnectivityPolicy == nil {
		var ret VnicLanConnectivityPolicyRelationship
		return ret
	}
	return *o.LanConnectivityPolicy
}

// GetLanConnectivityPolicyOk returns a tuple with the LanConnectivityPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthIfAllOf) GetLanConnectivityPolicyOk() (*VnicLanConnectivityPolicyRelationship, bool) {
	if o == nil || o.LanConnectivityPolicy == nil {
		return nil, false
	}
	return o.LanConnectivityPolicy, true
}

// HasLanConnectivityPolicy returns a boolean if a field has been set.
func (o *VnicEthIfAllOf) HasLanConnectivityPolicy() bool {
	if o != nil && o.LanConnectivityPolicy != nil {
		return true
	}

	return false
}

// SetLanConnectivityPolicy gets a reference to the given VnicLanConnectivityPolicyRelationship and assigns it to the LanConnectivityPolicy field.
func (o *VnicEthIfAllOf) SetLanConnectivityPolicy(v VnicLanConnectivityPolicyRelationship) {
	o.LanConnectivityPolicy = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *VnicEthIfAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthIfAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *VnicEthIfAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *VnicEthIfAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o VnicEthIfAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cdn != nil {
		toSerialize["Cdn"] = o.Cdn
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Order != nil {
		toSerialize["Order"] = o.Order
	}
	if o.Placement != nil {
		toSerialize["Placement"] = o.Placement
	}
	if o.UsnicSettings != nil {
		toSerialize["UsnicSettings"] = o.UsnicSettings
	}
	if o.VmqSettings != nil {
		toSerialize["VmqSettings"] = o.VmqSettings
	}
	if o.EthAdapterPolicy != nil {
		toSerialize["EthAdapterPolicy"] = o.EthAdapterPolicy
	}
	if o.EthNetworkPolicy != nil {
		toSerialize["EthNetworkPolicy"] = o.EthNetworkPolicy
	}
	if o.EthQosPolicy != nil {
		toSerialize["EthQosPolicy"] = o.EthQosPolicy
	}
	if o.LanConnectivityPolicy != nil {
		toSerialize["LanConnectivityPolicy"] = o.LanConnectivityPolicy
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	return json.Marshal(toSerialize)
}

type NullableVnicEthIfAllOf struct {
	value *VnicEthIfAllOf
	isSet bool
}

func (v NullableVnicEthIfAllOf) Get() *VnicEthIfAllOf {
	return v.value
}

func (v *NullableVnicEthIfAllOf) Set(val *VnicEthIfAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicEthIfAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicEthIfAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicEthIfAllOf(val *VnicEthIfAllOf) *NullableVnicEthIfAllOf {
	return &NullableVnicEthIfAllOf{value: val, isSet: true}
}

func (v NullableVnicEthIfAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicEthIfAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
