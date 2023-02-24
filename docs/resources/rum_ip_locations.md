---
layout: ""
page_title: dynatrace_rum_ip_locations Resource - terraform-provider-dynatrace"
description: |-
  The resource `dynatrace_rum_ip_locations` covers configuration for IP address mapping rules for real user monitoring
---

# dynatrace_rum_ip_locations (Resource)

## Dynatrace Documentation

- Map internal IP addresses to locations - https://www.dynatrace.com/support/help/how-to-use-dynatrace/real-user-monitoring/setup-and-configuration/web-applications/additional-configuration/map-internal-ip-addresses-to-locations-web

- Settings API - https://www.dynatrace.com/support/help/dynatrace-api/environment-api/settings (schemaId: `builtin:rum.ip-mappings`)

## Export Example Usage

- `terraform-provider-dynatrace -export dynatrace_rum_ip_locations` downloads all existing IP address mapping configuration

The full documentation of the export feature is available [here](https://registry.terraform.io/providers/dynatrace-oss/dynatrace/latest/docs/guides/export-v2).

## Resource Example Usage

```terraform
resource "dynatrace_rum_ip_locations" "#name#" {
  country_code = "ZA"
  ip           = "192.168.3.1"
  ip_to        = "192.168.3.100"
  region_code  = "05"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `country_code` (String) The country code of the location. 

 Use the alpha-2 code of the [ISO 3166-2 standard](https://dt-url.net/iso3166-2), (for example, `AT` for Austria or `PL` for Poland).
- `ip` (String) Single IP or IP range start address

### Optional

- `city` (String) The city name of the location.
- `ip_to` (String) IP range end
- `latitude` (Number) Latitude
- `longitude` (Number) Longitude
- `region_code` (String) The region code of the location. 

 For the [USA](https://dt-url.net/iso3166us) or [Canada](https://dt-url.net/iso3166ca) use ISO 3166-2 state codes without `US-` or `CA-` prefix. 

 For the rest of the world use [FIPS 10-4 codes](https://dt-url.net/fipscodes) without country prefix.

### Read-Only

- `id` (String) The ID of this resource.
 