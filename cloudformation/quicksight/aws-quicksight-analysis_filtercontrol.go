// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Analysis_FilterControl AWS CloudFormation Resource (AWS::QuickSight::Analysis.FilterControl)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filtercontrol.html
type Analysis_FilterControl struct {

	// DateTimePicker AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filtercontrol.html#cfn-quicksight-analysis-filtercontrol-datetimepicker
	DateTimePicker *Analysis_FilterDateTimePickerControl `json:"DateTimePicker,omitempty"`

	// Dropdown AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filtercontrol.html#cfn-quicksight-analysis-filtercontrol-dropdown
	Dropdown *Analysis_FilterDropDownControl `json:"Dropdown,omitempty"`

	// List AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filtercontrol.html#cfn-quicksight-analysis-filtercontrol-list
	List *Analysis_FilterListControl `json:"List,omitempty"`

	// RelativeDateTime AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filtercontrol.html#cfn-quicksight-analysis-filtercontrol-relativedatetime
	RelativeDateTime *Analysis_FilterRelativeDateTimeControl `json:"RelativeDateTime,omitempty"`

	// Slider AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filtercontrol.html#cfn-quicksight-analysis-filtercontrol-slider
	Slider *Analysis_FilterSliderControl `json:"Slider,omitempty"`

	// TextArea AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filtercontrol.html#cfn-quicksight-analysis-filtercontrol-textarea
	TextArea *Analysis_FilterTextAreaControl `json:"TextArea,omitempty"`

	// TextField AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filtercontrol.html#cfn-quicksight-analysis-filtercontrol-textfield
	TextField *Analysis_FilterTextFieldControl `json:"TextField,omitempty"`

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
func (r *Analysis_FilterControl) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.FilterControl"
}
