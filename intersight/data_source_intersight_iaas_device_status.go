package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceIaasDeviceStatus() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIaasDeviceStatusRead,
		Schema: map[string]*schema.Schema{
			"account_name": {
				Description: "The UCSD infra account name. Account Name is created when UCSD admin adds any new infra account (Physical/Virtual/Compute/Network) to be managed by UCSD.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"account_type": {
				Description: "The UCSD Infra Account type.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"claim_status": {
				Description: "Describes if the device is claimed in Intersight or not.",
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
			"connection_status": {
				Description: "Describes about the connection status between the UCSD and the actual end device.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"device_model": {
				Description: "Describes about the device model.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"device_vendor": {
				Description: "Describes about the device vendor/manufacturer of the device.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"device_version": {
				Description: "Describes about the current firmware version running on the device.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"guid": {
				Description: "A collection of references to the [iaas.UcsdInfo](mo://iaas.UcsdInfo) Managed Object.\nWhen this managed object is deleted, the referenced [iaas.UcsdInfo](mo://iaas.UcsdInfo) MO unsets its reference to this deleted MO.",
				Type:        schema.TypeList,
				MaxItems:    1,
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"ip_address": {
				Description: "The IPAddress of the device.",
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"pod": {
				Description: "Describes about the pod to which this device belongs to in UCSD.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"pod_type": {
				Description: "Describes about the podType of Pod to which this device belongs to in UCSD.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
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
				Computed: true,
			},
		},
	}
}
func dataSourceIaasDeviceStatusRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "iaas/DeviceStatuses"
	var o models.IaasDeviceStatus
	if v, ok := d.GetOk("account_name"); ok {
		x := (v.(string))
		o.AccountName = x
	}
	if v, ok := d.GetOk("account_type"); ok {
		x := (v.(string))
		o.AccountType = x
	}
	if v, ok := d.GetOk("claim_status"); ok {
		x := (v.(string))
		o.ClaimStatus = x
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.ClassID = x
	}
	if v, ok := d.GetOk("connection_status"); ok {
		x := (v.(string))
		o.ConnectionStatus = x
	}
	if v, ok := d.GetOk("device_model"); ok {
		x := (v.(string))
		o.DeviceModel = x
	}
	if v, ok := d.GetOk("device_vendor"); ok {
		x := (v.(string))
		o.DeviceVendor = x
	}
	if v, ok := d.GetOk("device_version"); ok {
		x := (v.(string))
		o.DeviceVersion = x
	}
	if v, ok := d.GetOk("ip_address"); ok {
		x := (v.(string))
		o.IPAddress = x
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.Moid = x
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.ObjectType = x
	}
	if v, ok := d.GetOk("pod"); ok {
		x := (v.(string))
		o.Pod = x
	}
	if v, ok := d.GetOk("pod_type"); ok {
		x := (v.(string))
		o.PodType = x
	}

	data, err := o.MarshalJSON()
	body, err := conn.SendGetRequest(url, data)
	if err != nil {
		return err
	}
	var x = make(map[string]interface{})
	if err = json.Unmarshal(body, &x); err != nil {
		return err
	}
	result := x["Results"]
	if result == nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s models.IaasDeviceStatus
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
				return err
			}
			if err := d.Set("account_name", (s.AccountName)); err != nil {
				return err
			}
			if err := d.Set("account_type", (s.AccountType)); err != nil {
				return err
			}
			if err := d.Set("claim_status", (s.ClaimStatus)); err != nil {
				return err
			}
			if err := d.Set("class_id", (s.ClassID)); err != nil {
				return err
			}
			if err := d.Set("connection_status", (s.ConnectionStatus)); err != nil {
				return err
			}
			if err := d.Set("device_model", (s.DeviceModel)); err != nil {
				return err
			}
			if err := d.Set("device_vendor", (s.DeviceVendor)); err != nil {
				return err
			}
			if err := d.Set("device_version", (s.DeviceVersion)); err != nil {
				return err
			}

			if err := d.Set("guid", flattenMapIaasUcsdInfoRef(s.GUID, d)); err != nil {
				return err
			}
			if err := d.Set("ip_address", (s.IPAddress)); err != nil {
				return err
			}
			if err := d.Set("moid", (s.Moid)); err != nil {
				return err
			}
			if err := d.Set("object_type", (s.ObjectType)); err != nil {
				return err
			}

			if err := d.Set("permission_resources", flattenListMoBaseMoRef(s.PermissionResources, d)); err != nil {
				return err
			}
			if err := d.Set("pod", (s.Pod)); err != nil {
				return err
			}
			if err := d.Set("pod_type", (s.PodType)); err != nil {
				return err
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}
			d.SetId(s.Moid)
		}
	}
	return nil
}
