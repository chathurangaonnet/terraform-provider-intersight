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

// AssetProductInformation Type for saving the product information.
type AssetProductInformation struct {
	MoBaseComplexType
	BillTo *AssetAddressInformation `json:"BillTo,omitempty"`
	// Short description of the Cisco product that helps identify the product easily. example \"DISTI:UCS 6248UP 1RU Fabric Int/No PSU/32 UP/ 12p LIC\".
	Description *string `json:"Description,omitempty"`
	// Family that the product belongs to. Example \"UCSB\".
	Family *string `json:"Family,omitempty"`
	// Group that the product belongs to. It is one higher level categorization above family. example \"Switch\".
	Group *string `json:"Group,omitempty"`
	// Product number that identifies the product. example PID. example \"UCS-FI-6248UP-CH2\".
	Number *string                  `json:"Number,omitempty"`
	ShipTo *AssetAddressInformation `json:"ShipTo,omitempty"`
	// Sub type of the product being specified. example \"UCS 6200 SER\".
	SubType *string `json:"SubType,omitempty"`
}

// NewAssetProductInformation instantiates a new AssetProductInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetProductInformation() *AssetProductInformation {
	this := AssetProductInformation{}
	return &this
}

// NewAssetProductInformationWithDefaults instantiates a new AssetProductInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetProductInformationWithDefaults() *AssetProductInformation {
	this := AssetProductInformation{}
	return &this
}

// GetBillTo returns the BillTo field value if set, zero value otherwise.
func (o *AssetProductInformation) GetBillTo() AssetAddressInformation {
	if o == nil || o.BillTo == nil {
		var ret AssetAddressInformation
		return ret
	}
	return *o.BillTo
}

// GetBillToOk returns a tuple with the BillTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetProductInformation) GetBillToOk() (*AssetAddressInformation, bool) {
	if o == nil || o.BillTo == nil {
		return nil, false
	}
	return o.BillTo, true
}

// HasBillTo returns a boolean if a field has been set.
func (o *AssetProductInformation) HasBillTo() bool {
	if o != nil && o.BillTo != nil {
		return true
	}

	return false
}

// SetBillTo gets a reference to the given AssetAddressInformation and assigns it to the BillTo field.
func (o *AssetProductInformation) SetBillTo(v AssetAddressInformation) {
	o.BillTo = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AssetProductInformation) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetProductInformation) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AssetProductInformation) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AssetProductInformation) SetDescription(v string) {
	o.Description = &v
}

// GetFamily returns the Family field value if set, zero value otherwise.
func (o *AssetProductInformation) GetFamily() string {
	if o == nil || o.Family == nil {
		var ret string
		return ret
	}
	return *o.Family
}

// GetFamilyOk returns a tuple with the Family field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetProductInformation) GetFamilyOk() (*string, bool) {
	if o == nil || o.Family == nil {
		return nil, false
	}
	return o.Family, true
}

// HasFamily returns a boolean if a field has been set.
func (o *AssetProductInformation) HasFamily() bool {
	if o != nil && o.Family != nil {
		return true
	}

	return false
}

// SetFamily gets a reference to the given string and assigns it to the Family field.
func (o *AssetProductInformation) SetFamily(v string) {
	o.Family = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *AssetProductInformation) GetGroup() string {
	if o == nil || o.Group == nil {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetProductInformation) GetGroupOk() (*string, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *AssetProductInformation) HasGroup() bool {
	if o != nil && o.Group != nil {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *AssetProductInformation) SetGroup(v string) {
	o.Group = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *AssetProductInformation) GetNumber() string {
	if o == nil || o.Number == nil {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetProductInformation) GetNumberOk() (*string, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *AssetProductInformation) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *AssetProductInformation) SetNumber(v string) {
	o.Number = &v
}

// GetShipTo returns the ShipTo field value if set, zero value otherwise.
func (o *AssetProductInformation) GetShipTo() AssetAddressInformation {
	if o == nil || o.ShipTo == nil {
		var ret AssetAddressInformation
		return ret
	}
	return *o.ShipTo
}

// GetShipToOk returns a tuple with the ShipTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetProductInformation) GetShipToOk() (*AssetAddressInformation, bool) {
	if o == nil || o.ShipTo == nil {
		return nil, false
	}
	return o.ShipTo, true
}

// HasShipTo returns a boolean if a field has been set.
func (o *AssetProductInformation) HasShipTo() bool {
	if o != nil && o.ShipTo != nil {
		return true
	}

	return false
}

// SetShipTo gets a reference to the given AssetAddressInformation and assigns it to the ShipTo field.
func (o *AssetProductInformation) SetShipTo(v AssetAddressInformation) {
	o.ShipTo = &v
}

// GetSubType returns the SubType field value if set, zero value otherwise.
func (o *AssetProductInformation) GetSubType() string {
	if o == nil || o.SubType == nil {
		var ret string
		return ret
	}
	return *o.SubType
}

// GetSubTypeOk returns a tuple with the SubType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetProductInformation) GetSubTypeOk() (*string, bool) {
	if o == nil || o.SubType == nil {
		return nil, false
	}
	return o.SubType, true
}

// HasSubType returns a boolean if a field has been set.
func (o *AssetProductInformation) HasSubType() bool {
	if o != nil && o.SubType != nil {
		return true
	}

	return false
}

// SetSubType gets a reference to the given string and assigns it to the SubType field.
func (o *AssetProductInformation) SetSubType(v string) {
	o.SubType = &v
}

func (o AssetProductInformation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.BillTo != nil {
		toSerialize["BillTo"] = o.BillTo
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Family != nil {
		toSerialize["Family"] = o.Family
	}
	if o.Group != nil {
		toSerialize["Group"] = o.Group
	}
	if o.Number != nil {
		toSerialize["Number"] = o.Number
	}
	if o.ShipTo != nil {
		toSerialize["ShipTo"] = o.ShipTo
	}
	if o.SubType != nil {
		toSerialize["SubType"] = o.SubType
	}
	return json.Marshal(toSerialize)
}

type NullableAssetProductInformation struct {
	value *AssetProductInformation
	isSet bool
}

func (v NullableAssetProductInformation) Get() *AssetProductInformation {
	return v.value
}

func (v *NullableAssetProductInformation) Set(val *AssetProductInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetProductInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetProductInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetProductInformation(val *AssetProductInformation) *NullableAssetProductInformation {
	return &NullableAssetProductInformation{value: val, isSet: true}
}

func (v NullableAssetProductInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetProductInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
