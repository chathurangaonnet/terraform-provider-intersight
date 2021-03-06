/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-06-17T02:04:50-07:00.
 *
 * API version: 1.0.9-1867
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// VnicEthAdapterPolicy An Ethernet adapter policy governs the host-side behavior of the adapter, including how the adapter handles traffic. For each VIC Virtual Ethernet Interface various features like VXLAN, NVGRE, ARFS, Interrupt settings, and TCP Offload settings can be configured.
type VnicEthAdapterPolicy struct {
	PolicyAbstractPolicy
	// Enables advanced filtering on the interface.
	AdvancedFilter          *bool                        `json:"AdvancedFilter,omitempty"`
	ArfsSettings            *VnicArfsSettings            `json:"ArfsSettings,omitempty"`
	CompletionQueueSettings *VnicCompletionQueueSettings `json:"CompletionQueueSettings,omitempty"`
	// Enables Interrupt Scaling on the interface.
	InterruptScaling  *bool                     `json:"InterruptScaling,omitempty"`
	InterruptSettings *VnicEthInterruptSettings `json:"InterruptSettings,omitempty"`
	NvgreSettings     *VnicNvgreSettings        `json:"NvgreSettings,omitempty"`
	RoceSettings      *VnicRoceSettings         `json:"RoceSettings,omitempty"`
	RssHashSettings   *VnicRssHashSettings      `json:"RssHashSettings,omitempty"`
	// Receive Side Scaling allows the incoming traffic to be spread across multiple CPU cores.
	RssSettings        *bool                   `json:"RssSettings,omitempty"`
	RxQueueSettings    *VnicEthRxQueueSettings `json:"RxQueueSettings,omitempty"`
	TcpOffloadSettings *VnicTcpOffloadSettings `json:"TcpOffloadSettings,omitempty"`
	TxQueueSettings    *VnicEthTxQueueSettings `json:"TxQueueSettings,omitempty"`
	// Uplink Failback Timeout in seconds when uplink failover is enabled for a vNIC. After a vNIC has started using its secondary interface, this setting controls how long the primary interface must be available before the system resumes using the primary interface for the vNIC.
	UplinkFailbackTimeout *int64                                `json:"UplinkFailbackTimeout,omitempty"`
	VxlanSettings         *VnicVxlanSettings                    `json:"VxlanSettings,omitempty"`
	Organization          *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _VnicEthAdapterPolicy VnicEthAdapterPolicy

// NewVnicEthAdapterPolicy instantiates a new VnicEthAdapterPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicEthAdapterPolicy() *VnicEthAdapterPolicy {
	this := VnicEthAdapterPolicy{}
	return &this
}

// NewVnicEthAdapterPolicyWithDefaults instantiates a new VnicEthAdapterPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicEthAdapterPolicyWithDefaults() *VnicEthAdapterPolicy {
	this := VnicEthAdapterPolicy{}
	return &this
}

// GetAdvancedFilter returns the AdvancedFilter field value if set, zero value otherwise.
func (o *VnicEthAdapterPolicy) GetAdvancedFilter() bool {
	if o == nil || o.AdvancedFilter == nil {
		var ret bool
		return ret
	}
	return *o.AdvancedFilter
}

// GetAdvancedFilterOk returns a tuple with the AdvancedFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthAdapterPolicy) GetAdvancedFilterOk() (*bool, bool) {
	if o == nil || o.AdvancedFilter == nil {
		return nil, false
	}
	return o.AdvancedFilter, true
}

// HasAdvancedFilter returns a boolean if a field has been set.
func (o *VnicEthAdapterPolicy) HasAdvancedFilter() bool {
	if o != nil && o.AdvancedFilter != nil {
		return true
	}

	return false
}

// SetAdvancedFilter gets a reference to the given bool and assigns it to the AdvancedFilter field.
func (o *VnicEthAdapterPolicy) SetAdvancedFilter(v bool) {
	o.AdvancedFilter = &v
}

// GetArfsSettings returns the ArfsSettings field value if set, zero value otherwise.
func (o *VnicEthAdapterPolicy) GetArfsSettings() VnicArfsSettings {
	if o == nil || o.ArfsSettings == nil {
		var ret VnicArfsSettings
		return ret
	}
	return *o.ArfsSettings
}

// GetArfsSettingsOk returns a tuple with the ArfsSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthAdapterPolicy) GetArfsSettingsOk() (*VnicArfsSettings, bool) {
	if o == nil || o.ArfsSettings == nil {
		return nil, false
	}
	return o.ArfsSettings, true
}

