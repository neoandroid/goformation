// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Dashboard_ParameterControl AWS CloudFormation Resource (AWS::QuickSight::Dashboard.ParameterControl)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-parametercontrol.html
type Dashboard_ParameterControl struct {

	// DateTimePicker AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-parametercontrol.html#cfn-quicksight-dashboard-parametercontrol-datetimepicker
	DateTimePicker *Dashboard_ParameterDateTimePickerControl `json:"DateTimePicker,omitempty"`

	// Dropdown AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-parametercontrol.html#cfn-quicksight-dashboard-parametercontrol-dropdown
	Dropdown *Dashboard_ParameterDropDownControl `json:"Dropdown,omitempty"`

	// List AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-parametercontrol.html#cfn-quicksight-dashboard-parametercontrol-list
	List *Dashboard_ParameterListControl `json:"List,omitempty"`

	// Slider AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-parametercontrol.html#cfn-quicksight-dashboard-parametercontrol-slider
	Slider *Dashboard_ParameterSliderControl `json:"Slider,omitempty"`

	// TextArea AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-parametercontrol.html#cfn-quicksight-dashboard-parametercontrol-textarea
	TextArea *Dashboard_ParameterTextAreaControl `json:"TextArea,omitempty"`

	// TextField AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-parametercontrol.html#cfn-quicksight-dashboard-parametercontrol-textfield
	TextField *Dashboard_ParameterTextFieldControl `json:"TextField,omitempty"`

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
func (r *Dashboard_ParameterControl) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.ParameterControl"
}
