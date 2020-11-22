---
layout: "metanetworks"
page_title: "Meta Networks: metanetworks_metaport_attachment"
description: |-
  Attach mapped subnets and mapped services to a metaport.
---

# Resource: metanetworks_metaport_attachment

Attach mapped subnets and mapped services to a metaport.

## Example Usage

```hcl
resource "metanetworks_metaport" "example" {
  name    = "example"
  enabled = false
}

resource "metanetworks_mapped_service" "example" {
  name           = "example"
  mapped_service = "example.com"
}

resource "metanetworks_metaport_attachment" "example" {
  metaport_id        = metanetworks_metaport.example.id
  network_element_id = metanetworks_mapped_service.example.id
}
```

## Argument Reference

The following arguments are supported:

* `metaport_id` - (Required) The ID of the metaport.
* `network_element_id` - (Required) The ID of the network element to attach to the Metaport.
