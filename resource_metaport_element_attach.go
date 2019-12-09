package main

import (
	"errors"
	"terraform-provider-metanetworks/metanetworks"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceMetaportElementAttach() *schema.Resource {
	return &schema.Resource{
		Create: resourceMetaportElementAttachCreate,
		Read:   resourceMetaportElementAttachRead,
		Delete: resourceMetaportElementAttachDelete,

		Schema: map[string]*schema.Schema{
			"metaport_id": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The ID of the MetaPort",
				Required:    true,
				ForceNew:    true,
			},
			"network_element_id": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The ID of the Network Element (Mapped Service etc)",
				Required:    true,
				ForceNew:    true,
			},
		},
	}
}

func resourceMetaportElementAttachCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*metanetworks.Client)

	element_id := d.Get("network_element_id").(string)
	metaport_id := d.Get("metaport_id").(string)

	mutex := metaportGetLock(d.Id())
	defer mutex.Unlock()

	var metaport *metanetworks.MetaPort
	metaport, err := client.GetMetaPort(metaport_id)
	if err != nil {
		return err
	}

	for i := 0; i < len(metaport.MappedElements); i++ {
		if metaport.MappedElements[i] == element_id {
			return errors.New("That network element is already mapped to this MetaPort")
		}

	}

	metaport.MappedElements = append(metaport.MappedElements, element_id)
	_, err = client.UpdateMetaPort(metaport_id, metaport)
	if err != nil {
		return err
	}

	d.SetId(element_id + metaport_id)

	return nil
}

func resourceMetaportElementAttachRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*metanetworks.Client)

	element_id := d.Get("network_element_id").(string)
	metaport_id := d.Get("metaport_id").(string)

	mutex := metaportGetLock(d.Id())
	defer mutex.Unlock()

	var metaport *metanetworks.MetaPort
	metaport, err := client.GetMetaPort(metaport_id)
	if err != nil {
		return err
	}

	found := false
	for i := 0; i < len(metaport.MappedElements); i++ {
		if metaport.MappedElements[i] == element_id {
			found = true
			break
		}

	}

	// If not present we need to
	// destroy the terraform resource so that it is recreated.
	if !found {
		d.SetId("")
	}

	return nil
}

func resourceMetaportElementAttachDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*metanetworks.Client)

	element_id := d.Get("network_element_id").(string)
	metaport_id := d.Get("metaport_id").(string)

	mutex := metaportGetLock(d.Id())
	defer mutex.Unlock()

	var metaport *metanetworks.MetaPort
	metaport, err := client.GetMetaPort(metaport_id)
	if err != nil {
		return err
	}

	// Note that if the entry has already been deleted this won't fail.
	for i := 0; i < len(metaport.MappedElements); i++ {
		if metaport.MappedElements[i] == element_id {
			metaport.MappedElements = append(metaport.MappedElements[i:], metaport.MappedElements[i+1:]...)
			break
		}
	}

	_, err = client.UpdateMetaPort(metaport_id, metaport)
	if err != nil {
		return err
	}

	return nil
}
