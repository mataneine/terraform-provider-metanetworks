package metanetworks

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEgressRoute() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"destinations": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Default:  true,
				Optional: true,
			},
			"exempt_sources": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"sources": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"via": {
				Type:     schema.TypeString,
				Required: true,
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

	description := d.Get("description").(string)
	destinations := resourceTypeSetToStringSlice(d.Get("destinations").(*schema.Set))
	enabled := d.Get("enabled").(bool)
	exemptSources := resourceTypeSetToStringSlice(d.Get("exempt_sources").(*schema.Set))
	name := d.Get("name").(string)
	sources := resourceTypeSetToStringSlice(d.Get("sources").(*schema.Set))
	via := d.Get("via").(string)

	egressRoute := EgressRoute{
		Description:   description,
		Destinations:  destinations,
		Enabled:       enabled,
		ExemptSources: exemptSources,
		Name:          name,
		Sources:       sources,
		Via:           via,
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

	description := d.Get("description").(string)
	destinations := resourceTypeSetToStringSlice(d.Get("destinations").(*schema.Set))
	enabled := d.Get("enabled").(bool)
	exemptSources := resourceTypeSetToStringSlice(d.Get("exempt_sources").(*schema.Set))
	name := d.Get("name").(string)
	sources := resourceTypeSetToStringSlice(d.Get("sources").(*schema.Set))
	via := d.Get("via").(string)

	egressRoute := EgressRoute{
		Description:   description,
		Destinations:  destinations,
		Enabled:       enabled,
		ExemptSources: exemptSources,
		Name:          name,
		Sources:       sources,
		Via:           via,
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
	if err != nil {
		return err
	}

	return nil
}

func egressRouteToResource(d *schema.ResourceData, m *EgressRoute) error {
	d.Set("description", m.Description)
	d.Set("destinations", m.Destinations)
	d.Set("enabled", m.Enabled)
	d.Set("exempt_sources", m.ExemptSources)
	d.Set("name", m.Name)
	d.Set("sources", m.Sources)
	d.Set("via", m.Via)
	d.Set("created_at", m.CreatedAt)
	d.Set("modified_at", m.ModifiedAt)
	d.Set("org_id", m.OrgID)

	d.SetId(m.ID)

	return nil
}
