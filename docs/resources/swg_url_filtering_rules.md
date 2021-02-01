---
layout: "metanetworks"
page_title: "Meta Networks: metanetworks_swg_url_filtering_rules"
description: |-
  Provides a Url Filtering Rule for SWG (Secure Web Gateway).
---

# Resource: metanetworks_swg_url_filtering_rules

Provides a Url Filtering Rule for SWG (Secure Web Gateway).

## Example Usage

```hcl
resource "metanetworks_swg_url_filtering_rules" "example" {
  name                          = "example"
  description                   = "example url filtering rule"
  action                        = "ISOLATION"
  priority                      = 1
  enabled                       = true
  sources                       = ["grp-exampleid"]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the URL Filtering Rule.
* `action` - (Required) Action to take when rule conditions are met. Enum: "ISOLATION", "BLOCK", "LOG"
* `description` - (Optional) The description of the URL Filtering Rule.
* `advanced_threat_protection` - (Optional) Whether to use ATP algorithms.
* `enabled` - (Optional, Default: true) Enable or Disable rule.
* `exempt_sources` - (Optional) Exclude entities from rule when applying to groups.
* `sources` - (Required) Entities to apply rule to.
* `forbidden_content_categories` - (Optional) <= 5. Unique list of content category ids for restriction.
* `priority` - (Required) 1-5000. Position of the rule. Lower numbers = higher priority.
* `threat_category` - (Required) Threat Category ID as string for restriction.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of the resource.
* `created_at` - Creation timestamp.
* `modified_at` - Modification timestamp.
* `org_id` - The ID of the organization.
