package metanetworks

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceProtocolGroup() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"protocols": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"from_port": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
						"to_port": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
						"proto": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
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
			"read_only": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
		Create: resourceProtocolGroupCreate,
		Read:   resourceProtocolGroupRead,
		Update: resourceProtocolGroupUpdate,
		Delete: resourceProtocolGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceProtocolGroupCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)

	protocolGroup := ProtocolGroup{
		Name:        name,
		Description: description,
	}

	if v, ok := d.GetOk("protocols"); ok {
		p, err := expandProtocols(v.([]interface{}), name)
		if err != nil {
			return err
		}
		protocolGroup.Protocols = p
	}

	var newProtocolGroup *ProtocolGroup
	newProtocolGroup, err := client.CreateProtocolGroup(&protocolGroup)
	if err != nil {
		return err
	}

	d.SetId(newProtocolGroup.ID)

	err = protocolGroupToResource(d, newProtocolGroup)
	if err != nil {
		return err
	}

	return resourceProtocolGroupRead(d, m)
}

func resourceProtocolGroupRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	var protocolGroup *ProtocolGroup
	protocolGroup, err := client.GetProtocolGroup(d.Id())
	if err != nil {
		d.SetId("")
		return nil
	}

	err = protocolGroupToResource(d, protocolGroup)
	if err != nil {
		return err
	}

	return nil
}

func resourceProtocolGroupUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)

	protocolGroup := ProtocolGroup{
		Name:        name,
		Description: description,
	}

	if v, ok := d.GetOk("protocols"); ok {
		p, err := expandProtocols(v.([]interface{}), name)
		if err != nil {
			return err
		}
		protocolGroup.Protocols = p
	}

	var updatedProtocolGroup *ProtocolGroup
	updatedProtocolGroup, err := client.UpdateProtocolGroup(d.Id(), &protocolGroup)
	if err != nil {
		return err
	}

	d.SetId(updatedProtocolGroup.ID)

	err = protocolGroupToResource(d, updatedProtocolGroup)
	if err != nil {
		return err
	}

	return resourceProtocolGroupRead(d, m)
}

func resourceProtocolGroupDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	err := client.DeleteProtocolGroup(d.Id())
	if err != nil {
		return err
	}

	return nil
}

func protocolGroupToResource(d *schema.ResourceData, m *ProtocolGroup) error {
	d.Set("description", m.Description)
	d.Set("name", m.Name)
	err := d.Set("protocols", flattenProtocols(m.Protocols))
	if err != nil {
		return err
	}
	d.Set("created_at", m.CreatedAt)
	d.Set("modified_at", m.ModifiedAt)
	d.Set("org_id", m.OrgID)
	d.Set("read_only", m.ReadOnly)

	d.SetId(m.ID)

	return nil
}

func flattenProtocols(in []Protocol) []map[string]interface{} {
	var out = make([]map[string]interface{}, len(in), len(in))
	for i, v := range in {
		m := make(map[string]interface{})
		m["from_port"] = v.FromPort
		m["to_port"] = v.ToPort
		m["proto"] = v.Protocol
		out[i] = m
	}
	log.Printf("flattenProtocols: %s", out)
	return out
}

func expandProtocols(data []interface{}, resourceID string) ([]Protocol, error) {
	protocols := make([]Protocol, 0, len(data))
	for _, d := range data {
		m, ok := d.(map[string]interface{})
		if !ok {
			continue
		}

		protocol := &Protocol{
			FromPort: int64(m["from_port"].(int)),
			ToPort:   int64(m["to_port"].(int)),
			Protocol: m["proto"].(string),
		}

		protocols = append(protocols, *protocol)
	}

	return protocols, nil
}
