
---
layout: "intersight"
page_title: "Intersight: intersight_storage_physical_disk"
sidebar_current: "docs-intersight-data-source-storage-physical-disk"
description: |-
Physical Disk on a server.
---

# Data Source: intersight_storage._physical_disk
Physical Disk on a server.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `block_size`:(string) The block size of the physical disk in bytes. 
* `bootable`:(string) This field identifies the disk drive as bootable if set to true. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `configuration_checkpoint`:(string) The current configuration checkpoint of the physical disk. 
* `configuration_state`:(string) The current configuration state of the physical disk. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `discovered_path`:(string) The discovered path of the physical disk. 
* `disk_id`:(string) This field identifies the ID assigned to physical disks. 
* `disk_state`:(string) This field identifies the health of the disk. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `drive_firmware`:(string) This field identifies the disk firmware running in the disk. 
* `drive_state`:(string) The drive state as reported by the controller. 
* `fde_capable`:(string) Full-Disk Encryption capability parameter of the physical disk. 
* `hot_spare_type`:(string) Type of hotspare configured on the physical disk. 
* `link_speed`:(string) The speed of the link between the drive and the controller. 
* `link_state`:(string) The current link state of the physical disk. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `num_blocks`:(string) The number of blocks present on the physical disk. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `oper_power_state`:(string) Operational power of the physical disk. 
* `oper_qualifier_reason`:(string)
* `operability`:(string) This field identifies the disk operability of the disk. 
* `physical_block_size`:(string) The block size of the installed physical disk. 
* `pid`:(string) This field identifies the Product ID for physicalDisk. 
* `presence`:(string) The presence state of the physical disk. 
* `protocol`:(string) This field identifies the disk protocol used for communication. 
* `raw_size`:(string) The raw size of the physical disk in MB. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `secured`:(string) This field identifies whether the disk is encrypted. 
* `serial`:(string) This field identifies the serial of the given component. 
* `size`:(string) The size of the physical disk in MB. 
* `thermal`:(string) Thermal state of the physical disk. 
* `type`:(string) This field identifies the type of the physical disk. 
* `variant_type`:(string) The variant type of the physical disk. 
* `vendor`:(string) This field identifies the vendor of the given component. 
