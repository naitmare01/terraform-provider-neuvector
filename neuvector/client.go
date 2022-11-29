package neuvector

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// NewClient creates common settings
func NewClient(neuvector_username string, neuvector_password string, neuvector_baseurl string, insecure_url bool) *Client {
	return &Client{
		username:   neuvector_username,
		password:   neuvector_password,
		baseurl:    neuvector_baseurl,
		insecure:   insecure_url,
		httpClient: &http.Client{},
	}
}

func (c *Client) SendRequest(method string, path string, payload interface{}, statusCode int) (value string, respheaders string, respCode int, err error) {
	Authenticate(c)

	apiurl := c.baseurl + "/" + path + "/"
	client := &http.Client{Timeout: 10 * time.Second}

	b := new(bytes.Buffer)
	err = json.NewEncoder(b).Encode(payload)
	if err != nil {
		return "", "", 0, err
	}

	if c.insecure {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	req, err := http.NewRequest(method, apiurl, b)
	if err != nil {
		return "", "", 0, err
	}

	req.Header.Add("Content-Type", "application/json")

	if path != "auth" {
		req.Header.Add("X-Auth-Token", c.apitoken)
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", "", resp.StatusCode, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", resp.StatusCode, err
	}
	resp.Body.Close()

	strbody := string(body)

	respHeaders := resp.Header
	headers, err := json.Marshal(respHeaders)
	if err != nil {
		return "", "", resp.StatusCode, err
	}

	if statusCode != 0 {
		if resp.StatusCode != statusCode {
			return "", "", 0, fmt.Errorf("[ERROR] unexpected status code got: %v expected: %v \n %v", resp.StatusCode, statusCode, strbody)
		}
	}

	DeleteAuthentication(c) //Important - If you are making REST API calls, please be sure make a DELETE call against /v1/auth when done. There is a maximum of 32 concurrent sessions for each user. If this is exceeded, an authentication failure will occur.

	return strbody, string(headers), resp.StatusCode, nil
}

func Authenticate(c *Client) (value string, respheaders string, respCode int, err error) {
	authpath := "auth"
	authmethod := "POST"
	authapiurl := c.baseurl + "/" + authpath + "/"
	authpayload := strings.NewReader(`{
		"password": {
			"username": "` + c.username + `",
			"password": "` + c.username + `"
		}
	}`)
	authclient := &http.Client{}

	if c.insecure {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	authreq, err := http.NewRequest(authmethod, authapiurl, authpayload)

	if err != nil {
		return "", "", 0, err
	}

	authreq.Header.Add("Content-Type", "application/json")

	res, err := authclient.Do(authreq)
	if err != nil {
		return "", "", 0, err
	}

	authbody, err := io.ReadAll(res.Body)
	if err != nil {
		return "", "", 0, err
	}
	res.Body.Close()
	string_body := string(authbody)

	var result map[string]any
	json.Unmarshal([]byte(string_body), &result)
	auth := result["token"].(map[string]interface{})["token"]
	stringauth := fmt.Sprintf("%v", auth)

	c.apitoken = stringauth

	return
}

func DeleteAuthentication(c *Client) (err error) {
	path := "auth"
	method := "DELETE"
	apiurl := c.baseurl + "/" + path + "/"
	authclient := &http.Client{}

	if c.insecure {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	req, err := http.NewRequest(method, apiurl, nil)

	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Auth-Token", c.apitoken)

	res, err := authclient.Do(req)
	if err != nil {
		return err
	}

	res.Body.Close()

	c.apitoken = ""

	return
}

func UserManagementPasswordBody(d *schema.ResourceData) UserManagementPasswordPost {
	body := UserManagementPasswordPost{}
	body.Config.Fullname = d.Get("fullname").(string)
	body.Config.Password = d.Get("password").(string)
	body.Config.NewPassword = d.Get("new_password").(string)

	return body
}

func UserManagementUserBody(d *schema.ResourceData) UserManagementUserPost {
	body := UserManagementUserPost{}
	body.User.BlockedForFailedLogin = d.Get("blocked_for_failed_login").(bool)
	body.User.BlockedForPasswordExpired = d.Get("blocked_for_password_expired").(bool)
	body.User.DefaultPassword = d.Get("default_password").(bool)
	body.User.Email = d.Get("email").(string)
	body.User.Fullname = d.Get("fullname").(string)
	body.User.LastLoginAt = d.Get("last_login_at").(string)
	body.User.LastLoginTimestamp = d.Get("last_login_timestamp").(int)
	body.User.Locale = d.Get("locale").(string)
	body.User.LoginCount = d.Get("login_count").(int)
	body.User.ModifyPassword = d.Get("modify_password").(bool)
	body.User.Password = d.Get("password").(string)
	body.User.Role = d.Get("role").(string)
	body.User.RoleDomains = d.Get("role_domains").(string)
	body.User.Server = d.Get("server").(string)
	body.User.Timeout = d.Get("timeout").(int)
	body.User.Username = d.Get("username").(string)

	return body
}
