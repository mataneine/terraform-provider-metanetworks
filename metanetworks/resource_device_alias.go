package metanetworks

import (
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDeviceAlias() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"device_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"alias": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
		Create: resourceDeviceAliasCreate,
		Read:   resourceDeviceAliasRead,
		Delete: resourceDeviceAliasDelete,
	}
}

func resourceDeviceAliasCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	deviceID := d.Get("device_id").(string)
	alias := d.Get("alias").(string)

	var networkElement *NetworkElement
	networkElement, err := client.GetNetworkElement(deviceID)
	if err != nil {
		return err
	}

	for i := 0; i < len(networkElement.Aliases); i++ {
		if networkElement.Aliases[i] == alias {
			return errors.New("That is alias is already present on the Native Service")
		}
	}

	_, err = client.SetNetworkElementAlias(deviceID, alias)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s_%s", deviceID, alias))

	return resourceDeviceAliasRead(d, m)
}

func resourceDeviceAliasRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceDeviceAliasDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	deviceID := d.Get("device_id").(string)
	alias := d.Get("alias").(string)
	var networkElement *NetworkElement
	networkElement, err := client.GetNetworkElement(deviceID)
	if err != nil {
		return err
	}

	for i := 0; i < len(networkElement.Aliases); i++ {
		if networkElement.Aliases[i] == alias {
			_, err = client.DeleteNetworkElementAlias(deviceID, alias)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
