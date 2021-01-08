---
layout: "metanetworks"
page_title: "Meta Networks: metanetworks_mapped_subnets_mapped_domain"
description: |-
  Set a mapped domain to a mapped subnet.
---

# Resource: metanetworks_mapped_subnets_mapped_domain

Set a mapped domain to a mapped subnet.

Hosts in the remote subnet will be addressable by dashed IP: `<xxx-xxx-xxx-xxx>.<DNS suffix>`, or by hostname: `<hostname>.<DNS Suffix>`

* Mapped DNS suffixes with enabled Enterprise DNS are marked with an asterisk
* Maps DNS Suffix to another suffix in the remote subnet. This allows replacing addressing scheme to a different DNS domain. I.e. when DNS Suffix = **'my.subnet'** and Remote DNS Suffix = **'acme.local'**, it means that **'example.my.subnet'** is resolved in the remote subnet as **'example.acme.local'**

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

* `mapped_subnets_id` - (Required) ID of the mapped subnet network element.
* `name` - (Required) Mapped DNS suffix.
* `mapped_domain` - (Required) Remote DNS suffix.
* `enterprise_dns` - (Optional) Resolve and route traffic according to routing group, default=false.
