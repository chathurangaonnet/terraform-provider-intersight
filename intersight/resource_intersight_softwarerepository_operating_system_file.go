package intersight

import (
	"log"

	models "github.com/cisco-intersight/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSoftwarerepositoryOperatingSystemFile() *schema.Resource {
	return &schema.Resource{
		Create: resourceSoftwarerepositoryOperatingSystemFileCreate,
		Read:   resourceSoftwarerepositoryOperatingSystemFileRead,
		Update: resourceSoftwarerepositoryOperatingSystemFileUpdate,
		Delete: resourceSoftwarerepositoryOperatingSystemFileDelete,
		Schema: map[string]*schema.Schema{
			"catalog": {
				Description: "A reference to a softwarerepositoryCatalog resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"class_id": {
				Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"description": {
				Description: "User provided description about the file. Cisco provided description for image inventoried from a Cisco repository.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"download_count": {
				Description: "The number of times this file has been downloaded from the local repository. It is used by the repository monitoring process to determine the files that are to be evicted from the cache.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"import_action": {
				Description: "The action to be performed on the imported file. If 'PreCache' is set, the image will be cached in Appliance. Applicable in Intersight appliance deployment. If 'Evict' is set, the cached file will be removed. Applicable in Intersight appliance deployment. If 'GeneratePreSignedUploadUrl' is set, generates pre signed URL (s) for the file to be imported into the repository. Applicable for local machine source. The URL (s) will be populated under LocalMachine file server. If 'CompleteImportProcess' is set, the ImportState is marked as 'Imported'. Applicable for local machine source. If 'Cancel' is set, the ImportState is marked as 'Failed'. Applicable for local machine source.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "None",
			},
			"import_state": {
				Description: "The state  of this file in the repository or Appliance. The importState is updated during the import operation and as part of the repository monitoring process.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"md5sum": {
				Description: "The md5sum checksum of the file. This information is available for all Cisco distributed images and files imported to the local repository.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
			"name": {
				Description: "The name of the file. It is populated as part of the image import operation.",
				Type:        schema.TypeString,
				Optional:    true,
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
			"sha512sum": {
				Description: "The sha512sum of the file. This information is available for all Cisco distributed images and files imported to the local repository.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"size": {
				Description: "The size (in bytes) of the file. This information is available for all Cisco distributed images and files imported to the local repository.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"software_advisory_url": {
				Description: "The software advisory, if any, provided by the vendor for this file.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"source": {
				Description: "Location of the file in an external repository.",
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
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
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
			"vendor": {
				Description: "The vendor or publisher of this file.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"version": {
				Description: "Vendor provided version for the file.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func resourceSoftwarerepositoryOperatingSystemFileCreate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewSoftwarerepositoryOperatingSystemFile()
	if v, ok := d.GetOk("catalog"); ok {
		p := make([]models.SoftwarerepositoryCatalogRelationship, 0, 1)
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
			o.SetObjectType("softwarerepository.Catalog")
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, o.AsSoftwarerepositoryCatalogRelationship())
		}
		x := p[0]
		o.SetCatalog(x)
	}

	o.SetClassId("softwarerepository.OperatingSystemFile")

	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}

	if v, ok := d.GetOk("download_count"); ok {
		x := int64(v.(int))
		o.SetDownloadCount(x)
	}

	if v, ok := d.GetOk("import_action"); ok {
		x := (v.(string))
		o.SetImportAction(x)
	}

	if v, ok := d.GetOk("import_state"); ok {
		x := (v.(string))
		o.SetImportState(x)
	}

	if v, ok := d.GetOk("md5sum"); ok {
		x := (v.(string))
		o.SetMd5sum(x)
	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}

	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}

	o.SetObjectType("softwarerepository.OperatingSystemFile")

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

	if v, ok := d.GetOk("sha512sum"); ok {
		x := (v.(string))
		o.SetSha512sum(x)
	}

	if v, ok := d.GetOk("size"); ok {
		x := int64(v.(int))
		o.SetSize(x)
	}

	if v, ok := d.GetOk("software_advisory_url"); ok {
		x := (v.(string))
		o.SetSoftwareAdvisoryUrl(x)
	}

	if v, ok := d.GetOk("source"); ok {
		p := make([]models.SoftwarerepositoryFileServer, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewSoftwarerepositoryFileServerWithDefaults()
			o.SetClassId("softwarerepository.FileServer")
			o.SetObjectType("softwarerepository.FileServer")
			p = append(p, *o)
		}
		x := p[0]
		o.SetSource(x)
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

	if v, ok := d.GetOk("vendor"); ok {
		x := (v.(string))
		o.SetVendor(x)
	}

	if v, ok := d.GetOk("version"); ok {
		x := (v.(string))
		o.SetVersion(x)
	}

	r := conn.ApiClient.SoftwarerepositoryApi.CreateSoftwarerepositoryOperatingSystemFile(conn.ctx).SoftwarerepositoryOperatingSystemFile(*o)
	result, _, err := r.Execute()
	if err != nil {
		log.Panicf("Failed to invoke operation: %v", err)
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return resourceSoftwarerepositoryOperatingSystemFileRead(d, meta)
}

func resourceSoftwarerepositoryOperatingSystemFileRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	r := conn.ApiClient.SoftwarerepositoryApi.GetSoftwarerepositoryOperatingSystemFileByMoid(conn.ctx, d.Id())
	s, _, err := r.Execute()

	if err != nil {
		log.Printf("error in unmarshaling model for read Error: %s", err.Error())
		return err
	}

	if err := d.Set("catalog", flattenMapSoftwarerepositoryCatalogRelationship(s.Catalog, d)); err != nil {
		return err
	}

	if err := d.Set("class_id", (s.ClassId)); err != nil {
		return err
	}

	if err := d.Set("description", (s.Description)); err != nil {
		return err
	}

	if err := d.Set("download_count", (s.DownloadCount)); err != nil {
		return err
	}

	if err := d.Set("import_action", (s.ImportAction)); err != nil {
		return err
	}

	if err := d.Set("import_state", (s.ImportState)); err != nil {
		return err
	}

	if err := d.Set("md5sum", (s.Md5sum)); err != nil {
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

	if err := d.Set("sha512sum", (s.Sha512sum)); err != nil {
		return err
	}

	if err := d.Set("size", (s.Size)); err != nil {
		return err
	}

	if err := d.Set("software_advisory_url", (s.SoftwareAdvisoryUrl)); err != nil {
		return err
	}

	if err := d.Set("source", flattenMapSoftwarerepositoryFileServer(s.Source, d)); err != nil {
		return err
	}

	if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
		return err
	}

	if err := d.Set("vendor", (s.Vendor)); err != nil {
		return err
	}

	if err := d.Set("version", (s.Version)); err != nil {
		return err
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.GetMoid())
	return nil
}

func resourceSoftwarerepositoryOperatingSystemFileUpdate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewSoftwarerepositoryOperatingSystemFile()
	if d.HasChange("catalog") {
		v := d.Get("catalog")
		p := make([]models.SoftwarerepositoryCatalogRelationship, 0, 1)
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
			o.SetObjectType("softwarerepository.Catalog")
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, o.AsSoftwarerepositoryCatalogRelationship())
		}
		x := p[0]
		o.SetCatalog(x)
	}

	if d.HasChange("description") {
		v := d.Get("description")
		x := (v.(string))
		o.SetDescription(x)
	}

	if d.HasChange("download_count") {
		v := d.Get("download_count")
		x := int64(v.(int))
		o.SetDownloadCount(x)
	}

	if d.HasChange("import_action") {
		v := d.Get("import_action")
		x := (v.(string))
		o.SetImportAction(x)
	}

	if d.HasChange("import_state") {
		v := d.Get("import_state")
		x := (v.(string))
		o.SetImportState(x)
	}

	if d.HasChange("md5sum") {
		v := d.Get("md5sum")
		x := (v.(string))
		o.SetMd5sum(x)
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

	if d.HasChange("sha512sum") {
		v := d.Get("sha512sum")
		x := (v.(string))
		o.SetSha512sum(x)
	}

	if d.HasChange("size") {
		v := d.Get("size")
		x := int64(v.(int))
		o.SetSize(x)
	}

	if d.HasChange("software_advisory_url") {
		v := d.Get("software_advisory_url")
		x := (v.(string))
		o.SetSoftwareAdvisoryUrl(x)
	}

	if d.HasChange("source") {
		v := d.Get("source")
		p := make([]models.SoftwarerepositoryFileServer, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewSoftwarerepositoryFileServerWithDefaults()
			o.SetClassId("softwarerepository.FileServer")
			o.SetObjectType("softwarerepository.FileServer")
			p = append(p, *o)
		}
		x := p[0]
		o.SetSource(x)
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

	if d.HasChange("vendor") {
		v := d.Get("vendor")
		x := (v.(string))
		o.SetVendor(x)
	}

	if d.HasChange("version") {
		v := d.Get("version")
		x := (v.(string))
		o.SetVersion(x)
	}

	r := conn.ApiClient.SoftwarerepositoryApi.UpdateSoftwarerepositoryOperatingSystemFile(conn.ctx, d.Id()).SoftwarerepositoryOperatingSystemFile(*o)
	result, _, err := r.Execute()
	if err != nil {
		log.Printf("error occurred while updating: %s", err.Error())
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return resourceSoftwarerepositoryOperatingSystemFileRead(d, meta)
}

func resourceSoftwarerepositoryOperatingSystemFileDelete(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	r := conn.ApiClient.SoftwarerepositoryApi.DeleteSoftwarerepositoryOperatingSystemFile(conn.ctx, d.Id())
	_, err := r.Execute()
	if err != nil {
		log.Printf("error occurred while deleting: %s", err.Error())
	}
	return err
}
