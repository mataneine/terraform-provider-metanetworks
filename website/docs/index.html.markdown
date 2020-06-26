---
layout: "metanetworks"
page_title: "Provider: metanetworks"
sidebar_current: "docs-metanetworks-index"
description: |-
  The metanetworks provider is used to interact with metanetworks resources.
---

# Metanetworks Provider

Use this paragraph to give a high-level overview of your provider, and any configuration it requires.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
provider "metanetworks" {
  org = "example_organization"
}

# Example resource configuration
resource "metanetworks_resource" "example" {
  # ...
}
```

## Authentication

The Metanetworks provider offers a flexible means of providing credentials for
authentication. The following methods are supported, in this order, and
explained below:

- Static credentials
- Environment variables
- Configuration file

### Static credentials

!> **Warning:** Hard-coding credentials into any Terraform configuration is not
recommended, and risks secret leakage should this file ever be committed to a
public version control system.

Static credentials can be provided by adding an `api_key` and `api_secret`
in-line in the Metanetworks provider block:

Usage:

```hcl
provider "metanetworks" {
  org        = "example_organization"
  access_key = "my-api-key"
  secret_key = "my-api-secret"
}
```

### Environment variables

You can provide your credentials via the `METANETWORKS_API_KEY`,
`METANETWORKS_API_SECRET` and `METANETWORKS_ORG`, environment variables,
representing your Metanetworks Access Key, Secret Key and the Organization name respectively.

Usage:

```hcl
provider "metanetworks" {}
```

```sh
$ export METANETWORKS_API_KEY="anapikey"
$ export METANETWORKS_API_SECRET="asecretkey"
$ export METANETWORKS_ORG="example_organization"
$ terraform plan
```

### Configuration file

You can use a configuration file to specify your credentials. The
file location is `$HOME/.metanetworks/credentials.json` on Linux and OS X, or
`"%USERPROFILE%\.metanetworks/credentials.json"` for Windows users.
If we fail to detect credentials inline, or in the environment, Terraform will check
this location.

Usage:

```hcl
provider "metanetworks" {}
```

credentials.json file:
```json
{
	"api_key":    "my-api-key",
	"api_secret": "my-api-secret",
	"org":        "example_organization"
}
```