# Resource: neuvector_eula

## Example Usage

```hcl
resource "neuvector_eula" "example" {
  accepted = true
}
```

## Argument Reference

The following arguments are supported:

* `accepted` - (Required) `true`/`false` if the EULA is accepted or not.

## Attributes Reference

In addition to all argument, the following attributes are exported:

* `id` - The ID of this resource

## Import

Pipedrive deals can be imported using the `id`(`id` will always be 0) eg,

`terraform import neuvector_eula.example 0`
