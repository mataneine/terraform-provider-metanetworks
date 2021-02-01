package metanetworks

import (
	"errors"
	"log"
)

const (
	swgUrlFilteringRulessEndpoint string = "/v1/url_filtering_rules"
)

// SwgUrlFilteringRules ...
type SwgUrlFilteringRules struct {
	Name                       string   `json:"name"`
	Description                string   `json:"description"`
	Action                     string   `json:"action"`
	AdvancedThreatProtection   bool     `json:"advanced_threat_protection,omitempty" type:"bool"`
	Enabled                    bool     `json:"enabled" type:"bool"`
	Priority                   int      `json:"priority"`
	ThreatCategory             string   `json:"threat_category,omitempty"`
	ExemptSources              []string `json:"exempt_sources,omitempty"`
	Sources                    []string `json:"sources,omitempty"`
	ForbiddenContentCategories []string `json:"forbidden_content_categories,omitempty"`
	CreatedAt                  string   `json:"created_at,omitempty" meta_api:"read_only"`
	ID                         string   `json:"id,omitempty" meta_api:"read_only"`
	ModifiedAt                 string   `json:"modified_at,omitempty" meta_api:"read_only"`
	OrgID                      string   `json:"org_id,omitempty" meta_api:"read_only"`
}

// GetSwgUrlFilteringRules ...
func (c *Client) GetSwgUrlFilteringRules(swgUrlFilteringRulesID string) (*SwgUrlFilteringRules, error) {
	var swgUrlFilteringRules SwgUrlFilteringRules
	err := c.Read(swgUrlFilteringRulessEndpoint+"/"+swgUrlFilteringRulesID+"?expand=true", &swgUrlFilteringRules)
	if err != nil {
		return nil, err
	}

	log.Printf("Returning SwgUrlFilteringRules from Get: %s", swgUrlFilteringRules.ID)
	return &swgUrlFilteringRules, nil
}

// UpdateSwgUrlFilteringRules ...
func (c *Client) UpdateSwgUrlFilteringRules(swgUrlFilteringRulesID string, swgUrlFilteringRules *SwgUrlFilteringRules) (*SwgUrlFilteringRules, error) {
	resp, err := c.Update(swgUrlFilteringRulessEndpoint+"/"+swgUrlFilteringRulesID, *swgUrlFilteringRules)
	if err != nil {
		return nil, err
	}
	updatedSwgUrlFilteringRules, _ := resp.(*SwgUrlFilteringRules)

	log.Printf("Returning SwgUrlFilteringRules from Update: %s", updatedSwgUrlFilteringRules.ID)
	return updatedSwgUrlFilteringRules, nil
}

// CreateSwgUrlFilteringRules ...
func (c *Client) CreateSwgUrlFilteringRules(swgUrlFilteringRules *SwgUrlFilteringRules) (*SwgUrlFilteringRules, error) {
	resp, err := c.Create(swgUrlFilteringRulessEndpoint, *swgUrlFilteringRules)
	if err != nil {
		return nil, err
	}

	createdSwgUrlFilteringRules, ok := resp.(*SwgUrlFilteringRules)
	if !ok {
		return nil, errors.New("Object returned from API was not a SwgUrlFilteringRules Pointer")
	}

	log.Printf("Returning SwgUrlFilteringRules from Create: %s", createdSwgUrlFilteringRules.ID)
	return createdSwgUrlFilteringRules, nil
}

//DeleteSwgUrlFilteringRules ...
func (c *Client) DeleteSwgUrlFilteringRules(swgUrlFilteringRulesID string) error {
	err := c.Delete(swgUrlFilteringRulessEndpoint + "/" + swgUrlFilteringRulesID)
	if err != nil {
		return err
	}

	return nil
}
