// Code generated by generators/resource/main.go; DO NOT EDIT.

package rds

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"

	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_rds_db_proxy", dBProxyResourceType)
}

// dBProxyResourceType returns the Terraform awscc_rds_db_proxy resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::RDS::DBProxy resource type.
func dBProxyResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"auth": {
			// Property: Auth
			// CloudFormation resource type schema:
			// {
			//   "description": "The authorization mechanism that the proxy uses.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "AuthScheme": {
			//         "description": "The type of authentication that the proxy uses for connections from the proxy to the underlying database. ",
			//         "enum": [
			//           "SECRETS"
			//         ],
			//         "type": "string"
			//       },
			//       "Description": {
			//         "description": "A user-specified description about the authentication used by a proxy to log in as a specific database user. ",
			//         "type": "string"
			//       },
			//       "IAMAuth": {
			//         "description": "Whether to require or disallow AWS Identity and Access Management (IAM) authentication for connections to the proxy. ",
			//         "enum": [
			//           "DISABLED",
			//           "REQUIRED"
			//         ],
			//         "type": "string"
			//       },
			//       "SecretArn": {
			//         "description": "The Amazon Resource Name (ARN) representing the secret that the proxy uses to authenticate to the RDS DB instance or Aurora DB cluster. These secrets are stored within Amazon Secrets Manager. ",
			//         "type": "string"
			//       },
			//       "UserName": {
			//         "description": "The name of the database user to which the proxy connects.",
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "minItems": 1,
			//   "type": "array"
			// }
			Description: "The authorization mechanism that the proxy uses.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"auth_scheme": {
						// Property: AuthScheme
						Description: "The type of authentication that the proxy uses for connections from the proxy to the underlying database. ",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"SECRETS",
							}),
						},
					},
					"description": {
						// Property: Description
						Description: "A user-specified description about the authentication used by a proxy to log in as a specific database user. ",
						Type:        types.StringType,
						Optional:    true,
					},
					"iam_auth": {
						// Property: IAMAuth
						Description: "Whether to require or disallow AWS Identity and Access Management (IAM) authentication for connections to the proxy. ",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"DISABLED",
								"REQUIRED",
							}),
						},
					},
					"secret_arn": {
						// Property: SecretArn
						Description: "The Amazon Resource Name (ARN) representing the secret that the proxy uses to authenticate to the RDS DB instance or Aurora DB cluster. These secrets are stored within Amazon Secrets Manager. ",
						Type:        types.StringType,
						Optional:    true,
					},
					"user_name": {
						// Property: UserName
						Description: "The name of the database user to which the proxy connects.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
				tfsdk.ListNestedAttributesOptions{
					MinItems: 1,
				},
			),
			Required: true,
		},
		"db_proxy_arn": {
			// Property: DBProxyArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) for the proxy.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) for the proxy.",
			Type:        types.StringType,
			Computed:    true,
		},
		"db_proxy_name": {
			// Property: DBProxyName
			// CloudFormation resource type schema:
			// {
			//   "description": "The identifier for the proxy. This name must be unique for all proxies owned by your AWS account in the specified AWS Region.",
			//   "maxLength": 64,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The identifier for the proxy. This name must be unique for all proxies owned by your AWS account in the specified AWS Region.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 64),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"debug_logging": {
			// Property: DebugLogging
			// CloudFormation resource type schema:
			// {
			//   "description": "Whether the proxy includes detailed information about SQL statements in its logs.",
			//   "type": "boolean"
			// }
			Description: "Whether the proxy includes detailed information about SQL statements in its logs.",
			Type:        types.BoolType,
			Optional:    true,
		},
		"endpoint": {
			// Property: Endpoint
			// CloudFormation resource type schema:
			// {
			//   "description": "The endpoint that you can use to connect to the proxy. You include the endpoint value in the connection string for a database client application.",
			//   "type": "string"
			// }
			Description: "The endpoint that you can use to connect to the proxy. You include the endpoint value in the connection string for a database client application.",
			Type:        types.StringType,
			Computed:    true,
		},
		"engine_family": {
			// Property: EngineFamily
			// CloudFormation resource type schema:
			// {
			//   "description": "The kinds of databases that the proxy can connect to.",
			//   "enum": [
			//     "MYSQL",
			//     "POSTGRESQL"
			//   ],
			//   "type": "string"
			// }
			Description: "The kinds of databases that the proxy can connect to.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"MYSQL",
					"POSTGRESQL",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"idle_client_timeout": {
			// Property: IdleClientTimeout
			// CloudFormation resource type schema:
			// {
			//   "description": "The number of seconds that a connection to the proxy can be inactive before the proxy disconnects it.",
			//   "type": "integer"
			// }
			Description: "The number of seconds that a connection to the proxy can be inactive before the proxy disconnects it.",
			Type:        types.NumberType,
			Optional:    true,
		},
		"require_tls": {
			// Property: RequireTLS
			// CloudFormation resource type schema:
			// {
			//   "description": "A Boolean parameter that specifies whether Transport Layer Security (TLS) encryption is required for connections to the proxy.",
			//   "type": "boolean"
			// }
			Description: "A Boolean parameter that specifies whether Transport Layer Security (TLS) encryption is required for connections to the proxy.",
			Type:        types.BoolType,
			Optional:    true,
		},
		"role_arn": {
			// Property: RoleArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the IAM role that the proxy uses to access secrets in AWS Secrets Manager.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the IAM role that the proxy uses to access secrets in AWS Secrets Manager.",
			Type:        types.StringType,
			Required:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An optional set of key-value pairs to associate arbitrary data of your choosing with the proxy.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 128,
			//         "pattern": "",
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "An optional set of key-value pairs to associate arbitrary data of your choosing with the proxy.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 128),
						},
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 128),
						},
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
		"vpc_id": {
			// Property: VpcId
			// CloudFormation resource type schema:
			// {
			//   "description": "VPC ID to associate with the new DB proxy.",
			//   "type": "string"
			// }
			Description: "VPC ID to associate with the new DB proxy.",
			Type:        types.StringType,
			Computed:    true,
		},
		"vpc_security_group_ids": {
			// Property: VpcSecurityGroupIds
			// CloudFormation resource type schema:
			// {
			//   "description": "VPC security group IDs to associate with the new proxy.",
			//   "insertionOrder": false,
			//   "items": {
			//     "type": "string"
			//   },
			//   "minItems": 1,
			//   "type": "array"
			// }
			Description: "VPC security group IDs to associate with the new proxy.",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtLeast(1),
			},
		},
		"vpc_subnet_ids": {
			// Property: VpcSubnetIds
			// CloudFormation resource type schema:
			// {
			//   "description": "VPC subnet IDs to associate with the new proxy.",
			//   "insertionOrder": false,
			//   "items": {
			//     "type": "string"
			//   },
			//   "minItems": 2,
			//   "type": "array"
			// }
			Description: "VPC subnet IDs to associate with the new proxy.",
			Type:        types.ListType{ElemType: types.StringType},
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtLeast(2),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource schema for AWS::RDS::DBProxy",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::RDS::DBProxy").WithTerraformTypeName("awscc_rds_db_proxy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"auth":                   "Auth",
		"auth_scheme":            "AuthScheme",
		"db_proxy_arn":           "DBProxyArn",
		"db_proxy_name":          "DBProxyName",
		"debug_logging":          "DebugLogging",
		"description":            "Description",
		"endpoint":               "Endpoint",
		"engine_family":          "EngineFamily",
		"iam_auth":               "IAMAuth",
		"idle_client_timeout":    "IdleClientTimeout",
		"key":                    "Key",
		"require_tls":            "RequireTLS",
		"role_arn":               "RoleArn",
		"secret_arn":             "SecretArn",
		"tags":                   "Tags",
		"user_name":              "UserName",
		"value":                  "Value",
		"vpc_id":                 "VpcId",
		"vpc_security_group_ids": "VpcSecurityGroupIds",
		"vpc_subnet_ids":         "VpcSubnetIds",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_rds_db_proxy", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
