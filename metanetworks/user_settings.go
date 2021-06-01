package metanetworks

import (
	"errors"
	"log"
)

const (
	userSettingsEndpoint string = "/v1/settings/auth"
)

type UserSettings struct {
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

func (c *Client) GetUserSettings(userSettingsID string) (*UserSettings, error) {
	var userSettings UserSettings
	err := c.Read(userSettingsEndpoint+"/"+userSettingsID, &userSettings)
	if err != nil {
		return nil, err
	}

	log.Printf("Returning Auth Setting from Get: %s", userSettings.ID)
	return &userSettings, nil
}

func (c *Client) UpdateUserSettings(userSettingsID string, userSettings *UserSettings) (*UserSettings, error) {
	resp, err := c.Update(userSettingsEndpoint+"/"+userSettingsID, *userSettings)
	if err != nil {
		return nil, err
	}
	updatedUserSettings, _ := resp.(*UserSettings)

	log.Printf("Returning Auth Setting from Update: %s", updatedUserSettings.ID)
	return updatedUserSettings, nil
}

func (c *Client) CreateUserSettings(userSettings *UserSettings) (*UserSettings, error) {
	resp, err := c.Create(userSettingsEndpoint, *userSettings)
	if err != nil {
		return nil, err
	}

	createdUserSettings, ok := resp.(*UserSettings)
	if !ok {
		return nil, errors.New("Object returned from API was not a Auth Setting Pointer")
	}

	log.Printf("Returning Auth Setting from Create: %s", createdUserSettings.ID)
	return createdUserSettings, nil
}

func (c *Client) DeleteUserSettings(userSettingsID string) error {
	err := c.Delete(userSettingsEndpoint + "/" + userSettingsID)
	if err != nil {
		return err
	}

	return nil
}
