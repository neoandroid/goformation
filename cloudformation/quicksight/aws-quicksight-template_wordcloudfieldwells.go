// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Template_WordCloudFieldWells AWS CloudFormation Resource (AWS::QuickSight::Template.WordCloudFieldWells)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-wordcloudfieldwells.html
type Template_WordCloudFieldWells struct {

	// WordCloudAggregatedFieldWells AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-wordcloudfieldwells.html#cfn-quicksight-template-wordcloudfieldwells-wordcloudaggregatedfieldwells
	WordCloudAggregatedFieldWells *Template_WordCloudAggregatedFieldWells `json:"WordCloudAggregatedFieldWells,omitempty"`

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
func (r *Template_WordCloudFieldWells) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.WordCloudFieldWells"
}
