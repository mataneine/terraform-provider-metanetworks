---
layout: "metanetworks"
page_title: "Metanetworks: metanetworks_group_resource"
sidebar_current: "docs-metanetworks-resource"
description: |-
  Create a new group.
---

# metanetworks_group_resource

Create a new group.

## Example Usage

```hcl
resource "metanetworks_group" "example" {
  name = "an_example_group name"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (required) The group name.
* `description` - The Group Description.

## Attributes Reference

The following attributes are exported:

* `expression` - Smart Groups Expression.
* `provisioned_by` - Name of the Identity Provider. ????
* `created_at` - Creation timestamp.
* `modified_at` - Modification timestamp.
* `org_id` - The Organization ID. ????
* `members` - The Group members.
* `roles` - Group roles
* `users` - Users belonging to the group.