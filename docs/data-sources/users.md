---
layout: "metanetworks"
page_title: "Meta Networks: metanetworks_user"
description: |-
  Return a user of the organization.
---

# Data Source: metanetworks_user

Return a user of the organization.

## Example Usage

```hcl
data "metanetworks_user" "example" {
  email = "user@example.com"
}
```

## Argument Reference

The following arguments are supported:

* `email` - (Required) The email address of the user.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of the resource.
* `description` - The description of the user.
* `enabled` - default=true.
* `family_name` - The last name of the user.
* `given_name` - The first name of the user.
* `name` - The full name of the user.
* `phone` - The phone number of the user.
* `phone_verified` - Meta account login 2-Step verification phone number verified.
* `provisioned_by` - Users can be provisioned in the system either by locally creating the users from the Admin portal or API. Another, more common practice, is to provision users from an organization directory service, by way of SCIM or LDAP protocols.
* `inventory` - Devices owned by the user.
* `mfa_enabled` - State of the API MFA setting.
* `overlay_mfa_enabled` - State of the Overlay MFA setting.
* `roles` - Roles assigned to the user or inherited from Groups.
* `tags` - Tags are key/value attributes that can be used to group elements together to Smart Groups, and placed as target or sources in Policies.
* `created_at` - Creation timestamp.
* `modified_at` - Modification timestamp.
* `org_id` - The ID of the organization.
