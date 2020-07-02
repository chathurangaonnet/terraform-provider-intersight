
---
layout: "intersight"
page_title: "Intersight: intersight_fabric_fcoe_uplink_pc_role"
sidebar_current: "docs-intersight-data-source-fabric-fcoe-uplink-pc-role"
description: |-
Object sent by user to configure a fcoe uplink port-channel on the collection of ports.
---

# Data Source: intersight_fabric._fcoe_uplink_pc_role
Object sent by user to configure a fcoe uplink port-channel on the collection of ports.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `admin_speed`:(string) Admin configured speed for the port. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `pc_id`:(int) Unique Identifier of the port-channel, local to this switch. 
* `udld_admin_state`:(string) Admin configured state for UDLD for this port. 
