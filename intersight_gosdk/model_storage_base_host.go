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

// StorageBaseHost Generic storage host object. It is a logical object to associate initiators to expose a volume as a LUN.
type StorageBaseHost struct {
	MoBaseMo
	// Short description about the host.
	Description *string                 `json:"Description,omitempty"`
	Initiators  *[]StorageBaseInitiator `json:"Initiators,omitempty"`
	// Name of the host in storage array.
	Name *string `json:"Name,omitempty"`
	// Operating system running on the host.
	OsType *string `json:"OsType,omitempty"`
}

// NewStorageBaseHost instantiates a new StorageBaseHost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageBaseHost() *StorageBaseHost {
	this := StorageBaseHost{}
	return &this
}

// NewStorageBaseHostWithDefaults instantiates a new StorageBaseHost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageBaseHostWithDefaults() *StorageBaseHost {
	this := StorageBaseHost{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *StorageBaseHost) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseHost) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *StorageBaseHost) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *StorageBaseHost) SetDescription(v string) {
	o.Description = &v
}

// GetInitiators returns the Initiators field value if set, zero value otherwise.
func (o *StorageBaseHost) GetInitiators() []StorageBaseInitiator {
	if o == nil || o.Initiators == nil {
		var ret []StorageBaseInitiator
		return ret
	}
	return *o.Initiators
}

// GetInitiatorsOk returns a tuple with the Initiators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseHost) GetInitiatorsOk() (*[]StorageBaseInitiator, bool) {
	if o == nil || o.Initiators == nil {
		return nil, false
	}
	return o.Initiators, true
}

// HasInitiators returns a boolean if a field has been set.
func (o *StorageBaseHost) HasInitiators() bool {
	if o != nil && o.Initiators != nil {
		return true
	}

	return false
}

// SetInitiators gets a reference to the given []StorageBaseInitiator and assigns it to the Initiators field.
func (o *StorageBaseHost) SetInitiators(v []StorageBaseInitiator) {
	o.Initiators = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageBaseHost) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseHost) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageBaseHost) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageBaseHost) SetName(v string) {
	o.Name = &v
}

// GetOsType returns the OsType field value if set, zero value otherwise.
func (o *StorageBaseHost) GetOsType() string {
	if o == nil || o.OsType == nil {
		var ret string
		return ret
	}
	return *o.OsType
}

// GetOsTypeOk returns a tuple with the OsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseHost) GetOsTypeOk() (*string, bool) {
	if o == nil || o.OsType == nil {
		return nil, false
	}
	return o.OsType, true
}

// HasOsType returns a boolean if a field has been set.
func (o *StorageBaseHost) HasOsType() bool {
	if o != nil && o.OsType != nil {
		return true
	}

	return false
}

// SetOsType gets a reference to the given string and assigns it to the OsType field.
func (o *StorageBaseHost) SetOsType(v string) {
	o.OsType = &v
}

func (o StorageBaseHost) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Initiators != nil {
		toSerialize["Initiators"] = o.Initiators
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.OsType != nil {
		toSerialize["OsType"] = o.OsType
	}
	return json.Marshal(toSerialize)
}

type NullableStorageBaseHost struct {
	value *StorageBaseHost
	isSet bool
}

func (v NullableStorageBaseHost) Get() *StorageBaseHost {
	return v.value
}

func (v *NullableStorageBaseHost) Set(val *StorageBaseHost) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageBaseHost) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageBaseHost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageBaseHost(val *StorageBaseHost) *NullableStorageBaseHost {
	return &NullableStorageBaseHost{value: val, isSet: true}
}

func (v NullableStorageBaseHost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageBaseHost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
