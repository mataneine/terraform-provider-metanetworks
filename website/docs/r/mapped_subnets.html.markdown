---
layout: "metanetworks"
page_title: "Metanetworks: metanetworks_mapped_subnets_resource"
sidebar_current: "docs-metanetworks-resource"
description: |-
  Create a new network element.
---

# metanetworks_mapped_subnet_resource

Create a new subnet network element.

## Example Usage

```hcl
resource "metanetworks_mapped_subnets" "example" {
  name           = "example"
  mapped_subnets = ["172.16.1.0/28"]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (required) Subnet network element to add to the organization.
* `mapped_subnets` - (required)

## Attributes Reference

The following attributes are exported:

* `name` - 
* `description` - 
* `tags` - 
* `mapped_domains` - 
* `created_at` - 
* `dns_name` - 
* `expires_at` - 
* `modified_at` - 
* `org_id` - 
* `net_id` - 
* `type` - 
* `version` - 