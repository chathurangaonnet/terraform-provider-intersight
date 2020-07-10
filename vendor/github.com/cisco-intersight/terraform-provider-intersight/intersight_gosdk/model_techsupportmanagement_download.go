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

// TechsupportmanagementDownload Download the techsupport file. The response to this API will be the actual techsupport file. This API needs to be invoked using the download link obtained by doing GET operation on TechsupportStatus. The techsupportDownloadUrl property in the TechsupportStatus API response will have the download link. No other link can be used to download the techsupport file.
type TechsupportmanagementDownload struct {
	MoBaseMo
	TechSupportStatus *TechsupportmanagementTechSupportStatusRelationship `json:"TechSupportStatus,omitempty"`
}

// NewTechsupportmanagementDownload instantiates a new TechsupportmanagementDownload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTechsupportmanagementDownload() *TechsupportmanagementDownload {
	this := TechsupportmanagementDownload{}
	return &this
}

// NewTechsupportmanagementDownloadWithDefaults instantiates a new TechsupportmanagementDownload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTechsupportmanagementDownloadWithDefaults() *TechsupportmanagementDownload {
	this := TechsupportmanagementDownload{}
	return &this
}

// GetTechSupportStatus returns the TechSupportStatus field value if set, zero value otherwise.
func (o *TechsupportmanagementDownload) GetTechSupportStatus() TechsupportmanagementTechSupportStatusRelationship {
	if o == nil || o.TechSupportStatus == nil {
		var ret TechsupportmanagementTechSupportStatusRelationship
		return ret
	}
	return *o.TechSupportStatus
}

// GetTechSupportStatusOk returns a tuple with the TechSupportStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementDownload) GetTechSupportStatusOk() (*TechsupportmanagementTechSupportStatusRelationship, bool) {
	if o == nil || o.TechSupportStatus == nil {
		return nil, false
	}
	return o.TechSupportStatus, true
}

// HasTechSupportStatus returns a boolean if a field has been set.
func (o *TechsupportmanagementDownload) HasTechSupportStatus() bool {
	if o != nil && o.TechSupportStatus != nil {
		return true
	}

	return false
}

// SetTechSupportStatus gets a reference to the given TechsupportmanagementTechSupportStatusRelationship and assigns it to the TechSupportStatus field.
func (o *TechsupportmanagementDownload) SetTechSupportStatus(v TechsupportmanagementTechSupportStatusRelationship) {
	o.TechSupportStatus = &v
}

func (o TechsupportmanagementDownload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.TechSupportStatus != nil {
		toSerialize["TechSupportStatus"] = o.TechSupportStatus
	}
	return json.Marshal(toSerialize)
}

type NullableTechsupportmanagementDownload struct {
	value *TechsupportmanagementDownload
	isSet bool
}

func (v NullableTechsupportmanagementDownload) Get() *TechsupportmanagementDownload {
	return v.value
}

func (v *NullableTechsupportmanagementDownload) Set(val *TechsupportmanagementDownload) {
	v.value = val
	v.isSet = true
}

func (v NullableTechsupportmanagementDownload) IsSet() bool {
	return v.isSet
}

func (v *NullableTechsupportmanagementDownload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTechsupportmanagementDownload(val *TechsupportmanagementDownload) *NullableTechsupportmanagementDownload {
	return &NullableTechsupportmanagementDownload{value: val, isSet: true}
}

func (v NullableTechsupportmanagementDownload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTechsupportmanagementDownload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
