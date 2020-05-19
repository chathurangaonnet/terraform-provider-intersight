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

// PortGroupAllOf Definition of the list of properties defined in 'port.Group', excluding properties defined in parent classes.
type PortGroupAllOf struct {
	Transport               *string                              `json:"Transport,omitempty"`
	EquipmentSharedIoModule *EquipmentSharedIoModuleRelationship `json:"EquipmentSharedIoModule,omitempty"`
	EquipmentSwitchCard     *EquipmentSwitchCardRelationship     `json:"EquipmentSwitchCard,omitempty"`
	// An array of relationships to etherPhysicalPort resources.
	EthernetPorts *[]EtherPhysicalPortRelationship `json:"EthernetPorts,omitempty"`
	// An array of relationships to fcPhysicalPort resources.
	FcPorts          *[]FcPhysicalPortRelationship        `json:"FcPorts,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	// An array of relationships to portSubGroup resources.
	SubGroups *[]PortSubGroupRelationship `json:"SubGroups,omitempty"`
}

// NewPortGroupAllOf instantiates a new PortGroupAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortGroupAllOf() *PortGroupAllOf {
	this := PortGroupAllOf{}
	return &this
}

// NewPortGroupAllOfWithDefaults instantiates a new PortGroupAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortGroupAllOfWithDefaults() *PortGroupAllOf {
	this := PortGroupAllOf{}
	return &this
}

// GetTransport returns the Transport field value if set, zero value otherwise.
func (o *PortGroupAllOf) GetTransport() string {
	if o == nil || o.Transport == nil {
		var ret string
		return ret
	}
	return *o.Transport
}

// GetTransportOk returns a tuple with the Transport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortGroupAllOf) GetTransportOk() (*string, bool) {
	if o == nil || o.Transport == nil {
		return nil, false
	}
	return o.Transport, true
}

// HasTransport returns a boolean if a field has been set.
func (o *PortGroupAllOf) HasTransport() bool {
	if o != nil && o.Transport != nil {
		return true
	}

	return false
}

// SetTransport gets a reference to the given string and assigns it to the Transport field.
func (o *PortGroupAllOf) SetTransport(v string) {
	o.Transport = &v
}

// GetEquipmentSharedIoModule returns the EquipmentSharedIoModule field value if set, zero value otherwise.
func (o *PortGroupAllOf) GetEquipmentSharedIoModule() EquipmentSharedIoModuleRelationship {
	if o == nil || o.EquipmentSharedIoModule == nil {
		var ret EquipmentSharedIoModuleRelationship
		return ret
	}
	return *o.EquipmentSharedIoModule
}

// GetEquipmentSharedIoModuleOk returns a tuple with the EquipmentSharedIoModule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortGroupAllOf) GetEquipmentSharedIoModuleOk() (*EquipmentSharedIoModuleRelationship, bool) {
	if o == nil || o.EquipmentSharedIoModule == nil {
		return nil, false
	}
	return o.EquipmentSharedIoModule, true
}

// HasEquipmentSharedIoModule returns a boolean if a field has been set.
func (o *PortGroupAllOf) HasEquipmentSharedIoModule() bool {
	if o != nil && o.EquipmentSharedIoModule != nil {
		return true
	}

	return false
}

// SetEquipmentSharedIoModule gets a reference to the given EquipmentSharedIoModuleRelationship and assigns it to the EquipmentSharedIoModule field.
func (o *PortGroupAllOf) SetEquipmentSharedIoModule(v EquipmentSharedIoModuleRelationship) {
	o.EquipmentSharedIoModule = &v
}

// GetEquipmentSwitchCard returns the EquipmentSwitchCard field value if set, zero value otherwise.
func (o *PortGroupAllOf) GetEquipmentSwitchCard() EquipmentSwitchCardRelationship {
	if o == nil || o.EquipmentSwitchCard == nil {
		var ret EquipmentSwitchCardRelationship
		return ret
	}
	return *o.EquipmentSwitchCard
}

// GetEquipmentSwitchCardOk returns a tuple with the EquipmentSwitchCard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortGroupAllOf) GetEquipmentSwitchCardOk() (*EquipmentSwitchCardRelationship, bool) {
	if o == nil || o.EquipmentSwitchCard == nil {
		return nil, false
	}
	return o.EquipmentSwitchCard, true
}

// HasEquipmentSwitchCard returns a boolean if a field has been set.
func (o *PortGroupAllOf) HasEquipmentSwitchCard() bool {
	if o != nil && o.EquipmentSwitchCard != nil {
		return true
	}

	return false
}

// SetEquipmentSwitchCard gets a reference to the given EquipmentSwitchCardRelationship and assigns it to the EquipmentSwitchCard field.
func (o *PortGroupAllOf) SetEquipmentSwitchCard(v EquipmentSwitchCardRelationship) {
	o.EquipmentSwitchCard = &v
}

// GetEthernetPorts returns the EthernetPorts field value if set, zero value otherwise.
func (o *PortGroupAllOf) GetEthernetPorts() []EtherPhysicalPortRelationship {
	if o == nil || o.EthernetPorts == nil {
		var ret []EtherPhysicalPortRelationship
		return ret
	}
	return *o.EthernetPorts
}

// GetEthernetPortsOk returns a tuple with the EthernetPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortGroupAllOf) GetEthernetPortsOk() (*[]EtherPhysicalPortRelationship, bool) {
	if o == nil || o.EthernetPorts == nil {
		return nil, false
	}
	return o.EthernetPorts, true
}

// HasEthernetPorts returns a boolean if a field has been set.
func (o *PortGroupAllOf) HasEthernetPorts() bool {
	if o != nil && o.EthernetPorts != nil {
		return true
	}

	return false
}

// SetEthernetPorts gets a reference to the given []EtherPhysicalPortRelationship and assigns it to the EthernetPorts field.
func (o *PortGroupAllOf) SetEthernetPorts(v []EtherPhysicalPortRelationship) {
	o.EthernetPorts = &v
}

// GetFcPorts returns the FcPorts field value if set, zero value otherwise.
func (o *PortGroupAllOf) GetFcPorts() []FcPhysicalPortRelationship {
	if o == nil || o.FcPorts == nil {
		var ret []FcPhysicalPortRelationship
		return ret
	}
	return *o.FcPorts
}

// GetFcPortsOk returns a tuple with the FcPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortGroupAllOf) GetFcPortsOk() (*[]FcPhysicalPortRelationship, bool) {
	if o == nil || o.FcPorts == nil {
		return nil, false
	}
	return o.FcPorts, true
}

// HasFcPorts returns a boolean if a field has been set.
func (o *PortGroupAllOf) HasFcPorts() bool {
	if o != nil && o.FcPorts != nil {
		return true
	}

	return false
}

// SetFcPorts gets a reference to the given []FcPhysicalPortRelationship and assigns it to the FcPorts field.
func (o *PortGroupAllOf) SetFcPorts(v []FcPhysicalPortRelationship) {
	o.FcPorts = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *PortGroupAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortGroupAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *PortGroupAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *PortGroupAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetSubGroups returns the SubGroups field value if set, zero value otherwise.
func (o *PortGroupAllOf) GetSubGroups() []PortSubGroupRelationship {
	if o == nil || o.SubGroups == nil {
		var ret []PortSubGroupRelationship
		return ret
	}
	return *o.SubGroups
}

// GetSubGroupsOk returns a tuple with the SubGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortGroupAllOf) GetSubGroupsOk() (*[]PortSubGroupRelationship, bool) {
	if o == nil || o.SubGroups == nil {
		return nil, false
	}
	return o.SubGroups, true
}

// HasSubGroups returns a boolean if a field has been set.
func (o *PortGroupAllOf) HasSubGroups() bool {
	if o != nil && o.SubGroups != nil {
		return true
	}

	return false
}

// SetSubGroups gets a reference to the given []PortSubGroupRelationship and assigns it to the SubGroups field.
func (o *PortGroupAllOf) SetSubGroups(v []PortSubGroupRelationship) {
	o.SubGroups = &v
}

func (o PortGroupAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Transport != nil {
		toSerialize["Transport"] = o.Transport
	}
	if o.EquipmentSharedIoModule != nil {
		toSerialize["EquipmentSharedIoModule"] = o.EquipmentSharedIoModule
	}
	if o.EquipmentSwitchCard != nil {
		toSerialize["EquipmentSwitchCard"] = o.EquipmentSwitchCard
	}
	if o.EthernetPorts != nil {
		toSerialize["EthernetPorts"] = o.EthernetPorts
	}
	if o.FcPorts != nil {
		toSerialize["FcPorts"] = o.FcPorts
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.SubGroups != nil {
		toSerialize["SubGroups"] = o.SubGroups
	}
	return json.Marshal(toSerialize)
}

type NullablePortGroupAllOf struct {
	value *PortGroupAllOf
	isSet bool
}

func (v NullablePortGroupAllOf) Get() *PortGroupAllOf {
	return v.value
}

func (v *NullablePortGroupAllOf) Set(val *PortGroupAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePortGroupAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePortGroupAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortGroupAllOf(val *PortGroupAllOf) *NullablePortGroupAllOf {
	return &NullablePortGroupAllOf{value: val, isSet: true}
}

func (v NullablePortGroupAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortGroupAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
