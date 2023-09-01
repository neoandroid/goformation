// Code generated by "go generate". Please don't change this file directly.

package devbatch

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// JobDefinition_LogConfiguration AWS CloudFormation Resource (AWS::DevBatch::JobDefinition.LogConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devbatch-jobdefinition-containerproperties-logconfiguration.html
type JobDefinition_LogConfiguration struct {

	// LogDriver AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devbatch-jobdefinition-containerproperties-logconfiguration.html#cfn-devbatch-jobdefinition-containerproperties-logconfiguration-logdriver
	LogDriver string `json:"LogDriver"`

	// Options AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devbatch-jobdefinition-containerproperties-logconfiguration.html#cfn-devbatch-jobdefinition-containerproperties-logconfiguration-options
	Options interface{} `json:"Options,omitempty"`

	// SecretOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devbatch-jobdefinition-containerproperties-logconfiguration.html#cfn-devbatch-jobdefinition-containerproperties-logconfiguration-secretoptions
	SecretOptions []JobDefinition_Secret `json:"SecretOptions,omitempty"`

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
func (r *JobDefinition_LogConfiguration) AWSCloudFormationType() string {
	return "AWS::DevBatch::JobDefinition.LogConfiguration"
}
