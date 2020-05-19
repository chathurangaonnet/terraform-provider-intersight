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

// IamPrivateKeySpec Parameters used to generate a private key.
type IamPrivateKeySpec struct {
	MoBaseMo
	Algorithm          *PkixKeyGenerationSpec             `json:"Algorithm,omitempty"`
	CertificateRequest *IamCertificateRequestRelationship `json:"CertificateRequest,omitempty"`
}

// NewIamPrivateKeySpec instantiates a new IamPrivateKeySpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamPrivateKeySpec() *IamPrivateKeySpec {
	this := IamPrivateKeySpec{}
	return &this
}

// NewIamPrivateKeySpecWithDefaults instantiates a new IamPrivateKeySpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamPrivateKeySpecWithDefaults() *IamPrivateKeySpec {
	this := IamPrivateKeySpec{}
	return &this
}

// GetAlgorithm returns the Algorithm field value if set, zero value otherwise.
func (o *IamPrivateKeySpec) GetAlgorithm() PkixKeyGenerationSpec {
	if o == nil || o.Algorithm == nil {
		var ret PkixKeyGenerationSpec
		return ret
	}
	return *o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamPrivateKeySpec) GetAlgorithmOk() (*PkixKeyGenerationSpec, bool) {
	if o == nil || o.Algorithm == nil {
		return nil, false
	}
	return o.Algorithm, true
}

// HasAlgorithm returns a boolean if a field has been set.
func (o *IamPrivateKeySpec) HasAlgorithm() bool {
	if o != nil && o.Algorithm != nil {
		return true
	}

	return false
}

// SetAlgorithm gets a reference to the given PkixKeyGenerationSpec and assigns it to the Algorithm field.
func (o *IamPrivateKeySpec) SetAlgorithm(v PkixKeyGenerationSpec) {
	o.Algorithm = &v
}

// GetCertificateRequest returns the CertificateRequest field value if set, zero value otherwise.
func (o *IamPrivateKeySpec) GetCertificateRequest() IamCertificateRequestRelationship {
	if o == nil || o.CertificateRequest == nil {
		var ret IamCertificateRequestRelationship
		return ret
	}
	return *o.CertificateRequest
}

// GetCertificateRequestOk returns a tuple with the CertificateRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamPrivateKeySpec) GetCertificateRequestOk() (*IamCertificateRequestRelationship, bool) {
	if o == nil || o.CertificateRequest == nil {
		return nil, false
	}
	return o.CertificateRequest, true
}

// HasCertificateRequest returns a boolean if a field has been set.
func (o *IamPrivateKeySpec) HasCertificateRequest() bool {
	if o != nil && o.CertificateRequest != nil {
		return true
	}

	return false
}

// SetCertificateRequest gets a reference to the given IamCertificateRequestRelationship and assigns it to the CertificateRequest field.
func (o *IamPrivateKeySpec) SetCertificateRequest(v IamCertificateRequestRelationship) {
	o.CertificateRequest = &v
}

func (o IamPrivateKeySpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.Algorithm != nil {
		toSerialize["Algorithm"] = o.Algorithm
	}
	if o.CertificateRequest != nil {
		toSerialize["CertificateRequest"] = o.CertificateRequest
	}
	return json.Marshal(toSerialize)
}

// AsIamPrivateKeySpecRelationship wraps this instance of IamPrivateKeySpec in IamPrivateKeySpecRelationship
func (s *IamPrivateKeySpec) AsIamPrivateKeySpecRelationship() IamPrivateKeySpecRelationship {
	return IamPrivateKeySpecRelationship{IamPrivateKeySpecRelationshipInterface: s}
}

type NullableIamPrivateKeySpec struct {
	value *IamPrivateKeySpec
	isSet bool
}

func (v NullableIamPrivateKeySpec) Get() *IamPrivateKeySpec {
	return v.value
}

func (v *NullableIamPrivateKeySpec) Set(val *IamPrivateKeySpec) {
	v.value = val
	v.isSet = true
}

func (v NullableIamPrivateKeySpec) IsSet() bool {
	return v.isSet
}

func (v *NullableIamPrivateKeySpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamPrivateKeySpec(val *IamPrivateKeySpec) *NullableIamPrivateKeySpec {
	return &NullableIamPrivateKeySpec{value: val, isSet: true}
}

func (v NullableIamPrivateKeySpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamPrivateKeySpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
