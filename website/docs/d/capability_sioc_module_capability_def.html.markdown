
---
layout: "intersight"
page_title: "Intersight: intersight_capability_sioc_module_capability_def"
sidebar_current: "docs-intersight-data-source-capability-sioc-module-capability-def"
description: |-
Chassis SIOC module capabilities.
---

# Data Source: intersight_capability._sioc_module_capability_def
Chassis SIOC module capabilities.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `dc_supported`:(bool) Device connector support on SIOC. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
