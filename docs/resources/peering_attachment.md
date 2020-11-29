---
layout: "metanetworks"
page_title: "Meta Networks: metanetworks_peering_attachment"
description: |-
  Attach mapped subnets and mapped services to a peering.
---

# Resource: metanetworks_peering_attachment

Attach mapped subnets and mapped services to a peering.

## Example Usage

```hcl
resource "metanetworks_peering" "example" {
  name    = "example"
}

resource "metanetworks_mapped_service" "example" {
  name           = "example"
  mapped_service = "example.com"
}

resource "metanetworks_peering_attachment" "example" {
  peering_id         = metanetworks_peering.example.id
  network_element_id = metanetworks_mapped_service.example.id
}
```

## Argument Reference

The following arguments are supported:

* `peering_id` - (Required) The ID of the peering.
* `network_element_id` - (Required) The ID of the network element to attach to the peering.
