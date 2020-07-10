
---
layout: "intersight"
page_title: "Intersight: intersight_iam_api_key"
sidebar_current: "docs-intersight-resource-iam-api-key"
description: |-
  API keys are used to programatically perform API calls. API keys can be created by passing purpose field, which is a description for the application using API keys. Maximum of 3 API keys per user is allowed. API key will have RSA key pair generated and as part of create request key MOID, private and public key in PEM format are returned. In Intersight only the public key is stored. Client side private key is stored and is used for signature calculation. API key requests are authenticated using RSA SHA 256 signature validations. If the incoming request for authorization has X-Starship-Token, then session based authorization is done, else API key based authorization is performed. Once User, API key and Account are found and signature verification succeeds, then the privilege validations are performed.
---

# Resource: intersight_iam._api_key
API keys are used to programatically perform API calls. API keys can be created by passing purpose field, which is a description for the application using API keys. Maximum of 3 API keys per user is allowed. API key will have RSA key pair generated and as part of create request key MOID, private and public key in PEM format are returned. In Intersight only the public key is stored. Client side private key is stored and is used for signature calculation. API key requests are authenticated using RSA SHA 256 signature validations. If the incoming request for authorization has X-Starship-Token, then session based authorization is done, else API key based authorization is performed. Once User, API key and Account are found and signature verification succeeds, then the privilege validations are performed.
## Argument Reference
The following arguments are supported:
* `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `hash_algorithm`:(string) The cryptographic hash algorithm to calculate the message digest. 
* `key_spec`:(Array with Maximum of one item) - The key generation specification provides the algorithm and the parameters required for this algorithm to generate a private key, public key pair. Supported key generation schemes include RSA, ECDSA and Edwards-Curve Digital Signature Algorithm (EdDSA). 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
  + `name`:(string)(Computed) Name of the key generation algorithm. 
  + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string)(Computed) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `permission`:(Array with Maximum of one item) -(Computed) A reference to a iamPermission resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `private_key`:(string) Holds the private key for the API key. 
* `purpose`:(string) The purpose of the API Key. 
* `signing_algorithm`:(string) The signing algorithm used by the client to authenticate API requests to Intersight.The signing algorithm must be compatible with the key generation specification. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `user`:(Array with Maximum of one item) -(Computed) A reference to a iamUser resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
