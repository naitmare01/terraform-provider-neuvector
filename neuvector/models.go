package neuvector

import (
	"net/http"
)

type Client struct {
	password   string
	username   string
	baseurl    string
	insecure   bool
	apitoken   string "omitempty"
	httpClient *http.Client
}
