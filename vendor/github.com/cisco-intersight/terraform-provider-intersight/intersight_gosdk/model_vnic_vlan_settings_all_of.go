/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-07-08T07:46:03Z.
 *
 * API version: 0.0.1-15983
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
)

// VnicVlanSettingsAllOf Definition of the list of properties defined in 'vnic.VlanSettings', excluding properties defined in parent classes.
type VnicVlanSettingsAllOf struct {
	// Allowed VLAN IDs of the virtual interface.
	AllowedVlans *string `json:"AllowedVlans,omitempty"`
	// Native VLAN ID of the virtual interface or the corresponding vethernet on the peer Fabric Interconnect to which the virtual interface is connected. Setting the ID to 0 will not associate any native VLAN to the traffic on the virtual interface.
	DefaultVlan *int64 `json:"DefaultVlan,omitempty"`
	// Option to determine if the port can carry single VLAN (Access) or multiple VLANs (Trunk) traffic.
	Mode *string `json:"Mode,omitempty"`
}

// NewVnicVlanSettingsAllOf instantiates a new VnicVlanSettingsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicVlanSettingsAllOf() *VnicVlanSettingsAllOf {
	this := VnicVlanSettingsAllOf{}
	var mode string = "ACCESS"
	this.Mode = &mode
	return &this
}

// NewVnicVlanSettingsAllOfWithDefaults instantiates a new VnicVlanSettingsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicVlanSettingsAllOfWithDefaults() *VnicVlanSettingsAllOf {
	this := VnicVlanSettingsAllOf{}
	var mode string = "ACCESS"
	this.Mode = &mode
	return &this
}

// GetAllowedVlans returns the AllowedVlans field value if set, zero value otherwise.
func (o *VnicVlanSettingsAllOf) GetAllowedVlans() string {
	if o == nil || o.AllowedVlans == nil {
		var ret string
		return ret
	}
	return *o.AllowedVlans
}

// GetAllowedVlansOk returns a tuple with the AllowedVlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicVlanSettingsAllOf) GetAllowedVlansOk() (*string, bool) {
	if o == nil || o.AllowedVlans == nil {
		return nil, false
	}
	return o.AllowedVlans, true
}

// HasAllowedVlans returns a boolean if a field has been set.
func (o *VnicVlanSettingsAllOf) HasAllowedVlans() bool {
	if o != nil && o.AllowedVlans != nil {
		return true
	}

	return false
}

// SetAllowedVlans gets a reference to the given string and assigns it to the AllowedVlans field.
func (o *VnicVlanSettingsAllOf) SetAllowedVlans(v string) {
	o.AllowedVlans = &v
}

// GetDefaultVlan returns the DefaultVlan field value if set, zero value otherwise.
func (o *VnicVlanSettingsAllOf) GetDefaultVlan() int64 {
	if o == nil || o.DefaultVlan == nil {
		var ret int64
		return ret
	}
	return *o.DefaultVlan
}

// GetDefaultVlanOk returns a tuple with the DefaultVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicVlanSettingsAllOf) GetDefaultVlanOk() (*int64, bool) {
	if o == nil || o.DefaultVlan == nil {
		return nil, false
	}
	return o.DefaultVlan, true
}

// HasDefaultVlan returns a boolean if a field has been set.
func (o *VnicVlanSettingsAllOf) HasDefaultVlan() bool {
	if o != nil && o.DefaultVlan != nil {
		return true
	}

	return false
}

// SetDefaultVlan gets a reference to the given int64 and assigns it to the DefaultVlan field.
func (o *VnicVlanSettingsAllOf) SetDefaultVlan(v int64) {
	o.DefaultVlan = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *VnicVlanSettingsAllOf) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicVlanSettingsAllOf) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *VnicVlanSettingsAllOf) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *VnicVlanSettingsAllOf) SetMode(v string) {
	o.Mode = &v
}

func (o VnicVlanSettingsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowedVlans != nil {
		toSerialize["AllowedVlans"] = o.AllowedVlans
	}
	if o.DefaultVlan != nil {
		toSerialize["DefaultVlan"] = o.DefaultVlan
	}
	if o.Mode != nil {
		toSerialize["Mode"] = o.Mode
	}
	return json.Marshal(toSerialize)
}

type NullableVnicVlanSettingsAllOf struct {
	value *VnicVlanSettingsAllOf
	isSet bool
}

func (v NullableVnicVlanSettingsAllOf) Get() *VnicVlanSettingsAllOf {
	return v.value
}

func (v *NullableVnicVlanSettingsAllOf) Set(val *VnicVlanSettingsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicVlanSettingsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicVlanSettingsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicVlanSettingsAllOf(val *VnicVlanSettingsAllOf) *NullableVnicVlanSettingsAllOf {
	return &NullableVnicVlanSettingsAllOf{value: val, isSet: true}
}

func (v NullableVnicVlanSettingsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicVlanSettingsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
