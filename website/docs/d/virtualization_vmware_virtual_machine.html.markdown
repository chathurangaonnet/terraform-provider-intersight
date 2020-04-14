---
layout: "intersight"
page_title: "Intersight: intersight_virtualization_vmware_virtual_machine"
sidebar_current: "docs-intersight-data-source-virtualizationVmwareVirtualMachine"
description: |-
The VMware Virtual machine. It has details such as power state, IP address, resource consumption, etc. Basic elements come from the base class and VMware specific details are provided here.
---

# Data Source: intersight_virtualization_vmware_virtual_machine
The VMware Virtual machine. It has details such as power state, IP address, resource consumption, etc. Basic elements come from the base class and VMware specific details are provided here.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `annotation`:(string)"List of annotations provided to this VM by user. Can be long."
* `boot_time`:(string)"Time when this VM booted up."
* `cpu_hot_add_enabled`:(bool)"Indicates if the capability to add CPUs to a running VM is enabled."
* `class_id`:(string)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
* `config_name`:(string)"The configuration name for this VM. This maybe the same as the guest hostname."
* `connection_state`:(string)"Shows if virtual machine is connected to vCenter. Values are Connected, Disconnected, Orphaned, Inaccessible, and Invalid."
* `default_power_off_type`:(string)"Indicates how the VM will be powered off (soft, hard etc.)."
* `dhcp_enabled`:(bool)"Shows if DHCP is used for IP/DNS on this VM."
* `folder`:(string)"The folder name associated with this VM."
* `guest_state`:(string)"The state of the guest OS running on this VM. Could be running, not running etc."
* `hypervisor_type`:(string)"Type of hypervisor where the virtual machine is hosted, for example VMware ESXi."
* `identity`:(string)"The internally generated identity of this VM. This entity is not manipulated by users. It aids in uniquely identifying the virtual machine object. For VMware, this is MOR (managed object reference)."
* `instance_uuid`:(string)"UUID assigned by vCenter to every VM."
* `is_template`:(bool)"If true, indicates that the entity refers to a template of a virtual machine and not a real virtual machine."
* `memory_hot_add_enabled`:(bool)"Adding memory to a running VM."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `name`:(string)"User-provided name to identify the virtual machine."
* `network_count`:(int)"Indicates how many networks are used by this VM."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `power_state`:(string)"Power state of the virtual machine."
* `protected_vm`:(bool)"Shows if this is a protected VM. VMs can be in protection groups."
* `remote_display_vnc_enabled`:(bool)"Shows if support for a remote VNC access is enabled."
* `resource_pool`:(string)"Name of the resource pool to which this VM belongs (optional)."
* `resource_pool_owner`:(string)"Who owns the resource pool."
* `resource_pool_parent`:(string)"The parent of the current resource pool to which this VM belongs."
* `tool_running_status`:(string)"Indicates if guest tools are running on this VM. Could be set to guestToolNotRunning or guestToolsRunning."
* `tools_version`:(string)"The version of the guest tools, usually not specified."
* `uuid`:(string)"The uuid of this virtual machine. The uuid is internally generated and not user specified."
* `vm_disk_count`:(int)"Shows the number of disks assigned to this VM."
* `vm_overall_status`:(string)"The operational state of the VM. Could be Available, Provisioned, Maintenance mode, Deleting, etc."
* `vm_path`:(string)"Example - [datastore3] VCSA-134/VCSA-134.vmx."
* `vm_version`:(string)"Information about the version of this VM (vmx-09, vmx-11 etc.)."
* `vm_vnic_count`:(int)"How many vnics are present."
* `vnic_device_config_id`:(string)"Information related to the guest info's VNIC virtual device. It is a comma-separated list."
