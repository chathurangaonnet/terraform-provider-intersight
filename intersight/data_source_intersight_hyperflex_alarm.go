package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	models "github.com/cisco-intersight/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceHyperflexAlarm() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHyperflexAlarmRead,
		Schema: map[string]*schema.Schema{
			"acknowledged": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"acknowledged_by": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"acknowledged_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"acknowledged_time_as_utc": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"class_id": {
				Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"cluster": {
				Description: "A reference to a hyperflexCluster resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"entity_data": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"entity_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"entity_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"entity_uu_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"message": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"triggered_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"triggered_time_as_utc": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func dataSourceHyperflexAlarmRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewHyperflexAlarm()
	if v, ok := d.GetOk("acknowledged"); ok {
		x := (v.(bool))
		o.SetAcknowledged(x)
	}
	if v, ok := d.GetOk("acknowledged_by"); ok {
		x := (v.(string))
		o.SetAcknowledgedBy(x)
	}
	if v, ok := d.GetOk("acknowledged_time"); ok {
		x := int64(v.(int))
		o.SetAcknowledgedTime(x)
	}
	if v, ok := d.GetOk("acknowledged_time_as_utc"); ok {
		x := (v.(string))
		o.SetAcknowledgedTimeAsUtc(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("entity_data"); ok {
		x := (v.(string))
		o.SetEntityData(x)
	}
	if v, ok := d.GetOk("entity_name"); ok {
		x := (v.(string))
		o.SetEntityName(x)
	}
	if v, ok := d.GetOk("entity_type"); ok {
		x := (v.(string))
		o.SetEntityType(x)
	}
	if v, ok := d.GetOk("entity_uu_id"); ok {
		x := (v.(string))
		o.SetEntityUuId(x)
	}
	if v, ok := d.GetOk("message"); ok {
		x := (v.(string))
		o.SetMessage(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.SetStatus(x)
	}
	if v, ok := d.GetOk("triggered_time"); ok {
		x := int64(v.(int))
		o.SetTriggeredTime(x)
	}
	if v, ok := d.GetOk("triggered_time_as_utc"); ok {
		x := (v.(string))
		o.SetTriggeredTimeAsUtc(x)
	}
	if v, ok := d.GetOk("uuid"); ok {
		x := (v.(string))
		o.SetUuid(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return fmt.Errorf("Json Marshalling of data source failed with error : %+v", err)
	}
	result, _, err := conn.ApiClient.HyperflexApi.GetHyperflexAlarmList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if err != nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = models.NewHyperflexAlarm()
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return err
			}
			if err := d.Set("acknowledged", (s.Acknowledged)); err != nil {
				return err
			}
			if err := d.Set("acknowledged_by", (s.AcknowledgedBy)); err != nil {
				return err
			}
			if err := d.Set("acknowledged_time", (s.AcknowledgedTime)); err != nil {
				return err
			}
			if err := d.Set("acknowledged_time_as_utc", (s.AcknowledgedTimeAsUtc)); err != nil {
				return err
			}
			if err := d.Set("class_id", (s.ClassId)); err != nil {
				return err
			}

			if err := d.Set("cluster", flattenMapHyperflexClusterRelationship(s.Cluster, d)); err != nil {
				return err
			}
			if err := d.Set("description", (s.Description)); err != nil {
				return err
			}
			if err := d.Set("entity_data", (s.EntityData)); err != nil {
				return err
			}
			if err := d.Set("entity_name", (s.EntityName)); err != nil {
				return err
			}
			if err := d.Set("entity_type", (s.EntityType)); err != nil {
				return err
			}
			if err := d.Set("entity_uu_id", (s.EntityUuId)); err != nil {
				return err
			}
			if err := d.Set("message", (s.Message)); err != nil {
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
			if err := d.Set("status", (s.Status)); err != nil {
				return err
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}
			if err := d.Set("triggered_time", (s.TriggeredTime)); err != nil {
				return err
			}
			if err := d.Set("triggered_time_as_utc", (s.TriggeredTimeAsUtc)); err != nil {
				return err
			}
			if err := d.Set("uuid", (s.Uuid)); err != nil {
				return err
			}
			d.SetId(s.GetMoid())
		}
	}
	return nil
}
