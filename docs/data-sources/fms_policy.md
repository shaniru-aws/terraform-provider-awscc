---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_fms_policy Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::FMS::Policy
---

# awscc_fms_policy (Data Source)

Data Source schema for AWS::FMS::Policy



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) Uniquely identifies the resource.

### Read-Only

- **arn** (String) A resource ARN.
- **delete_all_policy_resources** (Boolean)
- **exclude_map** (Attributes) An FMS includeMap or excludeMap. (see [below for nested schema](#nestedatt--exclude_map))
- **exclude_resource_tags** (Boolean)
- **include_map** (Attributes) An FMS includeMap or excludeMap. (see [below for nested schema](#nestedatt--include_map))
- **policy_name** (String)
- **remediation_enabled** (Boolean)
- **resource_tags** (Attributes List) (see [below for nested schema](#nestedatt--resource_tags))
- **resource_type** (String) An AWS resource type
- **resource_type_list** (List of String)
- **security_service_policy_data** (Attributes) (see [below for nested schema](#nestedatt--security_service_policy_data))
- **tags** (Attributes List) (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--exclude_map"></a>
### Nested Schema for `exclude_map`

Read-Only:

- **account** (List of String)
- **orgunit** (List of String)


<a id="nestedatt--include_map"></a>
### Nested Schema for `include_map`

Read-Only:

- **account** (List of String)
- **orgunit** (List of String)


<a id="nestedatt--resource_tags"></a>
### Nested Schema for `resource_tags`

Read-Only:

- **key** (String)
- **value** (String)


<a id="nestedatt--security_service_policy_data"></a>
### Nested Schema for `security_service_policy_data`

Read-Only:

- **managed_service_data** (String)
- **type** (String)


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- **key** (String)
- **value** (String)

