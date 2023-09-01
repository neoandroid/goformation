// Code generated by "go generate". Please don't change this file directly.

package devbatch

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// JobDefinition_Environment AWS CloudFormation Resource (AWS::DevBatch::JobDefinition.Environment)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devbatch-jobdefinition-environment.html
type JobDefinition_Environment struct {

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devbatch-jobdefinition-environment.html#cfn-devbatch-jobdefinition-environment-name
	Name *string `json:"Name,omitempty"`

	// Value AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devbatch-jobdefinition-environment.html#cfn-devbatch-jobdefinition-environment-value
	Value *string `json:"Value,omitempty"`

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
func (r *JobDefinition_Environment) AWSCloudFormationType() string {
	return "AWS::DevBatch::JobDefinition.Environment"
}
