package metanetworks

import (
	"errors"
	"log"
)

const (
	swgContentCategoriesEndpoint string = "/v1/content_categories"
)

// SwgContentCategories ...
type SwgContentCategories struct {
	Description             string   `json:"description"`
	ConfidenceLevel         string   `json:"confidence_level,omitempty"`
	Name                    string   `json:"name"`
	ForbidUncategorizedUrls bool     `json:"forbid_uncategorized_urls,omitempty" type:"bool"`
	Types                   []string `json:"types,omitempty"`
	Urls                    []string `json:"urls,omitempty"`
	Detail                  string   `json:"detail,omitempty" meta_api:"read_only"`
	ID                      string   `json:"id,omitempty" meta_api:"read_only"`
	Status                  string   `json:"status,omitempty" meta_api:"read_only"`
	Title                   string   `json:"title,omitempty" meta_api:"read_only"`
	OrgID                   string   `json:"org_id,omitempty" meta_api:"read_only"`
	Type                    string   `json:"type,omitempty" meta_api:"read_only"`
	CreatedAt               string   `json:"created_at,omitempty" meta_api:"read_only"`
	ModifiedAt              string   `json:"modified_at,omitempty" meta_api:"read_only"`
}

// GetSwgContentCategories ...
func (c *Client) GetSwgContentCategories(swgContentCategoriesID string) (*SwgContentCategories, error) {
	var SwgContentCategories SwgContentCategories
	err := c.Read(swgContentCategoriesEndpoint+"/"+swgContentCategoriesID+"?expand=true", &SwgContentCategories)
	if err != nil {
		return nil, err
	}

	log.Printf("Returning SwgContentCategories from Get: %s", SwgContentCategories.ID)
	return &SwgContentCategories, nil
}

// UpdateSwgContentCategories ...
func (c *Client) UpdateSwgContentCategories(swgContentCategoriesID string, swgContentCategories *SwgContentCategories) (*SwgContentCategories, error) {
	resp, err := c.Update(swgContentCategoriesEndpoint+"/"+swgContentCategoriesID, *swgContentCategories)
	if err != nil {
		return nil, err
	}
	updatedSwgContentCategories, _ := resp.(*SwgContentCategories)

	log.Printf("Returning SwgContentCategories from Update: %s", updatedSwgContentCategories.ID)
	return updatedSwgContentCategories, nil
}

// CreateSwgContentCategories ...
func (c *Client) CreateSwgContentCategories(swgContentCategories *SwgContentCategories) (*SwgContentCategories, error) {
	resp, err := c.Create(swgContentCategoriesEndpoint, *swgContentCategories)
	if err != nil {
		return nil, err
	}

	createdSwgContentCategories, ok := resp.(*SwgContentCategories)
	if !ok {
		return nil, errors.New("Object returned from API was not a SwgContentCategories Pointer")
	}

	log.Printf("Returning SwgContentCategories from Create: %s", createdSwgContentCategories.ID)
	return createdSwgContentCategories, nil
}

// DeleteSwgContentCategories ...
func (c *Client) DeleteSwgContentCategories(swgContentCategoriesID string) error {
	err := c.Delete(swgContentCategoriesEndpoint + "/" + swgContentCategoriesID)
	if err != nil {
		return err
	}

	return nil
}
