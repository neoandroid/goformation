// Code generated by "go generate". Please don't change this file directly.

package greengrassv2

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Deployment_DeploymentConfigurationValidationPolicy AWS CloudFormation Resource (AWS::GreengrassV2::Deployment.DeploymentConfigurationValidationPolicy)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-deployment-deploymentconfigurationvalidationpolicy.html
type Deployment_DeploymentConfigurationValidationPolicy struct {

	// TimeoutInSeconds AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-deployment-deploymentconfigurationvalidationpolicy.html#cfn-greengrassv2-deployment-deploymentconfigurationvalidationpolicy-timeoutinseconds
	TimeoutInSeconds *int `json:"TimeoutInSeconds,omitempty"`

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
func (r *Deployment_DeploymentConfigurationValidationPolicy) AWSCloudFormationType() string {
	return "AWS::GreengrassV2::Deployment.DeploymentConfigurationValidationPolicy"
}