# Data Source: neuvector_user_management

## Example Usage

```hcl
data "neuvector_user_management" "example" {
  username = "admin"
}

output "example" {
  value = data.neuvector_user_management.example
}
```

## Argument Reference

The following arguments are supported:

* **username** - (Required) The username.

## Attributes Reference

In addition to all argument, the following attributes are exported:

* **blocked_for_failed_login**
* **blocked_for_password_expired**
* **default_password**
* **email**
* **fullname**
* **id**
* **last_login_at**
* **last_login_timestamp**
* **locale**
* **login_count**
* **modify_password**
* **password**
* **role**
* **role_domains**
* **server**
* **timeout**
* **username**
