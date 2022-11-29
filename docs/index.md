# Neuvector Provider

This provider gives Terraform the ability to work with [Neuvector](https://neuvector.com).

## Authentication

```hcl
provider "neuvector" {
  username = "admin"
  password = "admin_password"
  url      = "https://127.0.0.1:10443/v1"
  insecure = true
}
```

## Argument Reference

* **username** - (Required) Admin username.
* **password** - (Required) Admin password.
* **url** - (Required) url and port to API.
* **insecure** - (Required) If the connection is insecure.
