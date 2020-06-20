---
layout: "metanetworks"
page_title: "Metanetworks: metanetworks_mapped_service_alias_resource"
sidebar_current: "docs-metanetworks-resource"
description: |-
  Create an Alias to a Mapped Service Network Element.
---

# metanetworks_mapped_service_alias_resource

Create an Alias to a Mapped Service Network Element

## Example Usage

```hcl
resource "metanetworks_mapped_service_alias" "example" {
  mapped_service_id = metanetworks_mapped_service.example.id
  alias             = "example.com"
}
resource "metanetworks_mapped_service" "example" {
  name = "example"
  mapped_service = example.com"
}
output "example" {
  value = metanetworks_mapped_service.example
}
```

## Argument Reference

The following arguments are supported:

* `mapped_service_id` - (required) the ID of the Mapped Service Network Element.
* `alias` - (required) Alias name
