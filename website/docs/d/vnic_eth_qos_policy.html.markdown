---
layout: "intersight"
page_title: "Intersight: intersight_vnic_eth_qos_policy"
sidebar_current: "docs-intersight-data-source-vnicEthQosPolicy"
description: |-
An Ethernet Quality of Service (QoS) policy assigns a system class to the outgoing traffic for a vNIC. This system class determines the quality of service for the outgoing traffic. For certain adapters you can also specify additional controls like burst and rate on the outgoing traffic.

---

# Data Source: intersight_vnic_eth_qos_policy
An Ethernet Quality of Service (QoS) policy assigns a system class to the outgoing traffic for a vNIC. This system class determines the quality of service for the outgoing traffic. For certain adapters you can also specify additional controls like burst and rate on the outgoing traffic.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `cos`:(int)Class of Service to be associated to the traffic on the virtual interface.
* `description`:(string)Description of the policy.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `mtu`:(int)The Maximum Transmission Unit (MTU) or packet size that the virtual interface accepts.
* `name`:(string)Name of the concrete policy.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `rate_limit`:(int)The value in Mbps (0-40000) to use for limiting the data rate on the virtual interface. Setting this to zero will turn rate limiting off.
* `trust_host_cos`:(bool)Enables usage of the Class of Service provided by the operating system.
