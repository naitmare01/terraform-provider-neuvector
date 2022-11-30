package neuvector

import (
	"net/http"
)

var UserManagementPath = "user"
var EulaPath = "eula"
var EulaID = "0"

type Client struct {
	password   string
	username   string
	baseurl    string
	insecure   bool
	apitoken   string "omitempty"
	httpClient *http.Client
}

type UserManagementPasswordPost struct {
	Config struct {
		Fullname    string `json:"fullname,omitempty"`
		Password    string `json:"password,omitempty"`
		NewPassword string `json:"new_password,omitempty"`
	} `json:"config,omitempty"`
}

type UserManagementUserPost struct {
	User struct {
		BlockedForFailedLogin     bool   `json:"blocked_for_failed_login,omitempty"`
		BlockedForPasswordExpired bool   `json:"blocked_for_password_expired,omitempty"`
		DefaultPassword           bool   `json:"default_password,omitempty"`
		Email                     string `json:"email,omitempty"`
		Fullname                  string `json:"fullname,omitempty"`
		LastLoginAt               string `json:"last_login_at,omitempty"`
		LastLoginTimestamp        int    `json:"last_login_timestamp,omitempty"`
		Locale                    string `json:"locale,omitempty"`
		LoginCount                int    `json:"login_count,omitempty"`
		ModifyPassword            bool   `json:"modify_password,omitempty"`
		Password                  string `json:"password,omitempty"`
		Role                      string `json:"role,omitempty"`
		RoleDomains               string `json:"role_domains,omitempty"`
		Server                    string `json:"server,omitempty"`
		Timeout                   int    `json:"timeout,omitempty"`
		Username                  string `json:"username,omitempty"`
	} `json:"user,omitempty"`
}

type EulaPost struct {
	Eula struct {
		Accepted bool `json:"accepted,omitempty"`
	} `json:"eula,omitempty"`
}

type OIDCPost struct {
	Config struct {
		Name string `json:"name,omitempty"`
		Oidc struct {
			Issuer       string   `json:"issuer,omitempty"`
			ClientId     string   `json:"client_id,omitempty"`
			ClientSecret string   `json:"client_secret,omitempty"`
			GroupClaim   string   `json:"group_claim,omitempty"`
			Scopes       []string `json:"scopes,omitempty"`
			Enable       bool     `json:"enable,omitempty"`
			DefaultRole  string   `json:"default_role,omitempty"`
			RoleGroups   struct {
				Role   string     `json:"role,omitempty"`
				Groups [][]string `json:"groups,omitempty"`
			} `json:"role_groups,omitempty"`
			GroupMappedRoles GroupMappedRolesStruct `json:"group_mapped_roles,omitempty"`
		} `json:"oidc,omitempty"`
	} `json:"config,omitempty"`
}

type GroupMappedRolesStruct struct {
	Group       string `json:"group,omitempty"`
	GlobalRole  string `json:"global_role,omitempty"`
	RoleDomains struct {
		Role    string     `json:"role"` //named role_domain in rest of the package due to duplicate key name
		Domains [][]string `json:"domains"`
	} `json:"role_domains"`
}
