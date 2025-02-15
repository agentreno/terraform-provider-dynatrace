---
layout: ""
page_title: "dynatrace_aws_credentials Data Source - terraform-provider-dynatrace"
description: |-
  The data source `dynatrace_aws_credentials` covers queries for AWS credentials
---

# dynatrace_aws_credentials (Data Source)

The `dynatrace_aws_credentials` data source allows the AWS credential ID to be retrieved by its label.

- `label` (String) - The label/name of the AWS credential

## Example Usage

```terraform
data "dynatrace_aws_credentials" "Example" {
  name = "Terraform Example"
}

output "id" {
  value = data.dynatrace_aws_credentials.Example.id
}

```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `label` (String)

### Read-Only

- `id` (String) The ID of this resource.