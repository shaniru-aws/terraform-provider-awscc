// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package refactorspaces

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_refactorspaces_service", serviceDataSource)
}

// serviceDataSource returns the Terraform awscc_refactorspaces_service data source.
// This Terraform data source corresponds to the CloudFormation AWS::RefactorSpaces::Service resource.
func serviceDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"application_identifier": {
			// Property: ApplicationIdentifier
			// CloudFormation resource type schema:
			//
			//	{
			//	  "maxLength": 14,
			//	  "minLength": 14,
			//	  "pattern": "^app-([0-9A-Za-z]{10}$)",
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "maxLength": 2048,
			//	  "minLength": 20,
			//	  "pattern": "^arn:(aws[a-zA-Z-]*)?:refactor-spaces:[a-zA-Z0-9\\-]+:\\w{12}:[a-zA-Z_0-9+=,.@\\-_/]+$",
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			//
			//	{
			//	  "maxLength": 256,
			//	  "minLength": 1,
			//	  "pattern": "^[a-zA-Z0-9-_\\s\\.\\!\\*\\#\\@\\']+$",
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"endpoint_type": {
			// Property: EndpointType
			// CloudFormation resource type schema:
			//
			//	{
			//	  "enum": [
			//	    "LAMBDA",
			//	    "URL"
			//	  ],
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"environment_identifier": {
			// Property: EnvironmentIdentifier
			// CloudFormation resource type schema:
			//
			//	{
			//	  "maxLength": 14,
			//	  "minLength": 14,
			//	  "pattern": "^env-([0-9A-Za-z]{10}$)",
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"lambda_endpoint": {
			// Property: LambdaEndpoint
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "properties": {
			//	    "Arn": {
			//	      "maxLength": 2048,
			//	      "minLength": 1,
			//	      "pattern": "^arn:(aws[a-zA-Z-]*)?:lambda:[a-z]{2}((-gov)|(-iso(b?)))?-[a-z]+-\\d{1}:\\d{12}:function:[a-zA-Z0-9-_]+(:(\\$LATEST|[a-zA-Z0-9-_]+))?$",
			//	      "type": "string"
			//	    }
			//	  },
			//	  "required": [
			//	    "Arn"
			//	  ],
			//	  "type": "object"
			//	}
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"arn": {
						// Property: Arn
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			//
			//	{
			//	  "maxLength": 63,
			//	  "minLength": 3,
			//	  "pattern": "",
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"service_identifier": {
			// Property: ServiceIdentifier
			// CloudFormation resource type schema:
			//
			//	{
			//	  "maxLength": 14,
			//	  "minLength": 14,
			//	  "pattern": "^svc-([0-9A-Za-z]{10}$)",
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Metadata that you can assign to help organize the frameworks that you create. Each tag is a key-value pair.",
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "description": "A label for tagging Environment resource",
			//	    "properties": {
			//	      "Key": {
			//	        "description": "A string used to identify this tag",
			//	        "maxLength": 128,
			//	        "minLength": 1,
			//	        "pattern": "",
			//	        "type": "string"
			//	      },
			//	      "Value": {
			//	        "description": "A string containing the value for the tag",
			//	        "maxLength": 256,
			//	        "minLength": 0,
			//	        "type": "string"
			//	      }
			//	    },
			//	    "required": [
			//	      "Key",
			//	      "Value"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "type": "array"
			//	}
			Description: "Metadata that you can assign to help organize the frameworks that you create. Each tag is a key-value pair.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "A string used to identify this tag",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "A string containing the value for the tag",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"url_endpoint": {
			// Property: UrlEndpoint
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "properties": {
			//	    "HealthUrl": {
			//	      "maxLength": 2048,
			//	      "minLength": 1,
			//	      "pattern": "^https?://[-a-zA-Z0-9+\\x38@#/%?=~_|!:,.;]*[-a-zA-Z0-9+\\x38@#/%=~_|]$",
			//	      "type": "string"
			//	    },
			//	    "Url": {
			//	      "maxLength": 2048,
			//	      "minLength": 1,
			//	      "pattern": "^https?://[-a-zA-Z0-9+\\x38@#/%?=~_|!:,.;]*[-a-zA-Z0-9+\\x38@#/%=~_|]$",
			//	      "type": "string"
			//	    }
			//	  },
			//	  "required": [
			//	    "Url"
			//	  ],
			//	  "type": "object"
			//	}
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"health_url": {
						// Property: HealthUrl
						Type:     types.StringType,
						Computed: true,
					},
					"url": {
						// Property: Url
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"vpc_id": {
			// Property: VpcId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "maxLength": 21,
			//	  "minLength": 12,
			//	  "pattern": "^vpc-[-a-f0-9]{8}([-a-f0-9]{9})?$",
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::RefactorSpaces::Service",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::RefactorSpaces::Service").WithTerraformTypeName("awscc_refactorspaces_service")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"application_identifier": "ApplicationIdentifier",
		"arn":                    "Arn",
		"description":            "Description",
		"endpoint_type":          "EndpointType",
		"environment_identifier": "EnvironmentIdentifier",
		"health_url":             "HealthUrl",
		"key":                    "Key",
		"lambda_endpoint":        "LambdaEndpoint",
		"name":                   "Name",
		"service_identifier":     "ServiceIdentifier",
		"tags":                   "Tags",
		"url":                    "Url",
		"url_endpoint":           "UrlEndpoint",
		"value":                  "Value",
		"vpc_id":                 "VpcId",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
