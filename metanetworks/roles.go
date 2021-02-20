package metanetworks

import (
	"errors"
	"log"
)

const (
	rolesEndpoint string = "/v1/roles"
)

// Roles ...
type Roles struct {
	Name              string        `json:"name"`
	Description       string        `json:"description"`
	AllSubOrgs        bool          `json:"all_suborgs"`
	SubOrgsExpression []interface{} `json:"suborgs_expression,omitempty"`
	ApplyToOrgs       []interface{} `json:"apply_to_orgs,omitempty"`
	Privileges        []interface{} `json:"privileges"`
	ReadOnly          bool          `json:"read_only"`
	CreatedAt         string        `json:"created_at,omitempty" meta_api:"read_only"`
	ID                string        `json:"id,omitempty" meta_api:"read_only"`
	ModifiedAt        string        `json:"modified_at,omitempty" meta_api:"read_only"`
}

// GetRoles ...
func (c *Client) GetRoles(rolesID string) (*Roles, error) {
	var roles Roles
	err := c.Read(rolesEndpoint+"/"+rolesID, &roles)
	if err != nil {
		return nil, err
	}

	log.Printf("Returning Roles Settings from Get: %s", roles.ID)
	return &roles, nil
}

// UpdateRoles ...
func (c *Client) UpdateRoles(rolesID string, roles *Roles) (*Roles, error) {
	resp, err := c.Update(rolesEndpoint+"/"+rolesID, *roles)
	if err != nil {
		return nil, err
	}
	updatedRoles, _ := resp.(*Roles)

	log.Printf("Returning Roles Settings from Update: %s", updatedRoles.ID)
	return updatedRoles, nil
}

// CreateRoles ...
func (c *Client) CreateRoles(roles *Roles) (*Roles, error) {
	resp, err := c.Create(rolesEndpoint, *roles)
	if err != nil {
		return nil, err
	}

	createdRoles, ok := resp.(*Roles)
	if !ok {
		return nil, errors.New("Object returned from API was not a Roles Pointer")
	}

	log.Printf("Returning Roles Settings from Create: %s", createdRoles.ID)
	return createdRoles, nil
}

// DeleteRoles ...
func (c *Client) DeleteRoles(rolesID string) error {
	err := c.Delete(rolesEndpoint + "/" + rolesID)
	if err != nil {
		return err
	}

	return nil
}
