---
layout: "metanetworks"
page_title: "Metanetworks: metanetworks_metaport_attachment_resource"
sidebar_current: "docs-metanetworks-resource"
description: |-
  metaports to add to the organization.
---

# metanetworks_metaport_attachment_resource

metaports to add to the organization.

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

* `metaport_id` - (required) the ID of the metaport.
* `network_element_id` - (required)
