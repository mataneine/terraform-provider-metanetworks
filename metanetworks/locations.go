package metanetworks

const (
	locationsEndpoint string = "/v1/locations"
)

// Location ...
type Location struct {
	City      string  `json:"city" meta_api:"read_only"`
	Country   string  `json:"country" meta_api:"read_only"`
	Latitude  float32 `json:"latitude" meta_api:"read_only"`
	Longitude float32 `json:"longitude" meta_api:"read_only"`
	Name      string  `json:"name" meta_api:"read_only"`
	State     string  `json:"state,omitempty" meta_api:"read_only"`
	Status    string  `json:"status" meta_api:"read_only"`
}

// GetLocations ...
func (c *Client) GetLocations() ([]Location, error) {
	var locations []Location
	err := c.Read(locationsEndpoint, &locations)
	if err != nil {
		return nil, err
	}

	return locations, nil
}
