---
layout: "metanetworks"
page_title: "Meta Networks: metanetworks_mapped_subnets_mapped_domain"
description: |-
  Set a mapped domain to a mapped subnet.
---

# Resource: metanetworks_mapped_subnets_mapped_domain

Set a mapped domain to a mapped subnet.

## Example Usage

```hcl
resource "metanetworks_mapped_subnets" "example" {
  name           = "example"
  mapped_subnets = ["172.16.0.0/22"]
}

resource "metanetworks_mapped_subnets_mapped_domain" "example" {
  mapped_subnets_id = metanetworks_mapped_subnets.example.id
  name              = "ec2.internal"
  mapped_domain     = "ec2.internal"
  enterprise_dns    = true
}
```

## Argument Reference

The following arguments are supported:

* `mapped_domain` - (Required) Mapped domain to set in the network element
* `enterprise_dns` - (Optional) Resolve and route IPv6 according to routing group
* `name` - (Required) Mapped domain name.
* `mapped_subnets_id` - (Required) ID of the subnets network element
