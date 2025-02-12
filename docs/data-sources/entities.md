---
layout: ""
page_title: "dynatrace_entities Data Source - terraform-provider-dynatrace"
description: |-
  The data source `dynatrace_entities` covers queries for the list of entities based off of type
---

# dynatrace_entities (Data Source)

The entities data source allows all entities to be retrieved by its type.

- `type` (String) Type of the entity, e.g. SERVICE. All available entity types can be retrieved with [/api/v2/entityTypes](https://www.dynatrace.com/support/help/dynatrace-api/environment-api/entity-v2/get-all-entity-types).

## Example Usage

```terraform
data "dynatrace_entities" "Test" {
  type = "SERVICE"
}

output "Service_List" {
  value = data.dynatrace_entities.Test.entities
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `type` (String)

### Optional

- `entities` (Block List) (see [below for nested schema](#nestedblock--entities))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--entities"></a>
### Nested Schema for `entities`

Optional:

- `display_name` (String) The name of the entity, displayed in the UI.
- `entity_id` (String) The ID of the entity.
- `tags` (Block List) A set of tags assigned to the entity. (see [below for nested schema](#nestedblock--entities--tags))
- `type` (String) The type of the entity.

Read-Only:

- `properties` (Map of String) Properties defining the entity.

<a id="nestedblock--entities--tags"></a>
### Nested Schema for `entities.tags`

Optional:

- `tag` (Block List) A tag assigned to the entity (see [below for nested schema](#nestedblock--entities--tags--tag))

<a id="nestedblock--entities--tags--tag"></a>
### Nested Schema for `entities.tags.tag`

Required:

- `context` (String) The origin of the tag, such as AWS or Cloud Foundry. Custom tags use the `CONTEXTLESS` value
- `key` (String) The key of the tag. Custom tags have the tag value here

Optional:

- `string_representation` (String) The string representation of the tag
- `value` (String) The value of the tag. Not applicable to custom tags