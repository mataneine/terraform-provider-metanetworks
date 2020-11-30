package metanetworks

import (
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNativeServiceAlias() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"native_service_id": {
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
		Create: resourceNativeServiceAliasCreate,
		Read:   resourceNativeServiceAliasRead,
		Delete: resourceNativeServiceAliasDelete,
	}
}

func resourceNativeServiceAliasCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	nativeServiceID := d.Get("native_service_id").(string)
	alias := d.Get("alias").(string)

	var networkElement *NetworkElement
	networkElement, err := client.GetNetworkElement(nativeServiceID)
	if err != nil {
		return err
	}

	for i := 0; i < len(networkElement.Aliases); i++ {
		if networkElement.Aliases[i] == alias {
			return errors.New("That is alias is already present on the Native Service")
		}
	}

	_, err = client.SetNetworkElementAlias(nativeServiceID, alias)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s_%s", nativeServiceID, alias))

	return resourceNativeServiceAliasRead(d, m)
}

func resourceNativeServiceAliasRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceNativeServiceAliasDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	nativeServiceID := d.Get("native_service_id").(string)
	alias := d.Get("alias").(string)
	var networkElement *NetworkElement
	networkElement, err := client.GetNetworkElement(nativeServiceID)
	if err != nil {
		return err
	}

	for i := 0; i < len(networkElement.Aliases); i++ {
		if networkElement.Aliases[i] == alias {
			_, err = client.DeleteNetworkElementAlias(nativeServiceID, alias)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
