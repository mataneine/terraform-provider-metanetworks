---
layout: "metanetworks"
page_title: "Meta Networks: metanetworks_user_settings"
description: |-
  Provides a user authentication settings resource.
---

# Resource: metanetworks_user_settings

Provides a user authentication settings resource.

## Example Usage

```hcl
resource "metanetworks_user_settings" "example" {
  name                  = "Default_Org_User_Settings"
  description           = "Example Description"
  apply_on_org          = true
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required). The name of the user settings.
* `description` - (Optional). The description of the user settings.
* `allowed_factors` - (Optional). List of allowed MFA factors when using Meta as idP. Values: `SMS`, `SOFTWARE_TOTP`, `VOICECALL`, `EMAIL`.
* `enabled` - (Optional) default=true.
* `apply_on_org` - (Optional; Required if `sources` is omitted). Applies setting to entire organization.
* `mfa_required` - (Optional). Force Multi-Factor Authentication. Only applies when using Meta as idP.
* `overlay_mfa_required` - (Optional). Specifies if MFA is required for overlay access.
* `sso_mandatory` - (Optional). All applicable users *MUST* use SSO. Useful if using Third Party idP (but not required).
* `sources` - (Optional; Required if `apply_on_org` is omitted). Applies setting to specified sources.
* `prohibited_os` - (Optional). List of operating systems to block from applying settings. Values: `Android`,`iOS`,`Windows`,`macOS`,`Linux`,`ChromeOS`.
* `max_devices_per_user` - (Optional). Maximum number of devices which any one user may enroll. Default `nul` (unlimited).
* `overlay_mfa_refresh_period` - (Optional). Time in *min* that an overlay MFA token is active for.
* `password_expiration` - (Optional). Password expiration in *days*

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of the resource.
* `created_at` - Creation timestamp.
* `modified_at` - Modification timestamp.
