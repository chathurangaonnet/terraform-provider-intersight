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

// HyperflexHxZoneResiliencyInfoDt struct for HyperflexHxZoneResiliencyInfoDt
type HyperflexHxZoneResiliencyInfoDt struct {
	MoBaseComplexType
	Name                 *string                      `json:"Name,omitempty"`
	ResiliencyInfo       *HyperflexHxResiliencyInfoDt `json:"ResiliencyInfo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexHxZoneResiliencyInfoDt HyperflexHxZoneResiliencyInfoDt

// NewHyperflexHxZoneResiliencyInfoDt instantiates a new HyperflexHxZoneResiliencyInfoDt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexHxZoneResiliencyInfoDt() *HyperflexHxZoneResiliencyInfoDt {
	this := HyperflexHxZoneResiliencyInfoDt{}
	return &this
}

// NewHyperflexHxZoneResiliencyInfoDtWithDefaults instantiates a new HyperflexHxZoneResiliencyInfoDt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexHxZoneResiliencyInfoDtWithDefaults() *HyperflexHxZoneResiliencyInfoDt {
	this := HyperflexHxZoneResiliencyInfoDt{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HyperflexHxZoneResiliencyInfoDt) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxZoneResiliencyInfoDt) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HyperflexHxZoneResiliencyInfoDt) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HyperflexHxZoneResiliencyInfoDt) SetName(v string) {
	o.Name = &v
}

// GetResiliencyInfo returns the ResiliencyInfo field value if set, zero value otherwise.
func (o *HyperflexHxZoneResiliencyInfoDt) GetResiliencyInfo() HyperflexHxResiliencyInfoDt {
	if o == nil || o.ResiliencyInfo == nil {
		var ret HyperflexHxResiliencyInfoDt
		return ret
	}
	return *o.ResiliencyInfo
}

// GetResiliencyInfoOk returns a tuple with the ResiliencyInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxZoneResiliencyInfoDt) GetResiliencyInfoOk() (*HyperflexHxResiliencyInfoDt, bool) {
	if o == nil || o.ResiliencyInfo == nil {
		return nil, false
	}
	return o.ResiliencyInfo, true
}

// HasResiliencyInfo returns a boolean if a field has been set.
func (o *HyperflexHxZoneResiliencyInfoDt) HasResiliencyInfo() bool {
	if o != nil && o.ResiliencyInfo != nil {
		return true
	}

	return false
}

// SetResiliencyInfo gets a reference to the given HyperflexHxResiliencyInfoDt and assigns it to the ResiliencyInfo field.
func (o *HyperflexHxZoneResiliencyInfoDt) SetResiliencyInfo(v HyperflexHxResiliencyInfoDt) {
	o.ResiliencyInfo = &v
}

func (o HyperflexHxZoneResiliencyInfoDt) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.ResiliencyInfo != nil {
		toSerialize["ResiliencyInfo"] = o.ResiliencyInfo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexHxZoneResiliencyInfoDt) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexHxZoneResiliencyInfoDt := _HyperflexHxZoneResiliencyInfoDt{}

	if err = json.Unmarshal(bytes, &varHyperflexHxZoneResiliencyInfoDt); err == nil {
		*o = HyperflexHxZoneResiliencyInfoDt(varHyperflexHxZoneResiliencyInfoDt)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Name")
		delete(additionalProperties, "ResiliencyInfo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexHxZoneResiliencyInfoDt struct {
	value *HyperflexHxZoneResiliencyInfoDt
	isSet bool
}

func (v NullableHyperflexHxZoneResiliencyInfoDt) Get() *HyperflexHxZoneResiliencyInfoDt {
	return v.value
}

func (v *NullableHyperflexHxZoneResiliencyInfoDt) Set(val *HyperflexHxZoneResiliencyInfoDt) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHxZoneResiliencyInfoDt) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHxZoneResiliencyInfoDt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHxZoneResiliencyInfoDt(val *HyperflexHxZoneResiliencyInfoDt) *NullableHyperflexHxZoneResiliencyInfoDt {
	return &NullableHyperflexHxZoneResiliencyInfoDt{value: val, isSet: true}
}

func (v NullableHyperflexHxZoneResiliencyInfoDt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHxZoneResiliencyInfoDt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
