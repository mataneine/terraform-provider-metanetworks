package metanetworks

import (
	"encoding/json"
)

type tag struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (c *Client) GetTags(endpoint string) (map[string]string, error) {
	var tags []tag
	err := c.Read(endpoint, &tags)
	if err != nil {
		return nil, err
	}

	tagMap := make(map[string]string)
	for _, tag := range tags {
		tagMap[tag.Name] = tag.Value
	}

	return tagMap, nil
}

func (c *Client) UpdateTags(endpoint string, tags map[string]string) error {
	tagsStruct := make([]tag, 0, len(tags))
	for key, value := range tags {
		tagsStruct = append(tagsStruct, tag{Name: key, Value: value})
	}

	jsonData, err := json.Marshal(&tagsStruct)
	if err != nil {
		return err
	}

	_, err = c.Request(endpoint, "PUT", jsonData, "application/json")
	if err != nil {
		return err
	}

	return nil
}
