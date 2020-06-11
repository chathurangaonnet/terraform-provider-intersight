---
layout: "intersight"
page_title: "Intersight: intersight_ether_physical_port"
sidebar_current: "docs-intersight-data-source-etherPhysicalPort"
description: |-
Physical ethernet port present on a FI.
---

# Data Source: intersight_ether_physical_port
Physical ethernet port present on a FI.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
* `device_mo_id`:(string)
* `dn`:(string)"The Distinguished Name unambiguously identifies an object in the system."
* `mac_address`:(string)
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `oper_state`:(string)
* `peer_dn`:(string)"PeerDn for ethernet physical port."
* `rn`:(string)"The Relative Name uniquely identifies an object within a given context."
* `role`:(string)
* `transceiver_type`:(string)
