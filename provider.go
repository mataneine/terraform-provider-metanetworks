package main

import (
	"errors"
	"terraform-provider-metanetworks/metanetworks"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				DefaultFunc: schema.EnvDefaultFunc("METANETWORKS_API_KEY", nil),
				Optional:    true,
			},
			"api_secret": {
				Type:        schema.TypeString,
				DefaultFunc: schema.EnvDefaultFunc("METANETWORKS_API_SECRET", nil),
				Optional:    true,
				Sensitive:   true,
			},
			"org": {
				Type:        schema.TypeString,
				DefaultFunc: schema.EnvDefaultFunc("METANETWORKS_ORG", nil),
				Optional:    true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"metanetworks_metaport":                        resourceMetaport(),
			"metanetworks_metaportOTAC":                    resourceMetaportOTAC(),
			"metanetworks_metaport_network_element_attach": resourceMetaportElementAttach(),
			"metanetworks_mapped_service":                  resourceMappedService(),
			"metanetworks_mapped_service_alias":            resourceMappedServiceAlias(),
			"metanetworks_group":                           resourceGroup(),
			"metanetworks_policy":                          resourcePolicy(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	api_key, haveAPIKey := d.GetOk("api_key")
	api_secret, haveAPISecret := d.GetOk("api_secret")
	org, haveOrg := d.GetOk("org")

	var client *metanetworks.Client
	var err error

	// If one is set
	if haveAPIKey || haveAPISecret || haveOrg {
		// but not all are set
		if !(haveAPIKey && haveAPISecret && haveOrg) {
			return nil, errors.New("Please provide an api_key, api_secret and org. Alternatively provide a configuration file and none of these parameters")
		} else {
			client, err = metanetworks.NewClient(api_key.(string), api_secret.(string), org.(string))
		}
	} else {
		client, err = metanetworks.NewClientFromConfig()
	}

	if err != nil {
		return nil, err
	}

	return client, nil
}
