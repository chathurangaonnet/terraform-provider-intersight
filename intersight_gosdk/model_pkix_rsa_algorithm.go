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

// PkixRsaAlgorithm The key pair is generated using the RSA algorithm and specified parameters.
type PkixRsaAlgorithm struct {
	PkixKeyGenerationSpec
	// The length of the RSA key, expressed in bits, for both public and private keys.
	Modulus *int32 `json:"Modulus,omitempty"`
}

// NewPkixRsaAlgorithm instantiates a new PkixRsaAlgorithm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPkixRsaAlgorithm() *PkixRsaAlgorithm {
	this := PkixRsaAlgorithm{}
	var modulus int32 = 2048
	this.Modulus = &modulus
	return &this
}

// NewPkixRsaAlgorithmWithDefaults instantiates a new PkixRsaAlgorithm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkixRsaAlgorithmWithDefaults() *PkixRsaAlgorithm {
	this := PkixRsaAlgorithm{}
	var modulus int32 = 2048
	this.Modulus = &modulus
	return &this
}

// GetModulus returns the Modulus field value if set, zero value otherwise.
func (o *PkixRsaAlgorithm) GetModulus() int32 {
	if o == nil || o.Modulus == nil {
		var ret int32
		return ret
	}
	return *o.Modulus
}

// GetModulusOk returns a tuple with the Modulus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkixRsaAlgorithm) GetModulusOk() (*int32, bool) {
	if o == nil || o.Modulus == nil {
		return nil, false
	}
	return o.Modulus, true
}

// HasModulus returns a boolean if a field has been set.
func (o *PkixRsaAlgorithm) HasModulus() bool {
	if o != nil && o.Modulus != nil {
		return true
	}

	return false
}

// SetModulus gets a reference to the given int32 and assigns it to the Modulus field.
func (o *PkixRsaAlgorithm) SetModulus(v int32) {
	o.Modulus = &v
}

func (o PkixRsaAlgorithm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPkixKeyGenerationSpec, errPkixKeyGenerationSpec := json.Marshal(o.PkixKeyGenerationSpec)
	if errPkixKeyGenerationSpec != nil {
		return []byte{}, errPkixKeyGenerationSpec
	}
	errPkixKeyGenerationSpec = json.Unmarshal([]byte(serializedPkixKeyGenerationSpec), &toSerialize)
	if errPkixKeyGenerationSpec != nil {
		return []byte{}, errPkixKeyGenerationSpec
	}
	if o.Modulus != nil {
		toSerialize["Modulus"] = o.Modulus
	}
	return json.Marshal(toSerialize)
}

type NullablePkixRsaAlgorithm struct {
	value *PkixRsaAlgorithm
	isSet bool
}

func (v NullablePkixRsaAlgorithm) Get() *PkixRsaAlgorithm {
	return v.value
}

func (v *NullablePkixRsaAlgorithm) Set(val *PkixRsaAlgorithm) {
	v.value = val
	v.isSet = true
}

func (v NullablePkixRsaAlgorithm) IsSet() bool {
	return v.isSet
}

func (v *NullablePkixRsaAlgorithm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePkixRsaAlgorithm(val *PkixRsaAlgorithm) *NullablePkixRsaAlgorithm {
	return &NullablePkixRsaAlgorithm{value: val, isSet: true}
}

func (v NullablePkixRsaAlgorithm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePkixRsaAlgorithm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
