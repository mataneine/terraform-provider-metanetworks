package metanetworks

import (
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePeeringAttachment() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"peering_id": {
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
		Create: resourcePeeringAttachmentCreate,
		Read:   resourcePeeringAttachmentRead,
		Delete: resourcePeeringAttachmentDelete,
	}
}

func resourcePeeringAttachmentCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	elementID := d.Get("network_element_id").(string)
	peeringID := d.Get("peering_id").(string)

	metanetworksMutexKV.Lock(d.Id())
	defer metanetworksMutexKV.Unlock(d.Id())

	var peering *Peering
	peering, err := client.GetPeering(peeringID)
	if err != nil {
		return err
	}

	for i := 0; i < len(peering.Peers); i++ {
		if peering.Peers[i] == elementID {
			return errors.New("That network element is already mapped to this Peering")
		}

	}

	peering.Peers = append(peering.Peers, elementID)
	_, err = client.UpdatePeering(peeringID, peering)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s_%s", peeringID, elementID))

	return resourcePeeringAttachmentRead(d, m)
}

func resourcePeeringAttachmentRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	elementID := d.Get("network_element_id").(string)
	peeringID := d.Get("peering_id").(string)

	metanetworksMutexKV.Lock(d.Id())
	defer metanetworksMutexKV.Unlock(d.Id())

	var peering *Peering
	peering, err := client.GetPeering(peeringID)
	if err != nil {
		return err
	}

	found := false
	for i := 0; i < len(peering.Peers); i++ {
		if peering.Peers[i] == elementID {
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

func resourcePeeringAttachmentDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	elementID := d.Get("network_element_id").(string)
	peeringID := d.Get("peering_id").(string)

	metanetworksMutexKV.Lock(d.Id())
	defer metanetworksMutexKV.Unlock(d.Id())

	var peering *Peering
	peering, err := client.GetPeering(peeringID)
	if err != nil {
		return err
	}

	// Note that if the entry has already been deleted this won't fail.
	for i := 0; i < len(peering.Peers); i++ {
		if peering.Peers[i] == elementID {
			peering.Peers = append(peering.Peers[:i], peering.Peers[i+1:]...)
			break
		}
	}

	_, err = client.UpdatePeering(peeringID, peering)
	return err
}
