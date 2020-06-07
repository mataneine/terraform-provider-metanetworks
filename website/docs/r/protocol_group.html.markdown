---
layout: "metanetworks"
page_title: "Metanetworks: metanetworks_protocol_group_resource"
sidebar_current: "docs-metanetworks-resource"
description: |-
  Modify a routing group of the organization.
---

# metanetworks_protocol_group

Modify a protocol group of the organization.

## Example Usage

```hcl
resource "metanetworks_protocol_group" "https" {
  name = "HTTPS"
  protocols {
    port  = 443
    proto = "tcp"
  }
}
```

## Argument Reference

The following arguments are supported:

* `description` - 
* `name` - 
* `protocols` -
			* `port` - 
			* `proto` - 

## Attributes Reference

The following attributes are exported:

* `created_at` - 
* `modified_at` - 
* `org_id` - protocol_group_id
* `read_only` - 
