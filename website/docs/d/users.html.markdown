---
layout: "metanetworks"
page_title: "Metanetworks: users_data_source"
sidebar_current: "docs-metanetwork-data-source"
description: |-
  Get information about a user.
---

# Data Source: user

This data source can be used to fetch information about a specific user. By using this data source, you can reference user properties without having to hard code IDs as input.

## Example Usage

```hcl
data "metanetworks_user" "example" {
  email = "an_example_user_email"
}
```

## Attributes Reference

* `description` - 
* `email` - 
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