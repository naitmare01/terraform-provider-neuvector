package neuvector

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// $ curl -s -k -H 'Content-Type: application/json' -H 'X-Auth-Token: c64125decb31e6d3125da45cba0f5025' https://127.0.0.1:10443/v1/user/admin -X PATCH -d '{"config":{"fullname":"admin","password":"admin","new_password":"NEWPASS"}}'

func resourceUserManagement() *schema.Resource {
	return &schema.Resource{
		Schema:        map[string]*schema.Schema{},
		CreateContext: resourceUserManagementCreate,
		ReadContext:   resourceUserManagementRead,
		UpdateContext: resourceUserManagementUpdate,
		DeleteContext: resourceUserManagementDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceUserManagementCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceUserManagementRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceUserManagementUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	//path := UserManagementPath
	//body := DealsBody(d)
	return diags
}

func resourceUserManagementDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}
