package neuvector

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceUserManagement() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"fullname": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"new_password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"blocked_for_failed_login": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"blocked_for_password_expired": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"default_password": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"email": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_login_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_login_timestamp": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"locale": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"login_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"modify_password": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"role": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"role_domains": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"server": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"timeout": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
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
	return resourceUserManagementRead(ctx, d, m) //Haven't found a way to create users yet..
}

func resourceUserManagementRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	path := UserManagementPath + "/" + d.Get("username").(string)
	resp, _, _, err := m.(*Client).SendRequest("GET", path, nil, 200)

	if err != nil {
		return diag.FromErr(err)
	}

	var result map[string]any
	err = json.Unmarshal([]byte(resp), &result)
	if err != nil {
		return diag.FromErr(err)
	}

	blocked_for_failed_login := result["user"].(map[string]interface{})["blocked_for_failed_login"]
	blocked_for_password_expired := result["user"].(map[string]interface{})["blocked_for_password_expired"]
	default_password := result["user"].(map[string]interface{})["default_password"]
	email := result["user"].(map[string]interface{})["email"]
	fullname := result["user"].(map[string]interface{})["fullname"]
	last_login_at := result["user"].(map[string]interface{})["last_login_at"]
	last_login_timestamp := result["user"].(map[string]interface{})["last_login_timestamp"]
	locale := result["user"].(map[string]interface{})["locale"]
	login_count := result["user"].(map[string]interface{})["login_count"]
	modify_password := result["user"].(map[string]interface{})["modify_password"]
	role := result["user"].(map[string]interface{})["role"]
	role_domains := result["user"].(map[string]interface{})["role_domains"]
	server := result["user"].(map[string]interface{})["server"]
	timeout := result["user"].(map[string]interface{})["timeout"]
	username := result["user"].(map[string]interface{})["username"]

	d.Set("blocked_for_failed_login", blocked_for_failed_login)
	d.Set("blocked_for_password_expired", blocked_for_password_expired)
	d.Set("default_password", default_password)
	d.Set("email", email)
	d.Set("fullname", fullname)
	d.Set("last_login_at", last_login_at)
	d.Set("last_login_timestamp", last_login_timestamp)
	d.Set("locale", locale)
	d.Set("login_count", login_count)
	d.Set("modify_password", modify_password)
	d.Set("role", role)
	d.Set("role_domains", role_domains)
	d.Set("server", server)
	d.Set("timeout", timeout)
	d.Set("username", username)

	d.SetId(username.(string))

	return diags
}

func resourceUserManagementUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	path := UserManagementPath + "/" + d.Get("username").(string)
	body := UserManagementPasswordBody(d)
	_, _, _, err := m.(*Client).SendRequest("PATCH", path, body, 200)

	if err != nil {
		return diag.FromErr(err)
	}

	new_password := d.Get("new_password").(string)
	Authenticate(m.(*Client))
	d.Set("password", new_password)

	return resourceUserManagementRead(ctx, d, m)
}

func resourceUserManagementDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// Not implemented
	return diags
}
