// Code generated by "go generate". Please don't change this file directly.

package securitylake

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// DataLake_ReplicationConfiguration AWS CloudFormation Resource (AWS::SecurityLake::DataLake.ReplicationConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-datalake-replicationconfiguration.html
type DataLake_ReplicationConfiguration struct {

	// Regions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-datalake-replicationconfiguration.html#cfn-securitylake-datalake-replicationconfiguration-regions
	Regions []string `json:"Regions,omitempty"`

	// RoleArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-datalake-replicationconfiguration.html#cfn-securitylake-datalake-replicationconfiguration-rolearn
	RoleArn *string `json:"RoleArn,omitempty"`

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
func (r *DataLake_ReplicationConfiguration) AWSCloudFormationType() string {
	return "AWS::SecurityLake::DataLake.ReplicationConfiguration"
}
