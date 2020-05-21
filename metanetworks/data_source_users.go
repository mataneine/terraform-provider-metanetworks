package metanetworks

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceUser() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"family_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"given_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"phone": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"provisioned_by": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"inventory": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
			"mfa_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"modified_at": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"org_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"overlay_mfa_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"phone_verified": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"roles": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
			"tags": &schema.Schema{
				Type:     schema.TypeMap,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
		},
		Read: dataSourceUserRead,
	}
}

func dataSourceUserRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	email := d.Get("email").(string)

	var user []User
	user, err := client.GetUsers(email)
	if err != nil {
		return err
	}
	err = userToResource(d, &user[0])
	if err != nil {
		return err
	}

	return nil
}
