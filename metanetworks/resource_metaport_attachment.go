package metanetworks

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMetaportAttachment() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"metaport_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"network_element_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
		Create: resourceMetaportAttachmentCreate,
		Read:   resourceMetaportAttachmentRead,
		Delete: resourceMetaportAttachmentDelete,
	}
}

func resourceMetaportAttachmentCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	elementID := d.Get("network_element_id").(string)
	metaportID := d.Get("metaport_id").(string)

	metanetworksMutexKV.Lock(d.Id())
	defer metanetworksMutexKV.Unlock(d.Id())

	var metaport *MetaPort
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

	_, err = WaitMetaportAttachmentCreate(client, metaportID, elementID)

	if err != nil {
		return fmt.Errorf("Error waiting for metaport attachment creation (%s) (%s)", metaportID, err)
	}

	d.SetId(fmt.Sprintf("%s_%s", metaportID, elementID))

	return resourceMetaportAttachmentRead(d, m)
}

func resourceMetaportAttachmentRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	elementID := d.Get("network_element_id").(string)
	metaportID := d.Get("metaport_id").(string)

	metanetworksMutexKV.Lock(d.Id())
	defer metanetworksMutexKV.Unlock(d.Id())

	var metaport *MetaPort
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

	// If not present we need to destroy the terraform resource so that it is recreated.
	if !found {
		d.SetId("")
	}

	return nil
}

func resourceMetaportAttachmentDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	elementID := d.Get("network_element_id").(string)
	metaportID := d.Get("metaport_id").(string)

	metanetworksMutexKV.Lock(d.Id())
	defer metanetworksMutexKV.Unlock(d.Id())

	var metaport *MetaPort
	metaport, err := client.GetMetaPort(metaportID)
	if err != nil {
		return err
	}

	// Note that if the entry has already been deleted this won't fail.
	for i := 0; i < len(metaport.MappedElements); i++ {
		if metaport.MappedElements[i] == elementID {
			metaport.MappedElements = append(metaport.MappedElements[:i], metaport.MappedElements[i+1:]...)
			break
		}
	}

	err = resource.Retry(5*time.Second, func() *resource.RetryError {
		if _, err := client.UpdateMetaPort(metaportID, metaport); err != nil {
			if !strings.Contains(err.Error(), "is busy. Try again later.") {
				return resource.NonRetryableError(err)
			}
			return resource.RetryableError(err)
		}
		return nil
	})

	if err != nil {
		return fmt.Errorf("Error in metaport attachment deletion (%s) (%s)", metaportID, err)
	}

	return nil
}
