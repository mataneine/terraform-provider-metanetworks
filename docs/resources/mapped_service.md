---
layout: "metanetworks"
page_title: "Meta Networks: metanetworks_mapped_service"
description: |-
  Provides a mapped service resource.
---

# Resource: metanetworks_mapped_service

Provides a mapped service resource.

## Example Usage

```hcl
resource "metanetworks_mapped_service" "example" {
  name           = "example"
  mapped_service = "example.com"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the mapped service.
* `mapped_service` - (Required) Mapped Service IP or Hostname
* `description` - (Optional) The description of the mapped service.
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
