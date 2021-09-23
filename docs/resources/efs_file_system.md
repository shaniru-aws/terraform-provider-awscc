---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_efs_file_system Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::EFS::FileSystem
---

# awscc_efs_file_system (Resource)

Resource Type definition for AWS::EFS::FileSystem



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **availability_zone_name** (String)
- **backup_policy** (Attributes) (see [below for nested schema](#nestedatt--backup_policy))
- **bypass_policy_lockout_safety_check** (Boolean) Whether to bypass the FileSystemPolicy lockout safety check. The policy lockout safety check determines whether the policy in the request will prevent the principal making the request to be locked out from making future PutFileSystemPolicy requests on the file system. Set BypassPolicyLockoutSafetyCheck to True only when you intend to prevent the principal that is making the request from making a subsequent PutFileSystemPolicy request on the file system. Defaults to false
- **encrypted** (Boolean)
- **file_system_policy** (Map of String)
- **file_system_tags** (Attributes List) (see [below for nested schema](#nestedatt--file_system_tags))
- **kms_key_id** (String)
- **lifecycle_policies** (Attributes List) (see [below for nested schema](#nestedatt--lifecycle_policies))
- **performance_mode** (String)
- **provisioned_throughput_in_mibps** (Number)
- **throughput_mode** (String)

### Read-Only

- **arn** (String)
- **file_system_id** (String)
- **id** (String) Uniquely identifies the resource.

<a id="nestedatt--backup_policy"></a>
### Nested Schema for `backup_policy`

Optional:

- **status** (String)


<a id="nestedatt--file_system_tags"></a>
### Nested Schema for `file_system_tags`

Optional:

- **key** (String)
- **value** (String)


<a id="nestedatt--lifecycle_policies"></a>
### Nested Schema for `lifecycle_policies`

Optional:

- **transition_to_ia** (String)

