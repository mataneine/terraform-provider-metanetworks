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

* `description` - 
* `enabled` - 
* `family_name` - 
* `given_name` - 
* `phone` - 
* `provisioned_by` - 
* `created_at` - 
* `inventory` - List of elements owned by the user
* `mfa_enabled` - State of the API MFA setting
* `modified_at` - 
* `name` - 
* `org_id` - 
* `overlay_mfa_enabled` - State of the Overlay MFA setting
* `phone_verified` - 
* `roles` - 
* `tags` - 