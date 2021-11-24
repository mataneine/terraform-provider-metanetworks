package metanetworks

import (
	"errors"
	"log"
)

const (
	swgThreatCategoriessEndpoint string = "/v1/threat_categories"
)

// SwgThreatCategories ...
type SwgThreatCategories struct {
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	Countries       []string `json:"countries,omitempty"`
	Types           []string `json:"types"`
	ConfidenceLevel string   `json:"confidence_level,omitempty"`
	RiskLevel       string   `json:"risk_level,omitempty"`
	CreatedAt       string   `json:"created_at,omitempty" meta_api:"read_only"`
	ID              string   `json:"id,omitempty" meta_api:"read_only"`
	ModifiedAt      string   `json:"modified_at,omitempty" meta_api:"read_only"`
	OrgID           string   `json:"org_id,omitempty" meta_api:"read_only"`
}

// GetSwgThreatCategories ...
func (c *Client) GetSwgThreatCategories(swgThreatCategoriesID string) (*SwgThreatCategories, error) {
	var swgThreatCategories SwgThreatCategories
	err := c.Read(swgThreatCategoriessEndpoint+"/"+swgThreatCategoriesID+"?expand=true", &swgThreatCategories)
	if err != nil {
		return nil, err
	}

	log.Printf("Returning SwgThreatCategories from Get: %s", swgThreatCategories.ID)
	return &swgThreatCategories, nil
}

// UpdateSwgThreatCategories ...
func (c *Client) UpdateSwgThreatCategories(swgThreatCategoriesID string, swgThreatCategories *SwgThreatCategories) (*SwgThreatCategories, error) {
	resp, err := c.Update(swgThreatCategoriessEndpoint+"/"+swgThreatCategoriesID, *swgThreatCategories)
	if err != nil {
		return nil, err
	}
	updatedSwgThreatCategories, _ := resp.(*SwgThreatCategories)

	log.Printf("Returning SwgThreatCategories from Update: %s", updatedSwgThreatCategories.ID)
	return updatedSwgThreatCategories, nil
}

// CreateSwgThreatCategories ...
func (c *Client) CreateSwgThreatCategories(swgThreatCategories *SwgThreatCategories) (*SwgThreatCategories, error) {
	resp, err := c.Create(swgThreatCategoriessEndpoint, *swgThreatCategories)
	if err != nil {
		return nil, err
	}

	createdSwgThreatCategories, ok := resp.(*SwgThreatCategories)
	if !ok {
		return nil, errors.New("Object returned from API was not a SwgThreatCategories Pointer")
	}

	log.Printf("Returning SwgThreatCategories from Create: %s", createdSwgThreatCategories.ID)
	return createdSwgThreatCategories, nil
}

//DeleteSwgThreatCategories ...
func (c *Client) DeleteSwgThreatCategories(swgThreatCategoriesID string) error {
	err := c.Delete(swgThreatCategoriessEndpoint + "/" + swgThreatCategoriesID)
	if err != nil {
		return err
	}

	return nil
}
