// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package apigatewayv2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_apigatewayv2_authorizer", authorizerDataSource)
}

// authorizerDataSource returns the Terraform awscc_apigatewayv2_authorizer data source.
// This Terraform data source corresponds to the CloudFormation AWS::ApiGatewayV2::Authorizer resource.
func authorizerDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ApiId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"api_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: AuthorizerCredentialsArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"authorizer_credentials_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: AuthorizerId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"authorizer_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: AuthorizerPayloadFormatVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"authorizer_payload_format_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: AuthorizerResultTtlInSeconds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "integer"
		//	}
		"authorizer_result_ttl_in_seconds": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: AuthorizerType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"authorizer_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: AuthorizerUri
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"authorizer_uri": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: EnableSimpleResponses
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"enable_simple_responses": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: IdentitySource
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"identity_source": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IdentityValidationExpression
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"identity_validation_expression": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: JwtConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "Audience": {
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": false
		//	    },
		//	    "Issuer": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"jwt_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Audience
				"audience": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Issuer
				"issuer": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::ApiGatewayV2::Authorizer",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ApiGatewayV2::Authorizer").WithTerraformTypeName("awscc_apigatewayv2_authorizer")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"api_id":                            "ApiId",
		"audience":                          "Audience",
		"authorizer_credentials_arn":        "AuthorizerCredentialsArn",
		"authorizer_id":                     "AuthorizerId",
		"authorizer_payload_format_version": "AuthorizerPayloadFormatVersion",
		"authorizer_result_ttl_in_seconds":  "AuthorizerResultTtlInSeconds",
		"authorizer_type":                   "AuthorizerType",
		"authorizer_uri":                    "AuthorizerUri",
		"enable_simple_responses":           "EnableSimpleResponses",
		"identity_source":                   "IdentitySource",
		"identity_validation_expression":    "IdentityValidationExpression",
		"issuer":                            "Issuer",
		"jwt_configuration":                 "JwtConfiguration",
		"name":                              "Name",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
