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

* `name` - (required) The name of the Routing Group.
* `description` - The description of the Routing Group.
* `mapped_elements_ids` - (required / attribute ????) List of Mapped Subnets and Services to attach to the Routing Group.
* `sources` - (List of ????) sources to attach to the Routing Group (Devices, Groups, Native Services, users) (concat users & groups ????)
no exempt_sources ????

## Attributes Reference

The following attributes are exported:

* `created_at` - Creation timestamp.
* `modified_at` - Modification timestamp.
* `org_id` - ID of the Orgabization
* `priority` - ????
id ???? - routing_group_id ????