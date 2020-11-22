package metanetworks

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMetaport() *schema.Resource {
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
			"mapped_elements": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
			"allow_support": {
				Type:     schema.TypeBool,
				Default:  true,
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
		},
		Create: resourceMetaportCreate,
		Read:   resourceMetaportRead,
		Update: resourceMetaportUpdate,
		Delete: resourceMetaportDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceMetaportCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	mappedElements := make([]string, 0)
	enabled := d.Get("enabled").(bool)
	allowSupport := d.Get("allow_support").(bool)

	metaport := MetaPort{
		Name:           name,
		Description:    description,
		MappedElements: mappedElements,
		Enabled:        enabled,
		AllowSupport:   allowSupport,
	}
	var newMetaport *MetaPort
	newMetaport, err := client.CreateMetaPort(&metaport)
	if err != nil {
		return err
	}

	d.SetId(newMetaport.ID)
	err = metaportToResource(d, newMetaport)
	if err != nil {
		return err
	}
	return resourceMetaportRead(d, m)
}

func resourceMetaportRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	metaport, err := client.GetMetaPort(d.Id())
	if err != nil {
		d.SetId("")
		return nil
	}

	err = metaportToResource(d, metaport)
	if err != nil {
		return err
	}

	return nil
}

func resourceMetaportUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	mappedElementsList := d.Get("mapped_elements").(*schema.Set).List()
	mappedElements := make([]string, len(mappedElementsList))
	for i, v := range mappedElements {
		mappedElements[i] = string(v[i])
	}

	enabled := d.Get("enabled").(bool)
	allowSupport := d.Get("allow_support").(bool)

	metaport := MetaPort{
		Name:           name,
		Description:    description,
		Enabled:        enabled,
		MappedElements: mappedElements,
		AllowSupport:   allowSupport,
	}

	var updatedMetaport *MetaPort
	updatedMetaport, err := client.UpdateMetaPort(d.Id(), &metaport)
	if err != nil {
		return err
	}
	err = metaportToResource(d, updatedMetaport)
	if err != nil {
		return err
	}

	return resourceMetaportRead(d, m)
}

func resourceMetaportDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	err := client.DeleteMetaPort(d.Id())
	if err != nil {
		return err
	}

	return nil
}

func metaportToResource(d *schema.ResourceData, m *MetaPort) error {
	err := d.Set("name", m.Name)
	if err != nil {
		return err
	}

	err = d.Set("description", m.Description)
	if err != nil {
		return err
	}

	err = d.Set("enabled", m.Enabled)
	if err != nil {
		return err
	}

	err = d.Set("mapped_elements", m.MappedElements)
	if err != nil {
		return err
	}

	err = d.Set("allow_support", m.AllowSupport)
	if err != nil {
		return err
	}

	err = d.Set("created_at", m.CreatedAt)
	if err != nil {
		return err
	}

	err = d.Set("dns_name", m.DNSName)
	if err != nil {
		return err
	}

	err = d.Set("expires_at", m.ExpiresAt)
	if err != nil {
		return err
	}

	err = d.Set("modified_at", m.ModifiedAt)
	if err != nil {
		return err
	}

	err = d.Set("org_id", m.OrgID)
	if err != nil {
		return err
	}

	d.SetId(m.ID)

	return nil
}
