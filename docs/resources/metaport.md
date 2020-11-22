---
layout: "metanetworks"
page_title: "Meta Networks: metanetworks_metaport"
description: |-
  Provides a metaport resource.
---

# Resource: metanetworks_metaport

Provides a metaport resource.

## Example Usage

```hcl
resource "metanetworks_metaport" "example" {
  name    = "example"
  enabled = false
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the metaport.
* `description` - (Optional) The description of the metaport.
* `enabled` - (Optional) default=true.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of the resource.
* `allow_support` - Enable external support to access to this metaport remotely, default=true.
* `mapped_elements` - Network elements attached to the metaport.
* `dns_name` - <metaport_id>.<org_id>.nsof
* `expires_at` - Expiry timestamp.
* `created_at` - Creation timestamp.
* `modified_at` - Modification timestamp.
* `org_id` - The ID of the organization.
