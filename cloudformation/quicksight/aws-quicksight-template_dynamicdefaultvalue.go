// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Template_DynamicDefaultValue AWS CloudFormation Resource (AWS::QuickSight::Template.DynamicDefaultValue)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-dynamicdefaultvalue.html
type Template_DynamicDefaultValue struct {

	// DefaultValueColumn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-dynamicdefaultvalue.html#cfn-quicksight-template-dynamicdefaultvalue-defaultvaluecolumn
	DefaultValueColumn *Template_ColumnIdentifier `json:"DefaultValueColumn"`

	// GroupNameColumn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-dynamicdefaultvalue.html#cfn-quicksight-template-dynamicdefaultvalue-groupnamecolumn
	GroupNameColumn *Template_ColumnIdentifier `json:"GroupNameColumn,omitempty"`

	// UserNameColumn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-dynamicdefaultvalue.html#cfn-quicksight-template-dynamicdefaultvalue-usernamecolumn
	UserNameColumn *Template_ColumnIdentifier `json:"UserNameColumn,omitempty"`

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
func (r *Template_DynamicDefaultValue) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.DynamicDefaultValue"
}
