// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package appflow

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_appflow_connector", connectorDataSource)
}

// connectorDataSource returns the Terraform awscc_appflow_connector data source.
// This Terraform data source corresponds to the CloudFormation AWS::AppFlow::Connector resource.
func connectorDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"connector_arn": {
			// Property: ConnectorArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": " The arn of the connector. The arn is unique for each ConnectorRegistration in your AWS account.",
			//	  "maxLength": 512,
			//	  "pattern": "arn:*:appflow:.*:[0-9]+:.*",
			//	  "type": "string"
			//	}
			Description: " The arn of the connector. The arn is unique for each ConnectorRegistration in your AWS account.",
			Type:        types.StringType,
			Computed:    true,
		},
		"connector_label": {
			// Property: ConnectorLabel
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": " The name of the connector. The name is unique for each ConnectorRegistration in your AWS account.",
			//	  "maxLength": 512,
			//	  "pattern": "[a-zA-Z0-9][\\w!@#.-]+",
			//	  "type": "string"
			//	}
			Description: " The name of the connector. The name is unique for each ConnectorRegistration in your AWS account.",
			Type:        types.StringType,
			Computed:    true,
		},
		"connector_provisioning_config": {
			// Property: ConnectorProvisioningConfig
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "Contains information about the configuration of the connector being registered.",
			//	  "properties": {
			//	    "Lambda": {
			//	      "additionalProperties": false,
			//	      "description": "Contains information about the configuration of the lambda which is being registered as the connector.",
			//	      "properties": {
			//	        "LambdaArn": {
			//	          "description": "Lambda ARN of the connector being registered.",
			//	          "maxLength": 512,
			//	          "pattern": "arn:*:.*:.*:[0-9]+:.*",
			//	          "type": "string"
			//	        }
			//	      },
			//	      "required": [
			//	        "LambdaArn"
			//	      ],
			//	      "type": "object"
			//	    }
			//	  },
			//	  "type": "object"
			//	}
			Description: "Contains information about the configuration of the connector being registered.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"lambda": {
						// Property: Lambda
						Description: "Contains information about the configuration of the lambda which is being registered as the connector.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"lambda_arn": {
									// Property: LambdaArn
									Description: "Lambda ARN of the connector being registered.",
									Type:        types.StringType,
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
		"connector_provisioning_type": {
			// Property: ConnectorProvisioningType
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The provisioning type of the connector. Currently the only supported value is LAMBDA. ",
			//	  "maxLength": 256,
			//	  "minLength": 1,
			//	  "pattern": "[a-zA-Z0-9][\\w!@#.-]+",
			//	  "type": "string"
			//	}
			Description: "The provisioning type of the connector. Currently the only supported value is LAMBDA. ",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "A description about the connector that's being registered.",
			//	  "maxLength": 2048,
			//	  "pattern": "[\\s\\w/!@#+=.-]*",
			//	  "type": "string"
			//	}
			Description: "A description about the connector that's being registered.",
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
		Description: "Data Source schema for AWS::AppFlow::Connector",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::AppFlow::Connector").WithTerraformTypeName("awscc_appflow_connector")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"connector_arn":                 "ConnectorArn",
		"connector_label":               "ConnectorLabel",
		"connector_provisioning_config": "ConnectorProvisioningConfig",
		"connector_provisioning_type":   "ConnectorProvisioningType",
		"description":                   "Description",
		"lambda":                        "Lambda",
		"lambda_arn":                    "LambdaArn",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}