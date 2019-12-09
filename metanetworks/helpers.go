package metanetworks

import (
	"encoding/json"
	"errors"
	"log"
	"reflect"
)

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

	json_data, err := json.Marshal(requestObject.Interface())
	if err != nil {
		return nil, err
	}

	resp, err := c.CreateRequest(endpoint, json_data)
	if err != nil {
		return nil, err
	}

	responseObject := reflect.New(t).Interface()
	err = json.Unmarshal(resp, &responseObject)
	id := reflect.Indirect(reflect.ValueOf(responseObject)).FieldByName("Id")

	log.Printf("Created Object with Id " + id.String())
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

func (c *Client) Read(endpoint string, o interface{}) error {
	resp, err := c.ReadRequest(endpoint)
	if err != nil {
		return err
	}

	err = json.Unmarshal(resp, o)
	if err != nil {
		return err
	}

	return nil
}

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

	json_data, err := json.Marshal(requestObject.Interface())
	if err != nil {
		return nil, err
	}

	resp, err := c.UpdateRequest(endpoint, json_data)
	if err != nil {
		return nil, err
	}

	responseObject := reflect.New(t).Interface()
	err = json.Unmarshal(resp, &responseObject)

	return responseObject, nil
}