// HasArfsSettings returns a boolean if a field has been set.
func (o *VnicEthAdapterPolicy) HasArfsSettings() bool {
	if o != nil && o.ArfsSettings != nil {
		return true
	}

	return false
}

// SetArfsSettings gets a reference to the given VnicArfsSettings and assigns it to the ArfsSettings field.
func (o *VnicEthAdapterPolicy) SetArfsSettings(v VnicArfsSettings) {
	o.ArfsSettings = &v
}

// GetCompletionQueueSettings returns the CompletionQueueSettings field value if set, zero value otherwise.
func (o *VnicEthAdapterPolicy) GetCompletionQueueSettings() VnicCompletionQueueSettings {
	if o == nil || o.CompletionQueueSettings == nil {
		var ret VnicCompletionQueueSettings
		return ret
	}
	return *o.CompletionQueueSettings
}

// GetCompletionQueueSettingsOk returns a tuple with the CompletionQueueSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthAdapterPolicy) GetCompletionQueueSettingsOk() (*VnicCompletionQueueSettings, bool) {
	if o == nil || o.CompletionQueueSettings == nil {
		return nil, false
	}
	return o.CompletionQueueSettings, true
}

// HasCompletionQueueSettings returns a boolean if a field has been set.
func (o *VnicEthAdapterPolicy) HasCompletionQueueSettings() bool {
	if o != nil && o.CompletionQueueSettings != nil {
		return true
	}

	return false
}

// SetCompletionQueueSettings gets a reference to the given VnicCompletionQueueSettings and assigns it to the CompletionQueueSettings field.
func (o *VnicEthAdapterPolicy) SetCompletionQueueSettings(v VnicCompletionQueueSettings) {
	o.CompletionQueueSettings = &v
}

// GetInterruptScaling returns the InterruptScaling field value if set, zero value otherwise.
func (o *VnicEthAdapterPolicy) GetInterruptScaling() bool {
	if o == nil || o.InterruptScaling == nil {
		var ret bool
		return ret
	}
	return *o.InterruptScaling
}

// GetInterruptScalingOk returns a tuple with the InterruptScaling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthAdapterPolicy) GetInterruptScalingOk() (*bool, bool) {
	if o == nil || o.InterruptScaling == nil {
		return nil, false
	}
	return o.InterruptScaling, true
}

// HasInterruptScaling returns a boolean if a field has been set.
func (o *VnicEthAdapterPolicy) HasInterruptScaling() bool {
	if o != nil && o.InterruptScaling != nil {
		return true
	}

	return false
}

// SetInterruptScaling gets a reference to the given bool and assigns it to the InterruptScaling field.
func (o *VnicEthAdapterPolicy) SetInterruptScaling(v bool) {
	o.InterruptScaling = &v
}

// GetInterruptSettings returns the InterruptSettings field value if set, zero value otherwise.
func (o *VnicEthAdapterPolicy) GetInterruptSettings() VnicEthInterruptSettings {
	if o == nil || o.InterruptSettings == nil {
		var ret VnicEthInterruptSettings
		return ret
	}
	return *o.InterruptSettings
}

// GetInterruptSettingsOk returns a tuple with the InterruptSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthAdapterPolicy) GetInterruptSettingsOk() (*VnicEthInterruptSettings, bool) {
	if o == nil || o.InterruptSettings == nil {
		return nil, false
	}
	return o.InterruptSettings, true
}

// HasInterruptSettings returns a boolean if a field has been set.
func (o *VnicEthAdapterPolicy) HasInterruptSettings() bool {
	if o != nil && o.InterruptSettings != nil {
		return true
	}

	return false
}

// SetInterruptSettings gets a reference to the given VnicEthInterruptSettings and assigns it to the InterruptSettings field.
func (o *VnicEthAdapterPolicy) SetInterruptSettings(v VnicEthInterruptSettings) {
	o.InterruptSettings = &v
}

// GetNvgreSettings returns the NvgreSettings field value if set, zero value otherwise.
func (o *VnicEthAdapterPolicy) GetNvgreSettings() VnicNvgreSettings {
	if o == nil || o.NvgreSettings == nil {
		var ret VnicNvgreSettings
		return ret
	}
	return *o.NvgreSettings
}

// GetNvgreSettingsOk returns a tuple with the NvgreSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthAdapterPolicy) GetNvgreSettingsOk() (*VnicNvgreSettings, bool) {
	if o == nil || o.NvgreSettings == nil {
		return nil, false
	}
	return o.NvgreSettings, true
}

