package main

import (
	"fmt"
	"terraform-provider-metanetworks/metanetworks"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourcePolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourcePolicyCreate,
		Read:   resourcePolicyRead,
		Update: resourcePolicyUpdate,
		Delete: resourcePolicyDelete,

		Schema: map[string]*schema.Schema{
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Default:     "",
				Description: "Brief description for identification purposes",
				Optional:    true,
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The name of the Service.",
				Required:    true,
			},
			"destinations": &schema.Schema{
				Type:        schema.TypeSet,
				Description: "List of Targets that the policy allows access to",
				Elem:        &schema.Schema{Type: schema.TypeString},
				Required:    true,
			},
			"enabled": &schema.Schema{
				Type:        schema.TypeBool,
				Default:     true,
				Description: "On/Off toggle to allow traffic to pass through this MetaPort",
				Optional:    true,
			},
			"protocol_groups": &schema.Schema{
				Type:        schema.TypeSet,
				Description: "List of Port/Protocols that the policy allows access to",
				Elem:        &schema.Schema{Type: schema.TypeString},
				Required:    true,
			},
			"sources": &schema.Schema{
				Type:        schema.TypeSet,
				Description: "List of Sources that the policy allows access from",
				Elem:        &schema.Schema{Type: schema.TypeString},
				Required:    true,
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
		},
	}
}

func resourceTypeSetToStringSlice(s *schema.Set) []string {

	values_list := s.List()
	values := make([]string, len(values_list))
	for i := 0; i < len(values_list); i++ {
		values[i] = fmt.Sprint(values_list[i])
	}

	return values

}

func resourcePolicyCreate(d *schema.ResourceData, m interface{}) error {

	client := m.(*metanetworks.Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	enabled := d.Get("enabled").(bool)
	sources := resourceTypeSetToStringSlice(d.Get("sources").(*schema.Set))
	destinations := resourceTypeSetToStringSlice(d.Get("destinations").(*schema.Set))
	protocol_groups := resourceTypeSetToStringSlice(d.Get("protocol_groups").(*schema.Set))

	policy := metanetworks.Policy{
		Name:           name,
		Description:    description,
		Enabled:        enabled,
		Destinations:   destinations,
		Sources:        sources,
		ProtocolGroups: protocol_groups,
	}

	var new_policy *metanetworks.Policy
	new_policy, err := client.CreatePolicy(&policy)
	if err != nil {
		return err
	}

	d.SetId(new_policy.Id)

	err = policyToResource(d, new_policy)
	if err != nil {
		return err
	}

	return nil
}

func resourcePolicyRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*metanetworks.Client)

	var new_policy *metanetworks.Policy
	new_policy, err := client.GetPolicy(d.Id())
	if err != nil {
		return err
	}
	err = policyToResource(d, new_policy)
	if err != nil {
		return err
	}

	return nil
}

func resourcePolicyUpdate(d *schema.ResourceData, m interface{}) error {

	client := m.(*metanetworks.Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	enabled := d.Get("enabled").(bool)
	sources := resourceTypeSetToStringSlice(d.Get("sources").(*schema.Set))
	destinations := resourceTypeSetToStringSlice(d.Get("destinations").(*schema.Set))
	protocol_groups := resourceTypeSetToStringSlice(d.Get("protocol_groups").(*schema.Set))

	policy := metanetworks.Policy{
		Name:           name,
		Description:    description,
		Enabled:        enabled,
		Destinations:   destinations,
		Sources:        sources,
		ProtocolGroups: protocol_groups,
	}

	var updated_policy *metanetworks.Policy
	updated_policy, err := client.UpdatePolicy(d.Id(), &policy)
	if err != nil {
		return err
	}

	d.SetId(updated_policy.Id)

	err = policyToResource(d, updated_policy)
	if err != nil {
		return err
	}

	return nil

}

func resourcePolicyDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*metanetworks.Client)

	err := client.DeletePolicy(d.Id())
	if err != nil {
		return err
	}

	return nil
}

func policyToResource(d *schema.ResourceData, m *metanetworks.Policy) error {

	d.Set("description", m.Description)
	d.Set("name", m.Name)

	d.Set("destinations", m.Destinations)
	d.Set("enabled", m.Enabled)
	d.Set("protocol_groups", m.ProtocolGroups)
	d.Set("sources", m.Sources)

	d.Set("created_at", m.CreatedAt)
	d.Set("modified_at", m.ModifiedAt)
	d.Set("org_id", m.OrgId)

	d.SetId(m.Id)

	return nil
}
