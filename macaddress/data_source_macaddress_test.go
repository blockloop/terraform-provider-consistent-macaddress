package macaddress

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceMacAddressGenerator(t *testing.T) {
	resourceName := "data.macaddress.test"
	seed := "test-seed"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMacAddressConfig(seed),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "seed", seed),
					resource.TestCheckResourceAttrSet(resourceName, "mac_address"),
				),
			},
		},
	})
}

func testAccDataSourceMacAddressConfig(seed string) string {
	return fmt.Sprintf(`
		provider "macaddress" {}

		data "macaddress" "test" {
		  seed = "%s"
		}
	`, seed)
}
