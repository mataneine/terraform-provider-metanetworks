---
layout: "metanetworks"
page_title: "Metanetworks: metanetworks_group_data_source"
sidebar_current: "docs-metanetworks-data-source"
description: |-
  Get information about a group.
---

# Data Source: metanetworks_group

Return a group of the organization

## Example Usage

```hcl
data "metanetworks_group" "example" {
  name = "an_example_group name"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (required) The group name.

## Attributes Reference

* `description` - The Group Description.
* `expression` - Smart Groups Expression.
* `provisioned_by` - Name of the Identity Provider. ????
* `created_at` - Creation Timestamp.
* `modified_at` - Modification Timestamp.
* `org_id` - The Organization ID. ????
* `members` - The Group members.
* `roles` - The Group Roles.
* `users` - Users belonging to the group.