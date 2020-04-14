---
layout: "intersight"
page_title: "Intersight: intersight_tam_advisory_info"
sidebar_current: "docs-intersight-data-source-tamAdvisoryInfo"
description: |-
State of an advisory in the context of a given account. Used to capture a given account's preferences regarding  associated advisory.
---

# Data Source: intersight_tam_advisory_info
State of an advisory in the context of a given account. Used to capture a given account's preferences regarding  associated advisory.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `state`:(string)"Current state of the advisory for the owner. Indicates if the user is interested in getting updates for the advisory."
