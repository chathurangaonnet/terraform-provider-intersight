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

// VnicEthQosPolicyAllOf Definition of the list of properties defined in 'vnic.EthQosPolicy', excluding properties defined in parent classes.
type VnicEthQosPolicyAllOf struct {
	// Class of Service to be associated to the traffic on the virtual interface.
	Cos *int64 `json:"Cos,omitempty"`
	// The Maximum Transmission Unit (MTU) or packet size that the virtual interface accepts.
	Mtu *int64 `json:"Mtu,omitempty"`
	// The value in Mbps (0-100000) to use for limiting the data rate on the virtual interface. Setting this to zero will turn rate limiting off.
	RateLimit *int64 `json:"RateLimit,omitempty"`
	// Enables usage of the Class of Service provided by the operating system.
	TrustHostCos *bool                                 `json:"TrustHostCos,omitempty"`
	Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
}

// NewVnicEthQosPolicyAllOf instantiates a new VnicEthQosPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicEthQosPolicyAllOf() *VnicEthQosPolicyAllOf {
	this := VnicEthQosPolicyAllOf{}
	return &this
}

// NewVnicEthQosPolicyAllOfWithDefaults instantiates a new VnicEthQosPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicEthQosPolicyAllOfWithDefaults() *VnicEthQosPolicyAllOf {
	this := VnicEthQosPolicyAllOf{}
	return &this
}

// GetCos returns the Cos field value if set, zero value otherwise.
func (o *VnicEthQosPolicyAllOf) GetCos() int64 {
	if o == nil || o.Cos == nil {
		var ret int64
		return ret
	}
	return *o.Cos
}

// GetCosOk returns a tuple with the Cos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthQosPolicyAllOf) GetCosOk() (*int64, bool) {
	if o == nil || o.Cos == nil {
		return nil, false
	}
	return o.Cos, true
}

// HasCos returns a boolean if a field has been set.
func (o *VnicEthQosPolicyAllOf) HasCos() bool {
	if o != nil && o.Cos != nil {
		return true
	}

	return false
}

// SetCos gets a reference to the given int64 and assigns it to the Cos field.
func (o *VnicEthQosPolicyAllOf) SetCos(v int64) {
	o.Cos = &v
}

// GetMtu returns the Mtu field value if set, zero value otherwise.
func (o *VnicEthQosPolicyAllOf) GetMtu() int64 {
	if o == nil || o.Mtu == nil {
		var ret int64
		return ret
	}
	return *o.Mtu
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthQosPolicyAllOf) GetMtuOk() (*int64, bool) {
	if o == nil || o.Mtu == nil {
		return nil, false
	}
	return o.Mtu, true
}

// HasMtu returns a boolean if a field has been set.
func (o *VnicEthQosPolicyAllOf) HasMtu() bool {
	if o != nil && o.Mtu != nil {
		return true
	}

	return false
}

// SetMtu gets a reference to the given int64 and assigns it to the Mtu field.
func (o *VnicEthQosPolicyAllOf) SetMtu(v int64) {
	o.Mtu = &v
}

// GetRateLimit returns the RateLimit field value if set, zero value otherwise.
func (o *VnicEthQosPolicyAllOf) GetRateLimit() int64 {
	if o == nil || o.RateLimit == nil {
		var ret int64
		return ret
	}
	return *o.RateLimit
}

// GetRateLimitOk returns a tuple with the RateLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthQosPolicyAllOf) GetRateLimitOk() (*int64, bool) {
	if o == nil || o.RateLimit == nil {
		return nil, false
	}
	return o.RateLimit, true
}

// HasRateLimit returns a boolean if a field has been set.
func (o *VnicEthQosPolicyAllOf) HasRateLimit() bool {
	if o != nil && o.RateLimit != nil {
		return true
	}

	return false
}

// SetRateLimit gets a reference to the given int64 and assigns it to the RateLimit field.
func (o *VnicEthQosPolicyAllOf) SetRateLimit(v int64) {
	o.RateLimit = &v
}

// GetTrustHostCos returns the TrustHostCos field value if set, zero value otherwise.
func (o *VnicEthQosPolicyAllOf) GetTrustHostCos() bool {
	if o == nil || o.TrustHostCos == nil {
		var ret bool
		return ret
	}
	return *o.TrustHostCos
}

// GetTrustHostCosOk returns a tuple with the TrustHostCos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthQosPolicyAllOf) GetTrustHostCosOk() (*bool, bool) {
	if o == nil || o.TrustHostCos == nil {
		return nil, false
	}
	return o.TrustHostCos, true
}

// HasTrustHostCos returns a boolean if a field has been set.
func (o *VnicEthQosPolicyAllOf) HasTrustHostCos() bool {
	if o != nil && o.TrustHostCos != nil {
		return true
	}

	return false
}

// SetTrustHostCos gets a reference to the given bool and assigns it to the TrustHostCos field.
func (o *VnicEthQosPolicyAllOf) SetTrustHostCos(v bool) {
	o.TrustHostCos = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *VnicEthQosPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthQosPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *VnicEthQosPolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *VnicEthQosPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o VnicEthQosPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cos != nil {
		toSerialize["Cos"] = o.Cos
	}
	if o.Mtu != nil {
		toSerialize["Mtu"] = o.Mtu
	}
	if o.RateLimit != nil {
		toSerialize["RateLimit"] = o.RateLimit
	}
	if o.TrustHostCos != nil {
		toSerialize["TrustHostCos"] = o.TrustHostCos
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	return json.Marshal(toSerialize)
}

type NullableVnicEthQosPolicyAllOf struct {
	value *VnicEthQosPolicyAllOf
	isSet bool
}

func (v NullableVnicEthQosPolicyAllOf) Get() *VnicEthQosPolicyAllOf {
	return v.value
}

func (v *NullableVnicEthQosPolicyAllOf) Set(val *VnicEthQosPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicEthQosPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicEthQosPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicEthQosPolicyAllOf(val *VnicEthQosPolicyAllOf) *NullableVnicEthQosPolicyAllOf {
	return &NullableVnicEthQosPolicyAllOf{value: val, isSet: true}
}

func (v NullableVnicEthQosPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicEthQosPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
