package metanetworks

import (
	"errors"
	"log"
)

const (
	egressRoutesEndpoint string = "/v1/egress_routes"
)

// EgressRoute ...
type EgressRoute struct {
	Description   string   `json:"description"`
	Destinations  []string `json:"destinations,omitempty"`
	Enabled       bool     `json:"enabled"`
	ExemptSources []string `json:"exempt_sources,omitempty"`
	Name          string   `json:"name"`
	Sources       []string `json:"sources,omitempty"`
	Via           string   `json:"via"`
	CreatedAt     string   `json:"created_at,omitempty" meta_api:"read_only"`
	ID            string   `json:"id,omitempty" meta_api:"read_only"`
	ModifiedAt    string   `json:"modified_at,omitempty" meta_api:"read_only"`
	OrgID         string   `json:"org_id,omitempty" meta_api:"read_only"`
}

// GetEgressRoute ...
func (c *Client) GetEgressRoute(egressRouteID string) (*EgressRoute, error) {
	var egressRoute EgressRoute
	err := c.Read(egressRoutesEndpoint+"/"+egressRouteID, &egressRoute)
	if err != nil {
		return nil, err
	}

	log.Printf("Returning EgressRoute from Get: %s", egressRoute.ID)
	return &egressRoute, nil
}

// UpdateEgressRoute ...
func (c *Client) UpdateEgressRoute(egressRouteID string, egressRoute *EgressRoute) (*EgressRoute, error) {
	resp, err := c.Update(egressRoutesEndpoint+"/"+egressRouteID, *egressRoute)
	if err != nil {
		return nil, err
	}
	updatedEgressRoute, _ := resp.(*EgressRoute)

	log.Printf("Returning EgressRoute from Update: %s", updatedEgressRoute.ID)
	return updatedEgressRoute, nil
}

// CreateEgressRoute ...
func (c *Client) CreateEgressRoute(egressRoute *EgressRoute) (*EgressRoute, error) {
	resp, err := c.Create(egressRoutesEndpoint, *egressRoute)
	if err != nil {
		return nil, err
	}

	createdEgressRoute, ok := resp.(*EgressRoute)
	if !ok {
		return nil, errors.New("Object returned from API was not a EgressRoute Pointer")
	}

	log.Printf("Returning EgressRoute from Create: %s", createdEgressRoute.ID)
	return createdEgressRoute, nil
}

// DeleteEgressRoute ...
func (c *Client) DeleteEgressRoute(egressRouteID string) error {
	err := c.Delete(egressRoutesEndpoint + "/" + egressRouteID)
	if err != nil {
		return err
	}

	return nil
}
