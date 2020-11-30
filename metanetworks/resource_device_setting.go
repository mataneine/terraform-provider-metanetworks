package metanetworks

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDeviceSetting() *schema.Resource {
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
			"direct_sso": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_server_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vpn_login_browser": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"protocol_selection_lifetime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"session_lifetime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"session_lifetime_grace": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Default:  true,
				Optional: true,
			},
			"split_tunnel": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"apply_on_org": &schema.Schema{
				Type:     schema.TypeBool,
				Default:  true,
				Optional: true,
			},
			"sources": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"search_domains": &schema.Schema{
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
		Create: resourceDeviceSettingCreate,
		Read:   resourceDeviceSettingRead,
		Update: resourceDeviceSettingUpdate,
		Delete: resourceDeviceSettingDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceDeviceSettingCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	directSSO := d.Get("direct_sso").(string)
	dnsServerType := d.Get("dns_server_type").(string)
	vpnLoginBrowser := d.Get("vpn_login_browser").(string)
	enabled := d.Get("enabled").(bool)
	applyOnOrg := d.Get("apply_on_org").(bool)
	splitTunnel := d.Get("split_tunnel").(bool)
	protocolSelectionLifetime := d.Get("protocol_selection_lifetime").(int)
	sessionLifetime := d.Get("session_lifetime").(int)
	sessionLifetimeGrace := d.Get("session_lifetime_grace").(int)
	searchDomains := resourceTypeSetToStringSlice(d.Get("search_domains").(*schema.Set))
	applyToEntities := resourceTypeSetToStringSlice(d.Get("sources").(*schema.Set))

	deviceSetting := DeviceSetting{
		Name:                      name,
		Description:               description,
		DirectSSO:                 directSSO,
		DNSServerType:             dnsServerType,
		VPNLoginBrowser:           vpnLoginBrowser,
		Enabled:                   enabled,
		ApplyOnOrg:                applyOnOrg,
		SplitTunnel:               splitTunnel,
		ProtocolSelectionLifetime: protocolSelectionLifetime,
		SessionLifetime:           sessionLifetime,
		SessionLifetimeGrace:      sessionLifetimeGrace,
		SearchDomains:             searchDomains,
		ApplyToEntities:           applyToEntities,
	}

	var newDeviceSetting *DeviceSetting
	newDeviceSetting, err := client.CreateDeviceSetting(&deviceSetting)
	if err != nil {
		return err
	}

	d.SetId(newDeviceSetting.ID)

	err = deviceSettingToResource(d, newDeviceSetting)
	if err != nil {
		return err
	}

	return resourceDeviceSettingRead(d, m)
}

func resourceDeviceSettingRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	deviceSetting, err := client.GetDeviceSetting(d.Id())
	if err != nil {
		d.SetId("")
		return nil
	}

	err = deviceSettingToResource(d, deviceSetting)
	if err != nil {
		return err
	}

	return nil
}

func resourceDeviceSettingUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	directSSO := d.Get("direct_sso").(string)
	dnsServerType := d.Get("dns_server_type").(string)
	vpnLoginBrowser := d.Get("vpn_login_browser").(string)
	enabled := d.Get("enabled").(bool)
	applyOnOrg := d.Get("apply_on_org").(bool)
	splitTunnel := d.Get("split_tunnel").(bool)
	protocolSelectionLifetime := d.Get("protocol_selection_lifetime").(int)
	sessionLifetime := d.Get("session_lifetime").(int)
	sessionLifetimeGrace := d.Get("session_lifetime_grace").(int)
	searchDomains := resourceTypeSetToStringSlice(d.Get("search_domains").(*schema.Set))
	applyToEntities := resourceTypeSetToStringSlice(d.Get("sources").(*schema.Set))

	deviceSetting := DeviceSetting{
		Name:                      name,
		Description:               description,
		DirectSSO:                 directSSO,
		DNSServerType:             dnsServerType,
		VPNLoginBrowser:           vpnLoginBrowser,
		Enabled:                   enabled,
		ApplyOnOrg:                applyOnOrg,
		SplitTunnel:               splitTunnel,
		ProtocolSelectionLifetime: protocolSelectionLifetime,
		SessionLifetime:           sessionLifetime,
		SessionLifetimeGrace:      sessionLifetimeGrace,
		SearchDomains:             searchDomains,
		ApplyToEntities:           applyToEntities,
	}

	var updatedDeviceSetting *DeviceSetting
	updatedDeviceSetting, err := client.UpdateDeviceSetting(d.Id(), &deviceSetting)
	if err != nil {
		return err
	}

	d.SetId(updatedDeviceSetting.ID)

	err = deviceSettingToResource(d, updatedDeviceSetting)
	if err != nil {
		return err
	}

	return resourceDeviceSettingRead(d, m)
}

func resourceDeviceSettingDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	err := client.DeleteDeviceSetting(d.Id())
	if err != nil {
		return err
	}

	return nil
}

func deviceSettingToResource(d *schema.ResourceData, m *DeviceSetting) error {
	d.Set("description", m.Description)
	d.Set("name", m.Name)
	d.Set("direct_sso", m.DirectSSO)
	d.Set("dns_server_type", m.DNSServerType)
	d.Set("vpn_login_browser", m.VPNLoginBrowser)
	d.Set("enabled", m.Enabled)
	d.Set("apply_on_org", m.ApplyOnOrg)
	d.Set("split_tunnel", m.SplitTunnel)
	d.Set("protocol_selection_lifetime", m.ProtocolSelectionLifetime)
	d.Set("session_lifetime", m.SessionLifetime)
	d.Set("session_lifetime_grace", m.SessionLifetimeGrace)
	d.Set("search_domains", m.SearchDomains)
	d.Set("apply_to_entities", m.ApplyToEntities)
	d.Set("created_at", m.CreatedAt)
	d.Set("modified_at", m.ModifiedAt)

	d.SetId(m.ID)

	return nil
}
