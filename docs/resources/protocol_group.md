---
layout: "metanetworks"
page_title: "Meta Networks: metanetworks_protocol_group"
description: |-
  Provides a protocol group resource.
---

# Resource: metanetworks_protocol_group

Provides a protocol group resource.

## Example Usage

```hcl
resource "metanetworks_protocol_group" "https" {
  name = "HTTPS"
  protocols {
    from_port = 443
    to_port   = 443
    proto     = "tcp"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the protocol group.
* `description` - (Optional) The description of the protocol group.
* `protocols` - (Optional) List of Protocols to attach to the protocol group.  Fields documented below.

### Protocols Arguments

For **protocols** the following attributes are supported.

  * `proto` - (Required) The protocol. Valid values are `tcp`, `udp` and `icmp`.
  * `from_port` - (Required) From port number.
  * `to_port` - (Required) To port number.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of the resource.
* `read_only` - Protocol Group is read only.
* `created_at` - Creation timestamp.
* `modified_at` - Modification timestamp.
* `org_id` - The ID of the organization.
