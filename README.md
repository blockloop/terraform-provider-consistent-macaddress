# Terraform Provider: MAC Address Generator

This Terraform provider generates a MAC address based on a given seed string. It's a simple and lightweight solution for creating consistent MAC addresses within your Terraform configurations.

## Features

- Generates a MAC address from a given seed string.
- Produces a consistent MAC address for the same seed.
- Outputs MAC addresses in uppercase.

## Usage

In your Terraform configuration, add the `macaddress` provider and use the `macaddress_generator` data source to generate a MAC address:

```
hcl
terraform {
  required_providers {
    macaddress = {
      source = "github.com/yourusername/terraform-provider-macaddress"
    }
  }
}

provider "macaddress" {}

data "macaddress_generator" "example" {
  seed = "my-seed-string"
}

output "mac_address" {
  value = data.macaddress_generator.example.mac_address
}
```

Replace `yourusername` with your actual GitHub username. Run `terraform init` and `terraform apply` to generate a MAC address.

## Testing

To run acceptance tests:

```
bash
TF_ACC=1 go test -v github.com/yourusername/terraform-provider-macaddress/macaddress -timeout 120m
```

To run unit tests:

```
bash
go test -v github.com/yourusername/terraform-provider-macaddress/macaddress
```

Replace `yourusername` with your actual GitHub username.

## Contributing

Contributions are welcome! If you find a bug or have a feature request, please create an issue. If you want to contribute code, feel free to open a pull request.