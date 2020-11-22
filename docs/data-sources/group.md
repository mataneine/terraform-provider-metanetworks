---
layout: "metanetworks"
page_title: "Meta Networks: metanetworks_group"
description: |-
  Return a group of the organization.
---

# Data Source: metanetworks_group

Return a group of the organization.

## Example Usage

```hcl
data "metanetworks_group" "example" {
  name = "example"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the group.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `description` - The description of the group.
* `expression` - Allows grouping entities by their tags. Filtering by tag value is also supported if provided. Supported operations: AND, OR, XOR, parenthesis.
* `roles` - The group roles.
* `users` - The group users.
* `provisioned_by` - Groups can be provisioned in the system either by locally creating the groups from the Admin portal or API. Another, more common practice, is to provision groups from an organization directory service, by way of SCIM or LDAP protocols.
* `created_at` - Creation Timestamp.
* `modified_at` - Modification Timestamp.
* `org_id` - The ID of the organization.
