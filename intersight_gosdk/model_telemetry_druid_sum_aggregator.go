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

// TelemetryDruidSumAggregator Computes the sum of values.
type TelemetryDruidSumAggregator struct {
	TelemetryDruidBaseAggregator
	// Output name for the summed value.
	Name string `json:"name"`
	// Name of the metric column to sum over.
	FieldName string `json:"fieldName"`
}

// NewTelemetryDruidSumAggregator instantiates a new TelemetryDruidSumAggregator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidSumAggregator(name string, fieldName string) *TelemetryDruidSumAggregator {
	this := TelemetryDruidSumAggregator{}
	this.Name = name
	this.FieldName = fieldName
	return &this
}

// NewTelemetryDruidSumAggregatorWithDefaults instantiates a new TelemetryDruidSumAggregator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidSumAggregatorWithDefaults() *TelemetryDruidSumAggregator {
	this := TelemetryDruidSumAggregator{}
	return &this
}

// GetName returns the Name field value
func (o *TelemetryDruidSumAggregator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidSumAggregator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TelemetryDruidSumAggregator) SetName(v string) {
	o.Name = v
}

// GetFieldName returns the FieldName field value
func (o *TelemetryDruidSumAggregator) GetFieldName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldName
}

// GetFieldNameOk returns a tuple with the FieldName field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidSumAggregator) GetFieldNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldName, true
}

// SetFieldName sets field value
func (o *TelemetryDruidSumAggregator) SetFieldName(v string) {
	o.FieldName = v
}

func (o TelemetryDruidSumAggregator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedTelemetryDruidBaseAggregator, errTelemetryDruidBaseAggregator := json.Marshal(o.TelemetryDruidBaseAggregator)
	if errTelemetryDruidBaseAggregator != nil {
		return []byte{}, errTelemetryDruidBaseAggregator
	}
	errTelemetryDruidBaseAggregator = json.Unmarshal([]byte(serializedTelemetryDruidBaseAggregator), &toSerialize)
	if errTelemetryDruidBaseAggregator != nil {
		return []byte{}, errTelemetryDruidBaseAggregator
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["fieldName"] = o.FieldName
	}
	return json.Marshal(toSerialize)
}

// AsTelemetryDruidAggregator wraps this instance of TelemetryDruidSumAggregator in TelemetryDruidAggregator
func (s *TelemetryDruidSumAggregator) AsTelemetryDruidAggregator() TelemetryDruidAggregator {
	return TelemetryDruidAggregator{TelemetryDruidAggregatorInterface: s}
}

type NullableTelemetryDruidSumAggregator struct {
	value *TelemetryDruidSumAggregator
	isSet bool
}

func (v NullableTelemetryDruidSumAggregator) Get() *TelemetryDruidSumAggregator {
	return v.value
}

func (v *NullableTelemetryDruidSumAggregator) Set(val *TelemetryDruidSumAggregator) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidSumAggregator) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidSumAggregator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidSumAggregator(val *TelemetryDruidSumAggregator) *NullableTelemetryDruidSumAggregator {
	return &NullableTelemetryDruidSumAggregator{value: val, isSet: true}
}

func (v NullableTelemetryDruidSumAggregator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidSumAggregator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
