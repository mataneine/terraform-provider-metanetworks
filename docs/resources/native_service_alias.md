---
layout: "metanetworks"
page_title: "Meta Networks: metanetworks_native_service_alias"
description: |-
  Set an alias to a native service.
---

# Resource: metanetworks_native_service_alias

  Set an alias to a native service.

## Example Usage

```hcl
resource "metanetworks_native_service" "example" {
  name = "example"
}

resource "metanetworks_native_service_alias" "example" {
  native_service_id = metanetworks_native_service.example.id
  alias             = "example.com"
}
```

## Argument Reference

The following arguments are supported:

* `native_service_id` - (Required) The ID of the native service.
* `alias` - (Required) Domain name.
