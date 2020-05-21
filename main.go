package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"terraform-provider-metanetworks/metanetworks"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: metanetworks.Provider})
}
