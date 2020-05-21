package metanetworks

import (
	"errors"
	"log"
)

const (
	routingGroupsEndpoint string = "/v1/routing_groups"
)

type RoutingGroup struct {
	Name           string   `json:"name"`
	Description    string   `json:"description"`
	MappedElements []string `json:"mapped_elements_ids,omitempty"`
	Sources        []string `json:"sources,omitempty"`
	CreatedAt      string   `json:"created_at,omitempty" meta_api:"read_only"`
	ID             string   `json:"id,omitempty" meta_api:"read_only"`
	ModifiedAt     string   `json:"modified_at,omitempty" meta_api:"read_only"`
	OrgID          string   `json:"org_id,omitempty" meta_api:"read_only"`
	Priority       int      `json:"priority,omitempty" meta_api:"read_only"`
}

func (c *Client) GetRoutingGroup(routingGroupID string) (*RoutingGroup, error) {
	var routingGroup RoutingGroup
	err := c.Read(routingGroupsEndpoint+"/"+routingGroupID, &routingGroup)
	if err != nil {
		return nil, err
	}

	log.Printf("Returning RoutingGroup from Get: %s", routingGroup.ID)
	return &routingGroup, nil
}

func (c *Client) UpdateRoutingGroup(routingGroupID string, routingGroup *RoutingGroup) (*RoutingGroup, error) {
	resp, err := c.Update(routingGroupsEndpoint+"/"+routingGroupID, *routingGroup)
	if err != nil {
		return nil, err
	}
	updatedRoutingGroup, _ := resp.(*RoutingGroup)

	log.Printf("Returning RoutingGroup from Update: %s", updatedRoutingGroup.ID)
	return updatedRoutingGroup, nil
}

func (c *Client) CreateRoutingGroup(routingGroup *RoutingGroup) (*RoutingGroup, error) {
	resp, err := c.Create(routingGroupsEndpoint, *routingGroup)
	if err != nil {
		return nil, err
	}

	createdRoutingGroup, ok := resp.(*RoutingGroup)
	if !ok {
		return nil, errors.New("Object returned from API was not a RoutingGroup Pointer")
	}

	log.Printf("Returning RoutingGroup from Create: %s", createdRoutingGroup.ID)
	return createdRoutingGroup, nil
}

func (c *Client) DeleteRoutingGroup(routingGroupID string) error {
	err := c.Delete(routingGroupsEndpoint + "/" + routingGroupID)
	if err != nil {
		return err
	}

	return nil
}
