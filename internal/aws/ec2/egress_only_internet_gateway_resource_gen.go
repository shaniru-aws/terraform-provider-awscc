// Code generated by generators/resource/main.go; DO NOT EDIT.

package ec2

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
	registry.AddResourceTypeFactory("awscc_ec2_egress_only_internet_gateway", egressOnlyInternetGatewayResourceType)
}

// egressOnlyInternetGatewayResourceType returns the Terraform awscc_ec2_egress_only_internet_gateway resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::EC2::EgressOnlyInternetGateway resource type.
func egressOnlyInternetGatewayResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "Service Generated ID of the EgressOnlyInternetGateway",
			//   "type": "string"
			// }
			Description: "Service Generated ID of the EgressOnlyInternetGateway",
			Type:        types.StringType,
			Computed:    true,
		},
		"vpc_id": {
			// Property: VpcId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the VPC for which to create the egress-only internet gateway.",
			//   "type": "string"
			// }
			Description: "The ID of the VPC for which to create the egress-only internet gateway.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::EC2::EgressOnlyInternetGateway",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::EgressOnlyInternetGateway").WithTerraformTypeName("awscc_ec2_egress_only_internet_gateway")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"id":     "Id",
		"vpc_id": "VpcId",
	})

	opts = opts.IsImmutableType(true)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_ec2_egress_only_internet_gateway", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
