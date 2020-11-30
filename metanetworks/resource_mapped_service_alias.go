package metanetworks

import (
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMappedServiceAlias() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"mapped_service_id": {
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
		Create: resourceMappedServiceAliasCreate,
		Read:   resourceMappedServiceAliasRead,
		Delete: resourceMappedServiceAliasDelete,
	}
}

func resourceMappedServiceAliasCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	mappedServiceID := d.Get("mapped_service_id").(string)
	alias := d.Get("alias").(string)

	var networkElement *NetworkElement
	networkElement, err := client.GetNetworkElement(mappedServiceID)
	if err != nil {
		return err
	}

	for i := 0; i < len(networkElement.Aliases); i++ {
		if networkElement.Aliases[i] == alias {
			return errors.New("That is alias is already present on the Mapped Service")
		}
	}

	_, err = client.SetNetworkElementAlias(mappedServiceID, alias)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s_%s", mappedServiceID, alias))

	return resourceMappedServiceAliasRead(d, m)
}

func resourceMappedServiceAliasRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceMappedServiceAliasDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	mappedServiceID := d.Get("mapped_service_id").(string)
	alias := d.Get("alias").(string)
	var networkElement *NetworkElement
	networkElement, err := client.GetNetworkElement(mappedServiceID)
	if err != nil {
		return err
	}

	for i := 0; i < len(networkElement.Aliases); i++ {
		if networkElement.Aliases[i] == alias {
			_, err = client.DeleteNetworkElementAlias(mappedServiceID, alias)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
