package neuvector

import (
	"net/http"
)

var UserManagementPath = "user/admin"

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
