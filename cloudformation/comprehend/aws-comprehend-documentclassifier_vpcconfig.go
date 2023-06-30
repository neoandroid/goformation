// Code generated by "go generate". Please don't change this file directly.

package comprehend

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// DocumentClassifier_VpcConfig AWS CloudFormation Resource (AWS::Comprehend::DocumentClassifier.VpcConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-documentclassifier-vpcconfig.html
type DocumentClassifier_VpcConfig struct {

	// SecurityGroupIds AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-documentclassifier-vpcconfig.html#cfn-comprehend-documentclassifier-vpcconfig-securitygroupids
	SecurityGroupIds []string `json:"SecurityGroupIds"`

	// Subnets AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-documentclassifier-vpcconfig.html#cfn-comprehend-documentclassifier-vpcconfig-subnets
	Subnets []string `json:"Subnets"`

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
func (r *DocumentClassifier_VpcConfig) AWSCloudFormationType() string {
	return "AWS::Comprehend::DocumentClassifier.VpcConfig"
}
