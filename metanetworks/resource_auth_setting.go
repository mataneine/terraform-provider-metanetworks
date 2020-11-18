package metanetworks

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAuthSetting() *schema.Resource {
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
			"allowed_factors": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Default:  true,
				Optional: true,
			},
			"apply_on_org": &schema.Schema{
				Type:     schema.TypeBool,
				Default:  true,
				Optional: true,
			},
			"mfa_required": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"overlay_mfa_required": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"sso_mandatory": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"apply_to_entities": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"prohibited_os": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"max_devices_per_user": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"overlay_mfa_refresh_period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"password_expiration": &schema.Schema{
				Type:     schema.TypeInt,
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
		Create: resourceAuthSettingCreate,
		Read:   resourceAuthSettingRead,
		Update: resourceAuthSettingUpdate,
		Delete: resourceAuthSettingDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceAuthSettingCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	enabled := d.Get("enabled").(bool)
	applyOnOrg := d.Get("apply_on_org").(bool)
	mfaRequired := d.Get("mfa_required").(bool)
	overlayMFARequired := d.Get("overlay_mfa_required").(bool)
	ssoMandatory := d.Get("sso_mandatory").(bool)
	maxDevicesPerUser := d.Get("max_devices_per_user").(int)
	overlayMFARefreshPeriod := d.Get("overlay_mfa_refresh_period").(int)
	passwordExpiration := d.Get("password_expiration").(int)
	allowedFactors := resourceTypeSetToStringSlice(d.Get("allowed_factors").(*schema.Set))
	applyToEntities := resourceTypeSetToStringSlice(d.Get("apply_to_entities").(*schema.Set))
	prohibitedOS := resourceTypeSetToStringSlice(d.Get("prohibited_os").(*schema.Set))

	authSetting := AuthSetting{
		Name:                    name,
		Description:             description,
		Enabled:                 enabled,
		ApplyOnOrg:              applyOnOrg,
		MFARequired:             mfaRequired,
		OverlayMFARequired:      overlayMFARequired,
		SSOMandatory:            ssoMandatory,
		MaxDevicesPerUser:       maxDevicesPerUser,
		OverlayMFARefreshPeriod: overlayMFARefreshPeriod,
		PasswordExpiration:      passwordExpiration,
		AllowedFactors:          allowedFactors,
		ApplyToEntities:         applyToEntities,
		ProhibitedOS:            prohibitedOS,
	}

	var newAuthSetting *AuthSetting
	newAuthSetting, err := client.CreateAuthSetting(&authSetting)
	if err != nil {
		return err
	}

	d.SetId(newAuthSetting.ID)

	err = authSettingToResource(d, newAuthSetting)
	if err != nil {
		return err
	}

	return resourceAuthSettingRead(d, m)
}

func resourceAuthSettingRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	authSetting, err := client.GetAuthSetting(d.Id())
	if err != nil {
		d.SetId("")
		return nil
	}

	err = authSettingToResource(d, authSetting)
	if err != nil {
		return err
	}

	return nil
}

func resourceAuthSettingUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	enabled := d.Get("enabled").(bool)
	applyOnOrg := d.Get("apply_on_org").(bool)
	mfaRequired := d.Get("mfa_required").(bool)
	overlayMFARequired := d.Get("overlay_mfa_required").(bool)
	ssoMandatory := d.Get("sso_mandatory").(bool)
	maxDevicesPerUser := d.Get("max_devices_per_user").(int)
	overlayMFARefreshPeriod := d.Get("overlay_mfa_refresh_period").(int)
	passwordExpiration := d.Get("password_expiration").(int)
	allowedFactors := resourceTypeSetToStringSlice(d.Get("allowed_factors").(*schema.Set))
	applyToEntities := resourceTypeSetToStringSlice(d.Get("apply_to_entities").(*schema.Set))
	prohibitedOS := resourceTypeSetToStringSlice(d.Get("prohibited_os").(*schema.Set))

	authSetting := AuthSetting{
		Name:                    name,
		Description:             description,
		Enabled:                 enabled,
		ApplyOnOrg:              applyOnOrg,
		MFARequired:             mfaRequired,
		OverlayMFARequired:      overlayMFARequired,
		SSOMandatory:            ssoMandatory,
		MaxDevicesPerUser:       maxDevicesPerUser,
		OverlayMFARefreshPeriod: overlayMFARefreshPeriod,
		PasswordExpiration:      passwordExpiration,
		AllowedFactors:          allowedFactors,
		ApplyToEntities:         applyToEntities,
		ProhibitedOS:            prohibitedOS,
	}

	var updatedAuthSetting *AuthSetting
	updatedAuthSetting, err := client.UpdateAuthSetting(d.Id(), &authSetting)
	if err != nil {
		return err
	}

	d.SetId(updatedAuthSetting.ID)

	err = authSettingToResource(d, updatedAuthSetting)
	if err != nil {
		return err
	}

	return resourceAuthSettingRead(d, m)
}

func resourceAuthSettingDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	err := client.DeleteAuthSetting(d.Id())
	if err != nil {
		return err
	}

	return nil
}

func authSettingToResource(d *schema.ResourceData, m *AuthSetting) error {
	d.Set("description", m.Description)
	d.Set("name", m.Name)
	d.Set("allowed_factors", m.AllowedFactors)
	d.Set("enabled", m.Enabled)
	d.Set("apply_to_entities", m.ApplyToEntities)
	d.Set("prohibited_os", m.ProhibitedOS)
	d.Set("created_at", m.CreatedAt)
	d.Set("modified_at", m.ModifiedAt)
	d.Set("apply_on_org", m.ApplyOnOrg)
	d.Set("mfa_required", m.MFARequired)
	d.Set("overlay_mfa_required", m.OverlayMFARequired)
	d.Set("sso_mandatory", m.SSOMandatory)
	d.Set("max_devices_per_user", m.MaxDevicesPerUser)
	d.Set("overlay_mfa_refresh_period", m.OverlayMFARefreshPeriod)
	d.Set("password_expiration", m.PasswordExpiration)
	d.SetId(m.ID)

	return nil
}
