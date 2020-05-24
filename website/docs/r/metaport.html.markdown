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

* `name` - (required)
* `description` - 
* `enabled` - (Not allowed for mapped service and mapped subnet)
* `mapped_elements` - network element IDs
* `allow_support` - 

## Attributes Reference

The following attributes are exported:

* `created_at` - 
* `dns_name` - 
* `expires_at` - 
* `modified_at` - 
* `network_element_id` - 
* `org_id` - 

