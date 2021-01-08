package metanetworks

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"time"
)

// Request ...
func (c *Client) Request(endpoint, method string, data []byte, contentType string) ([]byte, error) {
	if contentType == "" {
		contentType = "application/json"
	}
	// Do we need to refresh the token? Do this first because the token might be expired, but the refresh token is ok.
	now := time.Now()
	if ((c.TokenRefreshed + c.OAUTHToken.Expiry) - now.Unix()) < 30 {
		err := c.RefreshToken()
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, baseURL+endpoint, bytes.NewReader(data))

	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	req.Header.Add("Authorization", "Bearer "+c.OAUTHToken.Token)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(string(body))
	}

	log.Printf("Response: %s", body)
	return body, nil
}

// Create ...
func (c *Client) Create(endpoint string, o interface{}) (interface{}, error) {
	v := reflect.ValueOf(o)
	t := reflect.TypeOf(o)
	if t.Kind() != reflect.Struct {
		return nil, errors.New("Tried to create with a " + t.Kind().String() + " not a Struct")
	}

	requestObject := reflect.Indirect(reflect.New(t))
	needsUpdate := false

	for i := 0; i < t.NumField(); i++ {

		// this is a field struct
		structField := t.Field(i)
		// these are reflect values
		requestField := requestObject.FieldByName(structField.Name)
		valueField := v.FieldByName(structField.Name)

		switch n := structField.Tag.Get("meta_api"); n {
		case "read_only":
			continue
		case "update_only":
			needsUpdate = true
		default:
			log.Printf("Field Name " + structField.Name + " with " + valueField.String())
			if requestField.CanSet() {
				requestField.Set(valueField)
			}
		}
	}

	data, err := json.Marshal(requestObject.Interface())
	if err != nil {
		return nil, err
	}

	resp, err := c.Request(endpoint, "POST", data, "application/json")
	if err != nil {
		return nil, err
	}

	responseObject := reflect.New(t).Interface()
	err = json.Unmarshal(resp, &responseObject)
	id := reflect.Indirect(reflect.ValueOf(responseObject)).FieldByName("ID")

	log.Printf("Created Object with ID " + id.String())
	// we need to update at a different endpoint, luckily these follow a pattern of sticking the id on the end.
	if needsUpdate {
		responseObject, err = c.Update(endpoint+"/"+id.String(), o)
		if err != nil {
			return nil, err
		}
		return responseObject, nil
	}

	return responseObject, nil
}

// Read ...
func (c *Client) Read(endpoint string, o interface{}) error {
	resp, err := c.Request(endpoint, "GET", nil, "application/json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(resp, o)
	if err != nil {
		return err
	}

	return nil
}

// Update ...
func (c *Client) Update(endpoint string, o interface{}) (interface{}, error) {
	v := reflect.ValueOf(o)
	t := reflect.TypeOf(o)
	if t.Kind() != reflect.Struct {
		return nil, errors.New("Tried to update with a " + t.Kind().String() + " not a Struct")
	}

	requestObject := reflect.Indirect(reflect.New(t))
	for i := 0; i < t.NumField(); i++ {

		// this is a field struct
		structField := t.Field(i)
		// these are reflect values
		requestField := requestObject.FieldByName(structField.Name)
		valueField := v.FieldByName(structField.Name)

		switch n := structField.Tag.Get("meta_api"); n {
		case "read_only":
			continue
		case "update_only":
			log.Printf("Update Field Name " + structField.Name + " with " + valueField.String())
			if requestField.CanSet() {
				requestField.Set(valueField)
			}
		default:
			log.Printf("Field Name " + structField.Name + " with " + valueField.String())
			if requestField.CanSet() {
				requestField.Set(valueField)
			}
		}
	}

	data, err := json.Marshal(requestObject.Interface())
	if err != nil {
		return nil, err
	}

	resp, err := c.Request(endpoint, "PATCH", data, "application/merge-patch+json")
	if err != nil {
		return nil, err
	}

	responseObject := reflect.New(t).Interface()
	err = json.Unmarshal(resp, &responseObject)

	return responseObject, nil
}

// Delete ...
func (c *Client) Delete(endpoint string) error {
	_, err := c.Request(endpoint, "DELETE", nil, "application/json")
	if err != nil {
		return err
	}
	return nil
}
