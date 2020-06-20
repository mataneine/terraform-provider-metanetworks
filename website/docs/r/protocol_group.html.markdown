---
layout: "metanetworks"
page_title: "Metanetworks: metanetworks_protocol_group_resource"
sidebar_current: "docs-metanetworks-resource"
description: |-
  Modify a protocol group of the organization.
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

* `description` - Protocol Group Description.
* `name` - (required) Protocol Group Name.
* `protocols` - List ???? of Protocols ( and Ports ) to attach to the Protocol Group.
  * `proto` - (required) [ "ICMP' | "TCP" | "UDP" ]
  * `port` - port number (no 'to_port' ????)

## Attributes Reference

The following attributes are exported:

* `created_at` - Creation timestamp.
* `modified_at` - Modification timestamp.
* `org_id` - The ID of the Protocol Group.
* `read_only` - Protocol Group is read only.
