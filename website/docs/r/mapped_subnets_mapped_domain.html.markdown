---
layout: "metanetworks"
page_title: "Metanetworks: metanetworks_mapped_subnets_mapped_domain_resource"
sidebar_current: "docs-metanetworks-resource"
description: |-
  Set a mapped domain to a subnets ???? network element.
---

# metanetworks_mapped_subnets_mapped_domain_resource

Set a mapped domain to a subnets ???? network element.

## Example Usage

```hcl
resource "metanetworks_mapped_subnets_mapped_domain" "example" {
  mapped_subnets_id = metanetworks_mapped_subnets.example.id
  name              = "ec2.internal"
  mapped_domain     = "ec2.internal"
  enterprise_dns    = true
}
resource "metanetworks_mapped_subnets" "example" {
  name           = "example"
  mapped_subnets = ["172.16.0.0/22"]
}
output "example" {
  value = metanetworks_mapped_subnets.example
}
```

## Argument Reference

The following arguments are supported:

* `mapped_domain` - (required) Mapped domain to set in the network element
* `enterprise_dns` - Resolve and route IPv6 according to routing group (boolean)
* `name` - [(required) ???? (ForceNew)] Mapped domain name
* `mapped_subnets_id` - [(required) ???? (ForceNew)] ID of the subnets ???? network element
