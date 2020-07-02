/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-06-19T15:15:17Z.
 *
 * API version: 1.0.9-1903
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
)

// TamAction An action is used to react when an object satifies the condition specified in alert query. For e.g. an action in case of an object matching a fieldNotice query would be to create an alert instance of type 'fieldNotice' for the affected object.
type TamAction struct {
	MoBaseComplexType
	// Type of the managed object that should be marked with an instance of the Alert (when operation type is create) or that should have an alert instance removed (when operation type is remove).
	AffectedObjectType *string `json:"AffectedObjectType,omitempty"`
	// Alert type is used to denote the category of an Intersight alert (FieldNotice, equipment Fault etc.).
	AlertType   *string           `json:"AlertType,omitempty"`
	Identifiers *[]TamIdentifiers `json:"Identifiers,omitempty"`
	// Uniquely identifies a given action among the set of actions corresponding to an advisory. Primarily used to store and compare results of subsequent iterations corresponding to the action queries.
	Name *string `json:"Name,omitempty"`
	// Operation type for the alert action. An action is used to carry out the process of \"reacting\" to an alert condition. For e.g.in case of a fieldNotice alert, the intention may be to create a new alert (if the condition matches and there is no existing alert) or to remove an existing alert when the alert condition has been remedied.
	OperationType *string          `json:"OperationType,omitempty"`
	Queries       *[]TamQueryEntry `json:"Queries,omitempty"`
	// Type of Intersight alert. An alert in Intersight could be one of several kinds (FieldNotice, PSIRT etc.). Primarily used for filtering alerts based on the type.
	Type *string `json:"Type,omitempty"`
}

// NewTamAction instantiates a new TamAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTamAction() *TamAction {
	this := TamAction{}
	var alertType string = "psirt"
	this.AlertType = &alertType
	var operationType string = "create"
	this.OperationType = &operationType
	var type_ string = "restApi"
	this.Type = &type_
	return &this
}

// NewTamActionWithDefaults instantiates a new TamAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTamActionWithDefaults() *TamAction {
	this := TamAction{}
	var alertType string = "psirt"
	this.AlertType = &alertType
	var operationType string = "create"
	this.OperationType = &operationType
	var type_ string = "restApi"
	this.Type = &type_
	return &this
}

// GetAffectedObjectType returns the AffectedObjectType field value if set, zero value otherwise.
func (o *TamAction) GetAffectedObjectType() string {
	if o == nil || o.AffectedObjectType == nil {
		var ret string
		return ret
	}
	return *o.AffectedObjectType
}

// GetAffectedObjectTypeOk returns a tuple with the AffectedObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAction) GetAffectedObjectTypeOk() (*string, bool) {
	if o == nil || o.AffectedObjectType == nil {
		return nil, false
	}
	return o.AffectedObjectType, true
}

// HasAffectedObjectType returns a boolean if a field has been set.
func (o *TamAction) HasAffectedObjectType() bool {
	if o != nil && o.AffectedObjectType != nil {
		return true
	}

	return false
}

// SetAffectedObjectType gets a reference to the given string and assigns it to the AffectedObjectType field.
func (o *TamAction) SetAffectedObjectType(v string) {
	o.AffectedObjectType = &v
}

// GetAlertType returns the AlertType field value if set, zero value otherwise.
func (o *TamAction) GetAlertType() string {
	if o == nil || o.AlertType == nil {
		var ret string
		return ret
	}
	return *o.AlertType
}

// GetAlertTypeOk returns a tuple with the AlertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAction) GetAlertTypeOk() (*string, bool) {
	if o == nil || o.AlertType == nil {
		return nil, false
	}
	return o.AlertType, true
}

// HasAlertType returns a boolean if a field has been set.
func (o *TamAction) HasAlertType() bool {
	if o != nil && o.AlertType != nil {
		return true
	}

	return false
}

// SetAlertType gets a reference to the given string and assigns it to the AlertType field.
func (o *TamAction) SetAlertType(v string) {
	o.AlertType = &v
}

// GetIdentifiers returns the Identifiers field value if set, zero value otherwise.
func (o *TamAction) GetIdentifiers() []TamIdentifiers {
	if o == nil || o.Identifiers == nil {
		var ret []TamIdentifiers
		return ret
	}
	return *o.Identifiers
}

// GetIdentifiersOk returns a tuple with the Identifiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAction) GetIdentifiersOk() (*[]TamIdentifiers, bool) {
	if o == nil || o.Identifiers == nil {
		return nil, false
	}
	return o.Identifiers, true
}

// HasIdentifiers returns a boolean if a field has been set.
func (o *TamAction) HasIdentifiers() bool {
	if o != nil && o.Identifiers != nil {
		return true
	}

	return false
}

// SetIdentifiers gets a reference to the given []TamIdentifiers and assigns it to the Identifiers field.
func (o *TamAction) SetIdentifiers(v []TamIdentifiers) {
	o.Identifiers = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TamAction) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAction) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TamAction) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TamAction) SetName(v string) {
	o.Name = &v
}

// GetOperationType returns the OperationType field value if set, zero value otherwise.
func (o *TamAction) GetOperationType() string {
	if o == nil || o.OperationType == nil {
		var ret string
		return ret
	}
	return *o.OperationType
}

// GetOperationTypeOk returns a tuple with the OperationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAction) GetOperationTypeOk() (*string, bool) {
	if o == nil || o.OperationType == nil {
		return nil, false
	}
	return o.OperationType, true
}

// HasOperationType returns a boolean if a field has been set.
func (o *TamAction) HasOperationType() bool {
	if o != nil && o.OperationType != nil {
		return true
	}

	return false
}

// SetOperationType gets a reference to the given string and assigns it to the OperationType field.
func (o *TamAction) SetOperationType(v string) {
	o.OperationType = &v
}

// GetQueries returns the Queries field value if set, zero value otherwise.
func (o *TamAction) GetQueries() []TamQueryEntry {
	if o == nil || o.Queries == nil {
		var ret []TamQueryEntry
		return ret
	}
	return *o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAction) GetQueriesOk() (*[]TamQueryEntry, bool) {
	if o == nil || o.Queries == nil {
		return nil, false
	}
	return o.Queries, true
}

// HasQueries returns a boolean if a field has been set.
func (o *TamAction) HasQueries() bool {
	if o != nil && o.Queries != nil {
		return true
	}

	return false
}

// SetQueries gets a reference to the given []TamQueryEntry and assigns it to the Queries field.
func (o *TamAction) SetQueries(v []TamQueryEntry) {
	o.Queries = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TamAction) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAction) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TamAction) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TamAction) SetType(v string) {
	o.Type = &v
}

func (o TamAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.AffectedObjectType != nil {
		toSerialize["AffectedObjectType"] = o.AffectedObjectType
	}
	if o.AlertType != nil {
		toSerialize["AlertType"] = o.AlertType
	}
	if o.Identifiers != nil {
		toSerialize["Identifiers"] = o.Identifiers
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.OperationType != nil {
		toSerialize["OperationType"] = o.OperationType
	}
	if o.Queries != nil {
		toSerialize["Queries"] = o.Queries
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableTamAction struct {
	value *TamAction
	isSet bool
}

func (v NullableTamAction) Get() *TamAction {
	return v.value
}

func (v *NullableTamAction) Set(val *TamAction) {
	v.value = val
	v.isSet = true
}

func (v NullableTamAction) IsSet() bool {
	return v.isSet
}

func (v *NullableTamAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTamAction(val *TamAction) *NullableTamAction {
	return &NullableTamAction{value: val, isSet: true}
}

func (v NullableTamAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTamAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}