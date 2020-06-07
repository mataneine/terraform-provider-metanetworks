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
resource "metanetworks_routing_group" "organization" {
  name    = "organization"
  sources = data.metanetworks_group.example_group.id
}
output "organization" {
  value = metanetworks_routing_group.organization
```

## Argument Reference

The following arguments are supported:

* `name` - Required
* `description` - 
* `mapped_elements_ids` - (required)
* `sources` - 
exempt_sources???

## Attributes Reference

The following attributes are exported:

* `created_at` - 
* `modified_at` - 
* `org_id` - 
* `priority` -
id ???? - routing_group_id ????