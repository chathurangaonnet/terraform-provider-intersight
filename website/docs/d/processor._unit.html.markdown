---
layout: "intersight"
page_title: "Intersight: intersight_processor._unit"
sidebar_current: "docs-intersight-data-source-processor.Unit"
description: |-
The CPU present on a server.
---

# Data Source: intersight_processor._unit
The CPU present on a server.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `architecture`:(string)
* `class_id`:(string)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
* `device_mo_id`:(string)
* `dn`:(string)"The Distinguished Name unambiguously identifies an object in the system."
* `model`:(string)"This field identifies the model of the given component."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `num_cores`:(int)
* `num_cores_enabled`:(string)
* `num_threads`:(string)
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `oper_power_state`:(string)
* `oper_state`:(string)
* `operability`:(string)
* `presence`:(string)
* `processor_id`:(int)
* `revision`:(string)
* `rn`:(string)"The Relative Name uniquely identifies an object within a given context."
* `serial`:(string)"This field identifies the serial of the given component."
* `socket_designation`:(string)
* `speed`:(float)
* `stepping`:(string)
* `thermal`:(string)
* `vendor`:(string)"This field identifies the vendor of the given component."
