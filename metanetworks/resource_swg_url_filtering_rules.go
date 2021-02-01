package metanetworks

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSwgUrlFilteringRules() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"action": {
				Type:     schema.TypeString,
				Required: true,
			},
			"advanced_threat_protection": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"exempt_sources": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"sources": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"forbidden_content_categories": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				MaxItems: 5,
			},
			"priority": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"threat_category": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"modified_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"org_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		Create: resourceSwgUrlFilteringRulesCreate,
		Read:   resourceSwgUrlFilteringRulesRead,
		Update: resourceSwgUrlFilteringRulesUpdate,
		Delete: resourceSwgUrlFilteringRulesDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceSwgUrlFilteringRulesCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	action := d.Get("action").(string)
	advancedThreatProtection := d.Get("advanced_threat_protection").(bool)
	enabled := d.Get("enabled").(bool)
	priority := d.Get("priority").(int)
	threatCategory := d.Get("threat_category").(string)
	exemptSources := resourceTypeSetToStringSlice(d.Get("exempt_sources").(*schema.Set))
	sources := resourceTypeSetToStringSlice(d.Get("sources").(*schema.Set))
	forbiddenContentCategories := resourceTypeSetToStringSlice(d.Get("forbidden_content_categories").(*schema.Set))

	swgUrlFilteringRules := SwgUrlFilteringRules{
		Name:                       name,
		Description:                description,
		Action:                     action,
		AdvancedThreatProtection:   advancedThreatProtection,
		Enabled:                    enabled,
		Priority:                   priority,
		ThreatCategory:             threatCategory,
		ExemptSources:              exemptSources,
		Sources:                    sources,
		ForbiddenContentCategories: forbiddenContentCategories,
	}

	var newSwgUrlFilteringRules *SwgUrlFilteringRules
	newSwgUrlFilteringRules, err := client.CreateSwgUrlFilteringRules(&swgUrlFilteringRules)
	if err != nil {
		return err
	}

	d.SetId(newSwgUrlFilteringRules.ID)

	err = swgUrlFilteringRulesToResource(d, newSwgUrlFilteringRules)
	if err != nil {
		return err
	}

	return resourceSwgUrlFilteringRulesRead(d, m)
}

func resourceSwgUrlFilteringRulesRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	var swgUrlFilteringRules *SwgUrlFilteringRules
	swgUrlFilteringRules, err := client.GetSwgUrlFilteringRules(d.Id())
	if err != nil {
		d.SetId("")
		return nil
	}

	err = swgUrlFilteringRulesToResource(d, swgUrlFilteringRules)
	if err != nil {
		return err
	}

	return nil
}

func resourceSwgUrlFilteringRulesUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	action := d.Get("action").(string)
	advancedThreatProtection := d.Get("advanced_threat_protection").(bool)
	enabled := d.Get("enabled").(bool)
	priority := d.Get("priority").(int)
	threatCategory := d.Get("threat_category").(string)
	exemptSources := resourceTypeSetToStringSlice(d.Get("exempt_sources").(*schema.Set))
	sources := resourceTypeSetToStringSlice(d.Get("sources").(*schema.Set))
	forbiddenContentCategories := resourceTypeSetToStringSlice(d.Get("forbidden_content_categories").(*schema.Set))

	swgUrlFilteringRules := SwgUrlFilteringRules{
		Name:                       name,
		Description:                description,
		Action:                     action,
		AdvancedThreatProtection:   advancedThreatProtection,
		Enabled:                    enabled,
		Priority:                   priority,
		ThreatCategory:             threatCategory,
		ExemptSources:              exemptSources,
		Sources:                    sources,
		ForbiddenContentCategories: forbiddenContentCategories,
	}

	var updatedSwgUrlFilteringRules *SwgUrlFilteringRules
	updatedSwgUrlFilteringRules, err := client.UpdateSwgUrlFilteringRules(d.Id(), &swgUrlFilteringRules)
	if err != nil {
		return err
	}

	d.SetId(updatedSwgUrlFilteringRules.ID)

	err = swgUrlFilteringRulesToResource(d, updatedSwgUrlFilteringRules)
	if err != nil {
		return err
	}

	return resourceSwgUrlFilteringRulesRead(d, m)
}

func resourceSwgUrlFilteringRulesDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	err := client.DeleteSwgUrlFilteringRules(d.Id())
	return err
}

func swgUrlFilteringRulesToResource(d *schema.ResourceData, m *SwgUrlFilteringRules) error {
	d.Set("name", m.Name)
	d.Set("description", m.Description)
	d.Set("action", m.Action)
	d.Set("advanced_threat_protection", m.AdvancedThreatProtection)
	d.Set("enabled", m.Enabled)
	d.Set("priority", m.Priority)
	d.Set("threat_category", m.ThreatCategory)
	d.Set("exempt_sources", m.ExemptSources)
	d.Set("sources", m.Sources)
	d.Set("forbidden_content_categories", m.ForbiddenContentCategories)

	d.SetId(m.ID)

	return nil
}
