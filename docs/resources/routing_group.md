---
layout: "metanetworks"
page_title: "Meta Networks: metanetworks_routing_group"
description: |-
  Provides a routing group resource.
---

# Resource: metanetworks_routing_group

Provides a routing group resource.

## Example Usage

```hcl
data "metanetworks_group" "example" {
  name = "example"
}

resource "metanetworks_routing_group" "example" {
  name    = "example"
  sources = [
    data.metanetworks_group.example.id
  ]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the routing group.
* `description` - (Optional) The description of the routing group.
* `exempt_sources` - (Optional) Set of users and/or groups/devices to exempt from the routing group.
* `sources` - (Optional) Set of users and/or groups/devices to attach to the routing group.
* `priority` - (Optional) The priority of the routing group. Valid values are `0..256`.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of the resource.
* `mapped_elements_ids` - Mapped Subnets/Services.
* `created_at` - Creation timestamp.
* `modified_at` - Modification timestamp.
* `org_id` - The ID of the organization.
