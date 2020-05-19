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

// WorkflowTaskDefinitionAllOf Definition of the list of properties defined in 'workflow.TaskDefinition', excluding properties defined in parent classes.
type WorkflowTaskDefinitionAllOf struct {
	// When true this will be the task version that is used when a specific task definition version is not specified. The very first task definition created with a name will be set as the default version, after that user can explicitly set any version of the task definition as the default version.
	DefaultVersion *bool `json:"DefaultVersion,omitempty"`
	// The task definition description to describe what this task will do when executed.
	Description        *string                     `json:"Description,omitempty"`
	InternalProperties *WorkflowInternalProperties `json:"InternalProperties,omitempty"`
	// A user friendly short name to identify the task definition.
	Label *string `json:"Label,omitempty"`
	// License entitlement required to run this task. It is determined by license requirement of features.
	LicenseEntitlement *string `json:"LicenseEntitlement,omitempty"`
	// The name of the task definition. The name should follow this convention <Verb or Action><Category><Vendor><Product><Noun or object> Verb or Action is a required portion of the name and this must be part of the pre-approved verb list. Category is an optional field and this will refer to the broad category of the task referring to the type of resource or endpoint. If there is no specific category then use \"Generic\" if required. Vendor is an optional field and this will refer to the specific vendor this task applies to. If the task is generic and not tied to a vendor, then do not specify anything. Product is an optional field, this will contain the vendor product and model when desired. Noun or object is a required field and  this will contain the noun or object on which the action is being performed. Examples SendEmail  - This is a task in Generic category for sending email. NewStorageVolume - This is a vendor agnostic task under Storage device category for creating a new volume.
	Name       *string             `json:"Name,omitempty"`
	Properties *WorkflowProperties `json:"Properties,omitempty"`
	// If set to true, the task requires access to secure properties and uses an encyption token associated with a workflow moid to encrypt or decrypt the secure properties.
	SecurePropAccess *bool `json:"SecurePropAccess,omitempty"`
	// The version of the task definition so we can support multiple versions of a task definition.
	Version *int64                       `json:"Version,omitempty"`
	Catalog *WorkflowCatalogRelationship `json:"Catalog,omitempty"`
	// An array of relationships to workflowTaskDefinition resources.
	ImplementedTasks *[]WorkflowTaskDefinitionRelationship `json:"ImplementedTasks,omitempty"`
	InterfaceTask    *WorkflowTaskDefinitionRelationship   `json:"InterfaceTask,omitempty"`
}

// NewWorkflowTaskDefinitionAllOf instantiates a new WorkflowTaskDefinitionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowTaskDefinitionAllOf() *WorkflowTaskDefinitionAllOf {
	this := WorkflowTaskDefinitionAllOf{}
	var licenseEntitlement string = "Base"
	this.LicenseEntitlement = &licenseEntitlement
	return &this
}

// NewWorkflowTaskDefinitionAllOfWithDefaults instantiates a new WorkflowTaskDefinitionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowTaskDefinitionAllOfWithDefaults() *WorkflowTaskDefinitionAllOf {
	this := WorkflowTaskDefinitionAllOf{}
	var licenseEntitlement string = "Base"
	this.LicenseEntitlement = &licenseEntitlement
	return &this
}

// GetDefaultVersion returns the DefaultVersion field value if set, zero value otherwise.
func (o *WorkflowTaskDefinitionAllOf) GetDefaultVersion() bool {
	if o == nil || o.DefaultVersion == nil {
		var ret bool
		return ret
	}
	return *o.DefaultVersion
}

// GetDefaultVersionOk returns a tuple with the DefaultVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskDefinitionAllOf) GetDefaultVersionOk() (*bool, bool) {
	if o == nil || o.DefaultVersion == nil {
		return nil, false
	}
	return o.DefaultVersion, true
}

// HasDefaultVersion returns a boolean if a field has been set.
func (o *WorkflowTaskDefinitionAllOf) HasDefaultVersion() bool {
	if o != nil && o.DefaultVersion != nil {
		return true
	}

	return false
}

// SetDefaultVersion gets a reference to the given bool and assigns it to the DefaultVersion field.
func (o *WorkflowTaskDefinitionAllOf) SetDefaultVersion(v bool) {
	o.DefaultVersion = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowTaskDefinitionAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskDefinitionAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowTaskDefinitionAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowTaskDefinitionAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetInternalProperties returns the InternalProperties field value if set, zero value otherwise.
func (o *WorkflowTaskDefinitionAllOf) GetInternalProperties() WorkflowInternalProperties {
	if o == nil || o.InternalProperties == nil {
		var ret WorkflowInternalProperties
		return ret
	}
	return *o.InternalProperties
}

// GetInternalPropertiesOk returns a tuple with the InternalProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskDefinitionAllOf) GetInternalPropertiesOk() (*WorkflowInternalProperties, bool) {
	if o == nil || o.InternalProperties == nil {
		return nil, false
	}
	return o.InternalProperties, true
}

