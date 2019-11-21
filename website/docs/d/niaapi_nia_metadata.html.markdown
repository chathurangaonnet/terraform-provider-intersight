---
layout: "intersight"
page_title: "Intersight: intersight_niaapi_nia_metadata"
sidebar_current: "docs-intersight-data-source-niaapiNiaMetadata"
description: |-
Contains the latest Metadata available for download from server.

---

# Data Source: intersight_niaapi_nia_metadata
Contains the latest Metadata available for download from server.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `date`:(string)The date when this package is generated.
* `metadata_chksum`:(string)Chksum used to check the integrity of the Metadata file downloaded.
* `metadata_filename`:(string)The Filename of this Metadata package.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `version`:(int)The version number of the Metadata package.
