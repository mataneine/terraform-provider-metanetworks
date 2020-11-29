package metanetworks

import (
	"errors"
	"log"
)

const (
	peeringsEndpoint string = "/v1/peerings"
)

// Peering ...
type Peering struct {
	Description string   `json:"description"`
	EgressNAT   bool     `json:"egress_nat"`
	Enabled     bool     `json:"enabled"`
	Name        string   `json:"name"`
	Peers       []string `json:"peers,omitempty"`
	CreatedAt   string   `json:"created_at,omitempty" meta_api:"read_only"`
	ID          string   `json:"id,omitempty" meta_api:"read_only"`
	ModifiedAt  string   `json:"modified_at,omitempty" meta_api:"read_only"`
	OrgID       string   `json:"org_id,omitempty" meta_api:"read_only"`
}

// GetPeering ...
func (c *Client) GetPeering(peeringID string) (*Peering, error) {
	var peering Peering
	err := c.Read(peeringsEndpoint+"/"+peeringID, &peering)
	if err != nil {
		return nil, err
	}

	log.Printf("Returning Peering from Get: %s", peering.ID)
	return &peering, nil
}

// UpdatePeering ...
func (c *Client) UpdatePeering(peeringID string, peering *Peering) (*Peering, error) {
	resp, err := c.Update(peeringsEndpoint+"/"+peeringID, *peering)
	if err != nil {
		return nil, err
	}
	updatedPeering, _ := resp.(*Peering)

	log.Printf("Returning Peering from Update: %s", updatedPeering.ID)
	return updatedPeering, nil
}

// CreatePeering ...
func (c *Client) CreatePeering(peering *Peering) (*Peering, error) {
	resp, err := c.Create(peeringsEndpoint, *peering)
	if err != nil {
		return nil, err
	}

	createdPeering, ok := resp.(*Peering)
	if !ok {
		return nil, errors.New("Object returned from API was not a Peering Pointer")
	}

	log.Printf("Returning Peering from Create: %s", createdPeering.ID)
	return createdPeering, nil
}

// DeletePeering ...
func (c *Client) DeletePeering(peeringID string) error {
	err := c.Delete(peeringsEndpoint + "/" + peeringID)
	if err != nil {
		return err
	}

	return nil
}
