---
layout: "metanetworks"
page_title: "Metanetworks: metanetworks_metaport_otac_resource"
sidebar_current: "docs-metanetworks-resource"
description: |-
  Request a OTAC (Time-limited One-Time Password) for metaport installer download
---

# metanetworks_metaport_otac_resource

Request a OTAC (Time-limited One-Time Password) for metaport installer download.

## Example Usage

```hcl
resource "metanetworks_metaport_otac" "example" {
  metaport_id        = metanetworks_metaport.example-metaport.id
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

* `metaport_id` - 
* `triggers` - ????

## Attributes Reference

The following attributes are exported:

* `secret` - 
