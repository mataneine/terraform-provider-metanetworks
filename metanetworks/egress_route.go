package metanetworks

import (
	"errors"
	"log"
)

const (
	egressRouteEndpoint string = "/v1/egress_routes"
)

type EgressRoute struct {
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	Via           string   `json:"via"`
	Enabled       bool     `json:"enabled" type:"bool"`
	Destinations  []string `json:"destinations,omitempty"`
	Sources       []string `json:"sources,omitempty"`
	ExemptSources []string `json:"exempt_sources,omitempty"`
	CreatedAt     string   `json:"created_at,omitempty" meta_api:"read_only"`
	ID            string   `json:"id,omitempty" meta_api:"read_only"`
	ModifiedAt    string   `json:"modified_at,omitempty" meta_api:"read_only"`
	OrgID         string   `json:"org_id,omitempty" meta_api:"read_only"`
}

func (c *Client) GetEgressRoute(egressRouteID string) (*EgressRoute, error) {
	var egressRoute EgressRoute
	err := c.Read(egressRouteEndpoint+"/"+egressRouteID, &egressRoute)
	if err != nil {
		return nil, err
	}

	log.Printf("Returning EgressRoute from Get: %s", egressRoute.ID)
	return &egressRoute, nil
}

func (c *Client) UpdateEgressRoute(egressRouteID string, egressRoute *EgressRoute) (*EgressRoute, error) {
	resp, err := c.Update(egressRouteEndpoint+"/"+egressRouteID, *egressRoute)
	if err != nil {
		return nil, err
	}
	updatedEgressRoute, _ := resp.(*EgressRoute)

	log.Printf("Returning EgressRoute from Update: %s", updatedEgressRoute.ID)
	return updatedEgressRoute, nil
}

func (c *Client) CreateEgressRoute(egressRoute *EgressRoute) (*EgressRoute, error) {
	resp, err := c.Create(egressRouteEndpoint, *egressRoute)
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

func (c *Client) DeleteEgressRoute(egressRouteID string) error {
	err := c.Delete(egressRouteEndpoint + "/" + egressRouteID)
	if err != nil {
		return err
	}

	return nil
}
