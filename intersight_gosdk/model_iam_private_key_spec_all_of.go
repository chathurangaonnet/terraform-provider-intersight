/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-06-19T15:15:17Z.
 *
 * API version: 1.0.9-1903
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
)

// IamPrivateKeySpecAllOf Definition of the list of properties defined in 'iam.PrivateKeySpec', excluding properties defined in parent classes.
type IamPrivateKeySpecAllOf struct {
	Algorithm          *PkixKeyGenerationSpec             `json:"Algorithm,omitempty"`
	CertificateRequest *IamCertificateRequestRelationship `json:"CertificateRequest,omitempty"`
}

// NewIamPrivateKeySpecAllOf instantiates a new IamPrivateKeySpecAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamPrivateKeySpecAllOf() *IamPrivateKeySpecAllOf {
	this := IamPrivateKeySpecAllOf{}
	return &this
}

// NewIamPrivateKeySpecAllOfWithDefaults instantiates a new IamPrivateKeySpecAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamPrivateKeySpecAllOfWithDefaults() *IamPrivateKeySpecAllOf {
	this := IamPrivateKeySpecAllOf{}
	return &this
}

// GetAlgorithm returns the Algorithm field value if set, zero value otherwise.
func (o *IamPrivateKeySpecAllOf) GetAlgorithm() PkixKeyGenerationSpec {
	if o == nil || o.Algorithm == nil {
		var ret PkixKeyGenerationSpec
		return ret
	}
	return *o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamPrivateKeySpecAllOf) GetAlgorithmOk() (*PkixKeyGenerationSpec, bool) {
	if o == nil || o.Algorithm == nil {
		return nil, false
	}
	return o.Algorithm, true
}

// HasAlgorithm returns a boolean if a field has been set.
func (o *IamPrivateKeySpecAllOf) HasAlgorithm() bool {
	if o != nil && o.Algorithm != nil {
		return true
	}

	return false
}

// SetAlgorithm gets a reference to the given PkixKeyGenerationSpec and assigns it to the Algorithm field.
func (o *IamPrivateKeySpecAllOf) SetAlgorithm(v PkixKeyGenerationSpec) {
	o.Algorithm = &v
}

// GetCertificateRequest returns the CertificateRequest field value if set, zero value otherwise.
func (o *IamPrivateKeySpecAllOf) GetCertificateRequest() IamCertificateRequestRelationship {
	if o == nil || o.CertificateRequest == nil {
		var ret IamCertificateRequestRelationship
		return ret
	}
	return *o.CertificateRequest
}

// GetCertificateRequestOk returns a tuple with the CertificateRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamPrivateKeySpecAllOf) GetCertificateRequestOk() (*IamCertificateRequestRelationship, bool) {
	if o == nil || o.CertificateRequest == nil {
		return nil, false
	}
	return o.CertificateRequest, true
}

// HasCertificateRequest returns a boolean if a field has been set.
func (o *IamPrivateKeySpecAllOf) HasCertificateRequest() bool {
	if o != nil && o.CertificateRequest != nil {
		return true
	}

	return false
}

// SetCertificateRequest gets a reference to the given IamCertificateRequestRelationship and assigns it to the CertificateRequest field.
func (o *IamPrivateKeySpecAllOf) SetCertificateRequest(v IamCertificateRequestRelationship) {
	o.CertificateRequest = &v
}

func (o IamPrivateKeySpecAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Algorithm != nil {
		toSerialize["Algorithm"] = o.Algorithm
	}
	if o.CertificateRequest != nil {
		toSerialize["CertificateRequest"] = o.CertificateRequest
	}
	return json.Marshal(toSerialize)
}

type NullableIamPrivateKeySpecAllOf struct {
	value *IamPrivateKeySpecAllOf
	isSet bool
}

func (v NullableIamPrivateKeySpecAllOf) Get() *IamPrivateKeySpecAllOf {
	return v.value
}

func (v *NullableIamPrivateKeySpecAllOf) Set(val *IamPrivateKeySpecAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIamPrivateKeySpecAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIamPrivateKeySpecAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamPrivateKeySpecAllOf(val *IamPrivateKeySpecAllOf) *NullableIamPrivateKeySpecAllOf {
	return &NullableIamPrivateKeySpecAllOf{value: val, isSet: true}
}

func (v NullableIamPrivateKeySpecAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamPrivateKeySpecAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}