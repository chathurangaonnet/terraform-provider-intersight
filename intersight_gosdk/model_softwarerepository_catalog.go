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

// SoftwarerepositoryCatalog A container MO that holds references to the files in an account's image repository. It is internally created for each account and is used to hold information about all user uploaded files.
type SoftwarerepositoryCatalog struct {
	MoBaseMo
	// The name of the catalog. The names are populated and predefined during MO creation.
	Name         *string                               `json:"Name,omitempty"`
	Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	System       *IamSystemRelationship                `json:"System,omitempty"`
}

// NewSoftwarerepositoryCatalog instantiates a new SoftwarerepositoryCatalog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwarerepositoryCatalog() *SoftwarerepositoryCatalog {
	this := SoftwarerepositoryCatalog{}
	return &this
}

// NewSoftwarerepositoryCatalogWithDefaults instantiates a new SoftwarerepositoryCatalog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwarerepositoryCatalogWithDefaults() *SoftwarerepositoryCatalog {
	this := SoftwarerepositoryCatalog{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SoftwarerepositoryCatalog) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCatalog) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SoftwarerepositoryCatalog) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SoftwarerepositoryCatalog) SetName(v string) {
	o.Name = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *SoftwarerepositoryCatalog) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCatalog) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *SoftwarerepositoryCatalog) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *SoftwarerepositoryCatalog) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *SoftwarerepositoryCatalog) GetSystem() IamSystemRelationship {
	if o == nil || o.System == nil {
		var ret IamSystemRelationship
		return ret
	}
	return *o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCatalog) GetSystemOk() (*IamSystemRelationship, bool) {
	if o == nil || o.System == nil {
		return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *SoftwarerepositoryCatalog) HasSystem() bool {
	if o != nil && o.System != nil {
		return true
	}

	return false
}

// SetSystem gets a reference to the given IamSystemRelationship and assigns it to the System field.
func (o *SoftwarerepositoryCatalog) SetSystem(v IamSystemRelationship) {
	o.System = &v
}

func (o SoftwarerepositoryCatalog) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.System != nil {
		toSerialize["System"] = o.System
	}
	return json.Marshal(toSerialize)
}

// AsSoftwarerepositoryCatalogRelationship wraps this instance of SoftwarerepositoryCatalog in SoftwarerepositoryCatalogRelationship
func (s *SoftwarerepositoryCatalog) AsSoftwarerepositoryCatalogRelationship() SoftwarerepositoryCatalogRelationship {
	return SoftwarerepositoryCatalogRelationship{SoftwarerepositoryCatalogRelationshipInterface: s}
}

type NullableSoftwarerepositoryCatalog struct {
	value *SoftwarerepositoryCatalog
	isSet bool
}

func (v NullableSoftwarerepositoryCatalog) Get() *SoftwarerepositoryCatalog {
	return v.value
}

func (v *NullableSoftwarerepositoryCatalog) Set(val *SoftwarerepositoryCatalog) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwarerepositoryCatalog) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwarerepositoryCatalog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwarerepositoryCatalog(val *SoftwarerepositoryCatalog) *NullableSoftwarerepositoryCatalog {
	return &NullableSoftwarerepositoryCatalog{value: val, isSet: true}
}

func (v NullableSoftwarerepositoryCatalog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwarerepositoryCatalog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
