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

// CapabilityCatalogAllOf Definition of the list of properties defined in 'capability.Catalog', excluding properties defined in parent classes.
type CapabilityCatalogAllOf struct {
	// A unique name for the catalog.
	Name         *string                               `json:"Name,omitempty"`
	Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to capabilitySection resources.
	Sections []CapabilitySectionRelationship `json:"Sections,omitempty"`
}

// NewCapabilityCatalogAllOf instantiates a new CapabilityCatalogAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityCatalogAllOf() *CapabilityCatalogAllOf {
	this := CapabilityCatalogAllOf{}
	return &this
}

// NewCapabilityCatalogAllOfWithDefaults instantiates a new CapabilityCatalogAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityCatalogAllOfWithDefaults() *CapabilityCatalogAllOf {
	this := CapabilityCatalogAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CapabilityCatalogAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityCatalogAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CapabilityCatalogAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CapabilityCatalogAllOf) SetName(v string) {
	o.Name = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *CapabilityCatalogAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityCatalogAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *CapabilityCatalogAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *CapabilityCatalogAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetSections returns the Sections field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CapabilityCatalogAllOf) GetSections() []CapabilitySectionRelationship {
	if o == nil {
		var ret []CapabilitySectionRelationship
		return ret
	}
	return o.Sections
}

// GetSectionsOk returns a tuple with the Sections field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CapabilityCatalogAllOf) GetSectionsOk() (*[]CapabilitySectionRelationship, bool) {
	if o == nil || o.Sections == nil {
		return nil, false
	}
	return &o.Sections, true
}

// HasSections returns a boolean if a field has been set.
func (o *CapabilityCatalogAllOf) HasSections() bool {
	if o != nil && o.Sections != nil {
		return true
	}

	return false
}

// SetSections gets a reference to the given []CapabilitySectionRelationship and assigns it to the Sections field.
func (o *CapabilityCatalogAllOf) SetSections(v []CapabilitySectionRelationship) {
	o.Sections = v
}

func (o CapabilityCatalogAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Sections != nil {
		toSerialize["Sections"] = o.Sections
	}
	return json.Marshal(toSerialize)
}

type NullableCapabilityCatalogAllOf struct {
	value *CapabilityCatalogAllOf
	isSet bool
}

func (v NullableCapabilityCatalogAllOf) Get() *CapabilityCatalogAllOf {
	return v.value
}

func (v *NullableCapabilityCatalogAllOf) Set(val *CapabilityCatalogAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityCatalogAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityCatalogAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityCatalogAllOf(val *CapabilityCatalogAllOf) *NullableCapabilityCatalogAllOf {
	return &NullableCapabilityCatalogAllOf{value: val, isSet: true}
}

func (v NullableCapabilityCatalogAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityCatalogAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
