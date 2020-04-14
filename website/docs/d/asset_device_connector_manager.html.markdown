---
layout: "intersight"
page_title: "Intersight: intersight_asset_device_connector_manager"
sidebar_current: "docs-intersight-data-source-assetDeviceConnectorManager"
description: |-
Information pertaining to a Registered Intersight Assist Appliance Device in Intersight.
---

# Data Source: intersight_asset_device_connector_manager
Information pertaining to a Registered Intersight Assist Appliance Device in Intersight.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
