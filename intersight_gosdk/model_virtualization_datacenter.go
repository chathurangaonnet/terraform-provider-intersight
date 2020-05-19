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

// VirtualizationDatacenter Common attributes of any datacenter. Serves as a base class for all concrete hypervisor types. Datacenter (DC) is the container that brings together all other elements (Host, Datastore, Virtual Machine) A typical DC has datastores for storage for Virtual Machines, and a handful of hosts to run the Virtual Machines. In addition, there could be virtual switches and portgroups in those switches.
type VirtualizationDatacenter struct {
	VirtualizationSourceDevice
	// Internally generated identity of this datacenter. This entity is not manipulated by users. It aids in uniquely identifying the datacenter object. For VMware, this is a MOR (managed object reference).
	Identity *string `json:"Identity,omitempty"`
	// User provided name for the datacenter. Usually, this name is subject to manipulations by user. It is not the identity of the datacenter.
	Name *string `json:"Name,omitempty"`
}

// NewVirtualizationDatacenter instantiates a new VirtualizationDatacenter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationDatacenter() *VirtualizationDatacenter {
	this := VirtualizationDatacenter{}
	return &this
}

// NewVirtualizationDatacenterWithDefaults instantiates a new VirtualizationDatacenter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationDatacenterWithDefaults() *VirtualizationDatacenter {
	this := VirtualizationDatacenter{}
	return &this
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *VirtualizationDatacenter) GetIdentity() string {
	if o == nil || o.Identity == nil {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationDatacenter) GetIdentityOk() (*string, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *VirtualizationDatacenter) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *VirtualizationDatacenter) SetIdentity(v string) {
	o.Identity = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VirtualizationDatacenter) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationDatacenter) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VirtualizationDatacenter) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VirtualizationDatacenter) SetName(v string) {
	o.Name = &v
}

func (o VirtualizationDatacenter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationSourceDevice, errVirtualizationSourceDevice := json.Marshal(o.VirtualizationSourceDevice)
	if errVirtualizationSourceDevice != nil {
		return []byte{}, errVirtualizationSourceDevice
	}
	errVirtualizationSourceDevice = json.Unmarshal([]byte(serializedVirtualizationSourceDevice), &toSerialize)
	if errVirtualizationSourceDevice != nil {
		return []byte{}, errVirtualizationSourceDevice
	}
	if o.Identity != nil {
		toSerialize["Identity"] = o.Identity
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableVirtualizationDatacenter struct {
	value *VirtualizationDatacenter
	isSet bool
}

func (v NullableVirtualizationDatacenter) Get() *VirtualizationDatacenter {
	return v.value
}

func (v *NullableVirtualizationDatacenter) Set(val *VirtualizationDatacenter) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationDatacenter) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationDatacenter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationDatacenter(val *VirtualizationDatacenter) *NullableVirtualizationDatacenter {
	return &NullableVirtualizationDatacenter{value: val, isSet: true}
}

func (v NullableVirtualizationDatacenter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationDatacenter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
