---
layout: "metanetworks"
page_title: "Metanetworks: metanetworks_metaport_attachment_resource"
sidebar_current: "docs-metanetworks-resource"
description: |-
  attach a network element to a metaport.
---

# metanetworks_metaport_attachment_resource

attach a network element to a metaport.

## Example Usage

```hcl
resource "metanetworks_metaport_attachment" "example" {
  metaport_id        = metanetworks_metaport.example-metaport.id
  network_element_id = metanetworks_mapped_service.example.id
}
resource "metanetworks_mapped_service" "example" {
  name = "example"
  mapped_service = example.com"
}
output "example" {
  value = metanetworks_mapped_service.example
}
resource "metanetworks_metaport" "example-metaport" {
  name    = "example-metaport"
  enabled = false
}
output "example-metaport" {
  value = metanetworks_metaport.example-metaport
}
```

## Argument Reference

The following arguments are supported:

* `metaport_id` - (required) The ID of the metaport.
* `network_element_id` - (required) The ID of the network element to attach to the Metaport.
