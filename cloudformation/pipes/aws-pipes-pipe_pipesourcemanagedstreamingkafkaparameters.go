// Code generated by "go generate". Please don't change this file directly.

package pipes

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Pipe_PipeSourceManagedStreamingKafkaParameters AWS CloudFormation Resource (AWS::Pipes::Pipe.PipeSourceManagedStreamingKafkaParameters)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourcemanagedstreamingkafkaparameters.html
type Pipe_PipeSourceManagedStreamingKafkaParameters struct {

	// BatchSize AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourcemanagedstreamingkafkaparameters.html#cfn-pipes-pipe-pipesourcemanagedstreamingkafkaparameters-batchsize
	BatchSize *int `json:"BatchSize,omitempty"`

	// ConsumerGroupID AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourcemanagedstreamingkafkaparameters.html#cfn-pipes-pipe-pipesourcemanagedstreamingkafkaparameters-consumergroupid
	ConsumerGroupID *string `json:"ConsumerGroupID,omitempty"`

	// Credentials AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourcemanagedstreamingkafkaparameters.html#cfn-pipes-pipe-pipesourcemanagedstreamingkafkaparameters-credentials
	Credentials *Pipe_MSKAccessCredentials `json:"Credentials,omitempty"`

	// MaximumBatchingWindowInSeconds AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourcemanagedstreamingkafkaparameters.html#cfn-pipes-pipe-pipesourcemanagedstreamingkafkaparameters-maximumbatchingwindowinseconds
	MaximumBatchingWindowInSeconds *int `json:"MaximumBatchingWindowInSeconds,omitempty"`

	// StartingPosition AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourcemanagedstreamingkafkaparameters.html#cfn-pipes-pipe-pipesourcemanagedstreamingkafkaparameters-startingposition
	StartingPosition *string `json:"StartingPosition,omitempty"`

	// TopicName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourcemanagedstreamingkafkaparameters.html#cfn-pipes-pipe-pipesourcemanagedstreamingkafkaparameters-topicname
	TopicName string `json:"TopicName"`

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
func (r *Pipe_PipeSourceManagedStreamingKafkaParameters) AWSCloudFormationType() string {
	return "AWS::Pipes::Pipe.PipeSourceManagedStreamingKafkaParameters"
}
