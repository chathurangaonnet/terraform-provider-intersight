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

// AssetTarget Target represents an entity which can be managed by Intersight. This includes physical entities like UCS and HyperFlex servers and software entities like VMware vCenter and Microsoft Azure cloud accounts.
type AssetTarget struct {
	MoBaseMo
	Connections *[]AssetConnection `json:"Connections,omitempty"`
	Services    *[]AssetService    `json:"Services,omitempty"`
	// Status indicates if Intersight can establish a connection and authenticate with the managed target. Status does not include information about the functional health of the target.
	Status *string `json:"Status,omitempty"`
	// StatusErrorReason provides additional context for the Status.
	StatusErrorReason *string `json:"StatusErrorReason,omitempty"`
	// The type of the managed target. For example a UCS Server or Vmware Vcenter target.
	TargetType *string                              `json:"TargetType,omitempty"`
	Account    *IamAccountRelationship              `json:"Account,omitempty"`
	Assist     *AssetDeviceRegistrationRelationship `json:"Assist,omitempty"`
}

// NewAssetTarget instantiates a new AssetTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetTarget() *AssetTarget {
	this := AssetTarget{}
	var status string = ""
	this.Status = &status
	var targetType string = ""
	this.TargetType = &targetType
	return &this
}

// NewAssetTargetWithDefaults instantiates a new AssetTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetTargetWithDefaults() *AssetTarget {
	this := AssetTarget{}
	var status string = ""
	this.Status = &status
	var targetType string = ""
	this.TargetType = &targetType
	return &this
}

// GetConnections returns the Connections field value if set, zero value otherwise.
func (o *AssetTarget) GetConnections() []AssetConnection {
	if o == nil || o.Connections == nil {
		var ret []AssetConnection
		return ret
	}
	return *o.Connections
}

// GetConnectionsOk returns a tuple with the Connections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetTarget) GetConnectionsOk() (*[]AssetConnection, bool) {
	if o == nil || o.Connections == nil {
		return nil, false
	}
	return o.Connections, true
}

// HasConnections returns a boolean if a field has been set.
func (o *AssetTarget) HasConnections() bool {
	if o != nil && o.Connections != nil {
		return true
	}

	return false
}

// SetConnections gets a reference to the given []AssetConnection and assigns it to the Connections field.
func (o *AssetTarget) SetConnections(v []AssetConnection) {
	o.Connections = &v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *AssetTarget) GetServices() []AssetService {
	if o == nil || o.Services == nil {
		var ret []AssetService
		return ret
	}
	return *o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetTarget) GetServicesOk() (*[]AssetService, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *AssetTarget) HasServices() bool {
	if o != nil && o.Services != nil {
		return true
	}

	return false
}

// SetServices gets a reference to the given []AssetService and assigns it to the Services field.
func (o *AssetTarget) SetServices(v []AssetService) {
	o.Services = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AssetTarget) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetTarget) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AssetTarget) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AssetTarget) SetStatus(v string) {
	o.Status = &v
}

// GetStatusErrorReason returns the StatusErrorReason field value if set, zero value otherwise.
func (o *AssetTarget) GetStatusErrorReason() string {
	if o == nil || o.StatusErrorReason == nil {
		var ret string
		return ret
	}
	return *o.StatusErrorReason
}

// GetStatusErrorReasonOk returns a tuple with the StatusErrorReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetTarget) GetStatusErrorReasonOk() (*string, bool) {
	if o == nil || o.StatusErrorReason == nil {
		return nil, false
	}
	return o.StatusErrorReason, true
}

// HasStatusErrorReason returns a boolean if a field has been set.
func (o *AssetTarget) HasStatusErrorReason() bool {
	if o != nil && o.StatusErrorReason != nil {
		return true
	}

	return false
}

// SetStatusErrorReason gets a reference to the given string and assigns it to the StatusErrorReason field.
func (o *AssetTarget) SetStatusErrorReason(v string) {
	o.StatusErrorReason = &v
}

// GetTargetType returns the TargetType field value if set, zero value otherwise.
func (o *AssetTarget) GetTargetType() string {
	if o == nil || o.TargetType == nil {
		var ret string
		return ret
	}
	return *o.TargetType
}

// GetTargetTypeOk returns a tuple with the TargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetTarget) GetTargetTypeOk() (*string, bool) {
	if o == nil || o.TargetType == nil {
		return nil, false
	}
	return o.TargetType, true
}

// HasTargetType returns a boolean if a field has been set.
func (o *AssetTarget) HasTargetType() bool {
	if o != nil && o.TargetType != nil {
		return true
	}

	return false
}

// SetTargetType gets a reference to the given string and assigns it to the TargetType field.
func (o *AssetTarget) SetTargetType(v string) {
	o.TargetType = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *AssetTarget) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetTarget) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *AssetTarget) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *AssetTarget) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

// GetAssist returns the Assist field value if set, zero value otherwise.
func (o *AssetTarget) GetAssist() AssetDeviceRegistrationRelationship {
	if o == nil || o.Assist == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.Assist
}

// GetAssistOk returns a tuple with the Assist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetTarget) GetAssistOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.Assist == nil {
		return nil, false
	}
	return o.Assist, true
}

// HasAssist returns a boolean if a field has been set.
func (o *AssetTarget) HasAssist() bool {
	if o != nil && o.Assist != nil {
		return true
	}

	return false
}

// SetAssist gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the Assist field.
func (o *AssetTarget) SetAssist(v AssetDeviceRegistrationRelationship) {
	o.Assist = &v
}

func (o AssetTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.Connections != nil {
		toSerialize["Connections"] = o.Connections
	}
	if o.Services != nil {
		toSerialize["Services"] = o.Services
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.StatusErrorReason != nil {
		toSerialize["StatusErrorReason"] = o.StatusErrorReason
	}
	if o.TargetType != nil {
		toSerialize["TargetType"] = o.TargetType
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.Assist != nil {
		toSerialize["Assist"] = o.Assist
	}
	return json.Marshal(toSerialize)
}

type NullableAssetTarget struct {
	value *AssetTarget
	isSet bool
}

func (v NullableAssetTarget) Get() *AssetTarget {
	return v.value
}

func (v *NullableAssetTarget) Set(val *AssetTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetTarget(val *AssetTarget) *NullableAssetTarget {
	return &NullableAssetTarget{value: val, isSet: true}
}

func (v NullableAssetTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
