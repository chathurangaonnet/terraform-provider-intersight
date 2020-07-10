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

// FabricPortModeAllOf Definition of the list of properties defined in 'fabric.PortMode', excluding properties defined in parent classes.
type FabricPortModeAllOf struct {
	// Custom Port Mode specified for the port range.
	CustomMode *string `json:"CustomMode,omitempty"`
	// Ending range of the Port Identifier.
	PortIdEnd *int64 `json:"PortIdEnd,omitempty"`
	// Starting range of the Port Identifier.
	PortIdStart *int64 `json:"PortIdStart,omitempty"`
	// Slot Identifier of the switch.
	SlotId     *int64                        `json:"SlotId,omitempty"`
	PortPolicy *FabricPortPolicyRelationship `json:"PortPolicy,omitempty"`
}

// NewFabricPortModeAllOf instantiates a new FabricPortModeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricPortModeAllOf() *FabricPortModeAllOf {
	this := FabricPortModeAllOf{}
	var customMode string = "FibreChannel"
	this.CustomMode = &customMode
	return &this
}

// NewFabricPortModeAllOfWithDefaults instantiates a new FabricPortModeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricPortModeAllOfWithDefaults() *FabricPortModeAllOf {
	this := FabricPortModeAllOf{}
	var customMode string = "FibreChannel"
	this.CustomMode = &customMode
	return &this
}

// GetCustomMode returns the CustomMode field value if set, zero value otherwise.
func (o *FabricPortModeAllOf) GetCustomMode() string {
	if o == nil || o.CustomMode == nil {
		var ret string
		return ret
	}
	return *o.CustomMode
}

// GetCustomModeOk returns a tuple with the CustomMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricPortModeAllOf) GetCustomModeOk() (*string, bool) {
	if o == nil || o.CustomMode == nil {
		return nil, false
	}
	return o.CustomMode, true
}

// HasCustomMode returns a boolean if a field has been set.
func (o *FabricPortModeAllOf) HasCustomMode() bool {
	if o != nil && o.CustomMode != nil {
		return true
	}

	return false
}

// SetCustomMode gets a reference to the given string and assigns it to the CustomMode field.
func (o *FabricPortModeAllOf) SetCustomMode(v string) {
	o.CustomMode = &v
}

// GetPortIdEnd returns the PortIdEnd field value if set, zero value otherwise.
func (o *FabricPortModeAllOf) GetPortIdEnd() int64 {
	if o == nil || o.PortIdEnd == nil {
		var ret int64
		return ret
	}
	return *o.PortIdEnd
}

// GetPortIdEndOk returns a tuple with the PortIdEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricPortModeAllOf) GetPortIdEndOk() (*int64, bool) {
	if o == nil || o.PortIdEnd == nil {
		return nil, false
	}
	return o.PortIdEnd, true
}

// HasPortIdEnd returns a boolean if a field has been set.
func (o *FabricPortModeAllOf) HasPortIdEnd() bool {
	if o != nil && o.PortIdEnd != nil {
		return true
	}

	return false
}

// SetPortIdEnd gets a reference to the given int64 and assigns it to the PortIdEnd field.
func (o *FabricPortModeAllOf) SetPortIdEnd(v int64) {
	o.PortIdEnd = &v
}

// GetPortIdStart returns the PortIdStart field value if set, zero value otherwise.
func (o *FabricPortModeAllOf) GetPortIdStart() int64 {
	if o == nil || o.PortIdStart == nil {
		var ret int64
		return ret
	}
	return *o.PortIdStart
}

// GetPortIdStartOk returns a tuple with the PortIdStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricPortModeAllOf) GetPortIdStartOk() (*int64, bool) {
	if o == nil || o.PortIdStart == nil {
		return nil, false
	}
	return o.PortIdStart, true
}

// HasPortIdStart returns a boolean if a field has been set.
func (o *FabricPortModeAllOf) HasPortIdStart() bool {
	if o != nil && o.PortIdStart != nil {
		return true
	}

	return false
}

// SetPortIdStart gets a reference to the given int64 and assigns it to the PortIdStart field.
func (o *FabricPortModeAllOf) SetPortIdStart(v int64) {
	o.PortIdStart = &v
}

// GetSlotId returns the SlotId field value if set, zero value otherwise.
func (o *FabricPortModeAllOf) GetSlotId() int64 {
	if o == nil || o.SlotId == nil {
		var ret int64
		return ret
	}
	return *o.SlotId
}

// GetSlotIdOk returns a tuple with the SlotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricPortModeAllOf) GetSlotIdOk() (*int64, bool) {
	if o == nil || o.SlotId == nil {
		return nil, false
	}
	return o.SlotId, true
}

// HasSlotId returns a boolean if a field has been set.
func (o *FabricPortModeAllOf) HasSlotId() bool {
	if o != nil && o.SlotId != nil {
		return true
	}

	return false
}

// SetSlotId gets a reference to the given int64 and assigns it to the SlotId field.
func (o *FabricPortModeAllOf) SetSlotId(v int64) {
	o.SlotId = &v
}

// GetPortPolicy returns the PortPolicy field value if set, zero value otherwise.
func (o *FabricPortModeAllOf) GetPortPolicy() FabricPortPolicyRelationship {
	if o == nil || o.PortPolicy == nil {
		var ret FabricPortPolicyRelationship
		return ret
	}
	return *o.PortPolicy
}

// GetPortPolicyOk returns a tuple with the PortPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricPortModeAllOf) GetPortPolicyOk() (*FabricPortPolicyRelationship, bool) {
	if o == nil || o.PortPolicy == nil {
		return nil, false
	}
	return o.PortPolicy, true
}

// HasPortPolicy returns a boolean if a field has been set.
func (o *FabricPortModeAllOf) HasPortPolicy() bool {
	if o != nil && o.PortPolicy != nil {
		return true
	}

	return false
}

// SetPortPolicy gets a reference to the given FabricPortPolicyRelationship and assigns it to the PortPolicy field.
func (o *FabricPortModeAllOf) SetPortPolicy(v FabricPortPolicyRelationship) {
	o.PortPolicy = &v
}

func (o FabricPortModeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CustomMode != nil {
		toSerialize["CustomMode"] = o.CustomMode
	}
	if o.PortIdEnd != nil {
		toSerialize["PortIdEnd"] = o.PortIdEnd
	}
	if o.PortIdStart != nil {
		toSerialize["PortIdStart"] = o.PortIdStart
	}
	if o.SlotId != nil {
		toSerialize["SlotId"] = o.SlotId
	}
	if o.PortPolicy != nil {
		toSerialize["PortPolicy"] = o.PortPolicy
	}
	return json.Marshal(toSerialize)
}

type NullableFabricPortModeAllOf struct {
	value *FabricPortModeAllOf
	isSet bool
}

func (v NullableFabricPortModeAllOf) Get() *FabricPortModeAllOf {
	return v.value
}

func (v *NullableFabricPortModeAllOf) Set(val *FabricPortModeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricPortModeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricPortModeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricPortModeAllOf(val *FabricPortModeAllOf) *NullableFabricPortModeAllOf {
	return &NullableFabricPortModeAllOf{value: val, isSet: true}
}

func (v NullableFabricPortModeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricPortModeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
