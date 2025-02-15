---
layout: ""
page_title: "dynatrace_update_windows Data Source - terraform-provider-dynatrace"
description: |-
  The data source `dynatrace_update_windows` covers queries for maintenance windows for OneAgent updates 
---

# dynatrace_update_windows (Data Source)

The `dynatrace_update_windows` data source allows the OneAgent update maintenance window ID to be retrieved by its name.

- `name` (String) - The name of the OneAgent update maintenance window

## Example Usage

```terraform
data "dynatrace_update_windows" "Example" {
  name = "Terraform Example"
}

output "id" {
  value = data.dynatrace_update_windows.Example.id
}

```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)

### Read-Only

- `id` (String) The ID of this resource.