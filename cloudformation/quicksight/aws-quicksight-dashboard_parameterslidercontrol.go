// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Dashboard_ParameterSliderControl AWS CloudFormation Resource (AWS::QuickSight::Dashboard.ParameterSliderControl)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-parameterslidercontrol.html
type Dashboard_ParameterSliderControl struct {

	// DisplayOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-parameterslidercontrol.html#cfn-quicksight-dashboard-parameterslidercontrol-displayoptions
	DisplayOptions *Dashboard_SliderControlDisplayOptions `json:"DisplayOptions,omitempty"`

	// MaximumValue AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-parameterslidercontrol.html#cfn-quicksight-dashboard-parameterslidercontrol-maximumvalue
	MaximumValue float64 `json:"MaximumValue"`

	// MinimumValue AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-parameterslidercontrol.html#cfn-quicksight-dashboard-parameterslidercontrol-minimumvalue
	MinimumValue float64 `json:"MinimumValue"`

	// ParameterControlId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-parameterslidercontrol.html#cfn-quicksight-dashboard-parameterslidercontrol-parametercontrolid
	ParameterControlId string `json:"ParameterControlId"`

	// SourceParameterName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-parameterslidercontrol.html#cfn-quicksight-dashboard-parameterslidercontrol-sourceparametername
	SourceParameterName string `json:"SourceParameterName"`

	// StepSize AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-parameterslidercontrol.html#cfn-quicksight-dashboard-parameterslidercontrol-stepsize
	StepSize float64 `json:"StepSize"`

	// Title AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-parameterslidercontrol.html#cfn-quicksight-dashboard-parameterslidercontrol-title
	Title string `json:"Title"`

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
func (r *Dashboard_ParameterSliderControl) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.ParameterSliderControl"
}
