package metanetworks

import (
	"fmt"
	"net/url"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	usersEndpoint string = "/v1/users"
)

type User struct {
	Description       string              `json:"description"`
	Email             string              `json:"email"`
	Enabled           bool                `json:"enabled"`
	FamilyName        string              `json:"family_name"`
	GivenName         string              `json:"given_name"`
	Phone             string              `json:"phone,omitempty"`
	ProvisionedBy     string              `json:"provisioned_by,omitempty"`
	CreatedAt         string              `json:"created_at,omitempty" meta_api:"read_only"`
	ID                string              `json:"id,omitempty" meta_api:"read_only"`
	Inventory         []string            `json:"inventory,omitempty" meta_api:"read_only"`
	MFAEnabled        bool                `json:"mfa_enabled,omitempty" meta_api:"read_only"`
	ModifiedAt        string              `json:"modified_at,omitempty" meta_api:"read_only"`
	Name              string              `json:"name,omitempty"`
	OrgID             string              `json:"org_id,omitempty" meta_api:"read_only"`
	OverlayMFAEnabled bool                `json:"overlay_mfa_enabled,omitempty"`
	PhoneVerified     bool                `json:"phone_verified,omitempty"`
	Roles             []string            `json:"roles,omitempty" meta_api:"read_only"`
	Tags              []map[string]string `json:"tags,omitempty" meta_api:"read_only"`
}

// userToResource ...
func userToResource(d *schema.ResourceData, m *User) error {
	d.Set("description", m.Description)
	d.Set("email", m.Email)
	d.Set("enabled", m.Description)
	d.Set("family_name", m.FamilyName)
	d.Set("given_name", m.GivenName)
	d.Set("phone", m.Phone)
	d.Set("provisioned_by", m.ProvisionedBy)
	d.Set("created_at", m.CreatedAt)
	d.Set("inventory", m.Inventory)
	d.Set("mfa_enabled", m.MFAEnabled)
	d.Set("modified_at", m.ModifiedAt)
	d.Set("name", m.Name)
	d.Set("org_id", m.OrgID)
	d.Set("overlay_mfa_enabled", m.OverlayMFAEnabled)
	d.Set("phone_verified", m.PhoneVerified)
	d.Set("roles", m.Roles)
	d.Set("tags", m.Tags)

	d.SetId(m.ID)

	return nil
}

func (c *Client) GetUsers(email string) ([]User, error) {
	var users []User
	err := c.Read(usersEndpoint+"?email="+url.QueryEscape(email), &users)
	if err != nil {
		return nil, err
	}

	if email != "" && len(users) == 0 {
		return nil, fmt.Errorf("Not found: %s", email)
	}

	return users, nil
}

func (c *Client) GetUser(userID string) (*User, error) {
	var user User
	err := c.Read(usersEndpoint+"/"+userID, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (c *Client) UpdateUser(userID string, user *User) (*User, error) {
	resp, err := c.Update(usersEndpoint+"/"+userID, *user)
	if err != nil {
		return nil, err
	}
	updatedUser, _ := resp.(*User)

	return updatedUser, nil
}

func (c *Client) CreateUser(user *User) (*User, error) {
	resp, err := c.Create(usersEndpoint, *user)
	if err != nil {
		return nil, err
	}
	createdUser, _ := resp.(*User)

	return createdUser, nil
}

func (c *Client) DeleteUser(userID string) error {
	err := c.Delete(usersEndpoint + "/" + userID)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) GetUserTags(userID string) (map[string]string, error) {
	tagMap, err := c.GetTags(usersEndpoint + "/" + userID + "/tags")
	if err != nil {
		return nil, err
	}

	return tagMap, nil
}

func (c *Client) SetUserTags(userID string, tags map[string]string) error {
	err := c.UpdateTags(usersEndpoint+"/"+userID+"/tags", tags)
	if err != nil {
		return err
	}
	return nil
}
