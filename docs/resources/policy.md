---
layout: "metanetworks"
page_title: "Meta Networks: metanetworks_policy"
description: |-
  Provides a policy resource.
---

# Resource: metanetworks_policy

Provides a policy resource.

## Example Usage

```hcl
resource "metanetworks_mapped_service" "example" {
  name           = "example"
  mapped_service = "example.com"
}

resource "metanetworks_mapped_subnets" "example" {
  name           = "example"
  mapped_subnets = ["172.16.1.0/28"]
}

resource "metanetworks_protocol_group" "https" {
  name = "HTTPS"
  protocols {
    from_port = 443
    to_port   = 443
    proto     = "tcp"
  }
}

data "metanetworks_user" "example" {
  email = "user@example.com"
}

data "metanetworks_group" "example" {
  name = "example"
}

resource "metanetworks_policy" "example" {
  name            = " example"
  destinations    = [
    metanetworks_mapped_service.example.id,
    metanetworks_mapped_subnets.example.id
  ]
  protocol_groups = [
    metanetworks_protocol_group.https.id,
  ]
  sources         = [
    data.metanetworks_user.example.id,
    data.metanetworks_group.example.id
  ]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the policy.
* `description` - (Optional) The description of the policy.
* `destinations` - (Optional) Set of users and/or groups/devices/mapped subnets/mapped services to attach to the policy.
* `enabled` - (Optional) default=true.
* `protocol_groups` - (Optional) Set of protocol groups.
* `exempt_sources` - (Optional) Set of users and/or groups/devices/mapped subnets/mapped services to exempt from the policy.
* `sources` - (Optional) Set of users and/or groups/devices/mapped subnets/mapped services to attach to the policy.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of the resource.
* `created_at` - Creation timestamp.
* `modified_at` - Modification timestamp.
* `org_id` - The ID of the organization.
