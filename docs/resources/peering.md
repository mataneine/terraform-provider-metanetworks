---
layout: "metanetworks"
page_title: "Meta Networks: metanetworks_peering"
description: |-
  Provides a peering resource.
---

# Resource: metanetworks_peering

Provides a peering resource.

## Example Usage

```hcl
resource "metanetworks_mapped_service" "example" {
  name           = "example"
  mapped_service = "example.com"
}

resource "metanetworks_mapped_subnets" "example" {
  name           = "example"
  mapped_subnets = ["172.16.1.0/28"]
}

resource "metanetworks_peering" "example" {
  name     = " example"
  peers    = [
    metanetworks_mapped_service.example.id,
    metanetworks_mapped_subnets.example.id
  ]
}
```

## Argument Reference

The following arguments are supported:

* `description` - (Optional) The description of the peering.
* `egress_nat` - (Optional) default=true.
* `enabled` - (Optional) default=true.
* `name` - (Required) The name of the peering.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of the resource.
* `peers` - Mapped Subnets/Services.
* `created_at` - Creation timestamp.
* `modified_at` - Modification timestamp.
* `org_id` - The ID of the organization.
