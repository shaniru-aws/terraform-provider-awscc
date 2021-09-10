// Code generated by generators/resource/main.go; DO NOT EDIT.

package nimblestudio

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("awscc_nimblestudio_studio_component", studioComponentResourceType)
}

// studioComponentResourceType returns the Terraform awscc_nimblestudio_studio_component resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::NimbleStudio::StudioComponent resource type.
func studioComponentResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"configuration": {
			// Property: Configuration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "ActiveDirectoryConfiguration": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "ComputerAttributes": {
			//           "items": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "Name": {
			//                 "type": "string"
			//               },
			//               "Value": {
			//                 "type": "string"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "type": "array"
			//         },
			//         "DirectoryId": {
			//           "type": "string"
			//         },
			//         "OrganizationalUnitDistinguishedName": {
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "ComputeFarmConfiguration": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "ActiveDirectoryUser": {
			//           "type": "string"
			//         },
			//         "Endpoint": {
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "LicenseServiceConfiguration": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "Endpoint": {
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "SharedFileSystemConfiguration": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "Endpoint": {
			//           "type": "string"
			//         },
			//         "FileSystemId": {
			//           "type": "string"
			//         },
			//         "LinuxMountPoint": {
			//           "type": "string"
			//         },
			//         "ShareName": {
			//           "type": "string"
			//         },
			//         "WindowsMountDrive": {
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"active_directory_configuration": {
						// Property: ActiveDirectoryConfiguration
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"computer_attributes": {
									// Property: ComputerAttributes
									Attributes: tfsdk.ListNestedAttributes(
										map[string]tfsdk.Attribute{
											"name": {
												// Property: Name
												Type:     types.StringType,
												Optional: true,
											},
											"value": {
												// Property: Value
												Type:     types.StringType,
												Optional: true,
											},
										},
										tfsdk.ListNestedAttributesOptions{},
									),
									Optional: true,
								},
								"directory_id": {
									// Property: DirectoryId
									Type:     types.StringType,
									Optional: true,
								},
								"organizational_unit_distinguished_name": {
									// Property: OrganizationalUnitDistinguishedName
									Type:     types.StringType,
									Optional: true,
								},
							},
						),
						Optional: true,
					},
					"compute_farm_configuration": {
						// Property: ComputeFarmConfiguration
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"active_directory_user": {
									// Property: ActiveDirectoryUser
									Type:     types.StringType,
									Optional: true,
								},
								"endpoint": {
									// Property: Endpoint
									Type:     types.StringType,
									Optional: true,
								},
							},
						),
						Optional: true,
					},
					"license_service_configuration": {
						// Property: LicenseServiceConfiguration
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"endpoint": {
									// Property: Endpoint
									Type:     types.StringType,
									Optional: true,
								},
							},
						),
						Optional: true,
					},
					"shared_file_system_configuration": {
						// Property: SharedFileSystemConfiguration
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"endpoint": {
									// Property: Endpoint
									Type:     types.StringType,
									Optional: true,
								},
								"file_system_id": {
									// Property: FileSystemId
									Type:     types.StringType,
									Optional: true,
								},
								"linux_mount_point": {
									// Property: LinuxMountPoint
									Type:     types.StringType,
									Optional: true,
								},
								"share_name": {
									// Property: ShareName
									Type:     types.StringType,
									Optional: true,
								},
								"windows_mount_drive": {
									// Property: WindowsMountDrive
									Type:     types.StringType,
									Optional: true,
								},
							},
						),
						Optional: true,
					},
				},
			),
			Optional: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
		},
		"ec_2_security_group_ids": {
			// Property: Ec2SecurityGroupIds
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Type:     types.ListType{ElemType: types.StringType},
			Optional: true,
		},
		"initialization_scripts": {
			// Property: InitializationScripts
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "LaunchProfileProtocolVersion": {
			//         "type": "string"
			//       },
			//       "Platform": {
			//         "type": "string"
			//       },
			//       "RunContext": {
			//         "type": "string"
			//       },
			//       "Script": {
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"launch_profile_protocol_version": {
						// Property: LaunchProfileProtocolVersion
						Type:     types.StringType,
						Optional: true,
					},
					"platform": {
						// Property: Platform
						Type:     types.StringType,
						Optional: true,
					},
					"run_context": {
						// Property: RunContext
						Type:     types.StringType,
						Optional: true,
					},
					"script": {
						// Property: Script
						Type:     types.StringType,
						Optional: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
		},
		"script_parameters": {
			// Property: ScriptParameters
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "type": "string"
			//       },
			//       "Value": {
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Optional: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Optional: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
		"studio_component_id": {
			// Property: StudioComponentId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"studio_id": {
			// Property: StudioId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"subtype": {
			// Property: Subtype
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "patternProperties": {
			//     "": {
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"type": {
			// Property: Type
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource schema for AWS::NimbleStudio::StudioComponent.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::NimbleStudio::StudioComponent").WithTerraformTypeName("awscc_nimblestudio_studio_component")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"active_directory_configuration":         "ActiveDirectoryConfiguration",
		"active_directory_user":                  "ActiveDirectoryUser",
		"compute_farm_configuration":             "ComputeFarmConfiguration",
		"computer_attributes":                    "ComputerAttributes",
		"configuration":                          "Configuration",
		"description":                            "Description",
		"directory_id":                           "DirectoryId",
		"ec_2_security_group_ids":                "Ec2SecurityGroupIds",
		"endpoint":                               "Endpoint",
		"file_system_id":                         "FileSystemId",
		"initialization_scripts":                 "InitializationScripts",
		"key":                                    "Key",
		"launch_profile_protocol_version":        "LaunchProfileProtocolVersion",
		"license_service_configuration":          "LicenseServiceConfiguration",
		"linux_mount_point":                      "LinuxMountPoint",
		"name":                                   "Name",
		"organizational_unit_distinguished_name": "OrganizationalUnitDistinguishedName",
		"platform":                               "Platform",
		"run_context":                            "RunContext",
		"script":                                 "Script",
		"script_parameters":                      "ScriptParameters",
		"share_name":                             "ShareName",
		"shared_file_system_configuration":       "SharedFileSystemConfiguration",
		"studio_component_id":                    "StudioComponentId",
		"studio_id":                              "StudioId",
		"subtype":                                "Subtype",
		"tags":                                   "Tags",
		"type":                                   "Type",
		"value":                                  "Value",
		"windows_mount_drive":                    "WindowsMountDrive",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_nimblestudio_studio_component", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
