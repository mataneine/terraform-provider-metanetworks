package metanetworks

const (
	UserEndpoint string = "/v1/users"
)

type User struct {
	Description       string              `json:"description,omitempty"`
	Email             string              `json:"email"`
	Enabled           bool                `json:"enabled"`
	FamilyName        string              `json:"family_name"`
	GivenName         string              `json:"given_name"`
	Phone             string              `json:"phone,omitempty"`
	ProvisionedBy     string              `json:"provisioned_by,omitempty"`
	CreatedAt         string              `json:"created_at,omitempty" meta_api:"read_only"`
	Id                string              `json:"id,omitempty" meta_api:"read_only"`
	Inventory         []string            `json:"inventory,omitempty" meta_api:"read_only"`
	MFAEnabled        bool                `json:"mfa_enabled,omitempty" meta_api:"read_only"`
	ModifiedAt        string              `json:"modified_at,omitempty" meta_api:"read_only"`
	Name              string              `json:"name,omitempty"`
	OrgId             string              `json:"org_id,omitempty" meta_api:"read_only"`
	OverlayMFAEnabled bool                `json:"overlay_mfa_enabled,omitempty"`
	PhoneVerified     bool                `json:"phone_verified,omitempty"`
	Roles             []string            `json:"roles,omitempty" meta_api:"read_only"`
	Tags              []map[string]string `json:"tags,omitempty" meta_api:"read_only"`
}

func (c *Client) GetUser(user_id string) (*User, error) {

	var user User
	err := c.Read(UserEndpoint+"/"+user_id, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (c *Client) UpdateUser(user_id string, user *User) (*User, error) {

	resp, err := c.Update(UserEndpoint+"/"+user_id, *user)
	if err != nil {
		return nil, err
	}
	updated_user, _ := resp.(*User)

	return updated_user, nil

}

func (c *Client) CreateUser(user *User) (*User, error) {

	resp, err := c.Create(UserEndpoint, *user)
	if err != nil {
		return nil, err
	}
	created_user, _ := resp.(*User)

	return created_user, nil

}

func (c *Client) DeleteUser(user_id string) error {
	err := c.DeleteRequest(UserEndpoint + "/" + user_id)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) ListUsers() ([]User, error) {

	var userList []User
	err := c.Read(UserEndpoint, &userList)
	if err != nil {
		return nil, err
	}

	return userList, nil
}

func (c *Client) GetUserTags(user_id string) (map[string]string, error) {
	tagMap, err := c.GetTags(UserEndpoint + "/" + user_id + "/tags")
	if err != nil {
		return nil, err
	}

	return tagMap, nil
}

func (c *Client) SetUserTags(user_id string, tags map[string]string) error {
	err := c.SetTags(UserEndpoint+"/"+user_id+"/tags", tags)
	if err != nil {
		return err
	}
	return nil
}
