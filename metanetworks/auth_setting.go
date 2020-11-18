package metanetworks

import (
	"errors"
	"log"
)

const (
	authSettingEndpoint string = "/v1/settings/auth"
)

type AuthSetting struct {
	Description             string   `json:"description"`
	AllowedFactors          []string `json:"allowed_factors,omitempty"`
	ApplyToEntities         []string `json:"apply_to_entities,omitempty"`
	ProhibitedOS            []string `json:"prohibited_os,omitempty"`
	Enabled                 bool     `json:"enabled" type:"bool"`
	ApplyOnOrg              bool     `json:"apply_on_org,omitempty"`
	MFARequired             bool     `json:"mfa_required,omitempty"`
	OverlayMFARequired      bool     `json:"overlay_mfa_required,omitempty"`
	SSOMandatory            bool     `json:"sso_mandatory,omitempty"`
	Name                    string   `json:"name"`
	MaxDevicesPerUser       int      `json:"max_devices_per_user,omitempty"`
	OverlayMFARefreshPeriod int      `json:"overlay_mfa_refresh_period,omitempty"`
	PasswordExpiration      int      `json:"password_expiration,omitempty"`
	CreatedAt               string   `json:"created_at,omitempty" meta_api:"read_only"`
	ID                      string   `json:"id,omitempty" meta_api:"read_only"`
	ModifiedAt              string   `json:"modified_at,omitempty" meta_api:"read_only"`
}

func (c *Client) GetAuthSetting(authSettingID string) (*AuthSetting, error) {
	var authSetting AuthSetting
	err := c.Read(authSettingEndpoint+"/"+authSettingID, &authSetting)
	if err != nil {
		return nil, err
	}

	log.Printf("Returning Auth Setting from Get: %s", authSetting.ID)
	return &authSetting, nil
}

func (c *Client) UpdateAuthSetting(authSettingID string, authSetting *AuthSetting) (*AuthSetting, error) {
	resp, err := c.Update(authSettingEndpoint+"/"+authSettingID, *authSetting)
	if err != nil {
		return nil, err
	}
	updatedAuthSetting, _ := resp.(*AuthSetting)

	log.Printf("Returning Auth Setting from Update: %s", updatedAuthSetting.ID)
	return updatedAuthSetting, nil
}

func (c *Client) CreateAuthSetting(authSetting *AuthSetting) (*AuthSetting, error) {
	resp, err := c.Create(authSettingEndpoint, *authSetting)
	if err != nil {
		return nil, err
	}

	createdAuthSetting, ok := resp.(*AuthSetting)
	if !ok {
		return nil, errors.New("Object returned from API was not a Auth Setting Pointer")
	}

	log.Printf("Returning Auth Setting from Create: %s", createdAuthSetting.ID)
	return createdAuthSetting, nil
}

func (c *Client) DeleteAuthSetting(authSettingID string) error {
	err := c.Delete(authSettingEndpoint + "/" + authSettingID)
	if err != nil {
		return err
	}

	return nil
}
