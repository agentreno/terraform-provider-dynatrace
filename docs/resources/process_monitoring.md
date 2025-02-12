---
layout: ""
page_title: "dynatrace_process_monitoring Resource - terraform-provider-dynatrace"
description: |-
  The resource `dynatrace_process_monitoring` covers configuration to monitor key processes on a host
---

# dynatrace_process_monitoring (Resource)

## Dynatrace Documentation

- Process Availability - https://www.dynatrace.com/support/help/how-to-use-dynatrace/process-groups/configuration/pg-monitoring

- Settings API - https://www.dynatrace.com/support/help/dynatrace-api/environment-api/settings (schemaId: `builtin:process.process-monitoring`)

## Export Example Usage

- `terraform-provider-dynatrace -export dynatrace_process_monitoring` downloads all existing key processes configuration

The full documentation of the export feature is available [here](https://registry.terraform.io/providers/dynatrace-oss/dynatrace/latest/docs/guides/export-v2).

## Resource Example Usage

```terraform
resource "dynatrace_process_monitoring" "example" {
  host_group_id   = "HOST_GROUP-0000000000000000"
  auto_monitoring = true
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `auto_monitoring` (Boolean) By disabling automatic deep monitoring the Dynatrace OneAgent will only deep monitor processes that are covered by a respective deep monitoring rule or where monitoring is enabled explicitly.
Disabling only works if all installed Agents have version 1.123 or higher. 

 With automatic monitoring enabled, you can create rules that define exceptions to automatic process detection and monitoring. With automatic monitoring disabled, you can define rules that identify specific processes that should be monitored. Rules are applied in the order listed in the custom and built-in process monitoring rules settings. This means that you can construct complex operations for fine-grain control over the processes that are monitored in your environment. For example, you might define an inclusion rule that’s followed by an exclusion rule covering the same process.
Once created, monitoring rules can be enabled/disabled at any time. The rules will only take effect after restart of the processes in question. Alternatively, you can disable automatic monitoring entirely and instead define "Include" rules for those processes you want to monitor.

### Optional

- `host_group_id` (String) The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.

### Read-Only

- `id` (String) The ID of this resource.
 