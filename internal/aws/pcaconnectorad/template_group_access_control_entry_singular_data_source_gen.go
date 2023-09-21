// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package pcaconnectorad

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_pcaconnectorad_template_group_access_control_entry", templateGroupAccessControlEntryDataSource)
}

// templateGroupAccessControlEntryDataSource returns the Terraform awscc_pcaconnectorad_template_group_access_control_entry data source.
// This Terraform data source corresponds to the CloudFormation AWS::PCAConnectorAD::TemplateGroupAccessControlEntry resource.
func templateGroupAccessControlEntryDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AccessRights
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "AutoEnroll": {
		//	      "enum": [
		//	        "ALLOW",
		//	        "DENY"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "Enroll": {
		//	      "enum": [
		//	        "ALLOW",
		//	        "DENY"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"access_rights": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AutoEnroll
				"auto_enroll": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Enroll
				"enroll": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: GroupDisplayName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 256,
		//	  "minLength": 0,
		//	  "pattern": "^[\\x20-\\x7E]+$",
		//	  "type": "string"
		//	}
		"group_display_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: GroupSecurityIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 256,
		//	  "minLength": 7,
		//	  "pattern": "^S-[0-9]-([0-9]+-){1,14}[0-9]+$",
		//	  "type": "string"
		//	}
		"group_security_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: TemplateArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 200,
		//	  "minLength": 5,
		//	  "pattern": "^arn:[\\w-]+:pca-connector-ad:[\\w-]+:[0-9]+:connector(\\/[\\w-]+)\\/template(\\/[\\w-]+)$",
		//	  "type": "string"
		//	}
		"template_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::PCAConnectorAD::TemplateGroupAccessControlEntry",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::PCAConnectorAD::TemplateGroupAccessControlEntry").WithTerraformTypeName("awscc_pcaconnectorad_template_group_access_control_entry")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"access_rights":             "AccessRights",
		"auto_enroll":               "AutoEnroll",
		"enroll":                    "Enroll",
		"group_display_name":        "GroupDisplayName",
		"group_security_identifier": "GroupSecurityIdentifier",
		"template_arn":              "TemplateArn",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}