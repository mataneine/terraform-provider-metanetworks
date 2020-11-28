package metanetworks

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNativeService() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Default:  true,
				Optional: true,
			},
			"tags": {
				Type:     schema.TypeMap,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"expires_at": {
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
			"aliases": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
		},
		Create: resourceNativeServiceCreate,
		Read:   resourceNativeServiceRead,
		Update: resourceNativeServiceUpdate,
		Delete: resourceNativeServiceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceNativeServiceCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	enabled := d.Get("enabled").(bool)

	networkElement := NetworkElement{
		Name:        name,
		Description: description,
		Enabled:     &enabled,
	}
	var newNativeService *NetworkElement
	newNativeService, err := client.CreateNetworkElement(&networkElement)
	if err != nil {
		return err
	}

	d.SetId(newNativeService.ID)

	err = nativeServiceToResource(d, newNativeService)
	if err != nil {
		return err
	}
	err = client.SetNetworkElementTags(d)
	if err != nil {
		return err
	}

	return resourceNativeServiceRead(d, m)
}

func resourceNativeServiceRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	networkElement, err := client.GetNetworkElement(d.Id())
	if err != nil {
		d.SetId("")
		return nil
	}

	err = nativeServiceToResource(d, networkElement)
	if err != nil {
		return err
	}

	return nil
}

func resourceNativeServiceUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	enabled := d.Get("enabled").(bool)

	networkElement := NetworkElement{
		Name:        name,
		Description: description,
		Enabled:     &enabled,
	}
	var updatedNativeService *NetworkElement
	updatedNativeService, err := client.UpdateNetworkElement(d.Id(), &networkElement)
	if err != nil {
		return err
	}

	err = nativeServiceToResource(d, updatedNativeService)
	if err != nil {
		return err
	}
	err = client.SetNetworkElementTags(d)
	if err != nil {
		return err
	}

	return resourceNativeServiceRead(d, m)
}

func resourceNativeServiceDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	err := client.DeleteNetworkElement(d.Id())
	if err != nil {
		return err
	}

	return nil
}

func nativeServiceToResource(d *schema.ResourceData, m *NetworkElement) error {
	d.Set("name", m.Name)
	d.Set("description", m.Description)
	d.Set("enabled", m.Enabled)
	d.Set("aliases", m.Aliases)
	d.Set("created_at", m.CreatedAt)
	d.Set("dns_name", m.DNSName)
	d.Set("expires_at", m.ExpiresAt)
	d.Set("modified_at", m.ModifiedAt)
	d.Set("org_id", m.OrgID)

	d.SetId(m.ID)

	return nil
}
