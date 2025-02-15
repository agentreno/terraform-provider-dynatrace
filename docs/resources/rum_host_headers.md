---
layout: ""
page_title: dynatrace_rum_host_headers Resource - terraform-provider-dynatrace"
description: |-
  The resource `dynatrace_rum_host_headers` covers configuration for identifying host names for real user monitoring
---

# dynatrace_rum_host_headers (Resource)

## Dynatrace Documentation

- Define applications for Real User Monitoring - https://www.dynatrace.com/support/help/platform-modules/digital-experience/web-applications/setup-and-configuration/initial-configuration/define-your-applications-via-the-my-web-application-placeholder

- Settings API - https://www.dynatrace.com/support/help/dynatrace-api/environment-api/settings (schemaId: `builtin:rum.host-headers`)

## Export Example Usage

- `terraform-provider-dynatrace -export dynatrace_rum_host_headers` downloads all existing host name HTTP request header configuration

The full documentation of the export feature is available [here](https://registry.terraform.io/providers/dynatrace-oss/dynatrace/latest/docs/guides/export-v2).

## Resource Example Usage

```terraform
resource "dynatrace_rum_host_headers" "#name#" {
  header_name = "HostHeaderExample3"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `header_name` (String) HTTP header format

### Read-Only

- `id` (String) The ID of this resource.
 