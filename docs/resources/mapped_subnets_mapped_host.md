---
layout: "metanetworks"
page_title: "Meta Networks: metanetworks_mapped_subnets_mapped_host"
description: |-
  Set a mapped host to a mapped subnet.
---

# Resource: metanetworks_mapped_subnets_mapped_host

Set a mapped host to a mapped subnet.

You can optionally assign additional domain names for specific hosts in the mapped subnet.


## Example Usage

```hcl
resource "metanetworks_mapped_subnets" "example" {
  name           = "example"
  mapped_subnets = ["172.16.0.0/22"]
}

resource "metanetworks_mapped_subnets_mapped_host" "example" {
  mapped_subnets_id = metanetworks_mapped_subnets.example.id
  name              = "ec2.internal"
  mapped_host       = "ec2.internal"
  ignore_bounds     = true
}
```

## Argument Reference

The following arguments are supported:

* `mapped_subnets_id` - (Required) ID of the mapped subnet network element.
* `name` - (Required) Mapped hostname.
* `mapped_host` - (Required) Remote hostname or IP.
* `ignore_bounds` - (Optional) Allow setting mapped hosts outside of the defined mapped subnets, default=false.
