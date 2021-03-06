/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-06-17T02:04:50-07:00.
 *
 * API version: 1.0.9-1867
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// TamBaseAdvisory An Intersight Advisory. An advisory represents an identification of a potential issue and may also include  a recommendation for resolving the said issue. Advisories may be of different kind and severity. for e.g. It could be a security vulnerability or a performance issue or a hardware issue with different recommendations for resolving them.
type TamBaseAdvisory struct {
	MoBaseMo
	// Brief description of the advisory details.
	Description *string `json:"Description,omitempty"`
	// A user defined name for the Intersight Advisory.
	Name     *string      `json:"Name,omitempty"`
	Severity *TamSeverity `json:"Severity,omitempty"`
	// Current state of the advisory.
	State                *string `json:"State,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TamBaseAdvisory TamBaseAdvisory

// NewTamBaseAdvisory instantiates a new TamBaseAdvisory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTamBaseAdvisory() *TamBaseAdvisory {
	this := TamBaseAdvisory{}
	var state string = "ready"
	this.State = &state
	return &this
}

// NewTamBaseAdvisoryWithDefaults instantiates a new TamBaseAdvisory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTamBaseAdvisoryWithDefaults() *TamBaseAdvisory {
	this := TamBaseAdvisory{}
	var state string = "ready"
	this.State = &state
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TamBaseAdvisory) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamBaseAdvisory) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TamBaseAdvisory) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TamBaseAdvisory) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TamBaseAdvisory) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamBaseAdvisory) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TamBaseAdvisory) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TamBaseAdvisory) SetName(v string) {
	o.Name = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *TamBaseAdvisory) GetSeverity() TamSeverity {
	if o == nil || o.Severity == nil {
		var ret TamSeverity
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamBaseAdvisory) GetSeverityOk() (*TamSeverity, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *TamBaseAdvisory) HasSeverity() bool {
	if o != nil && o.Severity != nil {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given TamSeverity and assigns it to the Severity field.
func (o *TamBaseAdvisory) SetSeverity(v TamSeverity) {
	o.Severity = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *TamBaseAdvisory) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamBaseAdvisory) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *TamBaseAdvisory) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *TamBaseAdvisory) SetState(v string) {
	o.State = &v
}

func (o TamBaseAdvisory) MarshalJSON() ([]byte, error) {
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
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Severity != nil {
		toSerialize["Severity"] = o.Severity
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TamBaseAdvisory) UnmarshalJSON(bytes []byte) (err error) {
	type TamBaseAdvisoryWithoutEmbeddedStruct struct {
		// Brief description of the advisory details.
		Description *string `json:"Description,omitempty"`
		// A user defined name for the Intersight Advisory.
		Name     *string      `json:"Name,omitempty"`
		Severity *TamSeverity `json:"Severity,omitempty"`
		// Current state of the advisory.
		State *string `json:"State,omitempty"`
	}

	varTamBaseAdvisoryWithoutEmbeddedStruct := TamBaseAdvisoryWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varTamBaseAdvisoryWithoutEmbeddedStruct)
	if err == nil {
		varTamBaseAdvisory := _TamBaseAdvisory{}
		varTamBaseAdvisory.Description = varTamBaseAdvisoryWithoutEmbeddedStruct.Description
		varTamBaseAdvisory.Name = varTamBaseAdvisoryWithoutEmbeddedStruct.Name
		varTamBaseAdvisory.Severity = varTamBaseAdvisoryWithoutEmbeddedStruct.Severity
		varTamBaseAdvisory.State = varTamBaseAdvisoryWithoutEmbeddedStruct.State
		*o = TamBaseAdvisory(varTamBaseAdvisory)
	} else {
		return err
	}

	varTamBaseAdvisory := _TamBaseAdvisory{}

	err = json.Unmarshal(bytes, &varTamBaseAdvisory)
	if err == nil {
		o.MoBaseMo = varTamBaseAdvisory.MoBaseMo
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Severity")
		delete(additionalProperties, "State")

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

type NullableTamBaseAdvisory struct {
	value *TamBaseAdvisory
	isSet bool
}

func (v NullableTamBaseAdvisory) Get() *TamBaseAdvisory {
	return v.value
}

func (v *NullableTamBaseAdvisory) Set(val *TamBaseAdvisory) {
	v.value = val
	v.isSet = true
}

func (v NullableTamBaseAdvisory) IsSet() bool {
	return v.isSet
}

func (v *NullableTamBaseAdvisory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTamBaseAdvisory(val *TamBaseAdvisory) *NullableTamBaseAdvisory {
	return &NullableTamBaseAdvisory{value: val, isSet: true}
}

func (v NullableTamBaseAdvisory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTamBaseAdvisory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
