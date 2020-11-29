---
layout: "metanetworks"
page_title: "Meta Networks: metanetworks_device_alias"
description: |-
  Set an alias to a device.
---

# Resource: metanetworks_device_alias

  Set an alias to a device.

## Example Usage

```hcl
data "metanetworks_user" "example" {
  email = "user@example.com"
}

resource "metanetworks_device" "example" {
  name     = "example"
  owner_id = data.metanetworks_user.example.id,
  platform = "macOS"
}

resource "metanetworks_device_alias" "example" {
  device_id = metanetworks_device.example.id
  alias     = "example.com"
}
```

## Argument Reference

The following arguments are supported:

* `device_id` - (Required) The ID of the device.
* `alias` - (Required) Domain name.
