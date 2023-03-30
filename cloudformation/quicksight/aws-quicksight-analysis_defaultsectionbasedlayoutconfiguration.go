// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Analysis_DefaultSectionBasedLayoutConfiguration AWS CloudFormation Resource (AWS::QuickSight::Analysis.DefaultSectionBasedLayoutConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-defaultsectionbasedlayoutconfiguration.html
type Analysis_DefaultSectionBasedLayoutConfiguration struct {

	// CanvasSizeOptions AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-defaultsectionbasedlayoutconfiguration.html#cfn-quicksight-analysis-defaultsectionbasedlayoutconfiguration-canvassizeoptions
	CanvasSizeOptions *Analysis_SectionBasedLayoutCanvasSizeOptions `json:"CanvasSizeOptions"`

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
func (r *Analysis_DefaultSectionBasedLayoutConfiguration) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.DefaultSectionBasedLayoutConfiguration"
}
