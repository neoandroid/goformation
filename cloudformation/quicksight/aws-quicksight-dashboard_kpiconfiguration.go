// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Dashboard_KPIConfiguration AWS CloudFormation Resource (AWS::QuickSight::Dashboard.KPIConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-kpiconfiguration.html
type Dashboard_KPIConfiguration struct {

	// FieldWells AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-kpiconfiguration.html#cfn-quicksight-dashboard-kpiconfiguration-fieldwells
	FieldWells *Dashboard_KPIFieldWells `json:"FieldWells,omitempty"`

	// KPIOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-kpiconfiguration.html#cfn-quicksight-dashboard-kpiconfiguration-kpioptions
	KPIOptions *Dashboard_KPIOptions `json:"KPIOptions,omitempty"`

	// SortConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-kpiconfiguration.html#cfn-quicksight-dashboard-kpiconfiguration-sortconfiguration
	SortConfiguration *Dashboard_KPISortConfiguration `json:"SortConfiguration,omitempty"`

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
func (r *Dashboard_KPIConfiguration) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.KPIConfiguration"
}
