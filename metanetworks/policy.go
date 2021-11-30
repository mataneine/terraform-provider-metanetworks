package metanetworks

import (
	"errors"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const (
	policiesEndpoint string = "/v1/policies"
)

// Policy ...
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

// GetPolicy ...
func (c *Client) GetPolicy(policyID string) (*Policy, error) {
	var policy Policy
	err := c.Read(policiesEndpoint+"/"+policyID, &policy)
	if err != nil {
		return nil, err
	}

	log.Printf("Returning Policy from Get: %s", policy.ID)
	return &policy, nil
}

// UpdatePolicy ...
func (c *Client) UpdatePolicy(policyID string, policy *Policy) (*Policy, error) {
	resp, err := c.Update(policiesEndpoint+"/"+policyID, *policy)
	if err != nil {
		return nil, err
	}
	updatedPolicy, _ := resp.(*Policy)

	log.Printf("Returning Policy from Update: %s", updatedPolicy.ID)
	return updatedPolicy, nil
}

// CreatePolicy ...
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

// DeletePolicy ...
func (c *Client) DeletePolicy(policyID string) error {
	err := c.Delete(policiesEndpoint + "/" + policyID)
	if err != nil {
		return err
	}

	return nil
}

func StatusPolicyCreate(client *Client, PolicyID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		var metaport *MetaPort
		_, err := client.GetPolicy(PolicyID)
		if err != nil {
			return metaport, "Pending", nil
		}
		return metaport, "Completed", nil
	}
}

func WaitPolicyCreate(client *Client, PolicyID string) (*Client, error) {
	createStateConf := &resource.StateChangeConf{
		Pending:    []string{"Pending"},
		Target:     []string{"Completed"},
		Timeout:    30 * time.Second,
		MinTimeout: 5 * time.Second,
		Delay:      2 * time.Second,
		Refresh:    StatusPolicyCreate(client, PolicyID),
	}

	_, err := createStateConf.WaitForState()
	if err != nil {
		return nil, err
	}

	return client, err
}
