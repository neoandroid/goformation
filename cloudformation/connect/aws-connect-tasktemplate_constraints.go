// Code generated by "go generate". Please don't change this file directly.

package connect

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// TaskTemplate_Constraints AWS CloudFormation Resource (AWS::Connect::TaskTemplate.Constraints)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-tasktemplate-constraints.html
type TaskTemplate_Constraints struct {

	// InvisibleFields AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-tasktemplate-constraints.html#cfn-connect-tasktemplate-constraints-invisiblefields
	InvisibleFields []TaskTemplate_InvisibleFieldInfo `json:"InvisibleFields,omitempty"`

	// ReadOnlyFields AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-tasktemplate-constraints.html#cfn-connect-tasktemplate-constraints-readonlyfields
	ReadOnlyFields []TaskTemplate_ReadOnlyFieldInfo `json:"ReadOnlyFields,omitempty"`

	// RequiredFields AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-tasktemplate-constraints.html#cfn-connect-tasktemplate-constraints-requiredfields
	RequiredFields []TaskTemplate_RequiredFieldInfo `json:"RequiredFields,omitempty"`

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
func (r *TaskTemplate_Constraints) AWSCloudFormationType() string {
	return "AWS::Connect::TaskTemplate.Constraints"
}
