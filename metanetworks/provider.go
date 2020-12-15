package metanetworks

import (
	"errors"

	"github.com/hashicorp/terraform-plugin-sdk/helper/mutexkv"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider returns a terraform.ResourceProvider.
func Provider() *schema.Provider {

	// The actual provider
	provider := &schema.Provider{
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
		DataSourcesMap: map[string]*schema.Resource{
			"metanetworks_group":     dataSourceGroup(),
			"metanetworks_locations": dataSourceLocations(),
			"metanetworks_user":      dataSourceUser(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"metanetworks_egress_route":                 resourceEgressRoute(),
			"metanetworks_group":                        resourceGroup(),
			"metanetworks_device_alias":                 resourceDeviceAlias(),
			"metanetworks_device":                       resourceDevice(),
			"metanetworks_mapped_service_alias":         resourceMappedServiceAlias(),
			"metanetworks_mapped_service":               resourceMappedService(),
			"metanetworks_mapped_subnets_mapped_domain": resourceMappedSubnetsMappedDomain(),
			"metanetworks_mapped_subnets":               resourceMappedSubnets(),
			"metanetworks_metaport_attachment":          resourceMetaportAttachment(),
			"metanetworks_metaport_otac":                resourceMetaportOTAC(),
			"metanetworks_metaport":                     resourceMetaport(),
			"metanetworks_native_service_alias":         resourceNativeServiceAlias(),
			"metanetworks_native_service":               resourceNativeService(),
			"metanetworks_peering_attachment":           resourcePeeringAttachment(),
			"metanetworks_peering":                      resourcePeering(),
			"metanetworks_policy":                       resourcePolicy(),
			"metanetworks_protocol_group":               resourceProtocolGroup(),
			"metanetworks_routing_group_attachment":     resourceRoutingGroupAttachment(),
			"metanetworks_routing_group":                resourceRoutingGroup(),
			"metanetworks_user_settings":                resourceUserSettings(),
			"metanetworks_posture_check":                resourcePostureCheck(),
		},
	}

	provider.ConfigureFunc = func(d *schema.ResourceData) (interface{}, error) {
		terraformVersion := provider.TerraformVersion
		if terraformVersion == "" {
			// Terraform 0.12 introduced this field to the protocol
			// We can therefore assume that if it's missing it's 0.10 or 0.11
			terraformVersion = "0.11+compatible"
		}
		return providerConfigure(d, terraformVersion)
	}

	return provider
}

func providerConfigure(d *schema.ResourceData, terraformVersion string) (interface{}, error) {
	apiKey, haveAPIKey := d.GetOk("api_key")
	apiSecret, haveAPISecret := d.GetOk("api_secret")
	org, haveOrg := d.GetOk("org")

	var client *Client
	var err error

	// If one is set
	if haveAPIKey || haveAPISecret || haveOrg {
		// but not all are set
		if !(haveAPIKey && haveAPISecret && haveOrg) {
			return nil, errors.New("Please provide an api_key, api_secret and org. Alternatively provide a configuration file and none of these parameters")
		}
		{
			client, err = NewClient(apiKey.(string), apiSecret.(string), org.(string))
		}
	} else {
		client, err = NewClientFromConfig()
	}

	if err != nil {
		return nil, err
	}

	client.terraformVersion = terraformVersion

	return client, nil
}

// This is a global MutexKV for use within this plugin.
var metanetworksMutexKV = mutexkv.NewMutexKV()
