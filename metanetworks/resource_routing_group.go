package metanetworks

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRoutingGroup() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mapped_elements_ids": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
			"exempt_sources": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"sources": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
		Create: resourceRoutingGroupCreate,
		Read:   resourceRoutingGroupRead,
		Update: resourceRoutingGroupUpdate,
		Delete: resourceRoutingGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceRoutingGroupCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	mappedElementsIDs := resourceTypeSetToStringSlice(d.Get("mapped_elements_ids").(*schema.Set))
	exemptSources := resourceTypeSetToStringSlice(d.Get("exempt_sources").(*schema.Set))
	sources := resourceTypeSetToStringSlice(d.Get("sources").(*schema.Set))

	routingGroup := RoutingGroup{
		Name:           name,
		Description:    description,
		MappedElements: mappedElementsIDs,
		ExemptSources:  exemptSources,
		Sources:        sources,
	}

	var newRoutingGroup *RoutingGroup
	newRoutingGroup, err := client.CreateRoutingGroup(&routingGroup)
	if err != nil {
		return err
	}

	d.SetId(newRoutingGroup.ID)

	err = routingGroupToResource(d, newRoutingGroup)
	if err != nil {
		return err
	}

	return resourceRoutingGroupRead(d, m)
}

func resourceRoutingGroupRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	var routingGroup *RoutingGroup
	routingGroup, err := client.GetRoutingGroup(d.Id())
	if err != nil {
		d.SetId("")
		return nil
	}

	err = routingGroupToResource(d, routingGroup)
	if err != nil {
		return err
	}

	return nil
}

func resourceRoutingGroupUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	mappedElementsIDs := resourceTypeSetToStringSlice(d.Get("mapped_elements_ids").(*schema.Set))
	exemptSources := resourceTypeSetToStringSlice(d.Get("exempt_sources").(*schema.Set))
	sources := resourceTypeSetToStringSlice(d.Get("sources").(*schema.Set))

	routingGroup := RoutingGroup{
		Name:           name,
		Description:    description,
		MappedElements: mappedElementsIDs,
		ExemptSources:  exemptSources,
		Sources:        sources,
	}

	var updatedRoutingGroup *RoutingGroup
	updatedRoutingGroup, err := client.UpdateRoutingGroup(d.Id(), &routingGroup)
	if err != nil {
		return err
	}

	d.SetId(updatedRoutingGroup.ID)

	err = routingGroupToResource(d, updatedRoutingGroup)
	if err != nil {
		return err
	}

	return resourceRoutingGroupRead(d, m)
}

func resourceRoutingGroupDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	err := client.DeleteRoutingGroup(d.Id())
	return err
}

func routingGroupToResource(d *schema.ResourceData, m *RoutingGroup) error {
	d.Set("name", m.Name)
	d.Set("description", m.Description)
	d.Set("mapped_elements_ids", m.MappedElements)
	d.Set("exempt_sources", m.ExemptSources)
	d.Set("sources", m.Sources)

	d.SetId(m.ID)

	return nil
}