// HasNvgreSettings returns a boolean if a field has been set.
func (o *VnicEthAdapterPolicy) HasNvgreSettings() bool {
	if o != nil && o.NvgreSettings != nil {
		return true
	}

	return false
}

// SetNvgreSettings gets a reference to the given VnicNvgreSettings and assigns it to the NvgreSettings field.
func (o *VnicEthAdapterPolicy) SetNvgreSettings(v VnicNvgreSettings) {
	o.NvgreSettings = &v
}

// GetRoceSettings returns the RoceSettings field value if set, zero value otherwise.
func (o *VnicEthAdapterPolicy) GetRoceSettings() VnicRoceSettings {
	if o == nil || o.RoceSettings == nil {
		var ret VnicRoceSettings
		return ret
	}
	return *o.RoceSettings
}

// GetRoceSettingsOk returns a tuple with the RoceSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthAdapterPolicy) GetRoceSettingsOk() (*VnicRoceSettings, bool) {
	if o == nil || o.RoceSettings == nil {
		return nil, false
	}
	return o.RoceSettings, true
}

// HasRoceSettings returns a boolean if a field has been set.
func (o *VnicEthAdapterPolicy) HasRoceSettings() bool {
	if o != nil && o.RoceSettings != nil {
		return true
	}

	return false
}

// SetRoceSettings gets a reference to the given VnicRoceSettings and assigns it to the RoceSettings field.
func (o *VnicEthAdapterPolicy) SetRoceSettings(v VnicRoceSettings) {
	o.RoceSettings = &v
}

// GetRssHashSettings returns the RssHashSettings field value if set, zero value otherwise.
func (o *VnicEthAdapterPolicy) GetRssHashSettings() VnicRssHashSettings {
	if o == nil || o.RssHashSettings == nil {
		var ret VnicRssHashSettings
		return ret
	}
	return *o.RssHashSettings
}

// GetRssHashSettingsOk returns a tuple with the RssHashSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthAdapterPolicy) GetRssHashSettingsOk() (*VnicRssHashSettings, bool) {
	if o == nil || o.RssHashSettings == nil {
		return nil, false
	}
	return o.RssHashSettings, true
}

// HasRssHashSettings returns a boolean if a field has been set.
func (o *VnicEthAdapterPolicy) HasRssHashSettings() bool {
	if o != nil && o.RssHashSettings != nil {
		return true
	}

	return false
}

// SetRssHashSettings gets a reference to the given VnicRssHashSettings and assigns it to the RssHashSettings field.
func (o *VnicEthAdapterPolicy) SetRssHashSettings(v VnicRssHashSettings) {
	o.RssHashSettings = &v
}

// GetRssSettings returns the RssSettings field value if set, zero value otherwise.
func (o *VnicEthAdapterPolicy) GetRssSettings() bool {
	if o == nil || o.RssSettings == nil {
		var ret bool
		return ret
	}
	return *o.RssSettings
}

// GetRssSettingsOk returns a tuple with the RssSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthAdapterPolicy) GetRssSettingsOk() (*bool, bool) {
	if o == nil || o.RssSettings == nil {
		return nil, false
	}
	return o.RssSettings, true
}

// HasRssSettings returns a boolean if a field has been set.
func (o *VnicEthAdapterPolicy) HasRssSettings() bool {
	if o != nil && o.RssSettings != nil {
		return true
	}

	return false
}

// SetRssSettings gets a reference to the given bool and assigns it to the RssSettings field.
func (o *VnicEthAdapterPolicy) SetRssSettings(v bool) {
	o.RssSettings = &v
}

// GetRxQueueSettings returns the RxQueueSettings field value if set, zero value otherwise.
func (o *VnicEthAdapterPolicy) GetRxQueueSettings() VnicEthRxQueueSettings {
	if o == nil || o.RxQueueSettings == nil {
		var ret VnicEthRxQueueSettings
		return ret
	}
	return *o.RxQueueSettings
}

// GetRxQueueSettingsOk returns a tuple with the RxQueueSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthAdapterPolicy) GetRxQueueSettingsOk() (*VnicEthRxQueueSettings, bool) {
	if o == nil || o.RxQueueSettings == nil {
		return nil, false
	}
	return o.RxQueueSettings, true
}

// HasRxQueueSettings returns a boolean if a field has been set.
func (o *VnicEthAdapterPolicy) HasRxQueueSettings() bool {
	if o != nil && o.RxQueueSettings != nil {
		return true
	}

	return false
}

