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

* `name` - (required) Mapped Subnet Name.
* `mapped_subnets` - (required) CIDRs to map for this service.

## Attributes Reference

The following attributes are exported:

* `platform` ???? - ???? not available in provider 
* `description` - Mapped Subnet Description
* `tags` - List ???? of tags associated with the subnet.
* `mapped_domains` - List ???? of mapped domains
* `created_at` - Creation timestamp.
* `dns_name` - Hostname
* `expires_at` - Expiry timestamp
* `modified_at` - Modification timestamp.
* `org_id` - ID of the Organization. ????
* `net_id` - ID of the (subnet ????) network element. ????
* `type` - ["Device" | "Service" | "Mapped Service" | "Mapped Subnet"]
* `version` ???? (not in api) - 
