---
layout: "metanetworks"
page_title: "Meta Networks: metanetworks_version_controls"
description: |-
  Provides a posture check resource.
---

# Resource: metanetworks_version_controls

Provides a version controls resource.

## Example Usage

```hcl
resource "metanetworks_version_controls" "example" {
  name         = "Example"
  description  = "Example"
  enabled      = true
  apply_to_org = true
  windows_policy = {
    mode = "lastest_stable"
  }
  macos_policy = {
    mode = "lastet_beta"
  }
  linux_policy = {
    mode = "specific_version"
    version = "3.7.7"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required): The name of the posture check.
* `description` - (Optional): The description of the version control profile.
* `enabled` - (Optional): default=true.
* `apply_on_org` - (Optional: default=true. Required if `sources` is omitted). Applies setting to entire organization. Note: this attribute overrides apply_to_entities
* `apply_to_entities` - (Optional: List of entities which will use this version control
* `exempt_entities` - (Optional): List of entities `sources` to exclude from version control profile.
* `linux_policy` - (Optional): Apply version control to Linux operating system.
* `macos_policy` - (Optional): Apply version control to macOS operating system.
* `windows_policy` - (Optional): Apply version control to Windows operating system.
* `mode` - (Optional): The OS policy mode `"disable" "specific_version" "latest_stable" "latest_beta"`
* `version` - (Required): Required if `mode = specific_version`

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of the resource.
* `created_at` - Creation timestamp.
* `modified_at` - Modification timestamp.
