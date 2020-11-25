package metanetworks

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEgressRoute() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"via": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"exempt_sources": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"sources": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"destinations": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Default:  true,
				Optional: true,
			},
			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"modified_at": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"org_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		Create: resourceEgressRouteCreate,
		Read:   resourceEgressRouteRead,
		Update: resourceEgressRouteUpdate,
		Delete: resourceEgressRouteDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceEgressRouteCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	via := d.Get("via").(string)
	destinations := resourceTypeSetToStringSlice(d.Get("destinations").(*schema.Set))
	sources := resourceTypeSetToStringSlice(d.Get("sources").(*schema.Set))
	exemptSources := resourceTypeSetToStringSlice(d.Get("exempt_sources").(*schema.Set))
	enabled := d.Get("enabled").(bool)

	egressRoute := EgressRoute{
		Name:          name,
		Description:   description,
		Via:           via,
		Destinations:  destinations,
		Sources:       sources,
		ExemptSources: exemptSources,
		Enabled:       enabled,
	}

	var newEgressRoute *EgressRoute
	newEgressRoute, err := client.CreateEgressRoute(&egressRoute)
	if err != nil {
		return err
	}

	d.SetId(newEgressRoute.ID)

	err = egressRouteToResource(d, newEgressRoute)
	if err != nil {
		return err
	}

	return resourceEgressRouteRead(d, m)
}

func resourceEgressRouteRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	var egressRoute *EgressRoute
	egressRoute, err := client.GetEgressRoute(d.Id())
	if err != nil {
		d.SetId("")
		return nil
	}

	err = egressRouteToResource(d, egressRoute)
	if err != nil {
		return err
	}

	return nil
}

func resourceEgressRouteUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	via := d.Get("via").(string)
	destinations := resourceTypeSetToStringSlice(d.Get("destinations").(*schema.Set))
	sources := resourceTypeSetToStringSlice(d.Get("sources").(*schema.Set))
	exemptSources := resourceTypeSetToStringSlice(d.Get("exempt_sources").(*schema.Set))
	enabled := d.Get("enabled").(bool)

	egressRoute := EgressRoute{
		Name:          name,
		Description:   description,
		Via:           via,
		Destinations:  destinations,
		Sources:       sources,
		ExemptSources: exemptSources,
		Enabled:       enabled,
	}

	var updatedEgressRoute *EgressRoute
	updatedEgressRoute, err := client.UpdateEgressRoute(d.Id(), &egressRoute)
	if err != nil {
		return err
	}

	d.SetId(updatedEgressRoute.ID)

	err = egressRouteToResource(d, updatedEgressRoute)
	if err != nil {
		return err
	}

	return resourceEgressRouteRead(d, m)
}

func resourceEgressRouteDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	err := client.DeleteEgressRoute(d.Id())
	return err
}

func egressRouteToResource(d *schema.ResourceData, m *EgressRoute) error {
	d.Set("name", m.Name)
	d.Set("description", m.Description)
	d.Set("via", m.Via)
	d.Set("destinations", m.Destinations)
	d.Set("sources", m.Sources)
	d.Set("exempt_sources", m.ExemptSources)
	d.Set("enabled", m.Enabled)

	d.SetId(m.ID)

	return nil
}
