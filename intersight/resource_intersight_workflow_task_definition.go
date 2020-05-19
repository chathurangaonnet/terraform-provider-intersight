package intersight

import (
	"log"

	models "github.com/cisco-intersight/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceWorkflowTaskDefinition() *schema.Resource {
	return &schema.Resource{
		Create: resourceWorkflowTaskDefinitionCreate,
		Read:   resourceWorkflowTaskDefinitionRead,
		Update: resourceWorkflowTaskDefinitionUpdate,
		Delete: resourceWorkflowTaskDefinitionDelete,
		Schema: map[string]*schema.Schema{
			"catalog": {
				Description: "A reference to a workflowCatalog resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
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
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
			},
			"class_id": {
				Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"default_version": {
				Description: "When true this will be the task version that is used when a specific task definition version is not specified. The very first task definition created with a name will be set as the default version, after that user can explicitly set any version of the task definition as the default version.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"description": {
				Description: "The task definition description to describe what this task will do when executed.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"implemented_tasks": {
				Description: "An array of relationships to workflowTaskDefinition resources.",
				Type:        schema.TypeList,
				Optional:    true,
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
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
				ForceNew:   true,
			},
			"interface_task": {
				Description: "A reference to a workflowTaskDefinition resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
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
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
				ForceNew:   true,
			},
			"internal_properties": {
				Description: "Type to capture all the internal properties for the task definition.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"base_task_type": {
							Description: "This field will hold the base task type like HttpBaseTask or RemoteAnsibleBaseTask.",
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
						"constraints": {
							Description: "This field will hold any constraints a concrete task definition will specify in order to limit the environment where the task can execute.",
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
									"object_type": {
										Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"target_data_type": {
										Description: "List of property constraints that helps to narrow down task implementations based on target device input.",
										Type:        schema.TypeMap,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										}, Optional: true,
									},
								},
							},
							ConfigMode: schema.SchemaConfigModeAttr,
						},
						"internal": {
							Description: "Denotes this is an internal task. Internal tasks will be hidden from the UI when executing a workflow.",
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"owner": {
							Description: "The service that owns and is responsible for execution of the task.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
			},
			"label": {
				Description: "A user friendly short name to identify the task definition.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"license_entitlement": {
				Description: "License entitlement required to run this task. It is determined by license requirement of features.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
			"name": {
				Description: "The name of the task definition. The name should follow this convention <Verb or Action><Category><Vendor><Product><Noun or object> Verb or Action is a required portion of the name and this must be part of the pre-approved verb list. Category is an optional field and this will refer to the broad category of the task referring to the type of resource or endpoint. If there is no specific category then use \"Generic\" if required. Vendor is an optional field and this will refer to the specific vendor this task applies to. If the task is generic and not tied to a vendor, then do not specify anything. Product is an optional field, this will contain the vendor product and model when desired. Noun or object is a required field and  this will contain the noun or object on which the action is being performed. Examples SendEmail  - This is a task in Generic category for sending email. NewStorageVolume - This is a vendor agnostic task under Storage device category for creating a new volume.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
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
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
			},
			"properties": {
				Description: "Type to capture all the properties for the task definition.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"external_meta": {
							Description: "When set to false the task definition can only be used by internal system workflows. When set to true then the task can be included in user defined workflows.",
							Type:        schema.TypeBool,
							Optional:    true,
							ForceNew:    true,
						},
						"input_definition": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"class_id": {
										Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"default": {
										Description: "Default value for the data type. If default value was provided and the input was required the default value will be used as the input.",
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"class_id": {
													Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"object_type": {
													Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"override": {
													Description: "Override the default value provided for the data type. When true, allow the user to enter value for the data type.",
													Type:        schema.TypeBool,
													Optional:    true,
												},
												"value": {
													Description: "Default value for the data type. If default value was provided and the input was required the default value will be used as the input.",
													Type:        schema.TypeMap,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													}, Optional: true,
												},
											},
										},
										ConfigMode: schema.SchemaConfigModeAttr,
										Computed:   true,
									},
									"description": {
										Description: "Provide a detailed description of the data type.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"label": {
										Description: "Descriptive label for the data type. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), space ( ) or an underscore (_). The first and last character in label must be an alphanumeric character.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"name": {
										Description: "Descriptive name for the data type. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-) or an underscore (_). The first and last character in name must be an alphanumeric character.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"object_type": {
										Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"required": {
										Description: "Specifies whether this parameter is required. The field is applicable for task and workflow.",
										Type:        schema.TypeBool,
										Optional:    true,
									},
								},
							},
							ConfigMode: schema.SchemaConfigModeAttr,
							Computed:   true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"output_definition": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"class_id": {
										Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"default": {
										Description: "Default value for the data type. If default value was provided and the input was required the default value will be used as the input.",
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"class_id": {
													Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"object_type": {
													Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"override": {
													Description: "Override the default value provided for the data type. When true, allow the user to enter value for the data type.",
													Type:        schema.TypeBool,
													Optional:    true,
												},
												"value": {
													Description: "Default value for the data type. If default value was provided and the input was required the default value will be used as the input.",
													Type:        schema.TypeMap,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													}, Optional: true,
												},
											},
										},
										ConfigMode: schema.SchemaConfigModeAttr,
										Computed:   true,
									},
									"description": {
										Description: "Provide a detailed description of the data type.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"label": {
										Description: "Descriptive label for the data type. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), space ( ) or an underscore (_). The first and last character in label must be an alphanumeric character.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"name": {
										Description: "Descriptive name for the data type. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-) or an underscore (_). The first and last character in name must be an alphanumeric character.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"object_type": {
										Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"required": {
										Description: "Specifies whether this parameter is required. The field is applicable for task and workflow.",
										Type:        schema.TypeBool,
										Optional:    true,
									},
								},
							},
							ConfigMode: schema.SchemaConfigModeAttr,
							Computed:   true,
						},
						"retry_count": {
							Description: "The number of times a task should be tried before marking as failed.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"retry_delay": {
							Description: "The delay in seconds after which the the task is re-tried.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"retry_policy": {
							Description: "The retry policy for the task.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "Fixed",
						},
						"support_status": {
							Description: "Supported status of the definition.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "Supported",
						},
						"timeout": {
							Description: "The timeout value in seconds after which task will be marked as timed out. Max allowed value is 7 days.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"timeout_policy": {
							Description: "The timeout policy for the task.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "Timeout",
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
			},
			"secure_prop_access": {
				Description: "If set to true, the task requires access to secure properties and uses an encyption token associated with a workflow moid to encrypt or decrypt the secure properties.",
				Type:        schema.TypeBool,
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
			"version": {
				Description: "The version of the task definition so we can support multiple versions of a task definition.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
		},
	}
}

func resourceWorkflowTaskDefinitionCreate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewWorkflowTaskDefinition()
	if v, ok := d.GetOk("catalog"); ok {
		p := make([]models.WorkflowCatalogRelationship, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewMoMoRefWithDefaults()
			o.SetClassId("mo.MoRef")
			if v, ok := l["link"]; ok {
				{
					x := (v.(string))
					o.SetLink(x)
				}
			}
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			o.SetObjectType("workflow.Catalog")
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, o.AsWorkflowCatalogRelationship())
		}
		x := p[0]
		o.SetCatalog(x)
	}

	o.SetClassId("workflow.TaskDefinition")

	if v, ok := d.GetOk("default_version"); ok {
		x := (v.(bool))
		o.SetDefaultVersion(x)
	}

	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}

	if v, ok := d.GetOk("implemented_tasks"); ok {
		x := make([]models.WorkflowTaskDefinitionRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewMoMoRefWithDefaults()
			l := s[i].(map[string]interface{})
			o.SetClassId("mo.MoRef")
			if v, ok := l["link"]; ok {
				{
					x := (v.(string))
					o.SetLink(x)
				}
			}
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			o.SetObjectType("workflow.TaskDefinition")
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			x = append(x, o.AsWorkflowTaskDefinitionRelationship())
		}
		o.SetImplementedTasks(x)
	}

	if v, ok := d.GetOk("interface_task"); ok {
		p := make([]models.WorkflowTaskDefinitionRelationship, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewMoMoRefWithDefaults()
			o.SetClassId("mo.MoRef")
			if v, ok := l["link"]; ok {
				{
					x := (v.(string))
					o.SetLink(x)
				}
			}
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			o.SetObjectType("workflow.TaskDefinition")
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, o.AsWorkflowTaskDefinitionRelationship())
		}
		x := p[0]
		o.SetInterfaceTask(x)
	}

	if v, ok := d.GetOk("internal_properties"); ok {
		p := make([]models.WorkflowInternalProperties, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewWorkflowInternalPropertiesWithDefaults()
			if v, ok := l["base_task_type"]; ok {
				{
					x := (v.(string))
					o.SetBaseTaskType(x)
				}
			}
			o.SetClassId("workflow.InternalProperties")
			if v, ok := l["constraints"]; ok {
				{
					p := make([]models.WorkflowTaskConstraints, 0, 1)
					l := (v.([]interface{})[0]).(map[string]interface{})
					{
						o := models.NewWorkflowTaskConstraintsWithDefaults()
						o.SetClassId("workflow.TaskConstraints")
						o.SetObjectType("workflow.TaskConstraints")
						if v, ok := l["target_data_type"]; ok {
							{
								x := v.(map[string]interface{})
								o.SetTargetDataType(x)
							}
						}
						p = append(p, *o)
					}
					x := p[0]
					o.SetConstraints(x)
				}
			}
			if v, ok := l["internal"]; ok {
				{
					x := (v.(bool))
					o.SetInternal(x)
				}
			}
			o.SetObjectType("workflow.InternalProperties")
			if v, ok := l["owner"]; ok {
				{
					x := (v.(string))
					o.SetOwner(x)
				}
			}
			p = append(p, *o)
		}
		x := p[0]
		o.SetInternalProperties(x)
	}

	if v, ok := d.GetOk("label"); ok {
		x := (v.(string))
		o.SetLabel(x)
	}

	if v, ok := d.GetOk("license_entitlement"); ok {
		x := (v.(string))
		o.SetLicenseEntitlement(x)
	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}

	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}

	o.SetObjectType("workflow.TaskDefinition")

	if v, ok := d.GetOk("permission_resources"); ok {
		x := make([]models.MoBaseMoRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewMoMoRefWithDefaults()
			l := s[i].(map[string]interface{})
			o.SetClassId("mo.MoRef")
			if v, ok := l["link"]; ok {
				{
					x := (v.(string))
					o.SetLink(x)
				}
			}
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			o.SetObjectType("mo.BaseMo")
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			x = append(x, o.AsMoBaseMoRelationship())
		}
		o.SetPermissionResources(x)
	}

	if v, ok := d.GetOk("properties"); ok {
		p := make([]models.WorkflowProperties, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewWorkflowPropertiesWithDefaults()
			o.SetClassId("workflow.Properties")
			if v, ok := l["external_meta"]; ok {
				{
					x := (v.(bool))
					o.SetExternalMeta(x)
				}
			}
			if v, ok := l["input_definition"]; ok {
				{
					x := make([]models.WorkflowBaseDataType, 0)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						o := models.NewWorkflowBaseDataTypeWithDefaults()
						l := s[i].(map[string]interface{})
						o.SetClassId("workflow.BaseDataType")
						if v, ok := l["default"]; ok {
							{
								p := make([]models.WorkflowDefaultValue, 0, 1)
								l := (v.([]interface{})[0]).(map[string]interface{})
								{
									o := models.NewWorkflowDefaultValueWithDefaults()
									o.SetClassId("workflow.DefaultValue")
									o.SetObjectType("workflow.DefaultValue")
									if v, ok := l["override"]; ok {
										{
											x := (v.(bool))
											o.SetOverride(x)
										}
									}
									if v, ok := l["value"]; ok {
										{
											x := v.(map[string]interface{})
											o.SetValue(x)
										}
									}
									p = append(p, *o)
								}
								x := p[0]
								o.SetDefault(x)
							}
						}
						if v, ok := l["description"]; ok {
							{
								x := (v.(string))
								o.SetDescription(x)
							}
						}
						if v, ok := l["label"]; ok {
							{
								x := (v.(string))
								o.SetLabel(x)
							}
						}
						if v, ok := l["name"]; ok {
							{
								x := (v.(string))
								o.SetName(x)
							}
						}
						o.SetObjectType("workflow.BaseDataType")
						if v, ok := l["required"]; ok {
							{
								x := (v.(bool))
								o.SetRequired(x)
							}
						}
						x = append(x, *o)
					}
					o.SetInputDefinition(x)
				}
			}
			o.SetObjectType("workflow.Properties")
			if v, ok := l["output_definition"]; ok {
				{
					x := make([]models.WorkflowBaseDataType, 0)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						o := models.NewWorkflowBaseDataTypeWithDefaults()
						l := s[i].(map[string]interface{})
						o.SetClassId("workflow.BaseDataType")
						if v, ok := l["default"]; ok {
							{
								p := make([]models.WorkflowDefaultValue, 0, 1)
								l := (v.([]interface{})[0]).(map[string]interface{})
								{
									o := models.NewWorkflowDefaultValueWithDefaults()
									o.SetClassId("workflow.DefaultValue")
									o.SetObjectType("workflow.DefaultValue")
									if v, ok := l["override"]; ok {
										{
											x := (v.(bool))
											o.SetOverride(x)
										}
									}
									if v, ok := l["value"]; ok {
										{
											x := v.(map[string]interface{})
											o.SetValue(x)
										}
									}
									p = append(p, *o)
								}
								x := p[0]
								o.SetDefault(x)
							}
						}
						if v, ok := l["description"]; ok {
							{
								x := (v.(string))
								o.SetDescription(x)
							}
						}
						if v, ok := l["label"]; ok {
							{
								x := (v.(string))
								o.SetLabel(x)
							}
						}
						if v, ok := l["name"]; ok {
							{
								x := (v.(string))
								o.SetName(x)
							}
						}
						o.SetObjectType("workflow.BaseDataType")
						if v, ok := l["required"]; ok {
							{
								x := (v.(bool))
								o.SetRequired(x)
							}
						}
						x = append(x, *o)
					}
					o.SetOutputDefinition(x)
				}
			}
			if v, ok := l["retry_count"]; ok {
				{
					x := int64(v.(int))
					o.SetRetryCount(x)
				}
			}
			if v, ok := l["retry_delay"]; ok {
				{
					x := int64(v.(int))
					o.SetRetryDelay(x)
				}
			}
			if v, ok := l["retry_policy"]; ok {
				{
					x := (v.(string))
					o.SetRetryPolicy(x)
				}
			}
			if v, ok := l["support_status"]; ok {
				{
					x := (v.(string))
					o.SetSupportStatus(x)
				}
			}
			if v, ok := l["timeout"]; ok {
				{
					x := int64(v.(int))
					o.SetTimeout(x)
				}
			}
			if v, ok := l["timeout_policy"]; ok {
				{
					x := (v.(string))
					o.SetTimeoutPolicy(x)
				}
			}
			p = append(p, *o)
		}
		x := p[0]
		o.SetProperties(x)
	}

	if v, ok := d.GetOk("secure_prop_access"); ok {
		x := (v.(bool))
		o.SetSecurePropAccess(x)
	}

	if v, ok := d.GetOk("tags"); ok {
		x := make([]models.MoTag, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewMoTagWithDefaults()
			l := s[i].(map[string]interface{})
			if v, ok := l["key"]; ok {
				{
					x := (v.(string))
					o.SetKey(x)
				}
			}
			if v, ok := l["value"]; ok {
				{
					x := (v.(string))
					o.SetValue(x)
				}
			}
			x = append(x, *o)
		}
		o.SetTags(x)
	}

	if v, ok := d.GetOk("version"); ok {
		x := int64(v.(int))
		o.SetVersion(x)
	}

	r := conn.ApiClient.WorkflowApi.CreateWorkflowTaskDefinition(conn.ctx).WorkflowTaskDefinition(*o)
	result, _, err := r.Execute()
	if err != nil {
		log.Panicf("Failed to invoke operation: %v", err)
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return resourceWorkflowTaskDefinitionRead(d, meta)
}

func resourceWorkflowTaskDefinitionRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	r := conn.ApiClient.WorkflowApi.GetWorkflowTaskDefinitionByMoid(conn.ctx, d.Id())
	s, _, err := r.Execute()

	if err != nil {
		log.Printf("error in unmarshaling model for read Error: %s", err.Error())
		return err
	}

	if err := d.Set("catalog", flattenMapWorkflowCatalogRelationship(s.Catalog, d)); err != nil {
		return err
	}

	if err := d.Set("class_id", (s.ClassId)); err != nil {
		return err
	}

	if err := d.Set("default_version", (s.DefaultVersion)); err != nil {
		return err
	}

	if err := d.Set("description", (s.Description)); err != nil {
		return err
	}

	if err := d.Set("implemented_tasks", flattenListWorkflowTaskDefinitionRelationship(s.ImplementedTasks, d)); err != nil {
		return err
	}

	if err := d.Set("interface_task", flattenMapWorkflowTaskDefinitionRelationship(s.InterfaceTask, d)); err != nil {
		return err
	}

	if err := d.Set("internal_properties", flattenMapWorkflowInternalProperties(s.InternalProperties, d)); err != nil {
		return err
	}

	if err := d.Set("label", (s.Label)); err != nil {
		return err
	}

	if err := d.Set("license_entitlement", (s.LicenseEntitlement)); err != nil {
		return err
	}

	if err := d.Set("moid", (s.Moid)); err != nil {
		return err
	}

	if err := d.Set("name", (s.Name)); err != nil {
		return err
	}

	if err := d.Set("object_type", (s.ObjectType)); err != nil {
		return err
	}

	if err := d.Set("permission_resources", flattenListMoBaseMoRelationship(s.PermissionResources, d)); err != nil {
		return err
	}

	if err := d.Set("properties", flattenMapWorkflowProperties(s.Properties, d)); err != nil {
		return err
	}

	if err := d.Set("secure_prop_access", (s.SecurePropAccess)); err != nil {
		return err
	}

	if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
		return err
	}

	if err := d.Set("version", (s.Version)); err != nil {
		return err
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.GetMoid())
	return nil
}

