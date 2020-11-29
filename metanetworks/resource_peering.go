package metanetworks

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePeering() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"egress_nat": {
				Type:     schema.TypeBool,
				Default:  true,
				Optional: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Default:  true,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"peers": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"modified_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"org_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		Create: resourcePeeringCreate,
		Read:   resourcePeeringRead,
		Update: resourcePeeringUpdate,
		Delete: resourcePeeringDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourcePeeringCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	description := d.Get("description").(string)
	egressNAT := d.Get("egress_nat").(bool)
	enabled := d.Get("enabled").(bool)
	name := d.Get("name").(string)
	peers := resourceTypeSetToStringSlice(d.Get("peers").(*schema.Set))

	peering := Peering{
		Description: description,
		EgressNAT:   egressNAT,
		Enabled:     enabled,
		Name:        name,
		Peers:       peers,
	}

	var newPeering *Peering
	newPeering, err := client.CreatePeering(&peering)
	if err != nil {
		return err
	}

	d.SetId(newPeering.ID)

	err = peeringToResource(d, newPeering)
	if err != nil {
		return err
	}

	return resourcePeeringRead(d, m)
}

func resourcePeeringRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	peering, err := client.GetPeering(d.Id())
	if err != nil {
		d.SetId("")
		return nil
	}

	err = peeringToResource(d, peering)
	if err != nil {
		return err
	}

	return nil
}

func resourcePeeringUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	description := d.Get("description").(string)
	egressNAT := d.Get("egress_nat").(bool)
	enabled := d.Get("enabled").(bool)
	name := d.Get("name").(string)
	peers := resourceTypeSetToStringSlice(d.Get("peers").(*schema.Set))

	peering := Peering{
		Description: description,
		EgressNAT:   egressNAT,
		Enabled:     enabled,
		Name:        name,
		Peers:       peers,
	}

	var updatedPeering *Peering
	updatedPeering, err := client.UpdatePeering(d.Id(), &peering)
	if err != nil {
		return err
	}

	d.SetId(updatedPeering.ID)

	err = peeringToResource(d, updatedPeering)
	if err != nil {
		return err
	}

	return resourcePeeringRead(d, m)
}

func resourcePeeringDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	err := client.DeletePeering(d.Id())
	if err != nil {
		return err
	}

	return nil
}

func peeringToResource(d *schema.ResourceData, m *Peering) error {
	d.Set("description", m.Description)
	d.Set("egress_nat", m.EgressNAT)
	d.Set("enabled", m.Enabled)
	d.Set("name", m.Name)
	d.Set("peers", m.Peers)
	d.Set("created_at", m.CreatedAt)
	d.Set("modified_at", m.ModifiedAt)
	d.Set("org_id", m.OrgID)

	d.SetId(m.ID)

	return nil
}
