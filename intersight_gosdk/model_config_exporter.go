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
	"reflect"
	"strings"
)

// ConfigExporter All export operations are captured as Exporter instances. Users shall use this Exporter mo to track the export operation progress.
type ConfigExporter struct {
	MoBaseMo
	// Pre-signed URL to download the exported package, if the export operation has completed successfully. Regenerated during a GET request, if the existing pre-signed URL has expired.
	DownloadPath *string        `json:"DownloadPath,omitempty"`
	Items        *[]ConfigMoRef `json:"Items,omitempty"`
	// An identifier for the exporter instance.
	Name *string `json:"Name,omitempty"`
	// Status of the export operation.
	Status *string `json:"Status,omitempty"`
	// Status message associated with failures or progress indication.
	StatusMessage *string `json:"StatusMessage,omitempty"`
	// An array of relationships to configExportedItem resources.
	ExportedItems        []ConfigExportedItemRelationship      `json:"ExportedItems,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConfigExporter ConfigExporter

// NewConfigExporter instantiates a new ConfigExporter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigExporter() *ConfigExporter {
	this := ConfigExporter{}
	var status string = ""
	this.Status = &status
	return &this
}

// NewConfigExporterWithDefaults instantiates a new ConfigExporter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigExporterWithDefaults() *ConfigExporter {
	this := ConfigExporter{}
	var status string = ""
	this.Status = &status
	return &this
}

// GetDownloadPath returns the DownloadPath field value if set, zero value otherwise.
func (o *ConfigExporter) GetDownloadPath() string {
	if o == nil || o.DownloadPath == nil {
		var ret string
		return ret
	}
	return *o.DownloadPath
}

// GetDownloadPathOk returns a tuple with the DownloadPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigExporter) GetDownloadPathOk() (*string, bool) {
	if o == nil || o.DownloadPath == nil {
		return nil, false
	}
	return o.DownloadPath, true
}

// HasDownloadPath returns a boolean if a field has been set.
func (o *ConfigExporter) HasDownloadPath() bool {
	if o != nil && o.DownloadPath != nil {
		return true
	}

	return false
}

// SetDownloadPath gets a reference to the given string and assigns it to the DownloadPath field.
func (o *ConfigExporter) SetDownloadPath(v string) {
	o.DownloadPath = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ConfigExporter) GetItems() []ConfigMoRef {
	if o == nil || o.Items == nil {
		var ret []ConfigMoRef
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigExporter) GetItemsOk() (*[]ConfigMoRef, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ConfigExporter) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ConfigMoRef and assigns it to the Items field.
func (o *ConfigExporter) SetItems(v []ConfigMoRef) {
	o.Items = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConfigExporter) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigExporter) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConfigExporter) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConfigExporter) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ConfigExporter) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigExporter) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ConfigExporter) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ConfigExporter) SetStatus(v string) {
	o.Status = &v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise.
func (o *ConfigExporter) GetStatusMessage() string {
	if o == nil || o.StatusMessage == nil {
		var ret string
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigExporter) GetStatusMessageOk() (*string, bool) {
	if o == nil || o.StatusMessage == nil {
		return nil, false
	}
	return o.StatusMessage, true
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *ConfigExporter) HasStatusMessage() bool {
	if o != nil && o.StatusMessage != nil {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given string and assigns it to the StatusMessage field.
func (o *ConfigExporter) SetStatusMessage(v string) {
	o.StatusMessage = &v
}

// GetExportedItems returns the ExportedItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigExporter) GetExportedItems() []ConfigExportedItemRelationship {
	if o == nil {
		var ret []ConfigExportedItemRelationship
		return ret
	}
	return o.ExportedItems
}

// GetExportedItemsOk returns a tuple with the ExportedItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigExporter) GetExportedItemsOk() (*[]ConfigExportedItemRelationship, bool) {
	if o == nil || o.ExportedItems == nil {
		return nil, false
	}
	return &o.ExportedItems, true
}

// HasExportedItems returns a boolean if a field has been set.
func (o *ConfigExporter) HasExportedItems() bool {
	if o != nil && o.ExportedItems != nil {
		return true
	}

	return false
}

// SetExportedItems gets a reference to the given []ConfigExportedItemRelationship and assigns it to the ExportedItems field.
func (o *ConfigExporter) SetExportedItems(v []ConfigExportedItemRelationship) {
	o.ExportedItems = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *ConfigExporter) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigExporter) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *ConfigExporter) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *ConfigExporter) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o ConfigExporter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.DownloadPath != nil {
		toSerialize["DownloadPath"] = o.DownloadPath
	}
	if o.Items != nil {
		toSerialize["Items"] = o.Items
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.StatusMessage != nil {
		toSerialize["StatusMessage"] = o.StatusMessage
	}
	if o.ExportedItems != nil {
		toSerialize["ExportedItems"] = o.ExportedItems
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConfigExporter) UnmarshalJSON(bytes []byte) (err error) {
	type ConfigExporterWithoutEmbeddedStruct struct {
		// Pre-signed URL to download the exported package, if the export operation has completed successfully. Regenerated during a GET request, if the existing pre-signed URL has expired.
		DownloadPath *string        `json:"DownloadPath,omitempty"`
		Items        *[]ConfigMoRef `json:"Items,omitempty"`
		// An identifier for the exporter instance.
		Name *string `json:"Name,omitempty"`
		// Status of the export operation.
		Status *string `json:"Status,omitempty"`
		// Status message associated with failures or progress indication.
		StatusMessage *string `json:"StatusMessage,omitempty"`
		// An array of relationships to configExportedItem resources.
		ExportedItems []ConfigExportedItemRelationship      `json:"ExportedItems,omitempty"`
		Organization  *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varConfigExporterWithoutEmbeddedStruct := ConfigExporterWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varConfigExporterWithoutEmbeddedStruct)
	if err == nil {
		varConfigExporter := _ConfigExporter{}
		varConfigExporter.DownloadPath = varConfigExporterWithoutEmbeddedStruct.DownloadPath
		varConfigExporter.Items = varConfigExporterWithoutEmbeddedStruct.Items
		varConfigExporter.Name = varConfigExporterWithoutEmbeddedStruct.Name
		varConfigExporter.Status = varConfigExporterWithoutEmbeddedStruct.Status
		varConfigExporter.StatusMessage = varConfigExporterWithoutEmbeddedStruct.StatusMessage
		varConfigExporter.ExportedItems = varConfigExporterWithoutEmbeddedStruct.ExportedItems
		varConfigExporter.Organization = varConfigExporterWithoutEmbeddedStruct.Organization
		*o = ConfigExporter(varConfigExporter)
	} else {
		return err
	}

	varConfigExporter := _ConfigExporter{}

	err = json.Unmarshal(bytes, &varConfigExporter)
	if err == nil {
		o.MoBaseMo = varConfigExporter.MoBaseMo
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "DownloadPath")
		delete(additionalProperties, "Items")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "StatusMessage")
		delete(additionalProperties, "ExportedItems")
		delete(additionalProperties, "Organization")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConfigExporter struct {
	value *ConfigExporter
	isSet bool
}

func (v NullableConfigExporter) Get() *ConfigExporter {
	return v.value
}

func (v *NullableConfigExporter) Set(val *ConfigExporter) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigExporter) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigExporter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigExporter(val *ConfigExporter) *NullableConfigExporter {
	return &NullableConfigExporter{value: val, isSet: true}
}

func (v NullableConfigExporter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigExporter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
