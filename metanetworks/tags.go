package metanetworks

import (
	"encoding/json"
)

// Tag ...
type Tag struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// GetTags ...
func (c *Client) GetTags(endpoint string) (map[string]string, error) {
	var tags []Tag
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

// UpdateTags ...
func (c *Client) UpdateTags(endpoint string, tags map[string]string) error {
	tagsStruct := make([]Tag, 0, len(tags))
	for key, value := range tags {
		tagsStruct = append(tagsStruct, Tag{Name: key, Value: value})
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
