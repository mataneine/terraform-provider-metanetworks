---
layout: "metanetworks"
page_title: "Metanetworks: metanetworks_mapped_service_resource"
sidebar_current: "docs-metanetworks-resource"
description: |-
  Create a new mapped_service.
---

# metanetworks_mapped_service_resource

Create a new mapped_service.

## Example Usage

```hcl
resource "metanetworks_mapped_service" "example" {
  name = "an_example_mapped_service name"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (required) The mapped_service name.
* `mapped_service` - hostname-or-ipv4
* `description` - 

## Attributes Reference

The following attributes are exported:

* `tags` - 
* `created_at` - 
* `dns_name` - 
* `expires_at` - 
* `modified_at` - 
* `org_id` - 
* `aliases` - 
* `net_id` - 
* `type` - 
* `version` - 