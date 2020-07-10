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

// ConnectorFileChecksumAllOf Definition of the list of properties defined in 'connector.FileChecksum', excluding properties defined in parent classes.
type ConnectorFileChecksumAllOf struct {
	// The calculated hash of the contents using the algorithm.
	Hash *string `json:"Hash,omitempty"`
	// The hash algorithm used to calculate the checksum.
	HashAlgorithm *string `json:"HashAlgorithm,omitempty"`
}

// NewConnectorFileChecksumAllOf instantiates a new ConnectorFileChecksumAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorFileChecksumAllOf() *ConnectorFileChecksumAllOf {
	this := ConnectorFileChecksumAllOf{}
	var hashAlgorithm string = "crc"
	this.HashAlgorithm = &hashAlgorithm
	return &this
}

// NewConnectorFileChecksumAllOfWithDefaults instantiates a new ConnectorFileChecksumAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorFileChecksumAllOfWithDefaults() *ConnectorFileChecksumAllOf {
	this := ConnectorFileChecksumAllOf{}
	var hashAlgorithm string = "crc"
	this.HashAlgorithm = &hashAlgorithm
	return &this
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *ConnectorFileChecksumAllOf) GetHash() string {
	if o == nil || o.Hash == nil {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorFileChecksumAllOf) GetHashOk() (*string, bool) {
	if o == nil || o.Hash == nil {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *ConnectorFileChecksumAllOf) HasHash() bool {
	if o != nil && o.Hash != nil {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *ConnectorFileChecksumAllOf) SetHash(v string) {
	o.Hash = &v
}

// GetHashAlgorithm returns the HashAlgorithm field value if set, zero value otherwise.
func (o *ConnectorFileChecksumAllOf) GetHashAlgorithm() string {
	if o == nil || o.HashAlgorithm == nil {
		var ret string
		return ret
	}
	return *o.HashAlgorithm
}

// GetHashAlgorithmOk returns a tuple with the HashAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorFileChecksumAllOf) GetHashAlgorithmOk() (*string, bool) {
	if o == nil || o.HashAlgorithm == nil {
		return nil, false
	}
	return o.HashAlgorithm, true
}

// HasHashAlgorithm returns a boolean if a field has been set.
func (o *ConnectorFileChecksumAllOf) HasHashAlgorithm() bool {
	if o != nil && o.HashAlgorithm != nil {
		return true
	}

	return false
}

// SetHashAlgorithm gets a reference to the given string and assigns it to the HashAlgorithm field.
func (o *ConnectorFileChecksumAllOf) SetHashAlgorithm(v string) {
	o.HashAlgorithm = &v
}

func (o ConnectorFileChecksumAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Hash != nil {
		toSerialize["Hash"] = o.Hash
	}
	if o.HashAlgorithm != nil {
		toSerialize["HashAlgorithm"] = o.HashAlgorithm
	}
	return json.Marshal(toSerialize)
}

type NullableConnectorFileChecksumAllOf struct {
	value *ConnectorFileChecksumAllOf
	isSet bool
}

func (v NullableConnectorFileChecksumAllOf) Get() *ConnectorFileChecksumAllOf {
	return v.value
}

func (v *NullableConnectorFileChecksumAllOf) Set(val *ConnectorFileChecksumAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorFileChecksumAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorFileChecksumAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorFileChecksumAllOf(val *ConnectorFileChecksumAllOf) *NullableConnectorFileChecksumAllOf {
	return &NullableConnectorFileChecksumAllOf{value: val, isSet: true}
}

func (v NullableConnectorFileChecksumAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorFileChecksumAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
