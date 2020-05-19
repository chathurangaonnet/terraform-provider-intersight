package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"time"

	models "github.com/cisco-intersight/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceAssetDeviceContractInformation() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAssetDeviceContractInformationRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"contract": {
				Description: "Contract information for the Cisco support contract purchased for the Cisco device.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bill_to": {
							Description: "BillTo address of listed for the contract.",
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address1": {
										Description: "Address Line one of the address information. example \"PO BOX 641570\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"address2": {
										Description: "Address Line two of the address information. example \"Cisco Systems\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"city": {
										Description: "City in which the address resides. example \"San Jose\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"class_id": {
										Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"country": {
										Description: "Country in which the address resides. example \"US\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"location": {
										Description: "Location in which the address resides. example \"14852\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"name": {
										Description: "Name of the user whose address is being populated.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"object_type": {
										Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"postal_code": {
										Description: "Postal Code in which the address resides. example \"95164-1570\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"state": {
										Description: "State in which the address resides. example \"CA\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"bill_to_global_ultimate": {
							Description: "BillToGlobalUltimate information listed in the contract.",
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"class_id": {
										Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"id": {
										Description: "ID of the user in BillToGlobal.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"name": {
										Description: "Name of the user in BillToGlobal.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"object_type": {
										Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"contract_number": {
							Description: "Contract number for the Cisco support contract purchased for the Cisco device.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"line_status": {
							Description: "Contract status as per the Cisco Contract APIx.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"contract_status": {
				Description: "Calculated contract status that is derived based on the service line status and contract end date. It is different from serviceLineStatus property. serviceLineStatus gives us ACTIVE, OVERDUE, EXPIRED. These are transformed into Active, Expiring Soon and Not Covered.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"covered_product_line_end_date": {
				Description: "End date of the covered product line. The coverage end date is fetched from Cisco SN2INFO API.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"device_id": {
				Description: "Unique identifier of the Cisco device. This information is used to query Cisco APIx SN2INFO and CCWR databases.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"device_type": {
				Description: "Type used to classify the device in Cisco Intersight. Currently supported values are Server and FabricInterconnect. This will be expanded to support more types in future.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"end_customer": {
				Description: "End customer information for the Cisco support contract purchased for the Cisco device.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": {
							Description: "Address as per the information provided by the user.",
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address1": {
										Description: "Address Line one of the address information. example \"PO BOX 641570\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"address2": {
										Description: "Address Line two of the address information. example \"Cisco Systems\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"city": {
										Description: "City in which the address resides. example \"San Jose\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"class_id": {
										Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"country": {
										Description: "Country in which the address resides. example \"US\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"location": {
										Description: "Location in which the address resides. example \"14852\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"name": {
										Description: "Name of the user whose address is being populated.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"object_type": {
										Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"postal_code": {
										Description: "Postal Code in which the address resides. example \"95164-1570\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"state": {
										Description: "State in which the address resides. example \"CA\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Description: "Unique identifier for an end customer. This identifier is allocated by Cisco.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Description: "Name as per the information provided by the user.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"end_user_global_ultimate": {
				Description: "EndUserGlobalUltimate information listed in the contract.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Description: "ID of the user in BillToGlobal.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Description: "Name of the user in BillToGlobal.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"is_valid": {
				Description: "Validates if the device is a genuine Cisco device. Validated is done using the Cisco SN2INFO APIs.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"item_type": {
				Description: "Item type of this specific Cisco device. example \"Chassis\".",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"maintenance_purchase_order_number": {
				Description: "Maintenance purchase order number for the Cisco device.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"maintenance_sales_order_number": {
				Description: "Maintenance sales order number for the Cisco device.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"object_type": {
				Description: "The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"permission_resources": {
				Description: "An array of relationships to moBaseMo resources.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"link": {
							Description: "A URL to an instance of the 'mo.MoRef' class.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"platform_type": {
				Description: "The platform type of the Cisco device.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"product": {
				Description: "Product information of the offering under a contract.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bill_to": {
							Description: "Billing address provided by customer while buying this Cisco product.",
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address1": {
										Description: "Address Line one of the address information. example \"PO BOX 641570\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"address2": {
										Description: "Address Line two of the address information. example \"Cisco Systems\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"city": {
										Description: "City in which the address resides. example \"San Jose\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"class_id": {
										Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"country": {
										Description: "Country in which the address resides. example \"US\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"location": {
										Description: "Location in which the address resides. example \"14852\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"name": {
										Description: "Name of the user whose address is being populated.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"object_type": {
										Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"postal_code": {
										Description: "Postal Code in which the address resides. example \"95164-1570\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"state": {
										Description: "State in which the address resides. example \"CA\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"description": {
							Description: "Short description of the Cisco product that helps identify the product easily. example \"DISTI:UCS 6248UP 1RU Fabric Int/No PSU/32 UP/ 12p LIC\".",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"family": {
							Description: "Family that the product belongs to. Example \"UCSB\".",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"group": {
							Description: "Group that the product belongs to. It is one higher level categorization above family. example \"Switch\".",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"number": {
							Description: "Product number that identifies the product. example PID. example \"UCS-FI-6248UP-CH2\".",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"ship_to": {
							Description: "Shipping address provided by customer while buying this Cisco product.",
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address1": {
										Description: "Address Line one of the address information. example \"PO BOX 641570\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"address2": {
										Description: "Address Line two of the address information. example \"Cisco Systems\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"city": {
										Description: "City in which the address resides. example \"San Jose\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"class_id": {
										Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"country": {
										Description: "Country in which the address resides. example \"US\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"location": {
										Description: "Location in which the address resides. example \"14852\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"name": {
										Description: "Name of the user whose address is being populated.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"object_type": {
										Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"postal_code": {
										Description: "Postal Code in which the address resides. example \"95164-1570\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"state": {
										Description: "State in which the address resides. example \"CA\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"sub_type": {
							Description: "Sub type of the product being specified. example \"UCS 6200 SER\".",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"purchase_order_number": {
				Description: "Purchase order number for the Cisco device. It is a unique number assigned for every purchase.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"registered_device": {
				Description: "A reference to a assetDeviceRegistration resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"link": {
							Description: "A URL to an instance of the 'mo.MoRef' class.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"reseller_global_ultimate": {
				Description: "ResellerGlobalUltimate information listed in the contract.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Description: "ID of the user in BillToGlobal.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Description: "Name of the user in BillToGlobal.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"sales_order_number": {
				Description: "Sales order number for the Cisco device. It is a unique number assigned for every sale.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"service_description": {
				Description: "The type of service contract that covers the Cisco device.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"service_end_date": {
				Description: "End date for the Cisco service contract that covers this Cisco device.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"service_level": {
				Description: "The type of service contract that covers the Cisco device.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"service_sku": {
				Description: "The SKU of the service contract that covers the Cisco device.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"service_start_date": {
				Description: "Start date for the Cisco service contract that covers this Cisco device.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"state_contract": {
				Description: "Internal property used for triggering and tracking actions for contract information.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Description: "The string representation of a tag key.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"value": {
							Description: "The string representation of a tag value.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"warranty_end_date": {
				Description: "End date for the warranty that covers the Cisco device.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"warranty_type": {
				Description: "Type of warranty that covers the Cisco device.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func dataSourceAssetDeviceContractInformationRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewAssetDeviceContractInformation()
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("contract_status"); ok {
		x := (v.(string))
		o.SetContractStatus(x)
	}
	if v, ok := d.GetOk("covered_product_line_end_date"); ok {
		x := (v.(string))
		o.SetCoveredProductLineEndDate(x)
	}
	if v, ok := d.GetOk("device_id"); ok {
		x := (v.(string))
		o.SetDeviceId(x)
	}
	if v, ok := d.GetOk("device_type"); ok {
		x := (v.(string))
		o.SetDeviceType(x)
	}
	if v, ok := d.GetOk("is_valid"); ok {
		x := (v.(bool))
		o.SetIsValid(x)
	}
	if v, ok := d.GetOk("item_type"); ok {
		x := (v.(string))
		o.SetItemType(x)
	}
	if v, ok := d.GetOk("maintenance_purchase_order_number"); ok {
		x := (v.(string))
		o.SetMaintenancePurchaseOrderNumber(x)
	}
	if v, ok := d.GetOk("maintenance_sales_order_number"); ok {
		x := (v.(string))
		o.SetMaintenanceSalesOrderNumber(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("platform_type"); ok {
		x := (v.(string))
		o.SetPlatformType(x)
	}
	if v, ok := d.GetOk("purchase_order_number"); ok {
		x := (v.(string))
		o.SetPurchaseOrderNumber(x)
	}
	if v, ok := d.GetOk("sales_order_number"); ok {
		x := (v.(string))
		o.SetSalesOrderNumber(x)
	}
	if v, ok := d.GetOk("service_description"); ok {
		x := (v.(string))
		o.SetServiceDescription(x)
	}
	if v, ok := d.GetOk("service_end_date"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetServiceEndDate(x)
	}
	if v, ok := d.GetOk("service_level"); ok {
		x := (v.(string))
		o.SetServiceLevel(x)
	}
	if v, ok := d.GetOk("service_sku"); ok {
		x := (v.(string))
		o.SetServiceSku(x)
	}
	if v, ok := d.GetOk("service_start_date"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetServiceStartDate(x)
	}
	if v, ok := d.GetOk("state_contract"); ok {
		x := (v.(string))
		o.SetStateContract(x)
	}
	if v, ok := d.GetOk("warranty_end_date"); ok {
		x := (v.(string))
		o.SetWarrantyEndDate(x)
	}
	if v, ok := d.GetOk("warranty_type"); ok {
		x := (v.(string))
		o.SetWarrantyType(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return fmt.Errorf("Json Marshalling of data source failed with error : %+v", err)
	}
	result, _, err := conn.ApiClient.AssetApi.GetAssetDeviceContractInformationList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if err != nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = models.NewAssetDeviceContractInformation()
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return err
			}
			if err := d.Set("class_id", (s.ClassId)); err != nil {
				return err
			}

			if err := d.Set("contract", flattenMapAssetContractInformation(s.Contract, d)); err != nil {
				return err
			}
			if err := d.Set("contract_status", (s.ContractStatus)); err != nil {
				return err
			}
			if err := d.Set("covered_product_line_end_date", (s.CoveredProductLineEndDate)); err != nil {
				return err
			}
			if err := d.Set("device_id", (s.DeviceId)); err != nil {
				return err
			}
			if err := d.Set("device_type", (s.DeviceType)); err != nil {
				return err
			}

			if err := d.Set("end_customer", flattenMapAssetCustomerInformation(s.EndCustomer, d)); err != nil {
				return err
			}

			if err := d.Set("end_user_global_ultimate", flattenMapAssetGlobalUltimate(s.EndUserGlobalUltimate, d)); err != nil {
				return err
			}
			if err := d.Set("is_valid", (s.IsValid)); err != nil {
				return err
			}
			if err := d.Set("item_type", (s.ItemType)); err != nil {
				return err
			}
			if err := d.Set("maintenance_purchase_order_number", (s.MaintenancePurchaseOrderNumber)); err != nil {
				return err
			}
			if err := d.Set("maintenance_sales_order_number", (s.MaintenanceSalesOrderNumber)); err != nil {
				return err
			}
			if err := d.Set("moid", (s.Moid)); err != nil {
				return err
			}
			if err := d.Set("object_type", (s.ObjectType)); err != nil {
				return err
			}

			if err := d.Set("permission_resources", flattenListMoBaseMoRelationship(s.PermissionResources, d)); err != nil {
				return err
			}
			if err := d.Set("platform_type", (s.PlatformType)); err != nil {
				return err
			}

			if err := d.Set("product", flattenMapAssetProductInformation(s.Product, d)); err != nil {
				return err
			}
			if err := d.Set("purchase_order_number", (s.PurchaseOrderNumber)); err != nil {
				return err
			}

			if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRelationship(s.RegisteredDevice, d)); err != nil {
				return err
			}

			if err := d.Set("reseller_global_ultimate", flattenMapAssetGlobalUltimate(s.ResellerGlobalUltimate, d)); err != nil {
				return err
			}
			if err := d.Set("sales_order_number", (s.SalesOrderNumber)); err != nil {
				return err
			}
			if err := d.Set("service_description", (s.ServiceDescription)); err != nil {
				return err
			}

			if err := d.Set("service_end_date", (s.ServiceEndDate).String()); err != nil {
				return err
			}
			if err := d.Set("service_level", (s.ServiceLevel)); err != nil {
				return err
			}
			if err := d.Set("service_sku", (s.ServiceSku)); err != nil {
				return err
			}

			if err := d.Set("service_start_date", (s.ServiceStartDate).String()); err != nil {
				return err
			}
			if err := d.Set("state_contract", (s.StateContract)); err != nil {
				return err
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}
			if err := d.Set("warranty_end_date", (s.WarrantyEndDate)); err != nil {
				return err
			}
			if err := d.Set("warranty_type", (s.WarrantyType)); err != nil {
				return err
			}
			d.SetId(s.GetMoid())
		}
	}
	return nil
}
