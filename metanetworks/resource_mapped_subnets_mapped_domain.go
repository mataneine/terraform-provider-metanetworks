package metanetworks

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMappedSubnetsMappedDomain() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"mapped_subnets_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"enterprise_dns": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"mapped_domain": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
		Create: resourceMappedSubnetsMappedDomainCreate,
		Read:   resourceMappedSubnetsMappedDomainRead,
		Update: resourceMappedSubnetsMappedDomainUpdate,
		Delete: resourceMappedSubnetsMappedDomainDelete,
	}
}

func resourceMappedSubnetsMappedDomainSet(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	mappedSubnetsID := d.Get("mapped_subnets_id").(string)
	name := d.Get("name").(string)
	domain := d.Get("mapped_domain").(string)
	enterpriseDNS := d.Get("enterprise_dns").(bool)

	mappedDomain := MappedDomain{
		MappedDomain:  domain,
		EnterpriseDNS: enterpriseDNS,
	}
	_, err := client.SetNetworkElementMappedDomains(mappedSubnetsID, name, &mappedDomain)
	if err != nil {
		return err
	}

	d.SetId(name)

	return resourceMappedSubnetsMappedDomainRead(d, m)
}

func resourceMappedSubnetsMappedDomainCreate(d *schema.ResourceData, m interface{}) error {
	return resourceMappedSubnetsMappedDomainSet(d, m)
}

func resourceMappedSubnetsMappedDomainRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	var mappedSubnetsID string
	if v, ok := d.GetOk("mapped_subnets_id"); ok {
		mappedSubnetsID = v.(string)
	}

	mappedDomain, err := client.GetMappedDomain(mappedSubnetsID, d.Id())
	if err != nil {
		d.SetId("")
		return nil
	}

	err = mappedSubnetsMappedDomainToResource(d, mappedDomain)
	if err != nil {
		return err
	}

	return nil
}

func resourceMappedSubnetsMappedDomainUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceMappedSubnetsMappedDomainSet(d, m)
}

func resourceMappedSubnetsMappedDomainDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	mappedSubnetsID := d.Get("mapped_subnets_id").(string)
	name := d.Get("name").(string)

	err := client.DeleteNetworkElementMappedDomains(mappedSubnetsID, name)
	if err != nil {
		return err
	}

	return nil
}

func mappedSubnetsMappedDomainToResource(d *schema.ResourceData, m *MappedDomain) error {
	d.Set("name", m.Name)
	d.Set("mapped_domain", m.MappedDomain)
	d.Set("enterprise_dns", m.EnterpriseDNS)

	d.SetId(m.Name)

	return nil
}
