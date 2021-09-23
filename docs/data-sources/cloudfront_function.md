---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_cloudfront_function Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::CloudFront::Function
---

# awscc_cloudfront_function (Data Source)

Data Source schema for AWS::CloudFront::Function



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) Uniquely identifies the resource.

### Read-Only

- **auto_publish** (Boolean)
- **function_arn** (String)
- **function_code** (String)
- **function_config** (Attributes) (see [below for nested schema](#nestedatt--function_config))
- **function_metadata** (Attributes) (see [below for nested schema](#nestedatt--function_metadata))
- **name** (String)
- **stage** (String)

<a id="nestedatt--function_config"></a>
### Nested Schema for `function_config`

Read-Only:

- **comment** (String)
- **runtime** (String)


<a id="nestedatt--function_metadata"></a>
### Nested Schema for `function_metadata`

Read-Only:

- **function_arn** (String)

