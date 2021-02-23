package metanetworks

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVersionControls() *schema.Resource {
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
			"enabled": {
				Type:     schema.TypeBool,
				Default:  true,
				Optional: true,
			},
			"windows_policy": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
			},
			"macos_policy": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
			},
			"linux_policy": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
			},
			"apply_to_org": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"exempt_sources": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"apply_to_entities": &schema.Schema{
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
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		Create: resourceVersionControlsCreate,
		Read:   resourceVersionControlsRead,
		Update: resourceVersionControlsUpdate,
		Delete: resourceVersionControlsDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceVersionControlsCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	enabled := d.Get("enabled").(bool)
	windowspolicy := d.Get("windows_policy").(map[string]interface{})
	macospolicy := d.Get("macos_policy").(map[string]interface{})
	linuxpolicy := d.Get("linux_policy").(map[string]interface{})
	applyToOrg := d.Get("apply_to_org").(bool)
	applyToEntities := resourceTypeSetToStringSlice(d.Get("apply_to_entities").(*schema.Set))
	exemptEntities := resourceTypeSetToStringSlice(d.Get("exempt_sources").(*schema.Set))

	versionControls := VersionControls{
		Name:            name,
		Description:     description,
		Enabled:         enabled,
		ApplyToOrg:      applyToOrg,
		ExemptEntities:  exemptEntities,
		ApplyToEntities: applyToEntities,
		WindowsPolicy:   windowspolicy,
		MacOSPolicy:     macospolicy,
		LinuxPolicy:     linuxpolicy,
	}

	var newVersionControls *VersionControls
	newVersionControls, err := client.CreateVersionControls(&versionControls)
	if err != nil {
		return err
	}

	d.SetId(newVersionControls.ID)

	err = versionControlsToResource(d, newVersionControls)
	if err != nil {
		return err
	}

	return resourceVersionControlsRead(d, m)
}

func resourceVersionControlsRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	versionControls, err := client.GetVersionControls(d.Id())
	if err != nil {
		d.SetId("")
		return nil
	}

	err = versionControlsToResource(d, versionControls)
	if err != nil {
		return err
	}

	return nil
}

func resourceVersionControlsUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	enabled := d.Get("enabled").(bool)
	windowspolicy := d.Get("windows_policy").(map[string]interface{})
	macospolicy := d.Get("macos_policy").(map[string]interface{})
	linuxpolicy := d.Get("linux_policy").(map[string]interface{})
	applyToOrg := d.Get("apply_to_org").(bool)
	applyToEntities := resourceTypeSetToStringSlice(d.Get("apply_to_entities").(*schema.Set))
	exemptEntities := resourceTypeSetToStringSlice(d.Get("exempt_sources").(*schema.Set))

	versionControls := VersionControls{
		Name:            name,
		Description:     description,
		Enabled:         enabled,
		ApplyToOrg:      applyToOrg,
		ExemptEntities:  exemptEntities,
		ApplyToEntities: applyToEntities,
		WindowsPolicy:   windowspolicy,
		MacOSPolicy:     macospolicy,
		LinuxPolicy:     linuxpolicy,
	}

	var updatedVersionControls *VersionControls
	updatedVersionControls, err := client.UpdateVersionControls(d.Id(), &versionControls)
	if err != nil {
		return err
	}

	d.SetId(updatedVersionControls.ID)

	err = versionControlsToResource(d, updatedVersionControls)
	if err != nil {
		return err
	}

	return resourceVersionControlsRead(d, m)
}

func resourceVersionControlsDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	err := client.DeleteVersionControls(d.Id())
	if err != nil {
		return err
	}

	return nil
}

func versionControlsToResource(d *schema.ResourceData, m *VersionControls) error {
	d.Set("description", m.Description)
	d.Set("name", m.Name)
	d.Set("enabled", m.Enabled)
	d.Set("apply_to_org", m.ApplyToOrg)
	d.Set("exempt_entities", m.ExemptEntities)
	d.Set("linux_policy", m.LinuxPolicy)
	d.Set("macos_policy", m.MacOSPolicy)
	d.Set("windows_policy", m.WindowsPolicy)
	d.Set("apply_to_entities", m.ApplyToEntities)
	d.Set("created_at", m.CreatedAt)
	d.Set("modified_at", m.ModifiedAt)

	d.SetId(m.ID)

	return nil
}
