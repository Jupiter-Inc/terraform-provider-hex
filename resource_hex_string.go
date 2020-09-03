package main

import (
	"encoding/hex"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceHexString() *schema.Resource {
	return &schema.Resource{
		Create: CreateHexString,
		Read:   schema.Noop,
		Delete: schema.RemoveFromState,

		Schema: map[string]*schema.Schema{
			"data": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"result": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func CreateHexString(d *schema.ResourceData, m interface{}) error {
	data := d.Get("data").(string)
	result := hex.EncodeToString([]byte(data))
	d.Set("result", result)
	d.SetId(data)
	return nil
}
