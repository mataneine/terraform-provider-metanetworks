---
layout: "metanetworks"
page_title: "Meta Networks: metanetworks_device"
description: |-
  Provides a device resource.
---

# Resource: metanetworks_device

Provides a device resource.

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
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the device.
* `description` - (Optional) The description of the device.
* `enabled` - (Optional) default=true.
* `owner_id` - (Required) The ID of owner of the device.
* `platform` - (Required) The platform of the device. Valid values are `Android`, `macOS`, `iOS`, `Linux`, `Windows` and `ChromeOS`.
* `tags` - (Optional) Tags are key/value attributes that can be used to group elements together.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of the resource.
* `aliases` - The domain names of the device.
* `dns_name` - `<network_element_id>`.`<org_id>`.nsof
* `expires_at` - Expiration timesptamp.
* `created_at` - Creation timestamp.
* `modified_at` - Modification timestamp.
* `org_id` - The ID of the organization.
