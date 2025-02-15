---
layout: ""
page_title: dynatrace_disk_specific_anomalies_v2 Resource - terraform-provider-dynatrace"
description: |-
  The resource `dynatrace_disk_specific_anomalies_v2` covers configuration for disk specific anomaly detection
---

# dynatrace_disk_specific_anomalies_v2 (Resource)

## Dynatrace Documentation

- Adjust the sensitivity of anomaly detection for infrastructure - https://www.dynatrace.com/support/help/how-to-use-dynatrace/problem-detection-and-analysis/problem-detection/adjust-sensitivity-anomaly-detection/adjust-sensitivity-infastructure

- Settings API - https://www.dynatrace.com/support/help/dynatrace-api/environment-api/settings (schemaId: `builtin:anomaly-detection.infrastructure-disks.per-disk-override`)

## Export Example Usage

- `terraform-provider-dynatrace -export dynatrace_disk_specific_anomalies_v2` downloads all existing disk specific anomaly detection configuration

The full documentation of the export feature is available [here](https://registry.terraform.io/providers/dynatrace-oss/dynatrace/latest/docs/guides/export-v2).

## Resource Example Usage

```terraform
resource "dynatrace_disk_specific_anomalies_v2" "#name#" {
  disk_id                                  = "DISK-1234567890000000"
  override_disk_low_space_detection        = true
  override_low_inodes_detection            = true
  override_slow_writes_and_reads_detection = true
  disk_low_inodes_detection {
    enabled        = true
    detection_mode = "custom"
    custom_thresholds {
      free_inodes_percentage = 1
    }
  }
  disk_low_space_detection {
    enabled        = true
    detection_mode = "custom"
    custom_thresholds {
      free_space_percentage = 1
    }
  }
  disk_slow_writes_and_reads_detection {
    enabled        = true
    detection_mode = "custom"
    custom_thresholds {
      write_and_read_time = 300
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `disk_id` (String) The id for the disk anomaly detection
- `override_disk_low_space_detection` (Boolean) Override low disk space detection settings
- `override_low_inodes_detection` (Boolean) Override low inodes detection settings
- `override_slow_writes_and_reads_detection` (Boolean) Override slow writes and reads detection settings

### Optional

- `disk_low_inodes_detection` (Block List, Max: 1) no documentation available (see [below for nested schema](#nestedblock--disk_low_inodes_detection))
- `disk_low_space_detection` (Block List, Max: 1) no documentation available (see [below for nested schema](#nestedblock--disk_low_space_detection))
- `disk_slow_writes_and_reads_detection` (Block List, Max: 1) no documentation available (see [below for nested schema](#nestedblock--disk_slow_writes_and_reads_detection))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--disk_low_inodes_detection"></a>
### Nested Schema for `disk_low_inodes_detection`

Required:

- `enabled` (Boolean) Detect low inodes number available

Optional:

- `custom_thresholds` (Block List, Max: 1) no documentation available (see [below for nested schema](#nestedblock--disk_low_inodes_detection--custom_thresholds))
- `detection_mode` (String) Detection mode for low inodes number available

<a id="nestedblock--disk_low_inodes_detection--custom_thresholds"></a>
### Nested Schema for `disk_low_inodes_detection.custom_thresholds`

Required:

- `free_inodes_percentage` (Number) Alert if the percentage of available inodes is lower than this threshold in 3 out of 5 samples



<a id="nestedblock--disk_low_space_detection"></a>
### Nested Schema for `disk_low_space_detection`

Required:

- `enabled` (Boolean) Detect low disk space

Optional:

- `custom_thresholds` (Block List, Max: 1) no documentation available (see [below for nested schema](#nestedblock--disk_low_space_detection--custom_thresholds))
- `detection_mode` (String) Detection mode for low disk space

<a id="nestedblock--disk_low_space_detection--custom_thresholds"></a>
### Nested Schema for `disk_low_space_detection.custom_thresholds`

Required:

- `free_space_percentage` (Number) Alert if free disk space is lower than this percentage in 3 out of 5 samples



<a id="nestedblock--disk_slow_writes_and_reads_detection"></a>
### Nested Schema for `disk_slow_writes_and_reads_detection`

Required:

- `enabled` (Boolean) Detect slow-running disks

Optional:

- `custom_thresholds` (Block List, Max: 1) no documentation available (see [below for nested schema](#nestedblock--disk_slow_writes_and_reads_detection--custom_thresholds))
- `detection_mode` (String) Detection mode for slow running disks

<a id="nestedblock--disk_slow_writes_and_reads_detection--custom_thresholds"></a>
### Nested Schema for `disk_slow_writes_and_reads_detection.custom_thresholds`

Required:

- `write_and_read_time` (Number) Alert if disk read time or write time is higher than this threshold in 3 out of 5 samples
 