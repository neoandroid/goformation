// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Dashboard_DataFieldSeriesItem AWS CloudFormation Resource (AWS::QuickSight::Dashboard.DataFieldSeriesItem)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datafieldseriesitem.html
type Dashboard_DataFieldSeriesItem struct {

	// AxisBinding AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datafieldseriesitem.html#cfn-quicksight-dashboard-datafieldseriesitem-axisbinding
	AxisBinding string `json:"AxisBinding"`

	// FieldId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datafieldseriesitem.html#cfn-quicksight-dashboard-datafieldseriesitem-fieldid
	FieldId string `json:"FieldId"`

	// FieldValue AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datafieldseriesitem.html#cfn-quicksight-dashboard-datafieldseriesitem-fieldvalue
	FieldValue *string `json:"FieldValue,omitempty"`

	// Settings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datafieldseriesitem.html#cfn-quicksight-dashboard-datafieldseriesitem-settings
	Settings *Dashboard_LineChartSeriesSettings `json:"Settings,omitempty"`

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
func (r *Dashboard_DataFieldSeriesItem) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.DataFieldSeriesItem"
}
