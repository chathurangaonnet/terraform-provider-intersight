package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceIamAppRegistration() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIamAppRegistrationRead,
		Schema: map[string]*schema.Schema{
			"account": {
				Description: "A collection of references to the [iam.Account](mo://iam.Account) Managed Object.When this managed object is deleted, the referenced [iam.Account](mo://iam.Account) MO unsets its reference to this deleted MO.",
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"client_id": {
				Description: "A unique identifier for the OAuth2 client application.The client ID is auto-generated when the AppRegistration object is created.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"client_name": {
				Description: "App Registration name specified by user.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"client_secret": {
				Description: "The OAuth2 client secret.The value of this property is generated when grantType includes 'client-credentials'.Otherwise, no client-secret is generated.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"client_type": {
				Description: "The type of the OAuth2 client (public or confidential), as specified in https://tools.ietf.org/html/rfc6749#section-2.1.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": {
				Description: "Description of the application.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"grant_types": {
				Description: "The set of grant types that OAuth2 clients can use for this application.The grant type is used in the OAuth2 login flow to validate the grant type that has been requested by the client.See https://tools.ietf.org/html/rfc7591#page-9 for more details.# It is set automatically when AppRegistration is created since currently we do not provide option for the user.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString}},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"oauth_tokens": {
				Description: "Collection of the OAuth2 tokens. Each OAuth2 token represents valid OAuth session.OAuth2 token is created when login over OAuth2 is performed using Authorization Code grant type.",
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
			},
			"object_type": {
				Description: "The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"permission": {
				Description: "Permission associated with OAuth token issued through Client Credentials flow. Permission of the current session will be used.",
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
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
			},
			"redirect_uris": {
				Description: "After a user successfully authorizes an application, the OAuth2 authorization server will redirect the user back to theapplication with either an authorization code or access token in the URL.Registered redirect URLs may contain query string parameters, but must not contain anything in the fragment.The registration server rejects the request if a user tries to register a redirect URL that contains a fragment.For native and mobile apps, Intersight allows a user to register a URL scheme such as myapp:// which can then be usedin the redirect URL. The authorization server allows arbitrary URL schemes to be registered in order to supportregistering redirect URLs for native apps.Redirect URLs are a critical part of the OAuth flow. After a user successfully authorizes an application,the authorization server will redirect the user back to the application with either an authorization code or accesstoken in the URL. Because the redirect URL will contain sensitive information, it is critical that the servicedoesn’t redirect the user to arbitrary locations.The best way to ensure the user will only be directed to appropriate locations is to require the developer toregister one or more redirect URLs when they create the application.The redirection endpoint URI MUST be an absolute URI.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString}},
			"renew_client_secret": {
				Description: "Set value to true to renew the client-secret. Applicable to client_credentials grant type.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"response_types": {
				Description: "The set of response types that a OAuth2 client can use.This is static list and it is set automatically when AppRegistration is created.According to RFC, it is used in OAuth2 login flow to check that this AppRegistration supports response type from the request.See https://tools.ietf.org/html/rfc7591#page-9 for more details.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString}},
			"revocation_timestamp": {
				Description: "Used to perform revocation for tokens of AppRegistration.Updated only internally is case Revoke property come from UI with value true.On each request with OAuth2 access token the CreationTime of the OAuth2 token will be compared to RevokationTimestamp of thecorresponding App Registration.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"revoke": {
				Description: "Used to trigger update the revocationTimestamp value.If UI sent updating request with the Revoke value is true, then update RevocationTimestamp.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"roles": {
				Description: "The set of roles that can be used when a OAuth2 client is accessing this registered application.For example, multiple roles may be defined in your Intersight account, but you want users to loginwith the 'Read-Only' role when accessing Intersight through a registered application.In that case, the 'roles' property should contain a single element referencing the 'Read-Only' role.A user can only assign roles they already have.This relationship is deprecated. Authorization is now performed by passing the 'scope' query parameterin the first request of the Authorization Code OAuth2 flow.The value of the 'scope' parameter is a list of scope names separated by comma:ROLE.Account Administrator, ROLE.<any role name>.",
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
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
				Computed: true,
			},
			"user": {
				Description: "A collection of references to the [iam.User](mo://iam.User) Managed Object.When this managed object is deleted, the referenced [iam.User](mo://iam.User) MO unsets its reference to this deleted MO.",
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}
func dataSourceIamAppRegistrationRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "iam/AppRegistrations"
	var o models.IamAppRegistration
	if v, ok := d.GetOk("client_id"); ok {
		x := (v.(string))
		o.ClientID = x
	}
	if v, ok := d.GetOk("client_name"); ok {
		x := (v.(string))
		o.ClientName = x
	}
	if v, ok := d.GetOk("client_secret"); ok {
		x := (v.(string))
		o.ClientSecret = x
	}
	if v, ok := d.GetOk("client_type"); ok {
		x := (v.(string))
		o.ClientType = &x
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.Description = x
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.Moid = x
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.ObjectType = x
	}
	if v, ok := d.GetOk("renew_client_secret"); ok {
		x := (v.(bool))
		o.RenewClientSecret = &x
	}
	if v, ok := d.GetOk("revocation_timestamp"); ok {
		x, _ := strfmt.ParseDateTime(v.(string))
		o.RevocationTimestamp = x
	}
	if v, ok := d.GetOk("revoke"); ok {
		x := (v.(bool))
		o.Revoke = &x
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
			var s models.IamAppRegistration
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
				return err
			}

			if err := d.Set("account", flattenMapIamAccountRef(s.Account, d)); err != nil {
				return err
			}
			if err := d.Set("client_id", (s.ClientID)); err != nil {
				return err
			}
			if err := d.Set("client_name", (s.ClientName)); err != nil {
				return err
			}
			if err := d.Set("client_secret", (s.ClientSecret)); err != nil {
				return err
			}
			if err := d.Set("client_type", (s.ClientType)); err != nil {
				return err
			}
			if err := d.Set("description", (s.Description)); err != nil {
				return err
			}
			if err := d.Set("grant_types", (s.GrantTypes)); err != nil {
				return err
			}
			if err := d.Set("moid", (s.Moid)); err != nil {
				return err
			}

			if err := d.Set("oauth_tokens", flattenListIamOAuthTokenRef(s.OauthTokens, d)); err != nil {
				return err
			}
			if err := d.Set("object_type", (s.ObjectType)); err != nil {
				return err
			}

			if err := d.Set("permission", flattenMapIamPermissionRef(s.Permission, d)); err != nil {
				return err
			}

			if err := d.Set("permission_resources", flattenListMoBaseMoRef(s.PermissionResources, d)); err != nil {
				return err
			}
			if err := d.Set("redirect_uris", (s.RedirectUris)); err != nil {
				return err
			}
			if err := d.Set("renew_client_secret", (s.RenewClientSecret)); err != nil {
				return err
			}
			if err := d.Set("response_types", (s.ResponseTypes)); err != nil {
				return err
			}

			if err := d.Set("revocation_timestamp", (s.RevocationTimestamp).String()); err != nil {
				return err
			}
			if err := d.Set("revoke", (s.Revoke)); err != nil {
				return err
			}

			if err := d.Set("roles", flattenListIamRoleRef(s.Roles, d)); err != nil {
				return err
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}

			if err := d.Set("user", flattenMapIamUserRef(s.User, d)); err != nil {
				return err
			}
			d.SetId(s.Moid)
		}
	}
	return nil
}
