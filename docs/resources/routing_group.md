---
layout: "metanetworks"
page_title: "Meta Networks: metanetworks_routing_group_attachment"
description: |-
  Provides a routing group resource.
---

# Resource: metanetworks_routing_group_attachment

Provides a routing group resource.

## Example Usage

```hcl
data "metanetworks_group" "example" {
  name = "example"
}

resource "metanetworks_routing_group" "organization" {
  name    = "organization"
  sources = data.metanetworks_group.example.id
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the routing group.
* `description` - (Optional) The description of the routing group.
* `exempt_sources` - (Optional) Set of users and/or groups/devices to exempt from the routing group.
* `sources` - (Optional) Set of users and/or groups/devices to attach to the routing group.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of the resource.
* `priority` - The priority of the routing group.
* `mapped_elements_ids` - Set of Mapped Subnets and Services to attach to the routing group.
* `created_at` - Creation timestamp.
* `modified_at` - Modification timestamp.
* `org_id` - The ID of the organization.
