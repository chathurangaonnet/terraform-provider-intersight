
---
layout: "intersight"
page_title: "Intersight: intersight_capability_equipment_physical_def"
sidebar_current: "docs-intersight-data-source-capability-equipment-physical-def"
description: |-
Type to represent additional switch specific capabilities.
---

# Data Source: intersight_capability._equipment_physical_def
Type to represent additional switch specific capabilities.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `depth`:(float) Depth information for a Switch/Fabric-Interconnect. 
* `height`:(float) Height information for a Switch/Fabric-Interconnect. 
* `max_power`:(int) Max Power information for a Switch/Fabric-Interconnect. 
* `min_power`:(int) Min Power information for a Switch/Fabric-Interconnect. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `nominal_power`:(int) Nominal Power information for a Switch/Fabric-Interconnect. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `pid`:(string) Product Identifier for a Switch/Fabric-Interconnect. 
* `sku`:(string) SKU information for Switch/Fabric-Interconnect. 
* `vid`:(string) VID information for Switch/Fabric-Interconnect. 
* `weight`:(float) Weight information for a Switch/Fabric-Interconnect. 
* `width`:(float) Width information for a Switch/Fabric-Interconnect. 
