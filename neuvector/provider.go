package neuvector

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("NEUVECTOR_USERNAME", nil),
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("NEUVECTOR_PASSWORD", nil),
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"insecure": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"neuvector_eula": dataSourceEula(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	nv_user := d.Get("username").(string)
	nv_pass := d.Get("password").(string)
	nv_url := d.Get("url").(string)
	nv_insecure := d.Get("insecure").(bool)
	return NewClient(nv_user, nv_pass, nv_url, nv_insecure), nil
}
