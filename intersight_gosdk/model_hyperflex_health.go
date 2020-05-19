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

// HyperflexHealth struct for HyperflexHealth
type HyperflexHealth struct {
	MoBaseMo
	ArbitrationServiceState   *string                            `json:"ArbitrationServiceState,omitempty"`
	DataReplicationCompliance *string                            `json:"DataReplicationCompliance,omitempty"`
	ResiliencyDetails         *HyperflexHxResiliencyInfoDt       `json:"ResiliencyDetails,omitempty"`
	State                     *string                            `json:"State,omitempty"`
	Uuid                      *string                            `json:"Uuid,omitempty"`
	ZkHealth                  *string                            `json:"ZkHealth,omitempty"`
	ZoneResiliencyList        *[]HyperflexHxZoneResiliencyInfoDt `json:"ZoneResiliencyList,omitempty"`
	Cluster                   *HyperflexClusterRelationship      `json:"Cluster,omitempty"`
}

// NewHyperflexHealth instantiates a new HyperflexHealth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexHealth() *HyperflexHealth {
	this := HyperflexHealth{}
	var arbitrationServiceState string = "NOT_AVAILABLE"
	this.ArbitrationServiceState = &arbitrationServiceState
	var dataReplicationCompliance string = "UNKNOWN"
	this.DataReplicationCompliance = &dataReplicationCompliance
	var state string = "UNKNOWN"
	this.State = &state
	var zkHealth string = "NOT_AVAILABLE"
	this.ZkHealth = &zkHealth
	return &this
}

// NewHyperflexHealthWithDefaults instantiates a new HyperflexHealth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexHealthWithDefaults() *HyperflexHealth {
	this := HyperflexHealth{}
	var arbitrationServiceState string = "NOT_AVAILABLE"
	this.ArbitrationServiceState = &arbitrationServiceState
	var dataReplicationCompliance string = "UNKNOWN"
	this.DataReplicationCompliance = &dataReplicationCompliance
	var state string = "UNKNOWN"
	this.State = &state
	var zkHealth string = "NOT_AVAILABLE"
	this.ZkHealth = &zkHealth
	return &this
}

// GetArbitrationServiceState returns the ArbitrationServiceState field value if set, zero value otherwise.
func (o *HyperflexHealth) GetArbitrationServiceState() string {
	if o == nil || o.ArbitrationServiceState == nil {
		var ret string
		return ret
	}
	return *o.ArbitrationServiceState
}

// GetArbitrationServiceStateOk returns a tuple with the ArbitrationServiceState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealth) GetArbitrationServiceStateOk() (*string, bool) {
	if o == nil || o.ArbitrationServiceState == nil {
		return nil, false
	}
	return o.ArbitrationServiceState, true
}

// HasArbitrationServiceState returns a boolean if a field has been set.
func (o *HyperflexHealth) HasArbitrationServiceState() bool {
	if o != nil && o.ArbitrationServiceState != nil {
		return true
	}

	return false
}

// SetArbitrationServiceState gets a reference to the given string and assigns it to the ArbitrationServiceState field.
func (o *HyperflexHealth) SetArbitrationServiceState(v string) {
	o.ArbitrationServiceState = &v
}

// GetDataReplicationCompliance returns the DataReplicationCompliance field value if set, zero value otherwise.
func (o *HyperflexHealth) GetDataReplicationCompliance() string {
	if o == nil || o.DataReplicationCompliance == nil {
		var ret string
		return ret
	}
	return *o.DataReplicationCompliance
}

// GetDataReplicationComplianceOk returns a tuple with the DataReplicationCompliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealth) GetDataReplicationComplianceOk() (*string, bool) {
	if o == nil || o.DataReplicationCompliance == nil {
		return nil, false
	}
	return o.DataReplicationCompliance, true
}

// HasDataReplicationCompliance returns a boolean if a field has been set.
func (o *HyperflexHealth) HasDataReplicationCompliance() bool {
	if o != nil && o.DataReplicationCompliance != nil {
		return true
	}

	return false
}

// SetDataReplicationCompliance gets a reference to the given string and assigns it to the DataReplicationCompliance field.
func (o *HyperflexHealth) SetDataReplicationCompliance(v string) {
	o.DataReplicationCompliance = &v
}

// GetResiliencyDetails returns the ResiliencyDetails field value if set, zero value otherwise.
func (o *HyperflexHealth) GetResiliencyDetails() HyperflexHxResiliencyInfoDt {
	if o == nil || o.ResiliencyDetails == nil {
		var ret HyperflexHxResiliencyInfoDt
		return ret
	}
	return *o.ResiliencyDetails
}

// GetResiliencyDetailsOk returns a tuple with the ResiliencyDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealth) GetResiliencyDetailsOk() (*HyperflexHxResiliencyInfoDt, bool) {
	if o == nil || o.ResiliencyDetails == nil {
		return nil, false
	}
	return o.ResiliencyDetails, true
}

