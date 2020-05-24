---
layout: "metanetworks"
page_title: "Metanetworks: metanetworks_routing_group_attachment_resource"
sidebar_current: "docs-metanetworks-resource"
description: |-
  Modify a routing group of the organization.
---

# metanetworks_routing_group_attachment

Modify a routing group of the organization.

## Example Usage

```hcl
resource "metanetworks_routing_group_attachment" "example" {
  routing_group_id   = metanetworks_routing_group.organization.id
  network_element_id = metanetworks_mapped_service.example.id
}
resource "metanetworks_routing_group" "organization" {
  name    = "organization"
  sources = data.metanetworks_group.example_group.id
}
output "organization" {
  value = metanetworks_routing_group.organization
}
data "metanetworks_group" "example_group" {
  name = "example-group"
}
output "example_group" {
  value = data.metanetworks_group.example_group
}
resource "metanetworks_mapped_service" "example" {
  name = "example"
  mapped_service = example.com"
}
output "example" {
  value = metanetworks_mapped_service.example
}
```

## Argument Reference

The following arguments are supported:

* `routing_group_id` - (required)
* `network_element_id` - (required) IDs of groups and/or Users to attach to the routing group
