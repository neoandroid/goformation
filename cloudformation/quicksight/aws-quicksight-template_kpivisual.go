// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Template_KPIVisual AWS CloudFormation Resource (AWS::QuickSight::Template.KPIVisual)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-kpivisual.html
type Template_KPIVisual struct {

	// Actions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-kpivisual.html#cfn-quicksight-template-kpivisual-actions
	Actions []Template_VisualCustomAction `json:"Actions,omitempty"`

	// ChartConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-kpivisual.html#cfn-quicksight-template-kpivisual-chartconfiguration
	ChartConfiguration *Template_KPIConfiguration `json:"ChartConfiguration,omitempty"`

	// ColumnHierarchies AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-kpivisual.html#cfn-quicksight-template-kpivisual-columnhierarchies
	ColumnHierarchies []Template_ColumnHierarchy `json:"ColumnHierarchies,omitempty"`

	// ConditionalFormatting AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-kpivisual.html#cfn-quicksight-template-kpivisual-conditionalformatting
	ConditionalFormatting *Template_KPIConditionalFormatting `json:"ConditionalFormatting,omitempty"`

	// Subtitle AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-kpivisual.html#cfn-quicksight-template-kpivisual-subtitle
	Subtitle *Template_VisualSubtitleLabelOptions `json:"Subtitle,omitempty"`

	// Title AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-kpivisual.html#cfn-quicksight-template-kpivisual-title
	Title *Template_VisualTitleLabelOptions `json:"Title,omitempty"`

	// VisualId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-kpivisual.html#cfn-quicksight-template-kpivisual-visualid
	VisualId string `json:"VisualId"`

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
func (r *Template_KPIVisual) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.KPIVisual"
}
