---
layout: ""
page_title: dynatrace_k8s_pvc_anomalies Resource - terraform-provider-dynatrace"
description: |-
  The resource `dynatrace_k8s_pvc_anomalies` covers configuration for Kubernetes persistent volume claim anomalies
---

# dynatrace_k8s_pvc_anomalies (Resource)

## Dynatrace Documentation

- Alert on common Kubernetes/OpenShift issues - https://www.dynatrace.com/support/help/platform-modules/infrastructure-monitoring/container-platform-monitoring/kubernetes-monitoring/alert-on-kubernetes-issues

- Settings API - https://www.dynatrace.com/support/help/dynatrace-api/environment-api/settings (schemaId: `builtin:anomaly-detection.kubernetes.pvc`)

## Export Example Usage

- `terraform-provider-dynatrace -export dynatrace_k8s_pvc_anomalies` downloads all existing Kubernetes persistent volume claim anomaly configuration

The full documentation of the export feature is available [here](https://registry.terraform.io/providers/dynatrace-oss/dynatrace/latest/docs/guides/export-v2).

## Resource Example Usage

```terraform
resource "dynatrace_k8s_pvc_anomalies" "#name#" {
  scope = "environment"
  low_disk_space_critical {
    enabled = true
    configuration {
      observation_period_in_minutes = 5
      sample_period_in_minutes      = 3
      threshold                     = 100
    }
  }
  low_disk_space_critical_percentage {
    enabled = true
    configuration {
      observation_period_in_minutes = 5
      sample_period_in_minutes      = 3
      threshold                     = 10
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `low_disk_space_critical` (Block List, Min: 1, Max: 1) Alerts on low disk space in megabytes for a persistent volume claim. (see [below for nested schema](#nestedblock--low_disk_space_critical))
- `low_disk_space_critical_percentage` (Block List, Min: 1, Max: 1) Alerts on low disk space in % for a persistent volume claim. (see [below for nested schema](#nestedblock--low_disk_space_critical_percentage))

### Optional

- `scope` (String) The scope of this setting (CLOUD_APPLICATION_NAMESPACE, KUBERNETES_CLUSTER). Omit this property if you want to cover the whole environment.

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--low_disk_space_critical"></a>
### Nested Schema for `low_disk_space_critical`

Required:

- `enabled` (Boolean) This setting is enabled (`true`) or disabled (`false`)

Optional:

- `configuration` (Block List, Max: 1) Alert if (see [below for nested schema](#nestedblock--low_disk_space_critical--configuration))

<a id="nestedblock--low_disk_space_critical--configuration"></a>
### Nested Schema for `low_disk_space_critical.configuration`

Required:

- `observation_period_in_minutes` (Number) within the last
- `sample_period_in_minutes` (Number) for at least
- `threshold` (Number) the available disk space is below



<a id="nestedblock--low_disk_space_critical_percentage"></a>
### Nested Schema for `low_disk_space_critical_percentage`

Required:

- `enabled` (Boolean) This setting is enabled (`true`) or disabled (`false`)

Optional:

- `configuration` (Block List, Max: 1) Alert if (see [below for nested schema](#nestedblock--low_disk_space_critical_percentage--configuration))

<a id="nestedblock--low_disk_space_critical_percentage--configuration"></a>
### Nested Schema for `low_disk_space_critical_percentage.configuration`

Required:

- `observation_period_in_minutes` (Number) within the last
- `sample_period_in_minutes` (Number) for at least
- `threshold` (Number) the available disk space is below
 