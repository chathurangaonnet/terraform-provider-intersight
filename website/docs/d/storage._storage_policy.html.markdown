---
layout: "intersight"
page_title: "Intersight: intersight_storage._storage_policy"
sidebar_current: "docs-intersight-data-source-storage.StoragePolicy"
description: |-
The storage policy models the reusable storage related configuration that can be applied on many servers. This policy allows creation of RAID groups using existing disk group policies and virtual drives on the drive groups. The user has options to move all unused disks to JBOD or Unconfigured good state. The encryption of drives can be enabled through this policy using remote keys from a KMIP server.
---

# Data Source: intersight_storage._storage_policy
The storage policy models the reusable storage related configuration that can be applied on many servers. This policy allows creation of RAID groups using existing disk group policies and virtual drives on the drive groups. The user has options to move all unused disks to JBOD or Unconfigured good state. The encryption of drives can be enabled through this policy using remote keys from a KMIP server.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
* `description`:(string)"Description of the policy."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `name`:(string)"Name of the concrete policy."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `retain_policy_virtual_drives`:(bool)"Retains the virtual drives defined in policy if they exist already. If this flag is false, the existing virtual drives are removed and created again based on virtual drives in the policy."
* `unused_disks_state`:(string)"This is used to specify the state, unconfigured good or jbod, in which the disks that are not used in this policy should be moved."
