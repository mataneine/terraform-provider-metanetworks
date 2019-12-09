package metanetworks

import (
	"encoding/json"
	"errors"
	"log"
	"reflect"
)

const (
	MetaPortEndpoint string = "/v1/metaports"
)

type MetaPort struct {
	Name           string   `json:"name"`
	Description    string   `json:"description"`
	Enabled        bool     `json:"enabled" meta_api:"update_only"`
	MappedElements []string `json:"mapped_elements"`

	Connection *Connection `json:"connection,omitempty" meta_api:"read_only"`

	AllowSupport bool `json:"allow_support" meta_api:"update_only"`

	CreatedAt        string `json:"created_at,omitempty" meta_api:"read_only"`
	DNSName          string `json:"dns_name,omitempty" meta_api:"read_only"`
	ExpiresAt        string `json:"expires_at,omitempty" meta_api:"read_only"`
	Id               string `json:"id,omitempty" meta_api:"read_only"`
	ModifiedAt       string `json:"modified_at,omitempty" meta_api:"read_only"`
	NetworkElementId string `json:"network_element_id,omitempty" meta_api:"read_only"`
	OrgId            string `json:"org_id,omitempty" meta_api:"read_only"`
}

type OTAC struct {
	ExpiresIn int64  `json:"expires_in"`
	Secret    string `json:"secret"`
}

type Connection struct {
	Connected      bool   `json:"connected"`
	ConnectedAt    string `json:"connected_at"`
	DisconnectedAt string `json:"disconnected_at"`
	Location       string `json:"location"`
	VPNProto       string `json:"vpn_proto"`
}

func (c *Client) GetMetaPort(metaport_id string) (*MetaPort, error) {

	var metaport MetaPort
	err := c.Read(MetaPortEndpoint+"/"+metaport_id+"?connection=true", &metaport)
	if err != nil {
		return nil, err
	}

	log.Printf("Returning Metaport from Get " + string(metaport.Id))
	return &metaport, nil

}

func (c *Client) UpdateMetaPort(metaportId string, metaport *MetaPort) (*MetaPort, error) {

	resp, err := c.Update(MetaPortEndpoint+"/"+metaportId, *metaport)
	if err != nil {
		return nil, err
	}
	updatedMetaport, _ := resp.(*MetaPort)

	log.Printf("Returning Metaport from Update " + string(updatedMetaport.Id))
	return updatedMetaport, nil
}

func (c *Client) CreateMetaPort(metaport *MetaPort) (*MetaPort, error) {
	resp, err := c.Create(MetaPortEndpoint, *metaport)
	if err != nil {
		return nil, err
	}

	createdMetaport, ok := resp.(*MetaPort)
	if !ok {
		log.Printf("Returned Type is " + reflect.TypeOf(resp).Kind().String())
		return nil, errors.New("Object returned from API was not a Metaport Pointer")
	}

	log.Printf("Returning Metaport from Create " + string(createdMetaport.Id))
	return createdMetaport, nil

}

func (c *Client) DeleteMetaPort(metaport_id string) error {
	err := c.DeleteRequest(MetaPortEndpoint + "/" + metaport_id)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) GenerateMetaPortOTAC(metaport_id string) (string, error) {
	resp, err := c.MakeRequest(MetaPortEndpoint+"/"+metaport_id+"/otac", "POST", nil, "")
	if err != nil {
		return "", err
	}

	var created_otac OTAC
	err = json.Unmarshal(resp, &created_otac)
	if err != nil {
		return "", err
	}

	return created_otac.Secret, nil
}
