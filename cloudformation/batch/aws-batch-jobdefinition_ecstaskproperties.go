// Code generated by "go generate". Please don't change this file directly.

package batch

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// JobDefinition_EcsTaskProperties AWS CloudFormation Resource (AWS::Batch::JobDefinition.EcsTaskProperties)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ecstaskproperties.html
type JobDefinition_EcsTaskProperties struct {

	// Containers AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ecstaskproperties.html#cfn-batch-jobdefinition-ecstaskproperties-containers
	Containers []JobDefinition_TaskContainerProperties `json:"Containers,omitempty"`

	// EphemeralStorage AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ecstaskproperties.html#cfn-batch-jobdefinition-ecstaskproperties-ephemeralstorage
	EphemeralStorage *JobDefinition_EphemeralStorage `json:"EphemeralStorage,omitempty"`

	// ExecutionRoleArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ecstaskproperties.html#cfn-batch-jobdefinition-ecstaskproperties-executionrolearn
	ExecutionRoleArn *string `json:"ExecutionRoleArn,omitempty"`

	// IpcMode AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ecstaskproperties.html#cfn-batch-jobdefinition-ecstaskproperties-ipcmode
	IpcMode *string `json:"IpcMode,omitempty"`

	// NetworkConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ecstaskproperties.html#cfn-batch-jobdefinition-ecstaskproperties-networkconfiguration
	NetworkConfiguration *JobDefinition_NetworkConfiguration `json:"NetworkConfiguration,omitempty"`

	// PidMode AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ecstaskproperties.html#cfn-batch-jobdefinition-ecstaskproperties-pidmode
	PidMode *string `json:"PidMode,omitempty"`

	// PlatformVersion AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ecstaskproperties.html#cfn-batch-jobdefinition-ecstaskproperties-platformversion
	PlatformVersion *string `json:"PlatformVersion,omitempty"`

	// RuntimePlatform AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ecstaskproperties.html#cfn-batch-jobdefinition-ecstaskproperties-runtimeplatform
	RuntimePlatform *JobDefinition_RuntimePlatform `json:"RuntimePlatform,omitempty"`

	// TaskRoleArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ecstaskproperties.html#cfn-batch-jobdefinition-ecstaskproperties-taskrolearn
	TaskRoleArn *string `json:"TaskRoleArn,omitempty"`

	// Volumes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ecstaskproperties.html#cfn-batch-jobdefinition-ecstaskproperties-volumes
	Volumes []JobDefinition_Volumes `json:"Volumes,omitempty"`

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
func (r *JobDefinition_EcsTaskProperties) AWSCloudFormationType() string {
	return "AWS::Batch::JobDefinition.EcsTaskProperties"
}