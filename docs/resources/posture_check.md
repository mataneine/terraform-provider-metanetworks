---
layout: "metanetworks"
page_title: "Meta Networks: metanetworks_posture_check"
description: |-
  Provides a posture check resource.
---

# Resource: metanetworks_posture_check

Provides a posture check resource.

## Example Usage

```hcl
resource "metanetworks_metaport" "example" {
  name                  = "CrowdStrike Posture Check"
  description           = "Example Description"
  apply_on_org          = true
  osquery               = "select * from services where name='CSFalconService' and status='RUNNING';"
  platform              = "Windows"
  enabled               = true
  action                = "NONE"
  when                  = ["PRE_CONNECT"]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required). The name of the posture check.
* `description` - (Optional). The description of the posture check.
* `action`	- (Optional). What happens when a posture check is failed. Values: `DISCONNECT`, `NONE`
* `apply_on_org` - (Optional; Required if `sources` is omitted). Applies setting to entire organization.
* `sources` - (Optional; Required if `apply_on_org` is omitted). Applies setting to specified sources.
* `exempt_sources` (Optional). Sources to exclude from posture check.
* `osquery` - (Optional). OSQuery string to perform posture check.
* `enabled` - (Optional) default=true.
* `platform` - (Optional). Platform that the posture check is for. Values: `Android`, `macOS`, `iOS`, `Linux`, `Windows`, `ChromeOS`
* `user_message_on_fail` - (Optional). Failure message to display when posture check fails.
* `interval` - (Optional; Required if `when` contains `PERIODIC`). Time in *minutes* between checks. Values: `5-60`
* `check` - (Optional; instead of `osquery`). Templated scenario to posture check for. Requires two arguments
** `min_version` - (Optional). String of the version.
** `type` - (Required). Values: `jailbroken_rooted`, `screen_lock_enabled`, `minimum_app_version`, `minimum_os_version`, `malicious_app_detection`, `developer_mode_enabled`.
* `when` - (Required). When the posture check should run. Values: `PRE_CONNECT`, `PERIODIC`.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of the resource.
* `created_at` - Creation timestamp.
* `modified_at` - Modification timestamp.
