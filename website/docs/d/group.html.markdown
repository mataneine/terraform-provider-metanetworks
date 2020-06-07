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

* `description` - 
* `expression` - Allows grouping entities by their tags. Filtering by tag value is also supported if provided.
* `provisioned_by` - 
* `created_at` - 
* `modified_at` - 
* `org_id` - 
* `members` - Returned when Expand=True is provided to call
* `roles` - 
* `users` - 