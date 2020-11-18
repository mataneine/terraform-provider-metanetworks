---
layout: "metanetworks"
page_title: "Metanetworks: metanetworks_user_data_source"
sidebar_current: "docs-metanetwork-data-source"
description: |-
  Get information about a user.
---

# Data Source: metanetworks_user

Return a user of the organization

## Example Usage

```hcl
data "metanetworks_user" "example" {
  email = "an_example_user_email"
}
```

## Argument Reference

The following arguments are supported:

* `email` - (required) The user's email.

## Attributes Reference

* `description` - The User Description.
* `enabled` - User is enabled.
* `family_name` - Last Name.
* `given_name` - First Name.
* `phone` - The User's phone number.
* `provisioned_by` - Name of the Identity Provider. ????
* `created_at` - Creation timestamp.
* `inventory` - List of elements owned by the user.
* `mfa_enabled` - State of the API MFA setting.
* `modified_at` - Modification timestamp.
* `name` - Name of the User entity. ????
* `org_id` - ID of the Organization. ????
* `overlay_mfa_enabled` - State of the Overlay MFA setting.
* `phone_verified` - User verified via phone. ????
* `roles` - Roles assigned to the user or inherited from Groups. ????
* `tags` - List ???? of tags associated with the user.
