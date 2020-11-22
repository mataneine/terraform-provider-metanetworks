package metanetworks

import (
	"errors"
	"log"
)

const (
	policiesEndpoint string = "/v1/policies"
)

type Policy struct {
	Description    string   `json:"description"`
	Destinations   []string `json:"destinations,omitempty"`
	Enabled        bool     `json:"enabled,omitempty"`
	Name           string   `json:"name"`
	ProtocolGroups []string `json:"protocol_groups,omitempty"`
	ExemptSources  []string `json:"exempt_sources,omitempty"`
	Sources        []string `json:"sources,omitempty"`
	CreatedAt      string   `json:"created_at,omitempty" meta_api:"read_only"`
	ID             string   `json:"id,omitempty" meta_api:"read_only"`
	ModifiedAt     string   `json:"modified_at,omitempty" meta_api:"read_only"`
	OrgID          string   `json:"org_id,omitempty" meta_api:"read_only"`
}

func (c *Client) GetPolicy(policyID string) (*Policy, error) {
	var policy Policy
	err := c.Read(policiesEndpoint+"/"+policyID, &policy)
	if err != nil {
		return nil, err
	}

	log.Printf("Returning Policy from Get: %s", policy.ID)
	return &policy, nil
}

func (c *Client) UpdatePolicy(policyID string, policy *Policy) (*Policy, error) {
	resp, err := c.Update(policiesEndpoint+"/"+policyID, *policy)
	if err != nil {
		return nil, err
	}
	updatedPolicy, _ := resp.(*Policy)

	log.Printf("Returning Policy from Update: %s", updatedPolicy.ID)
	return updatedPolicy, nil
}

func (c *Client) CreatePolicy(policy *Policy) (*Policy, error) {
	resp, err := c.Create(policiesEndpoint, *policy)
	if err != nil {
		return nil, err
	}

	createdPolicy, ok := resp.(*Policy)
	if !ok {
		return nil, errors.New("Object returned from API was not a Policy Pointer")
	}

	log.Printf("Returning Policy from Create: %s", createdPolicy.ID)
	return createdPolicy, nil
}

func (c *Client) DeletePolicy(policyID string) error {
	err := c.Delete(policiesEndpoint + "/" + policyID)
	if err != nil {
		return err
	}

	return nil
}
