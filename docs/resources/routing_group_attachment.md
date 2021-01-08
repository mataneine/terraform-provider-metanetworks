---
layout: "metanetworks"
page_title: "Meta Networks: metanetworks_routing_group_attachment"
description: |-
  Attach mapped subnets and mapped services to a routing group.
---

# Resource: metanetworks_routing_group_attachment

Attach mapped subnets and mapped services to a routing group.

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

resource "metanetworks_mapped_service" "example" {
  name           = "example"
  mapped_service = "example.com"
}

resource "metanetworks_routing_group_attachment" "example" {
  routing_group_id   = metanetworks_routing_group.example.id
  network_element_id = metanetworks_mapped_service.example.id
}
```

## Argument Reference

The following arguments are supported:

* `routing_group_id` - (Required) The ID of the routing group.
* `network_element_id` - (Required) The ID of the network element to attach to the routing group.
