---
layout: "metanetworks"
page_title: "Meta Networks: metanetworks_version_controls"
description: |-
  Provides a version controls resource.
---

# Resource: metanetworks_version_controls

  Provides a version controls resource.

## Example Usage

```hcl
resource "metanetworks_version_controls" "example" {
  name                  = "Default_Org_Device_Setting"
  description           = "Example Description"
  apply_on_org          = true
  windows_policy = {
      mode = "latest_beta"
  }
  macos_policy = {
      mode = "latest_stable"
  }
  linux_policy = {
      mode = "specific_version"
      version = "3.7.6"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required). The name of the device setting.
* `description` - (Optional). The description of the device setting.
* `enabled` - (Optional) default=true.
* `apply_on_org` - (Optional; Required if `apply_to_entities` is omitted). Indicates whether this version control applies to the org. Note: this attribute overrides apply_to_entities.
* `apply_to_entities` - (Optional) - List of entities (`sources`) which will use this version control
* `exempt_entities` - (Optional) - Sources to exclude from version controls policy
* `windows_policy` (Optional) - Required if `mode` parameter is used.
* `macos_policy` (Optional) - Required if `mode` parameter is used.
* `linux_policy` (Optional) - Required if `mode` parameter is used.
* `mode` - (Optional) - `"disable" "specific_version" "latest_stable" "latest_beta"`
* `version` - (Optional): Required only if `mode` = `"specific_version"`

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of the resource.
* `created_at` - Creation timestamp.
* `modified_at` - Modification timestamp.