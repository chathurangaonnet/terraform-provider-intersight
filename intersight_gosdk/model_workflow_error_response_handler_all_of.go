/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-07-11T05:47:33Z.
 *
 * API version: 1.0.9-1999
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
)

// WorkflowErrorResponseHandlerAllOf Definition of the list of properties defined in 'workflow.ErrorResponseHandler', excluding properties defined in parent classes.
type WorkflowErrorResponseHandlerAllOf struct {
	// A detailed description about the error response handler.
	Description *string `json:"Description,omitempty"`
	// Name for the error response handler.
	Name       *string             `json:"Name,omitempty"`
	Parameters *[]ContentParameter `json:"Parameters,omitempty"`
	// The platform type for which the error response handler is defined.
	PlatformType         *string                      `json:"PlatformType,omitempty"`
	Types                *[]ContentComplexType        `json:"Types,omitempty"`
	Catalog              *WorkflowCatalogRelationship `json:"Catalog,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowErrorResponseHandlerAllOf WorkflowErrorResponseHandlerAllOf

// NewWorkflowErrorResponseHandlerAllOf instantiates a new WorkflowErrorResponseHandlerAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowErrorResponseHandlerAllOf() *WorkflowErrorResponseHandlerAllOf {
	this := WorkflowErrorResponseHandlerAllOf{}
	var platformType string = ""
	this.PlatformType = &platformType
	return &this
}

// NewWorkflowErrorResponseHandlerAllOfWithDefaults instantiates a new WorkflowErrorResponseHandlerAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowErrorResponseHandlerAllOfWithDefaults() *WorkflowErrorResponseHandlerAllOf {
	this := WorkflowErrorResponseHandlerAllOf{}
	var platformType string = ""
	this.PlatformType = &platformType
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowErrorResponseHandlerAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowErrorResponseHandlerAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowErrorResponseHandlerAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowErrorResponseHandlerAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowErrorResponseHandlerAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowErrorResponseHandlerAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowErrorResponseHandlerAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowErrorResponseHandlerAllOf) SetName(v string) {
	o.Name = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *WorkflowErrorResponseHandlerAllOf) GetParameters() []ContentParameter {
	if o == nil || o.Parameters == nil {
		var ret []ContentParameter
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowErrorResponseHandlerAllOf) GetParametersOk() (*[]ContentParameter, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *WorkflowErrorResponseHandlerAllOf) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []ContentParameter and assigns it to the Parameters field.
func (o *WorkflowErrorResponseHandlerAllOf) SetParameters(v []ContentParameter) {
	o.Parameters = &v
}

// GetPlatformType returns the PlatformType field value if set, zero value otherwise.
func (o *WorkflowErrorResponseHandlerAllOf) GetPlatformType() string {
	if o == nil || o.PlatformType == nil {
		var ret string
		return ret
	}
	return *o.PlatformType
}

// GetPlatformTypeOk returns a tuple with the PlatformType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowErrorResponseHandlerAllOf) GetPlatformTypeOk() (*string, bool) {
	if o == nil || o.PlatformType == nil {
		return nil, false
	}
	return o.PlatformType, true
}

// HasPlatformType returns a boolean if a field has been set.
func (o *WorkflowErrorResponseHandlerAllOf) HasPlatformType() bool {
	if o != nil && o.PlatformType != nil {
		return true
	}

	return false
}

// SetPlatformType gets a reference to the given string and assigns it to the PlatformType field.
func (o *WorkflowErrorResponseHandlerAllOf) SetPlatformType(v string) {
	o.PlatformType = &v
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *WorkflowErrorResponseHandlerAllOf) GetTypes() []ContentComplexType {
	if o == nil || o.Types == nil {
		var ret []ContentComplexType
		return ret
	}
	return *o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowErrorResponseHandlerAllOf) GetTypesOk() (*[]ContentComplexType, bool) {
	if o == nil || o.Types == nil {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *WorkflowErrorResponseHandlerAllOf) HasTypes() bool {
	if o != nil && o.Types != nil {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []ContentComplexType and assigns it to the Types field.
func (o *WorkflowErrorResponseHandlerAllOf) SetTypes(v []ContentComplexType) {
	o.Types = &v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise.
func (o *WorkflowErrorResponseHandlerAllOf) GetCatalog() WorkflowCatalogRelationship {
	if o == nil || o.Catalog == nil {
		var ret WorkflowCatalogRelationship
		return ret
	}
	return *o.Catalog
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowErrorResponseHandlerAllOf) GetCatalogOk() (*WorkflowCatalogRelationship, bool) {
	if o == nil || o.Catalog == nil {
		return nil, false
	}
	return o.Catalog, true
}

// HasCatalog returns a boolean if a field has been set.
func (o *WorkflowErrorResponseHandlerAllOf) HasCatalog() bool {
	if o != nil && o.Catalog != nil {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given WorkflowCatalogRelationship and assigns it to the Catalog field.
func (o *WorkflowErrorResponseHandlerAllOf) SetCatalog(v WorkflowCatalogRelationship) {
	o.Catalog = &v
}

func (o WorkflowErrorResponseHandlerAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Parameters != nil {
		toSerialize["Parameters"] = o.Parameters
	}
	if o.PlatformType != nil {
		toSerialize["PlatformType"] = o.PlatformType
	}
	if o.Types != nil {
		toSerialize["Types"] = o.Types
	}
	if o.Catalog != nil {
		toSerialize["Catalog"] = o.Catalog
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowErrorResponseHandlerAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowErrorResponseHandlerAllOf := _WorkflowErrorResponseHandlerAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowErrorResponseHandlerAllOf); err == nil {
		*o = WorkflowErrorResponseHandlerAllOf(varWorkflowErrorResponseHandlerAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Parameters")
		delete(additionalProperties, "PlatformType")
		delete(additionalProperties, "Types")
		delete(additionalProperties, "Catalog")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowErrorResponseHandlerAllOf struct {
	value *WorkflowErrorResponseHandlerAllOf
	isSet bool
}

func (v NullableWorkflowErrorResponseHandlerAllOf) Get() *WorkflowErrorResponseHandlerAllOf {
	return v.value
}

func (v *NullableWorkflowErrorResponseHandlerAllOf) Set(val *WorkflowErrorResponseHandlerAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowErrorResponseHandlerAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowErrorResponseHandlerAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowErrorResponseHandlerAllOf(val *WorkflowErrorResponseHandlerAllOf) *NullableWorkflowErrorResponseHandlerAllOf {
	return &NullableWorkflowErrorResponseHandlerAllOf{value: val, isSet: true}
}

func (v NullableWorkflowErrorResponseHandlerAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowErrorResponseHandlerAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
