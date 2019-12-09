package main

import (
	"errors"
	"terraform-provider-metanetworks/metanetworks"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceMetaportElementAttach() *schema.Resource {
	return &schema.Resource{
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
		Create: resourceMetaportElementAttachCreate,
		Read:   resourceMetaportElementAttachRead,
		Delete: resourceMetaportElementAttachDelete,
	}
}

func resourceMetaportElementAttachCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*metanetworks.Client)

	elementID := d.Get("network_element_id").(string)
	metaportID := d.Get("metaport_id").(string)

	mutex := metaportGetLock(d.Id())
	defer mutex.Unlock()

	var metaport *metanetworks.MetaPort
	metaport, err := client.GetMetaPort(metaportID)
	if err != nil {
		return err
	}

	for i := 0; i < len(metaport.MappedElements); i++ {
		if metaport.MappedElements[i] == elementID {
			return errors.New("That network element is already mapped to this MetaPort")
		}

	}

	metaport.MappedElements = append(metaport.MappedElements, elementID)
	_, err = client.UpdateMetaPort(metaportID, metaport)
	if err != nil {
		return err
	}

	d.SetId(elementID + metaportID)

	return nil
}

func resourceMetaportElementAttachRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*metanetworks.Client)

	elementID := d.Get("network_element_id").(string)
	metaportID := d.Get("metaport_id").(string)

	mutex := metaportGetLock(d.Id())
	defer mutex.Unlock()

	var metaport *metanetworks.MetaPort
	metaport, err := client.GetMetaPort(metaportID)
	if err != nil {
		return err
	}

	found := false
	for i := 0; i < len(metaport.MappedElements); i++ {
		if metaport.MappedElements[i] == elementID {
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

	elementID := d.Get("network_element_id").(string)
	metaportID := d.Get("metaport_id").(string)

	mutex := metaportGetLock(d.Id())
	defer mutex.Unlock()

	var metaport *metanetworks.MetaPort
	metaport, err := client.GetMetaPort(metaportID)
	if err != nil {
		return err
	}

	// Note that if the entry has already been deleted this won't fail.
	for i := 0; i < len(metaport.MappedElements); i++ {
		if metaport.MappedElements[i] == elementID {
			metaport.MappedElements = append(metaport.MappedElements[i:], metaport.MappedElements[i+1:]...)
			break
		}
	}

	_, err = client.UpdateMetaPort(metaportID, metaport)
	if err != nil {
		return err
	}

	return nil
}
