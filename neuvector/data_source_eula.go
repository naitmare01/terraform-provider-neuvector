package neuvector

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEula() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceEulaRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"accepted": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func dataSourceEulaRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	path := "eula"
	resp, _, _, err := m.(*Client).SendRequest("GET", path, nil, 200)

	if err != nil {
		return diag.FromErr(err)
	}

	var result map[string]any
	err = json.Unmarshal([]byte(resp), &result)
	if err != nil {
		return diag.FromErr(err)
	}

	accepted := result["eula"].(map[string]interface{})["accepted"]

	d.Set("accepted", accepted)

	d.SetId("1")

	return diags
}
