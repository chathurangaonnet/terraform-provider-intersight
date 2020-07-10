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

// VirtualizationVirtualMachine Depicts operations to control the life cycle of a VM on a Hypervisor.
type VirtualizationVirtualMachine struct {
	MoBaseMo
	// Action to be performed on a virtual machine (Create, PowerState, Migrate, Clone etc).
	Action                *string                        `json:"Action,omitempty"`
	ActionInfo            *VirtualizationActionInfo      `json:"ActionInfo,omitempty"`
	AffinitySelectors     *[]InfraMetaData               `json:"AffinitySelectors,omitempty"`
	AntiAffinitySelectors *[]InfraMetaData               `json:"AntiAffinitySelectors,omitempty"`
	CloudInitConfig       *VirtualizationCloudInitConfig `json:"CloudInitConfig,omitempty"`
	// Number of vCPUs allocated to virtual machine.
	Cpu *int64 `json:"Cpu,omitempty"`
	// Flag to indicate whether the configuration is created from inventory object.
	Discovered *bool                               `json:"Discovered,omitempty"`
	Disk       *[]VirtualizationVirtualMachineDisk `json:"Disk,omitempty"`
	// Guest operating system running on virtual machine.
	GuestOs   *string                           `json:"GuestOs,omitempty"`
	Interface *[]VirtualizationNetworkInterface `json:"Interface,omitempty"`
	Labels    *[]InfraMetaData                  `json:"Labels,omitempty"`
	// Virtual machine memory defined in mega bytes.
	Memory *int64 `json:"Memory,omitempty"`
	// Virtual machine name contains only letters, numbers, allowed special character and must be unique.
	Name *string `json:"Name,omitempty"`
	// Expected power state of virtual machine (PowerOn, PowerOff, Restart).
	PowerState       *string                                       `json:"PowerState,omitempty"`
	Cluster          *VirtualizationBaseClusterRelationship        `json:"Cluster,omitempty"`
	Host             *VirtualizationBaseHostRelationship           `json:"Host,omitempty"`
	Inventory        *VirtualizationBaseVirtualMachineRelationship `json:"Inventory,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship          `json:"RegisteredDevice,omitempty"`
	WorkflowInfo     *WorkflowWorkflowInfoRelationship             `json:"WorkflowInfo,omitempty"`
}

// NewVirtualizationVirtualMachine instantiates a new VirtualizationVirtualMachine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVirtualMachine() *VirtualizationVirtualMachine {
	this := VirtualizationVirtualMachine{}
	var action string = "None"
	this.Action = &action
	var guestOs string = "linux"
	this.GuestOs = &guestOs
	var powerState string = "PowerOff"
	this.PowerState = &powerState
	return &this
}

// NewVirtualizationVirtualMachineWithDefaults instantiates a new VirtualizationVirtualMachine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVirtualMachineWithDefaults() *VirtualizationVirtualMachine {
	this := VirtualizationVirtualMachine{}
	var action string = "None"
	this.Action = &action
	var guestOs string = "linux"
	this.GuestOs = &guestOs
	var powerState string = "PowerOff"
	this.PowerState = &powerState
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *VirtualizationVirtualMachine) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualMachine) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *VirtualizationVirtualMachine) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *VirtualizationVirtualMachine) SetAction(v string) {
	o.Action = &v
}

// GetActionInfo returns the ActionInfo field value if set, zero value otherwise.
func (o *VirtualizationVirtualMachine) GetActionInfo() VirtualizationActionInfo {
	if o == nil || o.ActionInfo == nil {
		var ret VirtualizationActionInfo
		return ret
	}
	return *o.ActionInfo
}

// GetActionInfoOk returns a tuple with the ActionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualMachine) GetActionInfoOk() (*VirtualizationActionInfo, bool) {
	if o == nil || o.ActionInfo == nil {
		return nil, false
	}
	return o.ActionInfo, true
}

// HasActionInfo returns a boolean if a field has been set.
func (o *VirtualizationVirtualMachine) HasActionInfo() bool {
	if o != nil && o.ActionInfo != nil {
		return true
	}

	return false
}

// SetActionInfo gets a reference to the given VirtualizationActionInfo and assigns it to the ActionInfo field.
func (o *VirtualizationVirtualMachine) SetActionInfo(v VirtualizationActionInfo) {
	o.ActionInfo = &v
}

// GetAffinitySelectors returns the AffinitySelectors field value if set, zero value otherwise.
func (o *VirtualizationVirtualMachine) GetAffinitySelectors() []InfraMetaData {
	if o == nil || o.AffinitySelectors == nil {
		var ret []InfraMetaData
		return ret
	}
	return *o.AffinitySelectors
}

// GetAffinitySelectorsOk returns a tuple with the AffinitySelectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualMachine) GetAffinitySelectorsOk() (*[]InfraMetaData, bool) {
	if o == nil || o.AffinitySelectors == nil {
		return nil, false
	}
	return o.AffinitySelectors, true
}

// HasAffinitySelectors returns a boolean if a field has been set.
func (o *VirtualizationVirtualMachine) HasAffinitySelectors() bool {
	if o != nil && o.AffinitySelectors != nil {
		return true
	}

	return false
}

// SetAffinitySelectors gets a reference to the given []InfraMetaData and assigns it to the AffinitySelectors field.
func (o *VirtualizationVirtualMachine) SetAffinitySelectors(v []InfraMetaData) {
	o.AffinitySelectors = &v
}

// GetAntiAffinitySelectors returns the AntiAffinitySelectors field value if set, zero value otherwise.
func (o *VirtualizationVirtualMachine) GetAntiAffinitySelectors() []InfraMetaData {
	if o == nil || o.AntiAffinitySelectors == nil {
		var ret []InfraMetaData
		return ret
	}
	return *o.AntiAffinitySelectors
}

// GetAntiAffinitySelectorsOk returns a tuple with the AntiAffinitySelectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualMachine) GetAntiAffinitySelectorsOk() (*[]InfraMetaData, bool) {
	if o == nil || o.AntiAffinitySelectors == nil {
		return nil, false
	}
	return o.AntiAffinitySelectors, true
}

// HasAntiAffinitySelectors returns a boolean if a field has been set.
func (o *VirtualizationVirtualMachine) HasAntiAffinitySelectors() bool {
	if o != nil && o.AntiAffinitySelectors != nil {
		return true
	}

	return false
}

// SetAntiAffinitySelectors gets a reference to the given []InfraMetaData and assigns it to the AntiAffinitySelectors field.
func (o *VirtualizationVirtualMachine) SetAntiAffinitySelectors(v []InfraMetaData) {
	o.AntiAffinitySelectors = &v
}

// GetCloudInitConfig returns the CloudInitConfig field value if set, zero value otherwise.
func (o *VirtualizationVirtualMachine) GetCloudInitConfig() VirtualizationCloudInitConfig {
	if o == nil || o.CloudInitConfig == nil {
		var ret VirtualizationCloudInitConfig
		return ret
	}
	return *o.CloudInitConfig
}

// GetCloudInitConfigOk returns a tuple with the CloudInitConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualMachine) GetCloudInitConfigOk() (*VirtualizationCloudInitConfig, bool) {
	if o == nil || o.CloudInitConfig == nil {
		return nil, false
	}
	return o.CloudInitConfig, true
}

// HasCloudInitConfig returns a boolean if a field has been set.
func (o *VirtualizationVirtualMachine) HasCloudInitConfig() bool {
	if o != nil && o.CloudInitConfig != nil {
		return true
	}

	return false
}

// SetCloudInitConfig gets a reference to the given VirtualizationCloudInitConfig and assigns it to the CloudInitConfig field.
func (o *VirtualizationVirtualMachine) SetCloudInitConfig(v VirtualizationCloudInitConfig) {
	o.CloudInitConfig = &v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *VirtualizationVirtualMachine) GetCpu() int64 {
	if o == nil || o.Cpu == nil {
		var ret int64
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualMachine) GetCpuOk() (*int64, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *VirtualizationVirtualMachine) HasCpu() bool {
	if o != nil && o.Cpu != nil {
		return true
	}

	return false
}

// SetCpu gets a reference to the given int64 and assigns it to the Cpu field.
func (o *VirtualizationVirtualMachine) SetCpu(v int64) {
	o.Cpu = &v
}

// GetDiscovered returns the Discovered field value if set, zero value otherwise.
func (o *VirtualizationVirtualMachine) GetDiscovered() bool {
	if o == nil || o.Discovered == nil {
		var ret bool
		return ret
	}
	return *o.Discovered
}

// GetDiscoveredOk returns a tuple with the Discovered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualMachine) GetDiscoveredOk() (*bool, bool) {
	if o == nil || o.Discovered == nil {
		return nil, false
	}
	return o.Discovered, true
}

// HasDiscovered returns a boolean if a field has been set.
func (o *VirtualizationVirtualMachine) HasDiscovered() bool {
	if o != nil && o.Discovered != nil {
		return true
	}

	return false
}

// SetDiscovered gets a reference to the given bool and assigns it to the Discovered field.
func (o *VirtualizationVirtualMachine) SetDiscovered(v bool) {
	o.Discovered = &v
}

// GetDisk returns the Disk field value if set, zero value otherwise.
func (o *VirtualizationVirtualMachine) GetDisk() []VirtualizationVirtualMachineDisk {
	if o == nil || o.Disk == nil {
		var ret []VirtualizationVirtualMachineDisk
		return ret
	}
	return *o.Disk
}

// GetDiskOk returns a tuple with the Disk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualMachine) GetDiskOk() (*[]VirtualizationVirtualMachineDisk, bool) {
	if o == nil || o.Disk == nil {
		return nil, false
	}
	return o.Disk, true
}

// HasDisk returns a boolean if a field has been set.
func (o *VirtualizationVirtualMachine) HasDisk() bool {
	if o != nil && o.Disk != nil {
		return true
	}

	return false
}

// SetDisk gets a reference to the given []VirtualizationVirtualMachineDisk and assigns it to the Disk field.
func (o *VirtualizationVirtualMachine) SetDisk(v []VirtualizationVirtualMachineDisk) {
	o.Disk = &v
}

// GetGuestOs returns the GuestOs field value if set, zero value otherwise.
func (o *VirtualizationVirtualMachine) GetGuestOs() string {
	if o == nil || o.GuestOs == nil {
		var ret string
		return ret
	}
	return *o.GuestOs
}

// GetGuestOsOk returns a tuple with the GuestOs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualMachine) GetGuestOsOk() (*string, bool) {
	if o == nil || o.GuestOs == nil {
		return nil, false
	}
	return o.GuestOs, true
}

// HasGuestOs returns a boolean if a field has been set.
func (o *VirtualizationVirtualMachine) HasGuestOs() bool {
	if o != nil && o.GuestOs != nil {
		return true
	}

	return false
}

// SetGuestOs gets a reference to the given string and assigns it to the GuestOs field.
func (o *VirtualizationVirtualMachine) SetGuestOs(v string) {
	o.GuestOs = &v
}

// GetInterface returns the Interface field value if set, zero value otherwise.
func (o *VirtualizationVirtualMachine) GetInterface() []VirtualizationNetworkInterface {
	if o == nil || o.Interface == nil {
		var ret []VirtualizationNetworkInterface
		return ret
	}
	return *o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualMachine) GetInterfaceOk() (*[]VirtualizationNetworkInterface, bool) {
	if o == nil || o.Interface == nil {
		return nil, false
	}
	return o.Interface, true
}

// HasInterface returns a boolean if a field has been set.
func (o *VirtualizationVirtualMachine) HasInterface() bool {
	if o != nil && o.Interface != nil {
		return true
	}

	return false
}

// SetInterface gets a reference to the given []VirtualizationNetworkInterface and assigns it to the Interface field.
func (o *VirtualizationVirtualMachine) SetInterface(v []VirtualizationNetworkInterface) {
	o.Interface = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *VirtualizationVirtualMachine) GetLabels() []InfraMetaData {
	if o == nil || o.Labels == nil {
		var ret []InfraMetaData
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualMachine) GetLabelsOk() (*[]InfraMetaData, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *VirtualizationVirtualMachine) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []InfraMetaData and assigns it to the Labels field.
func (o *VirtualizationVirtualMachine) SetLabels(v []InfraMetaData) {
	o.Labels = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *VirtualizationVirtualMachine) GetMemory() int64 {
	if o == nil || o.Memory == nil {
		var ret int64
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualMachine) GetMemoryOk() (*int64, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *VirtualizationVirtualMachine) HasMemory() bool {
	if o != nil && o.Memory != nil {
		return true
	}

	return false
}

// SetMemory gets a reference to the given int64 and assigns it to the Memory field.
func (o *VirtualizationVirtualMachine) SetMemory(v int64) {
	o.Memory = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VirtualizationVirtualMachine) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualMachine) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VirtualizationVirtualMachine) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VirtualizationVirtualMachine) SetName(v string) {
	o.Name = &v
}

// GetPowerState returns the PowerState field value if set, zero value otherwise.
func (o *VirtualizationVirtualMachine) GetPowerState() string {
	if o == nil || o.PowerState == nil {
		var ret string
		return ret
	}
	return *o.PowerState
}

// GetPowerStateOk returns a tuple with the PowerState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualMachine) GetPowerStateOk() (*string, bool) {
	if o == nil || o.PowerState == nil {
		return nil, false
	}
	return o.PowerState, true
}

// HasPowerState returns a boolean if a field has been set.
func (o *VirtualizationVirtualMachine) HasPowerState() bool {
	if o != nil && o.PowerState != nil {
		return true
	}

	return false
}

// SetPowerState gets a reference to the given string and assigns it to the PowerState field.
func (o *VirtualizationVirtualMachine) SetPowerState(v string) {
	o.PowerState = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *VirtualizationVirtualMachine) GetCluster() VirtualizationBaseClusterRelationship {
	if o == nil || o.Cluster == nil {
		var ret VirtualizationBaseClusterRelationship
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualMachine) GetClusterOk() (*VirtualizationBaseClusterRelationship, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *VirtualizationVirtualMachine) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given VirtualizationBaseClusterRelationship and assigns it to the Cluster field.
func (o *VirtualizationVirtualMachine) SetCluster(v VirtualizationBaseClusterRelationship) {
	o.Cluster = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *VirtualizationVirtualMachine) GetHost() VirtualizationBaseHostRelationship {
	if o == nil || o.Host == nil {
		var ret VirtualizationBaseHostRelationship
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualMachine) GetHostOk() (*VirtualizationBaseHostRelationship, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *VirtualizationVirtualMachine) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given VirtualizationBaseHostRelationship and assigns it to the Host field.
func (o *VirtualizationVirtualMachine) SetHost(v VirtualizationBaseHostRelationship) {
	o.Host = &v
}

// GetInventory returns the Inventory field value if set, zero value otherwise.
func (o *VirtualizationVirtualMachine) GetInventory() VirtualizationBaseVirtualMachineRelationship {
	if o == nil || o.Inventory == nil {
		var ret VirtualizationBaseVirtualMachineRelationship
		return ret
	}
	return *o.Inventory
}

// GetInventoryOk returns a tuple with the Inventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualMachine) GetInventoryOk() (*VirtualizationBaseVirtualMachineRelationship, bool) {
	if o == nil || o.Inventory == nil {
		return nil, false
	}
	return o.Inventory, true
}

// HasInventory returns a boolean if a field has been set.
func (o *VirtualizationVirtualMachine) HasInventory() bool {
	if o != nil && o.Inventory != nil {
		return true
	}

	return false
}

// SetInventory gets a reference to the given VirtualizationBaseVirtualMachineRelationship and assigns it to the Inventory field.
func (o *VirtualizationVirtualMachine) SetInventory(v VirtualizationBaseVirtualMachineRelationship) {
	o.Inventory = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *VirtualizationVirtualMachine) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualMachine) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *VirtualizationVirtualMachine) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *VirtualizationVirtualMachine) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetWorkflowInfo returns the WorkflowInfo field value if set, zero value otherwise.
func (o *VirtualizationVirtualMachine) GetWorkflowInfo() WorkflowWorkflowInfoRelationship {
	if o == nil || o.WorkflowInfo == nil {
		var ret WorkflowWorkflowInfoRelationship
		return ret
	}
	return *o.WorkflowInfo
}

// GetWorkflowInfoOk returns a tuple with the WorkflowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualMachine) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.WorkflowInfo == nil {
		return nil, false
	}
	return o.WorkflowInfo, true
}

// HasWorkflowInfo returns a boolean if a field has been set.
func (o *VirtualizationVirtualMachine) HasWorkflowInfo() bool {
	if o != nil && o.WorkflowInfo != nil {
		return true
	}

	return false
}

// SetWorkflowInfo gets a reference to the given WorkflowWorkflowInfoRelationship and assigns it to the WorkflowInfo field.
func (o *VirtualizationVirtualMachine) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship) {
	o.WorkflowInfo = &v
}

func (o VirtualizationVirtualMachine) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.Action != nil {
		toSerialize["Action"] = o.Action
	}
	if o.ActionInfo != nil {
		toSerialize["ActionInfo"] = o.ActionInfo
	}
	if o.AffinitySelectors != nil {
		toSerialize["AffinitySelectors"] = o.AffinitySelectors
	}
	if o.AntiAffinitySelectors != nil {
		toSerialize["AntiAffinitySelectors"] = o.AntiAffinitySelectors
	}
	if o.CloudInitConfig != nil {
		toSerialize["CloudInitConfig"] = o.CloudInitConfig
	}
	if o.Cpu != nil {
		toSerialize["Cpu"] = o.Cpu
	}
	if o.Discovered != nil {
		toSerialize["Discovered"] = o.Discovered
	}
	if o.Disk != nil {
		toSerialize["Disk"] = o.Disk
	}
	if o.GuestOs != nil {
		toSerialize["GuestOs"] = o.GuestOs
	}
	if o.Interface != nil {
		toSerialize["Interface"] = o.Interface
	}
	if o.Labels != nil {
		toSerialize["Labels"] = o.Labels
	}
	if o.Memory != nil {
		toSerialize["Memory"] = o.Memory
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.PowerState != nil {
		toSerialize["PowerState"] = o.PowerState
	}
	if o.Cluster != nil {
		toSerialize["Cluster"] = o.Cluster
	}
	if o.Host != nil {
		toSerialize["Host"] = o.Host
	}
	if o.Inventory != nil {
		toSerialize["Inventory"] = o.Inventory
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.WorkflowInfo != nil {
		toSerialize["WorkflowInfo"] = o.WorkflowInfo
	}
	return json.Marshal(toSerialize)
}

type NullableVirtualizationVirtualMachine struct {
	value *VirtualizationVirtualMachine
	isSet bool
}

func (v NullableVirtualizationVirtualMachine) Get() *VirtualizationVirtualMachine {
	return v.value
}

func (v *NullableVirtualizationVirtualMachine) Set(val *VirtualizationVirtualMachine) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVirtualMachine) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVirtualMachine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVirtualMachine(val *VirtualizationVirtualMachine) *NullableVirtualizationVirtualMachine {
	return &NullableVirtualizationVirtualMachine{value: val, isSet: true}
}

func (v NullableVirtualizationVirtualMachine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVirtualMachine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}