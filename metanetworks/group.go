package metanetworks

import (
	"encoding/json"
	"errors"
	"log"
)

const (
	GroupEndpoint string = "/v1/groups"
)

type Group struct {
	Description   string   `json:"description,omitempty"`
	Expression    string   `json:"expression,omitempty"`
	Name          string   `json:"name"`
	ProvisionedBy string   `json:"provisioned_by,omitempty" meta_api:"read_only"`
	CreatedAt     string   `json:"created_at,omitempty" meta_api:"read_only"`
	Id            string   `json:"id,omitempty" meta_api:"read_only"`
	Members       []string `json:"members,omitempty" meta_api:"read_only"`
	ModifiedAt    string   `json:"modified_at,omitempty" meta_api:"read_only"`
	OrgId         string   `json:"org_id,omitempty" meta_api:"read_only"`
	Roles         []string `json:"roles,omitempty" meta_api:"read_only"`
	Users         []string `json:"users,omitempty" meta_api:"read_only"`
}

func (c *Client) GetGroup(element_id string) (*Group, error) {

	var Group Group
	err := c.Read(GroupEndpoint+"/"+element_id, &Group)
	if err != nil {
		return nil, err
	}

	log.Printf("Returning Group from Get " + string(Group.Id))
	return &Group, nil
}

func (c *Client) UpdateGroup(group_id string, group *Group) (*Group, error) {

	resp, err := c.Update(GroupEndpoint+"/"+group_id, *group)
	if err != nil {
		return nil, err
	}
	updatedGroup, _ := resp.(*Group)

	log.Printf("Returning Group from Update " + string(updatedGroup.Id))
	return updatedGroup, nil

}

func (c *Client) CreateGroup(group *Group) (*Group, error) {

	resp, err := c.Create(GroupEndpoint, *group)
	if err != nil {
		return nil, err
	}

	createdGroup, ok := resp.(*Group)
	if !ok {
		return nil, errors.New("Object returned from API was not a Group Pointer")
	}

	log.Printf("Returning Group from Create " + string(createdGroup.Id))
	return createdGroup, nil

}

func (c *Client) DeleteGroup(group_id string) error {
	err := c.DeleteRequest(GroupEndpoint + "/" + group_id)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) modifyUsers(group_id string, users []string, operation string) (*Group, error) {

	json_data, err := json.Marshal(users)
	if err != nil {
		return nil, err
	}

	resp, err := c.MakeRequest(GroupEndpoint+"/"+group_id+"/"+operation+"/", "POST", json_data, "application/json")
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
func (c *Client) AddGroupUsers(group_id string, users []string) (*Group, error) {

	return c.modifyUsers(group_id, users, "add_users")
}

func (c *Client) RemoveGroupUsers(group_id string, users []string) (*Group, error) {

	return c.modifyUsers(group_id, users, "remove_users")
}

func (c *Client) SetGroupRoles(group_id string, roles []string) (*Group, error) {

	json_data, err := json.Marshal(roles)
	if err != nil {
		return nil, err
	}
	resp, err := c.MakeRequest(GroupEndpoint+"/"+group_id+"/roles/", "PUT", json_data, "application/json")
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
