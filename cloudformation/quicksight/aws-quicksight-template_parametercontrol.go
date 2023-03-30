// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Template_ParameterControl AWS CloudFormation Resource (AWS::QuickSight::Template.ParameterControl)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-parametercontrol.html
type Template_ParameterControl struct {

	// DateTimePicker AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-parametercontrol.html#cfn-quicksight-template-parametercontrol-datetimepicker
	DateTimePicker *Template_ParameterDateTimePickerControl `json:"DateTimePicker,omitempty"`

	// Dropdown AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-parametercontrol.html#cfn-quicksight-template-parametercontrol-dropdown
	Dropdown *Template_ParameterDropDownControl `json:"Dropdown,omitempty"`

	// List AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-parametercontrol.html#cfn-quicksight-template-parametercontrol-list
	List *Template_ParameterListControl `json:"List,omitempty"`

	// Slider AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-parametercontrol.html#cfn-quicksight-template-parametercontrol-slider
	Slider *Template_ParameterSliderControl `json:"Slider,omitempty"`

	// TextArea AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-parametercontrol.html#cfn-quicksight-template-parametercontrol-textarea
	TextArea *Template_ParameterTextAreaControl `json:"TextArea,omitempty"`

	// TextField AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-parametercontrol.html#cfn-quicksight-template-parametercontrol-textfield
	TextField *Template_ParameterTextFieldControl `json:"TextField,omitempty"`

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
func (r *Template_ParameterControl) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.ParameterControl"
}
