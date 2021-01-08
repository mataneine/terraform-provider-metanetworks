package metanetworks

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGroup() *schema.Resource {
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
			"expression": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"provisioned_by": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
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
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
			"roles": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"users": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
		},
		Create: resourceGroupCreate,
		Read:   resourceGroupRead,
		Update: resourceGroupUpdate,
		Delete: resourceGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceGroupCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	expression := d.Get("expression").(string)

	group := Group{
		Name:        name,
		Description: description,
		Expression:  expression,
	}

	var newGroup *Group
	newGroup, err := client.CreateGroup(&group)
	if err != nil {
		return err
	}

	d.SetId(newGroup.ID)
	err = groupToResource(d, newGroup)
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

	return resourceGroupRead(d, m)
}

func resourceGroupRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	group, err := client.GetGroup(d.Id())
	if err != nil {
		d.SetId("")
		return nil
	}

	err = groupToResource(d, group)
	if err != nil {
		return err
	}

	return nil
}

func resourceGroupUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	expression := d.Get("expression").(string)

	group := Group{
		Name:        name,
		Description: description,
		Expression:  expression,
	}

	var updatedGroup *Group
	updatedGroup, err := client.UpdateGroup(d.Id(), &group)
	if err != nil {
		return err
	}

	d.SetId(updatedGroup.ID)

	err = groupToResource(d, updatedGroup)
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

	return resourceGroupRead(d, m)
}

func resourceGroupDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	err := client.DeleteGroup(d.Id())
	if err != nil {
		return err
	}

	return nil
}

func setGroupRoles(d *schema.ResourceData, client *Client) error {
	if d.HasChange("roles") {
		roles := resourceTypeSetToStringSlice(d.Get("roles").(*schema.Set))
		group, err := client.SetGroupRoles(d.Id(), roles)
		if err != nil {
			return err
		}

		groupToResource(d, group)
	}

	return nil
}

func setGroupUsers(d *schema.ResourceData, client *Client) error {
	if d.HasChange("users") {
		old, new := d.GetChange("users")
		toAddSet := new.(*schema.Set).Difference(old.(*schema.Set))
		toRemoveSet := old.(*schema.Set).Difference(new.(*schema.Set))

		toAdd := resourceTypeSetToStringSlice(toAddSet)
		toRemove := resourceTypeSetToStringSlice(toRemoveSet)

		var group *Group
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

		groupToResource(d, group)
	}

	return nil
}
