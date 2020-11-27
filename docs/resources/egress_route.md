---
layout: "metanetworks"
page_title: "Meta Networks: metanetworks_egress_route_resource"
description: |-
  Provides an egress route resource.
---

# Resource: metanetworks_egress_route_resource

Provides an egress route resource.

## Example Usage

```hcl
data "metanetworks_group" "example" {
  name = "example"
}

resource "metanetworks_mapped_subnets" "example" {
  name           = "example"
  mapped_subnets = ["172.16.0.0/22"]
}

resource "metanetworks_egress_route" "example" {
  name = "example"
  destinations = [
    "example.com",
  ]
  sources = [
    data.metanetworks_group.example.id
  ]
  via = metanetworks_mapped_subnets.example.id
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the egress route.
* `description` - (Optional) The description of the egress route.
* `destinations` - (Optional) Target hostnames.
* `enabled` - (Optional) default=true.
* `exempt_sources` - (Optional) Set of users and/or groups/devices/mapped subnets to exempt from the egress route.
* `sources` - (Optional) Set of users and/or groups/devices/mapped subnets to attach to the egress route.
* `via` - (Required) Region or mapped subnet.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of the resource.
* `created_at` - Creation timestamp.
* `modified_at` - Modification timestamp.
* `org_id` - The ID of the organization.
