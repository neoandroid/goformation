// Code generated by "go generate". Please don't change this file directly.

package sagemaker

import (
	"github.com/awslabs/goformation/v6/cloudformation/policies"
)

// ModelPackage_DriftCheckModelQuality AWS CloudFormation Resource (AWS::SageMaker::ModelPackage.DriftCheckModelQuality)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-driftcheckmodelquality.html
type ModelPackage_DriftCheckModelQuality struct {

	// Constraints AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-driftcheckmodelquality.html#cfn-sagemaker-modelpackage-driftcheckmodelquality-constraints
	Constraints *ModelPackage_MetricsSource `json:"Constraints,omitempty"`

	// Statistics AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-driftcheckmodelquality.html#cfn-sagemaker-modelpackage-driftcheckmodelquality-statistics
	Statistics *ModelPackage_MetricsSource `json:"Statistics,omitempty"`

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
func (r *ModelPackage_DriftCheckModelQuality) AWSCloudFormationType() string {
	return "AWS::SageMaker::ModelPackage.DriftCheckModelQuality"
}
