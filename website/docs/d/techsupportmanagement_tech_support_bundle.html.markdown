
---
layout: "intersight"
page_title: "Intersight: intersight_techsupportmanagement_tech_support_bundle"
sidebar_current: "docs-intersight-data-source-techsupportmanagement-tech-support-bundle"
description: |-
A request to collect techsupport and upload it to Intersight Storage Service. The serial number, PID and/or relationship to the target resource provided by the user is used to determine the device type for techsupport collection. If the serial number, PID and target resource are specified in the request, the values must match. Valid values of device types are network.Element for fabric interconnect, compute.Blade for blade server, compute.RackUnit for rack server, equipment.Chassis for chassis, equipment.IoCard for IO Module, equipment.FEX for fabric extender and adapter.Unit for network adapter. UCSM techsupport is collected for device type network.Element. Chassis techsupport is collected for compute.Blade, equipment.Chassis, equipment.IoCard, and blade adapter.Unit. Server techsupport is collected for compute.RackUnit and rack adapter.Unit. Fabric extender techsupport is collected for device type equipment.FEX. Hyper Flex node level techsupport is collected when the request specifies the platform type (HX) and the device type is Hyperflex.Node.
---

# Data Source: intersight_techsupportmanagement._tech_support_bundle
A request to collect techsupport and upload it to Intersight Storage Service. The serial number, PID and/or relationship to the target resource provided by the user is used to determine the device type for techsupport collection. If the serial number, PID and target resource are specified in the request, the values must match. Valid values of device types are network.Element for fabric interconnect, compute.Blade for blade server, compute.RackUnit for rack server, equipment.Chassis for chassis, equipment.IoCard for IO Module, equipment.FEX for fabric extender and adapter.Unit for network adapter. UCSM techsupport is collected for device type network.Element. Chassis techsupport is collected for compute.Blade, equipment.Chassis, equipment.IoCard, and blade adapter.Unit. Server techsupport is collected for compute.RackUnit and rack adapter.Unit. Fabric extender techsupport is collected for device type equipment.FEX. Hyper Flex node level techsupport is collected when the request specifies the platform type (HX) and the device type is Hyperflex.Node.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `device_identifier`:(string) The device identifier used to uniquely identify an individual device. 
* `device_type`:(string) The device type obtained from the inventory. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `pid`:(string) Product identification of the device. 
* `platform_type`:(string) The platform type of the device. 
* `serial`:(string) Serial number of the device. 
