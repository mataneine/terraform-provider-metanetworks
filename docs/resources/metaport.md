---
layout: "metanetworks"
page_title: "Metanetworks: metanetworks_metaport_resource"
sidebar_current: "docs-metanetworks-resource"
description: |-
  Create a new metaport.
---

# metanetworks_metaport_resource

Create a new metaport

## Example Usage

```hcl
resource "metanetworks_metaport" "example-metaport" {
  name    = "example-metaport"
  enabled = false
}
```

## Argument Reference

The following arguments are supported:

* `name` - (required) The MetaPort name.
* `description` - The MetaPort description.
* `enabled` - (Not allowed for mapped service and mapped subnet ????)
* `mapped_elements` - Network element IDs to attach to the Metaport
* `allow_support` ???? attribute - ????

## Attributes Reference

The following attributes are exported:

* `created_at` - Creation timestamp.
* `dns_name` - ????
* `expires_at` - Expiry timestamp.
* `modified_at` - Modification timestamp.
* `network_element_id` - The ID of the Metaport ????
* `org_id` - The ID of the organization
