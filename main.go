package main

import (
	"github.com/blockloop/terraform-provider-consistent-macaddress/macaddress"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: macaddress.Provider,
	})
}
