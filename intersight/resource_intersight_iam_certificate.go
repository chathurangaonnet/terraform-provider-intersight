package intersight

import (
	"encoding/json"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceIamCertificate() *schema.Resource {
	return &schema.Resource{
		Create: resourceIamCertificateCreate,
		Read:   resourceIamCertificateRead,
		Update: resourceIamCertificateUpdate,
		Delete: resourceIamCertificateDelete,
		Schema: map[string]*schema.Schema{
			"certificate": {
				Description: "User-input pem-encoded certificate, signed by a CAcert.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"issuer": {
							Description: "The X.509 distinguished name of the issuer of this certificate.",
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"additional_properties": {
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: SuppressDiffAdditionProps,
									},
									"common_name": {
										Description: "A required component that identifies a person or an object.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"country": {
										Description: "Identifier for the country in which the entity resides.",
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString}},
									"locality": {
										Description: "Identifier for the place where the entry resides. The locality can be a city, county, township, or other geographic region.",
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString}},
									"object_type": {
										Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"organization": {
										Description: "Identifier for the organization in which the entity resides.",
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString}},
									"organizational_unit": {
										Description: "Identifier for a unit within the organization.",
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString}},
									"state": {
										Description: "Identifier for the state or province of the entity.",
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString}},
								},
							},
							ConfigMode: schema.SchemaConfigModeAttr,
						},
						"object_type": {
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"pem_certificate": {
							Description: "The base64 encoded certificate in PEM format.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"sha256_fingerprint": {
							Description: "The computed SHA-256 fingerprint of the certificate. Equivalent to 'openssl x509 -fingerprint -sha256'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"signature_algorithm": {
							Description: "Signature algorithm, as specified in [RFC 5280](https://tools.ietf.org/html/rfc5280).",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"subject": {
							Description: "The X.509 distinguished name of the subject of this certificate.",
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"additional_properties": {
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: SuppressDiffAdditionProps,
									},
									"common_name": {
										Description: "A required component that identifies a person or an object.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"country": {
										Description: "Identifier for the country in which the entity resides.",
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString}},
									"locality": {
										Description: "Identifier for the place where the entry resides. The locality can be a city, county, township, or other geographic region.",
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString}},
									"object_type": {
										Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"organization": {
										Description: "Identifier for the organization in which the entity resides.",
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString}},
									"organizational_unit": {
										Description: "Identifier for a unit within the organization.",
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString}},
									"state": {
										Description: "Identifier for the state or province of the entity.",
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString}},
								},
							},
							ConfigMode: schema.SchemaConfigModeAttr,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
			},
			"certificate_request": {
				Description: "The certificate signing request associated with this certificate.",
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.",
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
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
			"object_type": {
				Description: "The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"permission_resources": {
				Description: "A slice of all permission resources (organizations) associated with this object. Permission ties resources and its associated roles/privileges.These resources which can be specified in a permission is PermissionResource. Currently only organizations can be specified in permission.All logical and physical resources part of an organization will have organization in PermissionResources field.If DeviceRegistration contains another DeviceRegistration and if parent is in org1 and child is part of org2, then child objects willhave PermissionResources as org1 and org2. Parent Objects will have PermissionResources as org1.All profiles/policies created with in an organization will have the organization as PermissionResources.",
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
			},
			"status": {
				Description: "Status of the certificate.",
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
						"key": {
							Description: "The string representation of a tag key.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
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
		},
	}
}
func resourceIamCertificateCreate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o models.IamCertificate
	if v, ok := d.GetOk("certificate"); ok {
		p := models.X509Certificate{}
		if len(v.([]interface{})) > 0 {
			o := models.X509Certificate{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.X509CertificateAO1P1.X509CertificateAO1P1 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["issuer"]; ok {
				{
					p := models.PkixDistinguishedName{}
					if len(v.([]interface{})) > 0 {
						o := models.PkixDistinguishedName{}
						l := (v.([]interface{})[0]).(map[string]interface{})
						if v, ok := l["additional_properties"]; ok {
							{
								x := []byte(v.(string))
								var x1 interface{}
								err := json.Unmarshal(x, &x1)
								if err == nil && x1 != nil {
									o.PkixDistinguishedNameAO1P1.PkixDistinguishedNameAO1P1 = x1.(map[string]interface{})
								}
							}
						}
						if v, ok := l["common_name"]; ok {
							{
								x := (v.(string))
								o.CommonName = x
							}
						}
						if v, ok := l["country"]; ok {
							{
								x := make([]string, 0)
								y := reflect.ValueOf(v)
								for i := 0; i < y.Len(); i++ {
									x = append(x, y.Index(i).Interface().(string))
								}
								o.Country = x
							}
						}
						if v, ok := l["locality"]; ok {
							{
								x := make([]string, 0)
								y := reflect.ValueOf(v)
								for i := 0; i < y.Len(); i++ {
									x = append(x, y.Index(i).Interface().(string))
								}
								o.Locality = x
							}
						}
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.ObjectType = x
							}
						}
						if v, ok := l["organization"]; ok {
							{
								x := make([]string, 0)
								y := reflect.ValueOf(v)
								for i := 0; i < y.Len(); i++ {
									x = append(x, y.Index(i).Interface().(string))
								}
								o.Organization = x
							}
						}
						if v, ok := l["organizational_unit"]; ok {
							{
								x := make([]string, 0)
								y := reflect.ValueOf(v)
								for i := 0; i < y.Len(); i++ {
									x = append(x, y.Index(i).Interface().(string))
								}
								o.OrganizationalUnit = x
							}
						}
						if v, ok := l["state"]; ok {
							{
								x := make([]string, 0)
								y := reflect.ValueOf(v)
								for i := 0; i < y.Len(); i++ {
									x = append(x, y.Index(i).Interface().(string))
								}
								o.State = x
							}
						}

						p = o
					}
					x := p
					o.Issuer = &x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["pem_certificate"]; ok {
				{
					x := (v.(string))
					o.PemCertificate = x
				}
			}
			if v, ok := l["sha256_fingerprint"]; ok {
				{
					x := (v.(string))
					o.Sha256Fingerprint = x
				}
			}
			if v, ok := l["signature_algorithm"]; ok {
				{
					x := (v.(string))
					o.SignatureAlgorithm = x
				}
			}
			if v, ok := l["subject"]; ok {
				{
					p := models.PkixDistinguishedName{}
					if len(v.([]interface{})) > 0 {
						o := models.PkixDistinguishedName{}
						l := (v.([]interface{})[0]).(map[string]interface{})
						if v, ok := l["additional_properties"]; ok {
							{
								x := []byte(v.(string))
								var x1 interface{}
								err := json.Unmarshal(x, &x1)
								if err == nil && x1 != nil {
									o.PkixDistinguishedNameAO1P1.PkixDistinguishedNameAO1P1 = x1.(map[string]interface{})
								}
							}
						}
						if v, ok := l["common_name"]; ok {
							{
								x := (v.(string))
								o.CommonName = x
							}
						}
						if v, ok := l["country"]; ok {
							{
								x := make([]string, 0)
								y := reflect.ValueOf(v)
								for i := 0; i < y.Len(); i++ {
									x = append(x, y.Index(i).Interface().(string))
								}
								o.Country = x
							}
						}
						if v, ok := l["locality"]; ok {
							{
								x := make([]string, 0)
								y := reflect.ValueOf(v)
								for i := 0; i < y.Len(); i++ {
									x = append(x, y.Index(i).Interface().(string))
								}
								o.Locality = x
							}
						}
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.ObjectType = x
							}
						}
						if v, ok := l["organization"]; ok {
							{
								x := make([]string, 0)
								y := reflect.ValueOf(v)
								for i := 0; i < y.Len(); i++ {
									x = append(x, y.Index(i).Interface().(string))
								}
								o.Organization = x
							}
						}
						if v, ok := l["organizational_unit"]; ok {
							{
								x := make([]string, 0)
								y := reflect.ValueOf(v)
								for i := 0; i < y.Len(); i++ {
									x = append(x, y.Index(i).Interface().(string))
								}
								o.OrganizationalUnit = x
							}
						}
						if v, ok := l["state"]; ok {
							{
								x := make([]string, 0)
								y := reflect.ValueOf(v)
								for i := 0; i < y.Len(); i++ {
									x = append(x, y.Index(i).Interface().(string))
								}
								o.State = x
							}
						}

						p = o
					}
					x := p
					o.Subject = &x
				}
			}

			p = o
		}
		x := p
		o.Certificate = &x

	}

	if v, ok := d.GetOk("certificate_request"); ok {
		p := models.IamCertificateRequestRef{}
		if len(v.([]interface{})) > 0 {
			o := models.IamCertificateRequestRef{}
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
		o.CertificateRequest = &x

	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.Moid = x

	}

	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.ObjectType = x

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

	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.Status = x

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

	url := "iam/Certificates"
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
	return resourceIamCertificateRead(d, meta)
}

func resourceIamCertificateRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "iam/Certificates" + "/" + d.Id()

	body, err := conn.SendGetRequest(url, []byte(""))
	if err != nil {
		return err
	}
	var s models.IamCertificate
	err = s.UnmarshalJSON(body)
	if err != nil {
		log.Printf("error in unmarshaling model for read Error: %s", err.Error())
		return err
	}

	if err := d.Set("certificate", flattenMapX509Certificate(s.Certificate, d)); err != nil {
		return err
	}

	if err := d.Set("certificate_request", flattenMapIamCertificateRequestRef(s.CertificateRequest, d)); err != nil {
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

	if err := d.Set("status", (s.Status)); err != nil {
		return err
	}

	if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
		return err
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.Moid)
	return nil
}
func resourceIamCertificateUpdate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o models.IamCertificate
	if d.HasChange("certificate") {
		v := d.Get("certificate")
		p := models.X509Certificate{}
		if len(v.([]interface{})) > 0 {
			o := models.X509Certificate{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.X509CertificateAO1P1.X509CertificateAO1P1 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["issuer"]; ok {
				{
					p := models.PkixDistinguishedName{}
					if len(v.([]interface{})) > 0 {
						o := models.PkixDistinguishedName{}
						l := (v.([]interface{})[0]).(map[string]interface{})
						if v, ok := l["additional_properties"]; ok {
							{
								x := []byte(v.(string))
								var x1 interface{}
								err := json.Unmarshal(x, &x1)
								if err == nil && x1 != nil {
									o.PkixDistinguishedNameAO1P1.PkixDistinguishedNameAO1P1 = x1.(map[string]interface{})
								}
							}
						}
						if v, ok := l["common_name"]; ok {
							{
								x := (v.(string))
								o.CommonName = x
							}
						}
						if v, ok := l["country"]; ok {
							{
								x := make([]string, 0)
								y := reflect.ValueOf(v)
								for i := 0; i < y.Len(); i++ {
									x = append(x, y.Index(i).Interface().(string))
								}
								o.Country = x
							}
						}
						if v, ok := l["locality"]; ok {
							{
								x := make([]string, 0)
								y := reflect.ValueOf(v)
								for i := 0; i < y.Len(); i++ {
									x = append(x, y.Index(i).Interface().(string))
								}
								o.Locality = x
							}
						}
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.ObjectType = x
							}
						}
						if v, ok := l["organization"]; ok {
							{
								x := make([]string, 0)
								y := reflect.ValueOf(v)
								for i := 0; i < y.Len(); i++ {
									x = append(x, y.Index(i).Interface().(string))
								}
								o.Organization = x
							}
						}
						if v, ok := l["organizational_unit"]; ok {
							{
								x := make([]string, 0)
								y := reflect.ValueOf(v)
								for i := 0; i < y.Len(); i++ {
									x = append(x, y.Index(i).Interface().(string))
								}
								o.OrganizationalUnit = x
							}
						}
						if v, ok := l["state"]; ok {
							{
								x := make([]string, 0)
								y := reflect.ValueOf(v)
								for i := 0; i < y.Len(); i++ {
									x = append(x, y.Index(i).Interface().(string))
								}
								o.State = x
							}
						}

						p = o
					}
					x := p
					o.Issuer = &x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["pem_certificate"]; ok {
				{
					x := (v.(string))
					o.PemCertificate = x
				}
			}
			if v, ok := l["sha256_fingerprint"]; ok {
				{
					x := (v.(string))
					o.Sha256Fingerprint = x
				}
			}
			if v, ok := l["signature_algorithm"]; ok {
				{
					x := (v.(string))
					o.SignatureAlgorithm = x
				}
			}
			if v, ok := l["subject"]; ok {
				{
					p := models.PkixDistinguishedName{}
					if len(v.([]interface{})) > 0 {
						o := models.PkixDistinguishedName{}
						l := (v.([]interface{})[0]).(map[string]interface{})
						if v, ok := l["additional_properties"]; ok {
							{
								x := []byte(v.(string))
								var x1 interface{}
								err := json.Unmarshal(x, &x1)
								if err == nil && x1 != nil {
									o.PkixDistinguishedNameAO1P1.PkixDistinguishedNameAO1P1 = x1.(map[string]interface{})
								}
							}
						}
						if v, ok := l["common_name"]; ok {
							{
								x := (v.(string))
								o.CommonName = x
							}
						}
						if v, ok := l["country"]; ok {
							{
								x := make([]string, 0)
								y := reflect.ValueOf(v)
								for i := 0; i < y.Len(); i++ {
									x = append(x, y.Index(i).Interface().(string))
								}
								o.Country = x
							}
						}
						if v, ok := l["locality"]; ok {
							{
								x := make([]string, 0)
								y := reflect.ValueOf(v)
								for i := 0; i < y.Len(); i++ {
									x = append(x, y.Index(i).Interface().(string))
								}
								o.Locality = x
							}
						}
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.ObjectType = x
							}
						}
						if v, ok := l["organization"]; ok {
							{
								x := make([]string, 0)
								y := reflect.ValueOf(v)
								for i := 0; i < y.Len(); i++ {
									x = append(x, y.Index(i).Interface().(string))
								}
								o.Organization = x
							}
						}
						if v, ok := l["organizational_unit"]; ok {
							{
								x := make([]string, 0)
								y := reflect.ValueOf(v)
								for i := 0; i < y.Len(); i++ {
									x = append(x, y.Index(i).Interface().(string))
								}
								o.OrganizationalUnit = x
							}
						}
						if v, ok := l["state"]; ok {
							{
								x := make([]string, 0)
								y := reflect.ValueOf(v)
								for i := 0; i < y.Len(); i++ {
									x = append(x, y.Index(i).Interface().(string))
								}
								o.State = x
							}
						}

						p = o
					}
					x := p
					o.Subject = &x
				}
			}

			p = o
		}
		x := p
		o.Certificate = &x
	}

	if d.HasChange("certificate_request") {
		v := d.Get("certificate_request")
		p := models.IamCertificateRequestRef{}
		if len(v.([]interface{})) > 0 {
			o := models.IamCertificateRequestRef{}
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
		o.CertificateRequest = &x
	}

	if d.HasChange("moid") {
		v := d.Get("moid")
		x := (v.(string))
		o.Moid = x
	}

	if d.HasChange("object_type") {
		v := d.Get("object_type")
		x := (v.(string))
		o.ObjectType = x
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

	if d.HasChange("status") {
		v := d.Get("status")
		x := (v.(string))
		o.Status = x
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

	url := "iam/Certificates" + "/" + d.Id()
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
	return resourceIamCertificateRead(d, meta)
}

func resourceIamCertificateDelete(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	url := "iam/Certificates" + "/" + d.Id()
	_, err := conn.SendDeleteRequest(url)
	if err != nil {
		log.Printf("error occurred while deleting: %s", err.Error())
	}
	return err
}
