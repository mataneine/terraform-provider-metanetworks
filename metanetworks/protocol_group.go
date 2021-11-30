package metanetworks

import (
	"errors"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const (
	protocolGroupsEndpoint string = "/v1/protocol_groups"
)

// ProtocolGroup ...
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

// Protocol ...
type Protocol struct {
	FromPort int64  `json:"from_port" type:"integer"`
	ToPort   int64  `json:"to_port" type:"integer"`
	Protocol string `json:"proto"`
}

// GetProtocolGroup ...
func (c *Client) GetProtocolGroup(protocolGroupID string) (*ProtocolGroup, error) {
	var protocolGroup ProtocolGroup
	err := c.Read(protocolGroupsEndpoint+"/"+protocolGroupID, &protocolGroup)
	if err != nil {
		return nil, err
	}

	log.Printf("Returning ProtocolGroup from Get: %s", protocolGroup)
	return &protocolGroup, nil
}

// UpdateProtocolGroup ...
func (c *Client) UpdateProtocolGroup(protocolGroupID string, protocolGroup *ProtocolGroup) (*ProtocolGroup, error) {
	resp, err := c.Update(protocolGroupsEndpoint+"/"+protocolGroupID, *protocolGroup)
	if err != nil {
		return nil, err
	}
	updatedProtocolGroup, _ := resp.(*ProtocolGroup)

	log.Printf("Returning ProtocolGroup from Update: %s", updatedProtocolGroup.ID)
	return updatedProtocolGroup, nil
}

// CreateProtocolGroup ...
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

// DeleteProtocolGroup ...
func (c *Client) DeleteProtocolGroup(protocolGroupID string) error {
	err := c.Delete(protocolGroupsEndpoint + "/" + protocolGroupID)
	if err != nil {
		return err
	}

	return nil
}

func StatusProtocolGroupCreate(client *Client, ProtocolGroupID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		var metaport *MetaPort
		_, err := client.GetProtocolGroup(ProtocolGroupID)
		if err != nil {
			return metaport, "Pending", nil
		}
		return metaport, "Completed", nil
	}
}

func WaitProtocolGroupCreate(client *Client, ProtocolGroupID string) (*Client, error) {
	createStateConf := &resource.StateChangeConf{
		Pending:    []string{"Pending"},
		Target:     []string{"Completed"},
		Timeout:    30 * time.Second,
		MinTimeout: 5 * time.Second,
		Delay:      2 * time.Second,
		Refresh:    StatusProtocolGroupCreate(client, ProtocolGroupID),
	}

	_, err := createStateConf.WaitForState()
	if err != nil {
		return nil, err
	}

	return client, err
}
