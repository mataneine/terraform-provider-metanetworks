package metanetworks

import (
	"errors"
	"log"
)

const (
	deviceSettingsEndpoint string = "/v1/settings/device"
)

type DeviceSettings struct {
	Name                      string   `json:"name"`
	Description               string   `json:"description,omitempty"`
	DirectSSO                 string   `json:"direct_sso,omitempty"`
	VPNLoginBrowser           string   `json:"vpn_login_browser,omitempty"`
	DNSServerType             string   `json:"dns_server_type,omitempty"`
	Enabled                   bool     `json:"enabled" type:"bool"`
	ApplyOnOrg                bool     `json:"apply_on_org,omitempty"`
	SplitTunnel               bool     `json:"split_tunnel,omitempty" type:"bool"`
	ProtocolSelectionLifetime int      `json:"protocol_selection_lifetime,omitempty"`
	SessionLifetime           int      `json:"session_lifetime,omitempty"`
	SessionLifetimeGrace      int      `json:"session_lifetime_grace,omitempty"`
	SearchDomains             []string `json:"search_domains,omitempty"`
	ApplyToEntities           []string `json:"apply_to_entities,omitempty"`
	CreatedAt                 string   `json:"created_at,omitempty" meta_api:"read_only"`
	ID                        string   `json:"id,omitempty" meta_api:"read_only"`
	ModifiedAt                string   `json:"modified_at,omitempty" meta_api:"read_only"`
}

func (c *Client) GetDeviceSettings(deviceSettingsID string) (*DeviceSettings, error) {
	var deviceSettings DeviceSettings
	err := c.Read(deviceSettingsEndpoint+"/"+deviceSettingsID, &deviceSettings)
	if err != nil {
		return nil, err
	}

	log.Printf("Returning Auth Setting from Get: %s", deviceSettings.ID)
	return &deviceSettings, nil
}

func (c *Client) UpdateDeviceSettings(deviceSettingsID string, deviceSettings *DeviceSettings) (*DeviceSettings, error) {
	resp, err := c.Update(deviceSettingsEndpoint+"/"+deviceSettingsID, *deviceSettings)
	if err != nil {
		return nil, err
	}
	updatedDeviceSettings, _ := resp.(*DeviceSettings)

	log.Printf("Returning Auth Setting from Update: %s", updatedDeviceSettings.ID)
	return updatedDeviceSettings, nil
}

func (c *Client) CreateDeviceSettings(deviceSettings *DeviceSettings) (*DeviceSettings, error) {
	resp, err := c.Create(deviceSettingsEndpoint, *deviceSettings)
	if err != nil {
		return nil, err
	}

	createdDeviceSettings, ok := resp.(*DeviceSettings)
	if !ok {
		return nil, errors.New("Object returned from API was not a Auth Setting Pointer")
	}

	log.Printf("Returning Auth Setting from Create: %s", createdDeviceSettings.ID)
	return createdDeviceSettings, nil
}

func (c *Client) DeleteDeviceSettings(deviceSettingsID string) error {
	err := c.Delete(deviceSettingsEndpoint + "/" + deviceSettingsID)
	if err != nil {
		return err
	}

	return nil
}
