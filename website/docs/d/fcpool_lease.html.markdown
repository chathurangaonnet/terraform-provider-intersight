
---
layout: "intersight"
page_title: "Intersight: intersight_fcpool_lease"
sidebar_current: "docs-intersight-data-source-fcpool-lease"
description: |-
Lease represents a single WWN ID that is part of the universe, allocated either from a pool or through static assignment.
---

# Data Source: intersight_fcpool._lease
Lease represents a single WWN ID that is part of the universe, allocated either from a pool or through static assignment.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `assigned_to_moid`:(string) Moid of the entity/server profile that owns this ID. 
* `assigned_to_type`:(string) Type of the entity that owns this ID. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `pool_purpose`:(string) Purpose of this WWN pool. 
* `wwn_id`:(string) WWN ID allocated for pool based allocation. 
