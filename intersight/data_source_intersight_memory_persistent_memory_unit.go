package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	models "github.com/cisco-intersight/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceMemoryPersistentMemoryUnit() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMemoryPersistentMemoryUnitRead,
		Schema: map[string]*schema.Schema{
			"admin_state": {
				Description: "This represents the administrative state of the memory unit on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"app_direct_capacity": {
				Description: "AppDirect capacity in GiB of the Persistent Memory Module on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"array_id": {
				Description: "This represents the memory array to which the memory unit belongs to.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"bank": {
				Description: "This represents the memory bank of the memory unit on a server.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"capacity": {
				Description: "This represents the memory capacity in MiB of the memory unit on a server.",
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
			"clock": {
				Description: "This represents the clock of the memory unit on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"count_status": {
				Description: "Count status of the Persistent Memory Module on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"device_mo_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dn": {
				Description: "The Distinguished Name unambiguously identifies an object in the system.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"firmware_version": {
				Description: "Firmware version of the firware running on the Persistent Memory Module on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"form_factor": {
				Description: "This represents the form factor of the memory unit on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"frozen_status": {
				Description: "Frozen status of the Persistent Memory Module on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"health_state": {
				Description: "Health state of the Persistent Memory Module on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"latency": {
				Description: "This represents the latency of the memory unit on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"location": {
				Description: "This represents the location of the memory unit on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"lock_status": {
				Description: "Lock status of the Persistent Memory Module on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"memory_array": {
				Description: "A reference to a memoryArray resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"memory_capacity": {
				Description: "Memory capacity in GiB of the Persistent Memory Module on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"memory_id": {
				Description: "ID of the Persistent Memory Module on a server.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"model": {
				Description: "This field identifies the model of the given component.",
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
			"oper_power_state": {
				Description: "This represents the operational power state of the memory unit on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"oper_state": {
				Description: "This represents the operational state of the memory unit on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"operability": {
				Description: "This represents the operability of the memory unit on a server.",
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
			"persistent_memory_capacity": {
				Description: "Persistent Memory capacity in GiB of the Persistent Memory Module on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"presence": {
				Description: "This represents the presence state of the memory unit on a server.",
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
			"reserved_capacity": {
				Description: "Reserved capacity in GiB of the Persistent Memory Module on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"revision": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rn": {
				Description: "The Relative Name uniquely identifies an object within a given context.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"security_status": {
				Description: "Security status of the Persistent Memory Module on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"serial": {
				Description: "This field identifies the serial of the given component.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"set": {
				Description: "This represents the set of the memory unit on a server.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"socket_id": {
				Description: "Socket ID of the Persistent Memory Module on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"socket_memory_id": {
				Description: "Socket Memory ID of the Persistent Memory Module on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"speed": {
				Description: "This represents the speed of the memory unit on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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
			"thermal": {
				Description: "This represents the thremal state of the memory unit on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"total_capacity": {
				Description: "Total capacity in GiB of the Persistent Memory Module on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"type": {
				Description: "This represents the memory type of the memory unit on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"uid": {
				Description: "UID of the Persistent Memory Module on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"vendor": {
				Description: "This field identifies the vendor of the given component.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"visibility": {
				Description: "This represents the visibility of the memory unit on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"width": {
				Description: "This represents the width of the memory unit on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func dataSourceMemoryPersistentMemoryUnitRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewMemoryPersistentMemoryUnit()
	if v, ok := d.GetOk("admin_state"); ok {
		x := (v.(string))
		o.SetAdminState(x)
	}
	if v, ok := d.GetOk("app_direct_capacity"); ok {
		x := (v.(string))
		o.SetAppDirectCapacity(x)
	}
	if v, ok := d.GetOk("array_id"); ok {
		x := int64(v.(int))
		o.SetArrayId(x)
	}
	if v, ok := d.GetOk("bank"); ok {
		x := int64(v.(int))
		o.SetBank(x)
	}
	if v, ok := d.GetOk("capacity"); ok {
		x := (v.(string))
		o.SetCapacity(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("clock"); ok {
		x := (v.(string))
		o.SetClock(x)
	}
	if v, ok := d.GetOk("count_status"); ok {
		x := (v.(string))
		o.SetCountStatus(x)
	}
	if v, ok := d.GetOk("device_mo_id"); ok {
		x := (v.(string))
		o.SetDeviceMoId(x)
	}
	if v, ok := d.GetOk("dn"); ok {
		x := (v.(string))
		o.SetDn(x)
	}
	if v, ok := d.GetOk("firmware_version"); ok {
		x := (v.(string))
		o.SetFirmwareVersion(x)
	}
	if v, ok := d.GetOk("form_factor"); ok {
		x := (v.(string))
		o.SetFormFactor(x)
	}
	if v, ok := d.GetOk("frozen_status"); ok {
		x := (v.(string))
		o.SetFrozenStatus(x)
	}
	if v, ok := d.GetOk("health_state"); ok {
		x := (v.(string))
		o.SetHealthState(x)
	}
	if v, ok := d.GetOk("latency"); ok {
		x := (v.(string))
		o.SetLatency(x)
	}
	if v, ok := d.GetOk("location"); ok {
		x := (v.(string))
		o.SetLocation(x)
	}
	if v, ok := d.GetOk("lock_status"); ok {
		x := (v.(string))
		o.SetLockStatus(x)
	}
	if v, ok := d.GetOk("memory_capacity"); ok {
		x := (v.(string))
		o.SetMemoryCapacity(x)
	}
	if v, ok := d.GetOk("memory_id"); ok {
		x := int64(v.(int))
		o.SetMemoryId(x)
	}
	if v, ok := d.GetOk("model"); ok {
		x := (v.(string))
		o.SetModel(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("oper_power_state"); ok {
		x := (v.(string))
		o.SetOperPowerState(x)
	}
	if v, ok := d.GetOk("oper_state"); ok {
		x := (v.(string))
		o.SetOperState(x)
	}
	if v, ok := d.GetOk("operability"); ok {
		x := (v.(string))
		o.SetOperability(x)
	}
	if v, ok := d.GetOk("persistent_memory_capacity"); ok {
		x := (v.(string))
		o.SetPersistentMemoryCapacity(x)
	}
	if v, ok := d.GetOk("presence"); ok {
		x := (v.(string))
		o.SetPresence(x)
	}
	if v, ok := d.GetOk("reserved_capacity"); ok {
		x := (v.(string))
		o.SetReservedCapacity(x)
	}
	if v, ok := d.GetOk("revision"); ok {
		x := (v.(string))
		o.SetRevision(x)
	}
	if v, ok := d.GetOk("rn"); ok {
		x := (v.(string))
		o.SetRn(x)
	}
	if v, ok := d.GetOk("security_status"); ok {
		x := (v.(string))
		o.SetSecurityStatus(x)
	}
	if v, ok := d.GetOk("serial"); ok {
		x := (v.(string))
		o.SetSerial(x)
	}
	if v, ok := d.GetOk("set"); ok {
		x := int64(v.(int))
		o.SetSet(x)
	}
	if v, ok := d.GetOk("socket_id"); ok {
		x := (v.(string))
		o.SetSocketId(x)
	}
	if v, ok := d.GetOk("socket_memory_id"); ok {
		x := (v.(string))
		o.SetSocketMemoryId(x)
	}
	if v, ok := d.GetOk("speed"); ok {
		x := (v.(string))
		o.SetSpeed(x)
	}
	if v, ok := d.GetOk("thermal"); ok {
		x := (v.(string))
		o.SetThermal(x)
	}
	if v, ok := d.GetOk("total_capacity"); ok {
		x := (v.(string))
		o.SetTotalCapacity(x)
	}
	if v, ok := d.GetOk("type"); ok {
		x := (v.(string))
		o.SetType(x)
	}
	if v, ok := d.GetOk("uid"); ok {
		x := (v.(string))
		o.SetUid(x)
	}
	if v, ok := d.GetOk("vendor"); ok {
		x := (v.(string))
		o.SetVendor(x)
	}
	if v, ok := d.GetOk("visibility"); ok {
		x := (v.(string))
		o.SetVisibility(x)
	}
	if v, ok := d.GetOk("width"); ok {
		x := (v.(string))
		o.SetWidth(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return fmt.Errorf("Json Marshalling of data source failed with error : %+v", err)
	}
	result, _, err := conn.ApiClient.MemoryApi.GetMemoryPersistentMemoryUnitList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if err != nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = models.NewMemoryPersistentMemoryUnit()
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return err
			}
			if err := d.Set("admin_state", (s.AdminState)); err != nil {
				return err
			}
			if err := d.Set("app_direct_capacity", (s.AppDirectCapacity)); err != nil {
				return err
			}
			if err := d.Set("array_id", (s.ArrayId)); err != nil {
				return err
			}
			if err := d.Set("bank", (s.Bank)); err != nil {
				return err
			}
			if err := d.Set("capacity", (s.Capacity)); err != nil {
				return err
			}
			if err := d.Set("class_id", (s.ClassId)); err != nil {
				return err
			}
			if err := d.Set("clock", (s.Clock)); err != nil {
				return err
			}
			if err := d.Set("count_status", (s.CountStatus)); err != nil {
				return err
			}
			if err := d.Set("device_mo_id", (s.DeviceMoId)); err != nil {
				return err
			}
			if err := d.Set("dn", (s.Dn)); err != nil {
				return err
			}
			if err := d.Set("firmware_version", (s.FirmwareVersion)); err != nil {
				return err
			}
			if err := d.Set("form_factor", (s.FormFactor)); err != nil {
				return err
			}
			if err := d.Set("frozen_status", (s.FrozenStatus)); err != nil {
				return err
			}
			if err := d.Set("health_state", (s.HealthState)); err != nil {
				return err
			}
			if err := d.Set("latency", (s.Latency)); err != nil {
				return err
			}
			if err := d.Set("location", (s.Location)); err != nil {
				return err
			}
			if err := d.Set("lock_status", (s.LockStatus)); err != nil {
				return err
			}

			if err := d.Set("memory_array", flattenMapMemoryArrayRelationship(s.MemoryArray, d)); err != nil {
				return err
			}
			if err := d.Set("memory_capacity", (s.MemoryCapacity)); err != nil {
				return err
			}
			if err := d.Set("memory_id", (s.MemoryId)); err != nil {
				return err
			}
			if err := d.Set("model", (s.Model)); err != nil {
				return err
			}
			if err := d.Set("moid", (s.Moid)); err != nil {
				return err
			}
			if err := d.Set("object_type", (s.ObjectType)); err != nil {
				return err
			}
			if err := d.Set("oper_power_state", (s.OperPowerState)); err != nil {
				return err
			}
			if err := d.Set("oper_state", (s.OperState)); err != nil {
				return err
			}
			if err := d.Set("operability", (s.Operability)); err != nil {
				return err
			}

			if err := d.Set("permission_resources", flattenListMoBaseMoRelationship(s.PermissionResources, d)); err != nil {
				return err
			}
			if err := d.Set("persistent_memory_capacity", (s.PersistentMemoryCapacity)); err != nil {
				return err
			}
			if err := d.Set("presence", (s.Presence)); err != nil {
				return err
			}

			if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRelationship(s.RegisteredDevice, d)); err != nil {
				return err
			}
			if err := d.Set("reserved_capacity", (s.ReservedCapacity)); err != nil {
				return err
			}
			if err := d.Set("revision", (s.Revision)); err != nil {
				return err
			}
			if err := d.Set("rn", (s.Rn)); err != nil {
				return err
			}
			if err := d.Set("security_status", (s.SecurityStatus)); err != nil {
				return err
			}
			if err := d.Set("serial", (s.Serial)); err != nil {
				return err
			}
			if err := d.Set("set", (s.Set)); err != nil {
				return err
			}
			if err := d.Set("socket_id", (s.SocketId)); err != nil {
				return err
			}
			if err := d.Set("socket_memory_id", (s.SocketMemoryId)); err != nil {
				return err
			}
			if err := d.Set("speed", (s.Speed)); err != nil {
				return err
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}
			if err := d.Set("thermal", (s.Thermal)); err != nil {
				return err
			}
			if err := d.Set("total_capacity", (s.TotalCapacity)); err != nil {
				return err
			}
			if err := d.Set("type", (s.Type)); err != nil {
				return err
			}
			if err := d.Set("uid", (s.Uid)); err != nil {
				return err
			}
			if err := d.Set("vendor", (s.Vendor)); err != nil {
				return err
			}
			if err := d.Set("visibility", (s.Visibility)); err != nil {
				return err
			}
			if err := d.Set("width", (s.Width)); err != nil {
				return err
			}
			d.SetId(s.GetMoid())
		}
	}
	return nil
}
