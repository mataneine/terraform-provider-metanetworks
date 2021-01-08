---
layout: "metanetworks"
page_title: "Meta Networks: metanetworks_native_service"
description: |-
  Provides a native service resource.
---

# Resource: metanetworks_native_service

Provides a native service resource.

## Example Usage

```hcl
resource "metanetworks_native_service" "example" {
  name = "example"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the native service.
* `description` - (Optional) The description of the native service.
* `enabled` - (Optional) default=true.
* `tags` - (Optional) Tags are key/value attributes that can be used to group elements together.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of the resource.
* `aliases` - The domain names of the native service.
* `dns_name` - `<network_element_id>`.`<org_id>`.nsof
* `expires_at` - Expiration timesptamp.
* `created_at` - Creation timestamp.
* `modified_at` - Modification timestamp.
* `org_id` - The ID of the organization.
