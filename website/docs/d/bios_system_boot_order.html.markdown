
---
layout: "intersight"
page_title: "Intersight: intersight_bios_system_boot_order"
sidebar_current: "docs-intersight-data-source-bios-system-boot-order"
description: |-
Actual Boot Order of the system.
---

# Data Source: intersight_bios._system_boot_order
Actual Boot Order of the system.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `boot_mode`:(string) The BIOS boot mode. UEFI uses the GUID Partition Table (GPT) whereas Legacy mode uses the Master Boot Record (MBR) partitioning scheme. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `dn`:(string) The Distinguished Name for this object, used to uniquely identify this object. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `secure_boot`:(string) Secure boot if set to enabled, enforces that device boots using only software that is trusted by the Original Equipment Manufacturer (OEM). 
