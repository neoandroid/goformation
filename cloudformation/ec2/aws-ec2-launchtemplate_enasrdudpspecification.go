// Code generated by "go generate". Please don't change this file directly.

package ec2

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// LaunchTemplate_EnaSrdUdpSpecification AWS CloudFormation Resource (AWS::EC2::LaunchTemplate.EnaSrdUdpSpecification)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-enasrdudpspecification.html
type LaunchTemplate_EnaSrdUdpSpecification struct {

	// EnaSrdUdpEnabled AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-enasrdudpspecification.html#cfn-ec2-launchtemplate-enasrdudpspecification-enasrdudpenabled
	EnaSrdUdpEnabled *bool `json:"EnaSrdUdpEnabled,omitempty"`

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
func (r *LaunchTemplate_EnaSrdUdpSpecification) AWSCloudFormationType() string {
	return "AWS::EC2::LaunchTemplate.EnaSrdUdpSpecification"
}