// HasResiliencyDetails returns a boolean if a field has been set.
func (o *HyperflexHealth) HasResiliencyDetails() bool {
	if o != nil && o.ResiliencyDetails != nil {
		return true
	}

	return false
}

// SetResiliencyDetails gets a reference to the given HyperflexHxResiliencyInfoDt and assigns it to the ResiliencyDetails field.
func (o *HyperflexHealth) SetResiliencyDetails(v HyperflexHxResiliencyInfoDt) {
	o.ResiliencyDetails = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *HyperflexHealth) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealth) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *HyperflexHealth) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *HyperflexHealth) SetState(v string) {
	o.State = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *HyperflexHealth) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealth) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *HyperflexHealth) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *HyperflexHealth) SetUuid(v string) {
	o.Uuid = &v
}

// GetZkHealth returns the ZkHealth field value if set, zero value otherwise.
func (o *HyperflexHealth) GetZkHealth() string {
	if o == nil || o.ZkHealth == nil {
		var ret string
		return ret
	}
	return *o.ZkHealth
}

// GetZkHealthOk returns a tuple with the ZkHealth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealth) GetZkHealthOk() (*string, bool) {
	if o == nil || o.ZkHealth == nil {
		return nil, false
	}
	return o.ZkHealth, true
}

// HasZkHealth returns a boolean if a field has been set.
func (o *HyperflexHealth) HasZkHealth() bool {
	if o != nil && o.ZkHealth != nil {
		return true
	}

	return false
}

// SetZkHealth gets a reference to the given string and assigns it to the ZkHealth field.
func (o *HyperflexHealth) SetZkHealth(v string) {
	o.ZkHealth = &v
}

// GetZoneResiliencyList returns the ZoneResiliencyList field value if set, zero value otherwise.
func (o *HyperflexHealth) GetZoneResiliencyList() []HyperflexHxZoneResiliencyInfoDt {
	if o == nil || o.ZoneResiliencyList == nil {
		var ret []HyperflexHxZoneResiliencyInfoDt
		return ret
	}
	return *o.ZoneResiliencyList
}

// GetZoneResiliencyListOk returns a tuple with the ZoneResiliencyList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealth) GetZoneResiliencyListOk() (*[]HyperflexHxZoneResiliencyInfoDt, bool) {
	if o == nil || o.ZoneResiliencyList == nil {
		return nil, false
	}
	return o.ZoneResiliencyList, true
}

// HasZoneResiliencyList returns a boolean if a field has been set.
func (o *HyperflexHealth) HasZoneResiliencyList() bool {
	if o != nil && o.ZoneResiliencyList != nil {
		return true
	}

	return false
}

// SetZoneResiliencyList gets a reference to the given []HyperflexHxZoneResiliencyInfoDt and assigns it to the ZoneResiliencyList field.
func (o *HyperflexHealth) SetZoneResiliencyList(v []HyperflexHxZoneResiliencyInfoDt) {
	o.ZoneResiliencyList = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *HyperflexHealth) GetCluster() HyperflexClusterRelationship {
	if o == nil || o.Cluster == nil {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealth) GetClusterOk() (*HyperflexClusterRelationship, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *HyperflexHealth) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given HyperflexClusterRelationship and assigns it to the Cluster field.
func (o *HyperflexHealth) SetCluster(v HyperflexClusterRelationship) {
	o.Cluster = &v
}

func (o HyperflexHealth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.ArbitrationServiceState != nil {
		toSerialize["ArbitrationServiceState"] = o.ArbitrationServiceState
	}
	if o.DataReplicationCompliance != nil {
		toSerialize["DataReplicationCompliance"] = o.DataReplicationCompliance
	}
	if o.ResiliencyDetails != nil {
		toSerialize["ResiliencyDetails"] = o.ResiliencyDetails
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.ZkHealth != nil {
		toSerialize["ZkHealth"] = o.ZkHealth
	}
	if o.ZoneResiliencyList != nil {
		toSerialize["ZoneResiliencyList"] = o.ZoneResiliencyList
	}
	if o.Cluster != nil {
		toSerialize["Cluster"] = o.Cluster
	}
	return json.Marshal(toSerialize)
}

// AsHyperflexHealthRelationship wraps this instance of HyperflexHealth in HyperflexHealthRelationship
func (s *HyperflexHealth) AsHyperflexHealthRelationship() HyperflexHealthRelationship {
	return HyperflexHealthRelationship{HyperflexHealthRelationshipInterface: s}
}

type NullableHyperflexHealth struct {
	value *HyperflexHealth
	isSet bool
}

func (v NullableHyperflexHealth) Get() *HyperflexHealth {
	return v.value
}

func (v *NullableHyperflexHealth) Set(val *HyperflexHealth) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHealth) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHealth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHealth(val *HyperflexHealth) *NullableHyperflexHealth {
	return &NullableHyperflexHealth{value: val, isSet: true}
}

func (v NullableHyperflexHealth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHealth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
