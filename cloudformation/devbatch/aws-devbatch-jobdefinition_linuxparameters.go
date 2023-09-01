// Code generated by "go generate". Please don't change this file directly.

package devbatch

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// JobDefinition_LinuxParameters AWS CloudFormation Resource (AWS::DevBatch::JobDefinition.LinuxParameters)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devbatch-jobdefinition-containerproperties-linuxparameters.html
type JobDefinition_LinuxParameters struct {

	// Devices AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devbatch-jobdefinition-containerproperties-linuxparameters.html#cfn-devbatch-jobdefinition-containerproperties-linuxparameters-devices
	Devices []JobDefinition_Device `json:"Devices,omitempty"`

	// InitProcessEnabled AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devbatch-jobdefinition-containerproperties-linuxparameters.html#cfn-devbatch-jobdefinition-containerproperties-linuxparameters-initprocessenabled
	InitProcessEnabled *bool `json:"InitProcessEnabled,omitempty"`

	// MaxSwap AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devbatch-jobdefinition-containerproperties-linuxparameters.html#cfn-devbatch-jobdefinition-containerproperties-linuxparameters-maxswap
	MaxSwap *int `json:"MaxSwap,omitempty"`

	// SharedMemorySize AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devbatch-jobdefinition-containerproperties-linuxparameters.html#cfn-devbatch-jobdefinition-containerproperties-linuxparameters-sharedmemorysize
	SharedMemorySize *int `json:"SharedMemorySize,omitempty"`

	// Swappiness AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devbatch-jobdefinition-containerproperties-linuxparameters.html#cfn-devbatch-jobdefinition-containerproperties-linuxparameters-swappiness
	Swappiness *int `json:"Swappiness,omitempty"`

	// Tmpfs AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devbatch-jobdefinition-containerproperties-linuxparameters.html#cfn-devbatch-jobdefinition-containerproperties-linuxparameters-tmpfs
	Tmpfs []JobDefinition_Tmpfs `json:"Tmpfs,omitempty"`

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
func (r *JobDefinition_LinuxParameters) AWSCloudFormationType() string {
	return "AWS::DevBatch::JobDefinition.LinuxParameters"
}
