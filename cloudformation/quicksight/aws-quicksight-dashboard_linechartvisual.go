// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Dashboard_LineChartVisual AWS CloudFormation Resource (AWS::QuickSight::Dashboard.LineChartVisual)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartvisual.html
type Dashboard_LineChartVisual struct {

	// Actions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartvisual.html#cfn-quicksight-dashboard-linechartvisual-actions
	Actions []Dashboard_VisualCustomAction `json:"Actions,omitempty"`

	// ChartConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartvisual.html#cfn-quicksight-dashboard-linechartvisual-chartconfiguration
	ChartConfiguration *Dashboard_LineChartConfiguration `json:"ChartConfiguration,omitempty"`

	// ColumnHierarchies AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartvisual.html#cfn-quicksight-dashboard-linechartvisual-columnhierarchies
	ColumnHierarchies []Dashboard_ColumnHierarchy `json:"ColumnHierarchies,omitempty"`

	// Subtitle AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartvisual.html#cfn-quicksight-dashboard-linechartvisual-subtitle
	Subtitle *Dashboard_VisualSubtitleLabelOptions `json:"Subtitle,omitempty"`

	// Title AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartvisual.html#cfn-quicksight-dashboard-linechartvisual-title
	Title *Dashboard_VisualTitleLabelOptions `json:"Title,omitempty"`

	// VisualId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartvisual.html#cfn-quicksight-dashboard-linechartvisual-visualid
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
func (r *Dashboard_LineChartVisual) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.LineChartVisual"
}