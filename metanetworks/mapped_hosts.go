package metanetworks

import (
	"encoding/json"
	"log"
)

// MappedHost ...
type MappedHost struct {
	IgnoreBounds bool   `json:"ignore_bounds,omitempty" type:"bool"`
	MappedHost   string `json:"mapped_host"`
	Name         string `json:"name,omitempty"`
}

// GetMappedHost ...
func (c *Client) GetMappedHost(networkElementID string, name string) (*MappedHost, error) {
	var mappedHost MappedHost
	url := networkElementsEndpoint + " / " + networkElementID + " / mapped_hosts / " + name
	log.Printf("URL: %s", url)
	err := c.Read(networkElementsEndpoint+"/"+networkElementID+"/mapped_hosts/"+name, &mappedHost)
	if err != nil {
		return nil, err
	}

	return &mappedHost, nil
}

// SetMappedHost ...
func (c *Client) SetMappedHost(endpoint string, mappedHost MappedHost) (*MappedHost, error) {
	jsonData, err := json.Marshal(mappedHost)
	if err != nil {
		return nil, err
	}
	resp, err := c.Request(endpoint, "PUT", jsonData, "application/json")
	if err != nil {
		return nil, err
	}
	var newMappedHost MappedHost
	err = json.Unmarshal(resp, &newMappedHost)
	if err != nil {
		return nil, err
	}

	return &newMappedHost, nil
}
