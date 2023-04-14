package macaddress

import (
	"context"
	"crypto/md5"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"macaddress_generator": dataSourceMacAddress(),
		},
	}
}

func dataSourceMacAddress() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceMacAddressRead,

		Schema: map[string]*schema.Schema{
			"seed": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"mac_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceMacAddressRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	seed := d.Get("seed").(string)

	macAddress, err := generateMacAddress(seed)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(macAddress)
	d.Set("mac_address", macAddress)

	return nil
}

func generateMacAddress(seed string) (string, error) {
	hash := md5.Sum([]byte(seed))
	macAddress := fmt.Sprintf("02:%02x:%02x:%02x:%02x:%02x", hash[0], hash[1], hash[2], hash[3], hash[4])
	macAddress = strings.ToUpper(macAddress)

	log.Printf("[DEBUG] Generated MAC address: %s", macAddress)

	return macAddress, nil
}

