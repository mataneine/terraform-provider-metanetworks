package metanetworks

import "encoding/json"

type tag struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type tags []tag

func (c *Client) GetTags(endpoint string) (map[string]string, error) {
	var tags tags
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

func (c *Client) SetTags(endpoint string, tags map[string]string) error {

	tagsStruct := make([]tag, 0, len(tags))
	for key, value := range tags {
		tagsStruct = append(tagsStruct, tag{Name: key, Value: value})
	}

	json_data, err := json.Marshal(&tagsStruct)
	if err != nil {
		return err
	}

	_, err = c.MakeRequest(endpoint, "PUT", json_data, "application/json")
	if err != nil {
		return err
	}

	return nil
}
