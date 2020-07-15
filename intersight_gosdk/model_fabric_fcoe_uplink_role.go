/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-07-11T05:47:33Z.
 *
 * API version: 1.0.9-1999
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// FabricFcoeUplinkRole Configuration object sent by user to create a fcoe uplink port.
type FabricFcoeUplinkRole struct {
	FabricPortRole
	// Admin configured speed for the port.
	AdminSpeed *string `json:"AdminSpeed,omitempty"`
	// Forward error correction configuration for the port.
	Fec *string `json:"Fec,omitempty"`
	// Admin configured state for UDLD for this port.
	UdldAdminState       *string `json:"UdldAdminState,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricFcoeUplinkRole FabricFcoeUplinkRole

// NewFabricFcoeUplinkRole instantiates a new FabricFcoeUplinkRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricFcoeUplinkRole() *FabricFcoeUplinkRole {
	this := FabricFcoeUplinkRole{}
	var adminSpeed string = "Auto"
	this.AdminSpeed = &adminSpeed
	var fec string = "Auto"
	this.Fec = &fec
	var udldAdminState string = "Disabled"
	this.UdldAdminState = &udldAdminState
	return &this
}

// NewFabricFcoeUplinkRoleWithDefaults instantiates a new FabricFcoeUplinkRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricFcoeUplinkRoleWithDefaults() *FabricFcoeUplinkRole {
	this := FabricFcoeUplinkRole{}
	var adminSpeed string = "Auto"
	this.AdminSpeed = &adminSpeed
	var fec string = "Auto"
	this.Fec = &fec
	var udldAdminState string = "Disabled"
	this.UdldAdminState = &udldAdminState
	return &this
}

// GetAdminSpeed returns the AdminSpeed field value if set, zero value otherwise.
func (o *FabricFcoeUplinkRole) GetAdminSpeed() string {
	if o == nil || o.AdminSpeed == nil {
		var ret string
		return ret
	}
	return *o.AdminSpeed
}

// GetAdminSpeedOk returns a tuple with the AdminSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricFcoeUplinkRole) GetAdminSpeedOk() (*string, bool) {
	if o == nil || o.AdminSpeed == nil {
		return nil, false
	}
	return o.AdminSpeed, true
}

// HasAdminSpeed returns a boolean if a field has been set.
func (o *FabricFcoeUplinkRole) HasAdminSpeed() bool {
	if o != nil && o.AdminSpeed != nil {
		return true
	}

	return false
}

// SetAdminSpeed gets a reference to the given string and assigns it to the AdminSpeed field.
func (o *FabricFcoeUplinkRole) SetAdminSpeed(v string) {
	o.AdminSpeed = &v
}

// GetFec returns the Fec field value if set, zero value otherwise.
func (o *FabricFcoeUplinkRole) GetFec() string {
	if o == nil || o.Fec == nil {
		var ret string
		return ret
	}
	return *o.Fec
}

// GetFecOk returns a tuple with the Fec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricFcoeUplinkRole) GetFecOk() (*string, bool) {
	if o == nil || o.Fec == nil {
		return nil, false
	}
	return o.Fec, true
}

// HasFec returns a boolean if a field has been set.
func (o *FabricFcoeUplinkRole) HasFec() bool {
	if o != nil && o.Fec != nil {
		return true
	}

	return false
}

// SetFec gets a reference to the given string and assigns it to the Fec field.
func (o *FabricFcoeUplinkRole) SetFec(v string) {
	o.Fec = &v
}

// GetUdldAdminState returns the UdldAdminState field value if set, zero value otherwise.
func (o *FabricFcoeUplinkRole) GetUdldAdminState() string {
	if o == nil || o.UdldAdminState == nil {
		var ret string
		return ret
	}
	return *o.UdldAdminState
}

// GetUdldAdminStateOk returns a tuple with the UdldAdminState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricFcoeUplinkRole) GetUdldAdminStateOk() (*string, bool) {
	if o == nil || o.UdldAdminState == nil {
		return nil, false
	}
	return o.UdldAdminState, true
}

// HasUdldAdminState returns a boolean if a field has been set.
func (o *FabricFcoeUplinkRole) HasUdldAdminState() bool {
	if o != nil && o.UdldAdminState != nil {
		return true
	}

	return false
}

// SetUdldAdminState gets a reference to the given string and assigns it to the UdldAdminState field.
func (o *FabricFcoeUplinkRole) SetUdldAdminState(v string) {
	o.UdldAdminState = &v
}

func (o FabricFcoeUplinkRole) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedFabricPortRole, errFabricPortRole := json.Marshal(o.FabricPortRole)
	if errFabricPortRole != nil {
		return []byte{}, errFabricPortRole
	}
	errFabricPortRole = json.Unmarshal([]byte(serializedFabricPortRole), &toSerialize)
	if errFabricPortRole != nil {
		return []byte{}, errFabricPortRole
	}
	if o.AdminSpeed != nil {
		toSerialize["AdminSpeed"] = o.AdminSpeed
	}
	if o.Fec != nil {
		toSerialize["Fec"] = o.Fec
	}
	if o.UdldAdminState != nil {
		toSerialize["UdldAdminState"] = o.UdldAdminState
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricFcoeUplinkRole) UnmarshalJSON(bytes []byte) (err error) {
	type FabricFcoeUplinkRoleWithoutEmbeddedStruct struct {
		// Admin configured speed for the port.
		AdminSpeed *string `json:"AdminSpeed,omitempty"`
		// Forward error correction configuration for the port.
		Fec *string `json:"Fec,omitempty"`
		// Admin configured state for UDLD for this port.
		UdldAdminState *string `json:"UdldAdminState,omitempty"`
	}

	varFabricFcoeUplinkRoleWithoutEmbeddedStruct := FabricFcoeUplinkRoleWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFabricFcoeUplinkRoleWithoutEmbeddedStruct)
	if err == nil {
		varFabricFcoeUplinkRole := _FabricFcoeUplinkRole{}
		varFabricFcoeUplinkRole.AdminSpeed = varFabricFcoeUplinkRoleWithoutEmbeddedStruct.AdminSpeed
		varFabricFcoeUplinkRole.Fec = varFabricFcoeUplinkRoleWithoutEmbeddedStruct.Fec
		varFabricFcoeUplinkRole.UdldAdminState = varFabricFcoeUplinkRoleWithoutEmbeddedStruct.UdldAdminState
		*o = FabricFcoeUplinkRole(varFabricFcoeUplinkRole)
	} else {
		return err
	}

	varFabricFcoeUplinkRole := _FabricFcoeUplinkRole{}

	err = json.Unmarshal(bytes, &varFabricFcoeUplinkRole)
	if err == nil {
		o.FabricPortRole = varFabricFcoeUplinkRole.FabricPortRole
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "AdminSpeed")
		delete(additionalProperties, "Fec")
		delete(additionalProperties, "UdldAdminState")

		// remove fields from embedded structs
		reflectFabricPortRole := reflect.ValueOf(o.FabricPortRole)
		for i := 0; i < reflectFabricPortRole.Type().NumField(); i++ {
			t := reflectFabricPortRole.Type().Field(i)

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

type NullableFabricFcoeUplinkRole struct {
	value *FabricFcoeUplinkRole
	isSet bool
}

func (v NullableFabricFcoeUplinkRole) Get() *FabricFcoeUplinkRole {
	return v.value
}

func (v *NullableFabricFcoeUplinkRole) Set(val *FabricFcoeUplinkRole) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricFcoeUplinkRole) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricFcoeUplinkRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricFcoeUplinkRole(val *FabricFcoeUplinkRole) *NullableFabricFcoeUplinkRole {
	return &NullableFabricFcoeUplinkRole{value: val, isSet: true}
}

func (v NullableFabricFcoeUplinkRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricFcoeUplinkRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
