package metanetworks

import (
	"encoding/json"
	"errors"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const (
	metaportsEndpoint string = "/v1/metaports"
)

// MetaPort ...
type MetaPort struct {
	Name           string      `json:"name"`
	Description    string      `json:"description"`
	Enabled        bool        `json:"enabled"`
	AllowSupport   *bool       `json:"allow_support,omitempty" meta_api:"update_only"`
	MappedElements []string    `json:"mapped_elements,omitempty"`
	Connection     *Connection `json:"connection,omitempty" meta_api:"read_only"`
	CreatedAt      string      `json:"created_at,omitempty" meta_api:"read_only"`
	DNSName        string      `json:"dns_name,omitempty" meta_api:"read_only"`
	ExpiresAt      string      `json:"expires_at,omitempty" meta_api:"read_only"`
	ID             string      `json:"id,omitempty" meta_api:"read_only"`
	ModifiedAt     string      `json:"modified_at,omitempty" meta_api:"read_only"`
	OrgID          string      `json:"org_id,omitempty" meta_api:"read_only"`
}

// OTAC ...
type OTAC struct {
	ExpiresIn int64  `json:"expires_in"`
	Secret    string `json:"secret"`
}

// Connection ...
type Connection struct {
	Connected      bool   `json:"connected"`
	ConnectedAt    string `json:"connected_at"`
	DisconnectedAt string `json:"disconnected_at"`
	Location       string `json:"location"`
	VPNProto       string `json:"vpn_proto"`
}

// GetMetaPort ...
func (c *Client) GetMetaPort(metaportID string) (*MetaPort, error) {
	var metaport MetaPort
	err := c.Read(metaportsEndpoint+"/"+metaportID+"?connection=true", &metaport)
	if err != nil {
		return nil, err
	}

	log.Printf("Returning Metaport from Get: %s", metaport.ID)
	return &metaport, nil
}

// UpdateMetaPort ...
func (c *Client) UpdateMetaPort(metaportID string, metaport *MetaPort) (*MetaPort, error) {
	resp, err := c.Update(metaportsEndpoint+"/"+metaportID, *metaport)
	if err != nil {
		return nil, err
	}
	updatedMetaport, _ := resp.(*MetaPort)

	log.Printf("Returning Metaport from Update: %s", updatedMetaport.ID)
	return updatedMetaport, nil
}

// CreateMetaPort ...
func (c *Client) CreateMetaPort(metaport *MetaPort) (*MetaPort, error) {
	resp, err := c.Create(metaportsEndpoint, *metaport)
	if err != nil {
		return nil, err
	}

	createdMetaport, ok := resp.(*MetaPort)
	if !ok {
		log.Printf("Returned Type is " + reflect.TypeOf(resp).Kind().String())
		return nil, errors.New("Object returned from API was not a Metaport Pointer")
	}

	log.Printf("Returning Metaport from Create: %s", createdMetaport.ID)
	return createdMetaport, nil
}

// DeleteMetaPort ...
func (c *Client) DeleteMetaPort(metaportID string) error {
	err := c.Delete(metaportsEndpoint + "/" + metaportID)
	if err != nil {
		return err
	}

	return nil
}

// GenerateMetaPortOTAC ...
func (c *Client) GenerateMetaPortOTAC(metaportID string) (string, error) {
	resp, err := c.Request(metaportsEndpoint+"/"+metaportID+"/otac", "POST", nil, "")
	if err != nil {
		return "", err
	}

	var createdOTAC OTAC
	err = json.Unmarshal(resp, &createdOTAC)
	if err != nil {
		return "", err
	}

	return createdOTAC.Secret, nil
}

func StatusMetaportAttachmentCreate(client *Client, metaportID string, elementID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		var metaport *MetaPort
		metaport, err := client.GetMetaPort(metaportID)
		if err != nil {
			return 0, "", err
		}

		for i := 0; i < len(metaport.MappedElements); i++ {
			if metaport.MappedElements[i] == elementID {
				return metaport, "Completed", nil
			}
		}
		return metaport, "Pending", nil
	}
}

func WaitMetaportAttachmentCreate(client *Client, metaportID string, elementID string) (*Client, error) {
	createStateConf := &resource.StateChangeConf{
		Pending:    []string{"Pending"},
		Target:     []string{"Completed"},
		Timeout:    5 * time.Minute,
		MinTimeout: 5 * time.Second,
		Delay:      3 * time.Second,
		Refresh:    StatusMetaportAttachmentCreate(client, metaportID, elementID),
	}

	_, err := createStateConf.WaitForState()
	if err != nil {
		return nil, err
	}

	return client, err
}
