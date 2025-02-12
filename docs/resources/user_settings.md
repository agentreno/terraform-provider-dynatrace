---
layout: ""
page_title: dynatrace_user_settings Resource - terraform-provider-dynatrace"
description: |-
  The resource `dynatrace_user_settings` covers user settings of an individual user
---

# dynatrace_user_settings (Resource)

-> This resource is excluded by default in the export utility since it is scoped for an individual user and requires a personal access token.

## Dynatrace Documentation

- Settings API - https://www.dynatrace.com/support/help/dynatrace-api/environment-api/settings (schemaId: `builtin:user-settings`)

## Export Example Usage

- `terraform-provider-dynatrace -export dynatrace_user_settings` downloads user settings of the individual user

The full documentation of the export feature is available [here](https://registry.terraform.io/providers/dynatrace-oss/dynatrace/latest/docs/guides/export-v2).

## Resource Example Usage

```
resource "dynatrace_user_settings" "#name#" {
  language = "en"
  region   = "auto"
  scope    = "user-terraform@dynatrace.com"
  theme    = "auto"
  timezone = "UTC"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `language` (String) Possible Values: `Auto`, `En`, `Ja`
- `region` (String) Region
- `scope` (String) The scope of this setting (user, userdefaults)
- `theme` (String) Possible Values: `Auto`, `Dark`, `Light`
- `timezone` (String) Timezone

### Read-Only

- `id` (String) The ID of this resource.
 