package metanetworks

import (
	"errors"
	"log"
)

const (
	PolicyEndpoint string = "/v1/policies"
)

type Policy struct {
	Description    string   `json:"description,omitempty"`
	Destinations   []string `json:"destinations,omitempty"`
	Enabled        bool     `json:"enabled,omitempty"`
	Name           string   `json:"name"`
	ProtocolGroups []string `json:"protocol_groups,omitempty"`
	Sources        []string `json:"sources,omitempty"`
	CreatedAt      string   `json:"created_at,omitempty" meta_api:"read_only"`
	Id             string   `json:"id,omitempty" meta_api:"read_only"`
	ModifiedAt     string   `json:"modified_at,omitempty" meta_api:"read_only"`
	OrgId          string   `json:"org_id,omitempty" meta_api:"read_only"`
}

func (c *Client) GetPolicy(policy_id string) (*Policy, error) {

	var policy Policy
	err := c.Read(PolicyEndpoint+"/"+policy_id, &policy)
	if err != nil {
		return nil, err
	}

	log.Printf("Returning Policy from Get " + string(policy.Id))
	return &policy, nil

}

func (c *Client) UpdatePolicy(policy_id string, policy *Policy) (*Policy, error) {

	resp, err := c.Update(PolicyEndpoint+"/"+policy_id, *policy)
	if err != nil {
		return nil, err
	}
	updated_policy, _ := resp.(*Policy)

	log.Printf("Returning Policy from Update " + string(updated_policy.Id))
	return updated_policy, nil

}

func (c *Client) CreatePolicy(policy *Policy) (*Policy, error) {

	resp, err := c.Create(PolicyEndpoint, *policy)
	if err != nil {
		return nil, err
	}

	createdPolicy, ok := resp.(*Policy)
	if !ok {
		return nil, errors.New("Object returned from API was not a Policy Pointer")
	}

	log.Printf("Returning Policy from Create " + string(createdPolicy.Id))
	return createdPolicy, nil

}

func (c *Client) DeletePolicy(policy_id string) error {

	err := c.DeleteRequest(PolicyEndpoint + "/" + policy_id)
	if err != nil {
		return err
	}

	return nil

}
