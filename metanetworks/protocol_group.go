package metanetworks

import (
	"errors"
	"log"
	"reflect"
)

const (
	protocolGroupsEndpoint string = "/v1/protocol_groups"
)

type ProtocolGroup struct {
	Description string     `json:"description"`
	Name        string     `json:"name"`
	Protocols   []Protocol `json:"protocols,omitempty"`
	CreatedAt   string     `json:"created_at,omitempty"`
	ID          string     `json:"id,omitempty"`
	ModifiedAt  string     `json:"modified_at,omitempty"`
	OrgID       string     `json:"org_id,omitempty"`
	ReadOnly    bool       `json:"read_only,omitempty"`
}

type Protocol struct {
	FromPort int64  `json:"from_port" type:"integer"`
	ToPort   int64  `json:"to_port" type:"integer"`
	Protocol string `json:"proto"`
}

func (c *Client) GetProtocolGroup(protocolGroupID string) (*ProtocolGroup, error) {
	var protocolGroup ProtocolGroup
	err := c.Read(protocolGroupsEndpoint+"/"+protocolGroupID, &protocolGroup)
	if err != nil {
		return nil, err
	}

	log.Printf("Returning ProtocolGroup from Get: %s", protocolGroup)
	return &protocolGroup, nil
}

func (c *Client) UpdateProtocolGroup(protocolGroupID string, protocolGroup *ProtocolGroup) (*ProtocolGroup, error) {
	resp, err := c.Update(protocolGroupsEndpoint+"/"+protocolGroupID, *protocolGroup)
	if err != nil {
		return nil, err
	}
	updatedProtocolGroup, _ := resp.(*ProtocolGroup)

	log.Printf("Returning ProtocolGroup from Update: %s", updatedProtocolGroup.ID)
	return updatedProtocolGroup, nil
}

func (c *Client) CreateProtocolGroup(protocolGroup *ProtocolGroup) (*ProtocolGroup, error) {
	resp, err := c.Create(protocolGroupsEndpoint, *protocolGroup)
	if err != nil {
		return nil, err
	}

	createdProtocolGroup, ok := resp.(*ProtocolGroup)
	if !ok {
		log.Printf("Returned Type is " + reflect.TypeOf(resp).Kind().String())
		return nil, errors.New("Object returned from API was not a ProtocolGroup Pointer")
	}

	log.Printf("Returning ProtocolGroup from Create: %s", createdProtocolGroup.ID)
	return createdProtocolGroup, nil
}

func (c *Client) DeleteProtocolGroup(protocolGroupID string) error {
	err := c.Delete(protocolGroupsEndpoint + "/" + protocolGroupID)
	if err != nil {
		return err
	}

	return nil
}
