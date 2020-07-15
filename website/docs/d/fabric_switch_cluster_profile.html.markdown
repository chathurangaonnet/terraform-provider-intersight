
---
layout: "intersight"
page_title: "Intersight: intersight_fabric_switch_cluster_profile"
sidebar_current: "docs-intersight-data-source-fabric-switch-cluster-profile"
description: |-
This specifies the configuration policies for a cluster of switches.
---

# Data Source: intersight_fabric._switch_cluster_profile
This specifies the configuration policies for a cluster of switches.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `description`:(string) Description of the profile. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete profile. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `type`:(string) Defines the type of the profile. Accepted value is instance. 
