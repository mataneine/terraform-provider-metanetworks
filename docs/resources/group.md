---
layout: "metanetworks"
page_title: "Meta Networks: metanetworks_group"
description: |-
  Provides a group resource.
---

# Resource: metanetworks_group

Provides a group resource.

## Example Usage

```hcl
resource "metanetworks_group" "example" {
  name = "example"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the group.
* `description` - (Optional) The description of the group.
* `expression` - (Optional) Allows grouping entities by their tags. Filtering by tag value is also supported if provided. Supported operations: AND, OR, XOR, parenthesis.
* `roles` - (Optional) The group roles.
* `users` - (Optional) The group users.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of the resource.
* `provisioned_by` - Groups can be provisioned in the system either by locally creating the groups from the Admin portal or API. Another, more common practice, is to provision groups from an organization directory service, by way of SCIM or LDAP protocols.
* `created_at` - Creation timestamp.
* `modified_at` - Modification timestamp.
* `org_id` - The ID of the organization.
