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

// TelemetryDruidInvertedTopNMetricSpec Sort dimension values in inverted order, i.e inverts the order of the delegate metric spec. It can be used to sort the values in ascending order.
type TelemetryDruidInvertedTopNMetricSpec struct {
	TelemetryDruidBaseTopNMetricSpec
	Metric TelemetryDruidTopNMetricSpec `json:"metric"`
}

// NewTelemetryDruidInvertedTopNMetricSpec instantiates a new TelemetryDruidInvertedTopNMetricSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidInvertedTopNMetricSpec(metric TelemetryDruidTopNMetricSpec) *TelemetryDruidInvertedTopNMetricSpec {
	this := TelemetryDruidInvertedTopNMetricSpec{}
	this.Metric = metric
	return &this
}

// NewTelemetryDruidInvertedTopNMetricSpecWithDefaults instantiates a new TelemetryDruidInvertedTopNMetricSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidInvertedTopNMetricSpecWithDefaults() *TelemetryDruidInvertedTopNMetricSpec {
	this := TelemetryDruidInvertedTopNMetricSpec{}
	return &this
}

// GetMetric returns the Metric field value
func (o *TelemetryDruidInvertedTopNMetricSpec) GetMetric() TelemetryDruidTopNMetricSpec {
	if o == nil {
		var ret TelemetryDruidTopNMetricSpec
		return ret
	}

	return o.Metric
}

// GetMetricOk returns a tuple with the Metric field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidInvertedTopNMetricSpec) GetMetricOk() (*TelemetryDruidTopNMetricSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metric, true
}

// SetMetric sets field value
func (o *TelemetryDruidInvertedTopNMetricSpec) SetMetric(v TelemetryDruidTopNMetricSpec) {
	o.Metric = v
}

func (o TelemetryDruidInvertedTopNMetricSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedTelemetryDruidBaseTopNMetricSpec, errTelemetryDruidBaseTopNMetricSpec := json.Marshal(o.TelemetryDruidBaseTopNMetricSpec)
	if errTelemetryDruidBaseTopNMetricSpec != nil {
		return []byte{}, errTelemetryDruidBaseTopNMetricSpec
	}
	errTelemetryDruidBaseTopNMetricSpec = json.Unmarshal([]byte(serializedTelemetryDruidBaseTopNMetricSpec), &toSerialize)
	if errTelemetryDruidBaseTopNMetricSpec != nil {
		return []byte{}, errTelemetryDruidBaseTopNMetricSpec
	}
	if true {
		toSerialize["metric"] = o.Metric
	}
	return json.Marshal(toSerialize)
}

// AsTelemetryDruidTopNMetricSpec wraps this instance of TelemetryDruidInvertedTopNMetricSpec in TelemetryDruidTopNMetricSpec
func (s *TelemetryDruidInvertedTopNMetricSpec) AsTelemetryDruidTopNMetricSpec() TelemetryDruidTopNMetricSpec {
	return TelemetryDruidTopNMetricSpec{TelemetryDruidTopNMetricSpecInterface: s}
}

type NullableTelemetryDruidInvertedTopNMetricSpec struct {
	value *TelemetryDruidInvertedTopNMetricSpec
	isSet bool
}

func (v NullableTelemetryDruidInvertedTopNMetricSpec) Get() *TelemetryDruidInvertedTopNMetricSpec {
	return v.value
}

func (v *NullableTelemetryDruidInvertedTopNMetricSpec) Set(val *TelemetryDruidInvertedTopNMetricSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidInvertedTopNMetricSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidInvertedTopNMetricSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidInvertedTopNMetricSpec(val *TelemetryDruidInvertedTopNMetricSpec) *NullableTelemetryDruidInvertedTopNMetricSpec {
	return &NullableTelemetryDruidInvertedTopNMetricSpec{value: val, isSet: true}
}

func (v NullableTelemetryDruidInvertedTopNMetricSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidInvertedTopNMetricSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