func resourceWorkflowTaskDefinitionUpdate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewWorkflowTaskDefinition()
	if d.HasChange("catalog") {
		v := d.Get("catalog")
		p := make([]models.WorkflowCatalogRelationship, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewMoMoRefWithDefaults()
			o.SetClassId("mo.MoRef")
			if v, ok := l["link"]; ok {
				{
					x := (v.(string))
					o.SetLink(x)
				}
			}
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			o.SetObjectType("workflow.Catalog")
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, o.AsWorkflowCatalogRelationship())
		}
		x := p[0]
		o.SetCatalog(x)
	}

	if d.HasChange("default_version") {
		v := d.Get("default_version")
		x := (v.(bool))
		o.SetDefaultVersion(x)
	}

	if d.HasChange("description") {
		v := d.Get("description")
		x := (v.(string))
		o.SetDescription(x)
	}

	if d.HasChange("implemented_tasks") {
		v := d.Get("implemented_tasks")
		x := make([]models.WorkflowTaskDefinitionRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewMoMoRefWithDefaults()
			l := s[i].(map[string]interface{})
			o.SetClassId("mo.MoRef")
			if v, ok := l["link"]; ok {
				{
					x := (v.(string))
					o.SetLink(x)
				}
			}
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			o.SetObjectType("workflow.TaskDefinition")
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			x = append(x, o.AsWorkflowTaskDefinitionRelationship())
		}
		o.SetImplementedTasks(x)
	}

	if d.HasChange("interface_task") {
		v := d.Get("interface_task")
		p := make([]models.WorkflowTaskDefinitionRelationship, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewMoMoRefWithDefaults()
			o.SetClassId("mo.MoRef")
			if v, ok := l["link"]; ok {
				{
					x := (v.(string))
					o.SetLink(x)
				}
			}
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			o.SetObjectType("workflow.TaskDefinition")
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, o.AsWorkflowTaskDefinitionRelationship())
		}
		x := p[0]
		o.SetInterfaceTask(x)
	}

	if d.HasChange("internal_properties") {
		v := d.Get("internal_properties")
		p := make([]models.WorkflowInternalProperties, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewWorkflowInternalPropertiesWithDefaults()
			if v, ok := l["base_task_type"]; ok {
				{
					x := (v.(string))
					o.SetBaseTaskType(x)
				}
			}
			o.SetClassId("workflow.InternalProperties")
			if v, ok := l["constraints"]; ok {
				{
					p := make([]models.WorkflowTaskConstraints, 0, 1)
					l := (v.([]interface{})[0]).(map[string]interface{})
					{
						o := models.NewWorkflowTaskConstraintsWithDefaults()
						o.SetClassId("workflow.TaskConstraints")
						o.SetObjectType("workflow.TaskConstraints")
						if v, ok := l["target_data_type"]; ok {
							{
								x := v.(map[string]interface{})
								o.SetTargetDataType(x)
							}
						}
						p = append(p, *o)
					}
					x := p[0]
					o.SetConstraints(x)
				}
			}
			if v, ok := l["internal"]; ok {
				{
					x := (v.(bool))
					o.SetInternal(x)
				}
			}
			o.SetObjectType("workflow.InternalProperties")
			if v, ok := l["owner"]; ok {
				{
					x := (v.(string))
					o.SetOwner(x)
				}
			}
			p = append(p, *o)
		}
		x := p[0]
		o.SetInternalProperties(x)
	}

	if d.HasChange("label") {
		v := d.Get("label")
		x := (v.(string))
		o.SetLabel(x)
	}

	if d.HasChange("license_entitlement") {
		v := d.Get("license_entitlement")
		x := (v.(string))
		o.SetLicenseEntitlement(x)
	}

	if d.HasChange("moid") {
		v := d.Get("moid")
		x := (v.(string))
		o.SetMoid(x)
	}

	if d.HasChange("name") {
		v := d.Get("name")
		x := (v.(string))
		o.SetName(x)
	}

	if d.HasChange("permission_resources") {
		v := d.Get("permission_resources")
		x := make([]models.MoBaseMoRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewMoMoRefWithDefaults()
			l := s[i].(map[string]interface{})
			o.SetClassId("mo.MoRef")
			if v, ok := l["link"]; ok {
				{
					x := (v.(string))
					o.SetLink(x)
				}
			}
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			o.SetObjectType("mo.BaseMo")
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			x = append(x, o.AsMoBaseMoRelationship())
		}
		o.SetPermissionResources(x)
	}

	if d.HasChange("properties") {
		v := d.Get("properties")
		p := make([]models.WorkflowProperties, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewWorkflowPropertiesWithDefaults()
			o.SetClassId("workflow.Properties")
			if v, ok := l["external_meta"]; ok {
				{
					x := (v.(bool))
					o.SetExternalMeta(x)
				}
			}
			if v, ok := l["input_definition"]; ok {
				{
					x := make([]models.WorkflowBaseDataType, 0)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						o := models.NewWorkflowBaseDataTypeWithDefaults()
						l := s[i].(map[string]interface{})
						o.SetClassId("workflow.BaseDataType")
						if v, ok := l["default"]; ok {
							{
								p := make([]models.WorkflowDefaultValue, 0, 1)
								l := (v.([]interface{})[0]).(map[string]interface{})
								{
									o := models.NewWorkflowDefaultValueWithDefaults()
									o.SetClassId("workflow.DefaultValue")
									o.SetObjectType("workflow.DefaultValue")
									if v, ok := l["override"]; ok {
										{
											x := (v.(bool))
											o.SetOverride(x)
										}
									}
									if v, ok := l["value"]; ok {
										{
											x := v.(map[string]interface{})
											o.SetValue(x)
										}
									}
									p = append(p, *o)
								}
								x := p[0]
								o.SetDefault(x)
							}
						}
						if v, ok := l["description"]; ok {
							{
								x := (v.(string))
								o.SetDescription(x)
							}
						}
						if v, ok := l["label"]; ok {
							{
								x := (v.(string))
								o.SetLabel(x)
							}
						}
						if v, ok := l["name"]; ok {
							{
								x := (v.(string))
								o.SetName(x)
							}
						}
						o.SetObjectType("workflow.BaseDataType")
						if v, ok := l["required"]; ok {
							{
								x := (v.(bool))
								o.SetRequired(x)
							}
						}
						x = append(x, *o)
					}
					o.SetInputDefinition(x)
				}
			}
			o.SetObjectType("workflow.Properties")
			if v, ok := l["output_definition"]; ok {
				{
					x := make([]models.WorkflowBaseDataType, 0)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						o := models.NewWorkflowBaseDataTypeWithDefaults()
						l := s[i].(map[string]interface{})
						o.SetClassId("workflow.BaseDataType")
						if v, ok := l["default"]; ok {
							{
								p := make([]models.WorkflowDefaultValue, 0, 1)
								l := (v.([]interface{})[0]).(map[string]interface{})
								{
									o := models.NewWorkflowDefaultValueWithDefaults()
									o.SetClassId("workflow.DefaultValue")
									o.SetObjectType("workflow.DefaultValue")
									if v, ok := l["override"]; ok {
										{
											x := (v.(bool))
											o.SetOverride(x)
										}
									}
									if v, ok := l["value"]; ok {
										{
											x := v.(map[string]interface{})
											o.SetValue(x)
										}
									}
									p = append(p, *o)
								}
								x := p[0]
								o.SetDefault(x)
							}
						}
						if v, ok := l["description"]; ok {
							{
								x := (v.(string))
								o.SetDescription(x)
							}
						}
						if v, ok := l["label"]; ok {
							{
								x := (v.(string))
								o.SetLabel(x)
							}
						}
						if v, ok := l["name"]; ok {
							{
								x := (v.(string))
								o.SetName(x)
							}
						}
						o.SetObjectType("workflow.BaseDataType")
						if v, ok := l["required"]; ok {
							{
								x := (v.(bool))
								o.SetRequired(x)
							}
						}
						x = append(x, *o)
					}
					o.SetOutputDefinition(x)
				}
			}
			if v, ok := l["retry_count"]; ok {
				{
					x := int64(v.(int))
					o.SetRetryCount(x)
				}
			}
			if v, ok := l["retry_delay"]; ok {
				{
					x := int64(v.(int))
					o.SetRetryDelay(x)
				}
			}
			if v, ok := l["retry_policy"]; ok {
				{
					x := (v.(string))
					o.SetRetryPolicy(x)
				}
			}
			if v, ok := l["support_status"]; ok {
				{
					x := (v.(string))
					o.SetSupportStatus(x)
				}
			}
			if v, ok := l["timeout"]; ok {
				{
					x := int64(v.(int))
					o.SetTimeout(x)
				}
			}
			if v, ok := l["timeout_policy"]; ok {
				{
					x := (v.(string))
					o.SetTimeoutPolicy(x)
				}
			}
			p = append(p, *o)
		}
		x := p[0]
		o.SetProperties(x)
	}

	if d.HasChange("secure_prop_access") {
		v := d.Get("secure_prop_access")
		x := (v.(bool))
		o.SetSecurePropAccess(x)
	}

	if d.HasChange("tags") {
		v := d.Get("tags")
		x := make([]models.MoTag, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewMoTagWithDefaults()
			l := s[i].(map[string]interface{})
			if v, ok := l["key"]; ok {
				{
					x := (v.(string))
					o.SetKey(x)
				}
			}
			if v, ok := l["value"]; ok {
				{
					x := (v.(string))
					o.SetValue(x)
				}
			}
			x = append(x, *o)
		}
		o.SetTags(x)
	}

	if d.HasChange("version") {
		v := d.Get("version")
		x := int64(v.(int))
		o.SetVersion(x)
	}

	r := conn.ApiClient.WorkflowApi.UpdateWorkflowTaskDefinition(conn.ctx, d.Id()).WorkflowTaskDefinition(*o)
	result, _, err := r.Execute()
	if err != nil {
		log.Printf("error occurred while updating: %s", err.Error())
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return resourceWorkflowTaskDefinitionRead(d, meta)
}

func resourceWorkflowTaskDefinitionDelete(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	r := conn.ApiClient.WorkflowApi.DeleteWorkflowTaskDefinition(conn.ctx, d.Id())
	_, err := r.Execute()
	if err != nil {
		log.Printf("error occurred while deleting: %s", err.Error())
	}
	return err
}
