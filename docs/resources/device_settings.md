---
layout: "metanetworks"
page_title: "Meta Networks: metanetworks_device_settings"
description: |-
  Provides a device settings resource.
---

# Resource: metanetworks_device_settings

Provides a device settings resource.

## Example Usage

```hcl
resource "metanetworks_device_settings" "example" {
  name                  = "Default_Org_Device_Setting"
  description           = "Example Description"
  apply_on_org          = true
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required). The name of the device settings.
* `description` - (Optional). The description of the device settings.
* `direct_sso`	- (Optional). Auto-select SSO at login
* `dns_server_type` - (Optional). Values: `OVERLAY`,`UNDERLAY`.
* `enabled` - (Optional) default=true.
* `apply_on_org` - (Optional; Required if `sources` is omitted). Applies settings to entire organization.
* `vpn_login_browser` - (Optional). Type of login for VPN. Values: `AGENT`, `EXTERNAL`, `USER_DEFINED`.
* `protocol_selection_lifetime` - (Optional). Specifies time in *minutes* that the protocol selection is valix. Max `525600` (1 year).
* `search_domains` - (Optional). (Optional). Specify domains? 
* `sources` - (Optional; Required if `apply_on_org` is omitted). Applies setting to specified sources.
* `session_lifetime` - (Optional). Time in *minutes* before requiring reauthentication.
* `session_lifetime_grace` - (Optional). Time in *minutes* prior to session expiration to request reauthentication. Max: `60` (1 hour).
* `split_tunnel` - (Optional). Setting to `true` will route all traffic (including internet bound) through Meta. Requires a Default Route mapped_subnet, metaport and egress resources.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of the resource.
* `created_at` - Creation timestamp.
* `modified_at` - Modification timestamp.
