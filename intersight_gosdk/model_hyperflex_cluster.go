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

// HyperflexCluster A HyperFlex cluster. Contains inventory information concerning the health, software versions, storage, and nodes of the cluster.
type HyperflexCluster struct {
	MoBaseMo
	// The number of days remaining before the cluster's storage utilization reaches the recommended capacity limit of 76%. Default value is math.MaxInt32 to indicate that the capacity runway is \"Unknown\" for a cluster that is not connected or with not sufficient data.
	CapacityRunway *int64 `json:"CapacityRunway,omitempty"`
	// The name of this HyperFlex cluster.
	ClusterName *string `json:"ClusterName,omitempty"`
	// The storage type of this cluster (All Flash or Hybrid).
	ClusterType *int64 `json:"ClusterType,omitempty"`
	// The unique identifier for this HyperFlex cluster.
	ClusterUuid *string `json:"ClusterUuid,omitempty"`
	// The number of compute nodes that belong to this cluster.
	ComputeNodeCount *int64 `json:"ComputeNodeCount,omitempty"`
	// The number of converged nodes that belong to this cluster.
	ConvergedNodeCount *int64 `json:"ConvergedNodeCount,omitempty"`
	// The deployment type of the HyperFlex cluster. The cluster can have one of the following configurations: 1. Datacenter: The HyperFlex cluster consists of UCS Fabric Interconnect-attached nodes on a single site. 2. Stretched Cluster: The HyperFlex cluster consists of UCS Fabric Interconnect-attached nodes distributed across multiple sites. 3. Edge: The HyperFlex cluster consists of 2-4 standalone nodes. If the cluster is running a HyperFlex Data Platform version less than 4.0 or if the deployment type cannot be determined, the deployment type is set as 'NA' (not available).
	DeploymentType *string `json:"DeploymentType,omitempty"`
	// The unique identifier of the device registration that represents this HyperFlex cluster's connection to Intersight.
	DeviceId *string `json:"DeviceId,omitempty"`
	// The number of yellow (warning) and red (critical) alarms stored as an aggregate. The first 16 bits indicate the number of red alarms, and the last 16 bits contain the number of yellow alarms.
	FltAggr *int64 `json:"FltAggr,omitempty"`
	// The HyperFlex Data Platform version of this cluster.
	HxVersion *string `json:"HxVersion,omitempty"`
	// The version and build number of the HyperFlex Data Platform for this cluster. After a cluster upgrade, this version string will be updated on the next inventory cycle to reflect the newly installed version.
	HxdpBuildVersion *string `json:"HxdpBuildVersion,omitempty"`
	// The type of hypervisor running on this cluster.
	HypervisorType *string `json:"HypervisorType,omitempty"`
	// The version of hypervisor running on this cluster.
	HypervisorVersion *string           `json:"HypervisorVersion,omitempty"`
	Summary           *HyperflexSummary `json:"Summary,omitempty"`
	// The storage utilization percentage is computed based on total capacity and current capacity utilization.
	UtilizationPercentage *float32 `json:"UtilizationPercentage,omitempty"`
	// The storage utilization trend percentage represents the trend in percentage computed using the first and last point from historical data.
	UtilizationTrendPercentage *float32 `json:"UtilizationTrendPercentage,omitempty"`
	// The number of virtual machines present on this cluster.
	VmCount *int64 `json:"VmCount,omitempty"`
	// An array of relationships to hyperflexAlarm resources.
	Alarm  *[]HyperflexAlarmRelationship `json:"Alarm,omitempty"`
	Health *HyperflexHealthRelationship  `json:"Health,omitempty"`
	// An array of relationships to hyperflexNode resources.
	Nodes            *[]HyperflexNodeRelationship         `json:"Nodes,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
}

// NewHyperflexCluster instantiates a new HyperflexCluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexCluster() *HyperflexCluster {
	this := HyperflexCluster{}
	var deploymentType string = "NA"
	this.DeploymentType = &deploymentType
	var hypervisorType string = "ESXi"
	this.HypervisorType = &hypervisorType
	return &this
}

// NewHyperflexClusterWithDefaults instantiates a new HyperflexCluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexClusterWithDefaults() *HyperflexCluster {
	this := HyperflexCluster{}
	var deploymentType string = "NA"
	this.DeploymentType = &deploymentType
	var hypervisorType string = "ESXi"
	this.HypervisorType = &hypervisorType
	return &this
}

// GetCapacityRunway returns the CapacityRunway field value if set, zero value otherwise.
func (o *HyperflexCluster) GetCapacityRunway() int64 {
	if o == nil || o.CapacityRunway == nil {
		var ret int64
		return ret
	}
	return *o.CapacityRunway
}

// GetCapacityRunwayOk returns a tuple with the CapacityRunway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexCluster) GetCapacityRunwayOk() (*int64, bool) {
	if o == nil || o.CapacityRunway == nil {
		return nil, false
	}
	return o.CapacityRunway, true
}

// HasCapacityRunway returns a boolean if a field has been set.
func (o *HyperflexCluster) HasCapacityRunway() bool {
	if o != nil && o.CapacityRunway != nil {
		return true
	}

	return false
}

// SetCapacityRunway gets a reference to the given int64 and assigns it to the CapacityRunway field.
func (o *HyperflexCluster) SetCapacityRunway(v int64) {
	o.CapacityRunway = &v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *HyperflexCluster) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexCluster) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *HyperflexCluster) HasClusterName() bool {
	if o != nil && o.ClusterName != nil {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *HyperflexCluster) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetClusterType returns the ClusterType field value if set, zero value otherwise.
func (o *HyperflexCluster) GetClusterType() int64 {
	if o == nil || o.ClusterType == nil {
		var ret int64
		return ret
	}
	return *o.ClusterType
}

// GetClusterTypeOk returns a tuple with the ClusterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexCluster) GetClusterTypeOk() (*int64, bool) {
	if o == nil || o.ClusterType == nil {
		return nil, false
	}
	return o.ClusterType, true
}

// HasClusterType returns a boolean if a field has been set.
func (o *HyperflexCluster) HasClusterType() bool {
	if o != nil && o.ClusterType != nil {
		return true
	}

	return false
}

// SetClusterType gets a reference to the given int64 and assigns it to the ClusterType field.
func (o *HyperflexCluster) SetClusterType(v int64) {
	o.ClusterType = &v
}

// GetClusterUuid returns the ClusterUuid field value if set, zero value otherwise.
func (o *HyperflexCluster) GetClusterUuid() string {
	if o == nil || o.ClusterUuid == nil {
		var ret string
		return ret
	}
	return *o.ClusterUuid
}

// GetClusterUuidOk returns a tuple with the ClusterUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexCluster) GetClusterUuidOk() (*string, bool) {
	if o == nil || o.ClusterUuid == nil {
		return nil, false
	}
	return o.ClusterUuid, true
}

// HasClusterUuid returns a boolean if a field has been set.
func (o *HyperflexCluster) HasClusterUuid() bool {
	if o != nil && o.ClusterUuid != nil {
		return true
	}

	return false
}

// SetClusterUuid gets a reference to the given string and assigns it to the ClusterUuid field.
func (o *HyperflexCluster) SetClusterUuid(v string) {
	o.ClusterUuid = &v
}

// GetComputeNodeCount returns the ComputeNodeCount field value if set, zero value otherwise.
func (o *HyperflexCluster) GetComputeNodeCount() int64 {
	if o == nil || o.ComputeNodeCount == nil {
		var ret int64
		return ret
	}
	return *o.ComputeNodeCount
}

// GetComputeNodeCountOk returns a tuple with the ComputeNodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexCluster) GetComputeNodeCountOk() (*int64, bool) {
	if o == nil || o.ComputeNodeCount == nil {
		return nil, false
	}
	return o.ComputeNodeCount, true
}

// HasComputeNodeCount returns a boolean if a field has been set.
func (o *HyperflexCluster) HasComputeNodeCount() bool {
	if o != nil && o.ComputeNodeCount != nil {
		return true
	}

	return false
}

// SetComputeNodeCount gets a reference to the given int64 and assigns it to the ComputeNodeCount field.
func (o *HyperflexCluster) SetComputeNodeCount(v int64) {
	o.ComputeNodeCount = &v
}

// GetConvergedNodeCount returns the ConvergedNodeCount field value if set, zero value otherwise.
func (o *HyperflexCluster) GetConvergedNodeCount() int64 {
	if o == nil || o.ConvergedNodeCount == nil {
		var ret int64
		return ret
	}
	return *o.ConvergedNodeCount
}

// GetConvergedNodeCountOk returns a tuple with the ConvergedNodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexCluster) GetConvergedNodeCountOk() (*int64, bool) {
	if o == nil || o.ConvergedNodeCount == nil {
		return nil, false
	}
	return o.ConvergedNodeCount, true
}

// HasConvergedNodeCount returns a boolean if a field has been set.
func (o *HyperflexCluster) HasConvergedNodeCount() bool {
	if o != nil && o.ConvergedNodeCount != nil {
		return true
	}

	return false
}

// SetConvergedNodeCount gets a reference to the given int64 and assigns it to the ConvergedNodeCount field.
func (o *HyperflexCluster) SetConvergedNodeCount(v int64) {
	o.ConvergedNodeCount = &v
}

// GetDeploymentType returns the DeploymentType field value if set, zero value otherwise.
func (o *HyperflexCluster) GetDeploymentType() string {
	if o == nil || o.DeploymentType == nil {
		var ret string
		return ret
	}
	return *o.DeploymentType
}

// GetDeploymentTypeOk returns a tuple with the DeploymentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexCluster) GetDeploymentTypeOk() (*string, bool) {
	if o == nil || o.DeploymentType == nil {
		return nil, false
	}
	return o.DeploymentType, true
}

// HasDeploymentType returns a boolean if a field has been set.
func (o *HyperflexCluster) HasDeploymentType() bool {
	if o != nil && o.DeploymentType != nil {
		return true
	}

	return false
}

// SetDeploymentType gets a reference to the given string and assigns it to the DeploymentType field.
func (o *HyperflexCluster) SetDeploymentType(v string) {
	o.DeploymentType = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *HyperflexCluster) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexCluster) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *HyperflexCluster) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *HyperflexCluster) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetFltAggr returns the FltAggr field value if set, zero value otherwise.
func (o *HyperflexCluster) GetFltAggr() int64 {
	if o == nil || o.FltAggr == nil {
		var ret int64
		return ret
	}
	return *o.FltAggr
}

// GetFltAggrOk returns a tuple with the FltAggr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexCluster) GetFltAggrOk() (*int64, bool) {
	if o == nil || o.FltAggr == nil {
		return nil, false
	}
	return o.FltAggr, true
}

// HasFltAggr returns a boolean if a field has been set.
func (o *HyperflexCluster) HasFltAggr() bool {
	if o != nil && o.FltAggr != nil {
		return true
	}

	return false
}

// SetFltAggr gets a reference to the given int64 and assigns it to the FltAggr field.
func (o *HyperflexCluster) SetFltAggr(v int64) {
	o.FltAggr = &v
}

// GetHxVersion returns the HxVersion field value if set, zero value otherwise.
func (o *HyperflexCluster) GetHxVersion() string {
	if o == nil || o.HxVersion == nil {
		var ret string
		return ret
	}
	return *o.HxVersion
}

// GetHxVersionOk returns a tuple with the HxVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexCluster) GetHxVersionOk() (*string, bool) {
	if o == nil || o.HxVersion == nil {
		return nil, false
	}
	return o.HxVersion, true
}

// HasHxVersion returns a boolean if a field has been set.
func (o *HyperflexCluster) HasHxVersion() bool {
	if o != nil && o.HxVersion != nil {
		return true
	}

	return false
}

// SetHxVersion gets a reference to the given string and assigns it to the HxVersion field.
func (o *HyperflexCluster) SetHxVersion(v string) {
	o.HxVersion = &v
}

// GetHxdpBuildVersion returns the HxdpBuildVersion field value if set, zero value otherwise.
func (o *HyperflexCluster) GetHxdpBuildVersion() string {
	if o == nil || o.HxdpBuildVersion == nil {
		var ret string
		return ret
	}
	return *o.HxdpBuildVersion
}

// GetHxdpBuildVersionOk returns a tuple with the HxdpBuildVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexCluster) GetHxdpBuildVersionOk() (*string, bool) {
	if o == nil || o.HxdpBuildVersion == nil {
		return nil, false
	}
	return o.HxdpBuildVersion, true
}

// HasHxdpBuildVersion returns a boolean if a field has been set.
func (o *HyperflexCluster) HasHxdpBuildVersion() bool {
	if o != nil && o.HxdpBuildVersion != nil {
		return true
	}

	return false
}

// SetHxdpBuildVersion gets a reference to the given string and assigns it to the HxdpBuildVersion field.
func (o *HyperflexCluster) SetHxdpBuildVersion(v string) {
	o.HxdpBuildVersion = &v
}

// GetHypervisorType returns the HypervisorType field value if set, zero value otherwise.
func (o *HyperflexCluster) GetHypervisorType() string {
	if o == nil || o.HypervisorType == nil {
		var ret string
		return ret
	}
	return *o.HypervisorType
}

// GetHypervisorTypeOk returns a tuple with the HypervisorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexCluster) GetHypervisorTypeOk() (*string, bool) {
	if o == nil || o.HypervisorType == nil {
		return nil, false
	}
	return o.HypervisorType, true
}

// HasHypervisorType returns a boolean if a field has been set.
func (o *HyperflexCluster) HasHypervisorType() bool {
	if o != nil && o.HypervisorType != nil {
		return true
	}

	return false
}

// SetHypervisorType gets a reference to the given string and assigns it to the HypervisorType field.
func (o *HyperflexCluster) SetHypervisorType(v string) {
	o.HypervisorType = &v
}

// GetHypervisorVersion returns the HypervisorVersion field value if set, zero value otherwise.
func (o *HyperflexCluster) GetHypervisorVersion() string {
	if o == nil || o.HypervisorVersion == nil {
		var ret string
		return ret
	}
	return *o.HypervisorVersion
}

// GetHypervisorVersionOk returns a tuple with the HypervisorVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexCluster) GetHypervisorVersionOk() (*string, bool) {
	if o == nil || o.HypervisorVersion == nil {
		return nil, false
	}
	return o.HypervisorVersion, true
}

// HasHypervisorVersion returns a boolean if a field has been set.
func (o *HyperflexCluster) HasHypervisorVersion() bool {
	if o != nil && o.HypervisorVersion != nil {
		return true
	}

	return false
}

// SetHypervisorVersion gets a reference to the given string and assigns it to the HypervisorVersion field.
func (o *HyperflexCluster) SetHypervisorVersion(v string) {
	o.HypervisorVersion = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *HyperflexCluster) GetSummary() HyperflexSummary {
	if o == nil || o.Summary == nil {
		var ret HyperflexSummary
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexCluster) GetSummaryOk() (*HyperflexSummary, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *HyperflexCluster) HasSummary() bool {
	if o != nil && o.Summary != nil {
		return true
	}

	return false
}

// SetSummary gets a reference to the given HyperflexSummary and assigns it to the Summary field.
func (o *HyperflexCluster) SetSummary(v HyperflexSummary) {
	o.Summary = &v
}

// GetUtilizationPercentage returns the UtilizationPercentage field value if set, zero value otherwise.
func (o *HyperflexCluster) GetUtilizationPercentage() float32 {
	if o == nil || o.UtilizationPercentage == nil {
		var ret float32
		return ret
	}
	return *o.UtilizationPercentage
}

// GetUtilizationPercentageOk returns a tuple with the UtilizationPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexCluster) GetUtilizationPercentageOk() (*float32, bool) {
	if o == nil || o.UtilizationPercentage == nil {
		return nil, false
	}
	return o.UtilizationPercentage, true
}

// HasUtilizationPercentage returns a boolean if a field has been set.
func (o *HyperflexCluster) HasUtilizationPercentage() bool {
	if o != nil && o.UtilizationPercentage != nil {
		return true
	}

	return false
}

// SetUtilizationPercentage gets a reference to the given float32 and assigns it to the UtilizationPercentage field.
func (o *HyperflexCluster) SetUtilizationPercentage(v float32) {
	o.UtilizationPercentage = &v
}

// GetUtilizationTrendPercentage returns the UtilizationTrendPercentage field value if set, zero value otherwise.
func (o *HyperflexCluster) GetUtilizationTrendPercentage() float32 {
	if o == nil || o.UtilizationTrendPercentage == nil {
		var ret float32
		return ret
	}
	return *o.UtilizationTrendPercentage
}

// GetUtilizationTrendPercentageOk returns a tuple with the UtilizationTrendPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexCluster) GetUtilizationTrendPercentageOk() (*float32, bool) {
	if o == nil || o.UtilizationTrendPercentage == nil {
		return nil, false
	}
	return o.UtilizationTrendPercentage, true
}

// HasUtilizationTrendPercentage returns a boolean if a field has been set.
func (o *HyperflexCluster) HasUtilizationTrendPercentage() bool {
	if o != nil && o.UtilizationTrendPercentage != nil {
		return true
	}

	return false
}

// SetUtilizationTrendPercentage gets a reference to the given float32 and assigns it to the UtilizationTrendPercentage field.
func (o *HyperflexCluster) SetUtilizationTrendPercentage(v float32) {
	o.UtilizationTrendPercentage = &v
}

// GetVmCount returns the VmCount field value if set, zero value otherwise.
func (o *HyperflexCluster) GetVmCount() int64 {
	if o == nil || o.VmCount == nil {
		var ret int64
		return ret
	}
	return *o.VmCount
}

// GetVmCountOk returns a tuple with the VmCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexCluster) GetVmCountOk() (*int64, bool) {
	if o == nil || o.VmCount == nil {
		return nil, false
	}
	return o.VmCount, true
}

// HasVmCount returns a boolean if a field has been set.
func (o *HyperflexCluster) HasVmCount() bool {
	if o != nil && o.VmCount != nil {
		return true
	}

	return false
}

// SetVmCount gets a reference to the given int64 and assigns it to the VmCount field.
func (o *HyperflexCluster) SetVmCount(v int64) {
	o.VmCount = &v
}

// GetAlarm returns the Alarm field value if set, zero value otherwise.
func (o *HyperflexCluster) GetAlarm() []HyperflexAlarmRelationship {
	if o == nil || o.Alarm == nil {
		var ret []HyperflexAlarmRelationship
		return ret
	}
	return *o.Alarm
}

// GetAlarmOk returns a tuple with the Alarm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexCluster) GetAlarmOk() (*[]HyperflexAlarmRelationship, bool) {
	if o == nil || o.Alarm == nil {
		return nil, false
	}
	return o.Alarm, true
}

// HasAlarm returns a boolean if a field has been set.
func (o *HyperflexCluster) HasAlarm() bool {
	if o != nil && o.Alarm != nil {
		return true
	}

	return false
}

// SetAlarm gets a reference to the given []HyperflexAlarmRelationship and assigns it to the Alarm field.
func (o *HyperflexCluster) SetAlarm(v []HyperflexAlarmRelationship) {
	o.Alarm = &v
}

// GetHealth returns the Health field value if set, zero value otherwise.
func (o *HyperflexCluster) GetHealth() HyperflexHealthRelationship {
	if o == nil || o.Health == nil {
		var ret HyperflexHealthRelationship
		return ret
	}
	return *o.Health
}

// GetHealthOk returns a tuple with the Health field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexCluster) GetHealthOk() (*HyperflexHealthRelationship, bool) {
	if o == nil || o.Health == nil {
		return nil, false
	}
	return o.Health, true
}

// HasHealth returns a boolean if a field has been set.
func (o *HyperflexCluster) HasHealth() bool {
	if o != nil && o.Health != nil {
		return true
	}

	return false
}

// SetHealth gets a reference to the given HyperflexHealthRelationship and assigns it to the Health field.
func (o *HyperflexCluster) SetHealth(v HyperflexHealthRelationship) {
	o.Health = &v
}

// GetNodes returns the Nodes field value if set, zero value otherwise.
func (o *HyperflexCluster) GetNodes() []HyperflexNodeRelationship {
	if o == nil || o.Nodes == nil {
		var ret []HyperflexNodeRelationship
		return ret
	}
	return *o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexCluster) GetNodesOk() (*[]HyperflexNodeRelationship, bool) {
	if o == nil || o.Nodes == nil {
		return nil, false
	}
	return o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *HyperflexCluster) HasNodes() bool {
	if o != nil && o.Nodes != nil {
		return true
	}

	return false
}

// SetNodes gets a reference to the given []HyperflexNodeRelationship and assigns it to the Nodes field.
func (o *HyperflexCluster) SetNodes(v []HyperflexNodeRelationship) {
	o.Nodes = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *HyperflexCluster) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexCluster) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *HyperflexCluster) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *HyperflexCluster) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o HyperflexCluster) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.CapacityRunway != nil {
		toSerialize["CapacityRunway"] = o.CapacityRunway
	}
	if o.ClusterName != nil {
		toSerialize["ClusterName"] = o.ClusterName
	}
	if o.ClusterType != nil {
		toSerialize["ClusterType"] = o.ClusterType
	}
	if o.ClusterUuid != nil {
		toSerialize["ClusterUuid"] = o.ClusterUuid
	}
	if o.ComputeNodeCount != nil {
		toSerialize["ComputeNodeCount"] = o.ComputeNodeCount
	}
	if o.ConvergedNodeCount != nil {
		toSerialize["ConvergedNodeCount"] = o.ConvergedNodeCount
	}
	if o.DeploymentType != nil {
		toSerialize["DeploymentType"] = o.DeploymentType
	}
	if o.DeviceId != nil {
		toSerialize["DeviceId"] = o.DeviceId
	}
	if o.FltAggr != nil {
		toSerialize["FltAggr"] = o.FltAggr
	}
	if o.HxVersion != nil {
		toSerialize["HxVersion"] = o.HxVersion
	}
	if o.HxdpBuildVersion != nil {
		toSerialize["HxdpBuildVersion"] = o.HxdpBuildVersion
	}
	if o.HypervisorType != nil {
		toSerialize["HypervisorType"] = o.HypervisorType
	}
	if o.HypervisorVersion != nil {
		toSerialize["HypervisorVersion"] = o.HypervisorVersion
	}
	if o.Summary != nil {
		toSerialize["Summary"] = o.Summary
	}
	if o.UtilizationPercentage != nil {
		toSerialize["UtilizationPercentage"] = o.UtilizationPercentage
	}
	if o.UtilizationTrendPercentage != nil {
		toSerialize["UtilizationTrendPercentage"] = o.UtilizationTrendPercentage
	}
	if o.VmCount != nil {
		toSerialize["VmCount"] = o.VmCount
	}
	if o.Alarm != nil {
		toSerialize["Alarm"] = o.Alarm
	}
	if o.Health != nil {
		toSerialize["Health"] = o.Health
	}
	if o.Nodes != nil {
		toSerialize["Nodes"] = o.Nodes
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	return json.Marshal(toSerialize)
}

// AsHyperflexClusterRelationship wraps this instance of HyperflexCluster in HyperflexClusterRelationship
func (s *HyperflexCluster) AsHyperflexClusterRelationship() HyperflexClusterRelationship {
	return HyperflexClusterRelationship{HyperflexClusterRelationshipInterface: s}
}

type NullableHyperflexCluster struct {
	value *HyperflexCluster
	isSet bool
}

func (v NullableHyperflexCluster) Get() *HyperflexCluster {
	return v.value
}

func (v *NullableHyperflexCluster) Set(val *HyperflexCluster) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexCluster(val *HyperflexCluster) *NullableHyperflexCluster {
	return &NullableHyperflexCluster{value: val, isSet: true}
}

func (v NullableHyperflexCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
