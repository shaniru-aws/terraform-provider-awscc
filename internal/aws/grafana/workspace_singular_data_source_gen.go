// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package grafana

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_grafana_workspace", workspaceDataSource)
}

// workspaceDataSource returns the Terraform awscc_grafana_workspace data source.
// This Terraform data source corresponds to the CloudFormation AWS::Grafana::Workspace resource.
func workspaceDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"account_access_type": {
			// Property: AccountAccessType
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "These enums represent valid account access types. Specifically these enums determine whether the workspace can access AWS resources in the AWS account only, or whether it can also access resources in other accounts in the same organization. If the value CURRENT_ACCOUNT is used, a workspace role ARN must be provided. If the value is ORGANIZATION, a list of organizational units must be provided.",
			//	  "enum": [
			//	    "CURRENT_ACCOUNT",
			//	    "ORGANIZATION"
			//	  ],
			//	  "type": "string"
			//	}
			Description: "These enums represent valid account access types. Specifically these enums determine whether the workspace can access AWS resources in the AWS account only, or whether it can also access resources in other accounts in the same organization. If the value CURRENT_ACCOUNT is used, a workspace role ARN must be provided. If the value is ORGANIZATION, a list of organizational units must be provided.",
			Type:        types.StringType,
			Computed:    true,
		},
		"authentication_providers": {
			// Property: AuthenticationProviders
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "List of authentication providers to enable.",
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "description": "Valid workspace authentication providers.",
			//	    "enum": [
			//	      "AWS_SSO",
			//	      "SAML"
			//	    ],
			//	    "type": "string"
			//	  },
			//	  "minItems": 1,
			//	  "type": "array",
			//	  "uniqueItems": true
			//	}
			Description: "List of authentication providers to enable.",
			Type:        types.SetType{ElemType: types.StringType},
			Computed:    true,
		},
		"client_token": {
			// Property: ClientToken
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "A unique, case-sensitive, user-provided identifier to ensure the idempotency of the request.",
			//	  "pattern": "^[!-~]{1,64}$",
			//	  "type": "string"
			//	}
			Description: "A unique, case-sensitive, user-provided identifier to ensure the idempotency of the request.",
			Type:        types.StringType,
			Computed:    true,
		},
		"creation_timestamp": {
			// Property: CreationTimestamp
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Timestamp when the workspace was created.",
			//	  "format": "date-time",
			//	  "type": "string"
			//	}
			Description: "Timestamp when the workspace was created.",
			Type:        types.StringType,
			Computed:    true,
		},
		"data_sources": {
			// Property: DataSources
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "List of data sources on the service managed IAM role.",
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "description": "These enums represent valid AWS data sources that can be queried via the Grafana workspace. These data sources are primarily used to help customers visualize which data sources have been added to a service managed workspace IAM role.",
			//	    "enum": [
			//	      "AMAZON_OPENSEARCH_SERVICE",
			//	      "CLOUDWATCH",
			//	      "PROMETHEUS",
			//	      "XRAY",
			//	      "TIMESTREAM",
			//	      "SITEWISE",
			//	      "ATHENA",
			//	      "REDSHIFT"
			//	    ],
			//	    "type": "string"
			//	  },
			//	  "type": "array"
			//	}
			Description: "List of data sources on the service managed IAM role.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Description of a workspace.",
			//	  "maxLength": 2048,
			//	  "minLength": 0,
			//	  "type": "string"
			//	}
			Description: "Description of a workspace.",
			Type:        types.StringType,
			Computed:    true,
		},
		"endpoint": {
			// Property: Endpoint
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Endpoint for the Grafana workspace.",
			//	  "maxLength": 2048,
			//	  "minLength": 1,
			//	  "type": "string"
			//	}
			Description: "Endpoint for the Grafana workspace.",
			Type:        types.StringType,
			Computed:    true,
		},
		"grafana_version": {
			// Property: GrafanaVersion
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Version of Grafana the workspace is currently using.",
			//	  "maxLength": 255,
			//	  "minLength": 1,
			//	  "type": "string"
			//	}
			Description: "Version of Grafana the workspace is currently using.",
			Type:        types.StringType,
			Computed:    true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The id that uniquely identifies a Grafana workspace.",
			//	  "pattern": "^g-[0-9a-f]{10}$",
			//	  "type": "string"
			//	}
			Description: "The id that uniquely identifies a Grafana workspace.",
			Type:        types.StringType,
			Computed:    true,
		},
		"modification_timestamp": {
			// Property: ModificationTimestamp
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Timestamp when the workspace was last modified",
			//	  "format": "date-time",
			//	  "type": "string"
			//	}
			Description: "Timestamp when the workspace was last modified",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The user friendly name of a workspace.",
			//	  "pattern": "^[a-zA-Z0-9-._~]{1,255}$",
			//	  "type": "string"
			//	}
			Description: "The user friendly name of a workspace.",
			Type:        types.StringType,
			Computed:    true,
		},
		"notification_destinations": {
			// Property: NotificationDestinations
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "List of notification destinations on the customers service managed IAM role that the Grafana workspace can query.",
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "description": "These enums represent valid AWS notification destinations that the Grafana workspace has permission to use. These notification destinations are primarily used to help customers visualize which destinations have been added to a service managed IAM role.",
			//	    "enum": [
			//	      "SNS"
			//	    ],
			//	    "type": "string"
			//	  },
			//	  "type": "array"
			//	}
			Description: "List of notification destinations on the customers service managed IAM role that the Grafana workspace can query.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"organization_role_name": {
			// Property: OrganizationRoleName
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The name of an IAM role that already exists to use with AWS Organizations to access AWS data sources and notification channels in other accounts in an organization.",
			//	  "maxLength": 2048,
			//	  "minLength": 1,
			//	  "type": "string"
			//	}
			Description: "The name of an IAM role that already exists to use with AWS Organizations to access AWS data sources and notification channels in other accounts in an organization.",
			Type:        types.StringType,
			Computed:    true,
		},
		"organizational_units": {
			// Property: OrganizationalUnits
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "List of Organizational Units containing AWS accounts the Grafana workspace can pull data from.",
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "description": "Id of an organizational unit.",
			//	    "type": "string"
			//	  },
			//	  "type": "array"
			//	}
			Description: "List of Organizational Units containing AWS accounts the Grafana workspace can pull data from.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"permission_type": {
			// Property: PermissionType
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "These enums represent valid permission types to use when creating or configuring a Grafana workspace. The SERVICE_MANAGED permission type means the Managed Grafana service will create a workspace IAM role on your behalf. The CUSTOMER_MANAGED permission type means that the customer is expected to provide an IAM role that the Grafana workspace can use to query data sources.",
			//	  "enum": [
			//	    "CUSTOMER_MANAGED",
			//	    "SERVICE_MANAGED"
			//	  ],
			//	  "type": "string"
			//	}
			Description: "These enums represent valid permission types to use when creating or configuring a Grafana workspace. The SERVICE_MANAGED permission type means the Managed Grafana service will create a workspace IAM role on your behalf. The CUSTOMER_MANAGED permission type means that the customer is expected to provide an IAM role that the Grafana workspace can use to query data sources.",
			Type:        types.StringType,
			Computed:    true,
		},
		"role_arn": {
			// Property: RoleArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "IAM Role that will be used to grant the Grafana workspace access to a customers AWS resources.",
			//	  "maxLength": 2048,
			//	  "minLength": 1,
			//	  "type": "string"
			//	}
			Description: "IAM Role that will be used to grant the Grafana workspace access to a customers AWS resources.",
			Type:        types.StringType,
			Computed:    true,
		},
		"saml_configuration": {
			// Property: SamlConfiguration
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "SAML configuration data associated with an AMG workspace.",
			//	  "properties": {
			//	    "AllowedOrganizations": {
			//	      "description": "List of SAML organizations allowed to access Grafana.",
			//	      "insertionOrder": false,
			//	      "items": {
			//	        "description": "A single SAML organization.",
			//	        "maxLength": 256,
			//	        "minLength": 1,
			//	        "type": "string"
			//	      },
			//	      "type": "array"
			//	    },
			//	    "AssertionAttributes": {
			//	      "additionalProperties": false,
			//	      "description": "Maps Grafana friendly names to the IdPs SAML attributes.",
			//	      "properties": {
			//	        "Email": {
			//	          "description": "Name of the attribute within the SAML assert to use as the users email in Grafana.",
			//	          "maxLength": 256,
			//	          "minLength": 1,
			//	          "type": "string"
			//	        },
			//	        "Groups": {
			//	          "description": "Name of the attribute within the SAML assert to use as the users groups in Grafana.",
			//	          "maxLength": 256,
			//	          "minLength": 1,
			//	          "type": "string"
			//	        },
			//	        "Login": {
			//	          "description": "Name of the attribute within the SAML assert to use as the users login handle in Grafana.",
			//	          "maxLength": 256,
			//	          "minLength": 1,
			//	          "type": "string"
			//	        },
			//	        "Name": {
			//	          "description": "Name of the attribute within the SAML assert to use as the users name in Grafana.",
			//	          "maxLength": 256,
			//	          "minLength": 1,
			//	          "type": "string"
			//	        },
			//	        "Org": {
			//	          "description": "Name of the attribute within the SAML assert to use as the users organizations in Grafana.",
			//	          "maxLength": 256,
			//	          "minLength": 1,
			//	          "type": "string"
			//	        },
			//	        "Role": {
			//	          "description": "Name of the attribute within the SAML assert to use as the users roles in Grafana.",
			//	          "maxLength": 256,
			//	          "minLength": 1,
			//	          "type": "string"
			//	        }
			//	      },
			//	      "type": "object"
			//	    },
			//	    "IdpMetadata": {
			//	      "additionalProperties": false,
			//	      "description": "IdP Metadata used to configure SAML authentication in Grafana.",
			//	      "properties": {
			//	        "Url": {
			//	          "description": "URL that vends the IdPs metadata.",
			//	          "maxLength": 2048,
			//	          "minLength": 1,
			//	          "type": "string"
			//	        },
			//	        "Xml": {
			//	          "description": "XML blob of the IdPs metadata.",
			//	          "type": "string"
			//	        }
			//	      },
			//	      "type": "object"
			//	    },
			//	    "LoginValidityDuration": {
			//	      "description": "The maximum lifetime an authenticated user can be logged in (in minutes) before being required to re-authenticate.",
			//	      "type": "number"
			//	    },
			//	    "RoleValues": {
			//	      "additionalProperties": false,
			//	      "description": "Maps SAML roles to the Grafana Editor and Admin roles.",
			//	      "properties": {
			//	        "Admin": {
			//	          "description": "List of SAML roles which will be mapped into the Grafana Admin role.",
			//	          "insertionOrder": false,
			//	          "items": {
			//	            "description": "A single SAML role.",
			//	            "maxLength": 256,
			//	            "minLength": 1,
			//	            "type": "string"
			//	          },
			//	          "type": "array"
			//	        },
			//	        "Editor": {
			//	          "description": "List of SAML roles which will be mapped into the Grafana Editor role.",
			//	          "insertionOrder": false,
			//	          "items": {
			//	            "description": "A single SAML role.",
			//	            "maxLength": 256,
			//	            "minLength": 1,
			//	            "type": "string"
			//	          },
			//	          "type": "array"
			//	        }
			//	      },
			//	      "type": "object"
			//	    }
			//	  },
			//	  "required": [
			//	    "IdpMetadata"
			//	  ],
			//	  "type": "object"
			//	}
			Description: "SAML configuration data associated with an AMG workspace.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"allowed_organizations": {
						// Property: AllowedOrganizations
						Description: "List of SAML organizations allowed to access Grafana.",
						Type:        types.ListType{ElemType: types.StringType},
						Computed:    true,
					},
					"assertion_attributes": {
						// Property: AssertionAttributes
						Description: "Maps Grafana friendly names to the IdPs SAML attributes.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"email": {
									// Property: Email
									Description: "Name of the attribute within the SAML assert to use as the users email in Grafana.",
									Type:        types.StringType,
									Computed:    true,
								},
								"groups": {
									// Property: Groups
									Description: "Name of the attribute within the SAML assert to use as the users groups in Grafana.",
									Type:        types.StringType,
									Computed:    true,
								},
								"login": {
									// Property: Login
									Description: "Name of the attribute within the SAML assert to use as the users login handle in Grafana.",
									Type:        types.StringType,
									Computed:    true,
								},
								"name": {
									// Property: Name
									Description: "Name of the attribute within the SAML assert to use as the users name in Grafana.",
									Type:        types.StringType,
									Computed:    true,
								},
								"org": {
									// Property: Org
									Description: "Name of the attribute within the SAML assert to use as the users organizations in Grafana.",
									Type:        types.StringType,
									Computed:    true,
								},
								"role": {
									// Property: Role
									Description: "Name of the attribute within the SAML assert to use as the users roles in Grafana.",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
					"idp_metadata": {
						// Property: IdpMetadata
						Description: "IdP Metadata used to configure SAML authentication in Grafana.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"url": {
									// Property: Url
									Description: "URL that vends the IdPs metadata.",
									Type:        types.StringType,
									Computed:    true,
								},
								"xml": {
									// Property: Xml
									Description: "XML blob of the IdPs metadata.",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
					"login_validity_duration": {
						// Property: LoginValidityDuration
						Description: "The maximum lifetime an authenticated user can be logged in (in minutes) before being required to re-authenticate.",
						Type:        types.Float64Type,
						Computed:    true,
					},
					"role_values": {
						// Property: RoleValues
						Description: "Maps SAML roles to the Grafana Editor and Admin roles.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"admin": {
									// Property: Admin
									Description: "List of SAML roles which will be mapped into the Grafana Admin role.",
									Type:        types.ListType{ElemType: types.StringType},
									Computed:    true,
								},
								"editor": {
									// Property: Editor
									Description: "List of SAML roles which will be mapped into the Grafana Editor role.",
									Type:        types.ListType{ElemType: types.StringType},
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"saml_configuration_status": {
			// Property: SamlConfigurationStatus
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Valid SAML configuration statuses.",
			//	  "enum": [
			//	    "CONFIGURED",
			//	    "NOT_CONFIGURED"
			//	  ],
			//	  "type": "string"
			//	}
			Description: "Valid SAML configuration statuses.",
			Type:        types.StringType,
			Computed:    true,
		},
		"sso_client_id": {
			// Property: SsoClientId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The client ID of the AWS SSO Managed Application.",
			//	  "type": "string"
			//	}
			Description: "The client ID of the AWS SSO Managed Application.",
			Type:        types.StringType,
			Computed:    true,
		},
		"stack_set_name": {
			// Property: StackSetName
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The name of the AWS CloudFormation stack set to use to generate IAM roles to be used for this workspace.",
			//	  "type": "string"
			//	}
			Description: "The name of the AWS CloudFormation stack set to use to generate IAM roles to be used for this workspace.",
			Type:        types.StringType,
			Computed:    true,
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "These enums represent the status of a workspace.",
			//	  "enum": [
			//	    "ACTIVE",
			//	    "CREATING",
			//	    "DELETING",
			//	    "FAILED",
			//	    "UPDATING",
			//	    "UPGRADING",
			//	    "DELETION_FAILED",
			//	    "CREATION_FAILED",
			//	    "UPDATE_FAILED",
			//	    "UPGRADE_FAILED",
			//	    "LICENSE_REMOVAL_FAILED"
			//	  ],
			//	  "type": "string"
			//	}
			Description: "These enums represent the status of a workspace.",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::Grafana::Workspace",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Grafana::Workspace").WithTerraformTypeName("awscc_grafana_workspace")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"account_access_type":       "AccountAccessType",
		"admin":                     "Admin",
		"allowed_organizations":     "AllowedOrganizations",
		"assertion_attributes":      "AssertionAttributes",
		"authentication_providers":  "AuthenticationProviders",
		"client_token":              "ClientToken",
		"creation_timestamp":        "CreationTimestamp",
		"data_sources":              "DataSources",
		"description":               "Description",
		"editor":                    "Editor",
		"email":                     "Email",
		"endpoint":                  "Endpoint",
		"grafana_version":           "GrafanaVersion",
		"groups":                    "Groups",
		"id":                        "Id",
		"idp_metadata":              "IdpMetadata",
		"login":                     "Login",
		"login_validity_duration":   "LoginValidityDuration",
		"modification_timestamp":    "ModificationTimestamp",
		"name":                      "Name",
		"notification_destinations": "NotificationDestinations",
		"org":                       "Org",
		"organization_role_name":    "OrganizationRoleName",
		"organizational_units":      "OrganizationalUnits",
		"permission_type":           "PermissionType",
		"role":                      "Role",
		"role_arn":                  "RoleArn",
		"role_values":               "RoleValues",
		"saml_configuration":        "SamlConfiguration",
		"saml_configuration_status": "SamlConfigurationStatus",
		"sso_client_id":             "SsoClientId",
		"stack_set_name":            "StackSetName",
		"status":                    "Status",
		"url":                       "Url",
		"xml":                       "Xml",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
