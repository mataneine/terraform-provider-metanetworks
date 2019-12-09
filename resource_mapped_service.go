package main

import (
	"terraform-provider-metanetworks/metanetworks"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceMappedService() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The name of the Service.",
				Required:    true,
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Default:     "",
				Description: "Brief description for identification purposes",
				Optional:    true,
			},
			"tags": &schema.Schema{
				Type: schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
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
			"org_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mapped_service": &schema.Schema{
				Type:        schema.TypeString,
				Description: "DNS Name or IP Address of the target service",
				Required:    true,
			},
			"aliases": &schema.Schema{
				Type:        schema.TypeSet,
				Description: "List of aliases that resolve to the mapped service",
				Elem:        &schema.Schema{Type: schema.TypeString},
				Computed:    true,
			},
			"net_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		Create: resourceMappedServiceCreate,
		Read:   resourceMappedServiceRead,
		Update: resourceMappedServiceUpdate,
		Delete: resourceMappedServiceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceMappedServiceCreate(d *schema.ResourceData, m interface{}) error {

	client := m.(*metanetworks.Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	mappedService := d.Get("mapped_service").(string)

	networkElement := metanetworks.NetworkElement{
		Name:          name,
		Description:   description,
		MappedService: mappedService,
	}
	var newMappedService *metanetworks.NetworkElement
	newMappedService, err := client.CreateNetworkElement(&networkElement)
	if err != nil {
		return err
	}

	d.SetId(newMappedService.Id)

	err = networkElementToResource(d, newMappedService)
	if err != nil {
		return err
	}
	err = setTags(d, client)
	if err != nil {
		return err
	}

	return nil
}

func resourceMappedServiceRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*metanetworks.Client)

	var newNetworkElement *metanetworks.NetworkElement
	newNetworkElement, err := client.GetNetworkElement(d.Id())
	if err != nil {
		return err
	}
	err = networkElementToResource(d, newNetworkElement)
	if err != nil {
		return err
	}
	err = getTags(d, client)
	if err != nil {
		return err
	}

	return nil
}

func resourceMappedServiceUpdate(d *schema.ResourceData, m interface{}) error {

	client := m.(*metanetworks.Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	mappedService := d.Get("mapped_service").(string)

	networkElement := metanetworks.NetworkElement{
		Name:          name,
		Description:   description,
		MappedService: mappedService,
	}
	var updatedMappedService *metanetworks.NetworkElement
	updatedMappedService, err := client.UpdateNetworkElement(d.Id(), &networkElement)
	if err != nil {
		return err
	}

	err = networkElementToResource(d, updatedMappedService)
	if err != nil {
		return err
	}
	err = setTags(d, client)
	if err != nil {
		return err
	}

	return nil

}

func resourceMappedServiceDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*metanetworks.Client)

	err := client.DeleteNetworkElement(d.Id())
	if err != nil {
		return err
	}

	return nil
}

func networkElementToResource(d *schema.ResourceData, m *metanetworks.NetworkElement) error {

	d.Set("name", m.Name)
	d.Set("description", m.Description)
	d.Set("mapped_service", m.MappedService)
	d.Set("aliases", m.Aliases)

	d.Set("created_at", m.CreatedAt)
	d.Set("dns_name", m.DNSName)
	d.Set("expires_at", m.ExpiresAt)
	d.Set("modified_at", m.ModifiedAt)
	d.Set("org_id", m.OrgId)

	d.SetId(m.Id)

	return nil
}

func setTags(d *schema.ResourceData, client *metanetworks.Client) error {
	if d.HasChange("tags") {
		tagMapInterface := d.Get("tags").(map[string]interface{})
		tagMapString := make(map[string]string)
		for key, value := range tagMapInterface {
			tagMapString[key] = value.(string)
		}

		id := d.Id()
		err := client.SetNetworkElementTags(id, tagMapString)
		if err != nil {
			return err
		}
	}

	return nil
}

func getTags(d *schema.ResourceData, client *metanetworks.Client) error {

	id := d.Id()

	tagsMap, err := client.GetNetworkElementTags(id)
	if err != nil {
		return err
	}

	d.Set("tags", tagsMap)

	return nil
}
