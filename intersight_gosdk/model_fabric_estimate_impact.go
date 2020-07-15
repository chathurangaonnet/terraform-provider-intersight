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

// FabricEstimateImpact Before submitting switch profile deploy operation, the estimate impact helps to know the list of components be impacted and require switch reboot.
type FabricEstimateImpact struct {
	MoBaseMo
	Profile              *FabricSwitchProfileRelationship `json:"Profile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricEstimateImpact FabricEstimateImpact

// NewFabricEstimateImpact instantiates a new FabricEstimateImpact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricEstimateImpact() *FabricEstimateImpact {
	this := FabricEstimateImpact{}
	return &this
}

// NewFabricEstimateImpactWithDefaults instantiates a new FabricEstimateImpact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricEstimateImpactWithDefaults() *FabricEstimateImpact {
	this := FabricEstimateImpact{}
	return &this
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *FabricEstimateImpact) GetProfile() FabricSwitchProfileRelationship {
	if o == nil || o.Profile == nil {
		var ret FabricSwitchProfileRelationship
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricEstimateImpact) GetProfileOk() (*FabricSwitchProfileRelationship, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *FabricEstimateImpact) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given FabricSwitchProfileRelationship and assigns it to the Profile field.
func (o *FabricEstimateImpact) SetProfile(v FabricSwitchProfileRelationship) {
	o.Profile = &v
}

func (o FabricEstimateImpact) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.Profile != nil {
		toSerialize["Profile"] = o.Profile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricEstimateImpact) UnmarshalJSON(bytes []byte) (err error) {
	type FabricEstimateImpactWithoutEmbeddedStruct struct {
		Profile *FabricSwitchProfileRelationship `json:"Profile,omitempty"`
	}

	varFabricEstimateImpactWithoutEmbeddedStruct := FabricEstimateImpactWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFabricEstimateImpactWithoutEmbeddedStruct)
	if err == nil {
		varFabricEstimateImpact := _FabricEstimateImpact{}
		varFabricEstimateImpact.Profile = varFabricEstimateImpactWithoutEmbeddedStruct.Profile
		*o = FabricEstimateImpact(varFabricEstimateImpact)
	} else {
		return err
	}

	varFabricEstimateImpact := _FabricEstimateImpact{}

	err = json.Unmarshal(bytes, &varFabricEstimateImpact)
	if err == nil {
		o.MoBaseMo = varFabricEstimateImpact.MoBaseMo
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Profile")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableFabricEstimateImpact struct {
	value *FabricEstimateImpact
	isSet bool
}

func (v NullableFabricEstimateImpact) Get() *FabricEstimateImpact {
	return v.value
}

func (v *NullableFabricEstimateImpact) Set(val *FabricEstimateImpact) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricEstimateImpact) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricEstimateImpact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricEstimateImpact(val *FabricEstimateImpact) *NullableFabricEstimateImpact {
	return &NullableFabricEstimateImpact{value: val, isSet: true}
}

func (v NullableFabricEstimateImpact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricEstimateImpact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
