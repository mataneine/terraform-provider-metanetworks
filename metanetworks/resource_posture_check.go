package metanetworks

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePostureCheck() *schema.Resource {
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
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"osquery": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"platform": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_message_on_fail": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"check": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"min_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
				Optional: true,
			},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Default:  true,
				Optional: true,
			},
			"apply_to_org": &schema.Schema{
				Type:     schema.TypeBool,
				Default:  true,
				Optional: true,
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
			"when": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
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
		Create: resourcePostureCheckCreate,
		Read:   resourcePostureCheckRead,
		Update: resourcePostureCheckUpdate,
		Delete: resourcePostureCheckDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourcePostureCheckCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	action := d.Get("action").(string)
	osQuery := d.Get("osquery").(string)
	platform := d.Get("platform").(string)
	userMessageOnFail := d.Get("user_message_on_fail").(string)
	enabled := d.Get("enabled").(bool)
	applyToOrg := d.Get("apply_to_org").(bool)
	interval := d.Get("interval").(int)
	check := d.Get("check").([]interface{})
	when := resourceTypeSetToStringSlice(d.Get("when").(*schema.Set))
	applyToEntities := resourceTypeSetToStringSlice(d.Get("sources").(*schema.Set))
	exemptEntities := resourceTypeSetToStringSlice(d.Get("exempt_sources").(*schema.Set))

	postureCheck := PostureCheck{
		Name:              name,
		Description:       description,
		Action:            action,
		OSQuery:           osQuery,
		Platform:          platform,
		UserMessageOnFail: userMessageOnFail,
		Enabled:           enabled,
		ApplyToOrg:        applyToOrg,
		Interval:          interval,
		Check:             check,
		When:              when,
		ExemptEntities:    exemptEntities,
		ApplyToEntities:   applyToEntities,
	}

	var newPostureCheck *PostureCheck
	newPostureCheck, err := client.CreatePostureCheck(&postureCheck)
	if err != nil {
		return err
	}

	d.SetId(newPostureCheck.ID)

	err = postureCheckToResource(d, newPostureCheck)
	if err != nil {
		return err
	}

	return resourcePostureCheckRead(d, m)
}

func resourcePostureCheckRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	postureCheck, err := client.GetPostureCheck(d.Id())
	if err != nil {
		d.SetId("")
		return nil
	}

	err = postureCheckToResource(d, postureCheck)
	if err != nil {
		return err
	}

	return nil
}

func resourcePostureCheckUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	action := d.Get("action").(string)
	osQuery := d.Get("osquery").(string)
	platform := d.Get("platform").(string)
	userMessageOnFail := d.Get("user_message_on_fail").(string)
	enabled := d.Get("enabled").(bool)
	applyToOrg := d.Get("apply_to_org").(bool)
	interval := d.Get("interval").(int)
	check := d.Get("check").([]interface{})
	when := resourceTypeSetToStringSlice(d.Get("when").(*schema.Set))
	applyToEntities := resourceTypeSetToStringSlice(d.Get("sources").(*schema.Set))
	exemptEntities := resourceTypeSetToStringSlice(d.Get("exempt_sources").(*schema.Set))

	postureCheck := PostureCheck{
		Name:              name,
		Description:       description,
		Action:            action,
		OSQuery:           osQuery,
		Platform:          platform,
		UserMessageOnFail: userMessageOnFail,
		Enabled:           enabled,
		ApplyToOrg:        applyToOrg,
		Interval:          interval,
		Check:             check,
		When:              when,
		ExemptEntities:    exemptEntities,
		ApplyToEntities:   applyToEntities,
	}

	var updatedPostureCheck *PostureCheck
	updatedPostureCheck, err := client.UpdatePostureCheck(d.Id(), &postureCheck)
	if err != nil {
		return err
	}

	d.SetId(updatedPostureCheck.ID)

	err = postureCheckToResource(d, updatedPostureCheck)
	if err != nil {
		return err
	}

	return resourcePostureCheckRead(d, m)
}

func resourcePostureCheckDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	err := client.DeletePostureCheck(d.Id())
	if err != nil {
		return err
	}

	return nil
}

func postureCheckToResource(d *schema.ResourceData, m *PostureCheck) error {
	d.Set("description", m.Description)
	d.Set("name", m.Name)
	d.Set("action", m.Action)
	d.Set("osquery", m.OSQuery)
	d.Set("platform", m.Platform)
	d.Set("enabled", m.Enabled)
	d.Set("apply_to_org", m.ApplyToOrg)
	d.Set("user_message_on_fail", m.UserMessageOnFail)
	d.Set("interval", m.Interval)
	d.Set("check", m.Check)
	d.Set("when", m.When)
	d.Set("exempt_entities", m.ExemptEntities)
	d.Set("apply_to_entities", m.ApplyToEntities)
	d.Set("created_at", m.CreatedAt)
	d.Set("modified_at", m.ModifiedAt)

	d.SetId(m.ID)

	return nil
}