// HasInternalProperties returns a boolean if a field has been set.
func (o *WorkflowTaskDefinitionAllOf) HasInternalProperties() bool {
	if o != nil && o.InternalProperties != nil {
		return true
	}

	return false
}

// SetInternalProperties gets a reference to the given WorkflowInternalProperties and assigns it to the InternalProperties field.
func (o *WorkflowTaskDefinitionAllOf) SetInternalProperties(v WorkflowInternalProperties) {
	o.InternalProperties = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WorkflowTaskDefinitionAllOf) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskDefinitionAllOf) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WorkflowTaskDefinitionAllOf) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *WorkflowTaskDefinitionAllOf) SetLabel(v string) {
	o.Label = &v
}

// GetLicenseEntitlement returns the LicenseEntitlement field value if set, zero value otherwise.
func (o *WorkflowTaskDefinitionAllOf) GetLicenseEntitlement() string {
	if o == nil || o.LicenseEntitlement == nil {
		var ret string
		return ret
	}
	return *o.LicenseEntitlement
}

// GetLicenseEntitlementOk returns a tuple with the LicenseEntitlement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskDefinitionAllOf) GetLicenseEntitlementOk() (*string, bool) {
	if o == nil || o.LicenseEntitlement == nil {
		return nil, false
	}
	return o.LicenseEntitlement, true
}

// HasLicenseEntitlement returns a boolean if a field has been set.
func (o *WorkflowTaskDefinitionAllOf) HasLicenseEntitlement() bool {
	if o != nil && o.LicenseEntitlement != nil {
		return true
	}

	return false
}

// SetLicenseEntitlement gets a reference to the given string and assigns it to the LicenseEntitlement field.
func (o *WorkflowTaskDefinitionAllOf) SetLicenseEntitlement(v string) {
	o.LicenseEntitlement = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowTaskDefinitionAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskDefinitionAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowTaskDefinitionAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowTaskDefinitionAllOf) SetName(v string) {
	o.Name = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *WorkflowTaskDefinitionAllOf) GetProperties() WorkflowProperties {
	if o == nil || o.Properties == nil {
		var ret WorkflowProperties
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskDefinitionAllOf) GetPropertiesOk() (*WorkflowProperties, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *WorkflowTaskDefinitionAllOf) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given WorkflowProperties and assigns it to the Properties field.
func (o *WorkflowTaskDefinitionAllOf) SetProperties(v WorkflowProperties) {
	o.Properties = &v
}

// GetSecurePropAccess returns the SecurePropAccess field value if set, zero value otherwise.
func (o *WorkflowTaskDefinitionAllOf) GetSecurePropAccess() bool {
	if o == nil || o.SecurePropAccess == nil {
		var ret bool
		return ret
	}
	return *o.SecurePropAccess
}

// GetSecurePropAccessOk returns a tuple with the SecurePropAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskDefinitionAllOf) GetSecurePropAccessOk() (*bool, bool) {
	if o == nil || o.SecurePropAccess == nil {
		return nil, false
	}
	return o.SecurePropAccess, true
}

// HasSecurePropAccess returns a boolean if a field has been set.
func (o *WorkflowTaskDefinitionAllOf) HasSecurePropAccess() bool {
	if o != nil && o.SecurePropAccess != nil {
		return true
	}

	return false
}

