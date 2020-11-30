Terraform Provider for Meta Networks
======================================================

See the [Meta Networks Provider documentation](docs/index.md) to get started using the provider.

Requirements
------------

- [Terraform](https://www.terraform.io/downloads.html) 0.12.x or higher
- [Go](https://golang.org/doc/install) 1.13 (to build the provider plugin)
- [GOPATH](http://golang.org/doc/code.html#GOPATH)

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/mataneine/terraform-provider-metanetworks`

```shell
$ mkdir -p $GOPATH/src/github.com/mataneine; cd $GOPATH/src/github.com/mataneine
$ git clone git@github.com:mataneine/terraform-provider-metanetworks
```

Enter the provider directory and build the provider

```shell
$ cd $GOPATH/src/github.com/mataneine/terraform-provider-metanetworks
$ make build
```

Using the provider
----------------------

Enter the provider directory and install the provider

```shell
$ cd $GOPATH/src/github.com/mataneine/terraform-provider-metanetworks
$ make install
```
If you're building the provider, follow the instructions to [install it as a plugin.](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin) After placing it into your plugins directory,  run `terraform init` to initialize it.

See the [Meta Networks Provider documentation](docs/index.md) to get started using the provider.

Developing the Provider
---------------------------

Enter the provider directory.
```sh
$ cd $GOPATH/src/github.com/mataneine/terraform-provider-metanetworks
```

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/src/github.com/mataneine/terraform-provider-metanetworks/bin` directory.

```shell
$ make build
```

In order to test the provider, you can simply run `make test`.

```shell
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```shell
$ make testacc
```
