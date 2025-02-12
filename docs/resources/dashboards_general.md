---
layout: ""
page_title: dynatrace_dashboards_general Resource - terraform-provider-dynatrace"
description: |-
  The resource `dynatrace_dashboards_general` covers configuration for general dashboard settings
---

# dynatrace_dashboards_general (Resource)

## Dynatrace Documentation

- Dashboards - https://www.dynatrace.com/support/help/observe-and-explore/dashboards

- Settings API - https://www.dynatrace.com/support/help/dynatrace-api/environment-api/settings (schemaId: `builtin:dashboards.general`)

## Export Example Usage

- `terraform-provider-dynatrace -export dynatrace_dashboards_general` downloads all existing general dashboard settings

The full documentation of the export feature is available [here](https://registry.terraform.io/providers/dynatrace-oss/dynatrace/latest/docs/guides/export-v2).

## Resource Example Usage

```terraform
resource "dynatrace_dashboards_general" "#name#" {
  enable_public_sharing = false
  default_dashboard_list {
    default_dashboard {
      dashboard = "00000000-0000-0000-0000-000000000000"
      user_group = "00000000-0000-0000-0000-000000000000"
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `enable_public_sharing` (Boolean) Allow users to grant anonymous access to dashboards. No sign-in will be required to view those dashboards read-only.

### Optional

- `default_dashboard_list` (Block List, Max: 1) Configure home dashboard for selected user group. The selected preset dashboard will be loaded as default landing page for this environment. (see [below for nested schema](#nestedblock--default_dashboard_list))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--default_dashboard_list"></a>
### Nested Schema for `default_dashboard_list`

Required:

- `default_dashboard` (Block List, Min: 1) (see [below for nested schema](#nestedblock--default_dashboard_list--default_dashboard))

<a id="nestedblock--default_dashboard_list--default_dashboard"></a>
### Nested Schema for `default_dashboard_list.default_dashboard`

Required:

- `dashboard` (String) Preset dashboard to show as default landing page
- `user_group` (String) Show selected dashboard by default for this user group
 