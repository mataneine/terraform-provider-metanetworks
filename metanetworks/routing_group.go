package metanetworks

import (
	"errors"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const (
	routingGroupsEndpoint string = "/v1/routing_groups"
)

// RoutingGroup ...
type RoutingGroup struct {
	Name           string   `json:"name"`
	Description    string   `json:"description"`
	MappedElements []string `json:"mapped_elements_ids,omitempty"`
	ExemptSources  []string `json:"exempt_sources,omitempty"`
	Sources        []string `json:"sources,omitempty"`
	CreatedAt      string   `json:"created_at,omitempty" meta_api:"read_only"`
	ID             string   `json:"id,omitempty" meta_api:"read_only"`
	ModifiedAt     string   `json:"modified_at,omitempty" meta_api:"read_only"`
	OrgID          string   `json:"org_id,omitempty" meta_api:"read_only"`
	Priority       int      `json:"priority,omitempty" meta_api:"update_only"`
}

// GetRoutingGroup ...
func (c *Client) GetRoutingGroup(routingGroupID string) (*RoutingGroup, error) {
	var routingGroup RoutingGroup
	err := c.Read(routingGroupsEndpoint+"/"+routingGroupID, &routingGroup)
	if err != nil {
		return nil, err
	}

	log.Printf("Returning RoutingGroup from Get: %s", routingGroup.ID)
	return &routingGroup, nil
}

// UpdateRoutingGroup ...
func (c *Client) UpdateRoutingGroup(routingGroupID string, routingGroup *RoutingGroup) (*RoutingGroup, error) {
	resp, err := c.Update(routingGroupsEndpoint+"/"+routingGroupID, *routingGroup)
	if err != nil {
		return nil, err
	}
	updatedRoutingGroup, _ := resp.(*RoutingGroup)

	log.Printf("Returning RoutingGroup from Update: %s", updatedRoutingGroup.ID)
	return updatedRoutingGroup, nil
}

// CreateRoutingGroup ...
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

//DeleteRoutingGroup ...
func (c *Client) DeleteRoutingGroup(routingGroupID string) error {
	err := c.Delete(routingGroupsEndpoint + "/" + routingGroupID)
	if err != nil {
		return err
	}

	return nil
}

func StatusRoutingGroupAttachmentCreate(client *Client, routingGroupID string, elementID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		var routingGroup *RoutingGroup
		routingGroup, err := client.GetRoutingGroup(routingGroupID)
		if err != nil {
			return 0, "", err
		}

		for i := 0; i < len(routingGroup.MappedElements); i++ {
			if routingGroup.MappedElements[i] == elementID {
				return routingGroup, "Completed", nil
			}
		}
		return routingGroup, "Pending", nil
	}
}

func WaitRoutingGroupAttachmentCreate(client *Client, routingGroupID string, elementID string) (*Client, error) {
	createStateConf := &resource.StateChangeConf{
		Pending:    []string{"Pending"},
		Target:     []string{"Completed"},
		Timeout:    5 * time.Minute,
		MinTimeout: 5 * time.Second,
		Delay:      3 * time.Second,
		Refresh:    StatusRoutingGroupAttachmentCreate(client, routingGroupID, elementID),
	}

	_, err := createStateConf.WaitForState()
	if err != nil {
		return nil, err
	}

	return client, err
}
