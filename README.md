# Terraform Provider: MAC Address Generator

This Terraform provider generates a MAC address based on a given seed string. It's a simple and lightweight solution for creating consistent MAC addresses within your Terraform configurations.

## Features

- Generates a MAC address from a given seed string.
- Produces a consistent MAC address for the same seed.
- Outputs MAC addresses in uppercase.

## Usage

[Terraform Registry](https://registry.terraform.io/providers/blockloop/consistent-macaddress/latest)

In your Terraform configuration, add the `macaddress` provider and use the `macaddress` data source to generate a MAC address:

```hcl
terraform {
  required_providers {
    consistent-macaddress = {
      source = "blockloop/consistent-macaddress"
      version = "0.2.16"
    }
  }
}

provider "consistent-macaddress" {}

data "macaddress" "foobar" {
  seed = "foobar"
}

output "mac_address" {
  value = data.macaddress.foobar.mac_address
}
```

## Testing

To run acceptance tests:

```
bash
TF_ACC=1 go test -v ./... -timeout 120m
```

To run unit tests:

```
bash
go test -v ./...
```

## Contributing

Contributions are welcome! If you find a bug or have a feature request, please create an issue. If you want to contribute code, feel free to open a pull request.
