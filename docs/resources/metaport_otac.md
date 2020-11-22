---
layout: "metanetworks"
page_title: "Meta Networks: metanetworks_metaport_otac"
description: |-
  Request a OTAC (Time-limited One-Time Password) for metaport installer download.
---

# Resource: metanetworks_metaport_otac

Request a OTAC (Time-limited One-Time Password) for metaport installer download.

## Example Usage

```hcl
resource "metanetworks_metaport" "example" {
  name    = "example"
  enabled = false
}

resource "metanetworks_metaport_otac" "example" {
  metaport_id = metanetworks_metaport.example.id
}
```

## Argument Reference

The following arguments are supported:

* `metaport_id` - (Required) The ID of the metaport.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `expires_in`- expiration time, default=60minutes.
* `secret` - OTAC (Time-limited One-Time Password) for metaport installer download.
