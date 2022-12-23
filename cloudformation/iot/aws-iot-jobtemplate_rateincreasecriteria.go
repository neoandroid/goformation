// Code generated by "go generate". Please don't change this file directly.

package iot

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// JobTemplate_RateIncreaseCriteria AWS CloudFormation Resource (AWS::IoT::JobTemplate.RateIncreaseCriteria)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-jobtemplate-rateincreasecriteria.html
type JobTemplate_RateIncreaseCriteria struct {

	// NumberOfNotifiedThings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-jobtemplate-rateincreasecriteria.html#cfn-iot-jobtemplate-rateincreasecriteria-numberofnotifiedthings
	NumberOfNotifiedThings *int `json:"NumberOfNotifiedThings,omitempty"`

	// NumberOfSucceededThings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-jobtemplate-rateincreasecriteria.html#cfn-iot-jobtemplate-rateincreasecriteria-numberofsucceededthings
	NumberOfSucceededThings *int `json:"NumberOfSucceededThings,omitempty"`

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
func (r *JobTemplate_RateIncreaseCriteria) AWSCloudFormationType() string {
	return "AWS::IoT::JobTemplate.RateIncreaseCriteria"
}