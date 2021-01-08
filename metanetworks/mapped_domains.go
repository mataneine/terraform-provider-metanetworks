package metanetworks

import "encoding/json"

import "log"

type MappedDomain struct {
	EnterpriseDNS bool   `json:"enterprise_dns,omitempty" type:"bool"`
	MappedDomain  string `json:"mapped_domain"`
	Name          string `json:"name,omitempty"`
}

func (c *Client) GetMappedDomain(networkElementID string, name string) (*MappedDomain, error) {
	var mappedDomain MappedDomain
	url := networkElementsEndpoint + " / " + networkElementID + " / mapped_domains / " + name
	log.Printf("URL: %s", url)
	err := c.Read(networkElementsEndpoint+"/"+networkElementID+"/mapped_domains/"+name, &mappedDomain)
	if err != nil {
		return nil, err
	}

	return &mappedDomain, nil
}

func (c *Client) SetMappedDomain(endpoint string, mappedDomain MappedDomain) (*MappedDomain, error) {
	jsonData, err := json.Marshal(mappedDomain)
	if err != nil {
		return nil, err
	}
	resp, err := c.Request(endpoint, "PUT", jsonData, "application/json")
	if err != nil {
		return nil, err
	}
	var newMappedDomain MappedDomain
	err = json.Unmarshal(resp, &newMappedDomain)
	if err != nil {
		return nil, err
	}

	return &newMappedDomain, nil
}
