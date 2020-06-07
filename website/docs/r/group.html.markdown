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
* `description` - 

## Attributes Reference

The following attributes are exported:

* `expression` - 
* `provisioned_by` - 
* `created_at` - 
* `modified_at` - 
* `org_id` - 
* `members` - 
* `roles` - 
* `users` - 