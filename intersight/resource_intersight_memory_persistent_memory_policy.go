package intersight

import (
	"log"

	models "github.com/cisco-intersight/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceMemoryPersistentMemoryPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceMemoryPersistentMemoryPolicyCreate,
		Read:   resourceMemoryPersistentMemoryPolicyRead,
		Update: resourceMemoryPersistentMemoryPolicyUpdate,
		Delete: resourceMemoryPersistentMemoryPolicyDelete,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"description": {
				Description: "Description of the policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"goals": {
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
						"memory_mode_percentage": {
							Description: "Volatile memory percentage.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"persistent_memory_type": {
							Description: "Type of the Persistent Memory configuration where the Persistent Memory Modules are combined in an interleaved set or not.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "app-direct",
						},
						"socket_id": {
							Description: "CPU Socket ID to which this goal will be applied.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "All Sockets",
							ForceNew:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
			},
			"local_security": {
				Description: "Local security for the Persistent Memory Modules on the server.",
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
						"enabled": {
							Description: "Persistent Memory Security state.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"is_secure_passphrase_set": {
							Description: "Indicates whether the value of the 'securePassphrase' property has been set.",
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"secure_passphrase": {
							Description: "Secure passphrase to be applied on the Persistent Memory Modules on the server. The allowed characters are a-z, A to Z, 0-9, and special characters =, \\u0021, &, \\#, $, %, +, ^, @, _, *, -.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
			},
			"logical_namespaces": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"capacity": {
							Description: "Capacity of this Namespace that is created or modified.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"mode": {
							Description: "Mode of this Namespace that is created or modified.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "raw",
						},
						"name": {
							Description: "Name of this Namespace to be created on the server.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"socket_id": {
							Description: "Socket ID of the region on which this Namespace has to be created or modified.",
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     1,
						},
						"socket_memory_id": {
							Description: "Socket Memory ID of the region on which this Namespace has to be created or modified.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "Not Applicable",
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
			},
			"management_mode": {
				Description: "Management Mode of the policy. This can be either Configured from Intersight or Configured from Operating System.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "configured-from-intersight",
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
			"name": {
				Description: "Name of the concrete policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"organization": {
				Description: "A reference to a organizationOrganization resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"profiles": {
				Description: "An array of relationships to policyAbstractConfigProfile resources.",
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
			},
			"retain_namespaces": {
				Description: "Persistent Memory Namespaces to be retained or not.",
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
		},
	}
}

func resourceMemoryPersistentMemoryPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewMemoryPersistentMemoryPolicy()
	o.SetClassId("memory.PersistentMemoryPolicy")

	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}

	if v, ok := d.GetOk("goals"); ok {
		x := make([]models.MemoryPersistentMemoryGoal, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewMemoryPersistentMemoryGoalWithDefaults()
			l := s[i].(map[string]interface{})
			o.SetClassId("memory.PersistentMemoryGoal")
			if v, ok := l["memory_mode_percentage"]; ok {
				{
					x := int32(v.(int))
					o.SetMemoryModePercentage(x)
				}
			}
			o.SetObjectType("memory.PersistentMemoryGoal")
			if v, ok := l["persistent_memory_type"]; ok {
				{
					x := (v.(string))
					o.SetPersistentMemoryType(x)
				}
			}
			if v, ok := l["socket_id"]; ok {
				{
					x := (v.(string))
					o.SetSocketId(x)
				}
			}
			x = append(x, *o)
		}
		o.SetGoals(x)
	}

	if v, ok := d.GetOk("local_security"); ok {
		p := make([]models.MemoryPersistentMemoryLocalSecurity, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewMemoryPersistentMemoryLocalSecurityWithDefaults()
			o.SetClassId("memory.PersistentMemoryLocalSecurity")
			if v, ok := l["enabled"]; ok {
				{
					x := (v.(bool))
					o.SetEnabled(x)
				}
			}
			if v, ok := l["is_secure_passphrase_set"]; ok {
				{
					x := (v.(bool))
					o.SetIsSecurePassphraseSet(x)
				}
			}
			o.SetObjectType("memory.PersistentMemoryLocalSecurity")
			if v, ok := l["secure_passphrase"]; ok {
				{
					x := (v.(string))
					o.SetSecurePassphrase(x)
				}
			}
			p = append(p, *o)
		}
		x := p[0]
		o.SetLocalSecurity(x)
	}

	if v, ok := d.GetOk("logical_namespaces"); ok {
		x := make([]models.MemoryPersistentMemoryLogicalNamespace, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewMemoryPersistentMemoryLogicalNamespaceWithDefaults()
			l := s[i].(map[string]interface{})
			if v, ok := l["capacity"]; ok {
				{
					x := int32(v.(int))
					o.SetCapacity(x)
				}
			}
			o.SetClassId("memory.PersistentMemoryLogicalNamespace")
			if v, ok := l["mode"]; ok {
				{
					x := (v.(string))
					o.SetMode(x)
				}
			}
			if v, ok := l["name"]; ok {
				{
					x := (v.(string))
					o.SetName(x)
				}
			}
			o.SetObjectType("memory.PersistentMemoryLogicalNamespace")
			if v, ok := l["socket_id"]; ok {
				{
					x := int32(v.(int))
					o.SetSocketId(x)
				}
			}
			if v, ok := l["socket_memory_id"]; ok {
				{
					x := (v.(string))
					o.SetSocketMemoryId(x)
				}
			}
			x = append(x, *o)
		}
		o.SetLogicalNamespaces(x)
	}

	if v, ok := d.GetOk("management_mode"); ok {
		x := (v.(string))
		o.SetManagementMode(x)
	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}

	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}

	o.SetObjectType("memory.PersistentMemoryPolicy")

	if v, ok := d.GetOk("organization"); ok {
		p := make([]models.OrganizationOrganizationRelationship, 0, 1)
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
			o.SetObjectType("organization.Organization")
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, o.AsOrganizationOrganizationRelationship())
		}
		x := p[0]
		o.SetOrganization(x)
	}

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

	if v, ok := d.GetOk("profiles"); ok {
		x := make([]models.PolicyAbstractConfigProfileRelationship, 0)
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
			o.SetObjectType("policy.AbstractConfigProfile")
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			x = append(x, o.AsPolicyAbstractConfigProfileRelationship())
		}
		o.SetProfiles(x)
	}

	if v, ok := d.GetOk("retain_namespaces"); ok {
		x := (v.(bool))
		o.SetRetainNamespaces(x)
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

	r := conn.ApiClient.MemoryApi.CreateMemoryPersistentMemoryPolicy(conn.ctx).MemoryPersistentMemoryPolicy(*o)
	result, _, err := r.Execute()
	if err != nil {
		log.Panicf("Failed to invoke operation: %v", err)
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return resourceMemoryPersistentMemoryPolicyRead(d, meta)
}
func detachMemoryPersistentMemoryPolicyProfiles(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewMemoryPersistentMemoryPolicy()

	o.Profiles = new([]models.PolicyAbstractConfigProfileRelationship)

	r := conn.ApiClient.MemoryApi.UpdateMemoryPersistentMemoryPolicy(conn.ctx, d.Id()).MemoryPersistentMemoryPolicy(*o)
	_, _, err := r.Execute()
	if err != nil {
		log.Printf("error occurred while creating: %s", err.Error())
	}
	return err
}

func resourceMemoryPersistentMemoryPolicyRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	r := conn.ApiClient.MemoryApi.GetMemoryPersistentMemoryPolicyByMoid(conn.ctx, d.Id())
	s, _, err := r.Execute()

	if err != nil {
		log.Printf("error in unmarshaling model for read Error: %s", err.Error())
		return err
	}

	if err := d.Set("class_id", (s.ClassId)); err != nil {
		return err
	}

	if err := d.Set("description", (s.Description)); err != nil {
		return err
	}

	if err := d.Set("goals", flattenListMemoryPersistentMemoryGoal(s.Goals, d)); err != nil {
		return err
	}

	if err := d.Set("local_security", flattenMapMemoryPersistentMemoryLocalSecurity(s.LocalSecurity, d)); err != nil {
		return err
	}

	if err := d.Set("logical_namespaces", flattenListMemoryPersistentMemoryLogicalNamespace(s.LogicalNamespaces, d)); err != nil {
		return err
	}

	if err := d.Set("management_mode", (s.ManagementMode)); err != nil {
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

	if err := d.Set("organization", flattenMapOrganizationOrganizationRelationship(s.Organization, d)); err != nil {
		return err
	}

	if err := d.Set("permission_resources", flattenListMoBaseMoRelationship(s.PermissionResources, d)); err != nil {
		return err
	}

	if err := d.Set("profiles", flattenListPolicyAbstractConfigProfileRelationship(s.Profiles, d)); err != nil {
		return err
	}

	if err := d.Set("retain_namespaces", (s.RetainNamespaces)); err != nil {
		return err
	}

	if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
		return err
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.GetMoid())
	return nil
}

func resourceMemoryPersistentMemoryPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewMemoryPersistentMemoryPolicy()

	if d.HasChange("description") {
		v := d.Get("description")
		x := (v.(string))
		o.SetDescription(x)
	}

	if d.HasChange("goals") {
		v := d.Get("goals")
		x := make([]models.MemoryPersistentMemoryGoal, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewMemoryPersistentMemoryGoalWithDefaults()
			l := s[i].(map[string]interface{})
			o.SetClassId("memory.PersistentMemoryGoal")
			if v, ok := l["memory_mode_percentage"]; ok {
				{
					x := int32(v.(int))
					o.SetMemoryModePercentage(x)
				}
			}
			o.SetObjectType("memory.PersistentMemoryGoal")
			if v, ok := l["persistent_memory_type"]; ok {
				{
					x := (v.(string))
					o.SetPersistentMemoryType(x)
				}
			}
			if v, ok := l["socket_id"]; ok {
				{
					x := (v.(string))
					o.SetSocketId(x)
				}
			}
			x = append(x, *o)
		}
		o.SetGoals(x)
	}

	if d.HasChange("local_security") {
		v := d.Get("local_security")
		p := make([]models.MemoryPersistentMemoryLocalSecurity, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewMemoryPersistentMemoryLocalSecurityWithDefaults()
			o.SetClassId("memory.PersistentMemoryLocalSecurity")
			if v, ok := l["enabled"]; ok {
				{
					x := (v.(bool))
					o.SetEnabled(x)
				}
			}
			if v, ok := l["is_secure_passphrase_set"]; ok {
				{
					x := (v.(bool))
					o.SetIsSecurePassphraseSet(x)
				}
			}
			o.SetObjectType("memory.PersistentMemoryLocalSecurity")
			if v, ok := l["secure_passphrase"]; ok {
				{
					x := (v.(string))
					o.SetSecurePassphrase(x)
				}
			}
			p = append(p, *o)
		}
		x := p[0]
		o.SetLocalSecurity(x)
	}

	if d.HasChange("logical_namespaces") {
		v := d.Get("logical_namespaces")
		x := make([]models.MemoryPersistentMemoryLogicalNamespace, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewMemoryPersistentMemoryLogicalNamespaceWithDefaults()
			l := s[i].(map[string]interface{})
			if v, ok := l["capacity"]; ok {
				{
					x := int32(v.(int))
					o.SetCapacity(x)
				}
			}
			o.SetClassId("memory.PersistentMemoryLogicalNamespace")
			if v, ok := l["mode"]; ok {
				{
					x := (v.(string))
					o.SetMode(x)
				}
			}
			if v, ok := l["name"]; ok {
				{
					x := (v.(string))
					o.SetName(x)
				}
			}
			o.SetObjectType("memory.PersistentMemoryLogicalNamespace")
			if v, ok := l["socket_id"]; ok {
				{
					x := int32(v.(int))
					o.SetSocketId(x)
				}
			}
			if v, ok := l["socket_memory_id"]; ok {
				{
					x := (v.(string))
					o.SetSocketMemoryId(x)
				}
			}
			x = append(x, *o)
		}
		o.SetLogicalNamespaces(x)
	}

	if d.HasChange("management_mode") {
		v := d.Get("management_mode")
		x := (v.(string))
		o.SetManagementMode(x)
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

	if d.HasChange("organization") {
		v := d.Get("organization")
		p := make([]models.OrganizationOrganizationRelationship, 0, 1)
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
			o.SetObjectType("organization.Organization")
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, o.AsOrganizationOrganizationRelationship())
		}
		x := p[0]
		o.SetOrganization(x)
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

	if d.HasChange("profiles") {
		v := d.Get("profiles")
		x := make([]models.PolicyAbstractConfigProfileRelationship, 0)
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
			o.SetObjectType("policy.AbstractConfigProfile")
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			x = append(x, o.AsPolicyAbstractConfigProfileRelationship())
		}
		o.SetProfiles(x)
	}

	if d.HasChange("retain_namespaces") {
		v := d.Get("retain_namespaces")
		x := (v.(bool))
		o.SetRetainNamespaces(x)
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

	r := conn.ApiClient.MemoryApi.UpdateMemoryPersistentMemoryPolicy(conn.ctx, d.Id()).MemoryPersistentMemoryPolicy(*o)
	result, _, err := r.Execute()
	if err != nil {
		log.Printf("error occurred while updating: %s", err.Error())
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return resourceMemoryPersistentMemoryPolicyRead(d, meta)
}

func resourceMemoryPersistentMemoryPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	e := detachMemoryPersistentMemoryPolicyProfiles(d, meta)
	if e != nil {
		return e
	}

	r := conn.ApiClient.MemoryApi.DeleteMemoryPersistentMemoryPolicy(conn.ctx, d.Id())
	_, err := r.Execute()
	if err != nil {
		log.Printf("error occurred while deleting: %s", err.Error())
	}
	return err
}
