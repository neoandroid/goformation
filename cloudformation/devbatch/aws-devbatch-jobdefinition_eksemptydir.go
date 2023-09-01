// Code generated by "go generate". Please don't change this file directly.

package devbatch

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// JobDefinition_EksEmptyDir AWS CloudFormation Resource (AWS::DevBatch::JobDefinition.EksEmptyDir)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devbatch-jobdefinition-eksemptydir.html
type JobDefinition_EksEmptyDir struct {

	// Medium AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devbatch-jobdefinition-eksemptydir.html#cfn-devbatch-jobdefinition-eksemptydir-medium
	Medium *string `json:"Medium,omitempty"`

	// SizeLimit AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devbatch-jobdefinition-eksemptydir.html#cfn-devbatch-jobdefinition-eksemptydir-sizelimit
	SizeLimit *string `json:"SizeLimit,omitempty"`

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
func (r *JobDefinition_EksEmptyDir) AWSCloudFormationType() string {
	return "AWS::DevBatch::JobDefinition.EksEmptyDir"
}
