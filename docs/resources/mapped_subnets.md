---
layout: "metanetworks"
page_title: "Meta Networks: metanetworks_mapped_subnets"
description: |-
  Provides a mapped subnet resource.
---

# Resource: metanetworks_mapped_subnets

Provides a mapped subnet resource.

## Example Usage

```hcl
resource "metanetworks_mapped_subnets" "example" {
  name           = "example"
  mapped_subnets = ["172.16.1.0/28"]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Mapped Subnet Name.
* `description` - (Optional) Mapped Subnet Description
* `mapped_subnets` - (Required) CIDRs to map for this service.
* `tags` - (Optional) Tags are key/value attributes that can be used to group elements together.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of the resource.
* `dns_name` - `<network_element_id>`.`<org_id>`.nsof
* `expires_at` - Expiration timestamp
* `created_at` - Creation timestamp.
* `modified_at` - Modification timestamp.
* `org_id` - The ID of the organization.
* `mapped_domains` - List of mapped domains. Fields documented below.

### Mapped Domains Attributes

For **mapped_domains** the following attributes are exported.

* `name` - Mapped DNS suffix.
* `mapped_domain` - Remote DNS suffix.
* `enterprise_dns` - Resolve and route traffic according to routing group, default=false.

