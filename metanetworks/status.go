package metanetworks

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func StatusMetaportAttachmentCreate(client *Client, metaportID string, elementID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		var metaport *MetaPort
		metaport, err := client.GetMetaPort(metaportID)
		if err != nil {
			return 0, "", err
		}

		for i := 0; i < len(metaport.MappedElements); i++ {
			if metaport.MappedElements[i] == elementID {
				return metaport, "Completed", nil
			}
		}
		return metaport, "Pending", nil
	}
}

func StatusRoutingGroupAttachmentCreate(client *Client, routingGroupID string, elementID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		var routingGroup *RoutingGroup
		routingGroup, err := client.GetRoutingGroup(routingGroupID)
		if err != nil {
			return 0, "", err
		}

		for i := 0; i < len(routingGroup.MappedElements); i++ {
			if routingGroup.MappedElements[i] == elementID {
				return routingGroup, "Completed", nil
			}
		}
		return routingGroup, "Pending", nil
	}
}
