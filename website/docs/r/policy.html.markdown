---
layout: "metanetworks"
page_title: "Metanetworks: metanetworks_policy_resource"
sidebar_current: "docs-metanetworks-resource"
description: |-
  Create a new policy.
---

# metanetworks_policy_resource

Create a new policy.

## Example Usage

```hcl
resource "metanetworks_policy" "example_policy" {
  name = "an_example_policy name"
  destinations = [
    var.example_mapped_service.id,
    var.example_mapped_subnet.id
  ]
  protocol_groups = [
    var.protocol_groups.https.id,
  ]
  sources = [
    var.example_user.id,
    var.example_group.id
  ]
}
resource "metanetworks_mapped_service" "example_mapped_service" {
  name           = "example_mapped_service"
  mapped_service = "example_mapped_service.example.com"
}
output "example_mapped_service" {
  value = metanetworks_mapped_service.example_mapped_service
}
resource "metanetworks_mapped_subnets" "example_mapped_subnet" {
  name           = "example_mapped_subnet"
  mapped_subnets = ["172.16.1.0/28"]
}
output "example_mapped_subnet" {
  value = metanetworks_mapped_subnets.example_mapped_subnet
}
resource "metanetworks_protocol_group" "https" {
  name = "HTTPS"
  protocols {
    port  = 443
    proto = "tcp"
  }
}
output "https" {
  value = metanetworks_protocol_group.https
}
data "metanetworks_user" "example_user" {
  email = "example.user@example.com"
}
output "example_user" {
  value = data.metanetworks_user.example_user
}
data "metanetworks_group" "example_group" {
  name = "example group"
}
output "example_group" {
  value = data.metanetworks_group.example_group
}
```

## Argument Reference

The following arguments are supported:

* `description` - The Policy description.
* `name` - The Policy name.
* `destinations` - The Policy targets. (merge subnets & mapped services ????)
* `enabled` - Is the Policy enabled.
* `protocol_groups` - Protocols and Ports Restrictions
* `sources` - (List of ????) The Policy sources.

## Attributes Reference

The following attributes are exported:

* `created_at` - Creation timestamp.
* `modified_at` - Modification timestamp.
* `org_id` - The Organization ID.