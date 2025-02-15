---
layout: ""
page_title: dynatrace_dashboards_presets Resource - terraform-provider-dynatrace"
description: |-
  The resource `dynatrace_dashboards_presets` covers configuration for dashboard preset settings
---

# dynatrace_dashboards_presets (Resource)

## Dynatrace Documentation

- Preset Dynatrace dashboards - https://www.dynatrace.com/support/help/observe-and-explore/dashboards/dashboards/dashboards-preset

- Settings API - https://www.dynatrace.com/support/help/dynatrace-api/environment-api/settings (schemaId: `builtin:dashboards.presets`)

## Export Example Usage

- `terraform-provider-dynatrace -export dynatrace_dashboards_presets` downloads all existing dashboard preset settings

The full documentation of the export feature is available [here](https://registry.terraform.io/providers/dynatrace-oss/dynatrace/latest/docs/guides/export-v2).

## Resource Example Usage

```terraform
resource "dynatrace_dashboards_presets" "#name#" {
  enable_dashboard_presets = true
  dashboard_presets_list {
    dashboard_presets {
      dashboard_preset = "00000000-0000-0000-0000-000000000000"
      user_group = "00000000-0000-0000-0000-000000000000"
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `enable_dashboard_presets` (Boolean) Dashboard presets are visible to all users by default. For a pristine environment you may disable them entirely or opt to manually limit visibility to selected user groups.

### Optional

- `dashboard_presets_list` (Block List, Max: 1) Show selected preset to respective user group only. (see [below for nested schema](#nestedblock--dashboard_presets_list))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--dashboard_presets_list"></a>
### Nested Schema for `dashboard_presets_list`

Required:

- `dashboard_presets` (Block List, Min: 1) (see [below for nested schema](#nestedblock--dashboard_presets_list--dashboard_presets))

<a id="nestedblock--dashboard_presets_list--dashboard_presets"></a>
### Nested Schema for `dashboard_presets_list.dashboard_presets`

Required:

- `dashboard_preset` (String) Dashboard preset to limit visibility for
- `user_group` (String) User group to show selected dashboard preset to
 