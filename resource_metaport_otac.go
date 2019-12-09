package main

import (
	"terraform-provider-metanetworks/metanetworks"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceMetaportOTAC() *schema.Resource {
	return &schema.Resource{
		Create: resourceMetaportOTACCreate,
		Read:   resourceMetaportOTACRead,
		Delete: resourceMetaportOTACDelete,

		Schema: map[string]*schema.Schema{
			"metaport_id": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The ID of the MetaPort",
				Required:    true,
				ForceNew:    true,
			},
			"triggers": &schema.Schema{
				Type:        schema.TypeSet,
				Required:    true,
				Description: "The list of IDs for mapped elements that should be accessed through this MetaPort.",
				Elem:        &schema.Schema{Type: schema.TypeString},
				ForceNew:    true,
			},
			"secret": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Brief description for identification purposes",
				Computed:    true,
				Sensitive:   true,
			},
		},
	}
}
func resourceMetaportOTACCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*metanetworks.Client)

	metaport_id := d.Get("metaport_id").(string)
	otac_secret, err := client.GenerateMetaPortOTAC(metaport_id)
	if err != nil {
		return err
	}

	d.Set("secret", otac_secret)
	d.SetId(otac_secret[0:5])

	return nil
}

func resourceMetaportOTACRead(d *schema.ResourceData, m interface{}) error {
	// fire and forget. The code is valid for a short time, there is no state.
	return nil
}

func resourceMetaportOTACDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
