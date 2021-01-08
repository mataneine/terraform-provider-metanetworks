package metanetworks

import (
	"errors"
	"log"
)

const (
	versionControlsEndpoint string = "/v1/version_controls"
)

// VersionControls ...
type VersionControls struct {
	Name            string                 `json:"name"`
	Description     string                 `json:"description,omitempty"`
	Enabled         bool                   `json:"enabled" type:"bool"`
	ApplyToOrg      bool                   `json:"apply_to_org"`
	ExemptEntities  []string               `json:"exempt_entities,omitempty"`
	ApplyToEntities []string               `json:"apply_to_entities,omitempty"`
	WindowsPolicy   map[string]interface{} `json:"windows_policy,omitempty"`
	MacOSPolicy     map[string]interface{} `json:"macos_policy,omitempty"`
	LinuxPolicy     map[string]interface{} `json:"linux_policy,omitempty"`
	CreatedAt       string                 `json:"created_at,omitempty" meta_api:"read_only"`
	ID              string                 `json:"id,omitempty" meta_api:"read_only"`
	ModifiedAt      string                 `json:"modified_at,omitempty" meta_api:"read_only"`
}

// GetVersionControls ...
func (c *Client) GetVersionControls(versionControlsID string) (*VersionControls, error) {
	var versionControls VersionControls
	err := c.Read(versionControlsEndpoint+"/"+versionControlsID, &versionControls)
	if err != nil {
		return nil, err
	}

	log.Printf("Returning Version Control Settings from Get: %s", versionControls.ID)
	return &versionControls, nil
}

// UpdateVersionControls ...
func (c *Client) UpdateVersionControls(versionControlsID string, versionControls *VersionControls) (*VersionControls, error) {
	resp, err := c.Update(versionControlsEndpoint+"/"+versionControlsID, *versionControls)
	if err != nil {
		return nil, err
	}
	updatedVersionControls, _ := resp.(*VersionControls)

	log.Printf("Returning Version Control Settings from Update: %s", updatedVersionControls.ID)
	return updatedVersionControls, nil
}

// CreateVersionControls ...
func (c *Client) CreateVersionControls(versionControls *VersionControls) (*VersionControls, error) {
	resp, err := c.Create(versionControlsEndpoint, *versionControls)
	if err != nil {
		return nil, err
	}

	createdVersionControls, ok := resp.(*VersionControls)
	if !ok {
		return nil, errors.New("Object returned from API was not a Version Control Pointer")
	}

	log.Printf("Returning Version Control Settings from Create: %s", createdVersionControls.ID)
	return createdVersionControls, nil
}

// DeleteVersionControls ...
func (c *Client) DeleteVersionControls(versionControlsID string) error {
	err := c.Delete(versionControlsEndpoint + "/" + versionControlsID)
	if err != nil {
		return err
	}

	return nil
}
