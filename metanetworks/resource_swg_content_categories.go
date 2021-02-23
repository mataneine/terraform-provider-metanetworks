package metanetworks

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSwgContentCategories() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"confidence_level": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"forbid_uncategorized_urls": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"types": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"urls": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"detail": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"title": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
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
		Create: resourceSwgContentCategoriesCreate,
		Read:   resourceSwgContentCategoriesRead,
		Update: resourceSwgContentCategoriesUpdate,
		Delete: resourceSwgContentCategoriesDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceSwgContentCategoriesCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	confidenceLevel := d.Get("confidence_level").(string)
	forbidUncategorizedUrls := d.Get("forbid_uncategorized_urls").(bool)
	types := resourceTypeSetToStringSlice(d.Get("types").(*schema.Set))
	urls := resourceTypeSetToStringSlice(d.Get("urls").(*schema.Set))

	swgContentCategories := SwgContentCategories{
		Name:                    name,
		Description:             description,
		ConfidenceLevel:         confidenceLevel,
		ForbidUncategorizedUrls: forbidUncategorizedUrls,
		Types:                   types,
		Urls:                    urls,
	}

	var newSwgContentCategories *SwgContentCategories
	newSwgContentCategories, err := client.CreateSwgContentCategories(&swgContentCategories)
	if err != nil {
		return err
	}

	d.SetId(newSwgContentCategories.ID)
	err = swgContentCategoriesToResource(d, newSwgContentCategories)
	if err != nil {
		return err
	}

	return resourceSwgContentCategoriesRead(d, m)
}

func resourceSwgContentCategoriesRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	swgContentCategories, err := client.GetSwgContentCategories(d.Id())
	if err != nil {
		d.SetId("")
		return nil
	}

	err = swgContentCategoriesToResource(d, swgContentCategories)
	if err != nil {
		return err
	}

	return nil
}

func resourceSwgContentCategoriesUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	confidenceLevel := d.Get("confidence_level").(string)
	forbidUncategorizedUrls := d.Get("forbid_uncategorized_urls").(bool)
	types := resourceTypeSetToStringSlice(d.Get("types").(*schema.Set))
	urls := resourceTypeSetToStringSlice(d.Get("urls").(*schema.Set))

	swgContentCategories := SwgContentCategories{
		Name:                    name,
		Description:             description,
		ConfidenceLevel:         confidenceLevel,
		ForbidUncategorizedUrls: forbidUncategorizedUrls,
		Types:                   types,
		Urls:                    urls,
	}

	var updatedSwgContentCategories *SwgContentCategories
	updatedSwgContentCategories, err := client.UpdateSwgContentCategories(d.Id(), &swgContentCategories)
	if err != nil {
		return err
	}

	d.SetId(updatedSwgContentCategories.ID)

	err = swgContentCategoriesToResource(d, updatedSwgContentCategories)
	if err != nil {
		return err
	}

	return resourceSwgContentCategoriesRead(d, m)
}

func resourceSwgContentCategoriesDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	err := client.DeleteSwgContentCategories(d.Id())
	if err != nil {
		return err
	}

	return nil
}

func swgContentCategoriesToResource(d *schema.ResourceData, m *SwgContentCategories) error {
	d.Set("description", m.Description)
	d.Set("confidence_level", m.ConfidenceLevel)
	d.Set("forbid_uncategorized_urls", m.ForbidUncategorizedUrls)
	d.Set("types", m.Types)
	d.Set("urls", m.Urls)
	d.Set("detail", m.Detail)
	d.Set("status", m.Status)
	d.Set("title", m.Title)
	d.Set("type", m.Type)
	d.Set("created_at", m.CreatedAt)
	d.Set("modified_at", m.ModifiedAt)
	d.Set("org_id", m.OrgID)
	d.SetId(m.ID)

	return nil
}
