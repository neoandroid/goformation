// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Dashboard_AxisDisplayMinMaxRange AWS CloudFormation Resource (AWS::QuickSight::Dashboard.AxisDisplayMinMaxRange)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-axisdisplayminmaxrange.html
type Dashboard_AxisDisplayMinMaxRange struct {

	// Maximum AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-axisdisplayminmaxrange.html#cfn-quicksight-dashboard-axisdisplayminmaxrange-maximum
	Maximum *float64 `json:"Maximum,omitempty"`

	// Minimum AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-axisdisplayminmaxrange.html#cfn-quicksight-dashboard-axisdisplayminmaxrange-minimum
	Minimum *float64 `json:"Minimum,omitempty"`

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
func (r *Dashboard_AxisDisplayMinMaxRange) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.AxisDisplayMinMaxRange"
}
