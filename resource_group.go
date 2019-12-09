package main

import (
	"terraform-provider-metanetworks/metanetworks"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceGroupCreate,
		Read:   resourceGroupRead,
		Update: resourceGroupUpdate,
		Delete: resourceGroupDelete,

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
			"expression": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The expression if this is a smart group",
				Optional:    true,
			},
			"provisioned_by": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The name of the Service.",
				Computed:    true,
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
			"members": &schema.Schema{
				Type:        schema.TypeSet,
				Description: "Members of the group. This list is generated from the expression",
				Elem:        &schema.Schema{Type: schema.TypeString},
				Computed:    true,
			},
			"roles": &schema.Schema{
				Type:        schema.TypeSet,
				Description: "Roles (permissions) to attach to the group",
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
			},
			"users": &schema.Schema{
				Type:        schema.TypeSet,
				Description: "Users to add to the group",
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
			},
		},
	}
}

func resourceGroupCreate(d *schema.ResourceData, m interface{}) error {

	client := m.(*metanetworks.Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	expression := d.Get("expression").(string)

	group := metanetworks.Group{
		Name:        name,
		Description: description,
		Expression:  expression,
	}

	var new_group *metanetworks.Group
	new_group, err := client.CreateGroup(&group)
	if err != nil {
		return err
	}

	d.SetId(new_group.Id)
	err = GroupToResource(d, new_group)
	if err != nil {
		return err
	}

	err = setGroupRoles(d, client)
	if err != nil {
		return err
	}

	err = setGroupUsers(d, client)
	if err != nil {
		return err
	}

	return nil
}

func resourceGroupRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*metanetworks.Client)

	var new_Group *metanetworks.Group
	new_Group, err := client.GetGroup(d.Id())
	if err != nil {
		return err
	}
	err = GroupToResource(d, new_Group)
	if err != nil {
		return err
	}

	return nil
}

func resourceGroupUpdate(d *schema.ResourceData, m interface{}) error {

	client := m.(*metanetworks.Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	expression := d.Get("expression").(string)

	group := metanetworks.Group{
		Name:        name,
		Description: description,
		Expression:  expression,
	}

	var updated_group *metanetworks.Group
	updated_group, err := client.UpdateGroup(d.Id(), &group)
	if err != nil {
		return err
	}

	d.SetId(updated_group.Id)

	err = GroupToResource(d, updated_group)
	if err != nil {
		return err
	}

	err = setGroupRoles(d, client)
	if err != nil {
		return err
	}

	err = setGroupUsers(d, client)
	if err != nil {
		return err
	}

	return nil

}

func resourceGroupDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*metanetworks.Client)

	err := client.DeleteGroup(d.Id())
	if err != nil {
		return err
	}

	return nil
}

func GroupToResource(d *schema.ResourceData, m *metanetworks.Group) error {

	d.Set("description", m.Description)
	d.Set("name", m.Name)

	d.Set("destinations", m.Description)
	d.Set("enabled", m.Description)
	d.Set("protocol_groups", m.Description)
	d.Set("sources", m.Description)

	d.Set("created_at", m.CreatedAt)
	d.Set("modified_at", m.ModifiedAt)
	d.Set("org_id", m.OrgId)

	d.SetId(m.Id)

	return nil
}

func setGroupRoles(d *schema.ResourceData, client *metanetworks.Client) error {
	if d.HasChange("roles") {
		roles := resourceTypeSetToStringSlice(d.Get("roles").(*schema.Set))
		group, err := client.SetGroupRoles(d.Id(), roles)
		if err != nil {
			return err
		}

		GroupToResource(d, group)
	}

	return nil
}

func setGroupUsers(d *schema.ResourceData, client *metanetworks.Client) error {
	if d.HasChange("users") {
		old, new := d.GetChange("users")
		toAddSet := new.(*schema.Set).Difference(old.(*schema.Set))
		toRemoveSet := old.(*schema.Set).Difference(new.(*schema.Set))

		toAdd := resourceTypeSetToStringSlice(toAddSet)
		toRemove := resourceTypeSetToStringSlice(toRemoveSet)

		var group *metanetworks.Group
		var err error
		if len(toAdd) > 0 {
			group, err = client.AddGroupUsers(d.Id(), toAdd)
			if err != nil {
				return err
			}
		}
		if len(toRemove) > 0 {
			group, err = client.RemoveGroupUsers(d.Id(), toAdd)
			if err != nil {
				return err
			}
		}

		GroupToResource(d, group)
	}

	return nil
}
