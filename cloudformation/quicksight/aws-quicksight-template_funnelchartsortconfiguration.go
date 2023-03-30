// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Template_FunnelChartSortConfiguration AWS CloudFormation Resource (AWS::QuickSight::Template.FunnelChartSortConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-funnelchartsortconfiguration.html
type Template_FunnelChartSortConfiguration struct {

	// CategoryItemsLimit AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-funnelchartsortconfiguration.html#cfn-quicksight-template-funnelchartsortconfiguration-categoryitemslimit
	CategoryItemsLimit *Template_ItemsLimitConfiguration `json:"CategoryItemsLimit,omitempty"`

	// CategorySort AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-funnelchartsortconfiguration.html#cfn-quicksight-template-funnelchartsortconfiguration-categorysort
	CategorySort []Template_FieldSortOptions `json:"CategorySort,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *Template_FunnelChartSortConfiguration) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.FunnelChartSortConfiguration"
}
