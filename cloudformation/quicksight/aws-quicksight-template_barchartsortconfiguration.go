// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Template_BarChartSortConfiguration AWS CloudFormation Resource (AWS::QuickSight::Template.BarChartSortConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-barchartsortconfiguration.html
type Template_BarChartSortConfiguration struct {

	// CategoryItemsLimit AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-barchartsortconfiguration.html#cfn-quicksight-template-barchartsortconfiguration-categoryitemslimit
	CategoryItemsLimit *Template_ItemsLimitConfiguration `json:"CategoryItemsLimit,omitempty"`

	// CategorySort AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-barchartsortconfiguration.html#cfn-quicksight-template-barchartsortconfiguration-categorysort
	CategorySort []Template_FieldSortOptions `json:"CategorySort,omitempty"`

	// ColorItemsLimit AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-barchartsortconfiguration.html#cfn-quicksight-template-barchartsortconfiguration-coloritemslimit
	ColorItemsLimit *Template_ItemsLimitConfiguration `json:"ColorItemsLimit,omitempty"`

	// ColorSort AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-barchartsortconfiguration.html#cfn-quicksight-template-barchartsortconfiguration-colorsort
	ColorSort []Template_FieldSortOptions `json:"ColorSort,omitempty"`

	// SmallMultiplesLimitConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-barchartsortconfiguration.html#cfn-quicksight-template-barchartsortconfiguration-smallmultipleslimitconfiguration
	SmallMultiplesLimitConfiguration *Template_ItemsLimitConfiguration `json:"SmallMultiplesLimitConfiguration,omitempty"`

	// SmallMultiplesSort AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-barchartsortconfiguration.html#cfn-quicksight-template-barchartsortconfiguration-smallmultiplessort
	SmallMultiplesSort []Template_FieldSortOptions `json:"SmallMultiplesSort,omitempty"`

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
func (r *Template_BarChartSortConfiguration) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.BarChartSortConfiguration"
}
