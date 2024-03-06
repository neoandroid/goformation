// Code generated by "go generate". Please don't change this file directly.

package codepipeline

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Pipeline_GitPushFilter AWS CloudFormation Resource (AWS::CodePipeline::Pipeline.GitPushFilter)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-triggers-git-push-filter.html
type Pipeline_GitPushFilter struct {

	// Branches AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-triggers-git-push-filter.html#aws-properties-codepipeline-pipeline-triggers-git-branch-filter-criteria
	Branches *Pipeline_GitBranchFilterCriteria `json:"Branches,omitempty"`

	// FilePaths AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-triggers-git-push-filter.html#aws-properties-codepipeline-pipeline-triggers-git-file-path-filter-criteria
	FilePaths *Pipeline_GitFilePathFilterCriteria `json:"FilePaths,omitempty"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-triggers-git-push-filter.html#aws-properties-codepipeline-pipeline-triggers-git-tag-filter-criteria
	Tags *Pipeline_GitTagFilterCriteria `json:"Tags,omitempty"`

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
func (r *Pipeline_GitPushFilter) AWSCloudFormationType() string {
	return "AWS::CodePipeline::Pipeline.GitPushFilter"
}
