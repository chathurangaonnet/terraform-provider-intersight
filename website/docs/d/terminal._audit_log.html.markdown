---
layout: "intersight"
page_title: "Intersight: intersight_terminal._audit_log"
sidebar_current: "docs-intersight-data-source-terminal.AuditLog"
description: |-
Audit log of remote terminal user sessions.
---

# Data Source: intersight_terminal._audit_log
Audit log of remote terminal user sessions.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
* `end_time`:(string)"The time the terminal was closed. If terminal has not closed, value is zero time."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `start_time`:(string)"The time the terminal session was opened."
