package metanetworks

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSwgThreatCategories() *schema.Resource {
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
			"countries": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"confidence_level": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"risk_level": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"types": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
			"detail": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"title": {
				Type:     schema.TypeString,
				Computed: true,
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
		Create: resourceSwgThreatCategoriesCreate,
		Read:   resourceSwgThreatCategoriesRead,
		Update: resourceSwgThreatCategoriesUpdate,
		Delete: resourceSwgThreatCategoriesDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceSwgThreatCategoriesCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	countries := resourceTypeSetToStringSlice(d.Get("countries").(*schema.Set))
	types := resourceTypeSetToStringSlice(d.Get("types").(*schema.Set))
	confidenceLevel := d.Get("confidence_level").(string)
	riskLevel := d.Get("risk_level").(string)

	swgThreatCategories := SwgThreatCategories{
		Name:            name,
		Description:     description,
		Countries:       countries,
		Types:           types,
		ConfidenceLevel: confidenceLevel,
		RiskLevel:       riskLevel,
	}

	var newSwgThreatCategories *SwgThreatCategories
	newSwgThreatCategories, err := client.CreateSwgThreatCategories(&swgThreatCategories)
	if err != nil {
		return err
	}

	d.SetId(newSwgThreatCategories.ID)

	err = swgThreatCategoriesToResource(d, newSwgThreatCategories)
	if err != nil {
		return err
	}

	return resourceSwgThreatCategoriesRead(d, m)
}

func resourceSwgThreatCategoriesRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	var swgThreatCategories *SwgThreatCategories
	swgThreatCategories, err := client.GetSwgThreatCategories(d.Id())
	if err != nil {
		d.SetId("")
		return nil
	}

	err = swgThreatCategoriesToResource(d, swgThreatCategories)
	if err != nil {
		return err
	}

	return nil
}

func resourceSwgThreatCategoriesUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	countries := resourceTypeSetToStringSlice(d.Get("countries").(*schema.Set))
	types := resourceTypeSetToStringSlice(d.Get("types").(*schema.Set))
	confidenceLevel := d.Get("confidence_level").(string)
	riskLevel := d.Get("risk_level").(string)

	swgThreatCategories := SwgThreatCategories{
		Name:            name,
		Description:     description,
		Countries:       countries,
		Types:           types,
		ConfidenceLevel: confidenceLevel,
		RiskLevel:       riskLevel,
	}

	var updatedSwgThreatCategories *SwgThreatCategories
	updatedSwgThreatCategories, err := client.UpdateSwgThreatCategories(d.Id(), &swgThreatCategories)
	if err != nil {
		return err
	}

	d.SetId(updatedSwgThreatCategories.ID)

	err = swgThreatCategoriesToResource(d, updatedSwgThreatCategories)
	if err != nil {
		return err
	}

	return resourceSwgThreatCategoriesRead(d, m)
}

func resourceSwgThreatCategoriesDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	err := client.DeleteSwgThreatCategories(d.Id())
	return err
}

func swgThreatCategoriesToResource(d *schema.ResourceData, m *SwgThreatCategories) error {
	d.Set("name", m.Name)
	d.Set("description", m.Description)
	d.Set("countries", m.Countries)
	d.Set("types", m.Types)
	d.Set("confidence_level", m.ConfidenceLevel)
	d.Set("risk_level", m.RiskLevel)

	d.SetId(m.ID)

	return nil
}
