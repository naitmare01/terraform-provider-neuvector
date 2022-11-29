package neuvector

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEula() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"accepted": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		CreateContext: resourceEulaCreate,
		ReadContext:   resourceEulaRead,
		UpdateContext: resourceEulaUpdate,
		DeleteContext: resourceEulaDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceEulaCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	path := EulaPath
	body := EulaBody(d)
	resp, _, _, err := m.(*Client).SendRequest("POST", path, body, 200)

	if err != nil {
		return diag.FromErr(err)
	}

	var result map[string]any
	err = json.Unmarshal([]byte(resp), &result)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(EulaID)
	return resourceEulaRead(ctx, d, m)
}

func resourceEulaRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	path := EulaPath
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

	return diags
}

func resourceEulaUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	path := EulaPath
	body := EulaBody(d)
	_, _, _, err := m.(*Client).SendRequest("POST", path, body, 200)

	if err != nil {
		return diag.FromErr(err)
	}

	return resourceEulaRead(ctx, d, m)
}

func resourceEulaDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Not implemented
	var diags diag.Diagnostics
	return diags
}
