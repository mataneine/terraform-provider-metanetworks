package metanetworks

import (
	"errors"
	"log"
)

const (
	postureCheckEndpoint string = "/v1/posture_checks"
)

// PostureCheck Struct ...
type PostureCheck struct {
	Description       string        `json:"description,omitempty"`
	Name              string        `json:"name"`
	Action            string        `json:"action"`
	OSQuery           string        `json:"osquery,omitempty"`
	Platform          string        `json:"platform"`
	UserMessageOnFail string        `json:"user_message_on_fail,omitempty"`
	Enabled           bool          `json:"enabled" type:"bool"`
	ApplyToOrg        bool          `json:"apply_to_org"`
	Interval          int           `json:"interval,omitempty"`
	Check             []interface{} `json:"check,omitempty"`
	When              []string      `json:"when"`
	ExemptEntities    []string      `json:"exempt_entities,omitempty"`
	ApplyToEntities   []string      `json:"apply_to_entities,omitempty"`
	CreatedAt         string        `json:"created_at,omitempty" meta_api:"read_only"`
	ID                string        `json:"id,omitempty" meta_api:"read_only"`
	ModifiedAt        string        `json:"modified_at,omitempty" meta_api:"read_only"`
}

// GetPostureCheck ...
func (c *Client) GetPostureCheck(postureCheckID string) (*PostureCheck, error) {
	var postureCheck PostureCheck
	err := c.Read(postureCheckEndpoint+"/"+postureCheckID, &postureCheck)
	if err != nil {
		return nil, err
	}

	log.Printf("Returning Posture Check from Get: %s", postureCheck.ID)
	return &postureCheck, nil
}

// UpdatePostureCheck ...
func (c *Client) UpdatePostureCheck(postureCheckID string, postureCheck *PostureCheck) (*PostureCheck, error) {
	resp, err := c.Update(postureCheckEndpoint+"/"+postureCheckID, *postureCheck)
	if err != nil {
		return nil, err
	}
	updatedPostureCheck, _ := resp.(*PostureCheck)

	log.Printf("Returning Posture Check from Update: %s", updatedPostureCheck.ID)
	return updatedPostureCheck, nil
}

// CreatePostureCheck ...
func (c *Client) CreatePostureCheck(postureCheck *PostureCheck) (*PostureCheck, error) {
	resp, err := c.Create(postureCheckEndpoint, *postureCheck)
	if err != nil {
		return nil, err
	}

	createdPostureCheck, ok := resp.(*PostureCheck)
	if !ok {
		return nil, errors.New("Object returned from API was not a Posture Check Pointer")
	}

	log.Printf("Returning Posture Check from Create: %s", createdPostureCheck.ID)
	return createdPostureCheck, nil
}

// DeletePostureCheck ...
func (c *Client) DeletePostureCheck(postureCheckID string) error {
	err := c.Delete(postureCheckEndpoint + "/" + postureCheckID)
	if err != nil {
		return err
	}

	return nil
}
