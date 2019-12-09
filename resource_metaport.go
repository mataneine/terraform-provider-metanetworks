package main

import (
	"sync"
	"terraform-provider-metanetworks/metanetworks"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Force operations on the same metaport to happen sequentially
// to prevent race conditions
var resourceMetaPortMutex = map[string]*sync.Mutex{}

func metaportGetLock(id string) *sync.Mutex {

	_, ok := resourceMetaPortMutex[id]
	if !ok {
		resourceMetaPortMutex[id] = &sync.Mutex{}
	}

	resourceMetaPortMutex[id].Lock()

	return resourceMetaPortMutex[id]
}

func resourceMetaport() *schema.Resource {
	return &schema.Resource{
		Create: resourceMetaportCreate,
		Read:   resourceMetaportRead,
		Update: resourceMetaportUpdate,
		Delete: resourceMetaportDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The name of the MetaPort.",
				Required:    true,
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Default:     "",
				Description: "Brief description for identification purposes",
				Optional:    true,
			},
			"enabled": &schema.Schema{
				Type:        schema.TypeBool,
				Default:     true,
				Description: "On/Off toggle to allow traffic to pass through this MetaPort",
				Optional:    true,
			},
			"mapped_elements": &schema.Schema{
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "The list of IDs for mapped elements that should be accessed through this MetaPort.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"allow_support": &schema.Schema{
				Type:        schema.TypeBool,
				Default:     true,
				Description: "Allow metanetworks to remotely access this MetaPort for support purposes.",
				Optional:    true,
			},
			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"expires_at": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"modified_at": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"network_element_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"org_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceMetaportCreate(d *schema.ResourceData, m interface{}) error {

	client := m.(*metanetworks.Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	mapped_elements := make([]string, 0)
	enabled := d.Get("enabled").(bool)
	allow_support := d.Get("allow_support").(bool)

	meta_port := metanetworks.MetaPort{
		Name:           name,
		Description:    description,
		MappedElements: mapped_elements,
		Enabled:        enabled,
		AllowSupport:   allow_support,
	}
	var new_metaport *metanetworks.MetaPort
	new_metaport, err := client.CreateMetaPort(&meta_port)
	if err != nil {
		return err
	}

	d.SetId(new_metaport.Id)
	err = metaportToResource(d, new_metaport)
	if err != nil {
		return err
	}
	return nil
}

func resourceMetaportRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*metanetworks.Client)

	var new_metaport *metanetworks.MetaPort

	mutex := metaportGetLock(d.Id())
	defer mutex.Unlock()

	new_metaport, err := client.GetMetaPort(d.Id())
	if err != nil {
		return err
	}
	err = metaportToResource(d, new_metaport)
	if err != nil {
		return err
	}

	return nil
}

func resourceMetaportUpdate(d *schema.ResourceData, m interface{}) error {

	client := m.(*metanetworks.Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	mapped_elements_list := d.Get("mapped_elements").(*schema.Set).List()
	mapped_elements := make([]string, len(mapped_elements_list))
	for i, v := range mapped_elements {
		mapped_elements[i] = string(v[i])
	}

	enabled := d.Get("enabled").(bool)
	allow_support := d.Get("allow_support").(bool)

	meta_port := metanetworks.MetaPort{
		Name:           name,
		Description:    description,
		Enabled:        enabled,
		MappedElements: mapped_elements,
		AllowSupport:   allow_support,
	}

	mutex := metaportGetLock(d.Id())
	defer mutex.Unlock()

	var updated_metaport *metanetworks.MetaPort
	updated_metaport, err := client.UpdateMetaPort(d.Id(), &meta_port)
	if err != nil {
		return err
	}
	err = metaportToResource(d, updated_metaport)
	if err != nil {
		return err
	}

	return nil

}

func resourceMetaportDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*metanetworks.Client)

	mutex := metaportGetLock(d.Id())
	defer mutex.Unlock()

	err := client.DeleteMetaPort(d.Id())
	if err != nil {
		return err
	}

	return nil
}

func metaportToResource(d *schema.ResourceData, m *metanetworks.MetaPort) error {

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

	//d.Set("connection", m.Connection)

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
	err = d.Set("network_element_id", m.NetworkElementId)
	if err != nil {
		return err
	}
	err = d.Set("org_id", m.OrgId)
	if err != nil {
		return err
	}

	d.SetId(m.Id)

	return nil
}
