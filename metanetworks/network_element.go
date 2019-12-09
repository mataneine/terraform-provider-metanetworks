package metanetworks

import (
	"encoding/json"
	"errors"
	"log"
)

const (
	NetworkElementEndpoint string = "/v1/network_elements"
)

type NetworkElement struct {
	Aliases     []string `json:"aliases,omitempty" meta_api:"read_only"`
	CreatedAt   string   `json:"created_at,omitempty" meta_api:"read_only"`
	Description string   `json:"description"`
	DNSName     string   `json:"dns_name,omitempty" meta_api:"read_only"`
	// Disable/enable is not supported for mapped services.
	Enabled       bool     `json:"enabled,omitempty" meta_api:"read_only"`
	ExpiresAt     string   `json:"expires_at,omitempty" meta_api:"read_only"`
	Id            string   `json:"id,omitempty" meta_api:"read_only"`
	MappedService string   `json:"mapped_service,omitempty"`
	MappedSubnets []string `json:"mapped_subnets,omitempty"`
	ModifiedAt    string   `json:"modified_at,omitempty" meta_api:"read_only"`
	Name          string   `json:"name"`
	NetId         int64    `json:"net_id,omitempty" meta_api:"read_only"`
	OrgId         string   `json:"org_id,omitempty" meta_api:"read_only"`
	OwnerId       string   `json:"owner_id,omitempty" meta_api:"read_only"`
	Platform      string   `json:"platform,omitempty" meta_api:"read_only"`
	Type          string   `json:"type,omitempty" meta_api:"read_only"`
	Version       int64    `json:"version,omitempty" meta_api:"read_only"`
}

func (c *Client) GetNetworkElement(element_id string) (*NetworkElement, error) {

	var networkElement NetworkElement
	err := c.Read(NetworkElementEndpoint+"/"+element_id, &networkElement)
	if err != nil {
		return nil, err
	}

	log.Printf("Returning Network Element from Get " + string(networkElement.Id))
	return &networkElement, nil
}

func (c *Client) UpdateNetworkElement(network_element_id string, networkElement *NetworkElement) (*NetworkElement, error) {

	resp, err := c.Update(NetworkElementEndpoint+"/"+network_element_id, *networkElement)
	if err != nil {
		return nil, err
	}
	updatedNetworkElement, _ := resp.(*NetworkElement)

	log.Printf("Returning Network Element from Update " + string(updatedNetworkElement.Id))
	return updatedNetworkElement, nil

}

func (c *Client) CreateNetworkElement(networkElement *NetworkElement) (*NetworkElement, error) {

	resp, err := c.Create(NetworkElementEndpoint, *networkElement)
	if err != nil {
		return nil, err
	}

	createdNetworkElement, ok := resp.(*NetworkElement)
	if !ok {
		return nil, errors.New("Object returned from API was not a Network Element Pointer")
	}

	log.Printf("Returning Network Element from Create " + string(createdNetworkElement.Id))
	return createdNetworkElement, nil

}

func (c *Client) DeleteNetworkElement(network_element_id string) error {
	err := c.DeleteRequest(NetworkElementEndpoint + "/" + network_element_id)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) AddNetworkElementAlias(network_element_id string, alias string) (*NetworkElement, error) {

	resp, err := c.MakeRequest(NetworkElementEndpoint+"/"+network_element_id+"/aliases/"+alias, "PUT", nil, "")
	if err != nil {
		return nil, err
	}
	var network_element NetworkElement
	err = json.Unmarshal(resp, &network_element)
	if err != nil {
		return nil, err
	}

	return &network_element, nil
}

func (c *Client) RemoveNetworkElementAlias(network_element_id string, alias string) (*NetworkElement, error) {
	resp, err := c.MakeRequest(NetworkElementEndpoint+"/"+network_element_id+"/aliases/"+alias, "DELETE", nil, "")
	if err != nil {
		return nil, err
	}
	var network_element NetworkElement
	err = json.Unmarshal(resp, &network_element)
	if err != nil {
		return nil, err
	}
	return &network_element, nil
}

func (c *Client) GetNetworkElementTags(network_element_id string) (map[string]string, error) {
	tagMap, err := c.GetTags(NetworkElementEndpoint + "/" + network_element_id + "/tags")
	if err != nil {
		return nil, err
	}

	return tagMap, nil
}

func (c *Client) SetNetworkElementTags(network_element_id string, tags map[string]string) error {
	err := c.SetTags(NetworkElementEndpoint+"/"+network_element_id+"/tags", tags)
	if err != nil {
		return err
	}
	return nil
}
