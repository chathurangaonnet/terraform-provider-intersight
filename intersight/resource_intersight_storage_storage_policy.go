package intersight

import (
	"encoding/json"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceStorageStoragePolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceStorageStoragePolicyCreate,
		Read:   resourceStorageStoragePolicyRead,
		Update: resourceStorageStoragePolicyUpdate,
		Delete: resourceStorageStoragePolicyDelete,
		Schema: map[string]*schema.Schema{
			"description": {
				Description: "Description of the policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"disk_group_policies": {
				Description: "Relationship to the used disk group policies.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The Object Type of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'\nis set and 'moid' is empty/absent from the request, Intersight will determine the Moid of the\nresource matching the filter expression and populate it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request. An error is returned if the filter\nmatches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
			},
			"global_hot_spares": {
				Description: "A collection of disks used as hot spares globally for all the RAID groups.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"slot_number": {
							Description: "The slot number of the disk to be referenced. As this is a policy, this slot number may or may not be valid depending on the number of disks in the associated server.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
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
				Description: "Relationship to the Organization that owns the Managed Object.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The Object Type of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'\nis set and 'moid' is empty/absent from the request, Intersight will determine the Moid of the\nresource matching the filter expression and populate it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request. An error is returned if the filter\nmatches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
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
				Description: "A slice of all permission resources (organizations) associated with this object. Permission ties resources and its associated roles/privileges.\nThese resources which can be specified in a permission is PermissionResource. Currently only organizations can be specified in permission.\nAll logical and physical resources part of an organization will have organization in PermissionResources field.\nIf DeviceRegistration contains another DeviceRegistration and if parent is in org1 and child is part of org2, then child objects will\nhave PermissionResources as org1 and org2. Parent Objects will have PermissionResources as org1.\nAll profiles/policies created with in an organization will have the organization as PermissionResources.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The Object Type of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'\nis set and 'moid' is empty/absent from the request, Intersight will determine the Moid of the\nresource matching the filter expression and populate it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request. An error is returned if the filter\nmatches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
			},
			"profiles": {
				Description: "Relationship to the profile objects.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The Object Type of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'\nis set and 'moid' is empty/absent from the request, Intersight will determine the Moid of the\nresource matching the filter expression and populate it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request. An error is returned if the filter\nmatches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
			},
			"retain_policy_virtual_drives": {
				Description: "Retains the virtual drives defined in policy if they exist already. If this flag is false, the existing virtual drives are removed and created again based on virtual drives in the policy.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"tags": {
				Description: "The array of tags, which allow to add key, value meta-data to managed objects.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"key": {
							Description: "The string representation of a tag key.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"value": {
							Description: "The string representation of a tag value.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
			},
			"unused_disks_state": {
				Description: "This is used to specify the state, unconfigured good or jbod, in which the disks that are not used in this policy should be moved.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "UnconfiguredGood",
			},
			"virtual_drives": {
				Description: "The list of virtual drives and the disk groups that need to be created through this policy.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_policy": {
							Description: "Access policy that host has on this virtual drive.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "Default",
						},
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"boot_drive": {
							Description: "This flag enables the use of this virtual drive as a boot drive.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"disk_group_name": {
							Description: "Disk group policy that has the disk group in which this virtual drive needs to be created.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"disk_group_policy": {
							Description: "Disk group policy that has the disk group in which this virtual drive needs to be created.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"drive_cache": {
							Description: "This property expect disk cache policy.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "Default",
						},
						"expand_to_available": {
							Description: "This flag enables this virtual drive to use all the available space in the disk group. When this flag is configured, the size property is ignored.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"io_policy": {
							Description: "This property expects the desired IO mode - direct IO or cached IO.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "Default",
						},
						"name": {
							Description: "The name of the virtual drive. The name can be between 1 and 15 alphanumeric characters. Spaces or any special characters other than - (hyphen), _ (underscore), : (colon), and . (period) are not allowed.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"read_policy": {
							Description: "Read ahead mode to be used to read data from this virtual drive.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "Default",
						},
						"size": {
							Description: "Virtual drive size in MB. This is a required field unless the 'Expand to Available' option is enabled.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"write_policy": {
							Description: "Write mode to be used to write data to this virtual drive.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "Default",
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
			},
		},
	}
}
func resourceStorageStoragePolicyCreate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o models.StorageStoragePolicy
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.Description = x

	}

	if v, ok := d.GetOk("disk_group_policies"); ok {
		x := make([]*models.StorageDiskGroupPolicyRef, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.StorageDiskGroupPolicyRef{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["moid"]; ok {
					{
						x := (v.(string))
						o.Moid = x
					}
				}
				if v, ok := l["object_type"]; ok {
					{
						x := (v.(string))
						o.ObjectType = x
					}
				}
				if v, ok := l["selector"]; ok {
					{
						x := (v.(string))
						o.Selector = x
					}
				}
				x = append(x, &o)
			}
		}
		o.DiskGroupPolicies = x

	}

	if v, ok := d.GetOk("global_hot_spares"); ok {
		x := make([]*models.StorageLocalDisk, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.StorageLocalDisk{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["additional_properties"]; ok {
					{
						x := []byte(v.(string))
						var x1 interface{}
						err := json.Unmarshal(x, &x1)
						if err == nil && x1 != nil {
							o.StorageLocalDiskAO1P1.StorageLocalDiskAO1P1 = x1.(map[string]interface{})
						}
					}
				}
				if v, ok := l["object_type"]; ok {
					{
						x := (v.(string))
						o.ObjectType = x
					}
				}
				if v, ok := l["slot_number"]; ok {
					{
						x := int64(v.(int))
						o.SlotNumber = x
					}
				}
				x = append(x, &o)
			}
		}
		o.GlobalHotSpares = x

	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.Moid = x

	}

	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.Name = x

	}

	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.ObjectType = x

	}

	if v, ok := d.GetOk("organization"); ok {
		p := models.OrganizationOrganizationRef{}
		if len(v.([]interface{})) > 0 {
			o := models.OrganizationOrganizationRef{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.Moid = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.Selector = x
				}
			}

			p = o
		}
		x := p
		o.Organization = &x

	}

	if v, ok := d.GetOk("permission_resources"); ok {
		x := make([]*models.MoBaseMoRef, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.MoBaseMoRef{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["moid"]; ok {
					{
						x := (v.(string))
						o.Moid = x
					}
				}
				if v, ok := l["object_type"]; ok {
					{
						x := (v.(string))
						o.ObjectType = x
					}
				}
				if v, ok := l["selector"]; ok {
					{
						x := (v.(string))
						o.Selector = x
					}
				}
				x = append(x, &o)
			}
		}
		o.PermissionResources = x

	}

	if v, ok := d.GetOk("profiles"); ok {
		x := make([]*models.PolicyAbstractConfigProfileRef, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.PolicyAbstractConfigProfileRef{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["moid"]; ok {
					{
						x := (v.(string))
						o.Moid = x
					}
				}
				if v, ok := l["object_type"]; ok {
					{
						x := (v.(string))
						o.ObjectType = x
					}
				}
				if v, ok := l["selector"]; ok {
					{
						x := (v.(string))
						o.Selector = x
					}
				}
				x = append(x, &o)
			}
		}
		o.Profiles = x

	}

	if v, ok := d.GetOkExists("retain_policy_virtual_drives"); ok {
		x := v.(bool)
		o.RetainPolicyVirtualDrives = &x
	}

	if v, ok := d.GetOk("tags"); ok {
		x := make([]*models.MoTag, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.MoTag{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["additional_properties"]; ok {
					{
						x := []byte(v.(string))
						var x1 interface{}
						err := json.Unmarshal(x, &x1)
						if err == nil && x1 != nil {
							o.MoTagAO1P1.MoTagAO1P1 = x1.(map[string]interface{})
						}
					}
				}
				if v, ok := l["key"]; ok {
					{
						x := (v.(string))
						o.Key = x
					}
				}
				if v, ok := l["object_type"]; ok {
					{
						x := (v.(string))
						o.ObjectType = x
					}
				}
				if v, ok := l["value"]; ok {
					{
						x := (v.(string))
						o.Value = x
					}
				}
				x = append(x, &o)
			}
		}
		o.Tags = x

	}

	if v, ok := d.GetOk("unused_disks_state"); ok {
		x := (v.(string))
		o.UnusedDisksState = &x

	}

	if v, ok := d.GetOk("virtual_drives"); ok {
		x := make([]*models.StorageVirtualDriveConfig, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.StorageVirtualDriveConfig{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["access_policy"]; ok {
					{
						x := (v.(string))
						o.AccessPolicy = &x
					}
				}
				if v, ok := l["additional_properties"]; ok {
					{
						x := []byte(v.(string))
						var x1 interface{}
						err := json.Unmarshal(x, &x1)
						if err == nil && x1 != nil {
							o.StorageVirtualDriveConfigAO1P1.StorageVirtualDriveConfigAO1P1 = x1.(map[string]interface{})
						}
					}
				}
				if v, ok := l["boot_drive"]; ok {
					{
						x := (v.(bool))
						o.BootDrive = &x
					}
				}
				if v, ok := l["disk_group_name"]; ok {
					{
						x := (v.(string))
						o.DiskGroupName = x
					}
				}
				if v, ok := l["disk_group_policy"]; ok {
					{
						x := (v.(string))
						o.DiskGroupPolicy = x
					}
				}
				if v, ok := l["drive_cache"]; ok {
					{
						x := (v.(string))
						o.DriveCache = &x
					}
				}
				if v, ok := l["expand_to_available"]; ok {
					{
						x := (v.(bool))
						o.ExpandToAvailable = &x
					}
				}
				if v, ok := l["io_policy"]; ok {
					{
						x := (v.(string))
						o.IoPolicy = &x
					}
				}
				if v, ok := l["name"]; ok {
					{
						x := (v.(string))
						o.Name = x
					}
				}
				if v, ok := l["object_type"]; ok {
					{
						x := (v.(string))
						o.ObjectType = x
					}
				}
				if v, ok := l["read_policy"]; ok {
					{
						x := (v.(string))
						o.ReadPolicy = &x
					}
				}
				if v, ok := l["size"]; ok {
					{
						x := int64(v.(int))
						o.Size = x
					}
				}
				if v, ok := l["write_policy"]; ok {
					{
						x := (v.(string))
						o.WritePolicy = &x
					}
				}
				x = append(x, &o)
			}
		}
		o.VirtualDrives = x

	}

	url := "storage/StoragePolicies"
	data, err := o.MarshalJSON()
	if err != nil {
		log.Printf("error in marshaling model object. Error: %s", err.Error())
		return err
	}

	body, err := conn.SendRequest(url, data)
	if err != nil {
		return err
	}

	err = o.UnmarshalJSON(body)
	if err != nil {
		log.Printf("error in unmarshaling model object. Error: %s", err.Error())
		return err
	}
	log.Printf("Moid: %s", o.Moid)
	d.SetId(o.Moid)
	return resourceStorageStoragePolicyRead(d, meta)
}
func detachStorageStoragePolicyProfiles(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	url := "storage/StoragePolicies" + "/" + d.Id()
	var o models.StorageStoragePolicy

	o.Profiles = []*models.PolicyAbstractConfigProfileRef{}

	data, err := o.MarshalJSON()
	if err != nil {
		log.Printf("error in marshaling sol_policy. Error: %s", err.Error())
		return err
	}

	body, err := conn.SendUpdateRequest(url, data)
	if err != nil {
		log.Printf("error in sending update request. error %s", err.Error())
		return err
	}
	var s models.ServerProfile
	err = s.UnmarshalJSON(body)
	if err != nil {
		log.Printf("error in unmarshaling sol_policy. Error: %s", err.Error())
	}

	return err
}

func resourceStorageStoragePolicyRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "storage/StoragePolicies" + "/" + d.Id()

	body, err := conn.SendGetRequest(url, []byte(""))
	if err != nil {
		return err
	}
	var s models.StorageStoragePolicy
	err = s.UnmarshalJSON(body)
	if err != nil {
		log.Printf("error in unmarshaling model for read Error: %s", err.Error())
		return err
	}

	if err := d.Set("description", (s.Description)); err != nil {
		return err
	}

	if err := d.Set("disk_group_policies", flattenListStorageDiskGroupPolicyRef(s.DiskGroupPolicies, d)); err != nil {
		return err
	}

	if err := d.Set("global_hot_spares", flattenListStorageLocalDisk(s.GlobalHotSpares, d)); err != nil {
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

	if err := d.Set("organization", flattenMapOrganizationOrganizationRef(s.Organization, d)); err != nil {
		return err
	}

	if err := d.Set("permission_resources", flattenListMoBaseMoRef(s.PermissionResources, d)); err != nil {
		return err
	}

	if err := d.Set("profiles", flattenListPolicyAbstractConfigProfileRef(s.Profiles, d)); err != nil {
		return err
	}

	if err := d.Set("retain_policy_virtual_drives", (s.RetainPolicyVirtualDrives)); err != nil {
		return err
	}

	if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
		return err
	}

	if err := d.Set("unused_disks_state", (s.UnusedDisksState)); err != nil {
		return err
	}

	if err := d.Set("virtual_drives", flattenListStorageVirtualDriveConfig(s.VirtualDrives, d)); err != nil {
		return err
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.Moid)
	return nil
}
func resourceStorageStoragePolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o models.StorageStoragePolicy
	if d.HasChange("description") {
		v := d.Get("description")
		x := (v.(string))
		o.Description = x
	}

	if d.HasChange("disk_group_policies") {
		v := d.Get("disk_group_policies")
		x := make([]*models.StorageDiskGroupPolicyRef, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.StorageDiskGroupPolicyRef{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["moid"]; ok {
					{
						x := (v.(string))
						o.Moid = x
					}
				}
				if v, ok := l["object_type"]; ok {
					{
						x := (v.(string))
						o.ObjectType = x
					}
				}
				if v, ok := l["selector"]; ok {
					{
						x := (v.(string))
						o.Selector = x
					}
				}
				x = append(x, &o)
			}
		}
		o.DiskGroupPolicies = x
	}

	if d.HasChange("global_hot_spares") {
		v := d.Get("global_hot_spares")
		x := make([]*models.StorageLocalDisk, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.StorageLocalDisk{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["additional_properties"]; ok {
					{
						x := []byte(v.(string))
						var x1 interface{}
						err := json.Unmarshal(x, &x1)
						if err == nil && x1 != nil {
							o.StorageLocalDiskAO1P1.StorageLocalDiskAO1P1 = x1.(map[string]interface{})
						}
					}
				}
				if v, ok := l["object_type"]; ok {
					{
						x := (v.(string))
						o.ObjectType = x
					}
				}
				if v, ok := l["slot_number"]; ok {
					{
						x := int64(v.(int))
						o.SlotNumber = x
					}
				}
				x = append(x, &o)
			}
		}
		o.GlobalHotSpares = x
	}

	if d.HasChange("moid") {
		v := d.Get("moid")
		x := (v.(string))
		o.Moid = x
	}

	if d.HasChange("name") {
		v := d.Get("name")
		x := (v.(string))
		o.Name = x
	}

	if d.HasChange("object_type") {
		v := d.Get("object_type")
		x := (v.(string))
		o.ObjectType = x
	}

	if d.HasChange("organization") {
		v := d.Get("organization")
		p := models.OrganizationOrganizationRef{}
		if len(v.([]interface{})) > 0 {
			o := models.OrganizationOrganizationRef{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.Moid = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.Selector = x
				}
			}

			p = o
		}
		x := p
		o.Organization = &x
	}

	if d.HasChange("permission_resources") {
		v := d.Get("permission_resources")
		x := make([]*models.MoBaseMoRef, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.MoBaseMoRef{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["moid"]; ok {
					{
						x := (v.(string))
						o.Moid = x
					}
				}
				if v, ok := l["object_type"]; ok {
					{
						x := (v.(string))
						o.ObjectType = x
					}
				}
				if v, ok := l["selector"]; ok {
					{
						x := (v.(string))
						o.Selector = x
					}
				}
				x = append(x, &o)
			}
		}
		o.PermissionResources = x
	}

	if d.HasChange("profiles") {
		v := d.Get("profiles")
		x := make([]*models.PolicyAbstractConfigProfileRef, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.PolicyAbstractConfigProfileRef{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["moid"]; ok {
					{
						x := (v.(string))
						o.Moid = x
					}
				}
				if v, ok := l["object_type"]; ok {
					{
						x := (v.(string))
						o.ObjectType = x
					}
				}
				if v, ok := l["selector"]; ok {
					{
						x := (v.(string))
						o.Selector = x
					}
				}
				x = append(x, &o)
			}
		}
		o.Profiles = x
	}

	if d.HasChange("retain_policy_virtual_drives") {
		v := d.Get("retain_policy_virtual_drives")
		x := (v.(bool))
		o.RetainPolicyVirtualDrives = &x
	}

	if d.HasChange("tags") {
		v := d.Get("tags")
		x := make([]*models.MoTag, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.MoTag{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["additional_properties"]; ok {
					{
						x := []byte(v.(string))
						var x1 interface{}
						err := json.Unmarshal(x, &x1)
						if err == nil && x1 != nil {
							o.MoTagAO1P1.MoTagAO1P1 = x1.(map[string]interface{})
						}
					}
				}
				if v, ok := l["key"]; ok {
					{
						x := (v.(string))
						o.Key = x
					}
				}
				if v, ok := l["object_type"]; ok {
					{
						x := (v.(string))
						o.ObjectType = x
					}
				}
				if v, ok := l["value"]; ok {
					{
						x := (v.(string))
						o.Value = x
					}
				}
				x = append(x, &o)
			}
		}
		o.Tags = x
	}

	if d.HasChange("unused_disks_state") {
		v := d.Get("unused_disks_state")
		x := (v.(string))
		o.UnusedDisksState = &x
	}

	if d.HasChange("virtual_drives") {
		v := d.Get("virtual_drives")
		x := make([]*models.StorageVirtualDriveConfig, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.StorageVirtualDriveConfig{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["access_policy"]; ok {
					{
						x := (v.(string))
						o.AccessPolicy = &x
					}
				}
				if v, ok := l["additional_properties"]; ok {
					{
						x := []byte(v.(string))
						var x1 interface{}
						err := json.Unmarshal(x, &x1)
						if err == nil && x1 != nil {
							o.StorageVirtualDriveConfigAO1P1.StorageVirtualDriveConfigAO1P1 = x1.(map[string]interface{})
						}
					}
				}
				if v, ok := l["boot_drive"]; ok {
					{
						x := (v.(bool))
						o.BootDrive = &x
					}
				}
				if v, ok := l["disk_group_name"]; ok {
					{
						x := (v.(string))
						o.DiskGroupName = x
					}
				}
				if v, ok := l["disk_group_policy"]; ok {
					{
						x := (v.(string))
						o.DiskGroupPolicy = x
					}
				}
				if v, ok := l["drive_cache"]; ok {
					{
						x := (v.(string))
						o.DriveCache = &x
					}
				}
				if v, ok := l["expand_to_available"]; ok {
					{
						x := (v.(bool))
						o.ExpandToAvailable = &x
					}
				}
				if v, ok := l["io_policy"]; ok {
					{
						x := (v.(string))
						o.IoPolicy = &x
					}
				}
				if v, ok := l["name"]; ok {
					{
						x := (v.(string))
						o.Name = x
					}
				}
				if v, ok := l["object_type"]; ok {
					{
						x := (v.(string))
						o.ObjectType = x
					}
				}
				if v, ok := l["read_policy"]; ok {
					{
						x := (v.(string))
						o.ReadPolicy = &x
					}
				}
				if v, ok := l["size"]; ok {
					{
						x := int64(v.(int))
						o.Size = x
					}
				}
				if v, ok := l["write_policy"]; ok {
					{
						x := (v.(string))
						o.WritePolicy = &x
					}
				}
				x = append(x, &o)
			}
		}
		o.VirtualDrives = x
	}

	url := "storage/StoragePolicies" + "/" + d.Id()
	data, err := o.MarshalJSON()
	if err != nil {
		log.Printf("error in marshaling model object. Error: %s", err.Error())
		return err
	}

	body, err := conn.SendUpdateRequest(url, data)
	if err != nil {
		return err
	}

	err = o.UnmarshalJSON(body)
	if err != nil {
		log.Printf("error in unmarshaling model object. Error: %s", err.Error())
		return err
	}
	log.Printf("Moid: %s", o.Moid)
	d.SetId(o.Moid)
	return resourceStorageStoragePolicyRead(d, meta)
}

func resourceStorageStoragePolicyDelete(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	url := "storage/StoragePolicies" + "/" + d.Id()
	detachStorageStoragePolicyProfiles(d, meta)
	_, err := conn.SendDeleteRequest(url)
	if err != nil {
		log.Printf("error occurred while deleting: %s", err.Error())
	}
	return err
}
