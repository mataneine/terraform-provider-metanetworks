package metanetworks

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDeviceSettings() *schema.Resource {
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
			"block_lan_access": &schema.Schema{
				Type:     schema.TypeBool,
				Default:  true,
				Optional: true,
			},
			"swg_always_on": &schema.Schema{
				Type:     schema.TypeBool,
				Default:  false,
				Optional: true,
			},
			"ztna_always_on": &schema.Schema{
				Type:     schema.TypeBool,
				Default:  false,
				Optional: true,
			},
			"session_expired_action": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Default:  "DISCONNECT",
			},
		},
		Create: resourceDeviceSettingsCreate,
		Read:   resourceDeviceSettingsRead,
		Update: resourceDeviceSettingsUpdate,
		Delete: resourceDeviceSettingsDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

// block_lan_access bool BlockLanAccess
// proxy_always_on bool SWGAlwaysOn
// session_expired_action [REQUIRE_REAUTH,DISCONNECT] SessionExpiredAction
// ztna_always_on bool ZTNAAlwaysOn
func resourceDeviceSettingsCreate(d *schema.ResourceData, m interface{}) error {
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
	blockLanAccess := d.Get("block_lan_access").(bool)
	swgAlwaysOn := d.Get("swg_always_on").(bool)
	ztnaAlwaysOn := d.Get("ztna_always_on").(bool)
	sessionExpiredAction := d.Get("session_expired_action").(string)

	deviceSettings := DeviceSettings{
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
		BlockLanAccess:            blockLanAccess,
		SWGAlwaysOn:               swgAlwaysOn,
		ZTNAAlwaysOn:              ztnaAlwaysOn,
		SessionExpiredAction:      sessionExpiredAction,
	}

	var newDeviceSettings *DeviceSettings
	newDeviceSettings, err := client.CreateDeviceSettings(&deviceSettings)
	if err != nil {
		return err
	}

	d.SetId(newDeviceSettings.ID)

	err = deviceSettingsToResource(d, newDeviceSettings)
	if err != nil {
		return err
	}

	return resourceDeviceSettingsRead(d, m)
}

func resourceDeviceSettingsRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	deviceSettings, err := client.GetDeviceSettings(d.Id())
	if err != nil {
		d.SetId("")
		return nil
	}

	err = deviceSettingsToResource(d, deviceSettings)
	if err != nil {
		return err
	}

	return nil
}

func resourceDeviceSettingsUpdate(d *schema.ResourceData, m interface{}) error {
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
	blockLanAccess := d.Get("block_lan_access").(bool)
	swgAlwaysOn := d.Get("swg_always_on").(bool)
	ztnaAlwaysOn := d.Get("ztna_always_on").(bool)
	sessionExpiredAction := d.Get("session_expired_action").(string)

	deviceSettings := DeviceSettings{
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
		BlockLanAccess:            blockLanAccess,
		SWGAlwaysOn:               swgAlwaysOn,
		ZTNAAlwaysOn:              ztnaAlwaysOn,
		SessionExpiredAction:      sessionExpiredAction,
	}

	var updatedDeviceSettings *DeviceSettings
	updatedDeviceSettings, err := client.UpdateDeviceSettings(d.Id(), &deviceSettings)
	if err != nil {
		return err
	}

	d.SetId(updatedDeviceSettings.ID)

	err = deviceSettingsToResource(d, updatedDeviceSettings)
	if err != nil {
		return err
	}

	return resourceDeviceSettingsRead(d, m)
}

func resourceDeviceSettingsDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	err := client.DeleteDeviceSettings(d.Id())
	if err != nil {
		return err
	}

	return nil
}

func deviceSettingsToResource(d *schema.ResourceData, m *DeviceSettings) error {
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
	d.Set("block_lan_access", m.BlockLanAccess)
	d.Set("swg_always_on", m.SWGAlwaysOn)
	d.Set("ztna_always_on", m.ZTNAAlwaysOn)
	d.Set("session_expired_action", m.SessionExpiredAction)

	d.SetId(m.ID)

	return nil
}
