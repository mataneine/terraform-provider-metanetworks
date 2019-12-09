package main

import (
	"crypto/md5"
	"errors"
	"terraform-provider-metanetworks/metanetworks"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceMappedServiceAlias() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"mapped_service_id": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The ID of the Mapped Service",
				Required:    true,
				ForceNew:    true,
			},
			"alias": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The alias to add",
				Required:    true,
				ForceNew:    true,
			},
		},
		Create: resourceMappedServiceAliasCreate,
		Read:   resourceMappedServiceAliasRead,
		Delete: resourceMappedServiceAliasDelete,
	}
}

func resourceMappedServiceAliasCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*metanetworks.Client)

	mappedServiceID := d.Get("mapped_service_id").(string)
	alias := d.Get("alias").(string)
	var networkElement *metanetworks.NetworkElement
	networkElement, err := client.GetNetworkElement(mappedServiceID)
	if err != nil {
		return err
	}

	for i := 0; i < len(networkElement.Aliases); i++ {
		if networkElement.Aliases[i] == alias {
			return errors.New("That is alias is already present on the Mapped Service")
		}
	}

	_, err = client.AddNetworkElementAlias(mappedServiceID, alias)
	if err != nil {
		return err
	}

	d.SetId(makeID(mappedServiceID + alias))

	return nil
}

func makeID(x string) string {
	h := md5.Sum([]byte(x))
	return string(h[:])
}

func resourceMappedServiceAliasRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceMappedServiceAliasDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*metanetworks.Client)

	mappedServiceID := d.Get("mapped_service_id").(string)
	alias := d.Get("alias").(string)
	var networkElement *metanetworks.NetworkElement
	networkElement, err := client.GetNetworkElement(mappedServiceID)
	if err != nil {
		return err
	}

	for i := 0; i < len(networkElement.Aliases); i++ {
		if networkElement.Aliases[i] == alias {
			_, err = client.RemoveNetworkElementAlias(mappedServiceID, alias)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
