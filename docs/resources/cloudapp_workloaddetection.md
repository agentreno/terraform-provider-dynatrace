---
layout: ""
page_title: "dynatrace_cloudapp_workloaddetection Resource - terraform-provider-dynatrace"
description: |-
  The resource `dynatrace_cloudapp_workloaddetection` merges processes of similar workloads into process groups, and consequently, services. Please note that fine-grained process detection rules will still be applied, while ignoring container or platform specific properties.
---

# dynatrace_cloudapp_workloaddetection (Resource)


## Dynatrace Documentation

- Cloud application and workload detection - https://www.dynatrace.com/support/help/platform-modules/infrastructure-monitoring/process-groups/configuration/cloud-app-and-workload-detection

- Settings API - https://www.dynatrace.com/support/help/dynatrace-api/environment-api/settings (schemaId: `builtin:process-group.cloud-application-workload-detection`)

## Export Example Usage

- `terraform-provider-dynatrace -export dynatrace_cloudapp_workloaddetection` downloads all existing muted requests configuration

The full documentation of the export feature is available [here](https://registry.terraform.io/providers/dynatrace-oss/dynatrace/latest/docs/guides/export-v2).

## Resource Example Usage

```terraform
# ID vu9U3hXa3q0AAAABADpidWlsdGluOnByb2Nlc3MtZ3JvdXAuY2xvdWQtYXBwbGljYXRpb24td29ya2xvYWQtZGV0ZWN0aW9uAAZ0ZW5hbnQABnRlbmFudAAkYjcwNmY4NWYtNWFkNC0zY2ZmLWJhYzMtZDg4YzFmNTkzMjgwvu9U3hXa3q0
resource "dynatrace_cloudapp_workloaddetection" "cloud_app_workload_detection" {
  cloud_foundry {
    enabled = false
  }
  docker {
    enabled = true
  }
  kubernetes {
    enabled = true
    filters {
      filter {
        enabled = false
        inclusion_toggles {
          inc_basepod   = false
          inc_container = true
          inc_namespace = true
          inc_product   = true
          inc_stage     = true
        }
        match_filter {
          match_operator = "EXISTS"
        }
      }
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cloud_foundry` (Block List, Min: 1, Max: 1) Enable this setting to get 
 * Processes of Cloud Foundry application instances merged into process groups by Cloud Foundry application. 
 *  Container resource metrics (Container group instance entities) and [related screens](https://www.dynatrace.com/support/help/shortlink/container-groups). (see [below for nested schema](#nestedblock--cloud_foundry))
- `docker` (Block List, Min: 1, Max: 1) Enable this setting for plain Docker environments to get 
 * Container resource metrics (Container group instance entities) and [related screens](https://www.dynatrace.com/support/help/shortlink/container-groups). (see [below for nested schema](#nestedblock--docker))
- `kubernetes` (Block List, Min: 1, Max: 1) Enable this setting to get 
 * Insights into your Kubernetes namespaces, workloads and pods (cloud application namespace, cloud application and cloud application instance and entities). 
 * Container resource metrics (container group instance entities) and [related screens](https://www.dynatrace.com/support/help/shortlink/container-groups). 
 * Similar workloads merged into process groups based on defined rules (see below). 
 * Version detection for services that run in Kubernetes workloads. (see [below for nested schema](#nestedblock--kubernetes))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--cloud_foundry"></a>
### Nested Schema for `cloud_foundry`

Required:

- `enabled` (Boolean) This setting is enabled (`true`) or disabled (`false`)


<a id="nestedblock--docker"></a>
### Nested Schema for `docker`

Required:

- `enabled` (Boolean) This setting is enabled (`true`) or disabled (`false`)


<a id="nestedblock--kubernetes"></a>
### Nested Schema for `kubernetes`

Required:

- `enabled` (Boolean) This setting is enabled (`true`) or disabled (`false`)
- `filters` (Block List, Min: 1, Max: 1) Define rules to merge similar Kubernetes workloads into process groups. 

 You can use workload properties like namespace name, base pod name or container name as well as the [environment variables DT_RELEASE_STAGE and DT_RELEASE_PRODUCT](https://dt-url.net/sb02v2a) for grouping processes of similar workloads. The first applicable rule will be applied. If no rule matches, “Namespace name” + “Base pod name” + “Container name” is used as fallback. (see [below for nested schema](#nestedblock--kubernetes--filters))

<a id="nestedblock--kubernetes--filters"></a>
### Nested Schema for `kubernetes.filters`

Required:

- `filter` (Block List, Min: 1) (see [below for nested schema](#nestedblock--kubernetes--filters--filter))

<a id="nestedblock--kubernetes--filters--filter"></a>
### Nested Schema for `kubernetes.filters.filter`

Required:

- `enabled` (Boolean) This setting is enabled (`true`) or disabled (`false`)
- `inclusion_toggles` (Block List, Min: 1, Max: 1) ID calculation based on (see [below for nested schema](#nestedblock--kubernetes--filters--filter--inclusion_toggles))
- `match_filter` (Block List, Min: 1, Max: 1) When namespace (see [below for nested schema](#nestedblock--kubernetes--filters--filter--match_filter))

<a id="nestedblock--kubernetes--filters--filter--inclusion_toggles"></a>
### Nested Schema for `kubernetes.filters.filter.inclusion_toggles`

Required:

- `inc_basepod` (Boolean) E.g. "cloud-credential-operator-" for "cloud-credential-operator-5ff6dbff57-gszgq"
- `inc_container` (Boolean) Container name
- `inc_namespace` (Boolean) Namespace name
- `inc_product` (Boolean) If Product is enabled and has no value, it defaults to Base pod name
- `inc_stage` (Boolean) Stage


<a id="nestedblock--kubernetes--filters--filter--match_filter"></a>
### Nested Schema for `kubernetes.filters.filter.match_filter`

Required:

- `match_operator` (String) Possible Values: `CONTAINS`, `ENDS`, `EQUALS`, `EXISTS`, `NOT_CONTAINS`, `NOT_ENDS`, `NOT_EQUALS`, `NOT_STARTS`, `STARTS`

Optional:

- `namespace` (String) Namespace name
 