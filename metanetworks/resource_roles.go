package metanetworks

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRoles() *schema.Resource {
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
			"all_suborgs": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"suborgs_expression": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"apply_to_orgs": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"privileges": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"read_only": &schema.Schema{
				Type:     schema.TypeBool,
				Default:  true,
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
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		Create: resourceRolesCreate,
		Read:   resourceRolesRead,
		Update: resourceRolesUpdate,
		Delete: resourceRolesDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceRolesCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	allSuborgs := d.Get("all_suborgs").(bool)
	subOrgsExpression := d.Get("suborgs_expression").([]interface{})
	applyToOrgs := d.Get("apply_to_orgs").([]interface{})
	readOnly := d.Get("read_only").(bool)
	privileges := d.Get("privileges").([]interface{})

	roles := Roles{
		Name:              name,
		Description:       description,
		AllSubOrgs:        allSuborgs,
		SubOrgsExpression: subOrgsExpression,
		ApplyToOrgs:       applyToOrgs,
		ReadOnly:          readOnly,
		Privileges:        privileges,
	}

	var newRoles *Roles
	newRoles, err := client.CreateRoles(&roles)
	if err != nil {
		return err
	}

	d.SetId(newRoles.ID)

	err = rolesToResource(d, newRoles)
	if err != nil {
		return err
	}

	return resourceRolesRead(d, m)
}

func resourceRolesRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	roles, err := client.GetRoles(d.Id())
	if err != nil {
		d.SetId("")
		return nil
	}

	err = rolesToResource(d, roles)
	if err != nil {
		return err
	}

	return nil
}

func resourceRolesUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	allSuborgs := d.Get("all_suborgs").(bool)
	subOrgsExpression := d.Get("suborgs_expression").([]interface{})
	applyToOrgs := d.Get("apply_to_orgs").([]interface{})
	readOnly := d.Get("read_only").(bool)
	privileges := d.Get("privileges").([]interface{})

	roles := Roles{
		Name:              name,
		Description:       description,
		AllSubOrgs:        allSuborgs,
		SubOrgsExpression: subOrgsExpression,
		ApplyToOrgs:       applyToOrgs,
		ReadOnly:          readOnly,
		Privileges:        privileges,
	}

	var updatedRoles *Roles
	updatedRoles, err := client.UpdateRoles(d.Id(), &roles)
	if err != nil {
		return err
	}

	d.SetId(updatedRoles.ID)

	err = rolesToResource(d, updatedRoles)
	if err != nil {
		return err
	}

	return resourceRolesRead(d, m)
}

func resourceRolesDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	err := client.DeleteRoles(d.Id())
	if err != nil {
		return err
	}

	return nil
}

func rolesToResource(d *schema.ResourceData, m *Roles) error {
	d.Set("name", m.Name)
	d.Set("description", m.Description)
	d.Set("all_suborgs", m.AllSubOrgs)
	d.Set("suborgs_expression", m.SubOrgsExpression)
	d.Set("apply_on_orgs", m.ApplyToOrgs)
	d.Set("read_only", m.ReadOnly)
	d.Set("privileges", m.Privileges)
	d.Set("created_at", m.CreatedAt)
	d.Set("modified_at", m.ModifiedAt)
	d.SetId(m.ID)

	return nil
}
