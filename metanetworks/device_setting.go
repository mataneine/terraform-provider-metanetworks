package metanetworks

import (
	"errors"
	"log"
)

const (
	deviceSettingEndpoint string = "/v1/settings/device"
)

type DeviceSetting struct {
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

func (c *Client) GetDeviceSetting(deviceSettingID string) (*DeviceSetting, error) {
	var deviceSetting DeviceSetting
	err := c.Read(deviceSettingEndpoint+"/"+deviceSettingID, &deviceSetting)
	if err != nil {
		return nil, err
	}

	log.Printf("Returning Auth Setting from Get: %s", deviceSetting.ID)
	return &deviceSetting, nil
}

func (c *Client) UpdateDeviceSetting(deviceSettingID string, deviceSetting *DeviceSetting) (*DeviceSetting, error) {
	resp, err := c.Update(deviceSettingEndpoint+"/"+deviceSettingID, *deviceSetting)
	if err != nil {
		return nil, err
	}
	updatedDeviceSetting, _ := resp.(*DeviceSetting)

	log.Printf("Returning Auth Setting from Update: %s", updatedDeviceSetting.ID)
	return updatedDeviceSetting, nil
}

func (c *Client) CreateDeviceSetting(deviceSetting *DeviceSetting) (*DeviceSetting, error) {
	resp, err := c.Create(deviceSettingEndpoint, *deviceSetting)
	if err != nil {
		return nil, err
	}

	createdDeviceSetting, ok := resp.(*DeviceSetting)
	if !ok {
		return nil, errors.New("Object returned from API was not a Auth Setting Pointer")
	}

	log.Printf("Returning Auth Setting from Create: %s", createdDeviceSetting.ID)
	return createdDeviceSetting, nil
}

func (c *Client) DeleteDeviceSetting(deviceSettingID string) error {
	err := c.Delete(deviceSettingEndpoint + "/" + deviceSettingID)
	if err != nil {
		return err
	}

	return nil
}
