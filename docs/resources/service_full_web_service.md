---
layout: ""
page_title: dynatrace_service_full_web_service Resource - terraform-provider-dynatrace"
description: |-
  The resource `dynatrace_service_full_web_service` covers service detection rules for full web services
---

# dynatrace_service_full_web_service (Resource)

## Dynatrace Documentation

- Service detection rules - https://www.dynatrace.com/support/help/platform-modules/applications-and-microservices/services/service-detection-and-naming/customize-service-detection

- Settings API - https://www.dynatrace.com/support/help/dynatrace-api/environment-api/settings (schemaId: `builtin:service-detection.full-web-service`)

## Export Example Usage

- `terraform-provider-dynatrace -export dynatrace_service_full_web_service` downloads all existing service detection rules for full web services

The full documentation of the export feature is available [here](https://registry.terraform.io/providers/dynatrace-oss/dynatrace/latest/docs/guides/export-v2).

## Resource Example Usage

```terraform
resource "dynatrace_service_full_web_service" "#name#" {
  name             = "#name#"
  description      = "Created by Terraform"
  enabled          = true
  management_zones = [ "000000000000000000" ]
  conditions {
    condition {
      attribute              = "HostName"
      compare_operation_type = "StringEndsWith"
      ignore_case            = true
      text_values            = [ "Terraform" ]
    }
  }
  id_contributors {
    detect_as_web_request_service = true
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `enabled` (Boolean) This setting is enabled (`true`) or disabled (`false`)
- `id_contributors` (Block List, Min: 1, Max: 1) Contributors to the Service Identifier calculation. All of the Contributors are always applied. (see [below for nested schema](#nestedblock--id_contributors))
- `name` (String) Rule name

### Optional

- `conditions` (Block List, Max: 1) A list of conditions necessary for the rule to take effect. If multiple conditions are specified, they must **all** match a Request for the rule to apply. Conditions are evaluated against attributes, but do not modify them. (see [below for nested schema](#nestedblock--conditions))
- `description` (String) Description
- `management_zones` (Set of String) Define a management zone filter for this service detection rule.

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--id_contributors"></a>
### Nested Schema for `id_contributors`

Required:

- `detect_as_web_request_service` (Boolean) Detect the matching requests as full web services (false) or web request services (true).

Setting this field to true prevents detecting of matching requests as full web services. A web request service is created instead. If you need to further modify the resulting web request service, you need to create a separate [Full web request rule](builtin:service-detection.full-web-request).

Optional:

- `application_id` (Block List, Max: 1) Application identifier (see [below for nested schema](#nestedblock--id_contributors--application_id))
- `context_root` (Block List, Max: 1) The context root is the first segment of the request URL after the Server name. For example, in the `www.dynatrace.com/support/help/dynatrace-api/` URL the context root is `/support`. The context root value can be found on the **Service overview page** under **Properties and tags**. (see [below for nested schema](#nestedblock--id_contributors--context_root))
- `server_name` (Block List, Max: 1) Server name (see [below for nested schema](#nestedblock--id_contributors--server_name))
- `web_service_name` (Block List, Max: 1) Web service name (see [below for nested schema](#nestedblock--id_contributors--web_service_name))
- `web_service_namespace` (Block List, Max: 1) Web service namespace (see [below for nested schema](#nestedblock--id_contributors--web_service_namespace))

<a id="nestedblock--id_contributors--application_id"></a>
### Nested Schema for `id_contributors.application_id`

Required:

- `enable_id_contributor` (Boolean) Transform this value before letting it contribute to the Service Id

Optional:

- `service_id_contributor` (Block List, Max: 1) no documentation available (see [below for nested schema](#nestedblock--id_contributors--application_id--service_id_contributor))

<a id="nestedblock--id_contributors--application_id--service_id_contributor"></a>
### Nested Schema for `id_contributors.application_id.service_id_contributor`

Required:

- `contribution_type` (String) Possible Values: `OriginalValue`, `OverrideValue`, `TransformValue`

Optional:

- `transformations` (Block List, Max: 1) Choose how to transform a value before it contributes to the Service Id. Note that all of the Transformations are always applied. Transformations are applied in the order they are specified, and the output of the previous transformation is the input for the next one. The resulting value contributes to the Service Id and can be found on the **Service overview page** under **Properties and tags**. (see [below for nested schema](#nestedblock--id_contributors--application_id--service_id_contributor--transformations))
- `value_override` (Block List, Max: 1) The value to be used instead of the detected value. (see [below for nested schema](#nestedblock--id_contributors--application_id--service_id_contributor--value_override))

<a id="nestedblock--id_contributors--application_id--service_id_contributor--transformations"></a>
### Nested Schema for `id_contributors.application_id.service_id_contributor.transformations`

Required:

- `transformation` (Block List, Min: 1) (see [below for nested schema](#nestedblock--id_contributors--application_id--service_id_contributor--transformations--transformation))

<a id="nestedblock--id_contributors--application_id--service_id_contributor--transformations--transformation"></a>
### Nested Schema for `id_contributors.application_id.service_id_contributor.transformations.transformation`

Required:

- `transformation_type` (String) Possible Values: `AFTER`, `BEFORE`, `BETWEEN`, `REMOVE_CREDIT_CARDS`, `REMOVE_IBANS`, `REMOVE_IPS`, `REMOVE_NUMBERS`, `REPLACE_BETWEEN`, `SPLIT_SELECT`, `TAKE_SEGMENTS`

Optional:

- `include_hex_numbers` (Boolean) include hexadecimal numbers
- `min_digit_count` (Number) min digit count
- `prefix` (String) no documentation available
- `replacement_value` (String) replacement
- `segment_count` (Number) How many segments should be taken.
- `select_index` (Number) select index
- `split_delimiter` (String) split by
- `suffix` (String) no documentation available
- `take_from_end` (Boolean) take from end



<a id="nestedblock--id_contributors--application_id--service_id_contributor--value_override"></a>
### Nested Schema for `id_contributors.application_id.service_id_contributor.value_override`

Required:

- `value` (String) no documentation available




<a id="nestedblock--id_contributors--context_root"></a>
### Nested Schema for `id_contributors.context_root`

Required:

- `enable_id_contributor` (Boolean) Transform this value before letting it contribute to the Service Id

Optional:

- `service_id_contributor` (Block List, Max: 1) no documentation available (see [below for nested schema](#nestedblock--id_contributors--context_root--service_id_contributor))

<a id="nestedblock--id_contributors--context_root--service_id_contributor"></a>
### Nested Schema for `id_contributors.context_root.service_id_contributor`

Required:

- `contribution_type` (String) Possible Values: `OriginalValue`, `OverrideValue`, `TransformURL`, `TransformValue`

Optional:

- `segment_count` (Number) The number of segments of the URL to be kept. The URL is divided by slashes (/), the indexing starts with 1 at context root. For example, if you specify 2 for the `www.dynatrace.com/support/help/dynatrace-api/` URL, the value of `support/help` is used.
- `transformations` (Block List, Max: 1) Choose how to transform a value before it contributes to the Service Id. Note that all of the Transformations are always applied. Transformations are applied in the order they are specified, and the output of the previous transformation is the input for the next one. The resulting value contributes to the Service Id and can be found on the **Service overview page** under **Properties and tags**. (see [below for nested schema](#nestedblock--id_contributors--context_root--service_id_contributor--transformations))
- `value_override` (Block List, Max: 1) The value to be used instead of the detected value. (see [below for nested schema](#nestedblock--id_contributors--context_root--service_id_contributor--value_override))

<a id="nestedblock--id_contributors--context_root--service_id_contributor--transformations"></a>
### Nested Schema for `id_contributors.context_root.service_id_contributor.transformations`

Required:

- `transformation` (Block List, Min: 1) (see [below for nested schema](#nestedblock--id_contributors--context_root--service_id_contributor--transformations--transformation))

<a id="nestedblock--id_contributors--context_root--service_id_contributor--transformations--transformation"></a>
### Nested Schema for `id_contributors.context_root.service_id_contributor.transformations.transformation`

Required:

- `transformation_type` (String) Possible Values: `BEFORE`, `REMOVE_CREDIT_CARDS`, `REMOVE_IBANS`, `REMOVE_IPS`, `REMOVE_NUMBERS`, `REPLACE_BETWEEN`

Optional:

- `include_hex_numbers` (Boolean) include hexadecimal numbers
- `min_digit_count` (Number) min digit count
- `prefix` (String) no documentation available
- `replacement_value` (String) replacement
- `suffix` (String) no documentation available



<a id="nestedblock--id_contributors--context_root--service_id_contributor--value_override"></a>
### Nested Schema for `id_contributors.context_root.service_id_contributor.value_override`

Required:

- `value` (String) no documentation available




<a id="nestedblock--id_contributors--server_name"></a>
### Nested Schema for `id_contributors.server_name`

Required:

- `enable_id_contributor` (Boolean) Transform this value before letting it contribute to the Service Id

Optional:

- `service_id_contributor` (Block List, Max: 1) no documentation available (see [below for nested schema](#nestedblock--id_contributors--server_name--service_id_contributor))

<a id="nestedblock--id_contributors--server_name--service_id_contributor"></a>
### Nested Schema for `id_contributors.server_name.service_id_contributor`

Required:

- `contribution_type` (String) Possible Values: `OriginalValue`, `OverrideValue`, `TransformValue`

Optional:

- `transformations` (Block List, Max: 1) Choose how to transform a value before it contributes to the Service Id. Note that all of the Transformations are always applied. Transformations are applied in the order they are specified, and the output of the previous transformation is the input for the next one. The resulting value contributes to the Service Id and can be found on the **Service overview page** under **Properties and tags**. (see [below for nested schema](#nestedblock--id_contributors--server_name--service_id_contributor--transformations))
- `value_override` (Block List, Max: 1) The value to be used instead of the detected value. (see [below for nested schema](#nestedblock--id_contributors--server_name--service_id_contributor--value_override))

<a id="nestedblock--id_contributors--server_name--service_id_contributor--transformations"></a>
### Nested Schema for `id_contributors.server_name.service_id_contributor.transformations`

Required:

- `transformation` (Block List, Min: 1) (see [below for nested schema](#nestedblock--id_contributors--server_name--service_id_contributor--transformations--transformation))

<a id="nestedblock--id_contributors--server_name--service_id_contributor--transformations--transformation"></a>
### Nested Schema for `id_contributors.server_name.service_id_contributor.transformations.transformation`

Required:

- `transformation_type` (String) Possible Values: `AFTER`, `BEFORE`, `BETWEEN`, `REMOVE_CREDIT_CARDS`, `REMOVE_IBANS`, `REMOVE_IPS`, `REMOVE_NUMBERS`, `REPLACE_BETWEEN`, `SPLIT_SELECT`, `TAKE_SEGMENTS`

Optional:

- `include_hex_numbers` (Boolean) include hexadecimal numbers
- `min_digit_count` (Number) min digit count
- `prefix` (String) no documentation available
- `replacement_value` (String) replacement
- `segment_count` (Number) How many segments should be taken.
- `select_index` (Number) select index
- `split_delimiter` (String) split by
- `suffix` (String) no documentation available
- `take_from_end` (Boolean) take from end



<a id="nestedblock--id_contributors--server_name--service_id_contributor--value_override"></a>
### Nested Schema for `id_contributors.server_name.service_id_contributor.value_override`

Required:

- `value` (String) no documentation available




<a id="nestedblock--id_contributors--web_service_name"></a>
### Nested Schema for `id_contributors.web_service_name`

Required:

- `enable_id_contributor` (Boolean) Transform this value before letting it contribute to the Service Id

Optional:

- `service_id_contributor` (Block List, Max: 1) no documentation available (see [below for nested schema](#nestedblock--id_contributors--web_service_name--service_id_contributor))

<a id="nestedblock--id_contributors--web_service_name--service_id_contributor"></a>
### Nested Schema for `id_contributors.web_service_name.service_id_contributor`

Required:

- `contribution_type` (String) Possible Values: `OriginalValue`, `OverrideValue`, `TransformValue`

Optional:

- `transformations` (Block List, Max: 1) Choose how to transform a value before it contributes to the Service Id. Note that all of the Transformations are always applied. Transformations are applied in the order they are specified, and the output of the previous transformation is the input for the next one. The resulting value contributes to the Service Id and can be found on the **Service overview page** under **Properties and tags**. (see [below for nested schema](#nestedblock--id_contributors--web_service_name--service_id_contributor--transformations))
- `value_override` (Block List, Max: 1) The value to be used instead of the detected value. (see [below for nested schema](#nestedblock--id_contributors--web_service_name--service_id_contributor--value_override))

<a id="nestedblock--id_contributors--web_service_name--service_id_contributor--transformations"></a>
### Nested Schema for `id_contributors.web_service_name.service_id_contributor.transformations`

Required:

- `transformation` (Block List, Min: 1) (see [below for nested schema](#nestedblock--id_contributors--web_service_name--service_id_contributor--transformations--transformation))

<a id="nestedblock--id_contributors--web_service_name--service_id_contributor--transformations--transformation"></a>
### Nested Schema for `id_contributors.web_service_name.service_id_contributor.transformations.transformation`

Required:

- `transformation_type` (String) Possible Values: `AFTER`, `BEFORE`, `BETWEEN`, `REMOVE_CREDIT_CARDS`, `REMOVE_IBANS`, `REMOVE_IPS`, `REMOVE_NUMBERS`, `REPLACE_BETWEEN`, `SPLIT_SELECT`, `TAKE_SEGMENTS`

Optional:

- `include_hex_numbers` (Boolean) include hexadecimal numbers
- `min_digit_count` (Number) min digit count
- `prefix` (String) no documentation available
- `replacement_value` (String) replacement
- `segment_count` (Number) How many segments should be taken.
- `select_index` (Number) select index
- `split_delimiter` (String) split by
- `suffix` (String) no documentation available
- `take_from_end` (Boolean) take from end



<a id="nestedblock--id_contributors--web_service_name--service_id_contributor--value_override"></a>
### Nested Schema for `id_contributors.web_service_name.service_id_contributor.value_override`

Required:

- `value` (String) no documentation available




<a id="nestedblock--id_contributors--web_service_namespace"></a>
### Nested Schema for `id_contributors.web_service_namespace`

Required:

- `enable_id_contributor` (Boolean) Transform this value before letting it contribute to the Service Id

Optional:

- `service_id_contributor` (Block List, Max: 1) no documentation available (see [below for nested schema](#nestedblock--id_contributors--web_service_namespace--service_id_contributor))

<a id="nestedblock--id_contributors--web_service_namespace--service_id_contributor"></a>
### Nested Schema for `id_contributors.web_service_namespace.service_id_contributor`

Required:

- `contribution_type` (String) Possible Values: `OriginalValue`, `OverrideValue`, `TransformValue`

Optional:

- `transformations` (Block List, Max: 1) Choose how to transform a value before it contributes to the Service Id. Note that all of the Transformations are always applied. Transformations are applied in the order they are specified, and the output of the previous transformation is the input for the next one. The resulting value contributes to the Service Id and can be found on the **Service overview page** under **Properties and tags**. (see [below for nested schema](#nestedblock--id_contributors--web_service_namespace--service_id_contributor--transformations))
- `value_override` (Block List, Max: 1) The value to be used instead of the detected value. (see [below for nested schema](#nestedblock--id_contributors--web_service_namespace--service_id_contributor--value_override))

<a id="nestedblock--id_contributors--web_service_namespace--service_id_contributor--transformations"></a>
### Nested Schema for `id_contributors.web_service_namespace.service_id_contributor.transformations`

Required:

- `transformation` (Block List, Min: 1) (see [below for nested schema](#nestedblock--id_contributors--web_service_namespace--service_id_contributor--transformations--transformation))

<a id="nestedblock--id_contributors--web_service_namespace--service_id_contributor--transformations--transformation"></a>
### Nested Schema for `id_contributors.web_service_namespace.service_id_contributor.transformations.transformation`

Required:

- `transformation_type` (String) Possible Values: `AFTER`, `BEFORE`, `BETWEEN`, `REMOVE_CREDIT_CARDS`, `REMOVE_IBANS`, `REMOVE_IPS`, `REMOVE_NUMBERS`, `REPLACE_BETWEEN`, `SPLIT_SELECT`, `TAKE_SEGMENTS`

Optional:

- `include_hex_numbers` (Boolean) include hexadecimal numbers
- `min_digit_count` (Number) min digit count
- `prefix` (String) no documentation available
- `replacement_value` (String) replacement
- `segment_count` (Number) How many segments should be taken.
- `select_index` (Number) select index
- `split_delimiter` (String) split by
- `suffix` (String) no documentation available
- `take_from_end` (Boolean) take from end



<a id="nestedblock--id_contributors--web_service_namespace--service_id_contributor--value_override"></a>
### Nested Schema for `id_contributors.web_service_namespace.service_id_contributor.value_override`

Required:

- `value` (String) no documentation available





<a id="nestedblock--conditions"></a>
### Nested Schema for `conditions`

Required:

- `condition` (Block List, Min: 1) (see [below for nested schema](#nestedblock--conditions--condition))

<a id="nestedblock--conditions--condition"></a>
### Nested Schema for `conditions.condition`

Required:

- `attribute` (String) Take the value of this attribute
- `compare_operation_type` (String) Apply this operation

Optional:

- `framework` (Set of String) Technology
- `ignore_case` (Boolean) Ignore case sensitivity for texts.
- `int_value` (Number) Value
- `int_values` (Set of Number) Values
- `ip_range_from` (String) From
- `ip_range_to` (String) To
- `tag_values` (Set of String) If multiple values are specified, at least one of them must match for the condition to match
- `text_values` (Set of String) If multiple values are specified, at least one of them must match for the condition to match
 