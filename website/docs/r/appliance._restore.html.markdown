---
layout: "intersight"
page_title: "Intersight: intersight_appliance._restore"
sidebar_current: "docs-intersight-resource-appliance.Restore"
description: |-
  Restore tracks requests to restore the Intersight Appliance. There will be only
one Restore managed object with a 'Started' state at any time. All other Restore
managed objects will be in terminal states.
---

# Resource: intersight_appliance._restore
Restore tracks requests to restore the Intersight Appliance. There will be only
one Restore managed object with a 'Started' state at any time. All other Restore
managed objects will be in terminal states.
## Argument Reference
The following arguments are supported:
* `account`:(Array with Maximum of one item) -"A reference to a iamAccount resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline."
This complex property has following sub-properties:
  + `class_id`:(string)(Computed)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
  + `link`:(string)"A URL to an instance of the 'mo.MoRef' class."
  + `moid`:(string)(Computed)"The Moid of the referenced REST resource."
  + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
  + `selector`:(string)(Computed)"An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'."
* `class_id`:(string)(Computed)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
* `elapsed_time`:(int)(Computed)"Elapsed time in seconds since the restore process has started."
* `end_time`:(string)(Computed)"End date and time of the restore process."
* `filename`:(string)"Backup filename to backup or restore."
* `is_password_set`:(bool)(Computed)"Indicates whether the value of the 'password' property has been set."
* `messages`:
                (Array of schema.TypeString) -
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `object_type`:(string)(Computed)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `password`:(string)"Password for authenticating with the file server."
* `permission_resources`:(Array)(Computed)"An array of relationships to moBaseMo resources."
This complex property has following sub-properties:
  + `class_id`:(string)(Computed)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
  + `link`:(string)"A URL to an instance of the 'mo.MoRef' class."
  + `moid`:(string)(Computed)"The Moid of the referenced REST resource."
  + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
  + `selector`:(string)(Computed)"An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'."
* `protocol`:(string)"Communication protocol used by the file server (e.g. scp or sftp)."
* `remote_host`:(string)"Hostname of the remote file server."
* `remote_path`:(string)"File server directory to copy the file."
* `remote_port`:(int)"Remote TCP port on the file server (e.g. 22 for scp)."
* `start_time`:(string)(Computed)"Start date and time of the restore process."
* `status`:(string)(Computed)"Status of the restore managed object."
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string)"The string representation of a tag key."
  + `value`:(string)"The string representation of a tag value."
* `username`:(string)"Username to authenticate the fileserver."
