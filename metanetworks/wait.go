package metanetworks

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func WaitMetaportAttachmentCreate(client *Client, metaportID string, elementID string) (*Client, error) {
	createStateConf := &resource.StateChangeConf{
		Pending:    []string{"Pending"},
		Target:     []string{"Completed"},
		Timeout:    5 * time.Minute,
		MinTimeout: 5 * time.Second,
		Delay:      3 * time.Second,
		Refresh:    StatusMetaportAttachmentCreate(client, metaportID, elementID),
	}

	_, err := createStateConf.WaitForState()
	if err != nil {
		return nil, err
	}

	return client, err
}

func WaitRoutingGroupAttachmentCreate(client *Client, routingGroupID string, elementID string) (*Client, error) {
	createStateConf := &resource.StateChangeConf{
		Pending:    []string{"Pending"},
		Target:     []string{"Completed"},
		Timeout:    5 * time.Minute,
		MinTimeout: 5 * time.Second,
		Delay:      3 * time.Second,
		Refresh:    StatusRoutingGroupAttachmentCreate(client, routingGroupID, elementID),
	}

	_, err := createStateConf.WaitForState()
	if err != nil {
		return nil, err
	}

	return client, err
}
