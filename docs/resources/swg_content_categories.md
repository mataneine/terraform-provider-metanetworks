---
layout: "metanetworks"
page_title: "Meta Networks: metanetworks_swg_content_categories"
description: |-
  Provides a Content Categories resource for SWG (Secure Web Gateway).
---

# Resource: metanetworks_swg_content_categories

Provides a Content Categories resource for SWG (Secure Web Gateway).

## Example Usage

```hcl
resource "metanetworks_swg_content_categories" "example" {
  name                 = "example"
  description          = "example content category"
  types                = ["Advertising","Games"]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the mapped service.
* `description` - (Optional) The description of the mapped service.
* `confidence_level` - (Optional) Degree of confidence (threshold) that must be met when the classification engine decides on URL classification. Enum: "LOW", "MEDIUM", "HIGH"
* `forbid_uncategorized_urls` - (Optional) Boolean. Forbuid access to uncategorized URLs.
* `types` - (Optional) Array of strings with category types to restrict. Enum: "Abortion", "Adult Sex Education", "Advertising", "Alcohol Tobacco", "Anonymizer", "Blogs", "Computer Hacking", "Dead Sites", "Drugs", "Education", "Email Host", "Finance", "Food", "Gambling", "Games", "Government", "Health", "Hobbies Interests", "Illegal Or Questionable", "Job Employment", "Lingerie Bikini", "Military", "Militancy Hate And Extremism", "Music", "News And Media", "Nudity", "Politics", "Pornography", "Portals", "Real Estate", "Religion", "Search", "Shopping And Auctions", "Social Networking", "Society And Lifestyle", "Software Technology", "Sports", "Streaming Media", "Television Movies", "Translator", "Travel", "Vehicles", "Violence", "Weapons"
* `urls` - (Optional) URLs to include in the custom category.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of the resource.
* `created_at` - Creation timestamp.
* `modified_at` - Modification timestamp.
* `org_id` - The ID of the organization.
