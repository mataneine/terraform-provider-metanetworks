package metanetworks

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDevice() *schema.Resource {
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
			"owner_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"platform": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tags": {
				Type:     schema.TypeMap,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"aliases": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
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
		},
		Create: resourceDeviceCreate,
		Read:   resourceDeviceRead,
		Update: resourceDeviceUpdate,
		Delete: resourceDeviceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceDeviceCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	enabled := d.Get("enabled").(bool)
	ownerID := d.Get("owner_id").(string)
	platform := d.Get("platform").(string)

	networkElement := NetworkElement{
		Name:        name,
		Description: description,
		Enabled:     &enabled,
		OwnerID:     ownerID,
		Platform:    platform,
	}
	var newDevice *NetworkElement
	newDevice, err := client.CreateNetworkElement(&networkElement)
	if err != nil {
		return err
	}

	d.SetId(newDevice.ID)

	err = deviceToResource(d, newDevice)
	if err != nil {
		return err
	}
	err = client.SetNetworkElementTags(d)
	if err != nil {
		return err
	}

	return resourceDeviceRead(d, m)
}

func resourceDeviceRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	networkElement, err := client.GetNetworkElement(d.Id())
	if err != nil {
		d.SetId("")
		return nil
	}

	err = deviceToResource(d, networkElement)
	if err != nil {
		return err
	}

	return nil
}

func resourceDeviceUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	enabled := d.Get("enabled").(bool)
	ownerID := d.Get("owner_id").(string)
	platform := d.Get("platform").(string)

	networkElement := NetworkElement{
		Name:        name,
		Description: description,
		Enabled:     &enabled,
		OwnerID:     ownerID,
		Platform:    platform,
	}
	var updatedDevice *NetworkElement
	updatedDevice, err := client.UpdateNetworkElement(d.Id(), &networkElement)
	if err != nil {
		return err
	}

	err = deviceToResource(d, updatedDevice)
	if err != nil {
		return err
	}
	err = client.SetNetworkElementTags(d)
	if err != nil {
		return err
	}

	return resourceDeviceRead(d, m)
}

func resourceDeviceDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	err := client.DeleteNetworkElement(d.Id())
	if err != nil {
		return err
	}

	return nil
}

func deviceToResource(d *schema.ResourceData, m *NetworkElement) error {
	d.Set("name", m.Name)
	d.Set("description", m.Description)
	d.Set("enabled", m.Enabled)
	d.Set("owner_id", m.OwnerID)
	d.Set("platform", m.Platform)
	d.Set("aliases", m.Aliases)
	d.Set("created_at", m.CreatedAt)
	d.Set("dns_name", m.DNSName)
	d.Set("expires_at", m.ExpiresAt)
	d.Set("modified_at", m.ModifiedAt)
	d.Set("org_id", m.OrgID)

	d.SetId(m.ID)

	return nil
}
