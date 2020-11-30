---
layout: "metanetworks"
page_title: "Meta Networks: metanetworks_group_data_source"
description: |-
  Return a group of the organization.
---

# Data Source: metanetworks_group

Return a group of the organization.

## Example Usage

```hcl
data "metanetworks_locations" "all" {}

locals {
  locations = {
    for location in data.metanetworks_locations.all.locations :
    location.city => location
  }
}
```

## Attributes Reference

The following attributes are exported:

* `locations` - List of locations. Fields documented below.

### Locations Attributes

For **locations** the following attributes are exported.

* `city` - The city of the location.
* `country` - The country of the location.
* `latitude` - The latitude of the location.
* `longitude` - The longitude of the location.
* `name` - The name of the location.
* `state` - The state of the location.
* `status` - The status of the location. Valid values are `Operational`, `Degraded`, `Outage`, and `Maintenance`.
