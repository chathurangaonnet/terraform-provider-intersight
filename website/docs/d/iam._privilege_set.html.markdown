---
layout: "intersight"
page_title: "Intersight: intersight_iam._privilege_set"
sidebar_current: "docs-intersight-data-source-iam.PrivilegeSet"
description: |-
A collection of privileges.
---

# Data Source: intersight_iam._privilege_set
A collection of privileges.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
* `description`:(string)"Description of the privilege set."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `name`:(string)"Name of the privilege set."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
