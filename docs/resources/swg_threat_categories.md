---
layout: "metanetworks"
page_title: "Meta Networks: metanetworks_swg_threat_categories"
description: |-
  Provides a Threat Categories resource for SWG (Secure Web Gateway).
---

# Resource: metanetworks_swg_threat_categories

Provides a Threat Categories resource for SWG (Secure Web Gateway).

## Example Usage

```hcl
resource "metanetworks_swg_threat_categories" "example" {
  name              = "example"
  description       = "example threat category"
  types             = ["Bot","Brute Forcer"]
  confidence_level  = "HIGH"
  risk_level        = "HIGH"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the threat category
* `description` - (Optional) The description of the threat category
* `confidence_level` - (Optional) Confidence of the classification when the classification engine classifies a URL
* `countries` - (Optional) Access restricted countries. Enum by Alpha-2 code (ISO-3166). EG "AU" -> Australia, "US" -> United States
* `risk_level` - (Optional) Risk threshold that will not be tolerated while browsing URL categories under selected threat types. Enum: "LOW", "MEDIUM", "HIGH"
* `types` - (Required) Predefined threat types to protect against. Enum: "Abused TLD", "Bitcoin Related", "Blackhole", "Bot", "Brute Forcer", "Chat Server", "CnC", "Compromised", "DDoS Target", "Drive By Src", "Drop", "DynDNS", "EXE Source", "Fake AV", "IP Check", "Mobile CnC", "Mobile Spyware CnC", "Online Gaming", "P2P CnC", "P2P", "Parking", "Phishing", "Proxy", "Remote Access Service", "Scanner", "Self Signed SSL", "Spam", "Spyware CnC", "Tor", "Undesirable", "Utility", "VPN"

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of the resource.
* `created_at` - Creation timestamp.
* `modified_at` - Modification timestamp.
* `org_id` - The ID of the organization.
