package metanetworks

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMappedSubnets() *schema.Resource {
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
			"tags": {
				Type:     schema.TypeMap,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"mapped_domains": {
				Type: schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enterprise_dns": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"mapped_domain": {
							Type:     schema.TypeString,
							Required: true,
						},
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
				Computed: true,
			},
			"mapped_hosts": {
				Type: schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mapped_host": {
							Type:     schema.TypeString,
							Required: true,
						},
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
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
			"mapped_subnets": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
		Create: resourceMappedSubnetsCreate,
		Read:   resourceMappedSubnetsRead,
		Update: resourceMappedSubnetsUpdate,
		Delete: resourceMappedSubnetsDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceMappedSubnetsCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	mappedSubnets := resourceTypeSetToStringSlice(d.Get("mapped_subnets").(*schema.Set))

	networkElement := NetworkElement{
		Name:          name,
		Description:   description,
		MappedSubnets: mappedSubnets,
	}
	var newMappedSubnets *NetworkElement
	newMappedSubnets, err := client.CreateNetworkElement(&networkElement)
	if err != nil {
		return err
	}

	d.SetId(newMappedSubnets.ID)

	err = mappedSubnetsToResource(d, newMappedSubnets)
	if err != nil {
		return err
	}
	err = client.SetNetworkElementTags(d)
	if err != nil {
		return err
	}

	return resourceMappedSubnetsRead(d, m)
}

func resourceMappedSubnetsRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	networkElement, err := client.GetNetworkElement(d.Id())
	if err != nil {
		d.SetId("")
		return nil
	}

	err = mappedSubnetsToResource(d, networkElement)
	if err != nil {
		return err
	}

	return nil
}

func resourceMappedSubnetsUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	mappedSubnets := resourceTypeSetToStringSlice(d.Get("mapped_subnets").(*schema.Set))

	networkElement := NetworkElement{
		Name:          name,
		Description:   description,
		MappedSubnets: mappedSubnets,
	}
	var updatedMappedSubnets *NetworkElement
	updatedMappedSubnets, err := client.UpdateNetworkElement(d.Id(), &networkElement)
	if err != nil {
		return err
	}

	err = mappedSubnetsToResource(d, updatedMappedSubnets)
	if err != nil {
		return err
	}
	err = client.SetNetworkElementTags(d)
	if err != nil {
		return err
	}

	return resourceMappedSubnetsRead(d, m)
}

func resourceMappedSubnetsDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	err := client.DeleteNetworkElement(d.Id())
	if err != nil {
		return err
	}

	return nil
}

func mappedSubnetsToResource(d *schema.ResourceData, m *NetworkElement) error {
	d.Set("name", m.Name)
	d.Set("description", m.Description)
	d.Set("mapped_subnets", m.MappedSubnets)
	err := d.Set("mapped_domains", flattenMappedDomains(m.MappedDomains))
	if err != nil {
		return err
	}
	err2 := d.Set("mapped_hosts", flattenMappedHosts(m.MappedHosts))
	if err2 != nil {
		return err2
	}
	d.Set("created_at", m.CreatedAt)
	d.Set("dns_name", m.DNSName)
	d.Set("expires_at", m.ExpiresAt)
	d.Set("modified_at", m.ModifiedAt)
	d.Set("org_id", m.OrgID)

	d.SetId(m.ID)

	return nil
}

func flattenMappedDomains(in []MappedDomain) []map[string]interface{} {
	var out = make([]map[string]interface{}, len(in), len(in))
	for i, v := range in {
		m := make(map[string]interface{})
		m["enterprise_dns"] = v.EnterpriseDNS
		m["mapped_domain"] = v.MappedDomain
		m["name"] = v.Name
		out[i] = m
	}
	return out
}

func flattenMappedHosts(in []MappedHost) []map[string]interface{} {
	var out = make([]map[string]interface{}, len(in), len(in))
	for i, v := range in {
		m := make(map[string]interface{})
		m["mapped_host"] = v.MappedHost
		m["name"] = v.Name
		out[i] = m
	}
	return out
}
