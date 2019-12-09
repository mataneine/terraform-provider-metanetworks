package metanetworks

import "encoding/json"

const (
	ProtocolGroupEndpoint string = "/v1/protocol_groups"
)

type ProtocolGroup struct {
	Description string     `json:"description,omitempty"`
	Name        string     `json:"name"`
	Protocols   []Protocol `json:"protocols,omitempty"`
	CreatedAt   string     `json:"created_at,omitempty"`
	Id          string     `json:"id,omitempty"`
	ModifiedAt  string     `json:"modified_at,omitempty"`
	OrgId       string     `json:"org_id,omitempty"`
	ReadOnly    bool       `json:"read_only,omitempty"`
}

type Protocol struct {
	Port     uint16 `json:"port"`
	Protocol string `json:"protocol"`
}

func (c *Client) GetProtocolGroup(protocol_group_id string) (*ProtocolGroup, error) {

	resp, err := c.MakeRequest(ProtocolGroupEndpoint+"/"+protocol_group_id, "GET", nil, "")
	if err != nil {
		return nil, err
	}

	var protocol_group ProtocolGroup
	err = json.Unmarshal(resp, &protocol_group)
	if err != nil {
		return nil, err
	}

	return &protocol_group, nil
}

func (c *Client) UpdateProtocolGroup(protocol_group_id string, protocol_group *ProtocolGroup) (*ProtocolGroup, error) {

	// can't change these so the api throws an error when you try and patch them
	protocol_group.Id = ""
	protocol_group.CreatedAt = ""
	protocol_group.OrgId = ""
	protocol_group.ModifiedAt = ""

	json_data, err := json.Marshal(protocol_group)
	if err != nil {
		return nil, err
	}
	resp, err := c.MakeRequest(ProtocolGroupEndpoint+"/"+protocol_group_id, "PATCH", json_data, "application/merge-patch+json")
	if err != nil {
		return nil, err
	}

	var updated_protocol_group ProtocolGroup
	err = json.Unmarshal(resp, &updated_protocol_group)
	if err != nil {
		return nil, err
	}

	return &updated_protocol_group, nil

}

func (c *Client) CreateProtocolGroup(protocol_group *ProtocolGroup) (*ProtocolGroup, error) {

	json_data, err := json.Marshal(protocol_group)
	if err != nil {
		return nil, err
	}
	resp, err := c.MakeRequest(ProtocolGroupEndpoint, "POST", json_data, "application/json")
	if err != nil {
		return nil, err
	}

	var created_protocol_group ProtocolGroup
	err = json.Unmarshal(resp, &created_protocol_group)
	if err != nil {
		return nil, err
	}

	return &created_protocol_group, nil

}

func (c *Client) DeleteProtocolGroup(protocol_group_id string) error {
	_, err := c.MakeRequest(ProtocolGroupEndpoint+"/"+protocol_group_id, "DELETE", nil, "")
	if err != nil {
		return err
	}

	return nil
}
