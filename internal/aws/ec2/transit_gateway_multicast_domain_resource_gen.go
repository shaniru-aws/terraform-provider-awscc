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
	registry.AddResourceTypeFactory("awscc_ec2_transit_gateway_multicast_domain", transitGatewayMulticastDomainResourceType)
}

// transitGatewayMulticastDomainResourceType returns the Terraform awscc_ec2_transit_gateway_multicast_domain resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::EC2::TransitGatewayMulticastDomain resource type.
func transitGatewayMulticastDomainResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"creation_time": {
			// Property: CreationTime
			// CloudFormation resource type schema:
			// {
			//   "description": "The time the transit gateway multicast domain was created.",
			//   "format": "date-time",
			//   "type": "string"
			// }
			Description: "The time the transit gateway multicast domain was created.",
			Type:        types.StringType,
			Computed:    true,
		},
		"options": {
			// Property: Options
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The options for the transit gateway multicast domain.",
			//   "properties": {
			//     "AutoAcceptSharedAssociations": {
			//       "description": "Indicates whether to automatically cross-account subnet associations that are associated with the transit gateway multicast domain. Valid Values: enable | disable",
			//       "type": "string"
			//     },
			//     "Igmpv2Support": {
			//       "description": "Indicates whether Internet Group Management Protocol (IGMP) version 2 is turned on for the transit gateway multicast domain. Valid Values: enable | disable",
			//       "type": "string"
			//     },
			//     "StaticSourcesSupport": {
			//       "description": "Indicates whether support for statically configuring transit gateway multicast group sources is turned on. Valid Values: enable | disable",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The options for the transit gateway multicast domain.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"auto_accept_shared_associations": {
						// Property: AutoAcceptSharedAssociations
						Description: "Indicates whether to automatically cross-account subnet associations that are associated with the transit gateway multicast domain. Valid Values: enable | disable",
						Type:        types.StringType,
						Optional:    true,
					},
					"igmpv_2_support": {
						// Property: Igmpv2Support
						Description: "Indicates whether Internet Group Management Protocol (IGMP) version 2 is turned on for the transit gateway multicast domain. Valid Values: enable | disable",
						Type:        types.StringType,
						Optional:    true,
					},
					"static_sources_support": {
						// Property: StaticSourcesSupport
						Description: "Indicates whether support for statically configuring transit gateway multicast group sources is turned on. Valid Values: enable | disable",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Optional: true,
		},
		"state": {
			// Property: State
			// CloudFormation resource type schema:
			// {
			//   "description": "The state of the transit gateway multicast domain.",
			//   "type": "string"
			// }
			Description: "The state of the transit gateway multicast domain.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "The tags for the transit gateway multicast domain.",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "description": "The key of the tag. Constraints: Tag keys are case-sensitive and accept a maximum of 127 Unicode characters. May not begin with aws:.",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value of the tag. Constraints: Tag values are case-sensitive and accept a maximum of 255 Unicode characters.",
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "The tags for the transit gateway multicast domain.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key of the tag. Constraints: Tag keys are case-sensitive and accept a maximum of 127 Unicode characters. May not begin with aws:.",
						Type:        types.StringType,
						Optional:    true,
					},
					"value": {
						// Property: Value
						Description: "The value of the tag. Constraints: Tag values are case-sensitive and accept a maximum of 255 Unicode characters.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
		"transit_gateway_id": {
			// Property: TransitGatewayId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the transit gateway.",
			//   "type": "string"
			// }
			Description: "The ID of the transit gateway.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"transit_gateway_multicast_domain_arn": {
			// Property: TransitGatewayMulticastDomainArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the transit gateway multicast domain.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the transit gateway multicast domain.",
			Type:        types.StringType,
			Computed:    true,
		},
		"transit_gateway_multicast_domain_id": {
			// Property: TransitGatewayMulticastDomainId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the transit gateway multicast domain.",
			//   "type": "string"
			// }
			Description: "The ID of the transit gateway multicast domain.",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "The AWS::EC2::TransitGatewayMulticastDomain type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::TransitGatewayMulticastDomain").WithTerraformTypeName("awscc_ec2_transit_gateway_multicast_domain")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"auto_accept_shared_associations":      "AutoAcceptSharedAssociations",
		"creation_time":                        "CreationTime",
		"igmpv_2_support":                      "Igmpv2Support",
		"key":                                  "Key",
		"options":                              "Options",
		"state":                                "State",
		"static_sources_support":               "StaticSourcesSupport",
		"tags":                                 "Tags",
		"transit_gateway_id":                   "TransitGatewayId",
		"transit_gateway_multicast_domain_arn": "TransitGatewayMulticastDomainArn",
		"transit_gateway_multicast_domain_id":  "TransitGatewayMulticastDomainId",
		"value":                                "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_ec2_transit_gateway_multicast_domain", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
