package metanetworks

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceUser() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"email": {
				Type:     schema.TypeString,
				Required: true,
			},
			"enabled": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"family_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"given_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"phone": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"provisioned_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"inventory": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
			"mfa_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"modified_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"org_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"overlay_mfa_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"phone_verified": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"roles": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
			"tags": {
				Type:     schema.TypeSet,
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
