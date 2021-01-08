package metanetworks

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/url"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	groupsEndpoint string = "/v1/groups"
)

type Group struct {
	Description   string   `json:"description"`
	Expression    string   `json:"expression,omitempty"`
	Name          string   `json:"name"`
	ProvisionedBy string   `json:"provisioned_by,omitempty" meta_api:"read_only"`
	CreatedAt     string   `json:"created_at,omitempty" meta_api:"read_only"`
	ID            string   `json:"id,omitempty" meta_api:"read_only"`
	Members       []string `json:"members,omitempty" meta_api:"read_only"`
	ModifiedAt    string   `json:"modified_at,omitempty" meta_api:"read_only"`
	OrgID         string   `json:"org_id,omitempty" meta_api:"read_only"`
	Roles         []string `json:"roles,omitempty" meta_api:"read_only"`
	Users         []string `json:"users,omitempty" meta_api:"read_only"`
}

// groupToResource ...
func groupToResource(d *schema.ResourceData, m *Group) error {
	d.Set("description", m.Description)
	d.Set("name", m.Name)
	d.Set("destinations", m.Description)
	d.Set("enabled", m.Description)
	d.Set("protocol_groups", m.Description)
	d.Set("sources", m.Description)
	d.Set("created_at", m.CreatedAt)
	d.Set("modified_at", m.ModifiedAt)
	d.Set("org_id", m.OrgID)

	d.SetId(m.ID)

	return nil
}

// GetGroups ...
func (c *Client) GetGroups(name string) ([]Group, error) {
	var groups []Group
	err := c.Read(groupsEndpoint+"?name="+url.QueryEscape(name), &groups)

	if err != nil {
		return nil, err
	}

	if name != "" && len(groups) == 0 {
		return nil, fmt.Errorf("Not found: %s", name)
	}

	return groups, nil
}

// GetGroup ...
func (c *Client) GetGroup(elementID string) (*Group, error) {
	var Group Group
	err := c.Read(groupsEndpoint+"/"+elementID, &Group)
	if err != nil {
		return nil, err
	}

	log.Printf("Returning Group from Get: %s", Group.ID)
	return &Group, nil
}

// UpdateGroup ...
func (c *Client) UpdateGroup(groupID string, group *Group) (*Group, error) {
	resp, err := c.Update(groupsEndpoint+"/"+groupID, *group)
	if err != nil {
		return nil, err
	}
	updatedGroup, _ := resp.(*Group)

	log.Printf("Returning Group from Update: %s", updatedGroup.ID)
	return updatedGroup, nil
}

// CreateGroup ...
func (c *Client) CreateGroup(group *Group) (*Group, error) {
	resp, err := c.Create(groupsEndpoint, *group)
	if err != nil {
		return nil, err
	}

	createdGroup, ok := resp.(*Group)
	if !ok {
		return nil, errors.New("Object returned from API was not a Group Pointer")
	}

	log.Printf("Returning Group from Create: %s", createdGroup.ID)
	return createdGroup, nil
}

// DeleteGroup ...
func (c *Client) DeleteGroup(groupID string) error {
	err := c.Delete(groupsEndpoint + "/" + groupID)
	if err != nil {
		return err
	}

	return nil
}

// modifyUsers ...
func (c *Client) modifyUsers(groupID string, users []string, operation string) (*Group, error) {
	jsonData, err := json.Marshal(users)
	if err != nil {
		return nil, err
	}

	resp, err := c.Request(groupsEndpoint+"/"+groupID+"/"+operation+"/", "POST", jsonData, "application/json")
	if err != nil {
		return nil, err
	}

	var group Group
	err = json.Unmarshal(resp, &group)
	if err != nil {
		return nil, err
	}

	return &group, nil
}

// AddGroupUsers ...
func (c *Client) AddGroupUsers(groupID string, users []string) (*Group, error) {
	return c.modifyUsers(groupID, users, "add_users")
}

// RemoveGroupUsers ...
func (c *Client) RemoveGroupUsers(groupID string, users []string) (*Group, error) {
	return c.modifyUsers(groupID, users, "remove_users")
}

// SetGroupRoles ...
func (c *Client) SetGroupRoles(groupID string, roles []string) (*Group, error) {
	jsonData, err := json.Marshal(roles)
	if err != nil {
		return nil, err
	}
	resp, err := c.Request(groupsEndpoint+"/"+groupID+"/roles/", "PUT", jsonData, "application/json")
	if err != nil {
		return nil, err
	}
	var group Group
	err = json.Unmarshal(resp, &group)
	if err != nil {
		return nil, err
	}

	return &group, nil
}