// SetRxQueueSettings gets a reference to the given VnicEthRxQueueSettings and assigns it to the RxQueueSettings field.
func (o *VnicEthAdapterPolicy) SetRxQueueSettings(v VnicEthRxQueueSettings) {
	o.RxQueueSettings = &v
}

// GetTcpOffloadSettings returns the TcpOffloadSettings field value if set, zero value otherwise.
func (o *VnicEthAdapterPolicy) GetTcpOffloadSettings() VnicTcpOffloadSettings {
	if o == nil || o.TcpOffloadSettings == nil {
		var ret VnicTcpOffloadSettings
		return ret
	}
	return *o.TcpOffloadSettings
}

// GetTcpOffloadSettingsOk returns a tuple with the TcpOffloadSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthAdapterPolicy) GetTcpOffloadSettingsOk() (*VnicTcpOffloadSettings, bool) {
	if o == nil || o.TcpOffloadSettings == nil {
		return nil, false
	}
	return o.TcpOffloadSettings, true
}

// HasTcpOffloadSettings returns a boolean if a field has been set.
func (o *VnicEthAdapterPolicy) HasTcpOffloadSettings() bool {
	if o != nil && o.TcpOffloadSettings != nil {
		return true
	}

	return false
}

// SetTcpOffloadSettings gets a reference to the given VnicTcpOffloadSettings and assigns it to the TcpOffloadSettings field.
func (o *VnicEthAdapterPolicy) SetTcpOffloadSettings(v VnicTcpOffloadSettings) {
	o.TcpOffloadSettings = &v
}

// GetTxQueueSettings returns the TxQueueSettings field value if set, zero value otherwise.
func (o *VnicEthAdapterPolicy) GetTxQueueSettings() VnicEthTxQueueSettings {
	if o == nil || o.TxQueueSettings == nil {
		var ret VnicEthTxQueueSettings
		return ret
	}
	return *o.TxQueueSettings
}

// GetTxQueueSettingsOk returns a tuple with the TxQueueSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthAdapterPolicy) GetTxQueueSettingsOk() (*VnicEthTxQueueSettings, bool) {
	if o == nil || o.TxQueueSettings == nil {
		return nil, false
	}
	return o.TxQueueSettings, true
}

// HasTxQueueSettings returns a boolean if a field has been set.
func (o *VnicEthAdapterPolicy) HasTxQueueSettings() bool {
	if o != nil && o.TxQueueSettings != nil {
		return true
	}

	return false
}

// SetTxQueueSettings gets a reference to the given VnicEthTxQueueSettings and assigns it to the TxQueueSettings field.
func (o *VnicEthAdapterPolicy) SetTxQueueSettings(v VnicEthTxQueueSettings) {
	o.TxQueueSettings = &v
}

// GetUplinkFailbackTimeout returns the UplinkFailbackTimeout field value if set, zero value otherwise.
func (o *VnicEthAdapterPolicy) GetUplinkFailbackTimeout() int64 {
	if o == nil || o.UplinkFailbackTimeout == nil {
		var ret int64
		return ret
	}
	return *o.UplinkFailbackTimeout
}

// GetUplinkFailbackTimeoutOk returns a tuple with the UplinkFailbackTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthAdapterPolicy) GetUplinkFailbackTimeoutOk() (*int64, bool) {
	if o == nil || o.UplinkFailbackTimeout == nil {
		return nil, false
	}
	return o.UplinkFailbackTimeout, true
}

// HasUplinkFailbackTimeout returns a boolean if a field has been set.
func (o *VnicEthAdapterPolicy) HasUplinkFailbackTimeout() bool {
	if o != nil && o.UplinkFailbackTimeout != nil {
		return true
	}

	return false
}

// SetUplinkFailbackTimeout gets a reference to the given int64 and assigns it to the UplinkFailbackTimeout field.
func (o *VnicEthAdapterPolicy) SetUplinkFailbackTimeout(v int64) {
	o.UplinkFailbackTimeout = &v
}

// GetVxlanSettings returns the VxlanSettings field value if set, zero value otherwise.
func (o *VnicEthAdapterPolicy) GetVxlanSettings() VnicVxlanSettings {
	if o == nil || o.VxlanSettings == nil {
		var ret VnicVxlanSettings
		return ret
	}
	return *o.VxlanSettings
}

