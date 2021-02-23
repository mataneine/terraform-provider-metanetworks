package metanetworks

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMappedSubnetsMappedHost() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"mapped_subnets_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ignore_bounds": {
				Type:     schema.TypeBool,
				Default:  false,
				Optional: true,
			},
			"mapped_host": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
		Create: resourceMappedSubnetsMappedHostCreate,
		Read:   resourceMappedSubnetsMappedHostRead,
		Update: resourceMappedSubnetsMappedHostUpdate,
		Delete: resourceMappedSubnetsMappedHostDelete,
	}
}

func resourceMappedSubnetsMappedHostSet(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	mappedSubnetsID := d.Get("mapped_subnets_id").(string)
	name := d.Get("name").(string)
	host := d.Get("mapped_host").(string)
	ignoreBounds := d.Get("ignore_bounds").(bool)

	mappedHost := MappedHost{
		MappedHost:   host,
		IgnoreBounds: ignoreBounds,
	}
	_, err := client.SetNetworkElementMappedHosts(mappedSubnetsID, name, &mappedHost)
	if err != nil {
		return err
	}

	d.SetId(name)

	return resourceMappedSubnetsMappedHostRead(d, m)
}

func resourceMappedSubnetsMappedHostCreate(d *schema.ResourceData, m interface{}) error {
	return resourceMappedSubnetsMappedHostSet(d, m)
}

func resourceMappedSubnetsMappedHostRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	var mappedSubnetsID string
	if v, ok := d.GetOk("mapped_subnets_id"); ok {
		mappedSubnetsID = v.(string)
	}

	mappedHost, err := client.GetMappedHost(mappedSubnetsID, d.Id())
	if err != nil {
		d.SetId("")
		return nil
	}

	err = mappedSubnetsMappedHostToResource(d, mappedHost)
	if err != nil {
		return err
	}

	return nil
}

func resourceMappedSubnetsMappedHostUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceMappedSubnetsMappedHostSet(d, m)
}

func resourceMappedSubnetsMappedHostDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	mappedSubnetsID := d.Get("mapped_subnets_id").(string)
	name := d.Get("name").(string)

	err := client.DeleteNetworkElementMappedHosts(mappedSubnetsID, name)
	if err != nil {
		return err
	}

	return nil
}

func mappedSubnetsMappedHostToResource(d *schema.ResourceData, m *MappedHost) error {
	d.Set("name", m.Name)
	d.Set("mapped_host", m.MappedHost)
	d.Set("ignore_bounds", m.IgnoreBounds)

	d.SetId(m.Name)

	return nil
}
