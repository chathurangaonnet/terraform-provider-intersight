package intersight

import (
	"log"

	models "github.com/cisco-intersight/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceHyperflexVcenterConfigPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceHyperflexVcenterConfigPolicyCreate,
		Read:   resourceHyperflexVcenterConfigPolicyRead,
		Update: resourceHyperflexVcenterConfigPolicyUpdate,
		Delete: resourceHyperflexVcenterConfigPolicyDelete,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"cluster_profiles": {
				Description: "An array of relationships to hyperflexClusterProfile resources.",
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
			"data_center": {
				Description: "The vCenter datacenter name.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": {
				Description: "Description of the policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hostname": {
				Description: "The vCenter server FQDN or IP.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"is_password_set": {
				Description: "Indicates whether the value of the 'password' property has been set.",
				Type:        schema.TypeBool,
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
			"password": {
				Description: "The password for authenticating with vCenter. Follow the corresponding password policy governed by vCenter.",
				Type:        schema.TypeString,
				Optional:    true,
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
			"sso_url": {
				Description: "Overrides the default vCenter Single Sign-On URL. Do not specify unless instructed by Cisco TAC.",
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
			"username": {
				Description: "The vCenter username (e.g. administrator@vsphere.local).",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func resourceHyperflexVcenterConfigPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewHyperflexVcenterConfigPolicy()
	o.SetClassId("hyperflex.VcenterConfigPolicy")

	if v, ok := d.GetOk("cluster_profiles"); ok {
		x := make([]models.HyperflexClusterProfileRelationship, 0)
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
			o.SetObjectType("hyperflex.ClusterProfile")
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			x = append(x, o.AsHyperflexClusterProfileRelationship())
		}
		o.SetClusterProfiles(x)
	}

	if v, ok := d.GetOk("data_center"); ok {
		x := (v.(string))
		o.SetDataCenter(x)
	}

	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}

	if v, ok := d.GetOk("hostname"); ok {
		x := (v.(string))
		o.SetHostname(x)
	}

	if v, ok := d.GetOk("is_password_set"); ok {
		x := (v.(bool))
		o.SetIsPasswordSet(x)
	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}

	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}

	o.SetObjectType("hyperflex.VcenterConfigPolicy")

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

	if v, ok := d.GetOk("password"); ok {
		x := (v.(string))
		o.SetPassword(x)
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

	if v, ok := d.GetOk("sso_url"); ok {
		x := (v.(string))
		o.SetSsoUrl(x)
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

	if v, ok := d.GetOk("username"); ok {
		x := (v.(string))
		o.SetUsername(x)
	}

	r := conn.ApiClient.HyperflexApi.CreateHyperflexVcenterConfigPolicy(conn.ctx).HyperflexVcenterConfigPolicy(*o)
	result, _, err := r.Execute()
	if err != nil {
		log.Panicf("Failed to invoke operation: %v", err)
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return resourceHyperflexVcenterConfigPolicyRead(d, meta)
}

func resourceHyperflexVcenterConfigPolicyRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	r := conn.ApiClient.HyperflexApi.GetHyperflexVcenterConfigPolicyByMoid(conn.ctx, d.Id())
	s, _, err := r.Execute()

	if err != nil {
		log.Printf("error in unmarshaling model for read Error: %s", err.Error())
		return err
	}

	if err := d.Set("class_id", (s.ClassId)); err != nil {
		return err
	}

	if err := d.Set("cluster_profiles", flattenListHyperflexClusterProfileRelationship(s.ClusterProfiles, d)); err != nil {
		return err
	}

	if err := d.Set("data_center", (s.DataCenter)); err != nil {
		return err
	}

	if err := d.Set("description", (s.Description)); err != nil {
		return err
	}

	if err := d.Set("hostname", (s.Hostname)); err != nil {
		return err
	}

	if err := d.Set("is_password_set", (s.IsPasswordSet)); err != nil {
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

	if err := d.Set("sso_url", (s.SsoUrl)); err != nil {
		return err
	}

	if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
		return err
	}

	if err := d.Set("username", (s.Username)); err != nil {
		return err
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.GetMoid())
	return nil
}

func resourceHyperflexVcenterConfigPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewHyperflexVcenterConfigPolicy()

	if d.HasChange("cluster_profiles") {
		v := d.Get("cluster_profiles")
		x := make([]models.HyperflexClusterProfileRelationship, 0)
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
			o.SetObjectType("hyperflex.ClusterProfile")
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			x = append(x, o.AsHyperflexClusterProfileRelationship())
		}
		o.SetClusterProfiles(x)
	}

	if d.HasChange("data_center") {
		v := d.Get("data_center")
		x := (v.(string))
		o.SetDataCenter(x)
	}

	if d.HasChange("description") {
		v := d.Get("description")
		x := (v.(string))
		o.SetDescription(x)
	}

	if d.HasChange("hostname") {
		v := d.Get("hostname")
		x := (v.(string))
		o.SetHostname(x)
	}

	if d.HasChange("is_password_set") {
		v := d.Get("is_password_set")
		x := (v.(bool))
		o.SetIsPasswordSet(x)
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

	if d.HasChange("password") {
		v := d.Get("password")
		x := (v.(string))
		o.SetPassword(x)
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

	if d.HasChange("sso_url") {
		v := d.Get("sso_url")
		x := (v.(string))
		o.SetSsoUrl(x)
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

	if d.HasChange("username") {
		v := d.Get("username")
		x := (v.(string))
		o.SetUsername(x)
	}

	r := conn.ApiClient.HyperflexApi.UpdateHyperflexVcenterConfigPolicy(conn.ctx, d.Id()).HyperflexVcenterConfigPolicy(*o)
	result, _, err := r.Execute()
	if err != nil {
		log.Printf("error occurred while updating: %s", err.Error())
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return resourceHyperflexVcenterConfigPolicyRead(d, meta)
}

func resourceHyperflexVcenterConfigPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	r := conn.ApiClient.HyperflexApi.DeleteHyperflexVcenterConfigPolicy(conn.ctx, d.Id())
	_, err := r.Execute()
	if err != nil {
		log.Printf("error occurred while deleting: %s", err.Error())
	}
	return err
}
