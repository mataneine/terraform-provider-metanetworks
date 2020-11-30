package metanetworks

import (
	"errors"
	"log"
)

const (
	userSettingEndpoint string = "/v1/settings/auth"
)

type UserSetting struct {
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

func (c *Client) GetUserSetting(userSettingID string) (*UserSetting, error) {
	var userSetting UserSetting
	err := c.Read(userSettingEndpoint+"/"+userSettingID, &userSetting)
	if err != nil {
		return nil, err
	}

	log.Printf("Returning Auth Setting from Get: %s", userSetting.ID)
	return &userSetting, nil
}

func (c *Client) UpdateUserSetting(userSettingID string, userSetting *UserSetting) (*UserSetting, error) {
	resp, err := c.Update(userSettingEndpoint+"/"+userSettingID, *userSetting)
	if err != nil {
		return nil, err
	}
	updatedUserSetting, _ := resp.(*UserSetting)

	log.Printf("Returning Auth Setting from Update: %s", updatedUserSetting.ID)
	return updatedUserSetting, nil
}

func (c *Client) CreateUserSetting(userSetting *UserSetting) (*UserSetting, error) {
	resp, err := c.Create(userSettingEndpoint, *userSetting)
	if err != nil {
		return nil, err
	}

	createdUserSetting, ok := resp.(*UserSetting)
	if !ok {
		return nil, errors.New("Object returned from API was not a Auth Setting Pointer")
	}

	log.Printf("Returning Auth Setting from Create: %s", createdUserSetting.ID)
	return createdUserSetting, nil
}

func (c *Client) DeleteUserSetting(userSettingID string) error {
	err := c.Delete(userSettingEndpoint + "/" + userSettingID)
	if err != nil {
		return err
	}

	return nil
}