// GetVxlanSettingsOk returns a tuple with the VxlanSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthAdapterPolicy) GetVxlanSettingsOk() (*VnicVxlanSettings, bool) {
	if o == nil || o.VxlanSettings == nil {
		return nil, false
	}
	return o.VxlanSettings, true
}

// HasVxlanSettings returns a boolean if a field has been set.
func (o *VnicEthAdapterPolicy) HasVxlanSettings() bool {
	if o != nil && o.VxlanSettings != nil {
		return true
	}

	return false
}

// SetVxlanSettings gets a reference to the given VnicVxlanSettings and assigns it to the VxlanSettings field.
func (o *VnicEthAdapterPolicy) SetVxlanSettings(v VnicVxlanSettings) {
	o.VxlanSettings = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *VnicEthAdapterPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthAdapterPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *VnicEthAdapterPolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *VnicEthAdapterPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o VnicEthAdapterPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	if o.AdvancedFilter != nil {
		toSerialize["AdvancedFilter"] = o.AdvancedFilter
	}
	if o.ArfsSettings != nil {
		toSerialize["ArfsSettings"] = o.ArfsSettings
	}
	if o.CompletionQueueSettings != nil {
		toSerialize["CompletionQueueSettings"] = o.CompletionQueueSettings
	}
	if o.InterruptScaling != nil {
		toSerialize["InterruptScaling"] = o.InterruptScaling
	}
	if o.InterruptSettings != nil {
		toSerialize["InterruptSettings"] = o.InterruptSettings
	}
	if o.NvgreSettings != nil {
		toSerialize["NvgreSettings"] = o.NvgreSettings
	}
	if o.RoceSettings != nil {
		toSerialize["RoceSettings"] = o.RoceSettings
	}
	if o.RssHashSettings != nil {
		toSerialize["RssHashSettings"] = o.RssHashSettings
	}
	if o.RssSettings != nil {
		toSerialize["RssSettings"] = o.RssSettings
	}
	if o.RxQueueSettings != nil {
		toSerialize["RxQueueSettings"] = o.RxQueueSettings
	}
	if o.TcpOffloadSettings != nil {
		toSerialize["TcpOffloadSettings"] = o.TcpOffloadSettings
	}
	if o.TxQueueSettings != nil {
		toSerialize["TxQueueSettings"] = o.TxQueueSettings
	}
	if o.UplinkFailbackTimeout != nil {
		toSerialize["UplinkFailbackTimeout"] = o.UplinkFailbackTimeout
	}
	if o.VxlanSettings != nil {
		toSerialize["VxlanSettings"] = o.VxlanSettings
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicEthAdapterPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type VnicEthAdapterPolicyWithoutEmbeddedStruct struct {
		// Enables advanced filtering on the interface.
		AdvancedFilter          *bool                        `json:"AdvancedFilter,omitempty"`
		ArfsSettings            *VnicArfsSettings            `json:"ArfsSettings,omitempty"`
		CompletionQueueSettings *VnicCompletionQueueSettings `json:"CompletionQueueSettings,omitempty"`
		// Enables Interrupt Scaling on the interface.
		InterruptScaling  *bool                     `json:"InterruptScaling,omitempty"`
		InterruptSettings *VnicEthInterruptSettings `json:"InterruptSettings,omitempty"`
		NvgreSettings     *VnicNvgreSettings        `json:"NvgreSettings,omitempty"`
		RoceSettings      *VnicRoceSettings         `json:"RoceSettings,omitempty"`
		RssHashSettings   *VnicRssHashSettings      `json:"RssHashSettings,omitempty"`
		// Receive Side Scaling allows the incoming traffic to be spread across multiple CPU cores.
		RssSettings        *bool                   `json:"RssSettings,omitempty"`
		RxQueueSettings    *VnicEthRxQueueSettings `json:"RxQueueSettings,omitempty"`
		TcpOffloadSettings *VnicTcpOffloadSettings `json:"TcpOffloadSettings,omitempty"`
		TxQueueSettings    *VnicEthTxQueueSettings `json:"TxQueueSettings,omitempty"`
		// Uplink Failback Timeout in seconds when uplink failover is enabled for a vNIC. After a vNIC has started using its secondary interface, this setting controls how long the primary interface must be available before the system resumes using the primary interface for the vNIC.
		UplinkFailbackTimeout *int64                                `json:"UplinkFailbackTimeout,omitempty"`
		VxlanSettings         *VnicVxlanSettings                    `json:"VxlanSettings,omitempty"`
		Organization          *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varVnicEthAdapterPolicyWithoutEmbeddedStruct := VnicEthAdapterPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVnicEthAdapterPolicyWithoutEmbeddedStruct)
	if err == nil {
		varVnicEthAdapterPolicy := _VnicEthAdapterPolicy{}
		varVnicEthAdapterPolicy.AdvancedFilter = varVnicEthAdapterPolicyWithoutEmbeddedStruct.AdvancedFilter
		varVnicEthAdapterPolicy.ArfsSettings = varVnicEthAdapterPolicyWithoutEmbeddedStruct.ArfsSettings
		varVnicEthAdapterPolicy.CompletionQueueSettings = varVnicEthAdapterPolicyWithoutEmbeddedStruct.CompletionQueueSettings
		varVnicEthAdapterPolicy.InterruptScaling = varVnicEthAdapterPolicyWithoutEmbeddedStruct.InterruptScaling
		varVnicEthAdapterPolicy.InterruptSettings = varVnicEthAdapterPolicyWithoutEmbeddedStruct.InterruptSettings
		varVnicEthAdapterPolicy.NvgreSettings = varVnicEthAdapterPolicyWithoutEmbeddedStruct.NvgreSettings
		varVnicEthAdapterPolicy.RoceSettings = varVnicEthAdapterPolicyWithoutEmbeddedStruct.RoceSettings
		varVnicEthAdapterPolicy.RssHashSettings = varVnicEthAdapterPolicyWithoutEmbeddedStruct.RssHashSettings
		varVnicEthAdapterPolicy.RssSettings = varVnicEthAdapterPolicyWithoutEmbeddedStruct.RssSettings
		varVnicEthAdapterPolicy.RxQueueSettings = varVnicEthAdapterPolicyWithoutEmbeddedStruct.RxQueueSettings
		varVnicEthAdapterPolicy.TcpOffloadSettings = varVnicEthAdapterPolicyWithoutEmbeddedStruct.TcpOffloadSettings
		varVnicEthAdapterPolicy.TxQueueSettings = varVnicEthAdapterPolicyWithoutEmbeddedStruct.TxQueueSettings
		varVnicEthAdapterPolicy.UplinkFailbackTimeout = varVnicEthAdapterPolicyWithoutEmbeddedStruct.UplinkFailbackTimeout
		varVnicEthAdapterPolicy.VxlanSettings = varVnicEthAdapterPolicyWithoutEmbeddedStruct.VxlanSettings
		varVnicEthAdapterPolicy.Organization = varVnicEthAdapterPolicyWithoutEmbeddedStruct.Organization
		*o = VnicEthAdapterPolicy(varVnicEthAdapterPolicy)
	} else {
		return err
	}

	varVnicEthAdapterPolicy := _VnicEthAdapterPolicy{}

	err = json.Unmarshal(bytes, &varVnicEthAdapterPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varVnicEthAdapterPolicy.PolicyAbstractPolicy
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "AdvancedFilter")
		delete(additionalProperties, "ArfsSettings")
		delete(additionalProperties, "CompletionQueueSettings")
		delete(additionalProperties, "InterruptScaling")
		delete(additionalProperties, "InterruptSettings")
		delete(additionalProperties, "NvgreSettings")
		delete(additionalProperties, "RoceSettings")
		delete(additionalProperties, "RssHashSettings")
		delete(additionalProperties, "RssSettings")
		delete(additionalProperties, "RxQueueSettings")
		delete(additionalProperties, "TcpOffloadSettings")
		delete(additionalProperties, "TxQueueSettings")
		delete(additionalProperties, "UplinkFailbackTimeout")
		delete(additionalProperties, "VxlanSettings")
		delete(additionalProperties, "Organization")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicy := reflect.ValueOf(o.PolicyAbstractPolicy)
		for i := 0; i < reflectPolicyAbstractPolicy.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicy.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVnicEthAdapterPolicy struct {
	value *VnicEthAdapterPolicy
	isSet bool
}

func (v NullableVnicEthAdapterPolicy) Get() *VnicEthAdapterPolicy {
	return v.value
}

func (v *NullableVnicEthAdapterPolicy) Set(val *VnicEthAdapterPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicEthAdapterPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicEthAdapterPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicEthAdapterPolicy(val *VnicEthAdapterPolicy) *NullableVnicEthAdapterPolicy {
	return &NullableVnicEthAdapterPolicy{value: val, isSet: true}
}

func (v NullableVnicEthAdapterPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicEthAdapterPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
