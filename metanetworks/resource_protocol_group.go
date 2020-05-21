package metanetworks

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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
						"port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
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
				Type:     schema.TypeString,
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
		p, err := protocolsFromList(v.([]interface{}), name)
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
		p, err := protocolsFromList(v.([]interface{}), name)
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
		m["port"] = v.Port
		m["proto"] = v.Protocol
		out[i] = m
	}
	log.Printf("flattenProtocols: %s", out)
	return out
}

func protocolsFromList(vs []interface{}, resourceID string) ([]Protocol, error) {
	result := make([]Protocol, 0, len(vs))
	for _, protocol := range vs {
		attr, ok := protocol.(map[string]interface{})
		if !ok {
			continue
		}

		t, err := protocolFromMap(attr, resourceID)
		if err != nil {
			return nil, err
		}

		if t != nil {
			result = append(result, *t)
		}
	}

	return result, nil
}

func protocolFromMap(attr map[string]interface{}, resourceID string) (*Protocol, error) {

	var p Protocol

	b, _ := json.Marshal(attr)
	json.Unmarshal(b, &p)

	if _, ok := attr["port"]; !ok {
		return nil, fmt.Errorf("%s: invalid protocol attributes: port missing", resourceID)
	}

	if _, ok := attr["proto"]; !ok {
		return nil, fmt.Errorf("%s: invalid protocol attributes: proto missing", resourceID)
	}

	var port int64
	var err error

	if v, ok := attr["port"].(int); ok {
		port = int64(v)
	}

	if v, ok := attr["port"].(string); ok {
		if port, err = strconv.ParseInt(v, 10, 64); err != nil {
			return nil, fmt.Errorf("%s: invalid protocol attribute: invalid value for port: %s", resourceID, v)
		}
	}

	t := &Protocol{
		Port:     port,
		Protocol: attr["proto"].(string),
	}

	return t, nil
}
