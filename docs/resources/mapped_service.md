---
layout: "metanetworks"
page_title: "Metanetworks: metanetworks_mapped_service_resource"
sidebar_current: "docs-metanetworks-resource"
description: |-
  Create a new mapped_service.
---

# metanetworks_mapped_service_resource

Create a new mapped_service.

## Example Usage

```hcl
resource "metanetworks_mapped_service" "example" {
  name = "an_example_mapped_service name"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (required) The Mapped Service Name.
* `mapped_service` - Mapped Service IP or Hostname
* `description` - Mapped Service Description

## Attributes Reference

The following attributes are exported:

* `tags` - List ???? of tags associated with the Mapped Service.
* `created_at` - Creation timestamp.
* `dns_name` - New Domain Name
* `expires_at` - Expiry timesptamp.
* `modified_at` - Modification timestamp.
* `org_id` - ID of the Organization. ????
* `aliases` - Mapped Service IP or Hostname
* `net_id` - ????
* `type` - ["Device" | "Service" | "Mapped Service" | "Mapped Subnet"]
* `version` - [???? couldn't find in api????]