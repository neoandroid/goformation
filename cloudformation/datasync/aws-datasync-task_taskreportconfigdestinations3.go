// Code generated by "go generate". Please don't change this file directly.

package datasync

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Task_TaskReportConfigDestinationS3 AWS CloudFormation Resource (AWS::DataSync::Task.TaskReportConfigDestinationS3)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-taskreportconfigdestinations3.html
type Task_TaskReportConfigDestinationS3 struct {

	// BucketAccessRoleArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-taskreportconfigdestinations3.html#cfn-datasync-task-taskreportconfigdestinations3-bucketaccessrolearn
	BucketAccessRoleArn *string `json:"BucketAccessRoleArn,omitempty"`

	// S3BucketArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-taskreportconfigdestinations3.html#cfn-datasync-task-taskreportconfigdestinations3-s3bucketarn
	S3BucketArn *string `json:"S3BucketArn,omitempty"`

	// Subdirectory AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-taskreportconfigdestinations3.html#cfn-datasync-task-taskreportconfigdestinations3-subdirectory
	Subdirectory *string `json:"Subdirectory,omitempty"`

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
func (r *Task_TaskReportConfigDestinationS3) AWSCloudFormationType() string {
	return "AWS::DataSync::Task.TaskReportConfigDestinationS3"
}