// SetSecurePropAccess gets a reference to the given bool and assigns it to the SecurePropAccess field.
func (o *WorkflowTaskDefinitionAllOf) SetSecurePropAccess(v bool) {
	o.SecurePropAccess = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *WorkflowTaskDefinitionAllOf) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskDefinitionAllOf) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *WorkflowTaskDefinitionAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *WorkflowTaskDefinitionAllOf) SetVersion(v int64) {
	o.Version = &v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise.
func (o *WorkflowTaskDefinitionAllOf) GetCatalog() WorkflowCatalogRelationship {
	if o == nil || o.Catalog == nil {
		var ret WorkflowCatalogRelationship
		return ret
	}
	return *o.Catalog
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskDefinitionAllOf) GetCatalogOk() (*WorkflowCatalogRelationship, bool) {
	if o == nil || o.Catalog == nil {
		return nil, false
	}
	return o.Catalog, true
}

// HasCatalog returns a boolean if a field has been set.
func (o *WorkflowTaskDefinitionAllOf) HasCatalog() bool {
	if o != nil && o.Catalog != nil {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given WorkflowCatalogRelationship and assigns it to the Catalog field.
func (o *WorkflowTaskDefinitionAllOf) SetCatalog(v WorkflowCatalogRelationship) {
	o.Catalog = &v
}

// GetImplementedTasks returns the ImplementedTasks field value if set, zero value otherwise.
func (o *WorkflowTaskDefinitionAllOf) GetImplementedTasks() []WorkflowTaskDefinitionRelationship {
	if o == nil || o.ImplementedTasks == nil {
		var ret []WorkflowTaskDefinitionRelationship
		return ret
	}
	return *o.ImplementedTasks
}

// GetImplementedTasksOk returns a tuple with the ImplementedTasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskDefinitionAllOf) GetImplementedTasksOk() (*[]WorkflowTaskDefinitionRelationship, bool) {
	if o == nil || o.ImplementedTasks == nil {
		return nil, false
	}
	return o.ImplementedTasks, true
}

// HasImplementedTasks returns a boolean if a field has been set.
func (o *WorkflowTaskDefinitionAllOf) HasImplementedTasks() bool {
	if o != nil && o.ImplementedTasks != nil {
		return true
	}

	return false
}

// SetImplementedTasks gets a reference to the given []WorkflowTaskDefinitionRelationship and assigns it to the ImplementedTasks field.
func (o *WorkflowTaskDefinitionAllOf) SetImplementedTasks(v []WorkflowTaskDefinitionRelationship) {
	o.ImplementedTasks = &v
}

// GetInterfaceTask returns the InterfaceTask field value if set, zero value otherwise.
func (o *WorkflowTaskDefinitionAllOf) GetInterfaceTask() WorkflowTaskDefinitionRelationship {
	if o == nil || o.InterfaceTask == nil {
		var ret WorkflowTaskDefinitionRelationship
		return ret
	}
	return *o.InterfaceTask
}

// GetInterfaceTaskOk returns a tuple with the InterfaceTask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskDefinitionAllOf) GetInterfaceTaskOk() (*WorkflowTaskDefinitionRelationship, bool) {
	if o == nil || o.InterfaceTask == nil {
		return nil, false
	}
	return o.InterfaceTask, true
}

// HasInterfaceTask returns a boolean if a field has been set.
func (o *WorkflowTaskDefinitionAllOf) HasInterfaceTask() bool {
	if o != nil && o.InterfaceTask != nil {
		return true
	}

	return false
}

// SetInterfaceTask gets a reference to the given WorkflowTaskDefinitionRelationship and assigns it to the InterfaceTask field.
func (o *WorkflowTaskDefinitionAllOf) SetInterfaceTask(v WorkflowTaskDefinitionRelationship) {
	o.InterfaceTask = &v
}

func (o WorkflowTaskDefinitionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DefaultVersion != nil {
		toSerialize["DefaultVersion"] = o.DefaultVersion
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.InternalProperties != nil {
		toSerialize["InternalProperties"] = o.InternalProperties
	}
	if o.Label != nil {
		toSerialize["Label"] = o.Label
	}
	if o.LicenseEntitlement != nil {
		toSerialize["LicenseEntitlement"] = o.LicenseEntitlement
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Properties != nil {
		toSerialize["Properties"] = o.Properties
	}
	if o.SecurePropAccess != nil {
		toSerialize["SecurePropAccess"] = o.SecurePropAccess
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.Catalog != nil {
		toSerialize["Catalog"] = o.Catalog
	}
	if o.ImplementedTasks != nil {
		toSerialize["ImplementedTasks"] = o.ImplementedTasks
	}
	if o.InterfaceTask != nil {
		toSerialize["InterfaceTask"] = o.InterfaceTask
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowTaskDefinitionAllOf struct {
	value *WorkflowTaskDefinitionAllOf
	isSet bool
}

func (v NullableWorkflowTaskDefinitionAllOf) Get() *WorkflowTaskDefinitionAllOf {
	return v.value
}

func (v *NullableWorkflowTaskDefinitionAllOf) Set(val *WorkflowTaskDefinitionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTaskDefinitionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTaskDefinitionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTaskDefinitionAllOf(val *WorkflowTaskDefinitionAllOf) *NullableWorkflowTaskDefinitionAllOf {
	return &NullableWorkflowTaskDefinitionAllOf{value: val, isSet: true}
}

func (v NullableWorkflowTaskDefinitionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTaskDefinitionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
