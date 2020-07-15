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

// AssetService A service that is enabled on a managed target. For example, Intersight Orchestration or Intersight Workload Optimizer.
type AssetService struct {
	MoBaseComplexType
	Options *AssetServiceOptions `json:"Options,omitempty"`
	// Status indicates if the respective Service can establish a connection and authenticate with the managed target. Status does not include information about the functional health of the target.
	Status *string `json:"Status,omitempty"`
	// When 'Status' is not Connected, statusErrorReason provides further details about why the device is not connected with Intersight.
	StatusErrorReason    *string `json:"StatusErrorReason,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetService AssetService

// NewAssetService instantiates a new AssetService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetService() *AssetService {
	this := AssetService{}
	var status string = ""
	this.Status = &status
	return &this
}

// NewAssetServiceWithDefaults instantiates a new AssetService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetServiceWithDefaults() *AssetService {
	this := AssetService{}
	var status string = ""
	this.Status = &status
	return &this
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *AssetService) GetOptions() AssetServiceOptions {
	if o == nil || o.Options == nil {
		var ret AssetServiceOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetService) GetOptionsOk() (*AssetServiceOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *AssetService) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given AssetServiceOptions and assigns it to the Options field.
func (o *AssetService) SetOptions(v AssetServiceOptions) {
	o.Options = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AssetService) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetService) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AssetService) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AssetService) SetStatus(v string) {
	o.Status = &v
}

// GetStatusErrorReason returns the StatusErrorReason field value if set, zero value otherwise.
func (o *AssetService) GetStatusErrorReason() string {
	if o == nil || o.StatusErrorReason == nil {
		var ret string
		return ret
	}
	return *o.StatusErrorReason
}

// GetStatusErrorReasonOk returns a tuple with the StatusErrorReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetService) GetStatusErrorReasonOk() (*string, bool) {
	if o == nil || o.StatusErrorReason == nil {
		return nil, false
	}
	return o.StatusErrorReason, true
}

// HasStatusErrorReason returns a boolean if a field has been set.
func (o *AssetService) HasStatusErrorReason() bool {
	if o != nil && o.StatusErrorReason != nil {
		return true
	}

	return false
}

// SetStatusErrorReason gets a reference to the given string and assigns it to the StatusErrorReason field.
func (o *AssetService) SetStatusErrorReason(v string) {
	o.StatusErrorReason = &v
}

func (o AssetService) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.Options != nil {
		toSerialize["Options"] = o.Options
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.StatusErrorReason != nil {
		toSerialize["StatusErrorReason"] = o.StatusErrorReason
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetService) UnmarshalJSON(bytes []byte) (err error) {
	type AssetServiceWithoutEmbeddedStruct struct {
		Options *AssetServiceOptions `json:"Options,omitempty"`
		// Status indicates if the respective Service can establish a connection and authenticate with the managed target. Status does not include information about the functional health of the target.
		Status *string `json:"Status,omitempty"`
		// When 'Status' is not Connected, statusErrorReason provides further details about why the device is not connected with Intersight.
		StatusErrorReason *string `json:"StatusErrorReason,omitempty"`
	}

	varAssetServiceWithoutEmbeddedStruct := AssetServiceWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAssetServiceWithoutEmbeddedStruct)
	if err == nil {
		varAssetService := _AssetService{}
		varAssetService.Options = varAssetServiceWithoutEmbeddedStruct.Options
		varAssetService.Status = varAssetServiceWithoutEmbeddedStruct.Status
		varAssetService.StatusErrorReason = varAssetServiceWithoutEmbeddedStruct.StatusErrorReason
		*o = AssetService(varAssetService)
	} else {
		return err
	}

	varAssetService := _AssetService{}

	err = json.Unmarshal(bytes, &varAssetService)
	if err == nil {
		o.MoBaseComplexType = varAssetService.MoBaseComplexType
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Options")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "StatusErrorReason")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableAssetService struct {
	value *AssetService
	isSet bool
}

func (v NullableAssetService) Get() *AssetService {
	return v.value
}

func (v *NullableAssetService) Set(val *AssetService) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetService) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetService(val *AssetService) *NullableAssetService {
	return &NullableAssetService{value: val, isSet: true}
}

func (v NullableAssetService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
