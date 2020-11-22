---
layout: "metanetworks"
page_title: "Meta Networks: metanetworks_mapped_service_alias"
description: |-
  Set an alias to a mapped service.
---

# Resource: metanetworks_mapped_service_alias

  Set an alias to a mapped service.

## Example Usage

```hcl
resource "metanetworks_mapped_service" "example" {
  name           = "example"
  mapped_service = "example.com"
}

resource "metanetworks_mapped_service_alias" "example" {
  mapped_service_id = metanetworks_mapped_service.example.id
  alias             = "example.com"
}
```

## Argument Reference

The following arguments are supported:

* `mapped_service_id` - (Required) The ID of the mapped service.
* `alias` - (Required) Alias name.
