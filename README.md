# terraform-provider-metanetworks

Terraform provider for Metanetworks, includes an API Client.
Very rough around the edges. 

## Building

`go build`

## Installing

Copy the terraform-provider-metanetworks binary to ~/.terraform.d/plugins
You can then do `terraform init` as usual.

## Authentication

Create an API key through the metanetworks web portal, and place the credentials in a file:
~/.metanetworks/credentials.json
The format of the file is:

```json
{
    "api_key": "<key>",
    "api_secret": "<secret>",
    "org": "<company>"
}
```
